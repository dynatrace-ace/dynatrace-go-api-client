# RdsLowStorageDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**CustomThresholds** | Pointer to [**RdsLowStorageThresholds**](RdsLowStorageThresholds.md) |  | [optional] 

## Methods

### NewRdsLowStorageDetectionConfig

`func NewRdsLowStorageDetectionConfig(enabled bool, ) *RdsLowStorageDetectionConfig`

NewRdsLowStorageDetectionConfig instantiates a new RdsLowStorageDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdsLowStorageDetectionConfigWithDefaults

`func NewRdsLowStorageDetectionConfigWithDefaults() *RdsLowStorageDetectionConfig`

NewRdsLowStorageDetectionConfigWithDefaults instantiates a new RdsLowStorageDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RdsLowStorageDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RdsLowStorageDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RdsLowStorageDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCustomThresholds

`func (o *RdsLowStorageDetectionConfig) GetCustomThresholds() RdsLowStorageThresholds`

GetCustomThresholds returns the CustomThresholds field if non-nil, zero value otherwise.

### GetCustomThresholdsOk

`func (o *RdsLowStorageDetectionConfig) GetCustomThresholdsOk() (*RdsLowStorageThresholds, bool)`

GetCustomThresholdsOk returns a tuple with the CustomThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomThresholds

`func (o *RdsLowStorageDetectionConfig) SetCustomThresholds(v RdsLowStorageThresholds)`

SetCustomThresholds sets CustomThresholds field to given value.

### HasCustomThresholds

`func (o *RdsLowStorageDetectionConfig) HasCustomThresholds() bool`

HasCustomThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


