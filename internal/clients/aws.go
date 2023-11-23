/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
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
	//errNoProviderConfig     = "no providerConfigRef provided"
	//errGetProviderConfig    = "cannot get referenced ProviderConfig"
	//errTrackUsage           = "cannot track ProviderConfig usage"
	//errExtractCredentials   = "cannot extract credentials"
	//errUnmarshalCredentials = "cannot unmarshal aws credentials as JSON"
	//errRegionNotFound       = "can not found region for terraform provider config"

	//accessKeyID     = "access_key"
	keyAccessKeyID = "access_key"
	//secretAccessKey = "secret_key"
	keySecretAccessKey = "secret_key"

	keyRegion       = "region"
	keySessionToken = "token"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string, mta *schema.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {

		pc := &v1beta1.ProviderConfig{}
		var err error
		if err = client.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
			return terraform.Setup{}, errors.Wrapf(err, "cannot get referenced Provider: %s", mg.GetProviderConfigReference().Name)
		}
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		err = pushDownTerraformSetupBuilder(ctx, client, mg, pc, &ps)
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

func pushDownTerraformSetupBuilder(ctx context.Context, c client.Client, mg resource.Managed, pc *v1beta1.ProviderConfig, ps *terraform.Setup) error {
	s := pc.Spec.Credentials.Source
	cfg, err := getAWSConfig(ctx, c, mg)
	if err != nil {
		return errors.Wrap(err, "cannot get AWS config")
	}
	ps.Configuration = map[string]any{
		keyRegion: cfg.Region,
	}
	data, err := resource.CommonCredentialExtractor(ctx, s, c, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return errors.Wrap(err, "cannot get credentials")
	}
	cfg, err = UseProviderSecret(ctx, data, DefaultSection, cfg.Region)
	if err != nil {
		return errors.Wrap(err, errAWSConfig)
	}
	creds, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to retrieve aws credentials from aws config")
	}
	ps.Configuration = map[string]any{
		keyRegion:          cfg.Region,
		keyAccessKeyID:     creds.AccessKeyID,
		keySecretAccessKey: creds.SecretAccessKey,
		keySessionToken:    creds.SessionToken,
	}
	return nil
}

func configureNoForkAWSClient(ctx context.Context, c client.Client, mg resource.Managed, pc *v1beta1.ProviderConfig, ps *terraform.Setup) (xpprovider.AWSConfig, error) { //nolint:gocyclo

	cfg, err := getAWSConfig(ctx, c, mg)
	if err != nil {
		return xpprovider.AWSConfig{}, errors.Wrap(err, "cannot get AWS config")
	}

	awsConfig := xpprovider.AWSConfig{
		Region:           cfg.Region,
		TerraformVersion: ps.Version,
	}
	s := pc.Spec.Credentials.Source
	data, err := resource.CommonCredentialExtractor(ctx, s, c, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return xpprovider.AWSConfig{}, errors.Wrap(err, "cannot get credentials")
	}
	cfg, err = UseProviderSecret(ctx, data, DefaultSection, cfg.Region)
	if err != nil {
		return xpprovider.AWSConfig{}, errors.Wrap(err, errAWSConfig)
	}
	creds, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		return xpprovider.AWSConfig{}, errors.Wrap(err, "failed to retrieve aws credentials from aws config")
	}

	awsConfig.Region = cfg.Region
	awsConfig.AccessKey = creds.AccessKeyID
	awsConfig.SecretKey = creds.SecretAccessKey
	awsConfig.Token = creds.SessionToken

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
