package controller

import (
	"context"
	"sync"

	"github.com/crossplane/upjet/pkg/controller"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
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
	securitygroup "kubedb.dev/provider-aws/internal/controller/ec2/securitygroup"
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
	role "kubedb.dev/provider-aws/internal/controller/iam/role"
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
	secret "kubedb.dev/provider-aws/internal/controller/secretsmanager/secret"
	topic "kubedb.dev/provider-aws/internal/controller/sns/topic"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

var (
	setupFns = map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{"docdb.aws.kubedb.com", "Cluster"}:                             cluster.Setup,
		schema.GroupKind{"docdb.aws.kubedb.com", "ClusterInstance"}:                     clusterinstance.Setup,
		schema.GroupKind{"docdb.aws.kubedb.com", "ClusterParameterGroup"}:               clusterparametergroup.Setup,
		schema.GroupKind{"docdb.aws.kubedb.com", "ClusterSnapshot"}:                     clustersnapshot.Setup,
		schema.GroupKind{"docdb.aws.kubedb.com", "EventSubscription"}:                   eventsubscription.Setup,
		schema.GroupKind{"docdb.aws.kubedb.com", "GlobalCluster"}:                       globalcluster.Setup,
		schema.GroupKind{"docdb.aws.kubedb.com", "SubnetGroup"}:                         subnetgroup.Setup,
		schema.GroupKind{"dynamodb.aws.kubedb.com", "ContributorInsights"}:              contributorinsights.Setup,
		schema.GroupKind{"dynamodb.aws.kubedb.com", "GlobalTable"}:                      globaltable.Setup,
		schema.GroupKind{"dynamodb.aws.kubedb.com", "KinesisStreamingDestination"}:      kinesisstreamingdestination.Setup,
		schema.GroupKind{"dynamodb.aws.kubedb.com", "Table"}:                            table.Setup,
		schema.GroupKind{"dynamodb.aws.kubedb.com", "TableItem"}:                        tableitem.Setup,
		schema.GroupKind{"dynamodb.aws.kubedb.com", "TableReplica"}:                     tablereplica.Setup,
		schema.GroupKind{"dynamodb.aws.kubedb.com", "Tag"}:                              tag.Setup,
		schema.GroupKind{"ec2.aws.kubedb.com", "Route"}:                                 route.Setup,
		schema.GroupKind{"ec2.aws.kubedb.com", "SecurityGroup"}:                         securitygroup.Setup,
		schema.GroupKind{"ec2.aws.kubedb.com", "SecurityGroupRule"}:                     securitygrouprule.Setup,
		schema.GroupKind{"ec2.aws.kubedb.com", "Subnet"}:                                subnet.Setup,
		schema.GroupKind{"ec2.aws.kubedb.com", "VPC"}:                                   vpc.Setup,
		schema.GroupKind{"ec2.aws.kubedb.com", "VPCEndpoint"}:                           vpcendpoint.Setup,
		schema.GroupKind{"ec2.aws.kubedb.com", "VPCPeeringConnection"}:                  vpcpeeringconnection.Setup,
		schema.GroupKind{"elasticache.aws.kubedb.com", "Cluster"}:                       clusterelasticache.Setup,
		schema.GroupKind{"elasticache.aws.kubedb.com", "ParameterGroup"}:                parametergroup.Setup,
		schema.GroupKind{"elasticache.aws.kubedb.com", "ReplicationGroup"}:              replicationgroup.Setup,
		schema.GroupKind{"elasticache.aws.kubedb.com", "SubnetGroup"}:                   subnetgroupelasticache.Setup,
		schema.GroupKind{"elasticache.aws.kubedb.com", "User"}:                          user.Setup,
		schema.GroupKind{"elasticache.aws.kubedb.com", "UserGroup"}:                     usergroup.Setup,
		schema.GroupKind{"elasticsearch.aws.kubedb.com", "Domain"}:                      domain.Setup,
		schema.GroupKind{"elasticsearch.aws.kubedb.com", "DomainPolicy"}:                domainpolicy.Setup,
		schema.GroupKind{"elasticsearch.aws.kubedb.com", "DomainSAMLOptions"}:           domainsamloptions.Setup,
		schema.GroupKind{"iam.aws.kubedb.com", "Role"}:                                  role.Setup,
		schema.GroupKind{"kafka.aws.kubedb.com", "Cluster"}:                             clusterkafka.Setup,
		schema.GroupKind{"kafka.aws.kubedb.com", "Configuration"}:                       configuration.Setup,
		schema.GroupKind{"kinesis.aws.kubedb.com", "Stream"}:                            stream.Setup,
		schema.GroupKind{"kms.aws.kubedb.com", "Key"}:                                   key.Setup,
		schema.GroupKind{"memorydb.aws.kubedb.com", "ACL"}:                              acl.Setup,
		schema.GroupKind{"memorydb.aws.kubedb.com", "Cluster"}:                          clustermemorydb.Setup,
		schema.GroupKind{"memorydb.aws.kubedb.com", "ParameterGroup"}:                   parametergroupmemorydb.Setup,
		schema.GroupKind{"memorydb.aws.kubedb.com", "Snapshot"}:                         snapshot.Setup,
		schema.GroupKind{"memorydb.aws.kubedb.com", "SubnetGroup"}:                      subnetgroupmemorydb.Setup,
		schema.GroupKind{"providerconfig.aws.kubedb.com", ""}:                           providerconfig.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "Cluster"}:                               clusterrds.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "ClusterActivityStream"}:                 clusteractivitystream.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "ClusterEndpoint"}:                       clusterendpoint.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "ClusterInstance"}:                       clusterinstancerds.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "ClusterParameterGroup"}:                 clusterparametergrouprds.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "ClusterRoleAssociation"}:                clusterroleassociation.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "ClusterSnapshot"}:                       clustersnapshotrds.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "DBInstanceAutomatedBackupsReplication"}: dbinstanceautomatedbackupsreplication.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "DBSnapshotCopy"}:                        dbsnapshotcopy.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "EventSubscription"}:                     eventsubscriptionrds.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "GlobalCluster"}:                         globalclusterrds.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "Instance"}:                              instance.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "InstanceRoleAssociation"}:               instanceroleassociation.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "OptionGroup"}:                           optiongroup.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "ParameterGroup"}:                        parametergrouprds.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "Proxy"}:                                 proxy.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "ProxyDefaultTargetGroup"}:               proxydefaulttargetgroup.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "ProxyEndpoint"}:                         proxyendpoint.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "ProxyTarget"}:                           proxytarget.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "Snapshot"}:                              snapshotrds.Setup,
		schema.GroupKind{"rds.aws.kubedb.com", "SubnetGroup"}:                           subnetgrouprds.Setup,
		schema.GroupKind{"secretsmanager.aws.kubedb.com", "Secret"}:                     secret.Setup,
		schema.GroupKind{"sns.aws.kubedb.com", "Topic"}:                                 topic.Setup,
	}
)

//package controller

var (
	setupDone = map[schema.GroupKind]bool{}
	mu        sync.RWMutex
)

type CustomResourceReconciler struct {
	mgr ctrl.Manager
	o   controller.Options
}

func NewCustomResourceReconciler(mgr ctrl.Manager, o controller.Options) *CustomResourceReconciler {
	return &CustomResourceReconciler{mgr: mgr, o: o}
}

func (r *CustomResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	var crd apiextensions.CustomResourceDefinition
	if err := r.mgr.GetClient().Get(ctx, req.NamespacedName, &crd); err != nil {
		log.Error(err, "unable to fetch CustomResourceDefinition")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	gk := schema.GroupKind{
		Group: crd.Spec.Group,
		Kind:  crd.Spec.Names.Kind,
	}
	mu.Lock()
	defer mu.Unlock()
	_, found := setupDone[gk]
	if found {
		return ctrl.Result{}, nil
	}
	setup, found := setupFns[gk]
	if found {
		setup(r.mgr, r.o)
		setupDone[gk] = true
	}

	return ctrl.Result{}, nil
}

func (r *CustomResourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiextensions.CustomResourceDefinition{}).
		Complete(r)
}
