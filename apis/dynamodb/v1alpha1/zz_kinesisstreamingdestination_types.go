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

type KinesisStreamingDestinationObservation struct {

	// The table_name and stream_arn separated by a comma (,).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN for a Kinesis data stream. This must exist in the same account and region as the DynamoDB table.
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`

	// The name of the DynamoDB table. There
	// can only be one Kinesis streaming destination for a given DynamoDB table.
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`
}

type KinesisStreamingDestinationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`

	// The ARN for a Kinesis data stream. This must exist in the same account and region as the DynamoDB table.
	// +crossplane:generate:reference:type=kubedb.dev/provider-aws/apis/kinesis/v1alpha1.Stream
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	// +kubebuilder:validation:Optional
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`

	// Reference to a Stream in kinesis to populate streamArn.
	// +kubebuilder:validation:Optional
	StreamArnRef *v1.Reference `json:"streamArnRef,omitempty" tf:"-"`

	// Selector for a Stream in kinesis to populate streamArn.
	// +kubebuilder:validation:Optional
	StreamArnSelector *v1.Selector `json:"streamArnSelector,omitempty" tf:"-"`

	// The name of the DynamoDB table. There
	// can only be one Kinesis streaming destination for a given DynamoDB table.
	// +crossplane:generate:reference:type=Table
	// +kubebuilder:validation:Optional
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`

	// Reference to a Table to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameRef *v1.Reference `json:"tableNameRef,omitempty" tf:"-"`

	// Selector for a Table to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameSelector *v1.Selector `json:"tableNameSelector,omitempty" tf:"-"`
}

// KinesisStreamingDestinationSpec defines the desired state of KinesisStreamingDestination
type KinesisStreamingDestinationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KinesisStreamingDestinationParameters `json:"forProvider"`
}

// KinesisStreamingDestinationStatus defines the observed state of KinesisStreamingDestination.
type KinesisStreamingDestinationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KinesisStreamingDestinationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KinesisStreamingDestination is the Schema for the KinesisStreamingDestinations API. Enables a Kinesis streaming destination for a DynamoDB table
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type KinesisStreamingDestination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.region)",message="region is a required parameter"
	Spec   KinesisStreamingDestinationSpec   `json:"spec"`
	Status KinesisStreamingDestinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KinesisStreamingDestinationList contains a list of KinesisStreamingDestinations
type KinesisStreamingDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KinesisStreamingDestination `json:"items"`
}

// Repository type metadata.
var (
	KinesisStreamingDestination_Kind             = "KinesisStreamingDestination"
	KinesisStreamingDestination_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KinesisStreamingDestination_Kind}.String()
	KinesisStreamingDestination_KindAPIVersion   = KinesisStreamingDestination_Kind + "." + CRDGroupVersion.String()
	KinesisStreamingDestination_GroupVersionKind = CRDGroupVersion.WithKind(KinesisStreamingDestination_Kind)
)

func init() {
	SchemeBuilder.Register(&KinesisStreamingDestination{}, &KinesisStreamingDestinationList{})
}
