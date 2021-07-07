/*
 * Dynatrace Cluster API
 *
 * Dynatrace Managed exposes cluster-wide functionality via REST endpoints. This interactive documentation also acts as a REST client you can use to interact with Dynatrace Managed clusters.   To authorize, use the API Token generated in [Settings - API Tokens page](/cmc#cm/apiToken). The HTTP status code of the response shows the result of your request. The expected response code for a successful request is documented individually per REST endpoint. Additionally the following error response codes can occur in our REST interface:  * 400 - Bad Request: Some request parameters are not correct. See response body for details. * 401 - Unauthorized: A valid authorization header is required but is missing. * 403 - Forbidden: Execution of request is not allowed, e.g. api-token is invalid. * 404 - Not Found: Endpoint does not exist or some entities could not be found. * 500 - Internal Server Error: See response body for details. * 556 - Upgrade in progress: Operation couldn't be performed during the upgrade.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// CreateToken struct for CreateToken
type CreateToken struct {
	// The name of the token.
	Name string `json:"name"`
	ExpiresIn *Duration `json:"expiresIn,omitempty"`
	// The list of scopes to be assigned to the token.  * `DiagnosticExport`: DiagnosticExport.  * `ControlManagement`: ControlManagement.  * `UnattendedInstall`: UnattendedInstall.  * `ServiceProviderAPI`: Service Provider API.  * `ExternalSyntheticIntegration`: Create and read synthetic monitors, locations, and nodes.  * `ClusterTokenManagement`: Cluster token management.  * `ReadSyntheticData`: Read synthetic monitors, locations, and nodes.  * `Nodekeeper`: Nodekeeper access for node management.  * `EnvironmentTokenManagement`: \"Token Management\" Token creation for existing environments.  * `settings.read`: Read settings.  * `settings.write`: Write settings.  * `apiTokens.read`: Read API tokens.  * `apiTokens.write`: Write API tokens.  
	Scopes []string `json:"scopes"`
}

// NewCreateToken instantiates a new CreateToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateToken(name string, scopes []string) *CreateToken {
	this := CreateToken{}
	this.Name = name
	this.Scopes = scopes
	return &this
}

// NewCreateTokenWithDefaults instantiates a new CreateToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTokenWithDefaults() *CreateToken {
	this := CreateToken{}
	return &this
}

// GetName returns the Name field value
func (o *CreateToken) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateToken) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateToken) SetName(v string) {
	o.Name = v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *CreateToken) GetExpiresIn() Duration {
	if o == nil || o.ExpiresIn == nil {
		var ret Duration
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateToken) GetExpiresInOk() (*Duration, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *CreateToken) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given Duration and assigns it to the ExpiresIn field.
func (o *CreateToken) SetExpiresIn(v Duration) {
	o.ExpiresIn = &v
}

// GetScopes returns the Scopes field value
func (o *CreateToken) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *CreateToken) GetScopesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Scopes, true
}

// SetScopes sets field value
func (o *CreateToken) SetScopes(v []string) {
	o.Scopes = v
}

func (o CreateToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ExpiresIn != nil {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	if true {
		toSerialize["scopes"] = o.Scopes
	}
	return json.Marshal(toSerialize)
}

type NullableCreateToken struct {
	value *CreateToken
	isSet bool
}

func (v NullableCreateToken) Get() *CreateToken {
	return v.value
}

func (v *NullableCreateToken) Set(val *CreateToken) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateToken) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateToken(val *CreateToken) *NullableCreateToken {
	return &NullableCreateToken{value: val, isSet: true}
}

func (v NullableCreateToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


