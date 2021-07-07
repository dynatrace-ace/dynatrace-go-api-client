# SchemaDefinitionRestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dynatrace** | Pointer to **string** | The version of the data format. | [optional] 
**SchemaId** | Pointer to **string** | The ID of the schema. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the schema. | [optional] 
**Description** | Pointer to **string** | A short description of the schema. | [optional] 
**Documentation** | Pointer to **string** | An extended description of the schema and/or links to documentation. | [optional] 
**Version** | Pointer to **string** | The version of the schema. | [optional] 
**MultiObject** | Pointer to **bool** | Multiple (&#x60;true&#x60;) objects per scope are permitted or a single (&#x60;false&#x60;) object per scope is permitted. | [optional] 
**Ordered** | Pointer to **bool** | If &#x60;true&#x60; the order of objects has semantic significance.   Only applicable when **multiObject** is set to &#x60;true&#x60;. | [optional] 
**MaxObjects** | Pointer to **int32** | The maximum amount of objects per scope.   Only applicable when **multiObject** is set to &#x60;true&#x60;. | [optional] 
**AllowedScopes** | Pointer to **[]string** | A list of scopes where the schema can be used. | [optional] 
**Enums** | Pointer to [**map[string]EnumType**](EnumType.md) | A list of definitions of enum properties. | [optional] 
**Types** | Pointer to [**map[string]SchemaType**](SchemaType.md) | A list of definitions of types.    A type is a complex property that contains its own set of subproperties. | [optional] 
**Properties** | Pointer to [**map[string]PropertyDefinition**](PropertyDefinition.md) | A list of schema&#39;s properties. | [optional] 
**Constraints** | Pointer to [**[]ComplexConstraint**](ComplexConstraint.md) | A list of constrains limiting the values to be accepted by the schema. | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata of the setting. | [optional] 

## Methods

### NewSchemaDefinitionRestDto

`func NewSchemaDefinitionRestDto() *SchemaDefinitionRestDto`

NewSchemaDefinitionRestDto instantiates a new SchemaDefinitionRestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaDefinitionRestDtoWithDefaults

`func NewSchemaDefinitionRestDtoWithDefaults() *SchemaDefinitionRestDto`

NewSchemaDefinitionRestDtoWithDefaults instantiates a new SchemaDefinitionRestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDynatrace

`func (o *SchemaDefinitionRestDto) GetDynatrace() string`

GetDynatrace returns the Dynatrace field if non-nil, zero value otherwise.

### GetDynatraceOk

`func (o *SchemaDefinitionRestDto) GetDynatraceOk() (*string, bool)`

GetDynatraceOk returns a tuple with the Dynatrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynatrace

`func (o *SchemaDefinitionRestDto) SetDynatrace(v string)`

SetDynatrace sets Dynatrace field to given value.

### HasDynatrace

`func (o *SchemaDefinitionRestDto) HasDynatrace() bool`

HasDynatrace returns a boolean if a field has been set.

### GetSchemaId

`func (o *SchemaDefinitionRestDto) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *SchemaDefinitionRestDto) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *SchemaDefinitionRestDto) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *SchemaDefinitionRestDto) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetDisplayName

`func (o *SchemaDefinitionRestDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SchemaDefinitionRestDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SchemaDefinitionRestDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SchemaDefinitionRestDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *SchemaDefinitionRestDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaDefinitionRestDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaDefinitionRestDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemaDefinitionRestDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentation

`func (o *SchemaDefinitionRestDto) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *SchemaDefinitionRestDto) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *SchemaDefinitionRestDto) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *SchemaDefinitionRestDto) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetVersion

`func (o *SchemaDefinitionRestDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SchemaDefinitionRestDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SchemaDefinitionRestDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SchemaDefinitionRestDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMultiObject

`func (o *SchemaDefinitionRestDto) GetMultiObject() bool`

GetMultiObject returns the MultiObject field if non-nil, zero value otherwise.

### GetMultiObjectOk

`func (o *SchemaDefinitionRestDto) GetMultiObjectOk() (*bool, bool)`

GetMultiObjectOk returns a tuple with the MultiObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiObject

`func (o *SchemaDefinitionRestDto) SetMultiObject(v bool)`

SetMultiObject sets MultiObject field to given value.

### HasMultiObject

`func (o *SchemaDefinitionRestDto) HasMultiObject() bool`

HasMultiObject returns a boolean if a field has been set.

### GetOrdered

`func (o *SchemaDefinitionRestDto) GetOrdered() bool`

GetOrdered returns the Ordered field if non-nil, zero value otherwise.

### GetOrderedOk

`func (o *SchemaDefinitionRestDto) GetOrderedOk() (*bool, bool)`

GetOrderedOk returns a tuple with the Ordered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdered

`func (o *SchemaDefinitionRestDto) SetOrdered(v bool)`

SetOrdered sets Ordered field to given value.

### HasOrdered

`func (o *SchemaDefinitionRestDto) HasOrdered() bool`

HasOrdered returns a boolean if a field has been set.

### GetMaxObjects

`func (o *SchemaDefinitionRestDto) GetMaxObjects() int32`

GetMaxObjects returns the MaxObjects field if non-nil, zero value otherwise.

### GetMaxObjectsOk

`func (o *SchemaDefinitionRestDto) GetMaxObjectsOk() (*int32, bool)`

GetMaxObjectsOk returns a tuple with the MaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxObjects

`func (o *SchemaDefinitionRestDto) SetMaxObjects(v int32)`

SetMaxObjects sets MaxObjects field to given value.

### HasMaxObjects

`func (o *SchemaDefinitionRestDto) HasMaxObjects() bool`

HasMaxObjects returns a boolean if a field has been set.

### GetAllowedScopes

`func (o *SchemaDefinitionRestDto) GetAllowedScopes() []string`

GetAllowedScopes returns the AllowedScopes field if non-nil, zero value otherwise.

### GetAllowedScopesOk

`func (o *SchemaDefinitionRestDto) GetAllowedScopesOk() (*[]string, bool)`

GetAllowedScopesOk returns a tuple with the AllowedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedScopes

`func (o *SchemaDefinitionRestDto) SetAllowedScopes(v []string)`

SetAllowedScopes sets AllowedScopes field to given value.

### HasAllowedScopes

`func (o *SchemaDefinitionRestDto) HasAllowedScopes() bool`

HasAllowedScopes returns a boolean if a field has been set.

### GetEnums

`func (o *SchemaDefinitionRestDto) GetEnums() map[string]EnumType`

GetEnums returns the Enums field if non-nil, zero value otherwise.

### GetEnumsOk

`func (o *SchemaDefinitionRestDto) GetEnumsOk() (*map[string]EnumType, bool)`

GetEnumsOk returns a tuple with the Enums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnums

`func (o *SchemaDefinitionRestDto) SetEnums(v map[string]EnumType)`

SetEnums sets Enums field to given value.

### HasEnums

`func (o *SchemaDefinitionRestDto) HasEnums() bool`

HasEnums returns a boolean if a field has been set.

### GetTypes

`func (o *SchemaDefinitionRestDto) GetTypes() map[string]SchemaType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SchemaDefinitionRestDto) GetTypesOk() (*map[string]SchemaType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SchemaDefinitionRestDto) SetTypes(v map[string]SchemaType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *SchemaDefinitionRestDto) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetProperties

`func (o *SchemaDefinitionRestDto) GetProperties() map[string]PropertyDefinition`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SchemaDefinitionRestDto) GetPropertiesOk() (*map[string]PropertyDefinition, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SchemaDefinitionRestDto) SetProperties(v map[string]PropertyDefinition)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SchemaDefinitionRestDto) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetConstraints

`func (o *SchemaDefinitionRestDto) GetConstraints() []ComplexConstraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *SchemaDefinitionRestDto) GetConstraintsOk() (*[]ComplexConstraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *SchemaDefinitionRestDto) SetConstraints(v []ComplexConstraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *SchemaDefinitionRestDto) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetMetadata

`func (o *SchemaDefinitionRestDto) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SchemaDefinitionRestDto) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SchemaDefinitionRestDto) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SchemaDefinitionRestDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


