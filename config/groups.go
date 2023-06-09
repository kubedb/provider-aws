package config

import (
	"github.com/upbound/upjet/pkg/config"
)

var KindMap = map[string]string{
	"aws_security_group_rule":    "SecurityGroupRule",
	"aws_vpc_peering_connection": "VPCPeeringConnection",
}

// KindOverrides overrides the kind of the resources given in KindMap.
func KindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if k, ok := KindMap[r.Name]; ok {
			r.Kind = k
		}
	}
}
