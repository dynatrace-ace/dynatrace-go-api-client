# AdvancedJavaScriptTagSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncBeaconFirefox** | Pointer to **bool** | Send the beacon signal as a synchronous XMLHttpRequest using Firefox enabled/disabled. | [optional] 
**SyncBeaconInternetExplorer** | Pointer to **bool** | Send the beacon signal as a synchronous XMLHttpRequest using Internet Explorer enabled/disabled. | [optional] 
**InstrumentUnsupportedAjaxFrameworks** | **bool** | Instrumentation of unsupported Ajax frameworks enabled/disabled. | 
**SpecialCharactersToEscape** | **string** | Additional special characters that are to be escaped using non-alphanumeric characters in HTML escape format. | 
**MaxActionNameLength** | **int32** | Maximum character length for action names. Valid values range from 5 to 10000. | 
**MaxErrorsToCapture** | **int32** | Maximum number of errors to be captured per page. Valid values range from 0 to 50. | 
**AdditionalEventHandlers** | [**AdditionalEventHandlers**](AdditionalEventHandlers.md) |  | 
**EventWrapperSettings** | [**EventWrapperSettings**](EventWrapperSettings.md) |  | 
**GlobalEventCaptureSettings** | [**GlobalEventCaptureSettings**](GlobalEventCaptureSettings.md) |  | 

## Methods

### NewAdvancedJavaScriptTagSettings

`func NewAdvancedJavaScriptTagSettings(instrumentUnsupportedAjaxFrameworks bool, specialCharactersToEscape string, maxActionNameLength int32, maxErrorsToCapture int32, additionalEventHandlers AdditionalEventHandlers, eventWrapperSettings EventWrapperSettings, globalEventCaptureSettings GlobalEventCaptureSettings, ) *AdvancedJavaScriptTagSettings`

NewAdvancedJavaScriptTagSettings instantiates a new AdvancedJavaScriptTagSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedJavaScriptTagSettingsWithDefaults

`func NewAdvancedJavaScriptTagSettingsWithDefaults() *AdvancedJavaScriptTagSettings`

NewAdvancedJavaScriptTagSettingsWithDefaults instantiates a new AdvancedJavaScriptTagSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyncBeaconFirefox

`func (o *AdvancedJavaScriptTagSettings) GetSyncBeaconFirefox() bool`

GetSyncBeaconFirefox returns the SyncBeaconFirefox field if non-nil, zero value otherwise.

### GetSyncBeaconFirefoxOk

`func (o *AdvancedJavaScriptTagSettings) GetSyncBeaconFirefoxOk() (*bool, bool)`

GetSyncBeaconFirefoxOk returns a tuple with the SyncBeaconFirefox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncBeaconFirefox

`func (o *AdvancedJavaScriptTagSettings) SetSyncBeaconFirefox(v bool)`

SetSyncBeaconFirefox sets SyncBeaconFirefox field to given value.

### HasSyncBeaconFirefox

`func (o *AdvancedJavaScriptTagSettings) HasSyncBeaconFirefox() bool`

HasSyncBeaconFirefox returns a boolean if a field has been set.

### GetSyncBeaconInternetExplorer

`func (o *AdvancedJavaScriptTagSettings) GetSyncBeaconInternetExplorer() bool`

GetSyncBeaconInternetExplorer returns the SyncBeaconInternetExplorer field if non-nil, zero value otherwise.

### GetSyncBeaconInternetExplorerOk

`func (o *AdvancedJavaScriptTagSettings) GetSyncBeaconInternetExplorerOk() (*bool, bool)`

GetSyncBeaconInternetExplorerOk returns a tuple with the SyncBeaconInternetExplorer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncBeaconInternetExplorer

`func (o *AdvancedJavaScriptTagSettings) SetSyncBeaconInternetExplorer(v bool)`

SetSyncBeaconInternetExplorer sets SyncBeaconInternetExplorer field to given value.

### HasSyncBeaconInternetExplorer

`func (o *AdvancedJavaScriptTagSettings) HasSyncBeaconInternetExplorer() bool`

HasSyncBeaconInternetExplorer returns a boolean if a field has been set.

### GetInstrumentUnsupportedAjaxFrameworks

`func (o *AdvancedJavaScriptTagSettings) GetInstrumentUnsupportedAjaxFrameworks() bool`

GetInstrumentUnsupportedAjaxFrameworks returns the InstrumentUnsupportedAjaxFrameworks field if non-nil, zero value otherwise.

### GetInstrumentUnsupportedAjaxFrameworksOk

`func (o *AdvancedJavaScriptTagSettings) GetInstrumentUnsupportedAjaxFrameworksOk() (*bool, bool)`

GetInstrumentUnsupportedAjaxFrameworksOk returns a tuple with the InstrumentUnsupportedAjaxFrameworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentUnsupportedAjaxFrameworks

`func (o *AdvancedJavaScriptTagSettings) SetInstrumentUnsupportedAjaxFrameworks(v bool)`

SetInstrumentUnsupportedAjaxFrameworks sets InstrumentUnsupportedAjaxFrameworks field to given value.


### GetSpecialCharactersToEscape

`func (o *AdvancedJavaScriptTagSettings) GetSpecialCharactersToEscape() string`

GetSpecialCharactersToEscape returns the SpecialCharactersToEscape field if non-nil, zero value otherwise.

### GetSpecialCharactersToEscapeOk

`func (o *AdvancedJavaScriptTagSettings) GetSpecialCharactersToEscapeOk() (*string, bool)`

GetSpecialCharactersToEscapeOk returns a tuple with the SpecialCharactersToEscape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialCharactersToEscape

`func (o *AdvancedJavaScriptTagSettings) SetSpecialCharactersToEscape(v string)`

SetSpecialCharactersToEscape sets SpecialCharactersToEscape field to given value.


### GetMaxActionNameLength

`func (o *AdvancedJavaScriptTagSettings) GetMaxActionNameLength() int32`

GetMaxActionNameLength returns the MaxActionNameLength field if non-nil, zero value otherwise.

### GetMaxActionNameLengthOk

`func (o *AdvancedJavaScriptTagSettings) GetMaxActionNameLengthOk() (*int32, bool)`

GetMaxActionNameLengthOk returns a tuple with the MaxActionNameLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActionNameLength

`func (o *AdvancedJavaScriptTagSettings) SetMaxActionNameLength(v int32)`

SetMaxActionNameLength sets MaxActionNameLength field to given value.


### GetMaxErrorsToCapture

`func (o *AdvancedJavaScriptTagSettings) GetMaxErrorsToCapture() int32`

GetMaxErrorsToCapture returns the MaxErrorsToCapture field if non-nil, zero value otherwise.

### GetMaxErrorsToCaptureOk

`func (o *AdvancedJavaScriptTagSettings) GetMaxErrorsToCaptureOk() (*int32, bool)`

GetMaxErrorsToCaptureOk returns a tuple with the MaxErrorsToCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxErrorsToCapture

`func (o *AdvancedJavaScriptTagSettings) SetMaxErrorsToCapture(v int32)`

SetMaxErrorsToCapture sets MaxErrorsToCapture field to given value.


### GetAdditionalEventHandlers

`func (o *AdvancedJavaScriptTagSettings) GetAdditionalEventHandlers() AdditionalEventHandlers`

GetAdditionalEventHandlers returns the AdditionalEventHandlers field if non-nil, zero value otherwise.

### GetAdditionalEventHandlersOk

`func (o *AdvancedJavaScriptTagSettings) GetAdditionalEventHandlersOk() (*AdditionalEventHandlers, bool)`

GetAdditionalEventHandlersOk returns a tuple with the AdditionalEventHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEventHandlers

`func (o *AdvancedJavaScriptTagSettings) SetAdditionalEventHandlers(v AdditionalEventHandlers)`

SetAdditionalEventHandlers sets AdditionalEventHandlers field to given value.


### GetEventWrapperSettings

`func (o *AdvancedJavaScriptTagSettings) GetEventWrapperSettings() EventWrapperSettings`

GetEventWrapperSettings returns the EventWrapperSettings field if non-nil, zero value otherwise.

### GetEventWrapperSettingsOk

`func (o *AdvancedJavaScriptTagSettings) GetEventWrapperSettingsOk() (*EventWrapperSettings, bool)`

GetEventWrapperSettingsOk returns a tuple with the EventWrapperSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventWrapperSettings

`func (o *AdvancedJavaScriptTagSettings) SetEventWrapperSettings(v EventWrapperSettings)`

SetEventWrapperSettings sets EventWrapperSettings field to given value.


### GetGlobalEventCaptureSettings

`func (o *AdvancedJavaScriptTagSettings) GetGlobalEventCaptureSettings() GlobalEventCaptureSettings`

GetGlobalEventCaptureSettings returns the GlobalEventCaptureSettings field if non-nil, zero value otherwise.

### GetGlobalEventCaptureSettingsOk

`func (o *AdvancedJavaScriptTagSettings) GetGlobalEventCaptureSettingsOk() (*GlobalEventCaptureSettings, bool)`

GetGlobalEventCaptureSettingsOk returns a tuple with the GlobalEventCaptureSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalEventCaptureSettings

`func (o *AdvancedJavaScriptTagSettings) SetGlobalEventCaptureSettings(v GlobalEventCaptureSettings)`

SetGlobalEventCaptureSettings sets GlobalEventCaptureSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


