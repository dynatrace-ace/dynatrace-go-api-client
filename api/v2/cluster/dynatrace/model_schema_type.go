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

// SchemaType A list of definitions of types.    A type is a complex property that contains its own set of subproperties.
type SchemaType struct {
	// The pattern for the summary (for example, \"Alert after *X* minutes.\") of the configuration in the UI.
	SummaryPattern *string `json:"summaryPattern,omitempty"`
	// A list of constraints limiting the values to be accepted.
	Constraints *[]ComplexConstraint `json:"constraints,omitempty"`
	// A short description of the version.
	VersionInfo *string `json:"versionInfo,omitempty"`
	// The version of the type.
	Version *string `json:"version,omitempty"`
	// Definition of properties that can be persisted.
	Properties *map[string]PropertyDefinition `json:"properties,omitempty"`
	// An extended description and/or links to documentation.
	Documentation *string `json:"documentation,omitempty"`
	// A short description of the property.
	Description *string `json:"description,omitempty"`
	// The display name of the property.
	DisplayName *string `json:"displayName,omitempty"`
}

// NewSchemaType instantiates a new SchemaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaType() *SchemaType {
	this := SchemaType{}
	return &this
}

// NewSchemaTypeWithDefaults instantiates a new SchemaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaTypeWithDefaults() *SchemaType {
	this := SchemaType{}
	return &this
}

// GetSummaryPattern returns the SummaryPattern field value if set, zero value otherwise.
func (o *SchemaType) GetSummaryPattern() string {
	if o == nil || o.SummaryPattern == nil {
		var ret string
		return ret
	}
	return *o.SummaryPattern
}

// GetSummaryPatternOk returns a tuple with the SummaryPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetSummaryPatternOk() (*string, bool) {
	if o == nil || o.SummaryPattern == nil {
		return nil, false
	}
	return o.SummaryPattern, true
}

// HasSummaryPattern returns a boolean if a field has been set.
func (o *SchemaType) HasSummaryPattern() bool {
	if o != nil && o.SummaryPattern != nil {
		return true
	}

	return false
}

// SetSummaryPattern gets a reference to the given string and assigns it to the SummaryPattern field.
func (o *SchemaType) SetSummaryPattern(v string) {
	o.SummaryPattern = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *SchemaType) GetConstraints() []ComplexConstraint {
	if o == nil || o.Constraints == nil {
		var ret []ComplexConstraint
		return ret
	}
	return *o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetConstraintsOk() (*[]ComplexConstraint, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *SchemaType) HasConstraints() bool {
	if o != nil && o.Constraints != nil {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []ComplexConstraint and assigns it to the Constraints field.
func (o *SchemaType) SetConstraints(v []ComplexConstraint) {
	o.Constraints = &v
}

// GetVersionInfo returns the VersionInfo field value if set, zero value otherwise.
func (o *SchemaType) GetVersionInfo() string {
	if o == nil || o.VersionInfo == nil {
		var ret string
		return ret
	}
	return *o.VersionInfo
}

// GetVersionInfoOk returns a tuple with the VersionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetVersionInfoOk() (*string, bool) {
	if o == nil || o.VersionInfo == nil {
		return nil, false
	}
	return o.VersionInfo, true
}

// HasVersionInfo returns a boolean if a field has been set.
func (o *SchemaType) HasVersionInfo() bool {
	if o != nil && o.VersionInfo != nil {
		return true
	}

	return false
}

// SetVersionInfo gets a reference to the given string and assigns it to the VersionInfo field.
func (o *SchemaType) SetVersionInfo(v string) {
	o.VersionInfo = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SchemaType) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SchemaType) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SchemaType) SetVersion(v string) {
	o.Version = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *SchemaType) GetProperties() map[string]PropertyDefinition {
	if o == nil || o.Properties == nil {
		var ret map[string]PropertyDefinition
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetPropertiesOk() (*map[string]PropertyDefinition, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *SchemaType) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]PropertyDefinition and assigns it to the Properties field.
func (o *SchemaType) SetProperties(v map[string]PropertyDefinition) {
	o.Properties = &v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *SchemaType) GetDocumentation() string {
	if o == nil || o.Documentation == nil {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetDocumentationOk() (*string, bool) {
	if o == nil || o.Documentation == nil {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *SchemaType) HasDocumentation() bool {
	if o != nil && o.Documentation != nil {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *SchemaType) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SchemaType) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SchemaType) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SchemaType) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SchemaType) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SchemaType) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SchemaType) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o SchemaType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SummaryPattern != nil {
		toSerialize["summaryPattern"] = o.SummaryPattern
	}
	if o.Constraints != nil {
		toSerialize["constraints"] = o.Constraints
	}
	if o.VersionInfo != nil {
		toSerialize["versionInfo"] = o.VersionInfo
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Documentation != nil {
		toSerialize["documentation"] = o.Documentation
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaType struct {
	value *SchemaType
	isSet bool
}

func (v NullableSchemaType) Get() *SchemaType {
	return v.value
}

func (v *NullableSchemaType) Set(val *SchemaType) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaType) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaType(val *SchemaType) *NullableSchemaType {
	return &NullableSchemaType{value: val, isSet: true}
}

func (v NullableSchemaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


