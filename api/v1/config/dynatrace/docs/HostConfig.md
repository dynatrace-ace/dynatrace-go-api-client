# HostConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Dynatrace entity ID of the host where OneAgent is deployed. | [optional] [readonly] 
**MonitoringConfig** | Pointer to [**MonitoringConfig**](MonitoringConfig.md) |  | [optional] 
**AutoUpdateConfig** | Pointer to [**HostAutoUpdateConfig**](HostAutoUpdateConfig.md) |  | [optional] 
**TechMonitoringConfigList** | Pointer to [**TechMonitoringList**](TechMonitoringList.md) |  | [optional] 

## Methods

### NewHostConfig

`func NewHostConfig() *HostConfig`

NewHostConfig instantiates a new HostConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostConfigWithDefaults

`func NewHostConfigWithDefaults() *HostConfig`

NewHostConfigWithDefaults instantiates a new HostConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HostConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HostConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMonitoringConfig

`func (o *HostConfig) GetMonitoringConfig() MonitoringConfig`

GetMonitoringConfig returns the MonitoringConfig field if non-nil, zero value otherwise.

### GetMonitoringConfigOk

`func (o *HostConfig) GetMonitoringConfigOk() (*MonitoringConfig, bool)`

GetMonitoringConfigOk returns a tuple with the MonitoringConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringConfig

`func (o *HostConfig) SetMonitoringConfig(v MonitoringConfig)`

SetMonitoringConfig sets MonitoringConfig field to given value.

### HasMonitoringConfig

`func (o *HostConfig) HasMonitoringConfig() bool`

HasMonitoringConfig returns a boolean if a field has been set.

### GetAutoUpdateConfig

`func (o *HostConfig) GetAutoUpdateConfig() HostAutoUpdateConfig`

GetAutoUpdateConfig returns the AutoUpdateConfig field if non-nil, zero value otherwise.

### GetAutoUpdateConfigOk

`func (o *HostConfig) GetAutoUpdateConfigOk() (*HostAutoUpdateConfig, bool)`

GetAutoUpdateConfigOk returns a tuple with the AutoUpdateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdateConfig

`func (o *HostConfig) SetAutoUpdateConfig(v HostAutoUpdateConfig)`

SetAutoUpdateConfig sets AutoUpdateConfig field to given value.

### HasAutoUpdateConfig

`func (o *HostConfig) HasAutoUpdateConfig() bool`

HasAutoUpdateConfig returns a boolean if a field has been set.

### GetTechMonitoringConfigList

`func (o *HostConfig) GetTechMonitoringConfigList() TechMonitoringList`

GetTechMonitoringConfigList returns the TechMonitoringConfigList field if non-nil, zero value otherwise.

### GetTechMonitoringConfigListOk

`func (o *HostConfig) GetTechMonitoringConfigListOk() (*TechMonitoringList, bool)`

GetTechMonitoringConfigListOk returns a tuple with the TechMonitoringConfigList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechMonitoringConfigList

`func (o *HostConfig) SetTechMonitoringConfigList(v TechMonitoringList)`

SetTechMonitoringConfigList sets TechMonitoringConfigList field to given value.

### HasTechMonitoringConfigList

`func (o *HostConfig) HasTechMonitoringConfigList() bool`

HasTechMonitoringConfigList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


