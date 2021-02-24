# ElbHighConnectionErrorsThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionErrorsPerMinute** | **int32** | Alert if number of backend connection errors is higher than *X* per minute in 3 out of 5 samples. | 

## Methods

### NewElbHighConnectionErrorsThresholds

`func NewElbHighConnectionErrorsThresholds(connectionErrorsPerMinute int32, ) *ElbHighConnectionErrorsThresholds`

NewElbHighConnectionErrorsThresholds instantiates a new ElbHighConnectionErrorsThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElbHighConnectionErrorsThresholdsWithDefaults

`func NewElbHighConnectionErrorsThresholdsWithDefaults() *ElbHighConnectionErrorsThresholds`

NewElbHighConnectionErrorsThresholdsWithDefaults instantiates a new ElbHighConnectionErrorsThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionErrorsPerMinute

`func (o *ElbHighConnectionErrorsThresholds) GetConnectionErrorsPerMinute() int32`

GetConnectionErrorsPerMinute returns the ConnectionErrorsPerMinute field if non-nil, zero value otherwise.

### GetConnectionErrorsPerMinuteOk

`func (o *ElbHighConnectionErrorsThresholds) GetConnectionErrorsPerMinuteOk() (*int32, bool)`

GetConnectionErrorsPerMinuteOk returns a tuple with the ConnectionErrorsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionErrorsPerMinute

`func (o *ElbHighConnectionErrorsThresholds) SetConnectionErrorsPerMinute(v int32)`

SetConnectionErrorsPerMinute sets ConnectionErrorsPerMinute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


