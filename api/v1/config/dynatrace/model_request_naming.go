/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// RequestNaming The request naming rule.
type RequestNaming struct {
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// The ID of the request naming rule.
	Id *string `json:"id,omitempty"`
	// The order string. Sorting request namings alphabetically by their order string determines their relative ordering.  Typically this is managed by Dynatrace internally and will not be present in GET responses nor used if present in PUT/POST requests, except where noted otherwise.
	Order *string `json:"order,omitempty"`
	// The rule is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	// The name to be assigned to matching requests.
	NamingPattern string `json:"namingPattern"`
	// Specifies the management zones for which this rule should be applied.
	ManagementZones *[]string `json:"managementZones,omitempty"`
	// The set of conditions for the request naming rule usage.    You can specify several conditions. The request has to match **all** the specified conditions for the rule to trigger.
	Conditions []Condition `json:"conditions"`
	// The list of custom placeholders to be used in the naming pattern.    It enables you to extract a request attribute value or other request attribute and use it in the request naming pattern.
	Placeholders *[]Placeholder `json:"placeholders,omitempty"`
}

// NewRequestNaming instantiates a new RequestNaming object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestNaming(enabled bool, namingPattern string, conditions []Condition, ) *RequestNaming {
	this := RequestNaming{}
	this.Enabled = enabled
	this.NamingPattern = namingPattern
	this.Conditions = conditions
	return &this
}

// NewRequestNamingWithDefaults instantiates a new RequestNaming object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestNamingWithDefaults() *RequestNaming {
	this := RequestNaming{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RequestNaming) GetMetadata() ConfigurationMetadata {
	if o == nil || o.Metadata == nil {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestNaming) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RequestNaming) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *RequestNaming) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RequestNaming) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestNaming) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RequestNaming) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RequestNaming) SetId(v string) {
	o.Id = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *RequestNaming) GetOrder() string {
	if o == nil || o.Order == nil {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestNaming) GetOrderOk() (*string, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *RequestNaming) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *RequestNaming) SetOrder(v string) {
	o.Order = &v
}

// GetEnabled returns the Enabled field value
func (o *RequestNaming) GetEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *RequestNaming) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *RequestNaming) SetEnabled(v bool) {
	o.Enabled = v
}

// GetNamingPattern returns the NamingPattern field value
func (o *RequestNaming) GetNamingPattern() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.NamingPattern
}

// GetNamingPatternOk returns a tuple with the NamingPattern field value
// and a boolean to check if the value has been set.
func (o *RequestNaming) GetNamingPatternOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NamingPattern, true
}

// SetNamingPattern sets field value
func (o *RequestNaming) SetNamingPattern(v string) {
	o.NamingPattern = v
}

// GetManagementZones returns the ManagementZones field value if set, zero value otherwise.
func (o *RequestNaming) GetManagementZones() []string {
	if o == nil || o.ManagementZones == nil {
		var ret []string
		return ret
	}
	return *o.ManagementZones
}

// GetManagementZonesOk returns a tuple with the ManagementZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestNaming) GetManagementZonesOk() (*[]string, bool) {
	if o == nil || o.ManagementZones == nil {
		return nil, false
	}
	return o.ManagementZones, true
}

// HasManagementZones returns a boolean if a field has been set.
func (o *RequestNaming) HasManagementZones() bool {
	if o != nil && o.ManagementZones != nil {
		return true
	}

	return false
}

// SetManagementZones gets a reference to the given []string and assigns it to the ManagementZones field.
func (o *RequestNaming) SetManagementZones(v []string) {
	o.ManagementZones = &v
}

// GetConditions returns the Conditions field value
func (o *RequestNaming) GetConditions() []Condition {
	if o == nil  {
		var ret []Condition
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *RequestNaming) GetConditionsOk() (*[]Condition, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Conditions, true
}

// SetConditions sets field value
func (o *RequestNaming) SetConditions(v []Condition) {
	o.Conditions = v
}

// GetPlaceholders returns the Placeholders field value if set, zero value otherwise.
func (o *RequestNaming) GetPlaceholders() []Placeholder {
	if o == nil || o.Placeholders == nil {
		var ret []Placeholder
		return ret
	}
	return *o.Placeholders
}

// GetPlaceholdersOk returns a tuple with the Placeholders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestNaming) GetPlaceholdersOk() (*[]Placeholder, bool) {
	if o == nil || o.Placeholders == nil {
		return nil, false
	}
	return o.Placeholders, true
}

// HasPlaceholders returns a boolean if a field has been set.
func (o *RequestNaming) HasPlaceholders() bool {
	if o != nil && o.Placeholders != nil {
		return true
	}

	return false
}

// SetPlaceholders gets a reference to the given []Placeholder and assigns it to the Placeholders field.
func (o *RequestNaming) SetPlaceholders(v []Placeholder) {
	o.Placeholders = &v
}

func (o RequestNaming) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["namingPattern"] = o.NamingPattern
	}
	if o.ManagementZones != nil {
		toSerialize["managementZones"] = o.ManagementZones
	}
	if true {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Placeholders != nil {
		toSerialize["placeholders"] = o.Placeholders
	}
	return json.Marshal(toSerialize)
}

type NullableRequestNaming struct {
	value *RequestNaming
	isSet bool
}

func (v NullableRequestNaming) Get() *RequestNaming {
	return v.value
}

func (v *NullableRequestNaming) Set(val *RequestNaming) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestNaming) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestNaming) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestNaming(val *RequestNaming) *NullableRequestNaming {
	return &NullableRequestNaming{value: val, isSet: true}
}

func (v NullableRequestNaming) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestNaming) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


