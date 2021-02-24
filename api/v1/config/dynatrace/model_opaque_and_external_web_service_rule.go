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

// OpaqueAndExternalWebServiceRule The service detection rule of the `OPAQUE_AND_EXTERNAL_WEB_SERVICE` type
type OpaqueAndExternalWebServiceRule struct {
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
	Conditions *[]ConditionsOpaqueAndExternalWebServiceAttributeTypeDto `json:"conditions,omitempty"`
	// Detect the matching requests as web services (`false`) or web request services (`true`).   Setting this field to `true` prevents detecting of matching requests as opaque web services. An opaque web request service is created instead. If you need to further modify the resulting web request service, you need to create a separate rule of the `OPAQUE_AND_EXTERNAL_WEB_REQUEST` type.   Default is `false`, detecting matching requests as opaque web services.
	DetectAsWebRequestService *bool `json:"detectAsWebRequestService,omitempty"`
	UrlPath *UrlPath `json:"urlPath,omitempty"`
	Port *Port `json:"port,omitempty"`
}

// NewOpaqueAndExternalWebServiceRule instantiates a new OpaqueAndExternalWebServiceRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpaqueAndExternalWebServiceRule(type_ string, name string, enabled bool, ) *OpaqueAndExternalWebServiceRule {
	this := OpaqueAndExternalWebServiceRule{}
	this.Type = type_
	this.Name = name
	this.Enabled = enabled
	return &this
}

// NewOpaqueAndExternalWebServiceRuleWithDefaults instantiates a new OpaqueAndExternalWebServiceRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpaqueAndExternalWebServiceRuleWithDefaults() *OpaqueAndExternalWebServiceRule {
	this := OpaqueAndExternalWebServiceRule{}
	return &this
}

// GetType returns the Type field value
func (o *OpaqueAndExternalWebServiceRule) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebServiceRule) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OpaqueAndExternalWebServiceRule) SetType(v string) {
	o.Type = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebServiceRule) GetMetadata() ConfigurationMetadata {
	if o == nil || o.Metadata == nil {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebServiceRule) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebServiceRule) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *OpaqueAndExternalWebServiceRule) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetManagementZones returns the ManagementZones field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebServiceRule) GetManagementZones() []string {
	if o == nil || o.ManagementZones == nil {
		var ret []string
		return ret
	}
	return *o.ManagementZones
}

// GetManagementZonesOk returns a tuple with the ManagementZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebServiceRule) GetManagementZonesOk() (*[]string, bool) {
	if o == nil || o.ManagementZones == nil {
		return nil, false
	}
	return o.ManagementZones, true
}

// HasManagementZones returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebServiceRule) HasManagementZones() bool {
	if o != nil && o.ManagementZones != nil {
		return true
	}

	return false
}

// SetManagementZones gets a reference to the given []string and assigns it to the ManagementZones field.
func (o *OpaqueAndExternalWebServiceRule) SetManagementZones(v []string) {
	o.ManagementZones = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebServiceRule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebServiceRule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebServiceRule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OpaqueAndExternalWebServiceRule) SetId(v string) {
	o.Id = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebServiceRule) GetOrder() string {
	if o == nil || o.Order == nil {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebServiceRule) GetOrderOk() (*string, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebServiceRule) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *OpaqueAndExternalWebServiceRule) SetOrder(v string) {
	o.Order = &v
}

// GetName returns the Name field value
func (o *OpaqueAndExternalWebServiceRule) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebServiceRule) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OpaqueAndExternalWebServiceRule) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebServiceRule) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebServiceRule) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebServiceRule) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OpaqueAndExternalWebServiceRule) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *OpaqueAndExternalWebServiceRule) GetEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebServiceRule) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *OpaqueAndExternalWebServiceRule) SetEnabled(v bool) {
	o.Enabled = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebServiceRule) GetConditions() []ConditionsOpaqueAndExternalWebServiceAttributeTypeDto {
	if o == nil || o.Conditions == nil {
		var ret []ConditionsOpaqueAndExternalWebServiceAttributeTypeDto
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebServiceRule) GetConditionsOk() (*[]ConditionsOpaqueAndExternalWebServiceAttributeTypeDto, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebServiceRule) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ConditionsOpaqueAndExternalWebServiceAttributeTypeDto and assigns it to the Conditions field.
func (o *OpaqueAndExternalWebServiceRule) SetConditions(v []ConditionsOpaqueAndExternalWebServiceAttributeTypeDto) {
	o.Conditions = &v
}

// GetDetectAsWebRequestService returns the DetectAsWebRequestService field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebServiceRule) GetDetectAsWebRequestService() bool {
	if o == nil || o.DetectAsWebRequestService == nil {
		var ret bool
		return ret
	}
	return *o.DetectAsWebRequestService
}

// GetDetectAsWebRequestServiceOk returns a tuple with the DetectAsWebRequestService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebServiceRule) GetDetectAsWebRequestServiceOk() (*bool, bool) {
	if o == nil || o.DetectAsWebRequestService == nil {
		return nil, false
	}
	return o.DetectAsWebRequestService, true
}

// HasDetectAsWebRequestService returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebServiceRule) HasDetectAsWebRequestService() bool {
	if o != nil && o.DetectAsWebRequestService != nil {
		return true
	}

	return false
}

// SetDetectAsWebRequestService gets a reference to the given bool and assigns it to the DetectAsWebRequestService field.
func (o *OpaqueAndExternalWebServiceRule) SetDetectAsWebRequestService(v bool) {
	o.DetectAsWebRequestService = &v
}

// GetUrlPath returns the UrlPath field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebServiceRule) GetUrlPath() UrlPath {
	if o == nil || o.UrlPath == nil {
		var ret UrlPath
		return ret
	}
	return *o.UrlPath
}

// GetUrlPathOk returns a tuple with the UrlPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebServiceRule) GetUrlPathOk() (*UrlPath, bool) {
	if o == nil || o.UrlPath == nil {
		return nil, false
	}
	return o.UrlPath, true
}

// HasUrlPath returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebServiceRule) HasUrlPath() bool {
	if o != nil && o.UrlPath != nil {
		return true
	}

	return false
}

// SetUrlPath gets a reference to the given UrlPath and assigns it to the UrlPath field.
func (o *OpaqueAndExternalWebServiceRule) SetUrlPath(v UrlPath) {
	o.UrlPath = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebServiceRule) GetPort() Port {
	if o == nil || o.Port == nil {
		var ret Port
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebServiceRule) GetPortOk() (*Port, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebServiceRule) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given Port and assigns it to the Port field.
func (o *OpaqueAndExternalWebServiceRule) SetPort(v Port) {
	o.Port = &v
}

func (o OpaqueAndExternalWebServiceRule) MarshalJSON() ([]byte, error) {
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
	if o.DetectAsWebRequestService != nil {
		toSerialize["detectAsWebRequestService"] = o.DetectAsWebRequestService
	}
	if o.UrlPath != nil {
		toSerialize["urlPath"] = o.UrlPath
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableOpaqueAndExternalWebServiceRule struct {
	value *OpaqueAndExternalWebServiceRule
	isSet bool
}

func (v NullableOpaqueAndExternalWebServiceRule) Get() *OpaqueAndExternalWebServiceRule {
	return v.value
}

func (v *NullableOpaqueAndExternalWebServiceRule) Set(val *OpaqueAndExternalWebServiceRule) {
	v.value = val
	v.isSet = true
}

func (v NullableOpaqueAndExternalWebServiceRule) IsSet() bool {
	return v.isSet
}

func (v *NullableOpaqueAndExternalWebServiceRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpaqueAndExternalWebServiceRule(val *OpaqueAndExternalWebServiceRule) *NullableOpaqueAndExternalWebServiceRule {
	return &NullableOpaqueAndExternalWebServiceRule{value: val, isSet: true}
}

func (v NullableOpaqueAndExternalWebServiceRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpaqueAndExternalWebServiceRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


