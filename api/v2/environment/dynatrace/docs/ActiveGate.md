# ActiveGate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the ActiveGate. | [optional] [readonly] 
**NetworkAddresses** | Pointer to **[]string** | A list of network addresses of the ActiveGate. | [optional] [readonly] 
**LoadBalancerAddresses** | Pointer to **[]string** | A list of Load Balancer addresses of the ActiveGate. | [optional] [readonly] 
**OsType** | Pointer to **string** | The OS type that the ActiveGate is running on. | [optional] [readonly] 
**AutoUpdateStatus** | Pointer to **string** | The current status of auto-updates of the ActiveGate. | [optional] [readonly] 
**OfflineSince** | Pointer to **int64** | The timestamp since when the ActiveGate is offline.    The &#x60;null&#x60; value means the ActiveGate is online. | [optional] [readonly] 
**Version** | Pointer to **string** | The current version of the ActiveGate in the &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;.&lt;timestamp&gt;&#x60; format. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the ActiveGate. | [optional] [readonly] 
**Hostname** | Pointer to **string** | The name of the host the ActiveGate is running on. | [optional] [readonly] 
**MainEnvironment** | Pointer to **string** | The ID of the main environment for a multi-environment ActiveGate. | [optional] [readonly] 
**Environments** | Pointer to **[]string** | A list of environments (specified by IDs) the ActiveGate can connect to. | [optional] [readonly] 
**AutoUpdateSettings** | Pointer to [**ActiveGateAutoUpdateConfig**](ActiveGateAutoUpdateConfig.md) |  | [optional] 
**NetworkZone** | Pointer to **string** | The network zone of the ActiveGate. | [optional] [readonly] 
**Group** | Pointer to **string** | The group of the ActiveGate. | [optional] [readonly] 
**Modules** | Pointer to [**[]ActiveGateModule**](ActiveGateModule.md) | A list of modules of the ActiveGate. | [optional] [readonly] 
**Containerized** | Pointer to **bool** | ActiveGate is deployed in container (&#x60;true&#x60;) or not (&#x60;false&#x60;). | [optional] [readonly] 

## Methods

### NewActiveGate

`func NewActiveGate() *ActiveGate`

NewActiveGate instantiates a new ActiveGate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveGateWithDefaults

`func NewActiveGateWithDefaults() *ActiveGate`

NewActiveGateWithDefaults instantiates a new ActiveGate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActiveGate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActiveGate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActiveGate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActiveGate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkAddresses

`func (o *ActiveGate) GetNetworkAddresses() []string`

GetNetworkAddresses returns the NetworkAddresses field if non-nil, zero value otherwise.

### GetNetworkAddressesOk

`func (o *ActiveGate) GetNetworkAddressesOk() (*[]string, bool)`

GetNetworkAddressesOk returns a tuple with the NetworkAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddresses

`func (o *ActiveGate) SetNetworkAddresses(v []string)`

SetNetworkAddresses sets NetworkAddresses field to given value.

### HasNetworkAddresses

`func (o *ActiveGate) HasNetworkAddresses() bool`

HasNetworkAddresses returns a boolean if a field has been set.

### GetLoadBalancerAddresses

`func (o *ActiveGate) GetLoadBalancerAddresses() []string`

GetLoadBalancerAddresses returns the LoadBalancerAddresses field if non-nil, zero value otherwise.

### GetLoadBalancerAddressesOk

`func (o *ActiveGate) GetLoadBalancerAddressesOk() (*[]string, bool)`

GetLoadBalancerAddressesOk returns a tuple with the LoadBalancerAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerAddresses

`func (o *ActiveGate) SetLoadBalancerAddresses(v []string)`

SetLoadBalancerAddresses sets LoadBalancerAddresses field to given value.

### HasLoadBalancerAddresses

`func (o *ActiveGate) HasLoadBalancerAddresses() bool`

HasLoadBalancerAddresses returns a boolean if a field has been set.

### GetOsType

`func (o *ActiveGate) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ActiveGate) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ActiveGate) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ActiveGate) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetAutoUpdateStatus

`func (o *ActiveGate) GetAutoUpdateStatus() string`

GetAutoUpdateStatus returns the AutoUpdateStatus field if non-nil, zero value otherwise.

### GetAutoUpdateStatusOk

`func (o *ActiveGate) GetAutoUpdateStatusOk() (*string, bool)`

GetAutoUpdateStatusOk returns a tuple with the AutoUpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdateStatus

`func (o *ActiveGate) SetAutoUpdateStatus(v string)`

SetAutoUpdateStatus sets AutoUpdateStatus field to given value.

### HasAutoUpdateStatus

`func (o *ActiveGate) HasAutoUpdateStatus() bool`

HasAutoUpdateStatus returns a boolean if a field has been set.

### GetOfflineSince

`func (o *ActiveGate) GetOfflineSince() int64`

GetOfflineSince returns the OfflineSince field if non-nil, zero value otherwise.

### GetOfflineSinceOk

`func (o *ActiveGate) GetOfflineSinceOk() (*int64, bool)`

GetOfflineSinceOk returns a tuple with the OfflineSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineSince

`func (o *ActiveGate) SetOfflineSince(v int64)`

SetOfflineSince sets OfflineSince field to given value.

### HasOfflineSince

`func (o *ActiveGate) HasOfflineSince() bool`

HasOfflineSince returns a boolean if a field has been set.

### GetVersion

`func (o *ActiveGate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ActiveGate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ActiveGate) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ActiveGate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *ActiveGate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActiveGate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActiveGate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActiveGate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHostname

`func (o *ActiveGate) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ActiveGate) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ActiveGate) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ActiveGate) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetMainEnvironment

`func (o *ActiveGate) GetMainEnvironment() string`

GetMainEnvironment returns the MainEnvironment field if non-nil, zero value otherwise.

### GetMainEnvironmentOk

`func (o *ActiveGate) GetMainEnvironmentOk() (*string, bool)`

GetMainEnvironmentOk returns a tuple with the MainEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainEnvironment

`func (o *ActiveGate) SetMainEnvironment(v string)`

SetMainEnvironment sets MainEnvironment field to given value.

### HasMainEnvironment

`func (o *ActiveGate) HasMainEnvironment() bool`

HasMainEnvironment returns a boolean if a field has been set.

### GetEnvironments

`func (o *ActiveGate) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ActiveGate) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ActiveGate) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ActiveGate) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetAutoUpdateSettings

`func (o *ActiveGate) GetAutoUpdateSettings() ActiveGateAutoUpdateConfig`

GetAutoUpdateSettings returns the AutoUpdateSettings field if non-nil, zero value otherwise.

### GetAutoUpdateSettingsOk

`func (o *ActiveGate) GetAutoUpdateSettingsOk() (*ActiveGateAutoUpdateConfig, bool)`

GetAutoUpdateSettingsOk returns a tuple with the AutoUpdateSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdateSettings

`func (o *ActiveGate) SetAutoUpdateSettings(v ActiveGateAutoUpdateConfig)`

SetAutoUpdateSettings sets AutoUpdateSettings field to given value.

### HasAutoUpdateSettings

`func (o *ActiveGate) HasAutoUpdateSettings() bool`

HasAutoUpdateSettings returns a boolean if a field has been set.

### GetNetworkZone

`func (o *ActiveGate) GetNetworkZone() string`

GetNetworkZone returns the NetworkZone field if non-nil, zero value otherwise.

### GetNetworkZoneOk

`func (o *ActiveGate) GetNetworkZoneOk() (*string, bool)`

GetNetworkZoneOk returns a tuple with the NetworkZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkZone

`func (o *ActiveGate) SetNetworkZone(v string)`

SetNetworkZone sets NetworkZone field to given value.

### HasNetworkZone

`func (o *ActiveGate) HasNetworkZone() bool`

HasNetworkZone returns a boolean if a field has been set.

### GetGroup

`func (o *ActiveGate) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ActiveGate) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ActiveGate) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ActiveGate) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetModules

`func (o *ActiveGate) GetModules() []ActiveGateModule`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ActiveGate) GetModulesOk() (*[]ActiveGateModule, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ActiveGate) SetModules(v []ActiveGateModule)`

SetModules sets Modules field to given value.

### HasModules

`func (o *ActiveGate) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetContainerized

`func (o *ActiveGate) GetContainerized() bool`

GetContainerized returns the Containerized field if non-nil, zero value otherwise.

### GetContainerizedOk

`func (o *ActiveGate) GetContainerizedOk() (*bool, bool)`

GetContainerizedOk returns a tuple with the Containerized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerized

`func (o *ActiveGate) SetContainerized(v bool)`

SetContainerized sets Containerized field to given value.

### HasContainerized

`func (o *ActiveGate) HasContainerized() bool`

HasContainerized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


