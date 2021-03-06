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

// FullWebRequestRule The service detection rule of the `FULL_WEB_REQUEST` type.
type FullWebRequestRule struct {
	// The type of the service detection rule.
	Type string `json:"type"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// Specifies the management zones of the process group for which this service detection rule should be created.
	ManagementZones *[]string `json:"managementZones,omitempty"`
	// The ID of the service detection rule.
	Id *string `json:"id,omitempty"`
	// The order of the rule in the rules list.    The rules are evaluated from top to bottom. The first matching rule applies.
	Order *string `json:"order,omitempty"`
	// The name of the rule.
	Name string `json:"name"`
	// A short description of the rule.
	Description *string `json:"description,omitempty"`
	// The rule is enabled(`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	// A list of conditions of the rule.   If several conditions are specified, the AND logic applies.
	Conditions *[]ConditionsFullWebRequestAttributeTypeDto `json:"conditions,omitempty"`
	ApplicationId *ApplicationId `json:"applicationId,omitempty"`
	ContextRoot *ContextRoot `json:"contextRoot,omitempty"`
	ServerName *ServerName `json:"serverName,omitempty"`
}

// NewFullWebRequestRule instantiates a new FullWebRequestRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullWebRequestRule(type_ string, name string, enabled bool, ) *FullWebRequestRule {
	this := FullWebRequestRule{}
	this.Type = type_
	this.Name = name
	this.Enabled = enabled
	return &this
}

// NewFullWebRequestRuleWithDefaults instantiates a new FullWebRequestRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullWebRequestRuleWithDefaults() *FullWebRequestRule {
	this := FullWebRequestRule{}
	return &this
}

// GetType returns the Type field value
func (o *FullWebRequestRule) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FullWebRequestRule) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FullWebRequestRule) SetType(v string) {
	o.Type = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *FullWebRequestRule) GetMetadata() ConfigurationMetadata {
	if o == nil || o.Metadata == nil {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullWebRequestRule) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *FullWebRequestRule) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *FullWebRequestRule) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetManagementZones returns the ManagementZones field value if set, zero value otherwise.
func (o *FullWebRequestRule) GetManagementZones() []string {
	if o == nil || o.ManagementZones == nil {
		var ret []string
		return ret
	}
	return *o.ManagementZones
}

// GetManagementZonesOk returns a tuple with the ManagementZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullWebRequestRule) GetManagementZonesOk() (*[]string, bool) {
	if o == nil || o.ManagementZones == nil {
		return nil, false
	}
	return o.ManagementZones, true
}

// HasManagementZones returns a boolean if a field has been set.
func (o *FullWebRequestRule) HasManagementZones() bool {
	if o != nil && o.ManagementZones != nil {
		return true
	}

	return false
}

// SetManagementZones gets a reference to the given []string and assigns it to the ManagementZones field.
func (o *FullWebRequestRule) SetManagementZones(v []string) {
	o.ManagementZones = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FullWebRequestRule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullWebRequestRule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FullWebRequestRule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FullWebRequestRule) SetId(v string) {
	o.Id = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *FullWebRequestRule) GetOrder() string {
	if o == nil || o.Order == nil {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullWebRequestRule) GetOrderOk() (*string, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *FullWebRequestRule) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *FullWebRequestRule) SetOrder(v string) {
	o.Order = &v
}

// GetName returns the Name field value
func (o *FullWebRequestRule) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FullWebRequestRule) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FullWebRequestRule) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FullWebRequestRule) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullWebRequestRule) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FullWebRequestRule) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FullWebRequestRule) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *FullWebRequestRule) GetEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *FullWebRequestRule) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *FullWebRequestRule) SetEnabled(v bool) {
	o.Enabled = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *FullWebRequestRule) GetConditions() []ConditionsFullWebRequestAttributeTypeDto {
	if o == nil || o.Conditions == nil {
		var ret []ConditionsFullWebRequestAttributeTypeDto
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullWebRequestRule) GetConditionsOk() (*[]ConditionsFullWebRequestAttributeTypeDto, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *FullWebRequestRule) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ConditionsFullWebRequestAttributeTypeDto and assigns it to the Conditions field.
func (o *FullWebRequestRule) SetConditions(v []ConditionsFullWebRequestAttributeTypeDto) {
	o.Conditions = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *FullWebRequestRule) GetApplicationId() ApplicationId {
	if o == nil || o.ApplicationId == nil {
		var ret ApplicationId
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullWebRequestRule) GetApplicationIdOk() (*ApplicationId, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *FullWebRequestRule) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given ApplicationId and assigns it to the ApplicationId field.
func (o *FullWebRequestRule) SetApplicationId(v ApplicationId) {
	o.ApplicationId = &v
}

// GetContextRoot returns the ContextRoot field value if set, zero value otherwise.
func (o *FullWebRequestRule) GetContextRoot() ContextRoot {
	if o == nil || o.ContextRoot == nil {
		var ret ContextRoot
		return ret
	}
	return *o.ContextRoot
}

// GetContextRootOk returns a tuple with the ContextRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullWebRequestRule) GetContextRootOk() (*ContextRoot, bool) {
	if o == nil || o.ContextRoot == nil {
		return nil, false
	}
	return o.ContextRoot, true
}

// HasContextRoot returns a boolean if a field has been set.
func (o *FullWebRequestRule) HasContextRoot() bool {
	if o != nil && o.ContextRoot != nil {
		return true
	}

	return false
}

// SetContextRoot gets a reference to the given ContextRoot and assigns it to the ContextRoot field.
func (o *FullWebRequestRule) SetContextRoot(v ContextRoot) {
	o.ContextRoot = &v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *FullWebRequestRule) GetServerName() ServerName {
	if o == nil || o.ServerName == nil {
		var ret ServerName
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullWebRequestRule) GetServerNameOk() (*ServerName, bool) {
	if o == nil || o.ServerName == nil {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *FullWebRequestRule) HasServerName() bool {
	if o != nil && o.ServerName != nil {
		return true
	}

	return false
}

// SetServerName gets a reference to the given ServerName and assigns it to the ServerName field.
func (o *FullWebRequestRule) SetServerName(v ServerName) {
	o.ServerName = &v
}

func (o FullWebRequestRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.ManagementZones != nil {
		toSerialize["managementZones"] = o.ManagementZones
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.ApplicationId != nil {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if o.ContextRoot != nil {
		toSerialize["contextRoot"] = o.ContextRoot
	}
	if o.ServerName != nil {
		toSerialize["serverName"] = o.ServerName
	}
	return json.Marshal(toSerialize)
}

type NullableFullWebRequestRule struct {
	value *FullWebRequestRule
	isSet bool
}

func (v NullableFullWebRequestRule) Get() *FullWebRequestRule {
	return v.value
}

func (v *NullableFullWebRequestRule) Set(val *FullWebRequestRule) {
	v.value = val
	v.isSet = true
}

func (v NullableFullWebRequestRule) IsSet() bool {
	return v.isSet
}

func (v *NullableFullWebRequestRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullWebRequestRule(val *FullWebRequestRule) *NullableFullWebRequestRule {
	return &NullableFullWebRequestRule{value: val, isSet: true}
}

func (v NullableFullWebRequestRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullWebRequestRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


