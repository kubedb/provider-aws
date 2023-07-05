/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClusterSnapshotObservation struct {

	// Allocated storage size in gigabytes (GB).
	AllocatedStorage *float64 `json:"allocatedStorage,omitempty" tf:"allocated_storage,omitempty"`

	// List of EC2 Availability Zones that instances in the DB cluster snapshot can be restored in.
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// The DB Cluster Identifier from which to take the snapshot.
	DBClusterIdentifier *string `json:"dbClusterIdentifier,omitempty" tf:"db_cluster_identifier,omitempty"`

	// The Amazon Resource Name (ARN) for the DB Cluster Snapshot.
	DBClusterSnapshotArn *string `json:"dbClusterSnapshotArn,omitempty" tf:"db_cluster_snapshot_arn,omitempty"`

	// The Identifier for the snapshot.
	DBClusterSnapshotIdentifier *string `json:"dbClusterSnapshotIdentifier,omitempty" tf:"db_cluster_snapshot_identifier,omitempty"`

	// Name of the database engine.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Version of the database engine for this DB cluster snapshot.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// If storage_encrypted is true, the AWS KMS key identifier for the encrypted DB cluster snapshot.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// License model information for the restored DB cluster.
	LicenseModel *string `json:"licenseModel,omitempty" tf:"license_model,omitempty"`

	// Port that the DB cluster was listening on at the time of the snapshot.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	SnapshotType *string `json:"snapshotType,omitempty" tf:"snapshot_type,omitempty"`

	// The Amazon Resource Name (ARN) for the DB Cluster Snapshot.
	SourceDBClusterSnapshotArn *string `json:"sourceDbClusterSnapshotArn,omitempty" tf:"source_db_cluster_snapshot_arn,omitempty"`

	// The status of this DB Cluster Snapshot.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Whether the DB cluster snapshot is encrypted.
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`

	// A map of tags to assign to the DB cluster. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The VPC ID associated with the DB cluster snapshot.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type ClusterSnapshotParameters struct {

	// The DB Cluster Identifier from which to take the snapshot.
	// +kubebuilder:validation:Optional
	DBClusterIdentifier *string `json:"dbClusterIdentifier,omitempty" tf:"db_cluster_identifier,omitempty"`

	// The Identifier for the snapshot.
	// +kubebuilder:validation:Optional
	DBClusterSnapshotIdentifier *string `json:"dbClusterSnapshotIdentifier,omitempty" tf:"db_cluster_snapshot_identifier,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`

	// A map of tags to assign to the DB cluster. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +kubebuilder:validation:Optional
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

// ClusterSnapshotSpec defines the desired state of ClusterSnapshot
type ClusterSnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterSnapshotParameters `json:"forProvider"`
}

// ClusterSnapshotStatus defines the observed state of ClusterSnapshot.
type ClusterSnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterSnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterSnapshot is the Schema for the ClusterSnapshots API. Manages an RDS database cluster snapshot.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.dbClusterIdentifier)",message="dbClusterIdentifier is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.dbClusterSnapshotIdentifier)",message="dbClusterSnapshotIdentifier is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.region)",message="region is a required parameter"
	Spec   ClusterSnapshotSpec   `json:"spec"`
	Status ClusterSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterSnapshotList contains a list of ClusterSnapshots
type ClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterSnapshot `json:"items"`
}

// Repository type metadata.
var (
	ClusterSnapshot_Kind             = "ClusterSnapshot"
	ClusterSnapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterSnapshot_Kind}.String()
	ClusterSnapshot_KindAPIVersion   = ClusterSnapshot_Kind + "." + CRDGroupVersion.String()
	ClusterSnapshot_GroupVersionKind = CRDGroupVersion.WithKind(ClusterSnapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterSnapshot{}, &ClusterSnapshotList{})
}