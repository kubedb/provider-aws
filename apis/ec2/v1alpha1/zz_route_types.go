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




type RouteInitParameters struct {


// Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
CarrierGatewayID *string `json:"carrierGatewayId,omitempty" tf:"carrier_gateway_id,omitempty"`

// The Amazon Resource Name (ARN) of a core network.
CoreNetworkArn *string `json:"coreNetworkArn,omitempty" tf:"core_network_arn,omitempty"`

// The destination CIDR block.
DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty" tf:"destination_cidr_block,omitempty"`

// The destination IPv6 CIDR block.
DestinationIPv6CidrBlock *string `json:"destinationIpv6CidrBlock,omitempty" tf:"destination_ipv6_cidr_block,omitempty"`

// The ID of a managed prefix list destination.
DestinationPrefixListID *string `json:"destinationPrefixListId,omitempty" tf:"destination_prefix_list_id,omitempty"`

// Identifier of a VPC Egress Only Internet Gateway.
EgressOnlyGatewayID *string `json:"egressOnlyGatewayId,omitempty" tf:"egress_only_gateway_id,omitempty"`

// Identifier of a VPC internet gateway or a virtual private gateway. Specify local when updating a previously imported local route.
GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

// Identifier of a Outpost local gateway.
LocalGatewayID *string `json:"localGatewayId,omitempty" tf:"local_gateway_id,omitempty"`

// Identifier of a VPC NAT gateway.
NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

// Identifier of an EC2 network interface.
NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

// The ID of the routing table.
RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

// Identifier of an EC2 Transit Gateway.
TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

// Identifier of a VPC Endpoint.
VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

// Identifier of a VPC peering connection.
VPCPeeringConnectionID *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id,omitempty"`
}


type RouteObservation struct {


// Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
CarrierGatewayID *string `json:"carrierGatewayId,omitempty" tf:"carrier_gateway_id,omitempty"`

// The Amazon Resource Name (ARN) of a core network.
CoreNetworkArn *string `json:"coreNetworkArn,omitempty" tf:"core_network_arn,omitempty"`

// The destination CIDR block.
DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty" tf:"destination_cidr_block,omitempty"`

// The destination IPv6 CIDR block.
DestinationIPv6CidrBlock *string `json:"destinationIpv6CidrBlock,omitempty" tf:"destination_ipv6_cidr_block,omitempty"`

// The ID of a managed prefix list destination.
DestinationPrefixListID *string `json:"destinationPrefixListId,omitempty" tf:"destination_prefix_list_id,omitempty"`

// Identifier of a VPC Egress Only Internet Gateway.
EgressOnlyGatewayID *string `json:"egressOnlyGatewayId,omitempty" tf:"egress_only_gateway_id,omitempty"`

// Identifier of a VPC internet gateway or a virtual private gateway. Specify local when updating a previously imported local route.
GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

// Route identifier computed from the routing table identifier and route destination.
ID *string `json:"id,omitempty" tf:"id,omitempty"`

// Identifier of an EC2 instance.
InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

// The AWS account ID of the owner of the EC2 instance.
InstanceOwnerID *string `json:"instanceOwnerId,omitempty" tf:"instance_owner_id,omitempty"`

// Identifier of a Outpost local gateway.
LocalGatewayID *string `json:"localGatewayId,omitempty" tf:"local_gateway_id,omitempty"`

// Identifier of a VPC NAT gateway.
NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

// Identifier of an EC2 network interface.
NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

// How the route was created - CreateRouteTable, CreateRoute or EnableVgwRoutePropagation.
Origin *string `json:"origin,omitempty" tf:"origin,omitempty"`

// The ID of the routing table.
RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

// The state of the route - active or blackhole.
State *string `json:"state,omitempty" tf:"state,omitempty"`

// Identifier of an EC2 Transit Gateway.
TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

// Identifier of a VPC Endpoint.
VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

// Identifier of a VPC peering connection.
VPCPeeringConnectionID *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id,omitempty"`
}


type RouteParameters struct {


// Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
// +kubebuilder:validation:Optional
CarrierGatewayID *string `json:"carrierGatewayId,omitempty" tf:"carrier_gateway_id,omitempty"`

// The Amazon Resource Name (ARN) of a core network.
// +kubebuilder:validation:Optional
CoreNetworkArn *string `json:"coreNetworkArn,omitempty" tf:"core_network_arn,omitempty"`

// The destination CIDR block.
// +kubebuilder:validation:Optional
DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty" tf:"destination_cidr_block,omitempty"`

// The destination IPv6 CIDR block.
// +kubebuilder:validation:Optional
DestinationIPv6CidrBlock *string `json:"destinationIpv6CidrBlock,omitempty" tf:"destination_ipv6_cidr_block,omitempty"`

// The ID of a managed prefix list destination.
// +kubebuilder:validation:Optional
DestinationPrefixListID *string `json:"destinationPrefixListId,omitempty" tf:"destination_prefix_list_id,omitempty"`

// Identifier of a VPC Egress Only Internet Gateway.
// +kubebuilder:validation:Optional
EgressOnlyGatewayID *string `json:"egressOnlyGatewayId,omitempty" tf:"egress_only_gateway_id,omitempty"`

// Identifier of a VPC internet gateway or a virtual private gateway. Specify local when updating a previously imported local route.
// +kubebuilder:validation:Optional
GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

// Identifier of a Outpost local gateway.
// +kubebuilder:validation:Optional
LocalGatewayID *string `json:"localGatewayId,omitempty" tf:"local_gateway_id,omitempty"`

// Identifier of a VPC NAT gateway.
// +kubebuilder:validation:Optional
NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

// Identifier of an EC2 network interface.
// +kubebuilder:validation:Optional
NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

// Region is the region you'd like your resource to be created in.
// +upjet:crd:field:TFTag=-
// +kubebuilder:validation:Optional
Region *string `json:"region,omitempty" tf:"-"`

// The ID of the routing table.
// +kubebuilder:validation:Optional
RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

// Identifier of an EC2 Transit Gateway.
// +kubebuilder:validation:Optional
TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

// Identifier of a VPC Endpoint.
// +kubebuilder:validation:Optional
VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

// Identifier of a VPC peering connection.
// +kubebuilder:validation:Optional
VPCPeeringConnectionID *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id,omitempty"`
}

// RouteSpec defines the desired state of Route
type RouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       RouteParameters `json:"forProvider"`
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
	InitProvider       RouteInitParameters `json:"initProvider,omitempty"`
}

// RouteStatus defines the observed state of Route.
type RouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          RouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route is the Schema for the Routes API. Provides a resource to create a routing entry in a VPC routing table.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region)",message="spec.forProvider.region is a required parameter"
// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.routeTableId) || (has(self.initProvider) && has(self.initProvider.routeTableId))",message="spec.forProvider.routeTableId is a required parameter"
	Spec              RouteSpec   `json:"spec"`
	Status            RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteList contains a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

// Repository type metadata.
var (
	Route_Kind             = "Route"
	Route_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Route_Kind}.String()
	Route_KindAPIVersion   = Route_Kind + "." + CRDGroupVersion.String()
	Route_GroupVersionKind = CRDGroupVersion.WithKind(Route_Kind)
)

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}
