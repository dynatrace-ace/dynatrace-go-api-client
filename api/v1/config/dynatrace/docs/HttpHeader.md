# HttpHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the HTTP header. | 
**Value** | Pointer to **string** | The value of the HTTP header. May contain an empty value.    Required when creating a new notification.   For the **Authorization** header, GET requests return the &#x60;null&#x60; value.   If you want update a notification configuration with an **Authorization** header which you want to remain intact, set the **Authorization** header with the &#x60;null&#x60; value. | [optional] 

## Methods

### NewHttpHeader

`func NewHttpHeader(name string, ) *HttpHeader`

NewHttpHeader instantiates a new HttpHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpHeaderWithDefaults

`func NewHttpHeaderWithDefaults() *HttpHeader`

NewHttpHeaderWithDefaults instantiates a new HttpHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HttpHeader) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HttpHeader) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HttpHeader) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *HttpHeader) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HttpHeader) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HttpHeader) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *HttpHeader) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


