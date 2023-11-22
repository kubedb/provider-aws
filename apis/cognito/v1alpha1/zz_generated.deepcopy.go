//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsConfigurationInitParameters) DeepCopyInto(out *AnalyticsConfigurationInitParameters) {
	*out = *in
	if in.ApplicationArn != nil {
		in, out := &in.ApplicationArn, &out.ApplicationArn
		*out = new(string)
		**out = **in
	}
	if in.ApplicationID != nil {
		in, out := &in.ApplicationID, &out.ApplicationID
		*out = new(string)
		**out = **in
	}
	if in.ExternalID != nil {
		in, out := &in.ExternalID, &out.ExternalID
		*out = new(string)
		**out = **in
	}
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
	if in.UserDataShared != nil {
		in, out := &in.UserDataShared, &out.UserDataShared
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsConfigurationInitParameters.
func (in *AnalyticsConfigurationInitParameters) DeepCopy() *AnalyticsConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(AnalyticsConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsConfigurationObservation) DeepCopyInto(out *AnalyticsConfigurationObservation) {
	*out = *in
	if in.ApplicationArn != nil {
		in, out := &in.ApplicationArn, &out.ApplicationArn
		*out = new(string)
		**out = **in
	}
	if in.ApplicationID != nil {
		in, out := &in.ApplicationID, &out.ApplicationID
		*out = new(string)
		**out = **in
	}
	if in.ExternalID != nil {
		in, out := &in.ExternalID, &out.ExternalID
		*out = new(string)
		**out = **in
	}
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
	if in.UserDataShared != nil {
		in, out := &in.UserDataShared, &out.UserDataShared
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsConfigurationObservation.
func (in *AnalyticsConfigurationObservation) DeepCopy() *AnalyticsConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(AnalyticsConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsConfigurationParameters) DeepCopyInto(out *AnalyticsConfigurationParameters) {
	*out = *in
	if in.ApplicationArn != nil {
		in, out := &in.ApplicationArn, &out.ApplicationArn
		*out = new(string)
		**out = **in
	}
	if in.ApplicationID != nil {
		in, out := &in.ApplicationID, &out.ApplicationID
		*out = new(string)
		**out = **in
	}
	if in.ExternalID != nil {
		in, out := &in.ExternalID, &out.ExternalID
		*out = new(string)
		**out = **in
	}
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
	if in.UserDataShared != nil {
		in, out := &in.UserDataShared, &out.UserDataShared
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsConfigurationParameters.
func (in *AnalyticsConfigurationParameters) DeepCopy() *AnalyticsConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(AnalyticsConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenValidityUnitsInitParameters) DeepCopyInto(out *TokenValidityUnitsInitParameters) {
	*out = *in
	if in.AccessToken != nil {
		in, out := &in.AccessToken, &out.AccessToken
		*out = new(string)
		**out = **in
	}
	if in.IDToken != nil {
		in, out := &in.IDToken, &out.IDToken
		*out = new(string)
		**out = **in
	}
	if in.RefreshToken != nil {
		in, out := &in.RefreshToken, &out.RefreshToken
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenValidityUnitsInitParameters.
func (in *TokenValidityUnitsInitParameters) DeepCopy() *TokenValidityUnitsInitParameters {
	if in == nil {
		return nil
	}
	out := new(TokenValidityUnitsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenValidityUnitsObservation) DeepCopyInto(out *TokenValidityUnitsObservation) {
	*out = *in
	if in.AccessToken != nil {
		in, out := &in.AccessToken, &out.AccessToken
		*out = new(string)
		**out = **in
	}
	if in.IDToken != nil {
		in, out := &in.IDToken, &out.IDToken
		*out = new(string)
		**out = **in
	}
	if in.RefreshToken != nil {
		in, out := &in.RefreshToken, &out.RefreshToken
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenValidityUnitsObservation.
func (in *TokenValidityUnitsObservation) DeepCopy() *TokenValidityUnitsObservation {
	if in == nil {
		return nil
	}
	out := new(TokenValidityUnitsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenValidityUnitsParameters) DeepCopyInto(out *TokenValidityUnitsParameters) {
	*out = *in
	if in.AccessToken != nil {
		in, out := &in.AccessToken, &out.AccessToken
		*out = new(string)
		**out = **in
	}
	if in.IDToken != nil {
		in, out := &in.IDToken, &out.IDToken
		*out = new(string)
		**out = **in
	}
	if in.RefreshToken != nil {
		in, out := &in.RefreshToken, &out.RefreshToken
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenValidityUnitsParameters.
func (in *TokenValidityUnitsParameters) DeepCopy() *TokenValidityUnitsParameters {
	if in == nil {
		return nil
	}
	out := new(TokenValidityUnitsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolClient) DeepCopyInto(out *UserPoolClient) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolClient.
func (in *UserPoolClient) DeepCopy() *UserPoolClient {
	if in == nil {
		return nil
	}
	out := new(UserPoolClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserPoolClient) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolClientInitParameters) DeepCopyInto(out *UserPoolClientInitParameters) {
	*out = *in
	if in.AccessTokenValidity != nil {
		in, out := &in.AccessTokenValidity, &out.AccessTokenValidity
		*out = new(float64)
		**out = **in
	}
	if in.AllowedOauthFlows != nil {
		in, out := &in.AllowedOauthFlows, &out.AllowedOauthFlows
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedOauthFlowsUserPoolClient != nil {
		in, out := &in.AllowedOauthFlowsUserPoolClient, &out.AllowedOauthFlowsUserPoolClient
		*out = new(bool)
		**out = **in
	}
	if in.AllowedOauthScopes != nil {
		in, out := &in.AllowedOauthScopes, &out.AllowedOauthScopes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AnalyticsConfiguration != nil {
		in, out := &in.AnalyticsConfiguration, &out.AnalyticsConfiguration
		*out = make([]AnalyticsConfigurationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AuthSessionValidity != nil {
		in, out := &in.AuthSessionValidity, &out.AuthSessionValidity
		*out = new(float64)
		**out = **in
	}
	if in.CallbackUrls != nil {
		in, out := &in.CallbackUrls, &out.CallbackUrls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DefaultRedirectURI != nil {
		in, out := &in.DefaultRedirectURI, &out.DefaultRedirectURI
		*out = new(string)
		**out = **in
	}
	if in.EnablePropagateAdditionalUserContextData != nil {
		in, out := &in.EnablePropagateAdditionalUserContextData, &out.EnablePropagateAdditionalUserContextData
		*out = new(bool)
		**out = **in
	}
	if in.EnableTokenRevocation != nil {
		in, out := &in.EnableTokenRevocation, &out.EnableTokenRevocation
		*out = new(bool)
		**out = **in
	}
	if in.ExplicitAuthFlows != nil {
		in, out := &in.ExplicitAuthFlows, &out.ExplicitAuthFlows
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.GenerateSecret != nil {
		in, out := &in.GenerateSecret, &out.GenerateSecret
		*out = new(bool)
		**out = **in
	}
	if in.IDTokenValidity != nil {
		in, out := &in.IDTokenValidity, &out.IDTokenValidity
		*out = new(float64)
		**out = **in
	}
	if in.LogoutUrls != nil {
		in, out := &in.LogoutUrls, &out.LogoutUrls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PreventUserExistenceErrors != nil {
		in, out := &in.PreventUserExistenceErrors, &out.PreventUserExistenceErrors
		*out = new(string)
		**out = **in
	}
	if in.ReadAttributes != nil {
		in, out := &in.ReadAttributes, &out.ReadAttributes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RefreshTokenValidity != nil {
		in, out := &in.RefreshTokenValidity, &out.RefreshTokenValidity
		*out = new(float64)
		**out = **in
	}
	if in.SupportedIdentityProviders != nil {
		in, out := &in.SupportedIdentityProviders, &out.SupportedIdentityProviders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TokenValidityUnits != nil {
		in, out := &in.TokenValidityUnits, &out.TokenValidityUnits
		*out = make([]TokenValidityUnitsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserPoolID != nil {
		in, out := &in.UserPoolID, &out.UserPoolID
		*out = new(string)
		**out = **in
	}
	if in.WriteAttributes != nil {
		in, out := &in.WriteAttributes, &out.WriteAttributes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolClientInitParameters.
func (in *UserPoolClientInitParameters) DeepCopy() *UserPoolClientInitParameters {
	if in == nil {
		return nil
	}
	out := new(UserPoolClientInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolClientList) DeepCopyInto(out *UserPoolClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserPoolClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolClientList.
func (in *UserPoolClientList) DeepCopy() *UserPoolClientList {
	if in == nil {
		return nil
	}
	out := new(UserPoolClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserPoolClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolClientObservation) DeepCopyInto(out *UserPoolClientObservation) {
	*out = *in
	if in.AccessTokenValidity != nil {
		in, out := &in.AccessTokenValidity, &out.AccessTokenValidity
		*out = new(float64)
		**out = **in
	}
	if in.AllowedOauthFlows != nil {
		in, out := &in.AllowedOauthFlows, &out.AllowedOauthFlows
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedOauthFlowsUserPoolClient != nil {
		in, out := &in.AllowedOauthFlowsUserPoolClient, &out.AllowedOauthFlowsUserPoolClient
		*out = new(bool)
		**out = **in
	}
	if in.AllowedOauthScopes != nil {
		in, out := &in.AllowedOauthScopes, &out.AllowedOauthScopes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AnalyticsConfiguration != nil {
		in, out := &in.AnalyticsConfiguration, &out.AnalyticsConfiguration
		*out = make([]AnalyticsConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AuthSessionValidity != nil {
		in, out := &in.AuthSessionValidity, &out.AuthSessionValidity
		*out = new(float64)
		**out = **in
	}
	if in.CallbackUrls != nil {
		in, out := &in.CallbackUrls, &out.CallbackUrls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DefaultRedirectURI != nil {
		in, out := &in.DefaultRedirectURI, &out.DefaultRedirectURI
		*out = new(string)
		**out = **in
	}
	if in.EnablePropagateAdditionalUserContextData != nil {
		in, out := &in.EnablePropagateAdditionalUserContextData, &out.EnablePropagateAdditionalUserContextData
		*out = new(bool)
		**out = **in
	}
	if in.EnableTokenRevocation != nil {
		in, out := &in.EnableTokenRevocation, &out.EnableTokenRevocation
		*out = new(bool)
		**out = **in
	}
	if in.ExplicitAuthFlows != nil {
		in, out := &in.ExplicitAuthFlows, &out.ExplicitAuthFlows
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.GenerateSecret != nil {
		in, out := &in.GenerateSecret, &out.GenerateSecret
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IDTokenValidity != nil {
		in, out := &in.IDTokenValidity, &out.IDTokenValidity
		*out = new(float64)
		**out = **in
	}
	if in.LogoutUrls != nil {
		in, out := &in.LogoutUrls, &out.LogoutUrls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PreventUserExistenceErrors != nil {
		in, out := &in.PreventUserExistenceErrors, &out.PreventUserExistenceErrors
		*out = new(string)
		**out = **in
	}
	if in.ReadAttributes != nil {
		in, out := &in.ReadAttributes, &out.ReadAttributes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RefreshTokenValidity != nil {
		in, out := &in.RefreshTokenValidity, &out.RefreshTokenValidity
		*out = new(float64)
		**out = **in
	}
	if in.SupportedIdentityProviders != nil {
		in, out := &in.SupportedIdentityProviders, &out.SupportedIdentityProviders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TokenValidityUnits != nil {
		in, out := &in.TokenValidityUnits, &out.TokenValidityUnits
		*out = make([]TokenValidityUnitsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserPoolID != nil {
		in, out := &in.UserPoolID, &out.UserPoolID
		*out = new(string)
		**out = **in
	}
	if in.WriteAttributes != nil {
		in, out := &in.WriteAttributes, &out.WriteAttributes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolClientObservation.
func (in *UserPoolClientObservation) DeepCopy() *UserPoolClientObservation {
	if in == nil {
		return nil
	}
	out := new(UserPoolClientObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolClientParameters) DeepCopyInto(out *UserPoolClientParameters) {
	*out = *in
	if in.AccessTokenValidity != nil {
		in, out := &in.AccessTokenValidity, &out.AccessTokenValidity
		*out = new(float64)
		**out = **in
	}
	if in.AllowedOauthFlows != nil {
		in, out := &in.AllowedOauthFlows, &out.AllowedOauthFlows
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedOauthFlowsUserPoolClient != nil {
		in, out := &in.AllowedOauthFlowsUserPoolClient, &out.AllowedOauthFlowsUserPoolClient
		*out = new(bool)
		**out = **in
	}
	if in.AllowedOauthScopes != nil {
		in, out := &in.AllowedOauthScopes, &out.AllowedOauthScopes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AnalyticsConfiguration != nil {
		in, out := &in.AnalyticsConfiguration, &out.AnalyticsConfiguration
		*out = make([]AnalyticsConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AuthSessionValidity != nil {
		in, out := &in.AuthSessionValidity, &out.AuthSessionValidity
		*out = new(float64)
		**out = **in
	}
	if in.CallbackUrls != nil {
		in, out := &in.CallbackUrls, &out.CallbackUrls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DefaultRedirectURI != nil {
		in, out := &in.DefaultRedirectURI, &out.DefaultRedirectURI
		*out = new(string)
		**out = **in
	}
	if in.EnablePropagateAdditionalUserContextData != nil {
		in, out := &in.EnablePropagateAdditionalUserContextData, &out.EnablePropagateAdditionalUserContextData
		*out = new(bool)
		**out = **in
	}
	if in.EnableTokenRevocation != nil {
		in, out := &in.EnableTokenRevocation, &out.EnableTokenRevocation
		*out = new(bool)
		**out = **in
	}
	if in.ExplicitAuthFlows != nil {
		in, out := &in.ExplicitAuthFlows, &out.ExplicitAuthFlows
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.GenerateSecret != nil {
		in, out := &in.GenerateSecret, &out.GenerateSecret
		*out = new(bool)
		**out = **in
	}
	if in.IDTokenValidity != nil {
		in, out := &in.IDTokenValidity, &out.IDTokenValidity
		*out = new(float64)
		**out = **in
	}
	if in.LogoutUrls != nil {
		in, out := &in.LogoutUrls, &out.LogoutUrls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PreventUserExistenceErrors != nil {
		in, out := &in.PreventUserExistenceErrors, &out.PreventUserExistenceErrors
		*out = new(string)
		**out = **in
	}
	if in.ReadAttributes != nil {
		in, out := &in.ReadAttributes, &out.ReadAttributes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RefreshTokenValidity != nil {
		in, out := &in.RefreshTokenValidity, &out.RefreshTokenValidity
		*out = new(float64)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SupportedIdentityProviders != nil {
		in, out := &in.SupportedIdentityProviders, &out.SupportedIdentityProviders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TokenValidityUnits != nil {
		in, out := &in.TokenValidityUnits, &out.TokenValidityUnits
		*out = make([]TokenValidityUnitsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserPoolID != nil {
		in, out := &in.UserPoolID, &out.UserPoolID
		*out = new(string)
		**out = **in
	}
	if in.WriteAttributes != nil {
		in, out := &in.WriteAttributes, &out.WriteAttributes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolClientParameters.
func (in *UserPoolClientParameters) DeepCopy() *UserPoolClientParameters {
	if in == nil {
		return nil
	}
	out := new(UserPoolClientParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolClientSpec) DeepCopyInto(out *UserPoolClientSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolClientSpec.
func (in *UserPoolClientSpec) DeepCopy() *UserPoolClientSpec {
	if in == nil {
		return nil
	}
	out := new(UserPoolClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolClientStatus) DeepCopyInto(out *UserPoolClientStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolClientStatus.
func (in *UserPoolClientStatus) DeepCopy() *UserPoolClientStatus {
	if in == nil {
		return nil
	}
	out := new(UserPoolClientStatus)
	in.DeepCopyInto(out)
	return out
}