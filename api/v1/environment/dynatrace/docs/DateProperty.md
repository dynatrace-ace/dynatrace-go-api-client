# DateProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The custom key of the property. | [optional] 
**Value** | Pointer to **time.Time** | The date value of the property. | [optional] 

## Methods

### NewDateProperty

`func NewDateProperty() *DateProperty`

NewDateProperty instantiates a new DateProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatePropertyWithDefaults

`func NewDatePropertyWithDefaults() *DateProperty`

NewDatePropertyWithDefaults instantiates a new DateProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DateProperty) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DateProperty) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DateProperty) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DateProperty) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *DateProperty) GetValue() time.Time`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DateProperty) GetValueOk() (*time.Time, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DateProperty) SetValue(v time.Time)`

SetValue sets Value field to given value.

### HasValue

`func (o *DateProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


