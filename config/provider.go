/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/kubeform/provider-aws/config/peeringconnection"
	"github.com/kubeform/provider-aws/config/routetable"
	"github.com/kubeform/provider-aws/config/security"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "aws"
	modulePath     = "github.com/kubeform/provider-aws"
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
			KindOverrides(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		peeringconnection.Configure,
		security.Configure,
		routetable.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
