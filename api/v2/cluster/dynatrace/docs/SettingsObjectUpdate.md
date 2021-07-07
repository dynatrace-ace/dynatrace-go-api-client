# SettingsObjectUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateToken** | Pointer to **string** | The update token of the object. You can use it to detect simultaneous modifications by different users.   It is generated upon retrieval (GET requests). If set on update (PUT request) or deletion, the update/deletion will be allowed only if there wasn&#39;t any change between the retrieval and the update.   If omitted on update/deletion, the operation overrides the current value or deletes it without any checks. | [optional] 
**SchemaVersion** | Pointer to **string** | The version of the schema on which the object is based. | [optional] 
**Value** | **map[string]interface{}** | The value of the setting.    It defines the actual values of settings&#39; parameters.   The actual content depends on the object&#39;s schema. | 

## Methods

### NewSettingsObjectUpdate

`func NewSettingsObjectUpdate(value map[string]interface{}, ) *SettingsObjectUpdate`

NewSettingsObjectUpdate instantiates a new SettingsObjectUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsObjectUpdateWithDefaults

`func NewSettingsObjectUpdateWithDefaults() *SettingsObjectUpdate`

NewSettingsObjectUpdateWithDefaults instantiates a new SettingsObjectUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateToken

`func (o *SettingsObjectUpdate) GetUpdateToken() string`

GetUpdateToken returns the UpdateToken field if non-nil, zero value otherwise.

### GetUpdateTokenOk

`func (o *SettingsObjectUpdate) GetUpdateTokenOk() (*string, bool)`

GetUpdateTokenOk returns a tuple with the UpdateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateToken

`func (o *SettingsObjectUpdate) SetUpdateToken(v string)`

SetUpdateToken sets UpdateToken field to given value.

### HasUpdateToken

`func (o *SettingsObjectUpdate) HasUpdateToken() bool`

HasUpdateToken returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *SettingsObjectUpdate) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *SettingsObjectUpdate) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *SettingsObjectUpdate) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *SettingsObjectUpdate) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetValue

`func (o *SettingsObjectUpdate) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SettingsObjectUpdate) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SettingsObjectUpdate) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


