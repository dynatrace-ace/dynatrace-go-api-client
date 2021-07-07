# NodeCollectionElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The ID of a node. | 
**Hostname** | **string** | The hostname of a node. | 
**Ips** | **[]string** | The IP of a node. | 
**Version** | **string** | The version of a node | 
**BrowserMonitorsEnabled** | **bool** | Browser check capabilities enabled flag. | 
**ActiveGateVersion** | **string** | The version of the Active Gate. | 
**OneAgentRoutingEnabled** | **bool** | The Active Gate has the One Agent routing enabled (&#39;true&#39;) or not (&#39;false&#39;). | 
**OperatingSystem** | **string** | The Active Gate&#39;s host operating system. | 
**AutoUpdateEnabled** | **bool** | The Active Gate has the Auto update option enabled (&#39;true&#39;) or not (&#39;false&#39;) | 
**Status** | **string** | The status of the synthetic node. | 
**PlayerVersion** | **string** | The version of the synthetic player. | 
**HealthCheckStatus** | **string** | The health check status of the synthetic node. | 

## Methods

### NewNodeCollectionElement

`func NewNodeCollectionElement(entityId string, hostname string, ips []string, version string, browserMonitorsEnabled bool, activeGateVersion string, oneAgentRoutingEnabled bool, operatingSystem string, autoUpdateEnabled bool, status string, playerVersion string, healthCheckStatus string, ) *NodeCollectionElement`

NewNodeCollectionElement instantiates a new NodeCollectionElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeCollectionElementWithDefaults

`func NewNodeCollectionElementWithDefaults() *NodeCollectionElement`

NewNodeCollectionElementWithDefaults instantiates a new NodeCollectionElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *NodeCollectionElement) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *NodeCollectionElement) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *NodeCollectionElement) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetHostname

`func (o *NodeCollectionElement) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NodeCollectionElement) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NodeCollectionElement) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetIps

`func (o *NodeCollectionElement) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *NodeCollectionElement) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *NodeCollectionElement) SetIps(v []string)`

SetIps sets Ips field to given value.


### GetVersion

`func (o *NodeCollectionElement) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NodeCollectionElement) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NodeCollectionElement) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetBrowserMonitorsEnabled

`func (o *NodeCollectionElement) GetBrowserMonitorsEnabled() bool`

GetBrowserMonitorsEnabled returns the BrowserMonitorsEnabled field if non-nil, zero value otherwise.

### GetBrowserMonitorsEnabledOk

`func (o *NodeCollectionElement) GetBrowserMonitorsEnabledOk() (*bool, bool)`

GetBrowserMonitorsEnabledOk returns a tuple with the BrowserMonitorsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserMonitorsEnabled

`func (o *NodeCollectionElement) SetBrowserMonitorsEnabled(v bool)`

SetBrowserMonitorsEnabled sets BrowserMonitorsEnabled field to given value.


### GetActiveGateVersion

`func (o *NodeCollectionElement) GetActiveGateVersion() string`

GetActiveGateVersion returns the ActiveGateVersion field if non-nil, zero value otherwise.

### GetActiveGateVersionOk

`func (o *NodeCollectionElement) GetActiveGateVersionOk() (*string, bool)`

GetActiveGateVersionOk returns a tuple with the ActiveGateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGateVersion

`func (o *NodeCollectionElement) SetActiveGateVersion(v string)`

SetActiveGateVersion sets ActiveGateVersion field to given value.


### GetOneAgentRoutingEnabled

`func (o *NodeCollectionElement) GetOneAgentRoutingEnabled() bool`

GetOneAgentRoutingEnabled returns the OneAgentRoutingEnabled field if non-nil, zero value otherwise.

### GetOneAgentRoutingEnabledOk

`func (o *NodeCollectionElement) GetOneAgentRoutingEnabledOk() (*bool, bool)`

GetOneAgentRoutingEnabledOk returns a tuple with the OneAgentRoutingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneAgentRoutingEnabled

`func (o *NodeCollectionElement) SetOneAgentRoutingEnabled(v bool)`

SetOneAgentRoutingEnabled sets OneAgentRoutingEnabled field to given value.


### GetOperatingSystem

`func (o *NodeCollectionElement) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *NodeCollectionElement) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *NodeCollectionElement) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### GetAutoUpdateEnabled

`func (o *NodeCollectionElement) GetAutoUpdateEnabled() bool`

GetAutoUpdateEnabled returns the AutoUpdateEnabled field if non-nil, zero value otherwise.

### GetAutoUpdateEnabledOk

`func (o *NodeCollectionElement) GetAutoUpdateEnabledOk() (*bool, bool)`

GetAutoUpdateEnabledOk returns a tuple with the AutoUpdateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdateEnabled

`func (o *NodeCollectionElement) SetAutoUpdateEnabled(v bool)`

SetAutoUpdateEnabled sets AutoUpdateEnabled field to given value.


### GetStatus

`func (o *NodeCollectionElement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeCollectionElement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeCollectionElement) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlayerVersion

`func (o *NodeCollectionElement) GetPlayerVersion() string`

GetPlayerVersion returns the PlayerVersion field if non-nil, zero value otherwise.

### GetPlayerVersionOk

`func (o *NodeCollectionElement) GetPlayerVersionOk() (*string, bool)`

GetPlayerVersionOk returns a tuple with the PlayerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerVersion

`func (o *NodeCollectionElement) SetPlayerVersion(v string)`

SetPlayerVersion sets PlayerVersion field to given value.


### GetHealthCheckStatus

`func (o *NodeCollectionElement) GetHealthCheckStatus() string`

GetHealthCheckStatus returns the HealthCheckStatus field if non-nil, zero value otherwise.

### GetHealthCheckStatusOk

`func (o *NodeCollectionElement) GetHealthCheckStatusOk() (*string, bool)`

GetHealthCheckStatusOk returns a tuple with the HealthCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckStatus

`func (o *NodeCollectionElement) SetHealthCheckStatus(v string)`

SetHealthCheckStatus sets HealthCheckStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


