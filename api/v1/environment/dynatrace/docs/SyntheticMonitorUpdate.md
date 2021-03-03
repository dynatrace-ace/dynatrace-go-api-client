# SyntheticMonitorUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrequencyMin** | **int32** | The frequency of the monitor, in minutes.    You can use one of the following values: &#x60;5&#x60;, &#x60;10&#x60;, &#x60;15&#x60;, &#x60;30&#x60;, and &#x60;60&#x60;. | 
**AnomalyDetection** | Pointer to [**AnomalyDetection**](AnomalyDetection.md) |  | [optional] 
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;BROWSER&#x60; -&gt; BrowserSyntheticMonitorUpdate  * &#x60;HTTP&#x60; -&gt; HttpSyntheticMonitorUpdate   | 
**Name** | **string** | The name of the monitor. | 
**Locations** | **[]string** | A list of locations from which the monitor is executed.    To specify a location, use its entity ID. | 
**Enabled** | **bool** | The monitor is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**Script** | **map[string]interface{}** | The script of a [browser](https://dt-url.net/9c103rda) or HTTP monitor. | 
**Tags** | [**[]TagWithSourceInfo**](TagWithSourceInfo.md) | A set of tags assigned to the monitor.    You can specify only the value of the tag here and the &#x60;CONTEXTLESS&#x60; context and source &#39;USER&#39; will be added automatically. But preferred option is usage of TagWithSourceDto model. | 
**ManuallyAssignedApps** | **[]string** | A set of manually assigned applications. | 

## Methods

### NewSyntheticMonitorUpdate

`func NewSyntheticMonitorUpdate(frequencyMin int32, type_ string, name string, locations []string, enabled bool, script map[string]interface{}, tags []TagWithSourceInfo, manuallyAssignedApps []string, ) *SyntheticMonitorUpdate`

NewSyntheticMonitorUpdate instantiates a new SyntheticMonitorUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticMonitorUpdateWithDefaults

`func NewSyntheticMonitorUpdateWithDefaults() *SyntheticMonitorUpdate`

NewSyntheticMonitorUpdateWithDefaults instantiates a new SyntheticMonitorUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequencyMin

`func (o *SyntheticMonitorUpdate) GetFrequencyMin() int32`

GetFrequencyMin returns the FrequencyMin field if non-nil, zero value otherwise.

### GetFrequencyMinOk

`func (o *SyntheticMonitorUpdate) GetFrequencyMinOk() (*int32, bool)`

GetFrequencyMinOk returns a tuple with the FrequencyMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyMin

`func (o *SyntheticMonitorUpdate) SetFrequencyMin(v int32)`

SetFrequencyMin sets FrequencyMin field to given value.


### GetAnomalyDetection

`func (o *SyntheticMonitorUpdate) GetAnomalyDetection() AnomalyDetection`

GetAnomalyDetection returns the AnomalyDetection field if non-nil, zero value otherwise.

### GetAnomalyDetectionOk

`func (o *SyntheticMonitorUpdate) GetAnomalyDetectionOk() (*AnomalyDetection, bool)`

GetAnomalyDetectionOk returns a tuple with the AnomalyDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyDetection

`func (o *SyntheticMonitorUpdate) SetAnomalyDetection(v AnomalyDetection)`

SetAnomalyDetection sets AnomalyDetection field to given value.

### HasAnomalyDetection

`func (o *SyntheticMonitorUpdate) HasAnomalyDetection() bool`

HasAnomalyDetection returns a boolean if a field has been set.

### GetType

`func (o *SyntheticMonitorUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticMonitorUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticMonitorUpdate) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *SyntheticMonitorUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticMonitorUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticMonitorUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetLocations

`func (o *SyntheticMonitorUpdate) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SyntheticMonitorUpdate) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SyntheticMonitorUpdate) SetLocations(v []string)`

SetLocations sets Locations field to given value.


### GetEnabled

`func (o *SyntheticMonitorUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SyntheticMonitorUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SyntheticMonitorUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetScript

`func (o *SyntheticMonitorUpdate) GetScript() map[string]interface{}`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *SyntheticMonitorUpdate) GetScriptOk() (*map[string]interface{}, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *SyntheticMonitorUpdate) SetScript(v map[string]interface{})`

SetScript sets Script field to given value.


### GetTags

`func (o *SyntheticMonitorUpdate) GetTags() []TagWithSourceInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SyntheticMonitorUpdate) GetTagsOk() (*[]TagWithSourceInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SyntheticMonitorUpdate) SetTags(v []TagWithSourceInfo)`

SetTags sets Tags field to given value.


### GetManuallyAssignedApps

`func (o *SyntheticMonitorUpdate) GetManuallyAssignedApps() []string`

GetManuallyAssignedApps returns the ManuallyAssignedApps field if non-nil, zero value otherwise.

### GetManuallyAssignedAppsOk

`func (o *SyntheticMonitorUpdate) GetManuallyAssignedAppsOk() (*[]string, bool)`

GetManuallyAssignedAppsOk returns a tuple with the ManuallyAssignedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyAssignedApps

`func (o *SyntheticMonitorUpdate) SetManuallyAssignedApps(v []string)`

SetManuallyAssignedApps sets ManuallyAssignedApps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


