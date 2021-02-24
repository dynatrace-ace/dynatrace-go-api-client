# RdsLowStorageThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeStoragePercentage** | **int32** | Alert if free storage space divided by allocated storage is lower than *X*% in 3 out of 5 samples. | 

## Methods

### NewRdsLowStorageThresholds

`func NewRdsLowStorageThresholds(freeStoragePercentage int32, ) *RdsLowStorageThresholds`

NewRdsLowStorageThresholds instantiates a new RdsLowStorageThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdsLowStorageThresholdsWithDefaults

`func NewRdsLowStorageThresholdsWithDefaults() *RdsLowStorageThresholds`

NewRdsLowStorageThresholdsWithDefaults instantiates a new RdsLowStorageThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeStoragePercentage

`func (o *RdsLowStorageThresholds) GetFreeStoragePercentage() int32`

GetFreeStoragePercentage returns the FreeStoragePercentage field if non-nil, zero value otherwise.

### GetFreeStoragePercentageOk

`func (o *RdsLowStorageThresholds) GetFreeStoragePercentageOk() (*int32, bool)`

GetFreeStoragePercentageOk returns a tuple with the FreeStoragePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeStoragePercentage

`func (o *RdsLowStorageThresholds) SetFreeStoragePercentage(v int32)`

SetFreeStoragePercentage sets FreeStoragePercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


