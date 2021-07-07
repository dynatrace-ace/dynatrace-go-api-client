# FromPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromTypes** | Pointer to **[]string** | A list of monitored entity types that can occupy the FROM position. | [optional] 
**Id** | Pointer to **string** | The ID of the relationship. | [optional] 

## Methods

### NewFromPosition

`func NewFromPosition() *FromPosition`

NewFromPosition instantiates a new FromPosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFromPositionWithDefaults

`func NewFromPositionWithDefaults() *FromPosition`

NewFromPositionWithDefaults instantiates a new FromPosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromTypes

`func (o *FromPosition) GetFromTypes() []string`

GetFromTypes returns the FromTypes field if non-nil, zero value otherwise.

### GetFromTypesOk

`func (o *FromPosition) GetFromTypesOk() (*[]string, bool)`

GetFromTypesOk returns a tuple with the FromTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTypes

`func (o *FromPosition) SetFromTypes(v []string)`

SetFromTypes sets FromTypes field to given value.

### HasFromTypes

`func (o *FromPosition) HasFromTypes() bool`

HasFromTypes returns a boolean if a field has been set.

### GetId

`func (o *FromPosition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FromPosition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FromPosition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FromPosition) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


