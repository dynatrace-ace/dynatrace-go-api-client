# ContentCapture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceTimingSettings** | [**ResourceTimingSettings**](ResourceTimingSettings.md) |  | 
**JavaScriptErrors** | **bool** | JavaScript errors monitoring enabled/disabled. | 
**TimeoutSettings** | [**TimeoutSettings**](TimeoutSettings.md) |  | 
**VisuallyCompleteAndSpeedIndex** | **bool** | Visually complete and Speed index support enabled/disabled. | 
**VisuallyComplete2Settings** | Pointer to [**VisuallyComplete2Settings**](VisuallyComplete2Settings.md) |  | [optional] 

## Methods

### NewContentCapture

`func NewContentCapture(resourceTimingSettings ResourceTimingSettings, javaScriptErrors bool, timeoutSettings TimeoutSettings, visuallyCompleteAndSpeedIndex bool, ) *ContentCapture`

NewContentCapture instantiates a new ContentCapture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentCaptureWithDefaults

`func NewContentCaptureWithDefaults() *ContentCapture`

NewContentCaptureWithDefaults instantiates a new ContentCapture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceTimingSettings

`func (o *ContentCapture) GetResourceTimingSettings() ResourceTimingSettings`

GetResourceTimingSettings returns the ResourceTimingSettings field if non-nil, zero value otherwise.

### GetResourceTimingSettingsOk

`func (o *ContentCapture) GetResourceTimingSettingsOk() (*ResourceTimingSettings, bool)`

GetResourceTimingSettingsOk returns a tuple with the ResourceTimingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTimingSettings

`func (o *ContentCapture) SetResourceTimingSettings(v ResourceTimingSettings)`

SetResourceTimingSettings sets ResourceTimingSettings field to given value.


### GetJavaScriptErrors

`func (o *ContentCapture) GetJavaScriptErrors() bool`

GetJavaScriptErrors returns the JavaScriptErrors field if non-nil, zero value otherwise.

### GetJavaScriptErrorsOk

`func (o *ContentCapture) GetJavaScriptErrorsOk() (*bool, bool)`

GetJavaScriptErrorsOk returns a tuple with the JavaScriptErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaScriptErrors

`func (o *ContentCapture) SetJavaScriptErrors(v bool)`

SetJavaScriptErrors sets JavaScriptErrors field to given value.


### GetTimeoutSettings

`func (o *ContentCapture) GetTimeoutSettings() TimeoutSettings`

GetTimeoutSettings returns the TimeoutSettings field if non-nil, zero value otherwise.

### GetTimeoutSettingsOk

`func (o *ContentCapture) GetTimeoutSettingsOk() (*TimeoutSettings, bool)`

GetTimeoutSettingsOk returns a tuple with the TimeoutSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSettings

`func (o *ContentCapture) SetTimeoutSettings(v TimeoutSettings)`

SetTimeoutSettings sets TimeoutSettings field to given value.


### GetVisuallyCompleteAndSpeedIndex

`func (o *ContentCapture) GetVisuallyCompleteAndSpeedIndex() bool`

GetVisuallyCompleteAndSpeedIndex returns the VisuallyCompleteAndSpeedIndex field if non-nil, zero value otherwise.

### GetVisuallyCompleteAndSpeedIndexOk

`func (o *ContentCapture) GetVisuallyCompleteAndSpeedIndexOk() (*bool, bool)`

GetVisuallyCompleteAndSpeedIndexOk returns a tuple with the VisuallyCompleteAndSpeedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisuallyCompleteAndSpeedIndex

`func (o *ContentCapture) SetVisuallyCompleteAndSpeedIndex(v bool)`

SetVisuallyCompleteAndSpeedIndex sets VisuallyCompleteAndSpeedIndex field to given value.


### GetVisuallyComplete2Settings

`func (o *ContentCapture) GetVisuallyComplete2Settings() VisuallyComplete2Settings`

GetVisuallyComplete2Settings returns the VisuallyComplete2Settings field if non-nil, zero value otherwise.

### GetVisuallyComplete2SettingsOk

`func (o *ContentCapture) GetVisuallyComplete2SettingsOk() (*VisuallyComplete2Settings, bool)`

GetVisuallyComplete2SettingsOk returns a tuple with the VisuallyComplete2Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisuallyComplete2Settings

`func (o *ContentCapture) SetVisuallyComplete2Settings(v VisuallyComplete2Settings)`

SetVisuallyComplete2Settings sets VisuallyComplete2Settings field to given value.

### HasVisuallyComplete2Settings

`func (o *ContentCapture) HasVisuallyComplete2Settings() bool`

HasVisuallyComplete2Settings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


