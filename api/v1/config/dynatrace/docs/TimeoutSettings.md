# TimeoutSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimedActionSupport** | **bool** | Timed action support enabled/disabled.   Enable to detect actions that trigger sending of XHRs via *setTimout* methods. | 
**TemporaryActionLimit** | **int32** | Defines how deep temporary actions may cascade. 0 disables temporary actions completely. Recommended value if enabled is 3. | 
**TemporaryActionTotalTimeout** | **int32** | The total timeout of all cascaded timeouts that should still be able to create a temporary action | 

## Methods

### NewTimeoutSettings

`func NewTimeoutSettings(timedActionSupport bool, temporaryActionLimit int32, temporaryActionTotalTimeout int32, ) *TimeoutSettings`

NewTimeoutSettings instantiates a new TimeoutSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeoutSettingsWithDefaults

`func NewTimeoutSettingsWithDefaults() *TimeoutSettings`

NewTimeoutSettingsWithDefaults instantiates a new TimeoutSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimedActionSupport

`func (o *TimeoutSettings) GetTimedActionSupport() bool`

GetTimedActionSupport returns the TimedActionSupport field if non-nil, zero value otherwise.

### GetTimedActionSupportOk

`func (o *TimeoutSettings) GetTimedActionSupportOk() (*bool, bool)`

GetTimedActionSupportOk returns a tuple with the TimedActionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedActionSupport

`func (o *TimeoutSettings) SetTimedActionSupport(v bool)`

SetTimedActionSupport sets TimedActionSupport field to given value.


### GetTemporaryActionLimit

`func (o *TimeoutSettings) GetTemporaryActionLimit() int32`

GetTemporaryActionLimit returns the TemporaryActionLimit field if non-nil, zero value otherwise.

### GetTemporaryActionLimitOk

`func (o *TimeoutSettings) GetTemporaryActionLimitOk() (*int32, bool)`

GetTemporaryActionLimitOk returns a tuple with the TemporaryActionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryActionLimit

`func (o *TimeoutSettings) SetTemporaryActionLimit(v int32)`

SetTemporaryActionLimit sets TemporaryActionLimit field to given value.


### GetTemporaryActionTotalTimeout

`func (o *TimeoutSettings) GetTemporaryActionTotalTimeout() int32`

GetTemporaryActionTotalTimeout returns the TemporaryActionTotalTimeout field if non-nil, zero value otherwise.

### GetTemporaryActionTotalTimeoutOk

`func (o *TimeoutSettings) GetTemporaryActionTotalTimeoutOk() (*int32, bool)`

GetTemporaryActionTotalTimeoutOk returns a tuple with the TemporaryActionTotalTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryActionTotalTimeout

`func (o *TimeoutSettings) SetTemporaryActionTotalTimeout(v int32)`

SetTemporaryActionTotalTimeout sets TemporaryActionTotalTimeout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


