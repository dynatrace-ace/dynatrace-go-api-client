# OutOfMemoryThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutOfMemoryExceptionsNumber** | **int32** | Alert if the number of Java out of memory exceptions is *X* per minute or higher. | 

## Methods

### NewOutOfMemoryThresholds

`func NewOutOfMemoryThresholds(outOfMemoryExceptionsNumber int32, ) *OutOfMemoryThresholds`

NewOutOfMemoryThresholds instantiates a new OutOfMemoryThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutOfMemoryThresholdsWithDefaults

`func NewOutOfMemoryThresholdsWithDefaults() *OutOfMemoryThresholds`

NewOutOfMemoryThresholdsWithDefaults instantiates a new OutOfMemoryThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutOfMemoryExceptionsNumber

`func (o *OutOfMemoryThresholds) GetOutOfMemoryExceptionsNumber() int32`

GetOutOfMemoryExceptionsNumber returns the OutOfMemoryExceptionsNumber field if non-nil, zero value otherwise.

### GetOutOfMemoryExceptionsNumberOk

`func (o *OutOfMemoryThresholds) GetOutOfMemoryExceptionsNumberOk() (*int32, bool)`

GetOutOfMemoryExceptionsNumberOk returns a tuple with the OutOfMemoryExceptionsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfMemoryExceptionsNumber

`func (o *OutOfMemoryThresholds) SetOutOfMemoryExceptionsNumber(v int32)`

SetOutOfMemoryExceptionsNumber sets OutOfMemoryExceptionsNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


