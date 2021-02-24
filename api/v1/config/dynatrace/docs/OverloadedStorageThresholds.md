# OverloadedStorageThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandAbortsNumber** | **int32** | Alert if number of command aborts is higher than *X* in 3 out of 5 samples. | 

## Methods

### NewOverloadedStorageThresholds

`func NewOverloadedStorageThresholds(commandAbortsNumber int32, ) *OverloadedStorageThresholds`

NewOverloadedStorageThresholds instantiates a new OverloadedStorageThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverloadedStorageThresholdsWithDefaults

`func NewOverloadedStorageThresholdsWithDefaults() *OverloadedStorageThresholds`

NewOverloadedStorageThresholdsWithDefaults instantiates a new OverloadedStorageThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandAbortsNumber

`func (o *OverloadedStorageThresholds) GetCommandAbortsNumber() int32`

GetCommandAbortsNumber returns the CommandAbortsNumber field if non-nil, zero value otherwise.

### GetCommandAbortsNumberOk

`func (o *OverloadedStorageThresholds) GetCommandAbortsNumberOk() (*int32, bool)`

GetCommandAbortsNumberOk returns a tuple with the CommandAbortsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandAbortsNumber

`func (o *OverloadedStorageThresholds) SetCommandAbortsNumber(v int32)`

SetCommandAbortsNumber sets CommandAbortsNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


