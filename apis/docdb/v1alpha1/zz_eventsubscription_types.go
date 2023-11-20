// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type EventSubscriptionInitParameters struct {

	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Event.html or run aws docdb describe-event-categories.
	EventCategories []*string `json:"eventCategories,omitempty" tf:"event_categories,omitempty"`

	// The Amazon Resource Name of the DocumentDB event notification subscription
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a source_type must also be specified.
	SourceIds []*string `json:"sourceIds,omitempty" tf:"source_ids,omitempty"`

	// The type of source that will be generating the events. Valid options are db-instance, db-cluster, db-parameter-group, db-security-group, db-cluster-snapshot. If not set, all sources will be subscribed to.
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EventSubscriptionObservation struct {

	// The Amazon Resource Name of the DocumentDB event notification subscription
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The AWS customer account associated with the DocumentDB event notification subscription
	CustomerAwsID *string `json:"customerAwsId,omitempty" tf:"customer_aws_id,omitempty"`

	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Event.html or run aws docdb describe-event-categories.
	EventCategories []*string `json:"eventCategories,omitempty" tf:"event_categories,omitempty"`

	// The name of the DocumentDB event notification subscription
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon Resource Name of the DocumentDB event notification subscription
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a source_type must also be specified.
	SourceIds []*string `json:"sourceIds,omitempty" tf:"source_ids,omitempty"`

	// The type of source that will be generating the events. Valid options are db-instance, db-cluster, db-parameter-group, db-security-group, db-cluster-snapshot. If not set, all sources will be subscribed to.
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EventSubscriptionParameters struct {

	// A boolean flag to enable/disable the subscription. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Event.html or run aws docdb describe-event-categories.
	// +kubebuilder:validation:Optional
	EventCategories []*string `json:"eventCategories,omitempty" tf:"event_categories,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`

	// The Amazon Resource Name of the DocumentDB event notification subscription
	// +kubebuilder:validation:Optional
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a source_type must also be specified.
	// +kubebuilder:validation:Optional
	SourceIds []*string `json:"sourceIds,omitempty" tf:"source_ids,omitempty"`

	// The type of source that will be generating the events. Valid options are db-instance, db-cluster, db-parameter-group, db-security-group, db-cluster-snapshot. If not set, all sources will be subscribed to.
	// +kubebuilder:validation:Optional
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +kubebuilder:validation:Optional
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

// EventSubscriptionSpec defines the desired state of EventSubscription
type EventSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventSubscriptionParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider EventSubscriptionInitParameters `json:"initProvider,omitempty"`
}

// EventSubscriptionStatus defines the observed state of EventSubscription.
type EventSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventSubscription is the Schema for the EventSubscriptions API. Provides a DocumentDB event subscription resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EventSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region)",message="spec.forProvider.region is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.snsTopicArn) || (has(self.initProvider) && has(self.initProvider.snsTopicArn))",message="spec.forProvider.snsTopicArn is a required parameter"
	Spec   EventSubscriptionSpec   `json:"spec"`
	Status EventSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventSubscriptionList contains a list of EventSubscriptions
type EventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventSubscription `json:"items"`
}

// Repository type metadata.
var (
	EventSubscription_Kind             = "EventSubscription"
	EventSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventSubscription_Kind}.String()
	EventSubscription_KindAPIVersion   = EventSubscription_Kind + "." + CRDGroupVersion.String()
	EventSubscription_GroupVersionKind = CRDGroupVersion.WithKind(EventSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&EventSubscription{}, &EventSubscriptionList{})
}
