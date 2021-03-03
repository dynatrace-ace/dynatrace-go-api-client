# DoubleProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The custom key of the property. | [optional] 
**Value** | Pointer to **float64** | The floating-point numeric value of the property. | [optional] 

## Methods

### NewDoubleProperty

`func NewDoubleProperty() *DoubleProperty`

NewDoubleProperty instantiates a new DoubleProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDoublePropertyWithDefaults

`func NewDoublePropertyWithDefaults() *DoubleProperty`

NewDoublePropertyWithDefaults instantiates a new DoubleProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DoubleProperty) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DoubleProperty) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DoubleProperty) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DoubleProperty) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *DoubleProperty) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DoubleProperty) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DoubleProperty) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *DoubleProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


