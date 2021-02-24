# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int32** |  | [optional] 
**ConstraintViolations** | Pointer to [**[]ConstraintViolation**](ConstraintViolation.md) |  | [optional] 

## Methods

### NewError

`func NewError() *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Error) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *Error) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Error) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Error) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *Error) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetConstraintViolations

`func (o *Error) GetConstraintViolations() []ConstraintViolation`

GetConstraintViolations returns the ConstraintViolations field if non-nil, zero value otherwise.

### GetConstraintViolationsOk

`func (o *Error) GetConstraintViolationsOk() (*[]ConstraintViolation, bool)`

GetConstraintViolationsOk returns a tuple with the ConstraintViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintViolations

`func (o *Error) SetConstraintViolations(v []ConstraintViolation)`

SetConstraintViolations sets ConstraintViolations field to given value.

### HasConstraintViolations

`func (o *Error) HasConstraintViolations() bool`

HasConstraintViolations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


