/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	ujconfig "github.com/upbound/upjet/pkg/config"
	"kubedb.dev/provider-aws/config/docdb"
	"kubedb.dev/provider-aws/config/dynamodb"
	"kubedb.dev/provider-aws/config/ec2"
	"kubedb.dev/provider-aws/config/elasticache"
	"kubedb.dev/provider-aws/config/kafka"
	"kubedb.dev/provider-aws/config/kinesis"
	"kubedb.dev/provider-aws/config/memorydb"
	"kubedb.dev/provider-aws/config/rds"
)

const (
	resourcePrefix = "aws"
	modulePath     = "kubedb.dev/provider-aws"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithRootGroup("aws.kubedb.com"),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			KindOverrides(),
			RegionAddition(),
			TagsAllRemoval(),
			IdentifierAssignedByAWS(),
			KnownReferencers(),
			AddExternalTagsField(),
			ExternalNameConfigurations(),
			NamePrefixRemoval(),
			DocumentationForTags(),
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
