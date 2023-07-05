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

type TableReplicaObservation_2 struct {

	// ARN of the table replica.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ARN of the main or global table which this resource will replicate.
	GlobalTableArn *string `json:"globalTableArn,omitempty" tf:"global_table_arn,omitempty"`

	// Name of the table and region of the main global table joined with a semicolon (e.g., TableName:us-east-1).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, alias/aws/dynamodb. Note: This attribute will not be populated with the ARN of default keys.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Whether to enable Point In Time Recovery for the replica. Default is false.
	PointInTimeRecovery *bool `json:"pointInTimeRecovery,omitempty" tf:"point_in_time_recovery,omitempty"`

	// Storage class of the table replica. Valid values are STANDARD and STANDARD_INFREQUENT_ACCESS. If not used, the table replica will use the same class as the global table.
	TableClassOverride *string `json:"tableClassOverride,omitempty" tf:"table_class_override,omitempty"`

	// Map of tags to populate on the created table. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TableReplicaParameters_2 struct {

	// ARN of the main or global table which this resource will replicate.
	// +kubebuilder:validation:Optional
	GlobalTableArn *string `json:"globalTableArn,omitempty" tf:"global_table_arn,omitempty"`

	// ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, alias/aws/dynamodb. Note: This attribute will not be populated with the ARN of default keys.
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Whether to enable Point In Time Recovery for the replica. Default is false.
	// +kubebuilder:validation:Optional
	PointInTimeRecovery *bool `json:"pointInTimeRecovery,omitempty" tf:"point_in_time_recovery,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`

	// Storage class of the table replica. Valid values are STANDARD and STANDARD_INFREQUENT_ACCESS. If not used, the table replica will use the same class as the global table.
	// +kubebuilder:validation:Optional
	TableClassOverride *string `json:"tableClassOverride,omitempty" tf:"table_class_override,omitempty"`

	// Map of tags to populate on the created table. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +kubebuilder:validation:Optional
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

// TableReplicaSpec defines the desired state of TableReplica
type TableReplicaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableReplicaParameters_2 `json:"forProvider"`
}

// TableReplicaStatus defines the observed state of TableReplica.
type TableReplicaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableReplicaObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TableReplica is the Schema for the TableReplicas API. Provides a DynamoDB table replica resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TableReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.globalTableArn)",message="globalTableArn is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.region)",message="region is a required parameter"
	Spec   TableReplicaSpec   `json:"spec"`
	Status TableReplicaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableReplicaList contains a list of TableReplicas
type TableReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TableReplica `json:"items"`
}

// Repository type metadata.
var (
	TableReplica_Kind             = "TableReplica"
	TableReplica_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TableReplica_Kind}.String()
	TableReplica_KindAPIVersion   = TableReplica_Kind + "." + CRDGroupVersion.String()
	TableReplica_GroupVersionKind = CRDGroupVersion.WithKind(TableReplica_Kind)
)

func init() {
	SchemeBuilder.Register(&TableReplica{}, &TableReplicaList{})
}