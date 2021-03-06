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

// Placeholder The custom placeholder to be used as a naming value pattern.    It enables you to extract a request attribute value or other request attribute and use it in the naming pattern.
type Placeholder struct {
	// The name of the placeholder. Use it in the naming pattern as `{name}`.
	Name string `json:"name"`
	// The attribute to extract from. You can only use attributes of the **string** type.
	Attribute string `json:"attribute"`
	// The type of extraction.    Defines either usage of regular expression (`regex`) or the position of request attribute value to be extracted.   When the **attribute** is `SERVICE_REQUEST_ATTRIBUTE` attribute and **aggregation** is `COUNT`, needs to be set to `ORIGINAL_TEXT`
	Kind string `json:"kind"`
	// Depending on the **type** value:    * `REGEX_EXTRACTION`: The regular expression.   * `BETWEEN_DELIMITER`: The opening delimiter string to look for.   * All other values: The delimiter string to look for.
	DelimiterOrRegex *string `json:"delimiterOrRegex,omitempty"`
	// The closing delimiter string to look for.    Required if the **kind** value is `BETWEEN_DELIMITER`. Not applicable otherwise.
	EndDelimiter *string `json:"endDelimiter,omitempty"`
	// The request attribute to extract from.    Required if the **kind** value is `SERVICE_REQUEST_ATTRIBUTE`. Not applicable otherwise.
	RequestAttribute *string `json:"requestAttribute,omitempty"`
	// The format of the extracted string.
	Normalization *string `json:"normalization,omitempty"`
	// If `true` request attribute will be taken from a child service call.    Only applicable for the `SERVICE_REQUEST_ATTRIBUTE` attribute. Defaults to `false`.
	UseFromChildCalls *bool `json:"useFromChildCalls,omitempty"`
	// Which value of the request attribute must be used when it occurs across multiple child requests.   Only applicable for the `SERVICE_REQUEST_ATTRIBUTE` attribute, when **useFromChildCalls** is `true`.   For the `COUNT` aggregation, the **kind** field is not applicable.
	Aggregation *string `json:"aggregation,omitempty"`
	Source *PropagationSource `json:"source,omitempty"`
}

// NewPlaceholder instantiates a new Placeholder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceholder(name string, attribute string, kind string, ) *Placeholder {
	this := Placeholder{}
	this.Name = name
	this.Attribute = attribute
	this.Kind = kind
	return &this
}

// NewPlaceholderWithDefaults instantiates a new Placeholder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceholderWithDefaults() *Placeholder {
	this := Placeholder{}
	return &this
}

// GetName returns the Name field value
func (o *Placeholder) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Placeholder) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Placeholder) SetName(v string) {
	o.Name = v
}

// GetAttribute returns the Attribute field value
func (o *Placeholder) GetAttribute() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *Placeholder) GetAttributeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *Placeholder) SetAttribute(v string) {
	o.Attribute = v
}

// GetKind returns the Kind field value
func (o *Placeholder) GetKind() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *Placeholder) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *Placeholder) SetKind(v string) {
	o.Kind = v
}

// GetDelimiterOrRegex returns the DelimiterOrRegex field value if set, zero value otherwise.
func (o *Placeholder) GetDelimiterOrRegex() string {
	if o == nil || o.DelimiterOrRegex == nil {
		var ret string
		return ret
	}
	return *o.DelimiterOrRegex
}

// GetDelimiterOrRegexOk returns a tuple with the DelimiterOrRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Placeholder) GetDelimiterOrRegexOk() (*string, bool) {
	if o == nil || o.DelimiterOrRegex == nil {
		return nil, false
	}
	return o.DelimiterOrRegex, true
}

// HasDelimiterOrRegex returns a boolean if a field has been set.
func (o *Placeholder) HasDelimiterOrRegex() bool {
	if o != nil && o.DelimiterOrRegex != nil {
		return true
	}

	return false
}

// SetDelimiterOrRegex gets a reference to the given string and assigns it to the DelimiterOrRegex field.
func (o *Placeholder) SetDelimiterOrRegex(v string) {
	o.DelimiterOrRegex = &v
}

// GetEndDelimiter returns the EndDelimiter field value if set, zero value otherwise.
func (o *Placeholder) GetEndDelimiter() string {
	if o == nil || o.EndDelimiter == nil {
		var ret string
		return ret
	}
	return *o.EndDelimiter
}

// GetEndDelimiterOk returns a tuple with the EndDelimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Placeholder) GetEndDelimiterOk() (*string, bool) {
	if o == nil || o.EndDelimiter == nil {
		return nil, false
	}
	return o.EndDelimiter, true
}

// HasEndDelimiter returns a boolean if a field has been set.
func (o *Placeholder) HasEndDelimiter() bool {
	if o != nil && o.EndDelimiter != nil {
		return true
	}

	return false
}

// SetEndDelimiter gets a reference to the given string and assigns it to the EndDelimiter field.
func (o *Placeholder) SetEndDelimiter(v string) {
	o.EndDelimiter = &v
}

// GetRequestAttribute returns the RequestAttribute field value if set, zero value otherwise.
func (o *Placeholder) GetRequestAttribute() string {
	if o == nil || o.RequestAttribute == nil {
		var ret string
		return ret
	}
	return *o.RequestAttribute
}

// GetRequestAttributeOk returns a tuple with the RequestAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Placeholder) GetRequestAttributeOk() (*string, bool) {
	if o == nil || o.RequestAttribute == nil {
		return nil, false
	}
	return o.RequestAttribute, true
}

// HasRequestAttribute returns a boolean if a field has been set.
func (o *Placeholder) HasRequestAttribute() bool {
	if o != nil && o.RequestAttribute != nil {
		return true
	}

	return false
}

// SetRequestAttribute gets a reference to the given string and assigns it to the RequestAttribute field.
func (o *Placeholder) SetRequestAttribute(v string) {
	o.RequestAttribute = &v
}

// GetNormalization returns the Normalization field value if set, zero value otherwise.
func (o *Placeholder) GetNormalization() string {
	if o == nil || o.Normalization == nil {
		var ret string
		return ret
	}
	return *o.Normalization
}

// GetNormalizationOk returns a tuple with the Normalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Placeholder) GetNormalizationOk() (*string, bool) {
	if o == nil || o.Normalization == nil {
		return nil, false
	}
	return o.Normalization, true
}

// HasNormalization returns a boolean if a field has been set.
func (o *Placeholder) HasNormalization() bool {
	if o != nil && o.Normalization != nil {
		return true
	}

	return false
}

// SetNormalization gets a reference to the given string and assigns it to the Normalization field.
func (o *Placeholder) SetNormalization(v string) {
	o.Normalization = &v
}

// GetUseFromChildCalls returns the UseFromChildCalls field value if set, zero value otherwise.
func (o *Placeholder) GetUseFromChildCalls() bool {
	if o == nil || o.UseFromChildCalls == nil {
		var ret bool
		return ret
	}
	return *o.UseFromChildCalls
}

// GetUseFromChildCallsOk returns a tuple with the UseFromChildCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Placeholder) GetUseFromChildCallsOk() (*bool, bool) {
	if o == nil || o.UseFromChildCalls == nil {
		return nil, false
	}
	return o.UseFromChildCalls, true
}

// HasUseFromChildCalls returns a boolean if a field has been set.
func (o *Placeholder) HasUseFromChildCalls() bool {
	if o != nil && o.UseFromChildCalls != nil {
		return true
	}

	return false
}

// SetUseFromChildCalls gets a reference to the given bool and assigns it to the UseFromChildCalls field.
func (o *Placeholder) SetUseFromChildCalls(v bool) {
	o.UseFromChildCalls = &v
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *Placeholder) GetAggregation() string {
	if o == nil || o.Aggregation == nil {
		var ret string
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Placeholder) GetAggregationOk() (*string, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *Placeholder) HasAggregation() bool {
	if o != nil && o.Aggregation != nil {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given string and assigns it to the Aggregation field.
func (o *Placeholder) SetAggregation(v string) {
	o.Aggregation = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Placeholder) GetSource() PropagationSource {
	if o == nil || o.Source == nil {
		var ret PropagationSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Placeholder) GetSourceOk() (*PropagationSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Placeholder) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given PropagationSource and assigns it to the Source field.
func (o *Placeholder) SetSource(v PropagationSource) {
	o.Source = &v
}

func (o Placeholder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["attribute"] = o.Attribute
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if o.DelimiterOrRegex != nil {
		toSerialize["delimiterOrRegex"] = o.DelimiterOrRegex
	}
	if o.EndDelimiter != nil {
		toSerialize["endDelimiter"] = o.EndDelimiter
	}
	if o.RequestAttribute != nil {
		toSerialize["requestAttribute"] = o.RequestAttribute
	}
	if o.Normalization != nil {
		toSerialize["normalization"] = o.Normalization
	}
	if o.UseFromChildCalls != nil {
		toSerialize["useFromChildCalls"] = o.UseFromChildCalls
	}
	if o.Aggregation != nil {
		toSerialize["aggregation"] = o.Aggregation
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullablePlaceholder struct {
	value *Placeholder
	isSet bool
}

func (v NullablePlaceholder) Get() *Placeholder {
	return v.value
}

func (v *NullablePlaceholder) Set(val *Placeholder) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceholder) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceholder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceholder(val *Placeholder) *NullablePlaceholder {
	return &NullablePlaceholder{value: val, isSet: true}
}

func (v NullablePlaceholder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceholder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


