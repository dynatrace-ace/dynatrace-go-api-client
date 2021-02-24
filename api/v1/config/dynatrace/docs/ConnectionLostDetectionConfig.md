# ConnectionLostDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**EnabledOnGracefulShutdowns** | **bool** | Alert (&#x60;true&#x60;) on graceful host shutdowns. | 

## Methods

### NewConnectionLostDetectionConfig

`func NewConnectionLostDetectionConfig(enabled bool, enabledOnGracefulShutdowns bool, ) *ConnectionLostDetectionConfig`

NewConnectionLostDetectionConfig instantiates a new ConnectionLostDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionLostDetectionConfigWithDefaults

`func NewConnectionLostDetectionConfigWithDefaults() *ConnectionLostDetectionConfig`

NewConnectionLostDetectionConfigWithDefaults instantiates a new ConnectionLostDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ConnectionLostDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConnectionLostDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConnectionLostDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnabledOnGracefulShutdowns

`func (o *ConnectionLostDetectionConfig) GetEnabledOnGracefulShutdowns() bool`

GetEnabledOnGracefulShutdowns returns the EnabledOnGracefulShutdowns field if non-nil, zero value otherwise.

### GetEnabledOnGracefulShutdownsOk

`func (o *ConnectionLostDetectionConfig) GetEnabledOnGracefulShutdownsOk() (*bool, bool)`

GetEnabledOnGracefulShutdownsOk returns a tuple with the EnabledOnGracefulShutdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnGracefulShutdowns

`func (o *ConnectionLostDetectionConfig) SetEnabledOnGracefulShutdowns(v bool)`

SetEnabledOnGracefulShutdowns sets EnabledOnGracefulShutdowns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


