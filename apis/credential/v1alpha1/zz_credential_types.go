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

type CredentialInitParameters struct {

	// The domain to add credential of Mailgun.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The local-part of the email address to create.
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// The region where domain will be created. Default value is us.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type CredentialObservation struct {

	// The domain to add credential of Mailgun.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The local-part of the email address to create.
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// The region where domain will be created. Default value is us.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type CredentialParameters struct {

	// The domain to add credential of Mailgun.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The local-part of the email address to create.
	// +kubebuilder:validation:Optional
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// Password for user authentication.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The region where domain will be created. Default value is us.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// CredentialSpec defines the desired state of Credential
type CredentialSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CredentialParameters `json:"forProvider"`
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
	InitProvider CredentialInitParameters `json:"initProvider,omitempty"`
}

// CredentialStatus defines the observed state of Credential.
type CredentialStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CredentialObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Credential is the Schema for the Credentials API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mailgun}
type Credential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domain) || (has(self.initProvider) && has(self.initProvider.domain))",message="spec.forProvider.domain is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.login) || (has(self.initProvider) && has(self.initProvider.login))",message="spec.forProvider.login is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="spec.forProvider.passwordSecretRef is a required parameter"
	Spec   CredentialSpec   `json:"spec"`
	Status CredentialStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CredentialList contains a list of Credentials
type CredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Credential `json:"items"`
}

// Repository type metadata.
var (
	Credential_Kind             = "Credential"
	Credential_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Credential_Kind}.String()
	Credential_KindAPIVersion   = Credential_Kind + "." + CRDGroupVersion.String()
	Credential_GroupVersionKind = CRDGroupVersion.WithKind(Credential_Kind)
)

func init() {
	SchemeBuilder.Register(&Credential{}, &CredentialList{})
}