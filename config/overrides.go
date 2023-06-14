package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/types/comments"
	"strings"
)

var (
	resourceGroup = map[string]string{
		"aws_security_group_rule":    "ec2",
		"aws_vpc_peering_connection": "ec2",
		"aws_rote":                   "ec2",

		"aws_docdb_cluster":                 "docdb",
		"aws_docdb_global_cluster":          "docdb",
		"aws_docdb_cluster_instance":        "docdb",
		"aws_docdb_subnet_group":            "docdb",
		"aws_docdb_cluster_parameter_group": "docdb",
		"aws_docdb_cluster_snapshot":        "docdb",
		"aws_docdb_event_subscription":      "docdb",

		"aws_dynamodb_table_replica":                 "dynamodb",
		"aws_dynamodb_table":                         "dynamodb",
		"aws_dynamodb_global_table":                  "dynamodb",
		"aws_dynamodb_tag":                           "dynamodb",
		"aws_dynamodb_table_item":                    "dynamodb",
		"aws_dynamodb_contributor_insights":          "dynamodb",
		"aws_dynamodb_kinesis_streaming_destination": "dynamodb",

		"aws_elasticache_parameter_group":   "elasticache",
		"aws_elasticache_subnet_group":      "elasticache",
		"aws_elasticache_cluster":           "elasticache",
		"aws_elasticache_replication_group": "elasticache",
		"aws_elasticache_user":              "elasticache",
		"aws_elasticache_user_group":        "elasticache",

		"aws_elasticsearch_domain":              "elasticsearch",
		"aws_elasticsearch_domain_policy":       "elasticsearch",
		"aws_elasticsearch_domain_saml_options": "elasticsearch",

		"aws_msk_configuration": "kafka",
		"aws_msk_cluster":       "kafka",

		"aws_memorydb_parameter_group": "memorydb",
		"aws_memorydb_subnet_group":    "memorydb",
		"aws_memorydb_cluster":         "memorydb",
		"aws_memorydb_acl":             "memorydb",
		"aws_memorydb_snapshot":        "memorydb",

		"aws_rds_cluster":                               "rds",
		"aws_db_instance":                               "rds",
		"aws_db_parameter_group":                        "rds",
		"aws_db_subnet_group":                           "rds",
		"aws_db_instance_role_association":              "rds",
		"aws_db_option_group":                           "rds",
		"aws_db_proxy":                                  "rds",
		"aws_db_proxy_default_target_group":             "rds",
		"aws_db_proxy_endpoint":                         "rds",
		"aws_db_proxy_target":                           "rds",
		"aws_db_snapshot":                               "rds",
		"aws_rds_cluster_activity_stream":               "rds",
		"aws_rds_cluster_endpoint":                      "rds",
		"aws_rds_cluster_instance":                      "rds",
		"aws_rds_cluster_parameter_group":               "rds",
		"aws_rds_cluster_role_association":              "rds",
		"aws_rds_global_cluster":                        "rds",
		"aws_db_cluster_snapshot":                       "rds",
		"aws_db_event_subscription":                     "rds",
		"aws_db_instance_automated_backups_replication": "rds",
		"aws_db_snapshot_copy":                          "rds",
		"aws_kms_key":                                   "kms",
		"aws_kinesis_stream":                            "kinesis",
	}

	resourceKind = map[string]string{
		"aws_security_group_rule":    "SecurityGroupRule",
		"aws_vpc_peering_connection": "VPCPeeringConnection",
		"aws_rote":                   "Route",

		"aws_docdb_cluster":                 "Cluster",
		"aws_docdb_global_cluster":          "GlobalCluster",
		"aws_docdb_cluster_instance":        "ClusterInstance",
		"aws_docdb_subnet_group":            "SubnetGroup",
		"aws_docdb_cluster_parameter_group": "ClusterParameterGroup",
		"aws_docdb_cluster_snapshot":        "ClusterSnapshot",
		"aws_docdb_event_subscription":      "EventSubscription",

		"aws_dynamodb_table_replica":                 "TableReplica",
		"aws_dynamodb_table":                         "Table",
		"aws_dynamodb_global_table":                  "GlobalTable",
		"aws_dynamodb_tag":                           "Tag",
		"aws_dynamodb_table_item":                    "TableItem",
		"aws_dynamodb_contributor_insights":          "ContributorInsights",
		"aws_dynamodb_kinesis_streaming_destination": "KinesisStreamingDestination",

		"aws_elasticache_parameter_group":   "ParameterGroup",
		"aws_elasticache_subnet_group":      "SubnetGroup",
		"aws_elasticache_cluster":           "Cluster",
		"aws_elasticache_replication_group": "ReplicationGroup",
		"aws_elasticache_user":              "User",
		"aws_elasticache_user_group":        "UserGroup",

		"aws_elasticsearch_domain":              "Domain",
		"aws_elasticsearch_domain_policy":       "DomainPolicy",
		"aws_elasticsearch_domain_saml_options": "DomainSAMLOptions",

		"aws_msk_configuration": "Configuration",
		"aws_msk_cluster":       "Cluster",

		"aws_memorydb_parameter_group": "ParameterGroup",
		"aws_memorydb_subnet_group":    "SubnetGroup",
		"aws_memorydb_cluster":         "Cluster",
		"aws_memorydb_acl":             "ACL",
		"aws_memorydb_snapshot":        "Snapshot",

		"aws_rds_cluster":                               "Cluster",
		"aws_db_instance":                               "Instance",
		"aws_db_parameter_group":                        "ParameterGroup",
		"aws_db_subnet_group":                           "SubnetGroup",
		"aws_db_instance_role_association":              "InstanceRoleAssociation",
		"aws_db_option_group":                           "OptionGroup",
		"aws_db_proxy":                                  "Proxy",
		"aws_db_proxy_default_target_group":             "ProxyDefaultTargetGroup",
		"aws_db_proxy_endpoint":                         "ProxyEndpoint",
		"aws_db_proxy_target":                           "ProxyTarget",
		"aws_db_snapshot":                               "Snapshot",
		"aws_rds_cluster_activity_stream":               "ClusterActivityStream",
		"aws_rds_cluster_endpoint":                      "ClusterEndpoint",
		"aws_rds_cluster_instance":                      "ClusterInstance",
		"aws_rds_cluster_parameter_group":               "ClusterParameterGroup",
		"aws_rds_cluster_role_association":              "ClusterRoleAssociation",
		"aws_rds_global_cluster":                        "GlobalCluster",
		"aws_db_cluster_snapshot":                       "ClusterSnapshot",
		"aws_db_event_subscription":                     "EventSubscription",
		"aws_db_instance_automated_backups_replication": "DBInstanceAutomatedBackupsReplication",
		"aws_db_snapshot_copy":                          "DBSnapshotCopy",
		"aws_kms_key":                                   "Key",
		"aws_kinesis_stream":                            "Stream",
	}
)

func groupKindOverride(r *config.Resource) {
	if _, ok := resourceGroup[r.Name]; ok {
		r.ShortGroup = resourceGroup[r.Name]
	}

	if _, ok := resourceKind[r.Name]; ok {
		r.Kind = resourceKind[r.Name]
	}
}

// RegionAddition adds region to the spec of all resources except iam group which
// does not have a region notion.
func RegionAddition() config.ResourceOption {
	return func(r *config.Resource) {
		if r.ShortGroup == "iam" || r.ShortGroup == "opsworks" {
			return
		}
		c := "Region is the region you'd like your resource to be created in.\n"
		comment, err := comments.New(c, comments.WithTFTag("-"))
		if err != nil {
			panic(errors.Wrap(err, "cannot build comment for region"))
		}
		r.TerraformResource.Schema["region"] = &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: comment.String(),
		}
		if r.MetaResource == nil {
			return
		}
		for _, ex := range r.MetaResource.Examples {
			defaultRegion := "us-west-1"
			if err := ex.SetPathValue("region", defaultRegion); err != nil {
				panic(err)
			}
			for k := range ex.Dependencies {
				if strings.HasPrefix(k, "aws_iam") {
					continue
				}
				if err := ex.Dependencies.SetPathValue(k, "region", defaultRegion); err != nil {
					panic(err)
				}
			}
		}
	}
}
