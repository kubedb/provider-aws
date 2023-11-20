package kafka

import (
	"github.com/crossplane/upjet/pkg/config"

	"kubedb.dev/provider-aws/config/common"
)

// Configure adds configurations for kafka group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_msk_cluster", func(r *config.Resource) {
		r.References["encryption_info.encryption_at_rest_kms_key_arn"] = config.Reference{
			Type:      "kubedb.dev/provider-aws/apis/kms/v1alpha1.Key",
			Extractor: common.PathARNExtractor,
		}
		/*r.References["logging_info.broker_logs.s3.bucket"] = config.Reference{
			Type: "kubedb.dev/provider-aws/apis/s3/v1alpha1.Bucket",
		}*/
		/*r.References["logging_info.broker_logs.cloudwatch_logs.log_group"] = config.Reference{
			Type: "kubedb.dev/provider-aws/apis/cloudwatchlogs/v1alpha1.Group",
		}*/
		r.References["broker_node_group_info.client_subnets"] = config.Reference{
			Type: "kubedb.dev/provider-aws/apis/ec2/v1alpha1.Subnet",
		}
		r.References["broker_node_group_info.security_groups"] = config.Reference{
			Type: "kubedb.dev/provider-aws/apis/ec2/v1alpha1.SecurityGroup",
		}
		r.UseAsync = true
	})
}
