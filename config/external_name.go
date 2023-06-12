/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"
)

func route() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		rtb, ok := parameters["route_table_id"]
		if !ok {
			return "", errors.New("route_table_id cannot be empty")
		}
		switch {
		case parameters["destination_cidr_block"] != nil:
			return fmt.Sprintf("%s_%s", rtb.(string), parameters["destination_cidr_block"].(string)), nil
		case parameters["destination_ipv6_cidr_block"] != nil:
			return fmt.Sprintf("%s_%s", rtb.(string), parameters["destination_ipv6_cidr_block"].(string)), nil
		case parameters["destination_prefix_list_id"] != nil:
			return fmt.Sprintf("%s_%s", rtb.(string), parameters["destination_prefix_list_id"].(string)), nil
		}
		return "", errors.New("destination_cidr_block or destination_ipv6_cidr_block or destination_prefix_list_id has to be given")
	}
	return e
}

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"aws_vpc_peering_connection": config.IdentifierFromProvider,
	"aws_security_group_rule":    config.IdentifierFromProvider,
	"aws_route":                  route(),
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
