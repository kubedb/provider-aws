/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	ujconfig "github.com/upbound/upjet/pkg/config"
	"kubeform.dev/provider-aws/config/docdb"
	"kubeform.dev/provider-aws/config/dynamodb"
	"kubeform.dev/provider-aws/config/ec2"
	"kubeform.dev/provider-aws/config/elasticache"
	"kubeform.dev/provider-aws/config/kafka"
	"kubeform.dev/provider-aws/config/kinesis"
	"kubeform.dev/provider-aws/config/memorydb"
	"kubeform.dev/provider-aws/config/rds"
)

const (
	resourcePrefix = "aws"
	modulePath     = "kubeform.dev/provider-aws"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithRootGroup("aws.kubeform.com"),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			RegionAddition(),
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
	} {
		configure(pc)
	}
	pc.ConfigureResources()

	return pc
}
