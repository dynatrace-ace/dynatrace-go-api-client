# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The ID of the synthetic node. | 
**Hostname** | **string** | The hostname of the synthetic node. | 
**Ips** | **[]string** | The IP of the synthetic node. | 
**Version** | **string** | The version of the synthetic node. | 
**BrowserMonitorsEnabled** | **bool** | The synthetic node is able to execute browser monitors (&#x60;true&#x60;) or not (&#x60;false&#x60;). | 
**ActiveGateVersion** | **string** | The version of the Active Gate. | 
**OneAgentRoutingEnabled** | **bool** | The Active Gate has the One Agent routing enabled (&#39;true&#39;) or not (&#39;false&#39;). | 
**OperatingSystem** | **string** | The Active Gate&#39;s host operating system. | 
**AutoUpdateEnabled** | **bool** | The Active Gate has the Auto update option enabled (&#39;true&#39;) or not (&#39;false&#39;) | 
**Status** | **string** | The status of the synthetic node. | 
**PlayerVersion** | **string** | The version of the synthetic player. | 
**HealthCheckStatus** | **string** | The health check status of the synthetic node. | 
**BrowserType** | **string** | The browser type. | 
**BrowserVersion** | **string** | The browser version. | 

## Methods

### NewNode

`func NewNode(entityId string, hostname string, ips []string, version string, browserMonitorsEnabled bool, activeGateVersion string, oneAgentRoutingEnabled bool, operatingSystem string, autoUpdateEnabled bool, status string, playerVersion string, healthCheckStatus string, browserType string, browserVersion string, ) *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *Node) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Node) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Node) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetHostname

`func (o *Node) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Node) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Node) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetIps

`func (o *Node) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *Node) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *Node) SetIps(v []string)`

SetIps sets Ips field to given value.


### GetVersion

`func (o *Node) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Node) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Node) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetBrowserMonitorsEnabled

`func (o *Node) GetBrowserMonitorsEnabled() bool`

GetBrowserMonitorsEnabled returns the BrowserMonitorsEnabled field if non-nil, zero value otherwise.

### GetBrowserMonitorsEnabledOk

`func (o *Node) GetBrowserMonitorsEnabledOk() (*bool, bool)`

GetBrowserMonitorsEnabledOk returns a tuple with the BrowserMonitorsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserMonitorsEnabled

`func (o *Node) SetBrowserMonitorsEnabled(v bool)`

SetBrowserMonitorsEnabled sets BrowserMonitorsEnabled field to given value.


### GetActiveGateVersion

`func (o *Node) GetActiveGateVersion() string`

GetActiveGateVersion returns the ActiveGateVersion field if non-nil, zero value otherwise.

### GetActiveGateVersionOk

`func (o *Node) GetActiveGateVersionOk() (*string, bool)`

GetActiveGateVersionOk returns a tuple with the ActiveGateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGateVersion

`func (o *Node) SetActiveGateVersion(v string)`

SetActiveGateVersion sets ActiveGateVersion field to given value.


### GetOneAgentRoutingEnabled

`func (o *Node) GetOneAgentRoutingEnabled() bool`

GetOneAgentRoutingEnabled returns the OneAgentRoutingEnabled field if non-nil, zero value otherwise.

### GetOneAgentRoutingEnabledOk

`func (o *Node) GetOneAgentRoutingEnabledOk() (*bool, bool)`

GetOneAgentRoutingEnabledOk returns a tuple with the OneAgentRoutingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneAgentRoutingEnabled

`func (o *Node) SetOneAgentRoutingEnabled(v bool)`

SetOneAgentRoutingEnabled sets OneAgentRoutingEnabled field to given value.


### GetOperatingSystem

`func (o *Node) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *Node) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *Node) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### GetAutoUpdateEnabled

`func (o *Node) GetAutoUpdateEnabled() bool`

GetAutoUpdateEnabled returns the AutoUpdateEnabled field if non-nil, zero value otherwise.

### GetAutoUpdateEnabledOk

`func (o *Node) GetAutoUpdateEnabledOk() (*bool, bool)`

GetAutoUpdateEnabledOk returns a tuple with the AutoUpdateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdateEnabled

`func (o *Node) SetAutoUpdateEnabled(v bool)`

SetAutoUpdateEnabled sets AutoUpdateEnabled field to given value.


### GetStatus

`func (o *Node) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Node) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Node) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlayerVersion

`func (o *Node) GetPlayerVersion() string`

GetPlayerVersion returns the PlayerVersion field if non-nil, zero value otherwise.

### GetPlayerVersionOk

`func (o *Node) GetPlayerVersionOk() (*string, bool)`

GetPlayerVersionOk returns a tuple with the PlayerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerVersion

`func (o *Node) SetPlayerVersion(v string)`

SetPlayerVersion sets PlayerVersion field to given value.


### GetHealthCheckStatus

`func (o *Node) GetHealthCheckStatus() string`

GetHealthCheckStatus returns the HealthCheckStatus field if non-nil, zero value otherwise.

### GetHealthCheckStatusOk

`func (o *Node) GetHealthCheckStatusOk() (*string, bool)`

GetHealthCheckStatusOk returns a tuple with the HealthCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckStatus

`func (o *Node) SetHealthCheckStatus(v string)`

SetHealthCheckStatus sets HealthCheckStatus field to given value.


### GetBrowserType

`func (o *Node) GetBrowserType() string`

GetBrowserType returns the BrowserType field if non-nil, zero value otherwise.

### GetBrowserTypeOk

`func (o *Node) GetBrowserTypeOk() (*string, bool)`

GetBrowserTypeOk returns a tuple with the BrowserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserType

`func (o *Node) SetBrowserType(v string)`

SetBrowserType sets BrowserType field to given value.


### GetBrowserVersion

`func (o *Node) GetBrowserVersion() string`

GetBrowserVersion returns the BrowserVersion field if non-nil, zero value otherwise.

### GetBrowserVersionOk

`func (o *Node) GetBrowserVersionOk() (*string, bool)`

GetBrowserVersionOk returns a tuple with the BrowserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserVersion

`func (o *Node) SetBrowserVersion(v string)`

SetBrowserVersion sets BrowserVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


