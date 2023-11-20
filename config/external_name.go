/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"fmt"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/pkg/errors"
	"kubedb.dev/provider-aws/config/common"
	"strings"
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

// FormattedIdentifierFromProvider is a helper function to construct Terraform
// IDs that use elements from the parameters in a certain string format.
// It should be used in cases where all information in the ID is gathered from
// the spec and not user defined like name. For example, zone_id:vpc_id.
func FormattedIdentifierFromProvider(separator string, keys ...string) config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		vals := make([]string, len(keys))
		for i, key := range keys {
			val, ok := parameters[key]
			if !ok {
				return "", errors.Errorf("%s cannot be empty", key)
			}
			s, ok := val.(string)
			if !ok {
				return "", errors.Errorf("%s needs to be string", key)
			}
			vals[i] = s
		}
		return strings.Join(vals, separator), nil
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

	"aws_security_group": config.IdentifierFromProvider,

	//vpc
	//
	// Imported using the id: vpc-23123
	"aws_vpc": config.IdentifierFromProvider,
	// Imported using the vpc endpoint id: vpce-3ecf2a57
	"aws_vpc_endpoint": config.IdentifierFromProvider,
	// Imported using the subnet id: subnet-9d4a7b6c
	"aws_subnet": config.IdentifierFromProvider,

	// docdb
	//
	// DocDB Clusters can be imported using the cluster_identifier
	"aws_docdb_cluster": config.ParameterAsIdentifier("cluster_identifier"),
	// aws_docdb_global_cluster can be imported by using the Global Cluster id
	"aws_docdb_global_cluster": config.IdentifierFromProvider,
	// DocDB Cluster Instances can be imported using the identifier
	"aws_docdb_cluster_instance": config.ParameterAsIdentifier("identifier"),
	// DocumentDB Subnet groups can be imported using the name
	"aws_docdb_subnet_group": config.NameAsIdentifier,
	// DocumentDB Cluster Parameter Groups can be imported using the name
	"aws_docdb_cluster_parameter_group": config.NameAsIdentifier,
	// aws_docdb_cluster_snapshot can be imported by using the cluster snapshot identifier
	"aws_docdb_cluster_snapshot": config.ParameterAsIdentifier("db_cluster_snapshot_identifier"),
	// DocDB Event Subscriptions can be imported using the name
	"aws_docdb_event_subscription": config.NameAsIdentifier,

	// dynamodb
	//
	// DynamoDB table replicas can be imported using the table-name:main-region
	"aws_dynamodb_table_replica": config.IdentifierFromProvider,
	// DynamoDB tables can be imported using the name
	"aws_dynamodb_table": config.NameAsIdentifier,
	// DynamoDB Global Tables can be imported using the global table name
	"aws_dynamodb_global_table": config.NameAsIdentifier,
	// aws_dynamodb_tag can be imported by using the DynamoDB resource identifier and key, separated by a comma (,)
	"aws_dynamodb_tag": config.TemplatedStringAsIdentifier("", "{{ .parameters.resource_arn }},{{ .parameters.key }}"),
	// DynamoDB Table Items can be imported using the name
	"aws_dynamodb_table_item": config.IdentifierFromProvider,
	// DynamoDB contributor insights
	"aws_dynamodb_contributor_insights": config.IdentifierFromProvider,
	// Dynamodb Kinesis streaming destinations are imported using "table_name,stream_arn"
	"aws_dynamodb_kinesis_streaming_destination": config.IdentifierFromProvider,

	// elasticache
	//
	"aws_elasticache_parameter_group":   config.IdentifierFromProvider,
	"aws_elasticache_subnet_group":      config.NameAsIdentifier,
	"aws_elasticache_cluster":           config.ParameterAsIdentifier("cluster_id"),
	"aws_elasticache_replication_group": config.ParameterAsIdentifier("replication_group_id"),
	"aws_elasticache_user":              config.ParameterAsIdentifier("user_id"),
	"aws_elasticache_user_group":        config.ParameterAsIdentifier("user_group_id"),

	// elasticsearch
	//
	// Elasticsearch domains can be imported using the domain_name
	"aws_elasticsearch_domain": config.TemplatedStringAsIdentifier("domain_name", "arn:aws:es:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:domain/{{ .external_name }}"),
	// No import
	"aws_elasticsearch_domain_policy": config.IdentifierFromProvider,
	// Elasticsearch domains can be imported using the domain_name
	"aws_elasticsearch_domain_saml_options": config.ParameterAsIdentifier("domain_name"),

	// kafka
	//
	// MSK configurations can be imported using the configuration ARN that has
	// a random substring in the end.
	"aws_msk_configuration": config.IdentifierFromProvider,
	// MSK clusters can be imported using the cluster arn that has a random substring
	// in the end.
	"aws_msk_cluster": config.IdentifierFromProvider,

	// memorydb
	//
	// Use the name to import a parameter group
	"aws_memorydb_parameter_group": config.NameAsIdentifier,
	// Use the name to import a subnet group
	"aws_memorydb_subnet_group": config.NameAsIdentifier,
	// Use the name to import a cluster
	"aws_memorydb_cluster": config.NameAsIdentifier,
	// Use the name to import an ACL
	"aws_memorydb_acl": config.NameAsIdentifier,
	// Use the name to import a snapshot
	"aws_memorydb_snapshot": config.NameAsIdentifier,

	// rds
	//
	"aws_rds_cluster":        config.ParameterAsIdentifier("cluster_identifier"),
	"aws_db_instance":        config.ParameterAsIdentifier("identifier"),
	"aws_db_parameter_group": config.NameAsIdentifier,
	"aws_db_subnet_group":    config.NameAsIdentifier,
	// aws_db_instance_role_association can be imported using the DB Instance Identifier and IAM Role ARN separated by a comma
	// $ terraform import aws_db_instance_role_association.example my-db-instance,arn:aws:iam::123456789012:role/my-role
	"aws_db_instance_role_association": config.IdentifierFromProvider,
	// DB Option groups can be imported using the name
	"aws_db_option_group": config.NameAsIdentifier,
	// DB proxies can be imported using the name
	"aws_db_proxy": config.NameAsIdentifier,
	// DB proxy default target groups can be imported using the db_proxy_name
	"aws_db_proxy_default_target_group": config.IdentifierFromProvider,
	// DB proxy endpoints can be imported using the DB-PROXY-NAME/DB-PROXY-ENDPOINT-NAME
	"aws_db_proxy_endpoint": config.TemplatedStringAsIdentifier("db_proxy_endpoint_name", "{{ .external_name }}/{{ .parameters.db_proxy_name }}"),
	// RDS DB Proxy Targets can be imported using the db_proxy_name, target_group_name, target type (e.g., RDS_INSTANCE or TRACKED_CLUSTER), and resource identifier separated by forward slashes (/)
	"aws_db_proxy_target": config.IdentifierFromProvider,
	// NOTE(turkenf): The resource aws_db_security_group is deprecated,
	// Please see: https://kubedb.dev/provider-aws/issues/696
	// aws_db_snapshot can be imported by using the snapshot identifier
	"aws_db_snapshot": config.ParameterAsIdentifier("db_snapshot_identifier"),
	// RDS Aurora Cluster Database Activity Streams can be imported using the resource_arn
	"aws_rds_cluster_activity_stream": config.IdentifierFromProvider,
	// RDS Clusters Endpoint can be imported using the cluster_endpoint_identifier
	"aws_rds_cluster_endpoint": config.ParameterAsIdentifier("cluster_endpoint_identifier"),
	// RDS Cluster Instances can be imported using the identifier
	"aws_rds_cluster_instance": config.ParameterAsIdentifier("identifier"),
	// RDS Cluster Parameter Groups can be imported using the name
	"aws_rds_cluster_parameter_group": config.NameAsIdentifier,
	// aws_rds_cluster_role_association can be imported using the DB Cluster Identifier and IAM Role ARN separated by a comma (,)
	// $ terraform import aws_rds_cluster_role_association.example my-db-cluster,arn:aws:iam::123456789012:role/my-role
	"aws_rds_cluster_role_association": FormattedIdentifierFromProvider(",", "db_cluster_identifier", "role_arn"),
	// aws_rds_global_cluster can be imported by using the RDS Global Cluster identifie
	"aws_rds_global_cluster": config.ParameterAsIdentifier("global_cluster_identifier"),
	// aws_db_cluster_snapshot can be imported by using the cluster snapshot identifier
	"aws_db_cluster_snapshot": config.IdentifierFromProvider,
	// DB Event Subscriptions can be imported using the name
	"aws_db_event_subscription": config.NameAsIdentifier,
	// RDS instance automated backups replication can be imported using the arn
	"aws_db_instance_automated_backups_replication": config.IdentifierFromProvider,
	// aws_db_snapshot_copy can be imported by using the snapshot identifier
	"aws_db_snapshot_copy": config.IdentifierFromProvider,

	// kms
	//
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	"aws_kms_key": config.IdentifierFromProvider,

	// kinesis
	//
	// Even though the documentation says the ID is name, it uses ARN..
	"aws_kinesis_stream": config.TemplatedStringAsIdentifier("name", " arn:aws:kinesis:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:stream/{{ .external_name }}"),
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1alpha1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.Version = common.VersionV1Alpha1
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
