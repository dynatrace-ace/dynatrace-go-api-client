# SyntheticMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The entity ID of the monitor. | 
**Name** | **string** | The name of the monitor. | 
**FrequencyMin** | **int32** | The frequency of the monitor, in minutes.    You can use one of the following values: &#x60;5&#x60;, &#x60;10&#x60;, &#x60;15&#x60;, &#x60;30&#x60;, and &#x60;60&#x60;. | 
**Enabled** | **bool** | The monitor is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;BROWSER&#x60; -&gt; BrowserSyntheticMonitor  * &#x60;HTTP&#x60; -&gt; HttpSyntheticMonitor   | 
**CreatedFrom** | **string** | The origin of a monitor | 
**Script** | **map[string]interface{}** | The script of a [browser](https://dt-url.net/9c103rda) or HTTP monitor. | 
**Locations** | **[]string** | A list of locations from which the monitor is executed.    To specify a location, use its entity ID. | 
**AnomalyDetection** | Pointer to [**AnomalyDetection**](AnomalyDetection.md) |  | [optional] 
**Tags** | [**[]TagWithSourceInfo**](TagWithSourceInfo.md) | A set of tags assigned to the monitor. | 
**ManagementZones** | [**[]ManagementZoneDto**](ManagementZoneDto.md) | A set of management zones to which the monitor belongs to. | 
**AutomaticallyAssignedApps** | **[]string** | A set of automatically assigned applications. | 
**ManuallyAssignedApps** | **[]string** | A set of manually assigned applications. | 

## Methods

### NewSyntheticMonitor

`func NewSyntheticMonitor(entityId string, name string, frequencyMin int32, enabled bool, type_ string, createdFrom string, script map[string]interface{}, locations []string, tags []TagWithSourceInfo, managementZones []ManagementZoneDto, automaticallyAssignedApps []string, manuallyAssignedApps []string, ) *SyntheticMonitor`

NewSyntheticMonitor instantiates a new SyntheticMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticMonitorWithDefaults

`func NewSyntheticMonitorWithDefaults() *SyntheticMonitor`

NewSyntheticMonitorWithDefaults instantiates a new SyntheticMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *SyntheticMonitor) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SyntheticMonitor) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SyntheticMonitor) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetName

`func (o *SyntheticMonitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticMonitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticMonitor) SetName(v string)`

SetName sets Name field to given value.


### GetFrequencyMin

`func (o *SyntheticMonitor) GetFrequencyMin() int32`

GetFrequencyMin returns the FrequencyMin field if non-nil, zero value otherwise.

### GetFrequencyMinOk

`func (o *SyntheticMonitor) GetFrequencyMinOk() (*int32, bool)`

GetFrequencyMinOk returns a tuple with the FrequencyMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyMin

`func (o *SyntheticMonitor) SetFrequencyMin(v int32)`

SetFrequencyMin sets FrequencyMin field to given value.


### GetEnabled

`func (o *SyntheticMonitor) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SyntheticMonitor) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SyntheticMonitor) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetType

`func (o *SyntheticMonitor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticMonitor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticMonitor) SetType(v string)`

SetType sets Type field to given value.


### GetCreatedFrom

`func (o *SyntheticMonitor) GetCreatedFrom() string`

GetCreatedFrom returns the CreatedFrom field if non-nil, zero value otherwise.

### GetCreatedFromOk

`func (o *SyntheticMonitor) GetCreatedFromOk() (*string, bool)`

GetCreatedFromOk returns a tuple with the CreatedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedFrom

`func (o *SyntheticMonitor) SetCreatedFrom(v string)`

SetCreatedFrom sets CreatedFrom field to given value.


### GetScript

`func (o *SyntheticMonitor) GetScript() map[string]interface{}`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *SyntheticMonitor) GetScriptOk() (*map[string]interface{}, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *SyntheticMonitor) SetScript(v map[string]interface{})`

SetScript sets Script field to given value.


### GetLocations

`func (o *SyntheticMonitor) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SyntheticMonitor) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SyntheticMonitor) SetLocations(v []string)`

SetLocations sets Locations field to given value.


### GetAnomalyDetection

`func (o *SyntheticMonitor) GetAnomalyDetection() AnomalyDetection`

GetAnomalyDetection returns the AnomalyDetection field if non-nil, zero value otherwise.

### GetAnomalyDetectionOk

`func (o *SyntheticMonitor) GetAnomalyDetectionOk() (*AnomalyDetection, bool)`

GetAnomalyDetectionOk returns a tuple with the AnomalyDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyDetection

`func (o *SyntheticMonitor) SetAnomalyDetection(v AnomalyDetection)`

SetAnomalyDetection sets AnomalyDetection field to given value.

### HasAnomalyDetection

`func (o *SyntheticMonitor) HasAnomalyDetection() bool`

HasAnomalyDetection returns a boolean if a field has been set.

### GetTags

`func (o *SyntheticMonitor) GetTags() []TagWithSourceInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SyntheticMonitor) GetTagsOk() (*[]TagWithSourceInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SyntheticMonitor) SetTags(v []TagWithSourceInfo)`

SetTags sets Tags field to given value.


### GetManagementZones

`func (o *SyntheticMonitor) GetManagementZones() []ManagementZoneDto`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *SyntheticMonitor) GetManagementZonesOk() (*[]ManagementZoneDto, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *SyntheticMonitor) SetManagementZones(v []ManagementZoneDto)`

SetManagementZones sets ManagementZones field to given value.


### GetAutomaticallyAssignedApps

`func (o *SyntheticMonitor) GetAutomaticallyAssignedApps() []string`

GetAutomaticallyAssignedApps returns the AutomaticallyAssignedApps field if non-nil, zero value otherwise.

### GetAutomaticallyAssignedAppsOk

`func (o *SyntheticMonitor) GetAutomaticallyAssignedAppsOk() (*[]string, bool)`

GetAutomaticallyAssignedAppsOk returns a tuple with the AutomaticallyAssignedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyAssignedApps

`func (o *SyntheticMonitor) SetAutomaticallyAssignedApps(v []string)`

SetAutomaticallyAssignedApps sets AutomaticallyAssignedApps field to given value.


### GetManuallyAssignedApps

`func (o *SyntheticMonitor) GetManuallyAssignedApps() []string`

GetManuallyAssignedApps returns the ManuallyAssignedApps field if non-nil, zero value otherwise.

### GetManuallyAssignedAppsOk

`func (o *SyntheticMonitor) GetManuallyAssignedAppsOk() (*[]string, bool)`

GetManuallyAssignedAppsOk returns a tuple with the ManuallyAssignedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyAssignedApps

`func (o *SyntheticMonitor) SetManuallyAssignedApps(v []string)`

SetManuallyAssignedApps sets ManuallyAssignedApps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


