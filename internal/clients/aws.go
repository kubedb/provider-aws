/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/xpprovider"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"unsafe"

	"github.com/crossplane/upjet/pkg/terraform"

	"kubedb.dev/provider-aws/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal aws credentials as JSON"
	errRegionNotFound       = "can not found region for terraform provider config"

	accessKeyID     = "access_key"
	secretAccessKey = "secret_key"
	keyRegion       = "region"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string, mta *schema.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		err := pushDownTerraformSetupBuilder(ctx, client, mg, pc, &ps)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot build terraform configuration")
		}

		awsConfig, err := configureNoForkAWSClient(ctx, client, mg, pc, &ps)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "could not configure no-fork AWS client")
		}
		p := mta.Meta()
		tfClient, diag := awsConfig.GetClient(ctx, &xpprovider.AWSClient{
			// #nosec G103
			ServicePackages: (*xpprovider.AWSClient)(unsafe.Pointer(reflect.ValueOf(p).Pointer())).ServicePackages,
		})
		if diag != nil && diag.HasError() {
			return terraform.Setup{}, errors.Errorf("failed to configure the AWS client: %v", diag)
		}
		ps.Meta = tfClient
		return ps, nil
	}
}

func pushDownTerraformSetupBuilder(ctx context.Context, client client.Client, mg resource.Managed, pc *v1beta1.ProviderConfig, ps *terraform.Setup) error {
	data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return errors.Wrap(err, errExtractCredentials)
	}
	creds := map[string]string{}
	if err := json.Unmarshal(data, &creds); err != nil {
		return errors.Wrap(err, errUnmarshalCredentials)
	}

	region, err := getRegion(mg)
	if err != nil {
		return errors.Wrap(err, errRegionNotFound)
	}

	// Set credentials in Terraform provider configuration.
	ps.Configuration = map[string]any{
		accessKeyID:     creds[accessKeyID],
		secretAccessKey: creds[secretAccessKey],
		keyRegion:       region,
	}
	return nil
}

func configureNoForkAWSClient(ctx context.Context, c client.Client, mg resource.Managed, pc *v1beta1.ProviderConfig, ps *terraform.Setup) (xpprovider.AWSConfig, error) { //nolint:gocyclo

	if len(pc.Spec.AssumeRoleChain) > 1 || pc.Spec.Endpoint != nil {
		return xpprovider.AWSConfig{}, errors.New("cannot configure no-fork client because the length of assume role chain array " +
			"is more than 1 or endpoint configuration is not nil")
	}

	cfg, err := getAWSConfig(ctx, c, mg)
	if err != nil {
		return xpprovider.AWSConfig{}, errors.Wrap(err, "cannot get AWS config")
	}

	awsConfig := xpprovider.AWSConfig{
		Region:           cfg.Region,
		TerraformVersion: ps.Version,
	}

	data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, c, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return xpprovider.AWSConfig{}, errors.Wrap(err, errExtractCredentials)
	}
	creds := map[string]string{}
	if err := json.Unmarshal(data, &creds); err != nil {
		return xpprovider.AWSConfig{}, errors.Wrap(err, errUnmarshalCredentials)
	}

	region, err := getRegion(mg)
	if err != nil {
		return xpprovider.AWSConfig{}, errors.Wrap(err, errRegionNotFound)
	}

	// Set credentials in Terraform provider configuration.

	awsConfig.AccessKey = creds[accessKeyID]
	awsConfig.SecretKey = creds[secretAccessKey]
	awsConfig.Region = region

	return awsConfig, nil
}

func getAWSConfig(ctx context.Context, c client.Client, mg resource.Managed) (*aws.Config, error) {
	cfg, err := GetAWSConfig(ctx, c, mg)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get AWS config")
	}
	if cfg.Region == "" && mg.GetObjectKind().GroupVersionKind().Group == "iam.aws.upbound.io" {
		cfg.Region = "us-east-1"
	}
	return cfg, nil
}

func getRegion(obj runtime.Object) (string, error) {
	fromMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return "", errors.Wrap(err, "cannot convert to unstructured")
	}
	r, err := fieldpath.Pave(fromMap).GetString("spec.forProvider.region")
	if fieldpath.IsNotFound(err) {
		// Region is not required for all resources, e.g. resource in "iam"
		// group.
		return "", nil
	}
	return r, err
}
