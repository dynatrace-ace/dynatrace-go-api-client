# CapturedMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**MethodReference**](MethodReference.md) |  | 
**Capture** | **string** | What to capture from the method. | 
**ArgumentIndex** | Pointer to **int32** | The index of the argument to capture. Set &#x60;0&#x60; to capture the return value, &#x60;1&#x60; or higher to capture a mehtod argument.    Required if the **capture** is set to &#x60;ARGUMENT&#x60;.   Not applicable in other cases. | [optional] 
**DeepObjectAccess** | Pointer to **string** | The getter chain to apply to the captured object. It is required in one of the following cases:   The **capture** is set to &#x60;THIS&#x60;.    The **capture** is set to &#x60;ARGUMENT&#x60;, and the argument is not a primitive, a primitive wrapper class, a string, or an array.    Not applicable in other cases. | [optional] 

## Methods

### NewCapturedMethod

`func NewCapturedMethod(method MethodReference, capture string, ) *CapturedMethod`

NewCapturedMethod instantiates a new CapturedMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapturedMethodWithDefaults

`func NewCapturedMethodWithDefaults() *CapturedMethod`

NewCapturedMethodWithDefaults instantiates a new CapturedMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *CapturedMethod) GetMethod() MethodReference`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CapturedMethod) GetMethodOk() (*MethodReference, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CapturedMethod) SetMethod(v MethodReference)`

SetMethod sets Method field to given value.


### GetCapture

`func (o *CapturedMethod) GetCapture() string`

GetCapture returns the Capture field if non-nil, zero value otherwise.

### GetCaptureOk

`func (o *CapturedMethod) GetCaptureOk() (*string, bool)`

GetCaptureOk returns a tuple with the Capture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapture

`func (o *CapturedMethod) SetCapture(v string)`

SetCapture sets Capture field to given value.


### GetArgumentIndex

`func (o *CapturedMethod) GetArgumentIndex() int32`

GetArgumentIndex returns the ArgumentIndex field if non-nil, zero value otherwise.

### GetArgumentIndexOk

`func (o *CapturedMethod) GetArgumentIndexOk() (*int32, bool)`

GetArgumentIndexOk returns a tuple with the ArgumentIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgumentIndex

`func (o *CapturedMethod) SetArgumentIndex(v int32)`

SetArgumentIndex sets ArgumentIndex field to given value.

### HasArgumentIndex

`func (o *CapturedMethod) HasArgumentIndex() bool`

HasArgumentIndex returns a boolean if a field has been set.

### GetDeepObjectAccess

`func (o *CapturedMethod) GetDeepObjectAccess() string`

GetDeepObjectAccess returns the DeepObjectAccess field if non-nil, zero value otherwise.

### GetDeepObjectAccessOk

`func (o *CapturedMethod) GetDeepObjectAccessOk() (*string, bool)`

GetDeepObjectAccessOk returns a tuple with the DeepObjectAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepObjectAccess

`func (o *CapturedMethod) SetDeepObjectAccess(v string)`

SetDeepObjectAccess sets DeepObjectAccess field to given value.

### HasDeepObjectAccess

`func (o *CapturedMethod) HasDeepObjectAccess() bool`

HasDeepObjectAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


