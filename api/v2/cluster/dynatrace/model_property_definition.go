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

// PropertyDefinition Configuration of a property in a settings schema.
type PropertyDefinition struct {
	Items *Item `json:"items,omitempty"`
	// The type referenced by the property value
	ReferencedType *string `json:"referencedType,omitempty"`
	// An extended description and/or links to documentation.
	Documentation *string `json:"documentation,omitempty"`
	// The maximum number of **objects** in a collection property.    Has the value of `1` for singletons.
	MaxObjects *int32 `json:"maxObjects,omitempty"`
	Precondition *Precondition `json:"precondition,omitempty"`
	// The minimum number of **objects** in a collection property.
	MinObjects *int32 `json:"minObjects,omitempty"`
	// A list of constraints limiting the values to be accepted.
	Constraints *[]Constraint `json:"constraints,omitempty"`
	// The subtype of the property's value.
	SubType *string `json:"subType,omitempty"`
	// A short description of the property.
	Description *string `json:"description,omitempty"`
	// Metadata of the property.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// The value can (`true`) or can't (`false`) be `null`.
	Nullable *bool `json:"nullable,omitempty"`
	// The default value to be used when no value is provided.   If a non-singleton has the value of `null`, it means an empty collection.
	Default *map[string]interface{} `json:"default,omitempty"`
	// The type of the property's value.
	Type *map[string]interface{} `json:"type,omitempty"`
	// The display name of the property.
	DisplayName *string `json:"displayName,omitempty"`
}

// NewPropertyDefinition instantiates a new PropertyDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyDefinition() *PropertyDefinition {
	this := PropertyDefinition{}
	return &this
}

// NewPropertyDefinitionWithDefaults instantiates a new PropertyDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyDefinitionWithDefaults() *PropertyDefinition {
	this := PropertyDefinition{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PropertyDefinition) GetItems() Item {
	if o == nil || o.Items == nil {
		var ret Item
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDefinition) GetItemsOk() (*Item, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PropertyDefinition) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given Item and assigns it to the Items field.
func (o *PropertyDefinition) SetItems(v Item) {
	o.Items = &v
}

// GetReferencedType returns the ReferencedType field value if set, zero value otherwise.
func (o *PropertyDefinition) GetReferencedType() string {
	if o == nil || o.ReferencedType == nil {
		var ret string
		return ret
	}
	return *o.ReferencedType
}

// GetReferencedTypeOk returns a tuple with the ReferencedType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDefinition) GetReferencedTypeOk() (*string, bool) {
	if o == nil || o.ReferencedType == nil {
		return nil, false
	}
	return o.ReferencedType, true
}

// HasReferencedType returns a boolean if a field has been set.
func (o *PropertyDefinition) HasReferencedType() bool {
	if o != nil && o.ReferencedType != nil {
		return true
	}

	return false
}

// SetReferencedType gets a reference to the given string and assigns it to the ReferencedType field.
func (o *PropertyDefinition) SetReferencedType(v string) {
	o.ReferencedType = &v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *PropertyDefinition) GetDocumentation() string {
	if o == nil || o.Documentation == nil {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDefinition) GetDocumentationOk() (*string, bool) {
	if o == nil || o.Documentation == nil {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *PropertyDefinition) HasDocumentation() bool {
	if o != nil && o.Documentation != nil {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *PropertyDefinition) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetMaxObjects returns the MaxObjects field value if set, zero value otherwise.
func (o *PropertyDefinition) GetMaxObjects() int32 {
	if o == nil || o.MaxObjects == nil {
		var ret int32
		return ret
	}
	return *o.MaxObjects
}

// GetMaxObjectsOk returns a tuple with the MaxObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDefinition) GetMaxObjectsOk() (*int32, bool) {
	if o == nil || o.MaxObjects == nil {
		return nil, false
	}
	return o.MaxObjects, true
}

// HasMaxObjects returns a boolean if a field has been set.
func (o *PropertyDefinition) HasMaxObjects() bool {
	if o != nil && o.MaxObjects != nil {
		return true
	}

	return false
}

// SetMaxObjects gets a reference to the given int32 and assigns it to the MaxObjects field.
func (o *PropertyDefinition) SetMaxObjects(v int32) {
	o.MaxObjects = &v
}

// GetPrecondition returns the Precondition field value if set, zero value otherwise.
func (o *PropertyDefinition) GetPrecondition() Precondition {
	if o == nil || o.Precondition == nil {
		var ret Precondition
		return ret
	}
	return *o.Precondition
}

// GetPreconditionOk returns a tuple with the Precondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDefinition) GetPreconditionOk() (*Precondition, bool) {
	if o == nil || o.Precondition == nil {
		return nil, false
	}
	return o.Precondition, true
}

// HasPrecondition returns a boolean if a field has been set.
func (o *PropertyDefinition) HasPrecondition() bool {
	if o != nil && o.Precondition != nil {
		return true
	}

	return false
}

// SetPrecondition gets a reference to the given Precondition and assigns it to the Precondition field.
func (o *PropertyDefinition) SetPrecondition(v Precondition) {
	o.Precondition = &v
}

// GetMinObjects returns the MinObjects field value if set, zero value otherwise.
func (o *PropertyDefinition) GetMinObjects() int32 {
	if o == nil || o.MinObjects == nil {
		var ret int32
		return ret
	}
	return *o.MinObjects
}

// GetMinObjectsOk returns a tuple with the MinObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDefinition) GetMinObjectsOk() (*int32, bool) {
	if o == nil || o.MinObjects == nil {
		return nil, false
	}
	return o.MinObjects, true
}

// HasMinObjects returns a boolean if a field has been set.
func (o *PropertyDefinition) HasMinObjects() bool {
	if o != nil && o.MinObjects != nil {
		return true
	}

	return false
}

// SetMinObjects gets a reference to the given int32 and assigns it to the MinObjects field.
func (o *PropertyDefinition) SetMinObjects(v int32) {
	o.MinObjects = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *PropertyDefinition) GetConstraints() []Constraint {
	if o == nil || o.Constraints == nil {
		var ret []Constraint
		return ret
	}
	return *o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDefinition) GetConstraintsOk() (*[]Constraint, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *PropertyDefinition) HasConstraints() bool {
	if o != nil && o.Constraints != nil {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []Constraint and assigns it to the Constraints field.
func (o *PropertyDefinition) SetConstraints(v []Constraint) {
	o.Constraints = &v
}

// GetSubType returns the SubType field value if set, zero value otherwise.
func (o *PropertyDefinition) GetSubType() string {
	if o == nil || o.SubType == nil {
		var ret string
		return ret
	}
	return *o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDefinition) GetSubTypeOk() (*string, bool) {
	if o == nil || o.SubType == nil {
		return nil, false
	}
	return o.SubType, true
}

// HasSubType returns a boolean if a field has been set.
func (o *PropertyDefinition) HasSubType() bool {
	if o != nil && o.SubType != nil {
		return true
	}

	return false
}

// SetSubType gets a reference to the given string and assigns it to the SubType field.
func (o *PropertyDefinition) SetSubType(v string) {
	o.SubType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PropertyDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PropertyDefinition) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PropertyDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PropertyDefinition) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDefinition) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PropertyDefinition) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *PropertyDefinition) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetNullable returns the Nullable field value if set, zero value otherwise.
func (o *PropertyDefinition) GetNullable() bool {
	if o == nil || o.Nullable == nil {
		var ret bool
		return ret
	}
	return *o.Nullable
}

// GetNullableOk returns a tuple with the Nullable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDefinition) GetNullableOk() (*bool, bool) {
	if o == nil || o.Nullable == nil {
		return nil, false
	}
	return o.Nullable, true
}

// HasNullable returns a boolean if a field has been set.
func (o *PropertyDefinition) HasNullable() bool {
	if o != nil && o.Nullable != nil {
		return true
	}

	return false
}

// SetNullable gets a reference to the given bool and assigns it to the Nullable field.
func (o *PropertyDefinition) SetNullable(v bool) {
	o.Nullable = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *PropertyDefinition) GetDefault() map[string]interface{} {
	if o == nil || o.Default == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDefinition) GetDefaultOk() (*map[string]interface{}, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *PropertyDefinition) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given map[string]interface{} and assigns it to the Default field.
func (o *PropertyDefinition) SetDefault(v map[string]interface{}) {
	o.Default = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PropertyDefinition) GetType() map[string]interface{} {
	if o == nil || o.Type == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDefinition) GetTypeOk() (*map[string]interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PropertyDefinition) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given map[string]interface{} and assigns it to the Type field.
func (o *PropertyDefinition) SetType(v map[string]interface{}) {
	o.Type = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PropertyDefinition) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDefinition) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PropertyDefinition) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PropertyDefinition) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o PropertyDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.ReferencedType != nil {
		toSerialize["referencedType"] = o.ReferencedType
	}
	if o.Documentation != nil {
		toSerialize["documentation"] = o.Documentation
	}
	if o.MaxObjects != nil {
		toSerialize["maxObjects"] = o.MaxObjects
	}
	if o.Precondition != nil {
		toSerialize["precondition"] = o.Precondition
	}
	if o.MinObjects != nil {
		toSerialize["minObjects"] = o.MinObjects
	}
	if o.Constraints != nil {
		toSerialize["constraints"] = o.Constraints
	}
	if o.SubType != nil {
		toSerialize["subType"] = o.SubType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Nullable != nil {
		toSerialize["nullable"] = o.Nullable
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyDefinition struct {
	value *PropertyDefinition
	isSet bool
}

func (v NullablePropertyDefinition) Get() *PropertyDefinition {
	return v.value
}

func (v *NullablePropertyDefinition) Set(val *PropertyDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyDefinition(val *PropertyDefinition) *NullablePropertyDefinition {
	return &NullablePropertyDefinition{value: val, isSet: true}
}

func (v NullablePropertyDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


