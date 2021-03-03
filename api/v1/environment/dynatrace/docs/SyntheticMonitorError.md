# SyntheticMonitorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | The error code. | 
**Message** | **string** | The error message. | 

## Methods

### NewSyntheticMonitorError

`func NewSyntheticMonitorError(code int32, message string, ) *SyntheticMonitorError`

NewSyntheticMonitorError instantiates a new SyntheticMonitorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticMonitorErrorWithDefaults

`func NewSyntheticMonitorErrorWithDefaults() *SyntheticMonitorError`

NewSyntheticMonitorErrorWithDefaults instantiates a new SyntheticMonitorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SyntheticMonitorError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SyntheticMonitorError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SyntheticMonitorError) SetCode(v int32)`

SetCode sets Code field to given value.


### GetMessage

`func (o *SyntheticMonitorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SyntheticMonitorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SyntheticMonitorError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


