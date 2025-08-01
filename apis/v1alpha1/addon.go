// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AddonSpec defines the desired state of Addon.
//
// An Amazon EKS add-on. For more information, see Amazon EKS add-ons (https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html)
// in the Amazon EKS User Guide.
type AddonSpec struct {

	// The version of the add-on. The version must match one of the versions returned
	// by DescribeAddonVersions (https://docs.aws.amazon.com/eks/latest/APIReference/API_DescribeAddonVersions.html).
	AddonVersion *string `json:"addonVersion,omitempty"`
	// A unique, case-sensitive identifier that you provide to ensurethe idempotency
	// of the request.
	ClientRequestToken *string `json:"clientRequestToken,omitempty"`
	// The name of your cluster.
	//
	// Regex Pattern: `^[0-9A-Za-z][A-Za-z0-9\-_]*$`
	ClusterName *string                                  `json:"clusterName,omitempty"`
	ClusterRef  *ackv1alpha1.AWSResourceReferenceWrapper `json:"clusterRef,omitempty"`
	// The set of configuration values for the add-on that's created. The values
	// that you provide are validated against the schema returned by DescribeAddonConfiguration.
	ConfigurationValues *string `json:"configurationValues,omitempty"`
	// The name of the add-on. The name must match one of the names returned by
	// DescribeAddonVersions.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// An array of Pod Identity Assocations to be created. Each EKS Pod Identity
	// association maps a Kubernetes service account to an IAM Role.
	//
	// For more information, see Attach an IAM Role to an Amazon EKS add-on using
	// Pod Identity (https://docs.aws.amazon.com/eks/latest/userguide/add-ons-iam.html)
	// in the Amazon EKS User Guide.
	PodIdentityAssociations []*AddonPodIdentityAssociations `json:"podIdentityAssociations,omitempty"`
	// How to resolve field value conflicts for an Amazon EKS add-on. Conflicts
	// are handled based on the value you choose:
	//
	//   - None – If the self-managed version of the add-on is installed on your
	//     cluster, Amazon EKS doesn't change the value. Creation of the add-on might
	//     fail.
	//
	//   - Overwrite – If the self-managed version of the add-on is installed
	//     on your cluster and the Amazon EKS default value is different than the
	//     existing value, Amazon EKS changes the value to the Amazon EKS default
	//     value.
	//
	//   - Preserve – This is similar to the NONE option. If the self-managed
	//     version of the add-on is installed on your cluster Amazon EKS doesn't
	//     change the add-on resource properties. Creation of the add-on might fail
	//     if conflicts are detected. This option works differently during the update
	//     operation. For more information, see UpdateAddon (https://docs.aws.amazon.com/eks/latest/APIReference/API_UpdateAddon.html).
	//
	// If you don't currently have the self-managed version of the add-on installed
	// on your cluster, the Amazon EKS add-on is installed. Amazon EKS sets all
	// values to default values, regardless of the option that you specify.
	ResolveConflicts *string `json:"resolveConflicts,omitempty"`
	// The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's
	// service account. The role must be assigned the IAM permissions required by
	// the add-on. If you don't specify an existing IAM role, then the add-on uses
	// the permissions assigned to the node IAM role. For more information, see
	// Amazon EKS node IAM role (https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html)
	// in the Amazon EKS User Guide.
	//
	// To specify an existing IAM role, you must have an IAM OpenID Connect (OIDC)
	// provider created for your cluster. For more information, see Enabling IAM
	// roles for service accounts on your cluster (https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html)
	// in the Amazon EKS User Guide.
	ServiceAccountRoleARN *string                                  `json:"serviceAccountRoleARN,omitempty"`
	ServiceAccountRoleRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"serviceAccountRoleRef,omitempty"`
	// Metadata that assists with categorization and organization. Each tag consists
	// of a key and an optional value. You define both. Tags don't propagate to
	// any other cluster or Amazon Web Services resources.
	Tags map[string]*string `json:"tags,omitempty"`
}

// AddonStatus defines the observed state of Addon
type AddonStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The Unix epoch timestamp at object creation.
	// +kubebuilder:validation:Optional
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
	// An object that represents the health of the add-on.
	// +kubebuilder:validation:Optional
	Health *AddonHealth `json:"health,omitempty"`
	// Information about an Amazon EKS add-on from the Amazon Web Services Marketplace.
	// +kubebuilder:validation:Optional
	MarketplaceInformation *MarketplaceInformation `json:"marketplaceInformation,omitempty"`
	// The Unix epoch timestamp for the last modification to the object.
	// +kubebuilder:validation:Optional
	ModifiedAt *metav1.Time `json:"modifiedAt,omitempty"`
	// The owner of the add-on.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty"`
	// The publisher of the add-on.
	// +kubebuilder:validation:Optional
	Publisher *string `json:"publisher,omitempty"`
	// The status of the add-on.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty"`
}

// Addon is the Schema for the Addons API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="CLUSTER",type=string,priority=0,JSONPath=`.spec.clusterName`
// +kubebuilder:printcolumn:name="ADDONVERSION",type=string,priority=0,JSONPath=`.spec.addonVersion`
// +kubebuilder:printcolumn:name="RESOLVECONFLICTS",type=string,priority=0,JSONPath=`.spec.resolveConflicts`
// +kubebuilder:printcolumn:name="STATUS",type=string,priority=1,JSONPath=`.status.status`
// +kubebuilder:printcolumn:name="SERVICEACCOUNTROLEARN",type=string,priority=1,JSONPath=`.spec.serviceAccountRoleARN`
// +kubebuilder:printcolumn:name="Synced",type="string",priority=0,JSONPath=".status.conditions[?(@.type==\"ACK.ResourceSynced\")].status"
// +kubebuilder:printcolumn:name="Age",type="date",priority=0,JSONPath=".metadata.creationTimestamp"
type Addon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AddonSpec   `json:"spec,omitempty"`
	Status            AddonStatus `json:"status,omitempty"`
}

// AddonList contains a list of Addon
// +kubebuilder:object:root=true
type AddonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Addon `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Addon{}, &AddonList{})
}
