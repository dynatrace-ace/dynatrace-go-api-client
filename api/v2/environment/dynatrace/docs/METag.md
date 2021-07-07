# METag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StringRepresentation** | Pointer to **string** | The string representation of the tag. | [optional] 
**Value** | Pointer to **string** | The value of the tag. | [optional] 
**Key** | Pointer to **string** | The key of the tag. | [optional] 
**Context** | Pointer to **string** | The origin of the tag, such as AWS or Cloud Foundry.    Custom tags use the &#x60;CONTEXTLESS&#x60; value. | [optional] 

## Methods

### NewMETag

`func NewMETag() *METag`

NewMETag instantiates a new METag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMETagWithDefaults

`func NewMETagWithDefaults() *METag`

NewMETagWithDefaults instantiates a new METag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStringRepresentation

`func (o *METag) GetStringRepresentation() string`

GetStringRepresentation returns the StringRepresentation field if non-nil, zero value otherwise.

### GetStringRepresentationOk

`func (o *METag) GetStringRepresentationOk() (*string, bool)`

GetStringRepresentationOk returns a tuple with the StringRepresentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringRepresentation

`func (o *METag) SetStringRepresentation(v string)`

SetStringRepresentation sets StringRepresentation field to given value.

### HasStringRepresentation

`func (o *METag) HasStringRepresentation() bool`

HasStringRepresentation returns a boolean if a field has been set.

### GetValue

`func (o *METag) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *METag) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *METag) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *METag) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetKey

`func (o *METag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *METag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *METag) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *METag) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetContext

`func (o *METag) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *METag) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *METag) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *METag) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


