# WebApplicationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Identifier** | Pointer to **string** | Dynatrace entity ID of the web application. | [optional] [readonly] 
**Name** | **string** | The name of the web application, displayed in the UI. | 
**Type** | Pointer to **string** | The type of the web application. | [optional] 
**RealUserMonitoringEnabled** | **bool** | Real user monitoring enabled/disabled. | 
**CostControlUserSessionPercentage** | **float32** | Analize *X*% of user sessions. | 
**LoadActionKeyPerformanceMetric** | **string** | The key performance metric of load actions. | 
**SessionReplayConfig** | Pointer to [**SessionReplaySetting**](SessionReplaySetting.md) |  | [optional] 
**XhrActionKeyPerformanceMetric** | **string** | The key performance metric of XHR actions. | 
**LoadActionApdexSettings** | [**Apdex**](Apdex.md) |  | 
**XhrActionApdexSettings** | [**Apdex**](Apdex.md) |  | 
**CustomActionApdexSettings** | [**Apdex**](Apdex.md) |  | 
**WaterfallSettings** | [**WaterfallSettings**](WaterfallSettings.md) |  | 
**MonitoringSettings** | [**MonitoringSettings**](MonitoringSettings.md) |  | 
**UserTags** | Pointer to [**[]UserTag**](UserTag.md) | User tags settings. | [optional] 
**UserActionAndSessionProperties** | Pointer to [**[]UserActionAndSessionProperties**](UserActionAndSessionProperties.md) | User action and session properties settings. Empty List means no change | [optional] 
**UserActionNamingSettings** | Pointer to [**UserActionNamingSettings**](UserActionNamingSettings.md) |  | [optional] 
**MetaDataCaptureSettings** | Pointer to [**[]MetaDataCapturing**](MetaDataCapturing.md) | Java script agent meta data capture settings. | [optional] 
**ConversionGoals** | Pointer to [**[]ConversionGoal**](ConversionGoal.md) | A list of conversion goals of the application. | [optional] 
**UrlInjectionPattern** | Pointer to **string** | Url injection pattern for manual web application. | [optional] 

## Methods

### NewWebApplicationConfig

`func NewWebApplicationConfig(name string, realUserMonitoringEnabled bool, costControlUserSessionPercentage float32, loadActionKeyPerformanceMetric string, xhrActionKeyPerformanceMetric string, loadActionApdexSettings Apdex, xhrActionApdexSettings Apdex, customActionApdexSettings Apdex, waterfallSettings WaterfallSettings, monitoringSettings MonitoringSettings, ) *WebApplicationConfig`

NewWebApplicationConfig instantiates a new WebApplicationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebApplicationConfigWithDefaults

`func NewWebApplicationConfigWithDefaults() *WebApplicationConfig`

NewWebApplicationConfigWithDefaults instantiates a new WebApplicationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *WebApplicationConfig) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebApplicationConfig) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebApplicationConfig) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebApplicationConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetIdentifier

`func (o *WebApplicationConfig) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *WebApplicationConfig) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *WebApplicationConfig) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *WebApplicationConfig) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *WebApplicationConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebApplicationConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebApplicationConfig) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *WebApplicationConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebApplicationConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebApplicationConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WebApplicationConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRealUserMonitoringEnabled

`func (o *WebApplicationConfig) GetRealUserMonitoringEnabled() bool`

GetRealUserMonitoringEnabled returns the RealUserMonitoringEnabled field if non-nil, zero value otherwise.

### GetRealUserMonitoringEnabledOk

`func (o *WebApplicationConfig) GetRealUserMonitoringEnabledOk() (*bool, bool)`

GetRealUserMonitoringEnabledOk returns a tuple with the RealUserMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealUserMonitoringEnabled

`func (o *WebApplicationConfig) SetRealUserMonitoringEnabled(v bool)`

SetRealUserMonitoringEnabled sets RealUserMonitoringEnabled field to given value.


### GetCostControlUserSessionPercentage

`func (o *WebApplicationConfig) GetCostControlUserSessionPercentage() float32`

GetCostControlUserSessionPercentage returns the CostControlUserSessionPercentage field if non-nil, zero value otherwise.

### GetCostControlUserSessionPercentageOk

`func (o *WebApplicationConfig) GetCostControlUserSessionPercentageOk() (*float32, bool)`

GetCostControlUserSessionPercentageOk returns a tuple with the CostControlUserSessionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostControlUserSessionPercentage

`func (o *WebApplicationConfig) SetCostControlUserSessionPercentage(v float32)`

SetCostControlUserSessionPercentage sets CostControlUserSessionPercentage field to given value.


### GetLoadActionKeyPerformanceMetric

`func (o *WebApplicationConfig) GetLoadActionKeyPerformanceMetric() string`

GetLoadActionKeyPerformanceMetric returns the LoadActionKeyPerformanceMetric field if non-nil, zero value otherwise.

### GetLoadActionKeyPerformanceMetricOk

`func (o *WebApplicationConfig) GetLoadActionKeyPerformanceMetricOk() (*string, bool)`

GetLoadActionKeyPerformanceMetricOk returns a tuple with the LoadActionKeyPerformanceMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadActionKeyPerformanceMetric

`func (o *WebApplicationConfig) SetLoadActionKeyPerformanceMetric(v string)`

SetLoadActionKeyPerformanceMetric sets LoadActionKeyPerformanceMetric field to given value.


### GetSessionReplayConfig

`func (o *WebApplicationConfig) GetSessionReplayConfig() SessionReplaySetting`

GetSessionReplayConfig returns the SessionReplayConfig field if non-nil, zero value otherwise.

### GetSessionReplayConfigOk

`func (o *WebApplicationConfig) GetSessionReplayConfigOk() (*SessionReplaySetting, bool)`

GetSessionReplayConfigOk returns a tuple with the SessionReplayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionReplayConfig

`func (o *WebApplicationConfig) SetSessionReplayConfig(v SessionReplaySetting)`

SetSessionReplayConfig sets SessionReplayConfig field to given value.

### HasSessionReplayConfig

`func (o *WebApplicationConfig) HasSessionReplayConfig() bool`

HasSessionReplayConfig returns a boolean if a field has been set.

### GetXhrActionKeyPerformanceMetric

`func (o *WebApplicationConfig) GetXhrActionKeyPerformanceMetric() string`

GetXhrActionKeyPerformanceMetric returns the XhrActionKeyPerformanceMetric field if non-nil, zero value otherwise.

### GetXhrActionKeyPerformanceMetricOk

`func (o *WebApplicationConfig) GetXhrActionKeyPerformanceMetricOk() (*string, bool)`

GetXhrActionKeyPerformanceMetricOk returns a tuple with the XhrActionKeyPerformanceMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXhrActionKeyPerformanceMetric

`func (o *WebApplicationConfig) SetXhrActionKeyPerformanceMetric(v string)`

SetXhrActionKeyPerformanceMetric sets XhrActionKeyPerformanceMetric field to given value.


### GetLoadActionApdexSettings

`func (o *WebApplicationConfig) GetLoadActionApdexSettings() Apdex`

GetLoadActionApdexSettings returns the LoadActionApdexSettings field if non-nil, zero value otherwise.

### GetLoadActionApdexSettingsOk

`func (o *WebApplicationConfig) GetLoadActionApdexSettingsOk() (*Apdex, bool)`

GetLoadActionApdexSettingsOk returns a tuple with the LoadActionApdexSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadActionApdexSettings

`func (o *WebApplicationConfig) SetLoadActionApdexSettings(v Apdex)`

SetLoadActionApdexSettings sets LoadActionApdexSettings field to given value.


### GetXhrActionApdexSettings

`func (o *WebApplicationConfig) GetXhrActionApdexSettings() Apdex`

GetXhrActionApdexSettings returns the XhrActionApdexSettings field if non-nil, zero value otherwise.

### GetXhrActionApdexSettingsOk

`func (o *WebApplicationConfig) GetXhrActionApdexSettingsOk() (*Apdex, bool)`

GetXhrActionApdexSettingsOk returns a tuple with the XhrActionApdexSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXhrActionApdexSettings

`func (o *WebApplicationConfig) SetXhrActionApdexSettings(v Apdex)`

SetXhrActionApdexSettings sets XhrActionApdexSettings field to given value.


### GetCustomActionApdexSettings

`func (o *WebApplicationConfig) GetCustomActionApdexSettings() Apdex`

GetCustomActionApdexSettings returns the CustomActionApdexSettings field if non-nil, zero value otherwise.

### GetCustomActionApdexSettingsOk

`func (o *WebApplicationConfig) GetCustomActionApdexSettingsOk() (*Apdex, bool)`

GetCustomActionApdexSettingsOk returns a tuple with the CustomActionApdexSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomActionApdexSettings

`func (o *WebApplicationConfig) SetCustomActionApdexSettings(v Apdex)`

SetCustomActionApdexSettings sets CustomActionApdexSettings field to given value.


### GetWaterfallSettings

`func (o *WebApplicationConfig) GetWaterfallSettings() WaterfallSettings`

GetWaterfallSettings returns the WaterfallSettings field if non-nil, zero value otherwise.

### GetWaterfallSettingsOk

`func (o *WebApplicationConfig) GetWaterfallSettingsOk() (*WaterfallSettings, bool)`

GetWaterfallSettingsOk returns a tuple with the WaterfallSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaterfallSettings

`func (o *WebApplicationConfig) SetWaterfallSettings(v WaterfallSettings)`

SetWaterfallSettings sets WaterfallSettings field to given value.


### GetMonitoringSettings

`func (o *WebApplicationConfig) GetMonitoringSettings() MonitoringSettings`

GetMonitoringSettings returns the MonitoringSettings field if non-nil, zero value otherwise.

### GetMonitoringSettingsOk

`func (o *WebApplicationConfig) GetMonitoringSettingsOk() (*MonitoringSettings, bool)`

GetMonitoringSettingsOk returns a tuple with the MonitoringSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringSettings

`func (o *WebApplicationConfig) SetMonitoringSettings(v MonitoringSettings)`

SetMonitoringSettings sets MonitoringSettings field to given value.


### GetUserTags

`func (o *WebApplicationConfig) GetUserTags() []UserTag`

GetUserTags returns the UserTags field if non-nil, zero value otherwise.

### GetUserTagsOk

`func (o *WebApplicationConfig) GetUserTagsOk() (*[]UserTag, bool)`

GetUserTagsOk returns a tuple with the UserTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTags

`func (o *WebApplicationConfig) SetUserTags(v []UserTag)`

SetUserTags sets UserTags field to given value.

### HasUserTags

`func (o *WebApplicationConfig) HasUserTags() bool`

HasUserTags returns a boolean if a field has been set.

### GetUserActionAndSessionProperties

`func (o *WebApplicationConfig) GetUserActionAndSessionProperties() []UserActionAndSessionProperties`

GetUserActionAndSessionProperties returns the UserActionAndSessionProperties field if non-nil, zero value otherwise.

### GetUserActionAndSessionPropertiesOk

`func (o *WebApplicationConfig) GetUserActionAndSessionPropertiesOk() (*[]UserActionAndSessionProperties, bool)`

GetUserActionAndSessionPropertiesOk returns a tuple with the UserActionAndSessionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionAndSessionProperties

`func (o *WebApplicationConfig) SetUserActionAndSessionProperties(v []UserActionAndSessionProperties)`

SetUserActionAndSessionProperties sets UserActionAndSessionProperties field to given value.

### HasUserActionAndSessionProperties

`func (o *WebApplicationConfig) HasUserActionAndSessionProperties() bool`

HasUserActionAndSessionProperties returns a boolean if a field has been set.

### GetUserActionNamingSettings

`func (o *WebApplicationConfig) GetUserActionNamingSettings() UserActionNamingSettings`

GetUserActionNamingSettings returns the UserActionNamingSettings field if non-nil, zero value otherwise.

### GetUserActionNamingSettingsOk

`func (o *WebApplicationConfig) GetUserActionNamingSettingsOk() (*UserActionNamingSettings, bool)`

GetUserActionNamingSettingsOk returns a tuple with the UserActionNamingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionNamingSettings

`func (o *WebApplicationConfig) SetUserActionNamingSettings(v UserActionNamingSettings)`

SetUserActionNamingSettings sets UserActionNamingSettings field to given value.

### HasUserActionNamingSettings

`func (o *WebApplicationConfig) HasUserActionNamingSettings() bool`

HasUserActionNamingSettings returns a boolean if a field has been set.

### GetMetaDataCaptureSettings

`func (o *WebApplicationConfig) GetMetaDataCaptureSettings() []MetaDataCapturing`

GetMetaDataCaptureSettings returns the MetaDataCaptureSettings field if non-nil, zero value otherwise.

### GetMetaDataCaptureSettingsOk

`func (o *WebApplicationConfig) GetMetaDataCaptureSettingsOk() (*[]MetaDataCapturing, bool)`

GetMetaDataCaptureSettingsOk returns a tuple with the MetaDataCaptureSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDataCaptureSettings

`func (o *WebApplicationConfig) SetMetaDataCaptureSettings(v []MetaDataCapturing)`

SetMetaDataCaptureSettings sets MetaDataCaptureSettings field to given value.

### HasMetaDataCaptureSettings

`func (o *WebApplicationConfig) HasMetaDataCaptureSettings() bool`

HasMetaDataCaptureSettings returns a boolean if a field has been set.

### GetConversionGoals

`func (o *WebApplicationConfig) GetConversionGoals() []ConversionGoal`

GetConversionGoals returns the ConversionGoals field if non-nil, zero value otherwise.

### GetConversionGoalsOk

`func (o *WebApplicationConfig) GetConversionGoalsOk() (*[]ConversionGoal, bool)`

GetConversionGoalsOk returns a tuple with the ConversionGoals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoals

`func (o *WebApplicationConfig) SetConversionGoals(v []ConversionGoal)`

SetConversionGoals sets ConversionGoals field to given value.

### HasConversionGoals

`func (o *WebApplicationConfig) HasConversionGoals() bool`

HasConversionGoals returns a boolean if a field has been set.

### GetUrlInjectionPattern

`func (o *WebApplicationConfig) GetUrlInjectionPattern() string`

GetUrlInjectionPattern returns the UrlInjectionPattern field if non-nil, zero value otherwise.

### GetUrlInjectionPatternOk

`func (o *WebApplicationConfig) GetUrlInjectionPatternOk() (*string, bool)`

GetUrlInjectionPatternOk returns a tuple with the UrlInjectionPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlInjectionPattern

`func (o *WebApplicationConfig) SetUrlInjectionPattern(v string)`

SetUrlInjectionPattern sets UrlInjectionPattern field to given value.

### HasUrlInjectionPattern

`func (o *WebApplicationConfig) HasUrlInjectionPattern() bool`

HasUrlInjectionPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


