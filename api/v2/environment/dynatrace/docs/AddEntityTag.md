# AddEntityTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The value of the custom tag to be added to monitored entities. May be null | [optional] 
**Key** | **string** | The key of the custom tag to be added to monitored entities. | 

## Methods

### NewAddEntityTag

`func NewAddEntityTag(key string, ) *AddEntityTag`

NewAddEntityTag instantiates a new AddEntityTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddEntityTagWithDefaults

`func NewAddEntityTagWithDefaults() *AddEntityTag`

NewAddEntityTagWithDefaults instantiates a new AddEntityTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *AddEntityTag) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddEntityTag) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AddEntityTag) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AddEntityTag) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetKey

`func (o *AddEntityTag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AddEntityTag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AddEntityTag) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


