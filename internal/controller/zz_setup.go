/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	cluster "kubedb.dev/provider-aws/internal/controller/docdb/cluster"
	clusterinstance "kubedb.dev/provider-aws/internal/controller/docdb/clusterinstance"
	clusterparametergroup "kubedb.dev/provider-aws/internal/controller/docdb/clusterparametergroup"
	clustersnapshot "kubedb.dev/provider-aws/internal/controller/docdb/clustersnapshot"
	eventsubscription "kubedb.dev/provider-aws/internal/controller/docdb/eventsubscription"
	globalcluster "kubedb.dev/provider-aws/internal/controller/docdb/globalcluster"
	subnetgroup "kubedb.dev/provider-aws/internal/controller/docdb/subnetgroup"
	contributorinsights "kubedb.dev/provider-aws/internal/controller/dynamodb/contributorinsights"
	globaltable "kubedb.dev/provider-aws/internal/controller/dynamodb/globaltable"
	kinesisstreamingdestination "kubedb.dev/provider-aws/internal/controller/dynamodb/kinesisstreamingdestination"
	table "kubedb.dev/provider-aws/internal/controller/dynamodb/table"
	tableitem "kubedb.dev/provider-aws/internal/controller/dynamodb/tableitem"
	tablereplica "kubedb.dev/provider-aws/internal/controller/dynamodb/tablereplica"
	tag "kubedb.dev/provider-aws/internal/controller/dynamodb/tag"
	route "kubedb.dev/provider-aws/internal/controller/ec2/route"
	securitygrouprule "kubedb.dev/provider-aws/internal/controller/ec2/securitygrouprule"
	subnet "kubedb.dev/provider-aws/internal/controller/ec2/subnet"
	vpc "kubedb.dev/provider-aws/internal/controller/ec2/vpc"
	vpcendpoint "kubedb.dev/provider-aws/internal/controller/ec2/vpcendpoint"
	vpcpeeringconnection "kubedb.dev/provider-aws/internal/controller/ec2/vpcpeeringconnection"
	clusterelasticache "kubedb.dev/provider-aws/internal/controller/elasticache/cluster"
	parametergroup "kubedb.dev/provider-aws/internal/controller/elasticache/parametergroup"
	replicationgroup "kubedb.dev/provider-aws/internal/controller/elasticache/replicationgroup"
	subnetgroupelasticache "kubedb.dev/provider-aws/internal/controller/elasticache/subnetgroup"
	user "kubedb.dev/provider-aws/internal/controller/elasticache/user"
	usergroup "kubedb.dev/provider-aws/internal/controller/elasticache/usergroup"
	domain "kubedb.dev/provider-aws/internal/controller/elasticsearch/domain"
	domainpolicy "kubedb.dev/provider-aws/internal/controller/elasticsearch/domainpolicy"
	domainsamloptions "kubedb.dev/provider-aws/internal/controller/elasticsearch/domainsamloptions"
	clusterkafka "kubedb.dev/provider-aws/internal/controller/kafka/cluster"
	configuration "kubedb.dev/provider-aws/internal/controller/kafka/configuration"
	stream "kubedb.dev/provider-aws/internal/controller/kinesis/stream"
	key "kubedb.dev/provider-aws/internal/controller/kms/key"
	acl "kubedb.dev/provider-aws/internal/controller/memorydb/acl"
	clustermemorydb "kubedb.dev/provider-aws/internal/controller/memorydb/cluster"
	parametergroupmemorydb "kubedb.dev/provider-aws/internal/controller/memorydb/parametergroup"
	snapshot "kubedb.dev/provider-aws/internal/controller/memorydb/snapshot"
	subnetgroupmemorydb "kubedb.dev/provider-aws/internal/controller/memorydb/subnetgroup"
	providerconfig "kubedb.dev/provider-aws/internal/controller/providerconfig"
	clusterrds "kubedb.dev/provider-aws/internal/controller/rds/cluster"
	clusteractivitystream "kubedb.dev/provider-aws/internal/controller/rds/clusteractivitystream"
	clusterendpoint "kubedb.dev/provider-aws/internal/controller/rds/clusterendpoint"
	clusterinstancerds "kubedb.dev/provider-aws/internal/controller/rds/clusterinstance"
	clusterparametergrouprds "kubedb.dev/provider-aws/internal/controller/rds/clusterparametergroup"
	clusterroleassociation "kubedb.dev/provider-aws/internal/controller/rds/clusterroleassociation"
	clustersnapshotrds "kubedb.dev/provider-aws/internal/controller/rds/clustersnapshot"
	dbinstanceautomatedbackupsreplication "kubedb.dev/provider-aws/internal/controller/rds/dbinstanceautomatedbackupsreplication"
	dbsnapshotcopy "kubedb.dev/provider-aws/internal/controller/rds/dbsnapshotcopy"
	eventsubscriptionrds "kubedb.dev/provider-aws/internal/controller/rds/eventsubscription"
	globalclusterrds "kubedb.dev/provider-aws/internal/controller/rds/globalcluster"
	instance "kubedb.dev/provider-aws/internal/controller/rds/instance"
	instanceroleassociation "kubedb.dev/provider-aws/internal/controller/rds/instanceroleassociation"
	optiongroup "kubedb.dev/provider-aws/internal/controller/rds/optiongroup"
	parametergrouprds "kubedb.dev/provider-aws/internal/controller/rds/parametergroup"
	proxy "kubedb.dev/provider-aws/internal/controller/rds/proxy"
	proxydefaulttargetgroup "kubedb.dev/provider-aws/internal/controller/rds/proxydefaulttargetgroup"
	proxyendpoint "kubedb.dev/provider-aws/internal/controller/rds/proxyendpoint"
	proxytarget "kubedb.dev/provider-aws/internal/controller/rds/proxytarget"
	snapshotrds "kubedb.dev/provider-aws/internal/controller/rds/snapshot"
	subnetgrouprds "kubedb.dev/provider-aws/internal/controller/rds/subnetgroup"
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
		subnet.Setup,
		vpc.Setup,
		vpcendpoint.Setup,
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
