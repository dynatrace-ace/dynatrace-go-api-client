/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace Environment API v1. To read about use cases and examples, refer to the [help page](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// NodeCollectionElement The short representation of a synthetic object. Only contains the ID and the display name.
type NodeCollectionElement struct {
	// The ID of a node.
	EntityId string `json:"entityId"`
	// The hostname of a node.
	Hostname string `json:"hostname"`
	// The IP of a node.
	Ips []string `json:"ips"`
	// The version of a node
	Version string `json:"version"`
	// Browser check capabilities enabled flag.
	BrowserMonitorsEnabled bool `json:"browserMonitorsEnabled"`
	// The version of the Active Gate.
	ActiveGateVersion string `json:"activeGateVersion"`
	// The Active Gate has the One Agent routing enabled ('true') or not ('false').
	OneAgentRoutingEnabled bool `json:"oneAgentRoutingEnabled "`
	// The Active Gate's host operating system.
	OperatingSystem string `json:"operatingSystem"`
	// The Active Gate has the Auto update option enabled ('true') or not ('false')
	AutoUpdateEnabled bool `json:"autoUpdateEnabled"`
	// The status of the synthetic node.
	Status string `json:"status"`
	// The version of the synthetic player.
	PlayerVersion string `json:"playerVersion"`
	// The health check status of the synthetic node.
	HealthCheckStatus string `json:"healthCheckStatus"`
}

// NewNodeCollectionElement instantiates a new NodeCollectionElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeCollectionElement(entityId string, hostname string, ips []string, version string, browserMonitorsEnabled bool, activeGateVersion string, oneAgentRoutingEnabled bool, operatingSystem string, autoUpdateEnabled bool, status string, playerVersion string, healthCheckStatus string) *NodeCollectionElement {
	this := NodeCollectionElement{}
	this.EntityId = entityId
	this.Hostname = hostname
	this.Ips = ips
	this.Version = version
	this.BrowserMonitorsEnabled = browserMonitorsEnabled
	this.ActiveGateVersion = activeGateVersion
	this.OneAgentRoutingEnabled = oneAgentRoutingEnabled
	this.OperatingSystem = operatingSystem
	this.AutoUpdateEnabled = autoUpdateEnabled
	this.Status = status
	this.PlayerVersion = playerVersion
	this.HealthCheckStatus = healthCheckStatus
	return &this
}

// NewNodeCollectionElementWithDefaults instantiates a new NodeCollectionElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeCollectionElementWithDefaults() *NodeCollectionElement {
	this := NodeCollectionElement{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *NodeCollectionElement) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *NodeCollectionElement) GetEntityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *NodeCollectionElement) SetEntityId(v string) {
	o.EntityId = v
}

// GetHostname returns the Hostname field value
func (o *NodeCollectionElement) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *NodeCollectionElement) GetHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *NodeCollectionElement) SetHostname(v string) {
	o.Hostname = v
}

// GetIps returns the Ips field value
func (o *NodeCollectionElement) GetIps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value
// and a boolean to check if the value has been set.
func (o *NodeCollectionElement) GetIpsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ips, true
}

// SetIps sets field value
func (o *NodeCollectionElement) SetIps(v []string) {
	o.Ips = v
}

// GetVersion returns the Version field value
func (o *NodeCollectionElement) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *NodeCollectionElement) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *NodeCollectionElement) SetVersion(v string) {
	o.Version = v
}

// GetBrowserMonitorsEnabled returns the BrowserMonitorsEnabled field value
func (o *NodeCollectionElement) GetBrowserMonitorsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BrowserMonitorsEnabled
}

// GetBrowserMonitorsEnabledOk returns a tuple with the BrowserMonitorsEnabled field value
// and a boolean to check if the value has been set.
func (o *NodeCollectionElement) GetBrowserMonitorsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BrowserMonitorsEnabled, true
}

// SetBrowserMonitorsEnabled sets field value
func (o *NodeCollectionElement) SetBrowserMonitorsEnabled(v bool) {
	o.BrowserMonitorsEnabled = v
}

// GetActiveGateVersion returns the ActiveGateVersion field value
func (o *NodeCollectionElement) GetActiveGateVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActiveGateVersion
}

// GetActiveGateVersionOk returns a tuple with the ActiveGateVersion field value
// and a boolean to check if the value has been set.
func (o *NodeCollectionElement) GetActiveGateVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ActiveGateVersion, true
}

// SetActiveGateVersion sets field value
func (o *NodeCollectionElement) SetActiveGateVersion(v string) {
	o.ActiveGateVersion = v
}

// GetOneAgentRoutingEnabled returns the OneAgentRoutingEnabled field value
func (o *NodeCollectionElement) GetOneAgentRoutingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OneAgentRoutingEnabled
}

// GetOneAgentRoutingEnabledOk returns a tuple with the OneAgentRoutingEnabled field value
// and a boolean to check if the value has been set.
func (o *NodeCollectionElement) GetOneAgentRoutingEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OneAgentRoutingEnabled, true
}

// SetOneAgentRoutingEnabled sets field value
func (o *NodeCollectionElement) SetOneAgentRoutingEnabled(v bool) {
	o.OneAgentRoutingEnabled = v
}

// GetOperatingSystem returns the OperatingSystem field value
func (o *NodeCollectionElement) GetOperatingSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value
// and a boolean to check if the value has been set.
func (o *NodeCollectionElement) GetOperatingSystemOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OperatingSystem, true
}

// SetOperatingSystem sets field value
func (o *NodeCollectionElement) SetOperatingSystem(v string) {
	o.OperatingSystem = v
}

// GetAutoUpdateEnabled returns the AutoUpdateEnabled field value
func (o *NodeCollectionElement) GetAutoUpdateEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoUpdateEnabled
}

// GetAutoUpdateEnabledOk returns a tuple with the AutoUpdateEnabled field value
// and a boolean to check if the value has been set.
func (o *NodeCollectionElement) GetAutoUpdateEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AutoUpdateEnabled, true
}

// SetAutoUpdateEnabled sets field value
func (o *NodeCollectionElement) SetAutoUpdateEnabled(v bool) {
	o.AutoUpdateEnabled = v
}

// GetStatus returns the Status field value
func (o *NodeCollectionElement) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *NodeCollectionElement) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *NodeCollectionElement) SetStatus(v string) {
	o.Status = v
}

// GetPlayerVersion returns the PlayerVersion field value
func (o *NodeCollectionElement) GetPlayerVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlayerVersion
}

// GetPlayerVersionOk returns a tuple with the PlayerVersion field value
// and a boolean to check if the value has been set.
func (o *NodeCollectionElement) GetPlayerVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlayerVersion, true
}

// SetPlayerVersion sets field value
func (o *NodeCollectionElement) SetPlayerVersion(v string) {
	o.PlayerVersion = v
}

// GetHealthCheckStatus returns the HealthCheckStatus field value
func (o *NodeCollectionElement) GetHealthCheckStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HealthCheckStatus
}

// GetHealthCheckStatusOk returns a tuple with the HealthCheckStatus field value
// and a boolean to check if the value has been set.
func (o *NodeCollectionElement) GetHealthCheckStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HealthCheckStatus, true
}

// SetHealthCheckStatus sets field value
func (o *NodeCollectionElement) SetHealthCheckStatus(v string) {
	o.HealthCheckStatus = v
}

func (o NodeCollectionElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entityId"] = o.EntityId
	}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	if true {
		toSerialize["ips"] = o.Ips
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["browserMonitorsEnabled"] = o.BrowserMonitorsEnabled
	}
	if true {
		toSerialize["activeGateVersion"] = o.ActiveGateVersion
	}
	if true {
		toSerialize["oneAgentRoutingEnabled "] = o.OneAgentRoutingEnabled
	}
	if true {
		toSerialize["operatingSystem"] = o.OperatingSystem
	}
	if true {
		toSerialize["autoUpdateEnabled"] = o.AutoUpdateEnabled
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["playerVersion"] = o.PlayerVersion
	}
	if true {
		toSerialize["healthCheckStatus"] = o.HealthCheckStatus
	}
	return json.Marshal(toSerialize)
}

type NullableNodeCollectionElement struct {
	value *NodeCollectionElement
	isSet bool
}

func (v NullableNodeCollectionElement) Get() *NodeCollectionElement {
	return v.value
}

func (v *NullableNodeCollectionElement) Set(val *NodeCollectionElement) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeCollectionElement) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeCollectionElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeCollectionElement(val *NodeCollectionElement) *NullableNodeCollectionElement {
	return &NullableNodeCollectionElement{value: val, isSet: true}
}

func (v NullableNodeCollectionElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeCollectionElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


