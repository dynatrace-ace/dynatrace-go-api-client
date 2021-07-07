# SettingsObjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | Pointer to **string** | The ID of the settings object. | [optional] 
**Code** | Pointer to **int32** | The HTTP status code for the object. | [optional] 

## Methods

### NewSettingsObjectResponse

`func NewSettingsObjectResponse() *SettingsObjectResponse`

NewSettingsObjectResponse instantiates a new SettingsObjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsObjectResponseWithDefaults

`func NewSettingsObjectResponseWithDefaults() *SettingsObjectResponse`

NewSettingsObjectResponseWithDefaults instantiates a new SettingsObjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *SettingsObjectResponse) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *SettingsObjectResponse) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *SettingsObjectResponse) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *SettingsObjectResponse) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetCode

`func (o *SettingsObjectResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SettingsObjectResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SettingsObjectResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *SettingsObjectResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


