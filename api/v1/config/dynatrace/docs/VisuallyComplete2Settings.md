# VisuallyComplete2Settings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageUrlBlacklist** | Pointer to **string** | A RegularExpression used to exclude images and iframes from being detected by the VC module. | [optional] 
**MutationBlacklist** | Pointer to **string** | Query selector for mutation nodes to ignore in VC and SI calculation | [optional] 
**MutationTimeout** | Pointer to **int32** | Determines the time in ms VC waits after an action closes to start calculation. Defaults to 50. | [optional] 
**InactivityTimeout** | Pointer to **int32** | The time in ms the VC module waits for no mutations happening on the page after the load action. Defaults to 1000. | [optional] 
**Threshold** | Pointer to **int32** | Minimum visible area in pixels of elements to be counted towards VC and SI. Defaults to 50. | [optional] 

## Methods

### NewVisuallyComplete2Settings

`func NewVisuallyComplete2Settings() *VisuallyComplete2Settings`

NewVisuallyComplete2Settings instantiates a new VisuallyComplete2Settings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisuallyComplete2SettingsWithDefaults

`func NewVisuallyComplete2SettingsWithDefaults() *VisuallyComplete2Settings`

NewVisuallyComplete2SettingsWithDefaults instantiates a new VisuallyComplete2Settings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageUrlBlacklist

`func (o *VisuallyComplete2Settings) GetImageUrlBlacklist() string`

GetImageUrlBlacklist returns the ImageUrlBlacklist field if non-nil, zero value otherwise.

### GetImageUrlBlacklistOk

`func (o *VisuallyComplete2Settings) GetImageUrlBlacklistOk() (*string, bool)`

GetImageUrlBlacklistOk returns a tuple with the ImageUrlBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrlBlacklist

`func (o *VisuallyComplete2Settings) SetImageUrlBlacklist(v string)`

SetImageUrlBlacklist sets ImageUrlBlacklist field to given value.

### HasImageUrlBlacklist

`func (o *VisuallyComplete2Settings) HasImageUrlBlacklist() bool`

HasImageUrlBlacklist returns a boolean if a field has been set.

### GetMutationBlacklist

`func (o *VisuallyComplete2Settings) GetMutationBlacklist() string`

GetMutationBlacklist returns the MutationBlacklist field if non-nil, zero value otherwise.

### GetMutationBlacklistOk

`func (o *VisuallyComplete2Settings) GetMutationBlacklistOk() (*string, bool)`

GetMutationBlacklistOk returns a tuple with the MutationBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutationBlacklist

`func (o *VisuallyComplete2Settings) SetMutationBlacklist(v string)`

SetMutationBlacklist sets MutationBlacklist field to given value.

### HasMutationBlacklist

`func (o *VisuallyComplete2Settings) HasMutationBlacklist() bool`

HasMutationBlacklist returns a boolean if a field has been set.

### GetMutationTimeout

`func (o *VisuallyComplete2Settings) GetMutationTimeout() int32`

GetMutationTimeout returns the MutationTimeout field if non-nil, zero value otherwise.

### GetMutationTimeoutOk

`func (o *VisuallyComplete2Settings) GetMutationTimeoutOk() (*int32, bool)`

GetMutationTimeoutOk returns a tuple with the MutationTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutationTimeout

`func (o *VisuallyComplete2Settings) SetMutationTimeout(v int32)`

SetMutationTimeout sets MutationTimeout field to given value.

### HasMutationTimeout

`func (o *VisuallyComplete2Settings) HasMutationTimeout() bool`

HasMutationTimeout returns a boolean if a field has been set.

### GetInactivityTimeout

`func (o *VisuallyComplete2Settings) GetInactivityTimeout() int32`

GetInactivityTimeout returns the InactivityTimeout field if non-nil, zero value otherwise.

### GetInactivityTimeoutOk

`func (o *VisuallyComplete2Settings) GetInactivityTimeoutOk() (*int32, bool)`

GetInactivityTimeoutOk returns a tuple with the InactivityTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityTimeout

`func (o *VisuallyComplete2Settings) SetInactivityTimeout(v int32)`

SetInactivityTimeout sets InactivityTimeout field to given value.

### HasInactivityTimeout

`func (o *VisuallyComplete2Settings) HasInactivityTimeout() bool`

HasInactivityTimeout returns a boolean if a field has been set.

### GetThreshold

`func (o *VisuallyComplete2Settings) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *VisuallyComplete2Settings) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *VisuallyComplete2Settings) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *VisuallyComplete2Settings) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


