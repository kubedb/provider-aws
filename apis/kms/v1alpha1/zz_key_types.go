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

type KeyObservation struct {

	// The Amazon Resource Name (ARN) of the key.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the Default Key Policy section in the AWS Key Management Service Developer Guide.
	// The default value is false.
	BypassPolicyLockoutSafetyCheck *bool `json:"bypassPolicyLockoutSafetyCheck,omitempty" tf:"bypass_policy_lockout_safety_check,omitempty"`

	// ID of the KMS Custom Key Store where the key will be stored instead of KMS (eg CloudHSM).
	CustomKeyStoreID *string `json:"customKeyStoreId,omitempty" tf:"custom_key_store_id,omitempty"`

	// Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
	// Valid values: SYMMETRIC_DEFAULT,  RSA_2048, RSA_3072, RSA_4096, HMAC_256, ECC_NIST_P256, ECC_NIST_P384, ECC_NIST_P521, or ECC_SECG_P256K1. Defaults to SYMMETRIC_DEFAULT. For help with choosing a key spec, see the AWS KMS Developer Guide.
	CustomerMasterKeySpec *string `json:"customerMasterKeySpec,omitempty" tf:"customer_master_key_spec,omitempty"`

	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between 7 and 30, inclusive. If you do not specify a value, it defaults to 30.
	// If the KMS key is a multi-Region primary key with replicas, the waiting period begins when the last of its replica keys is deleted. Otherwise, the waiting period begins immediately.
	DeletionWindowInDays *float64 `json:"deletionWindowInDays,omitempty" tf:"deletion_window_in_days,omitempty"`

	// The description of the key as viewed in AWS console.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether key rotation is enabled. Defaults to false.
	EnableKeyRotation *bool `json:"enableKeyRotation,omitempty" tf:"enable_key_rotation,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies whether the key is enabled. Defaults to true.
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// The globally unique identifier for the key.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Specifies the intended use of the key. Valid values: ENCRYPT_DECRYPT, SIGN_VERIFY, or GENERATE_VERIFY_MAC.
	// Defaults to ENCRYPT_DECRYPT.
	KeyUsage *string `json:"keyUsage,omitempty" tf:"key_usage,omitempty"`

	// Indicates whether the KMS key is a multi-Region (true) or regional (false) key. Defaults to false.
	MultiRegion *bool `json:"multiRegion,omitempty" tf:"multi_region,omitempty"`

	// A valid policy JSON document. Although this is a key policy, not an IAM policy, an aws_iam_policy_document, in the form that designates a principal, can be used.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// A map of tags to assign to the object. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type KeyParameters struct {

	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the Default Key Policy section in the AWS Key Management Service Developer Guide.
	// The default value is false.
	// +kubebuilder:validation:Optional
	BypassPolicyLockoutSafetyCheck *bool `json:"bypassPolicyLockoutSafetyCheck,omitempty" tf:"bypass_policy_lockout_safety_check,omitempty"`

	// ID of the KMS Custom Key Store where the key will be stored instead of KMS (eg CloudHSM).
	// +kubebuilder:validation:Optional
	CustomKeyStoreID *string `json:"customKeyStoreId,omitempty" tf:"custom_key_store_id,omitempty"`

	// Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
	// Valid values: SYMMETRIC_DEFAULT,  RSA_2048, RSA_3072, RSA_4096, HMAC_256, ECC_NIST_P256, ECC_NIST_P384, ECC_NIST_P521, or ECC_SECG_P256K1. Defaults to SYMMETRIC_DEFAULT. For help with choosing a key spec, see the AWS KMS Developer Guide.
	// +kubebuilder:validation:Optional
	CustomerMasterKeySpec *string `json:"customerMasterKeySpec,omitempty" tf:"customer_master_key_spec,omitempty"`

	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between 7 and 30, inclusive. If you do not specify a value, it defaults to 30.
	// If the KMS key is a multi-Region primary key with replicas, the waiting period begins when the last of its replica keys is deleted. Otherwise, the waiting period begins immediately.
	// +kubebuilder:validation:Optional
	DeletionWindowInDays *float64 `json:"deletionWindowInDays,omitempty" tf:"deletion_window_in_days,omitempty"`

	// The description of the key as viewed in AWS console.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether key rotation is enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableKeyRotation *bool `json:"enableKeyRotation,omitempty" tf:"enable_key_rotation,omitempty"`

	// Specifies whether the key is enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// Specifies the intended use of the key. Valid values: ENCRYPT_DECRYPT, SIGN_VERIFY, or GENERATE_VERIFY_MAC.
	// Defaults to ENCRYPT_DECRYPT.
	// +kubebuilder:validation:Optional
	KeyUsage *string `json:"keyUsage,omitempty" tf:"key_usage,omitempty"`

	// Indicates whether the KMS key is a multi-Region (true) or regional (false) key. Defaults to false.
	// +kubebuilder:validation:Optional
	MultiRegion *bool `json:"multiRegion,omitempty" tf:"multi_region,omitempty"`

	// A valid policy JSON document. Although this is a key policy, not an IAM policy, an aws_iam_policy_document, in the form that designates a principal, can be used.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`

	// A map of tags to assign to the object. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +kubebuilder:validation:Optional
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

// KeySpec defines the desired state of Key
type KeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyParameters `json:"forProvider"`
}

// KeyStatus defines the observed state of Key.
type KeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Key is the Schema for the Keys API. Manages a single-Region or multi-Region primary KMS key.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Key struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.region)",message="region is a required parameter"
	Spec   KeySpec   `json:"spec"`
	Status KeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyList contains a list of Keys
type KeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Key `json:"items"`
}

// Repository type metadata.
var (
	Key_Kind             = "Key"
	Key_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Key_Kind}.String()
	Key_KindAPIVersion   = Key_Kind + "." + CRDGroupVersion.String()
	Key_GroupVersionKind = CRDGroupVersion.WithKind(Key_Kind)
)

func init() {
	SchemeBuilder.Register(&Key{}, &KeyList{})
}
