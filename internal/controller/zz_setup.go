/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	providerconfig "github.com/appscode/provider-aws/internal/controller/providerconfig"
	grouprule "github.com/appscode/provider-aws/internal/controller/vpc/grouprule"
	peeringconnection "github.com/appscode/provider-aws/internal/controller/vpc/peeringconnection"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		grouprule.Setup,
		peeringconnection.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
