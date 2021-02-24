# NetworkHighRetransmissionDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**CustomThresholds** | Pointer to [**NetworkHighRetransmissionThresholds**](NetworkHighRetransmissionThresholds.md) |  | [optional] 

## Methods

### NewNetworkHighRetransmissionDetectionConfig

`func NewNetworkHighRetransmissionDetectionConfig(enabled bool, ) *NetworkHighRetransmissionDetectionConfig`

NewNetworkHighRetransmissionDetectionConfig instantiates a new NetworkHighRetransmissionDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkHighRetransmissionDetectionConfigWithDefaults

`func NewNetworkHighRetransmissionDetectionConfigWithDefaults() *NetworkHighRetransmissionDetectionConfig`

NewNetworkHighRetransmissionDetectionConfigWithDefaults instantiates a new NetworkHighRetransmissionDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NetworkHighRetransmissionDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkHighRetransmissionDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkHighRetransmissionDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCustomThresholds

`func (o *NetworkHighRetransmissionDetectionConfig) GetCustomThresholds() NetworkHighRetransmissionThresholds`

GetCustomThresholds returns the CustomThresholds field if non-nil, zero value otherwise.

### GetCustomThresholdsOk

`func (o *NetworkHighRetransmissionDetectionConfig) GetCustomThresholdsOk() (*NetworkHighRetransmissionThresholds, bool)`

GetCustomThresholdsOk returns a tuple with the CustomThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomThresholds

`func (o *NetworkHighRetransmissionDetectionConfig) SetCustomThresholds(v NetworkHighRetransmissionThresholds)`

SetCustomThresholds sets CustomThresholds field to given value.

### HasCustomThresholds

`func (o *NetworkHighRetransmissionDetectionConfig) HasCustomThresholds() bool`

HasCustomThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


