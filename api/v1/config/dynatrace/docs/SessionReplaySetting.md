# SessionReplaySetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | SessionReplay Enabled. | [optional] 
**CostControlPercentage** | Pointer to **int32** | Session replay sampling rating in percentage. | [optional] 

## Methods

### NewSessionReplaySetting

`func NewSessionReplaySetting() *SessionReplaySetting`

NewSessionReplaySetting instantiates a new SessionReplaySetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionReplaySettingWithDefaults

`func NewSessionReplaySettingWithDefaults() *SessionReplaySetting`

NewSessionReplaySettingWithDefaults instantiates a new SessionReplaySetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SessionReplaySetting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SessionReplaySetting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SessionReplaySetting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SessionReplaySetting) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCostControlPercentage

`func (o *SessionReplaySetting) GetCostControlPercentage() int32`

GetCostControlPercentage returns the CostControlPercentage field if non-nil, zero value otherwise.

### GetCostControlPercentageOk

`func (o *SessionReplaySetting) GetCostControlPercentageOk() (*int32, bool)`

GetCostControlPercentageOk returns a tuple with the CostControlPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostControlPercentage

`func (o *SessionReplaySetting) SetCostControlPercentage(v int32)`

SetCostControlPercentage sets CostControlPercentage field to given value.

### HasCostControlPercentage

`func (o *SessionReplaySetting) HasCostControlPercentage() bool`

HasCostControlPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


