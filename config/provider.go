/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"context"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry/reference"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	"github.com/hashicorp/terraform-provider-aws/xpprovider"
	"kubedb.dev/provider-aws/config/docdb"
	"kubedb.dev/provider-aws/config/dynamodb"
	"kubedb.dev/provider-aws/config/ec2"
	"kubedb.dev/provider-aws/config/elasticache"
	"kubedb.dev/provider-aws/config/iam"
	"kubedb.dev/provider-aws/config/kafka"
	"kubedb.dev/provider-aws/config/kinesis"
	"kubedb.dev/provider-aws/config/memorydb"
	"kubedb.dev/provider-aws/config/rds"
	"kubedb.dev/provider-aws/config/secretsmanager"
	"kubedb.dev/provider-aws/config/sns"
)

const (
	resourcePrefix = "aws"
	modulePath     = "kubedb.dev/provider-aws"
)

var (
	//go:embed schema.json
	providerSchema string

	//go:embed provider-metadata.yaml
	providerMetadata string
)

// workaround for the TF AWS v4.67.0-based no-fork release: We would like to
// keep the types in the generated CRDs intact
// (prevent number->int type replacements).
func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(ctx context.Context, generationProvider bool) (*ujconfig.Provider, error) {
	var p *schema.Provider
	var err error
	if generationProvider {
		p, err = getProviderSchema(providerSchema)
	} else {
		p, err = xpprovider.GetProviderSchema(ctx)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(CLIReconciledResourceList()),
		ujconfig.WithNoForkIncludeList(NoForkResourceList()),
		ujconfig.WithRootGroup("aws.kubedb.com"),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithTerraformProvider(p),
		ujconfig.WithDefaultResourceOptions(
			RegionAddition(),
			KnownReferencers(),
			ExternalNameConfigurations(),
		))

	// API group overrides from Terraform import statements
	for _, r := range pc.Resources {
		groupKindOverride(r)
	}

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		ec2.Configure,
		docdb.Configure,
		dynamodb.Configure,
		elasticache.Configure,
		kafka.Configure,
		memorydb.Configure,
		rds.Configure,
		kinesis.Configure,
		iam.Configure,
		secretsmanager.Configure,
		sns.Configure,
	} {
		configure(pc)
	}
	pc.ConfigureResources()

	return pc, nil
}
