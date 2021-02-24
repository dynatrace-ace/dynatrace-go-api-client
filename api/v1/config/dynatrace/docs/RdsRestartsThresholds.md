# RdsRestartsThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestartsPerMinute** | **int32** | Alert if number of restarts is *X* per minute or higher in 3 out of 20 samples. | 

## Methods

### NewRdsRestartsThresholds

`func NewRdsRestartsThresholds(restartsPerMinute int32, ) *RdsRestartsThresholds`

NewRdsRestartsThresholds instantiates a new RdsRestartsThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdsRestartsThresholdsWithDefaults

`func NewRdsRestartsThresholdsWithDefaults() *RdsRestartsThresholds`

NewRdsRestartsThresholdsWithDefaults instantiates a new RdsRestartsThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestartsPerMinute

`func (o *RdsRestartsThresholds) GetRestartsPerMinute() int32`

GetRestartsPerMinute returns the RestartsPerMinute field if non-nil, zero value otherwise.

### GetRestartsPerMinuteOk

`func (o *RdsRestartsThresholds) GetRestartsPerMinuteOk() (*int32, bool)`

GetRestartsPerMinuteOk returns a tuple with the RestartsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartsPerMinute

`func (o *RdsRestartsThresholds) SetRestartsPerMinute(v int32)`

SetRestartsPerMinute sets RestartsPerMinute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


