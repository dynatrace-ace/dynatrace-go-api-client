# SettingsObjectCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** | The version of the schema on which the object is based. | [optional] 
**InsertAfter** | Pointer to **string** | The position of the new object. The new object will be added after the specified one.   If &#x60;null&#x60;, the new object will be placed in the last position.   If set to empty string, the new object will be placed in the first position.   Only applicable for objects based on schemas with ordered objects (schema&#39;s **ordered** parameter is set to &#x60;true&#x60;). | [optional] 
**Scope** | **string** | The scope that the object targets. | 
**SchemaId** | **string** | The schema on which the object is based. | 
**Value** | **map[string]interface{}** | The value of the setting.    It defines the actual values of settings&#39; parameters.   The actual content depends on the object&#39;s schema. | 

## Methods

### NewSettingsObjectCreate

`func NewSettingsObjectCreate(scope string, schemaId string, value map[string]interface{}, ) *SettingsObjectCreate`

NewSettingsObjectCreate instantiates a new SettingsObjectCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsObjectCreateWithDefaults

`func NewSettingsObjectCreateWithDefaults() *SettingsObjectCreate`

NewSettingsObjectCreateWithDefaults instantiates a new SettingsObjectCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *SettingsObjectCreate) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *SettingsObjectCreate) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *SettingsObjectCreate) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *SettingsObjectCreate) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetInsertAfter

`func (o *SettingsObjectCreate) GetInsertAfter() string`

GetInsertAfter returns the InsertAfter field if non-nil, zero value otherwise.

### GetInsertAfterOk

`func (o *SettingsObjectCreate) GetInsertAfterOk() (*string, bool)`

GetInsertAfterOk returns a tuple with the InsertAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertAfter

`func (o *SettingsObjectCreate) SetInsertAfter(v string)`

SetInsertAfter sets InsertAfter field to given value.

### HasInsertAfter

`func (o *SettingsObjectCreate) HasInsertAfter() bool`

HasInsertAfter returns a boolean if a field has been set.

### GetScope

`func (o *SettingsObjectCreate) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SettingsObjectCreate) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SettingsObjectCreate) SetScope(v string)`

SetScope sets Scope field to given value.


### GetSchemaId

`func (o *SettingsObjectCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *SettingsObjectCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *SettingsObjectCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.


### GetValue

`func (o *SettingsObjectCreate) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SettingsObjectCreate) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SettingsObjectCreate) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


