# SlowPhysicalStorageDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**CustomThresholds** | Pointer to [**SlowPhysicalStorageThresholds**](SlowPhysicalStorageThresholds.md) |  | [optional] 

## Methods

### NewSlowPhysicalStorageDetectionConfig

`func NewSlowPhysicalStorageDetectionConfig(enabled bool, ) *SlowPhysicalStorageDetectionConfig`

NewSlowPhysicalStorageDetectionConfig instantiates a new SlowPhysicalStorageDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlowPhysicalStorageDetectionConfigWithDefaults

`func NewSlowPhysicalStorageDetectionConfigWithDefaults() *SlowPhysicalStorageDetectionConfig`

NewSlowPhysicalStorageDetectionConfigWithDefaults instantiates a new SlowPhysicalStorageDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SlowPhysicalStorageDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SlowPhysicalStorageDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SlowPhysicalStorageDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCustomThresholds

`func (o *SlowPhysicalStorageDetectionConfig) GetCustomThresholds() SlowPhysicalStorageThresholds`

GetCustomThresholds returns the CustomThresholds field if non-nil, zero value otherwise.

### GetCustomThresholdsOk

`func (o *SlowPhysicalStorageDetectionConfig) GetCustomThresholdsOk() (*SlowPhysicalStorageThresholds, bool)`

GetCustomThresholdsOk returns a tuple with the CustomThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomThresholds

`func (o *SlowPhysicalStorageDetectionConfig) SetCustomThresholds(v SlowPhysicalStorageThresholds)`

SetCustomThresholds sets CustomThresholds field to given value.

### HasCustomThresholds

`func (o *SlowPhysicalStorageDetectionConfig) HasCustomThresholds() bool`

HasCustomThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


