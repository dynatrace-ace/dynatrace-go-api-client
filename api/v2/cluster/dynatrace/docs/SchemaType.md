# SchemaType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SummaryPattern** | Pointer to **string** | The pattern for the summary (for example, \&quot;Alert after *X* minutes.\&quot;) of the configuration in the UI. | [optional] 
**Constraints** | Pointer to [**[]ComplexConstraint**](ComplexConstraint.md) | A list of constraints limiting the values to be accepted. | [optional] 
**VersionInfo** | Pointer to **string** | A short description of the version. | [optional] 
**Version** | Pointer to **string** | The version of the type. | [optional] 
**Properties** | Pointer to [**map[string]PropertyDefinition**](PropertyDefinition.md) | Definition of properties that can be persisted. | [optional] 
**Documentation** | Pointer to **string** | An extended description and/or links to documentation. | [optional] 
**Description** | Pointer to **string** | A short description of the property. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the property. | [optional] 

## Methods

### NewSchemaType

`func NewSchemaType() *SchemaType`

NewSchemaType instantiates a new SchemaType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaTypeWithDefaults

`func NewSchemaTypeWithDefaults() *SchemaType`

NewSchemaTypeWithDefaults instantiates a new SchemaType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummaryPattern

`func (o *SchemaType) GetSummaryPattern() string`

GetSummaryPattern returns the SummaryPattern field if non-nil, zero value otherwise.

### GetSummaryPatternOk

`func (o *SchemaType) GetSummaryPatternOk() (*string, bool)`

GetSummaryPatternOk returns a tuple with the SummaryPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryPattern

`func (o *SchemaType) SetSummaryPattern(v string)`

SetSummaryPattern sets SummaryPattern field to given value.

### HasSummaryPattern

`func (o *SchemaType) HasSummaryPattern() bool`

HasSummaryPattern returns a boolean if a field has been set.

### GetConstraints

`func (o *SchemaType) GetConstraints() []ComplexConstraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *SchemaType) GetConstraintsOk() (*[]ComplexConstraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *SchemaType) SetConstraints(v []ComplexConstraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *SchemaType) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetVersionInfo

`func (o *SchemaType) GetVersionInfo() string`

GetVersionInfo returns the VersionInfo field if non-nil, zero value otherwise.

### GetVersionInfoOk

`func (o *SchemaType) GetVersionInfoOk() (*string, bool)`

GetVersionInfoOk returns a tuple with the VersionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionInfo

`func (o *SchemaType) SetVersionInfo(v string)`

SetVersionInfo sets VersionInfo field to given value.

### HasVersionInfo

`func (o *SchemaType) HasVersionInfo() bool`

HasVersionInfo returns a boolean if a field has been set.

### GetVersion

`func (o *SchemaType) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SchemaType) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SchemaType) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SchemaType) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetProperties

`func (o *SchemaType) GetProperties() map[string]PropertyDefinition`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SchemaType) GetPropertiesOk() (*map[string]PropertyDefinition, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SchemaType) SetProperties(v map[string]PropertyDefinition)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SchemaType) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetDocumentation

`func (o *SchemaType) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *SchemaType) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *SchemaType) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *SchemaType) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetDescription

`func (o *SchemaType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemaType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *SchemaType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SchemaType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SchemaType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SchemaType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


