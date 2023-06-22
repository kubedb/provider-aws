package ec2

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_route", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}

		r.UseAsync = true
		r.ShortGroup = "ec2"
	})
}
