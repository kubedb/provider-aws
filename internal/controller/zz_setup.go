/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	cluster "kubeform.dev/provider-aws/internal/controller/docdb/cluster"
	clusterinstance "kubeform.dev/provider-aws/internal/controller/docdb/clusterinstance"
	clusterparametergroup "kubeform.dev/provider-aws/internal/controller/docdb/clusterparametergroup"
	clustersnapshot "kubeform.dev/provider-aws/internal/controller/docdb/clustersnapshot"
	eventsubscription "kubeform.dev/provider-aws/internal/controller/docdb/eventsubscription"
	globalcluster "kubeform.dev/provider-aws/internal/controller/docdb/globalcluster"
	subnetgroup "kubeform.dev/provider-aws/internal/controller/docdb/subnetgroup"
	contributorinsights "kubeform.dev/provider-aws/internal/controller/dynamodb/contributorinsights"
	globaltable "kubeform.dev/provider-aws/internal/controller/dynamodb/globaltable"
	kinesisstreamingdestination "kubeform.dev/provider-aws/internal/controller/dynamodb/kinesisstreamingdestination"
	table "kubeform.dev/provider-aws/internal/controller/dynamodb/table"
	tableitem "kubeform.dev/provider-aws/internal/controller/dynamodb/tableitem"
	tablereplica "kubeform.dev/provider-aws/internal/controller/dynamodb/tablereplica"
	tag "kubeform.dev/provider-aws/internal/controller/dynamodb/tag"
	route "kubeform.dev/provider-aws/internal/controller/ec2/route"
	securitygrouprule "kubeform.dev/provider-aws/internal/controller/ec2/securitygrouprule"
	vpcpeeringconnection "kubeform.dev/provider-aws/internal/controller/ec2/vpcpeeringconnection"
	clusterelasticache "kubeform.dev/provider-aws/internal/controller/elasticache/cluster"
	parametergroup "kubeform.dev/provider-aws/internal/controller/elasticache/parametergroup"
	replicationgroup "kubeform.dev/provider-aws/internal/controller/elasticache/replicationgroup"
	subnetgroupelasticache "kubeform.dev/provider-aws/internal/controller/elasticache/subnetgroup"
	user "kubeform.dev/provider-aws/internal/controller/elasticache/user"
	usergroup "kubeform.dev/provider-aws/internal/controller/elasticache/usergroup"
	domain "kubeform.dev/provider-aws/internal/controller/elasticsearch/domain"
	domainpolicy "kubeform.dev/provider-aws/internal/controller/elasticsearch/domainpolicy"
	domainsamloptions "kubeform.dev/provider-aws/internal/controller/elasticsearch/domainsamloptions"
	clusterkafka "kubeform.dev/provider-aws/internal/controller/kafka/cluster"
	configuration "kubeform.dev/provider-aws/internal/controller/kafka/configuration"
	stream "kubeform.dev/provider-aws/internal/controller/kinesis/stream"
	key "kubeform.dev/provider-aws/internal/controller/kms/key"
	acl "kubeform.dev/provider-aws/internal/controller/memorydb/acl"
	clustermemorydb "kubeform.dev/provider-aws/internal/controller/memorydb/cluster"
	parametergroupmemorydb "kubeform.dev/provider-aws/internal/controller/memorydb/parametergroup"
	snapshot "kubeform.dev/provider-aws/internal/controller/memorydb/snapshot"
	subnetgroupmemorydb "kubeform.dev/provider-aws/internal/controller/memorydb/subnetgroup"
	providerconfig "kubeform.dev/provider-aws/internal/controller/providerconfig"
	clusterrds "kubeform.dev/provider-aws/internal/controller/rds/cluster"
	clusteractivitystream "kubeform.dev/provider-aws/internal/controller/rds/clusteractivitystream"
	clusterendpoint "kubeform.dev/provider-aws/internal/controller/rds/clusterendpoint"
	clusterinstancerds "kubeform.dev/provider-aws/internal/controller/rds/clusterinstance"
	clusterparametergrouprds "kubeform.dev/provider-aws/internal/controller/rds/clusterparametergroup"
	clusterroleassociation "kubeform.dev/provider-aws/internal/controller/rds/clusterroleassociation"
	clustersnapshotrds "kubeform.dev/provider-aws/internal/controller/rds/clustersnapshot"
	dbinstanceautomatedbackupsreplication "kubeform.dev/provider-aws/internal/controller/rds/dbinstanceautomatedbackupsreplication"
	dbsnapshotcopy "kubeform.dev/provider-aws/internal/controller/rds/dbsnapshotcopy"
	eventsubscriptionrds "kubeform.dev/provider-aws/internal/controller/rds/eventsubscription"
	globalclusterrds "kubeform.dev/provider-aws/internal/controller/rds/globalcluster"
	instance "kubeform.dev/provider-aws/internal/controller/rds/instance"
	instanceroleassociation "kubeform.dev/provider-aws/internal/controller/rds/instanceroleassociation"
	optiongroup "kubeform.dev/provider-aws/internal/controller/rds/optiongroup"
	parametergrouprds "kubeform.dev/provider-aws/internal/controller/rds/parametergroup"
	proxy "kubeform.dev/provider-aws/internal/controller/rds/proxy"
	proxydefaulttargetgroup "kubeform.dev/provider-aws/internal/controller/rds/proxydefaulttargetgroup"
	proxyendpoint "kubeform.dev/provider-aws/internal/controller/rds/proxyendpoint"
	proxytarget "kubeform.dev/provider-aws/internal/controller/rds/proxytarget"
	snapshotrds "kubeform.dev/provider-aws/internal/controller/rds/snapshot"
	subnetgrouprds "kubeform.dev/provider-aws/internal/controller/rds/subnetgroup"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		clusterinstance.Setup,
		clusterparametergroup.Setup,
		clustersnapshot.Setup,
		eventsubscription.Setup,
		globalcluster.Setup,
		subnetgroup.Setup,
		contributorinsights.Setup,
		globaltable.Setup,
		kinesisstreamingdestination.Setup,
		table.Setup,
		tableitem.Setup,
		tablereplica.Setup,
		tag.Setup,
		route.Setup,
		securitygrouprule.Setup,
		vpcpeeringconnection.Setup,
		clusterelasticache.Setup,
		parametergroup.Setup,
		replicationgroup.Setup,
		subnetgroupelasticache.Setup,
		user.Setup,
		usergroup.Setup,
		domain.Setup,
		domainpolicy.Setup,
		domainsamloptions.Setup,
		clusterkafka.Setup,
		configuration.Setup,
		stream.Setup,
		key.Setup,
		acl.Setup,
		clustermemorydb.Setup,
		parametergroupmemorydb.Setup,
		snapshot.Setup,
		subnetgroupmemorydb.Setup,
		providerconfig.Setup,
		clusterrds.Setup,
		clusteractivitystream.Setup,
		clusterendpoint.Setup,
		clusterinstancerds.Setup,
		clusterparametergrouprds.Setup,
		clusterroleassociation.Setup,
		clustersnapshotrds.Setup,
		dbinstanceautomatedbackupsreplication.Setup,
		dbsnapshotcopy.Setup,
		eventsubscriptionrds.Setup,
		globalclusterrds.Setup,
		instance.Setup,
		instanceroleassociation.Setup,
		optiongroup.Setup,
		parametergrouprds.Setup,
		proxy.Setup,
		proxydefaulttargetgroup.Setup,
		proxyendpoint.Setup,
		proxytarget.Setup,
		snapshotrds.Setup,
		subnetgrouprds.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
