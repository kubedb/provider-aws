/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/appscode/provider-aws/config/peeringconnection"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "aws"
	modulePath     = "github.com/appscode/provider-aws"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			RegionAddition(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		peeringconnection.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
