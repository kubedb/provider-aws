package kinesis

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for kinesis group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_kinesis_stream", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"enforce_consumer_deletion"},
		}
		config.MoveToStatus(r.TerraformResource, "arn")
	})
}
