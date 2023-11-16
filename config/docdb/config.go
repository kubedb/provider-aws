package docdb

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for docdb group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_docdb_cluster", func(r *config.Resource) {
		config.MoveToStatus(r.TerraformResource, "cluster_members")
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_docdb_cluster_instance", func(r *config.Resource) {
		r.References["cluster_identifier"] = config.Reference{
			Type: "Cluster",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_docdb_subnet_group", func(r *config.Resource) {
		r.References["subnet_ids"] = config.Reference{
			Type: "kubedb.dev/provider-aws/apis/ec2/v1alpha1.Subnet",
		}
	})
}
