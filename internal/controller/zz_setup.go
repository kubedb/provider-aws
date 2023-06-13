/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	route "kubeform.dev/provider-aws/internal/controller/ec2/route"
	providerconfig "kubeform.dev/provider-aws/internal/controller/providerconfig"
	securitygrouprule "kubeform.dev/provider-aws/internal/controller/vpc/securitygrouprule"
	vpcpeeringconnection "kubeform.dev/provider-aws/internal/controller/vpc/vpcpeeringconnection"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		route.Setup,
		providerconfig.Setup,
		securitygrouprule.Setup,
		vpcpeeringconnection.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
