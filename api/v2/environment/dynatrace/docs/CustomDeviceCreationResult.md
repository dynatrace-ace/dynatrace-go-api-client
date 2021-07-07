# CustomDeviceCreationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** | The Dynatrace entity ID of the custom device. | [optional] 
**GroupId** | Pointer to **string** | The Dynatrace entity ID of the custom device group. | [optional] 

## Methods

### NewCustomDeviceCreationResult

`func NewCustomDeviceCreationResult() *CustomDeviceCreationResult`

NewCustomDeviceCreationResult instantiates a new CustomDeviceCreationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDeviceCreationResultWithDefaults

`func NewCustomDeviceCreationResultWithDefaults() *CustomDeviceCreationResult`

NewCustomDeviceCreationResultWithDefaults instantiates a new CustomDeviceCreationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *CustomDeviceCreationResult) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *CustomDeviceCreationResult) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *CustomDeviceCreationResult) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *CustomDeviceCreationResult) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetGroupId

`func (o *CustomDeviceCreationResult) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CustomDeviceCreationResult) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CustomDeviceCreationResult) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CustomDeviceCreationResult) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


