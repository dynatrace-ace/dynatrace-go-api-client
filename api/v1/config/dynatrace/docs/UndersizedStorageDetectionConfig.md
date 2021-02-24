# UndersizedStorageDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**CustomThresholds** | Pointer to [**UndersizedStorageThresholds**](UndersizedStorageThresholds.md) |  | [optional] 

## Methods

### NewUndersizedStorageDetectionConfig

`func NewUndersizedStorageDetectionConfig(enabled bool, ) *UndersizedStorageDetectionConfig`

NewUndersizedStorageDetectionConfig instantiates a new UndersizedStorageDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUndersizedStorageDetectionConfigWithDefaults

`func NewUndersizedStorageDetectionConfigWithDefaults() *UndersizedStorageDetectionConfig`

NewUndersizedStorageDetectionConfigWithDefaults instantiates a new UndersizedStorageDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UndersizedStorageDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UndersizedStorageDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UndersizedStorageDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCustomThresholds

`func (o *UndersizedStorageDetectionConfig) GetCustomThresholds() UndersizedStorageThresholds`

GetCustomThresholds returns the CustomThresholds field if non-nil, zero value otherwise.

### GetCustomThresholdsOk

`func (o *UndersizedStorageDetectionConfig) GetCustomThresholdsOk() (*UndersizedStorageThresholds, bool)`

GetCustomThresholdsOk returns a tuple with the CustomThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomThresholds

`func (o *UndersizedStorageDetectionConfig) SetCustomThresholds(v UndersizedStorageThresholds)`

SetCustomThresholds sets CustomThresholds field to given value.

### HasCustomThresholds

`func (o *UndersizedStorageDetectionConfig) HasCustomThresholds() bool`

HasCustomThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


