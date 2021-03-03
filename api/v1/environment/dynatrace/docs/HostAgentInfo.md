# HostAgentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostInfo** | Pointer to [**Host**](Host.md) |  | [optional] 
**FaultyVersion** | Pointer to **bool** | OneAgent version is faulty (&#x60;true&#x60;) or not (&#x60;false&#x60;). | [optional] 
**Active** | Pointer to **bool** | OneAgent is active (&#x60;true&#x60;) or inactive (&#x60;false&#x60;). | [optional] 
**ConfiguredMonitoringMode** | Pointer to **string** | Configured monitoring mode of OneAgent. | [optional] 
**MonitoringType** | Pointer to **string** | The monitoring mode of OneAgent. | [optional] 
**AutoUpdateSetting** | Pointer to **string** | The effective auto-update setting of OneAgent. For host with inherited configuration it is calculated from its parent&#39;s configuration | [optional] 
**UpdateStatus** | Pointer to **string** | The current update status of OneAgent. | [optional] 
**AvailableVersions** | Pointer to **[]string** | A list of versions OneAgent can be updated to. | [optional] 
**ConfiguredMonitoringEnabled** | Pointer to **bool** | Monitoring is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;) in the OneAgent configuration. | [optional] 
**AvailabilityState** | Pointer to **string** | The availability state of OneAgent. | [optional] 
**CurrentActiveGateId** | Pointer to **int32** | The ID of an ActiveGate to which OneAgent is connected. | [optional] 
**CurrentNetworkZoneId** | Pointer to **string** | The ID of the network zone that OneAgent is using. | [optional] 
**Modules** | Pointer to [**[]ModuleInfo**](ModuleInfo.md) | A list of code modules deployed on the host. | [optional] 
**Plugins** | Pointer to [**[]PluginInfo**](PluginInfo.md) | A list of plugins deployed on the host. | [optional] 

## Methods

### NewHostAgentInfo

`func NewHostAgentInfo() *HostAgentInfo`

NewHostAgentInfo instantiates a new HostAgentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostAgentInfoWithDefaults

`func NewHostAgentInfoWithDefaults() *HostAgentInfo`

NewHostAgentInfoWithDefaults instantiates a new HostAgentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostInfo

`func (o *HostAgentInfo) GetHostInfo() Host`

GetHostInfo returns the HostInfo field if non-nil, zero value otherwise.

### GetHostInfoOk

`func (o *HostAgentInfo) GetHostInfoOk() (*Host, bool)`

GetHostInfoOk returns a tuple with the HostInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostInfo

`func (o *HostAgentInfo) SetHostInfo(v Host)`

SetHostInfo sets HostInfo field to given value.

### HasHostInfo

`func (o *HostAgentInfo) HasHostInfo() bool`

HasHostInfo returns a boolean if a field has been set.

### GetFaultyVersion

`func (o *HostAgentInfo) GetFaultyVersion() bool`

GetFaultyVersion returns the FaultyVersion field if non-nil, zero value otherwise.

### GetFaultyVersionOk

`func (o *HostAgentInfo) GetFaultyVersionOk() (*bool, bool)`

GetFaultyVersionOk returns a tuple with the FaultyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultyVersion

`func (o *HostAgentInfo) SetFaultyVersion(v bool)`

SetFaultyVersion sets FaultyVersion field to given value.

### HasFaultyVersion

`func (o *HostAgentInfo) HasFaultyVersion() bool`

HasFaultyVersion returns a boolean if a field has been set.

### GetActive

`func (o *HostAgentInfo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *HostAgentInfo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *HostAgentInfo) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *HostAgentInfo) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConfiguredMonitoringMode

`func (o *HostAgentInfo) GetConfiguredMonitoringMode() string`

GetConfiguredMonitoringMode returns the ConfiguredMonitoringMode field if non-nil, zero value otherwise.

### GetConfiguredMonitoringModeOk

`func (o *HostAgentInfo) GetConfiguredMonitoringModeOk() (*string, bool)`

GetConfiguredMonitoringModeOk returns a tuple with the ConfiguredMonitoringMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredMonitoringMode

`func (o *HostAgentInfo) SetConfiguredMonitoringMode(v string)`

SetConfiguredMonitoringMode sets ConfiguredMonitoringMode field to given value.

### HasConfiguredMonitoringMode

`func (o *HostAgentInfo) HasConfiguredMonitoringMode() bool`

HasConfiguredMonitoringMode returns a boolean if a field has been set.

### GetMonitoringType

`func (o *HostAgentInfo) GetMonitoringType() string`

GetMonitoringType returns the MonitoringType field if non-nil, zero value otherwise.

### GetMonitoringTypeOk

`func (o *HostAgentInfo) GetMonitoringTypeOk() (*string, bool)`

GetMonitoringTypeOk returns a tuple with the MonitoringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringType

`func (o *HostAgentInfo) SetMonitoringType(v string)`

SetMonitoringType sets MonitoringType field to given value.

### HasMonitoringType

`func (o *HostAgentInfo) HasMonitoringType() bool`

HasMonitoringType returns a boolean if a field has been set.

### GetAutoUpdateSetting

`func (o *HostAgentInfo) GetAutoUpdateSetting() string`

GetAutoUpdateSetting returns the AutoUpdateSetting field if non-nil, zero value otherwise.

### GetAutoUpdateSettingOk

`func (o *HostAgentInfo) GetAutoUpdateSettingOk() (*string, bool)`

GetAutoUpdateSettingOk returns a tuple with the AutoUpdateSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdateSetting

`func (o *HostAgentInfo) SetAutoUpdateSetting(v string)`

SetAutoUpdateSetting sets AutoUpdateSetting field to given value.

### HasAutoUpdateSetting

`func (o *HostAgentInfo) HasAutoUpdateSetting() bool`

HasAutoUpdateSetting returns a boolean if a field has been set.

### GetUpdateStatus

`func (o *HostAgentInfo) GetUpdateStatus() string`

GetUpdateStatus returns the UpdateStatus field if non-nil, zero value otherwise.

### GetUpdateStatusOk

`func (o *HostAgentInfo) GetUpdateStatusOk() (*string, bool)`

GetUpdateStatusOk returns a tuple with the UpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStatus

`func (o *HostAgentInfo) SetUpdateStatus(v string)`

SetUpdateStatus sets UpdateStatus field to given value.

### HasUpdateStatus

`func (o *HostAgentInfo) HasUpdateStatus() bool`

HasUpdateStatus returns a boolean if a field has been set.

### GetAvailableVersions

`func (o *HostAgentInfo) GetAvailableVersions() []string`

GetAvailableVersions returns the AvailableVersions field if non-nil, zero value otherwise.

### GetAvailableVersionsOk

`func (o *HostAgentInfo) GetAvailableVersionsOk() (*[]string, bool)`

GetAvailableVersionsOk returns a tuple with the AvailableVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVersions

`func (o *HostAgentInfo) SetAvailableVersions(v []string)`

SetAvailableVersions sets AvailableVersions field to given value.

### HasAvailableVersions

`func (o *HostAgentInfo) HasAvailableVersions() bool`

HasAvailableVersions returns a boolean if a field has been set.

### GetConfiguredMonitoringEnabled

`func (o *HostAgentInfo) GetConfiguredMonitoringEnabled() bool`

GetConfiguredMonitoringEnabled returns the ConfiguredMonitoringEnabled field if non-nil, zero value otherwise.

### GetConfiguredMonitoringEnabledOk

`func (o *HostAgentInfo) GetConfiguredMonitoringEnabledOk() (*bool, bool)`

GetConfiguredMonitoringEnabledOk returns a tuple with the ConfiguredMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredMonitoringEnabled

`func (o *HostAgentInfo) SetConfiguredMonitoringEnabled(v bool)`

SetConfiguredMonitoringEnabled sets ConfiguredMonitoringEnabled field to given value.

### HasConfiguredMonitoringEnabled

`func (o *HostAgentInfo) HasConfiguredMonitoringEnabled() bool`

HasConfiguredMonitoringEnabled returns a boolean if a field has been set.

### GetAvailabilityState

`func (o *HostAgentInfo) GetAvailabilityState() string`

GetAvailabilityState returns the AvailabilityState field if non-nil, zero value otherwise.

### GetAvailabilityStateOk

`func (o *HostAgentInfo) GetAvailabilityStateOk() (*string, bool)`

GetAvailabilityStateOk returns a tuple with the AvailabilityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityState

`func (o *HostAgentInfo) SetAvailabilityState(v string)`

SetAvailabilityState sets AvailabilityState field to given value.

### HasAvailabilityState

`func (o *HostAgentInfo) HasAvailabilityState() bool`

HasAvailabilityState returns a boolean if a field has been set.

### GetCurrentActiveGateId

`func (o *HostAgentInfo) GetCurrentActiveGateId() int32`

GetCurrentActiveGateId returns the CurrentActiveGateId field if non-nil, zero value otherwise.

### GetCurrentActiveGateIdOk

`func (o *HostAgentInfo) GetCurrentActiveGateIdOk() (*int32, bool)`

GetCurrentActiveGateIdOk returns a tuple with the CurrentActiveGateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentActiveGateId

`func (o *HostAgentInfo) SetCurrentActiveGateId(v int32)`

SetCurrentActiveGateId sets CurrentActiveGateId field to given value.

### HasCurrentActiveGateId

`func (o *HostAgentInfo) HasCurrentActiveGateId() bool`

HasCurrentActiveGateId returns a boolean if a field has been set.

### GetCurrentNetworkZoneId

`func (o *HostAgentInfo) GetCurrentNetworkZoneId() string`

GetCurrentNetworkZoneId returns the CurrentNetworkZoneId field if non-nil, zero value otherwise.

### GetCurrentNetworkZoneIdOk

`func (o *HostAgentInfo) GetCurrentNetworkZoneIdOk() (*string, bool)`

GetCurrentNetworkZoneIdOk returns a tuple with the CurrentNetworkZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNetworkZoneId

`func (o *HostAgentInfo) SetCurrentNetworkZoneId(v string)`

SetCurrentNetworkZoneId sets CurrentNetworkZoneId field to given value.

### HasCurrentNetworkZoneId

`func (o *HostAgentInfo) HasCurrentNetworkZoneId() bool`

HasCurrentNetworkZoneId returns a boolean if a field has been set.

### GetModules

`func (o *HostAgentInfo) GetModules() []ModuleInfo`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *HostAgentInfo) GetModulesOk() (*[]ModuleInfo, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *HostAgentInfo) SetModules(v []ModuleInfo)`

SetModules sets Modules field to given value.

### HasModules

`func (o *HostAgentInfo) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetPlugins

`func (o *HostAgentInfo) GetPlugins() []PluginInfo`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *HostAgentInfo) GetPluginsOk() (*[]PluginInfo, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *HostAgentInfo) SetPlugins(v []PluginInfo)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *HostAgentInfo) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


