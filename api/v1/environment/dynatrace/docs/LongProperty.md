# LongProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The custom key of the property. | [optional] 
**Value** | Pointer to **int64** | The Long value of the property. | [optional] 

## Methods

### NewLongProperty

`func NewLongProperty() *LongProperty`

NewLongProperty instantiates a new LongProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLongPropertyWithDefaults

`func NewLongPropertyWithDefaults() *LongProperty`

NewLongPropertyWithDefaults instantiates a new LongProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *LongProperty) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LongProperty) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LongProperty) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *LongProperty) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *LongProperty) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LongProperty) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LongProperty) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *LongProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


