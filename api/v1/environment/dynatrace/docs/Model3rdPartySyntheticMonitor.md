# Model3rdPartySyntheticMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the monitor. | 
**Title** | **string** | The name of the monitor. | 
**Description** | Pointer to **string** | A description of the monitor. | [optional] 
**TestSetup** | Pointer to **string** | The information on monitor setup, for example &#x60;browser&#x60;. | [optional] 
**ExpirationTimestamp** | Pointer to **int64** | The timestamp of the monitor expiration, in UTC milliseconds. | [optional] 
**DrilldownLink** | Pointer to **string** | The URL to the results of monitor execution. | [optional] 
**EditLink** | Pointer to **string** | The URL to edit the monitor in the original UI. | [optional] 
**Enabled** | Pointer to **bool** | The monitor is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). Default is &#x60;true&#x60;.   If &#x60;true&#x60;, set the **deleted** parameter to &#x60;false&#x60;. | [optional] 
**Deleted** | Pointer to **bool** | The flag of the deleted monitor. Default is &#x60;false&#x60;.    If &#x60;true&#x60;, set the **enabled** parameter to &#x60;false&#x60;. | [optional] 
**Locations** | [**[]SyntheticTestLocation**](SyntheticTestLocation.md) | Locations from which the synthetic monitor runs. | 
**Steps** | Pointer to [**[]SyntheticTestStep**](SyntheticTestStep.md) | Steps of the third-party monitor. | [optional] 
**ScheduleIntervalInSeconds** | **int32** | The frequency of the monitor, in seconds. The monitor is repeated with the specified interval at the third-party source.   Dynatrace expects results of a monitor execution with the specified interval. If you report results to Dynatrace less often, adjust the **noDataTimeout** value accordingly. | 
**NoDataTimeout** | Pointer to **int32** | The timeout of the monitor, in seconds. If no result is reported within this time, the availability state switches to unmonitored. Default is doubled frequency of the monitor. | [optional] 

## Methods

### NewModel3rdPartySyntheticMonitor

`func NewModel3rdPartySyntheticMonitor(id string, title string, locations []SyntheticTestLocation, scheduleIntervalInSeconds int32, ) *Model3rdPartySyntheticMonitor`

NewModel3rdPartySyntheticMonitor instantiates a new Model3rdPartySyntheticMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel3rdPartySyntheticMonitorWithDefaults

`func NewModel3rdPartySyntheticMonitorWithDefaults() *Model3rdPartySyntheticMonitor`

NewModel3rdPartySyntheticMonitorWithDefaults instantiates a new Model3rdPartySyntheticMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Model3rdPartySyntheticMonitor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Model3rdPartySyntheticMonitor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Model3rdPartySyntheticMonitor) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *Model3rdPartySyntheticMonitor) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Model3rdPartySyntheticMonitor) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Model3rdPartySyntheticMonitor) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Model3rdPartySyntheticMonitor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Model3rdPartySyntheticMonitor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Model3rdPartySyntheticMonitor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Model3rdPartySyntheticMonitor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTestSetup

`func (o *Model3rdPartySyntheticMonitor) GetTestSetup() string`

GetTestSetup returns the TestSetup field if non-nil, zero value otherwise.

### GetTestSetupOk

`func (o *Model3rdPartySyntheticMonitor) GetTestSetupOk() (*string, bool)`

GetTestSetupOk returns a tuple with the TestSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSetup

`func (o *Model3rdPartySyntheticMonitor) SetTestSetup(v string)`

SetTestSetup sets TestSetup field to given value.

### HasTestSetup

`func (o *Model3rdPartySyntheticMonitor) HasTestSetup() bool`

HasTestSetup returns a boolean if a field has been set.

### GetExpirationTimestamp

`func (o *Model3rdPartySyntheticMonitor) GetExpirationTimestamp() int64`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *Model3rdPartySyntheticMonitor) GetExpirationTimestampOk() (*int64, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *Model3rdPartySyntheticMonitor) SetExpirationTimestamp(v int64)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.

### HasExpirationTimestamp

`func (o *Model3rdPartySyntheticMonitor) HasExpirationTimestamp() bool`

HasExpirationTimestamp returns a boolean if a field has been set.

### GetDrilldownLink

`func (o *Model3rdPartySyntheticMonitor) GetDrilldownLink() string`

GetDrilldownLink returns the DrilldownLink field if non-nil, zero value otherwise.

### GetDrilldownLinkOk

`func (o *Model3rdPartySyntheticMonitor) GetDrilldownLinkOk() (*string, bool)`

GetDrilldownLinkOk returns a tuple with the DrilldownLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrilldownLink

`func (o *Model3rdPartySyntheticMonitor) SetDrilldownLink(v string)`

SetDrilldownLink sets DrilldownLink field to given value.

### HasDrilldownLink

`func (o *Model3rdPartySyntheticMonitor) HasDrilldownLink() bool`

HasDrilldownLink returns a boolean if a field has been set.

### GetEditLink

`func (o *Model3rdPartySyntheticMonitor) GetEditLink() string`

GetEditLink returns the EditLink field if non-nil, zero value otherwise.

### GetEditLinkOk

`func (o *Model3rdPartySyntheticMonitor) GetEditLinkOk() (*string, bool)`

GetEditLinkOk returns a tuple with the EditLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditLink

`func (o *Model3rdPartySyntheticMonitor) SetEditLink(v string)`

SetEditLink sets EditLink field to given value.

### HasEditLink

`func (o *Model3rdPartySyntheticMonitor) HasEditLink() bool`

HasEditLink returns a boolean if a field has been set.

### GetEnabled

`func (o *Model3rdPartySyntheticMonitor) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Model3rdPartySyntheticMonitor) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Model3rdPartySyntheticMonitor) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Model3rdPartySyntheticMonitor) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDeleted

`func (o *Model3rdPartySyntheticMonitor) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Model3rdPartySyntheticMonitor) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Model3rdPartySyntheticMonitor) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Model3rdPartySyntheticMonitor) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetLocations

`func (o *Model3rdPartySyntheticMonitor) GetLocations() []SyntheticTestLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *Model3rdPartySyntheticMonitor) GetLocationsOk() (*[]SyntheticTestLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *Model3rdPartySyntheticMonitor) SetLocations(v []SyntheticTestLocation)`

SetLocations sets Locations field to given value.


### GetSteps

`func (o *Model3rdPartySyntheticMonitor) GetSteps() []SyntheticTestStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Model3rdPartySyntheticMonitor) GetStepsOk() (*[]SyntheticTestStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Model3rdPartySyntheticMonitor) SetSteps(v []SyntheticTestStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *Model3rdPartySyntheticMonitor) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetScheduleIntervalInSeconds

`func (o *Model3rdPartySyntheticMonitor) GetScheduleIntervalInSeconds() int32`

GetScheduleIntervalInSeconds returns the ScheduleIntervalInSeconds field if non-nil, zero value otherwise.

### GetScheduleIntervalInSecondsOk

`func (o *Model3rdPartySyntheticMonitor) GetScheduleIntervalInSecondsOk() (*int32, bool)`

GetScheduleIntervalInSecondsOk returns a tuple with the ScheduleIntervalInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleIntervalInSeconds

`func (o *Model3rdPartySyntheticMonitor) SetScheduleIntervalInSeconds(v int32)`

SetScheduleIntervalInSeconds sets ScheduleIntervalInSeconds field to given value.


### GetNoDataTimeout

`func (o *Model3rdPartySyntheticMonitor) GetNoDataTimeout() int32`

GetNoDataTimeout returns the NoDataTimeout field if non-nil, zero value otherwise.

### GetNoDataTimeoutOk

`func (o *Model3rdPartySyntheticMonitor) GetNoDataTimeoutOk() (*int32, bool)`

GetNoDataTimeoutOk returns a tuple with the NoDataTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataTimeout

`func (o *Model3rdPartySyntheticMonitor) SetNoDataTimeout(v int32)`

SetNoDataTimeout sets NoDataTimeout field to given value.

### HasNoDataTimeout

`func (o *Model3rdPartySyntheticMonitor) HasNoDataTimeout() bool`

HasNoDataTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


