/*
Copyright 2021 Upbound Inc.
*/

package iam

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the iam group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_iam_access_key", func(r *config.Resource) {
		r.References = config.References{
			"user": config.Reference{
				Type: "User",
			},
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["id"].(string); ok {
				conn["username"] = []byte(a)
			}
			if a, ok := attr["secret"].(string); ok {
				conn["password"] = []byte(a)
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("aws_iam_role", func(r *config.Resource) {
		r.MetaResource.ArgumentDocs["inline_policy"] = `Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. See below. If no blocks are configured, Crossplane will not manage any inline policies in this resource. Configuring one empty block (i.e., inline_policy {}) will cause Crossplane to remove all inline policies added out of band on apply.`
		r.MetaResource.ArgumentDocs["managed_policy_arns"] = `Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, Crossplane will ignore policy attachments to this resource. When configured, Crossplane will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., managed_policy_arns = []) will cause Crossplane to remove all managed policy attachments.`
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"managed_policy_arns"},
		}
	})
}
