/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"github.com/ory/hydra-maester/hydra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OAuth2ClientSpec defines the desired state of OAuth2Client
type OAuth2ClientSpec struct {
	// +kubebuilder:validation:MaxItems=4
	// +kubebuilder:validation:MinItems=1
	//
	// GrantTypes is an array of grant types the client is allowed to use.
	GrantTypes []GrantType `json:"grantTypes"`

	// +kubebuilder:validation:MaxItems=3
	// +kubebuilder:validation:MinItems=1
	//
	// ResponseTypes is an array of the OAuth 2.0 response type strings that the client can
	// use at the authorization endpoint.
	ResponseTypes []ResponseType `json:"responseTypes,omitempty"`

	// +kubebuilder:validation:Pattern=([a-zA-Z0-9\.\*]+\s?)+
	//
	// Scope is a string containing a space-separated list of scope values (as
	// described in Section 3.3 of OAuth 2.0 [RFC6749]) that the client
	// can use when requesting access tokens.
	Scope string `json:"scope"`
}

// +kubebuilder:validation:Enum=client_credentials;authorization_code;implicit;refresh_token
// GrantType represents an OAuth 2.0 grant type
type GrantType string

// +kubebuilder:validation:Enum=id_token;code;token
// ResponseType represents an OAuth 2.0 response type strings
type ResponseType string

// OAuth2ClientStatus defines the observed state of OAuth2Client
type OAuth2ClientStatus struct {
	// Secret points to the K8s secret that contains this client's id and password
	Secret *string `json:"secret,omitempty"`
	// ClientID is the id for this client.
	ClientID *string `json:"clientID,omitempty"`
	// ObservedGeneration represents the most recent generation observed by the daemon set controller.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// OAuth2Client is the Schema for the oauth2clients API
type OAuth2Client struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OAuth2ClientSpec   `json:"spec,omitempty"`
	Status OAuth2ClientStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OAuth2ClientList contains a list of OAuth2Client
type OAuth2ClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OAuth2Client `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OAuth2Client{}, &OAuth2ClientList{})
}

// ToOAuth2ClientJSON converts an OAuth2Client into a OAuth2ClientJSON object that represents an OAuth2 client digestible by ORY Hydra
func (c *OAuth2Client) ToOAuth2ClientJSON() *hydra.OAuth2ClientJSON {
	return &hydra.OAuth2ClientJSON{
		Name:          c.Name,
		ClientID:      c.Status.ClientID,
		GrantTypes:    grantToStringSlice(c.Spec.GrantTypes),
		ResponseTypes: responseToStringSlice(c.Spec.ResponseTypes),
		Scope:         c.Spec.Scope,
	}
}

func responseToStringSlice(rt []ResponseType) []string {
	var output = make([]string, len(rt))
	for i, elem := range rt {
		output[i] = string(elem)
	}
	return output
}

func grantToStringSlice(gt []GrantType) []string {
	var output = make([]string, len(gt))
	for i, elem := range gt {
		output[i] = string(elem)
	}
	return output
}