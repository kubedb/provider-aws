//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Route) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteList) DeepCopyInto(out *RouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteList.
func (in *RouteList) DeepCopy() *RouteList {
	if in == nil {
		return nil
	}
	out := new(RouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteObservation) DeepCopyInto(out *RouteObservation) {
	*out = *in
	if in.CarrierGatewayID != nil {
		in, out := &in.CarrierGatewayID, &out.CarrierGatewayID
		*out = new(string)
		**out = **in
	}
	if in.CoreNetworkArn != nil {
		in, out := &in.CoreNetworkArn, &out.CoreNetworkArn
		*out = new(string)
		**out = **in
	}
	if in.DestinationCidrBlock != nil {
		in, out := &in.DestinationCidrBlock, &out.DestinationCidrBlock
		*out = new(string)
		**out = **in
	}
	if in.DestinationIPv6CidrBlock != nil {
		in, out := &in.DestinationIPv6CidrBlock, &out.DestinationIPv6CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.DestinationPrefixListID != nil {
		in, out := &in.DestinationPrefixListID, &out.DestinationPrefixListID
		*out = new(string)
		**out = **in
	}
	if in.EgressOnlyGatewayID != nil {
		in, out := &in.EgressOnlyGatewayID, &out.EgressOnlyGatewayID
		*out = new(string)
		**out = **in
	}
	if in.GatewayID != nil {
		in, out := &in.GatewayID, &out.GatewayID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceOwnerID != nil {
		in, out := &in.InstanceOwnerID, &out.InstanceOwnerID
		*out = new(string)
		**out = **in
	}
	if in.LocalGatewayID != nil {
		in, out := &in.LocalGatewayID, &out.LocalGatewayID
		*out = new(string)
		**out = **in
	}
	if in.NATGatewayID != nil {
		in, out := &in.NATGatewayID, &out.NATGatewayID
		*out = new(string)
		**out = **in
	}
	if in.NetworkInterfaceID != nil {
		in, out := &in.NetworkInterfaceID, &out.NetworkInterfaceID
		*out = new(string)
		**out = **in
	}
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = new(string)
		**out = **in
	}
	if in.RouteTableID != nil {
		in, out := &in.RouteTableID, &out.RouteTableID
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.TransitGatewayID != nil {
		in, out := &in.TransitGatewayID, &out.TransitGatewayID
		*out = new(string)
		**out = **in
	}
	if in.VPCEndpointID != nil {
		in, out := &in.VPCEndpointID, &out.VPCEndpointID
		*out = new(string)
		**out = **in
	}
	if in.VPCPeeringConnectionID != nil {
		in, out := &in.VPCPeeringConnectionID, &out.VPCPeeringConnectionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteObservation.
func (in *RouteObservation) DeepCopy() *RouteObservation {
	if in == nil {
		return nil
	}
	out := new(RouteObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteParameters) DeepCopyInto(out *RouteParameters) {
	*out = *in
	if in.CarrierGatewayID != nil {
		in, out := &in.CarrierGatewayID, &out.CarrierGatewayID
		*out = new(string)
		**out = **in
	}
	if in.CoreNetworkArn != nil {
		in, out := &in.CoreNetworkArn, &out.CoreNetworkArn
		*out = new(string)
		**out = **in
	}
	if in.DestinationCidrBlock != nil {
		in, out := &in.DestinationCidrBlock, &out.DestinationCidrBlock
		*out = new(string)
		**out = **in
	}
	if in.DestinationIPv6CidrBlock != nil {
		in, out := &in.DestinationIPv6CidrBlock, &out.DestinationIPv6CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.DestinationPrefixListID != nil {
		in, out := &in.DestinationPrefixListID, &out.DestinationPrefixListID
		*out = new(string)
		**out = **in
	}
	if in.EgressOnlyGatewayID != nil {
		in, out := &in.EgressOnlyGatewayID, &out.EgressOnlyGatewayID
		*out = new(string)
		**out = **in
	}
	if in.GatewayID != nil {
		in, out := &in.GatewayID, &out.GatewayID
		*out = new(string)
		**out = **in
	}
	if in.LocalGatewayID != nil {
		in, out := &in.LocalGatewayID, &out.LocalGatewayID
		*out = new(string)
		**out = **in
	}
	if in.NATGatewayID != nil {
		in, out := &in.NATGatewayID, &out.NATGatewayID
		*out = new(string)
		**out = **in
	}
	if in.NetworkInterfaceID != nil {
		in, out := &in.NetworkInterfaceID, &out.NetworkInterfaceID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RouteTableID != nil {
		in, out := &in.RouteTableID, &out.RouteTableID
		*out = new(string)
		**out = **in
	}
	if in.TransitGatewayID != nil {
		in, out := &in.TransitGatewayID, &out.TransitGatewayID
		*out = new(string)
		**out = **in
	}
	if in.VPCEndpointID != nil {
		in, out := &in.VPCEndpointID, &out.VPCEndpointID
		*out = new(string)
		**out = **in
	}
	if in.VPCPeeringConnectionID != nil {
		in, out := &in.VPCPeeringConnectionID, &out.VPCPeeringConnectionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteParameters.
func (in *RouteParameters) DeepCopy() *RouteParameters {
	if in == nil {
		return nil
	}
	out := new(RouteParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSpec) DeepCopyInto(out *RouteSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSpec.
func (in *RouteSpec) DeepCopy() *RouteSpec {
	if in == nil {
		return nil
	}
	out := new(RouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteStatus) DeepCopyInto(out *RouteStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteStatus.
func (in *RouteStatus) DeepCopy() *RouteStatus {
	if in == nil {
		return nil
	}
	out := new(RouteStatus)
	in.DeepCopyInto(out)
	return out
}