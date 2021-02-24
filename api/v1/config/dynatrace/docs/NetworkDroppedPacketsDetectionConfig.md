# NetworkDroppedPacketsDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**CustomThresholds** | Pointer to [**NetworkDroppedPacketsThresholds**](NetworkDroppedPacketsThresholds.md) |  | [optional] 

## Methods

### NewNetworkDroppedPacketsDetectionConfig

`func NewNetworkDroppedPacketsDetectionConfig(enabled bool, ) *NetworkDroppedPacketsDetectionConfig`

NewNetworkDroppedPacketsDetectionConfig instantiates a new NetworkDroppedPacketsDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDroppedPacketsDetectionConfigWithDefaults

`func NewNetworkDroppedPacketsDetectionConfigWithDefaults() *NetworkDroppedPacketsDetectionConfig`

NewNetworkDroppedPacketsDetectionConfigWithDefaults instantiates a new NetworkDroppedPacketsDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NetworkDroppedPacketsDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkDroppedPacketsDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkDroppedPacketsDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCustomThresholds

`func (o *NetworkDroppedPacketsDetectionConfig) GetCustomThresholds() NetworkDroppedPacketsThresholds`

GetCustomThresholds returns the CustomThresholds field if non-nil, zero value otherwise.

### GetCustomThresholdsOk

`func (o *NetworkDroppedPacketsDetectionConfig) GetCustomThresholdsOk() (*NetworkDroppedPacketsThresholds, bool)`

GetCustomThresholdsOk returns a tuple with the CustomThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomThresholds

`func (o *NetworkDroppedPacketsDetectionConfig) SetCustomThresholds(v NetworkDroppedPacketsThresholds)`

SetCustomThresholds sets CustomThresholds field to given value.

### HasCustomThresholds

`func (o *NetworkDroppedPacketsDetectionConfig) HasCustomThresholds() bool`

HasCustomThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


