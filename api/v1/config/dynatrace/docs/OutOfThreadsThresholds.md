# OutOfThreadsThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutOfThreadsExceptionsNumber** | **int32** | Alert if the number of Java out of threads exceptions is *X* per minute or higher. | 

## Methods

### NewOutOfThreadsThresholds

`func NewOutOfThreadsThresholds(outOfThreadsExceptionsNumber int32, ) *OutOfThreadsThresholds`

NewOutOfThreadsThresholds instantiates a new OutOfThreadsThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutOfThreadsThresholdsWithDefaults

`func NewOutOfThreadsThresholdsWithDefaults() *OutOfThreadsThresholds`

NewOutOfThreadsThresholdsWithDefaults instantiates a new OutOfThreadsThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutOfThreadsExceptionsNumber

`func (o *OutOfThreadsThresholds) GetOutOfThreadsExceptionsNumber() int32`

GetOutOfThreadsExceptionsNumber returns the OutOfThreadsExceptionsNumber field if non-nil, zero value otherwise.

### GetOutOfThreadsExceptionsNumberOk

`func (o *OutOfThreadsThresholds) GetOutOfThreadsExceptionsNumberOk() (*int32, bool)`

GetOutOfThreadsExceptionsNumberOk returns a tuple with the OutOfThreadsExceptionsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfThreadsExceptionsNumber

`func (o *OutOfThreadsThresholds) SetOutOfThreadsExceptionsNumber(v int32)`

SetOutOfThreadsExceptionsNumber sets OutOfThreadsExceptionsNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


