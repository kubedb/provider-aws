/*
Copyright 2022 Upbound Inc.
*/

package clients

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	stscredstypesv2 "github.com/aws/aws-sdk-go-v2/service/sts/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/go-ini/ini"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"kubedb.dev/provider-aws/apis/v1beta1"
	"kubedb.dev/provider-aws/internal/version"
)

const (
	// DefaultSection for INI files.
	DefaultSection = "DEFAULT"
	//ff=ini.DefaultSection

	// authentication types
	errRoleChainConfig = "failed to load assumed role AWS config"
	errAWSConfig       = "failed to get AWS config"
)

// GlobalRegion is the region name used for AWS services that do not have a notion
// of region.
const GlobalRegion = "aws-global"

// Endpoint URL configuration types.
const (
	URLConfigTypeStatic  = "Static"
	URLConfigTypeDynamic = "Dynamic"
)

// userAgentV2 constructs the Crossplane user agent for AWS v2 clients
var userAgentV2 = config.WithAPIOptions([]func(*middleware.Stack) error{
	awsmiddleware.AddUserAgentKeyValue("upbound-provider-aws", version.Version),
	awsmiddleware.AddUserAgentKeyValue("crossplane-provider-aws", version.Version),
})

// GetAWSConfig to produce a config that can be used to authenticate to AWS.
func GetAWSConfig(ctx context.Context, c client.Client, mg resource.Managed) (*aws.Config, error) { // nolint:gocyclo
	if mg.GetProviderConfigReference() == nil {
		return nil, errors.New("no providerConfigRef provided")
	}
	region, err := getRegion(mg)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get region")
	}
	pc := &v1beta1.ProviderConfig{}
	if err := c.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
		return nil, errors.Wrap(err, "cannot get referenced Provider")
	}

	t := resource.NewProviderConfigUsageTracker(c, &v1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, "cannot track ProviderConfig usage")
	}

	var cfg *aws.Config
	s := pc.Spec.Credentials.Source //nolint:exhaustive

	data, err := resource.CommonCredentialExtractor(ctx, s, c, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get credentials")
	}
	cfg, err = UseProviderSecret(ctx, data, DefaultSection, region)
	if err != nil {
		return nil, errors.Wrap(err, errAWSConfig)
	}

	cfg, err = GetRoleChainConfig(ctx, &pc.Spec, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get credentials")
	}
	return SetResolver(pc, cfg), nil
}

type awsEndpointResolverAdaptorWithOptions func(service, region string, options interface{}) (aws.Endpoint, error)

func (a awsEndpointResolverAdaptorWithOptions) ResolveEndpoint(service, region string, options ...interface{}) (aws.Endpoint, error) {
	return a(service, region, options)
}

// SetResolver parses annotations from the managed resource
// and returns a configuration accordingly.
func SetResolver(pc *v1beta1.ProviderConfig, cfg *aws.Config) *aws.Config { // nolint:gocyclo
	if pc.Spec.Endpoint == nil {
		return cfg
	}
	cfg.EndpointResolverWithOptions = awsEndpointResolverAdaptorWithOptions(func(service, region string, options interface{}) (aws.Endpoint, error) { //nolint:staticcheck
		fullURL := ""
		switch pc.Spec.Endpoint.URL.Type {
		case URLConfigTypeStatic:
			if pc.Spec.Endpoint.URL.Static == nil {
				return aws.Endpoint{}, errors.New("static type is chosen but static field does not have a value")
			}
			fullURL = aws.ToString(pc.Spec.Endpoint.URL.Static)
		case URLConfigTypeDynamic:
			if pc.Spec.Endpoint.URL.Dynamic == nil {
				return aws.Endpoint{}, errors.New("dynamic type is chosen but dynamic configuration is not given")
			}
			// NOTE(muvaf): IAM does not have any region.
			if service == "IAM" {
				fullURL = fmt.Sprintf("%s://%s.%s", pc.Spec.Endpoint.URL.Dynamic.Protocol, strings.ToLower(service), pc.Spec.Endpoint.URL.Dynamic.Host)
			} else {
				fullURL = fmt.Sprintf("%s://%s.%s.%s", pc.Spec.Endpoint.URL.Dynamic.Protocol, strings.ToLower(service), region, pc.Spec.Endpoint.URL.Dynamic.Host)
			}
		default:
			return aws.Endpoint{}, errors.New("unsupported url config type is chosen")
		}
		e := aws.Endpoint{
			URL:               fullURL,
			HostnameImmutable: aws.ToBool(pc.Spec.Endpoint.HostnameImmutable),
			PartitionID:       aws.ToString(pc.Spec.Endpoint.PartitionID),
			SigningName:       aws.ToString(pc.Spec.Endpoint.SigningName),
			SigningRegion:     aws.ToString(LateInitializeStringPtr(pc.Spec.Endpoint.SigningRegion, &region)),
			SigningMethod:     aws.ToString(pc.Spec.Endpoint.SigningMethod),
		}
		// Only IAM does not have a region parameter and "aws-global" is used in
		// SDK setup. However, signing region has to be us-east-1 and it needs
		// to be set.
		if region == "aws-global" {
			switch aws.ToString(pc.Spec.Endpoint.PartitionID) {
			case "aws-us-gov", "aws-cn", "aws-iso", "aws-iso-b":
				e.SigningRegion = aws.ToString(LateInitializeStringPtr(pc.Spec.Endpoint.SigningRegion, &region))
			default:
				e.SigningRegion = "us-east-1"
			}
		}
		if pc.Spec.Endpoint.Source != nil {
			switch *pc.Spec.Endpoint.Source {
			case "ServiceMetadata":
				e.Source = aws.EndpointSourceServiceMetadata
			case "Custom":
				e.Source = aws.EndpointSourceCustom
			}
		}
		return e, nil
	})
	return cfg
}

// CredentialsIDSecret retrieves AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY from the data which contains
// aws credentials under given profile
// Example:
// [default]
// aws_access_key_id = <YOUR_ACCESS_KEY_ID>
// aws_secret_access_key = <YOUR_SECRET_ACCESS_KEY>
func CredentialsIDSecret(data []byte, profile string) (aws.Credentials, error) {
	awsConfig, err := ini.InsensitiveLoad(data)
	if err != nil {
		return aws.Credentials{}, errors.Wrap(err, "cannot parse credentials secret")
	}

	iniProfile, err := awsConfig.GetSection(profile)
	if err != nil {
		return aws.Credentials{}, errors.Wrap(err, fmt.Sprintf("cannot get %s profile in credentials secret", profile))
	}

	accessKeyID := iniProfile.Key("aws_access_key_id")
	//accessKeyID := iniProfile.Key("access_key")
	secretAccessKey := iniProfile.Key("aws_secret_access_key")
	//secretAccessKey := iniProfile.Key("secret_key")
	//sessionToken := iniProfile.Key("aws_session_token")

	// NOTE(muvaf): Key function implementation never returns nil but still its
	// type is pointer so we check to make sure its next versions doesn't break
	// that implicit contract.
	if accessKeyID == nil || secretAccessKey == nil {
		return aws.Credentials{}, errors.New("returned key can be empty but cannot be nil")
	}

	return aws.Credentials{
		AccessKeyID:     accessKeyID.Value(),
		SecretAccessKey: secretAccessKey.Value(),
		//SessionToken:    sessionToken.Value(),
	}, nil
}

// AuthMethod is a method of authenticating to the AWS API
type AuthMethod func(context.Context, []byte, string, string) (*aws.Config, error)

// stsRegionOrDefault sets the STS client region to the passed region, or
// defaults to the global region.
func stsRegionOrDefault(region string) func(*sts.Options) {
	return func(o *sts.Options) {
		if region == "" {
			o.Region = GlobalRegion
		}
	}
}

// UseProviderSecret - AWS configuration which can be used to issue requests against AWS API
func UseProviderSecret(ctx context.Context, data []byte, profile, region string) (*aws.Config, error) {
	creds, err := CredentialsIDSecret(data, profile)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse credentials secret")
	}

	awsConfig, err := config.LoadDefaultConfig(
		ctx,
		userAgentV2,
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: creds,
		}),
	)
	if err != nil {
		return nil, errors.Wrap(err, "cannot load default AWS config")
	}
	return &awsConfig, nil
}

// GetRoleChainConfig returns an aws.Config capable of doing role chaining with
// AssumeRoleWithWebIdentity & AssumeRoles.
func GetRoleChainConfig(ctx context.Context, pcs *v1beta1.ProviderConfigSpec, cfg *aws.Config) (*aws.Config, error) {
	pCfg := cfg
	for _, aro := range pcs.AssumeRoleChain {
		stsAssume := stscreds.NewAssumeRoleProvider(
			sts.NewFromConfig(*pCfg, stsRegionOrDefault(cfg.Region)), //nolint:contextcheck
			aws.ToString(aro.RoleARN),
			SetAssumeRoleOptions(aro),
		)
		cfgWithAssumeRole, err := config.LoadDefaultConfig(
			ctx,
			userAgentV2,
			config.WithRegion(cfg.Region),
			config.WithCredentialsProvider(aws.NewCredentialsCache(stsAssume)),
		)
		if err != nil {
			return nil, errors.Wrap(err, errRoleChainConfig)
		}
		pCfg = &cfgWithAssumeRole
	}
	return pCfg, nil
}

// SetAssumeRoleOptions sets options when Assuming an IAM Role
func SetAssumeRoleOptions(aro v1beta1.AssumeRoleOptions) func(*stscreds.AssumeRoleOptions) {
	return func(opt *stscreds.AssumeRoleOptions) {
		opt.ExternalID = aro.ExternalID
		for _, t := range aro.Tags {
			opt.Tags = append(
				opt.Tags,
				stscredstypesv2.Tag{
					Key:   t.Key,
					Value: t.Value,
				})
		}
		opt.TransitiveTagKeys = append(opt.TransitiveTagKeys, aro.TransitiveTagKeys...)
	}
}

// LateInitializeStringPtr returns in if it's non-nil, otherwise returns from
// which is the backup for the cases in is nil.
func LateInitializeStringPtr(in *string, from *string) *string {
	if in != nil {
		return in
	}
	return from
}
