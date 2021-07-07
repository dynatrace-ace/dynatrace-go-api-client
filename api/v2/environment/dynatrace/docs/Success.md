# Success

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Detailed message | [optional] 
**Code** | Pointer to **int32** | The HTTP status code | [optional] 

## Methods

### NewSuccess

`func NewSuccess() *Success`

NewSuccess instantiates a new Success object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccessWithDefaults

`func NewSuccessWithDefaults() *Success`

NewSuccessWithDefaults instantiates a new Success object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Success) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Success) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Success) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Success) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *Success) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Success) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Success) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *Success) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


