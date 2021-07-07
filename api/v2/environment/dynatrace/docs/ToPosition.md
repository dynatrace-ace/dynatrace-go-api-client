# ToPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToTypes** | Pointer to **[]string** | A list of monitored entity types that can occupy the TO position. | [optional] 
**Id** | Pointer to **string** | The ID of the relationship. | [optional] 

## Methods

### NewToPosition

`func NewToPosition() *ToPosition`

NewToPosition instantiates a new ToPosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToPositionWithDefaults

`func NewToPositionWithDefaults() *ToPosition`

NewToPositionWithDefaults instantiates a new ToPosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToTypes

`func (o *ToPosition) GetToTypes() []string`

GetToTypes returns the ToTypes field if non-nil, zero value otherwise.

### GetToTypesOk

`func (o *ToPosition) GetToTypesOk() (*[]string, bool)`

GetToTypesOk returns a tuple with the ToTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTypes

`func (o *ToPosition) SetToTypes(v []string)`

SetToTypes sets ToTypes field to given value.

### HasToTypes

`func (o *ToPosition) HasToTypes() bool`

HasToTypes returns a boolean if a field has been set.

### GetId

`func (o *ToPosition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToPosition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToPosition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ToPosition) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


