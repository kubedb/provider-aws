package ec2

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_route", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}

		r.UseAsync = true
		r.ShortGroup = "ec2"
	})

	p.AddResourceConfigurator("aws_subnet", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			// NOTE(muvaf): Conflicts with AvailabilityZone. See the following
			// for more details: https://github.com/crossplane/upjet/issues/107
			IgnoredFields: []string{
				"availability_zone_id",
			},
		}
		r.UseAsync = true
	})
}
