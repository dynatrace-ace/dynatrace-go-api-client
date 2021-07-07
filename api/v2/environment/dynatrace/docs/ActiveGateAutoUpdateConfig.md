# ActiveGateAutoUpdateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Setting** | **string** | The state of the ActiveGate auto-update: enabled, disabled, or inherited.   If set to &#x60;INHERITED&#x60;, the setting is inherited from the global configuration set on the environment or Managed cluster level. | 
**EffectiveSetting** | Pointer to **string** | The actual state of the ActiveGate auto-update.   Applicable only if the **setting** parameter is set to &#x60;INHERITED&#x60;. In that case, the value is taken from the parent setting. Otherwise, it&#39;s just a duplicate of the **setting** value. | [optional] [readonly] 

## Methods

### NewActiveGateAutoUpdateConfig

`func NewActiveGateAutoUpdateConfig(setting string, ) *ActiveGateAutoUpdateConfig`

NewActiveGateAutoUpdateConfig instantiates a new ActiveGateAutoUpdateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveGateAutoUpdateConfigWithDefaults

`func NewActiveGateAutoUpdateConfigWithDefaults() *ActiveGateAutoUpdateConfig`

NewActiveGateAutoUpdateConfigWithDefaults instantiates a new ActiveGateAutoUpdateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSetting

`func (o *ActiveGateAutoUpdateConfig) GetSetting() string`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *ActiveGateAutoUpdateConfig) GetSettingOk() (*string, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *ActiveGateAutoUpdateConfig) SetSetting(v string)`

SetSetting sets Setting field to given value.


### GetEffectiveSetting

`func (o *ActiveGateAutoUpdateConfig) GetEffectiveSetting() string`

GetEffectiveSetting returns the EffectiveSetting field if non-nil, zero value otherwise.

### GetEffectiveSettingOk

`func (o *ActiveGateAutoUpdateConfig) GetEffectiveSettingOk() (*string, bool)`

GetEffectiveSettingOk returns a tuple with the EffectiveSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveSetting

`func (o *ActiveGateAutoUpdateConfig) SetEffectiveSetting(v string)`

SetEffectiveSetting sets EffectiveSetting field to given value.

### HasEffectiveSetting

`func (o *ActiveGateAutoUpdateConfig) HasEffectiveSetting() bool`

HasEffectiveSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


