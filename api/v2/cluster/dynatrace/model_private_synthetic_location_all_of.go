/*
 * Dynatrace Cluster API
 *
 * Dynatrace Managed exposes cluster-wide functionality via REST endpoints. This interactive documentation also acts as a REST client you can use to interact with Dynatrace Managed clusters.   To authorize, use the API Token generated in [Settings - API Tokens page](/cmc#cm/apiToken). The HTTP status code of the response shows the result of your request. The expected response code for a successful request is documented individually per REST endpoint. Additionally the following error response codes can occur in our REST interface:  * 400 - Bad Request: Some request parameters are not correct. See response body for details. * 401 - Unauthorized: A valid authorization header is required but is missing. * 403 - Forbidden: Execution of request is not allowed, e.g. api-token is invalid. * 404 - Not Found: Endpoint does not exist or some entities could not be found. * 500 - Internal Server Error: See response body for details. * 556 - Upgrade in progress: Operation couldn't be performed during the upgrade.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// PrivateSyntheticLocationAllOf struct for PrivateSyntheticLocationAllOf
type PrivateSyntheticLocationAllOf struct {
	// A list of synthetic nodes belonging to the location.    You can retrieve the list of available nodes with the [GET all nodes](https://dt-url.net/miy3rpl) call.
	Nodes *[]string `json:"nodes,omitempty"`
	// The alerting of location outage is enabled (`true`) or disabled (`false`).
	AvailabilityLocationOutage *bool `json:"availabilityLocationOutage,omitempty"`
	// The alerting of node outage is enabled (`true`) or disabled (`false`).    If enabled, the outage of *any* node in the location triggers an alert.
	AvailabilityNodeOutage *bool `json:"availabilityNodeOutage,omitempty"`
	// Alert if the location or node outage lasts longer than *X* minutes.    Only applicable when **availabilityLocationOutage** or **availabilityNodeOutage** is set to `true`.
	LocationNodeOutageDelayInMinutes *int32 `json:"locationNodeOutageDelayInMinutes,omitempty"`
	// The notifications of location and node outage is enabled (`true`) or disabled (`false`).
	AvailabilityNotificationsEnabled *bool `json:"availabilityNotificationsEnabled,omitempty"`
}

// NewPrivateSyntheticLocationAllOf instantiates a new PrivateSyntheticLocationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateSyntheticLocationAllOf() *PrivateSyntheticLocationAllOf {
	this := PrivateSyntheticLocationAllOf{}
	return &this
}

// NewPrivateSyntheticLocationAllOfWithDefaults instantiates a new PrivateSyntheticLocationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateSyntheticLocationAllOfWithDefaults() *PrivateSyntheticLocationAllOf {
	this := PrivateSyntheticLocationAllOf{}
	return &this
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *PrivateSyntheticLocationAllOf) GetNodes() []string {
	if o == nil || o.Nodes == nil {
		var ret []string
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSyntheticLocationAllOf) GetNodesOk() (*[]string, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *PrivateSyntheticLocationAllOf) HasNodes() bool {
	if o != nil && o.Nodes != nil {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []string and assigns it to the Nodes field.
func (o *PrivateSyntheticLocationAllOf) SetNodes(v []string) {
	o.Nodes = &v
}

// GetAvailabilityLocationOutage returns the AvailabilityLocationOutage field value if set, zero value otherwise.
func (o *PrivateSyntheticLocationAllOf) GetAvailabilityLocationOutage() bool {
	if o == nil || o.AvailabilityLocationOutage == nil {
		var ret bool
		return ret
	}
	return *o.AvailabilityLocationOutage
}

// GetAvailabilityLocationOutageOk returns a tuple with the AvailabilityLocationOutage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSyntheticLocationAllOf) GetAvailabilityLocationOutageOk() (*bool, bool) {
	if o == nil || o.AvailabilityLocationOutage == nil {
		return nil, false
	}
	return o.AvailabilityLocationOutage, true
}

// HasAvailabilityLocationOutage returns a boolean if a field has been set.
func (o *PrivateSyntheticLocationAllOf) HasAvailabilityLocationOutage() bool {
	if o != nil && o.AvailabilityLocationOutage != nil {
		return true
	}

	return false
}

// SetAvailabilityLocationOutage gets a reference to the given bool and assigns it to the AvailabilityLocationOutage field.
func (o *PrivateSyntheticLocationAllOf) SetAvailabilityLocationOutage(v bool) {
	o.AvailabilityLocationOutage = &v
}

// GetAvailabilityNodeOutage returns the AvailabilityNodeOutage field value if set, zero value otherwise.
func (o *PrivateSyntheticLocationAllOf) GetAvailabilityNodeOutage() bool {
	if o == nil || o.AvailabilityNodeOutage == nil {
		var ret bool
		return ret
	}
	return *o.AvailabilityNodeOutage
}

// GetAvailabilityNodeOutageOk returns a tuple with the AvailabilityNodeOutage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSyntheticLocationAllOf) GetAvailabilityNodeOutageOk() (*bool, bool) {
	if o == nil || o.AvailabilityNodeOutage == nil {
		return nil, false
	}
	return o.AvailabilityNodeOutage, true
}

// HasAvailabilityNodeOutage returns a boolean if a field has been set.
func (o *PrivateSyntheticLocationAllOf) HasAvailabilityNodeOutage() bool {
	if o != nil && o.AvailabilityNodeOutage != nil {
		return true
	}

	return false
}

// SetAvailabilityNodeOutage gets a reference to the given bool and assigns it to the AvailabilityNodeOutage field.
func (o *PrivateSyntheticLocationAllOf) SetAvailabilityNodeOutage(v bool) {
	o.AvailabilityNodeOutage = &v
}

// GetLocationNodeOutageDelayInMinutes returns the LocationNodeOutageDelayInMinutes field value if set, zero value otherwise.
func (o *PrivateSyntheticLocationAllOf) GetLocationNodeOutageDelayInMinutes() int32 {
	if o == nil || o.LocationNodeOutageDelayInMinutes == nil {
		var ret int32
		return ret
	}
	return *o.LocationNodeOutageDelayInMinutes
}

// GetLocationNodeOutageDelayInMinutesOk returns a tuple with the LocationNodeOutageDelayInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSyntheticLocationAllOf) GetLocationNodeOutageDelayInMinutesOk() (*int32, bool) {
	if o == nil || o.LocationNodeOutageDelayInMinutes == nil {
		return nil, false
	}
	return o.LocationNodeOutageDelayInMinutes, true
}

// HasLocationNodeOutageDelayInMinutes returns a boolean if a field has been set.
func (o *PrivateSyntheticLocationAllOf) HasLocationNodeOutageDelayInMinutes() bool {
	if o != nil && o.LocationNodeOutageDelayInMinutes != nil {
		return true
	}

	return false
}

// SetLocationNodeOutageDelayInMinutes gets a reference to the given int32 and assigns it to the LocationNodeOutageDelayInMinutes field.
func (o *PrivateSyntheticLocationAllOf) SetLocationNodeOutageDelayInMinutes(v int32) {
	o.LocationNodeOutageDelayInMinutes = &v
}

// GetAvailabilityNotificationsEnabled returns the AvailabilityNotificationsEnabled field value if set, zero value otherwise.
func (o *PrivateSyntheticLocationAllOf) GetAvailabilityNotificationsEnabled() bool {
	if o == nil || o.AvailabilityNotificationsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AvailabilityNotificationsEnabled
}

// GetAvailabilityNotificationsEnabledOk returns a tuple with the AvailabilityNotificationsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSyntheticLocationAllOf) GetAvailabilityNotificationsEnabledOk() (*bool, bool) {
	if o == nil || o.AvailabilityNotificationsEnabled == nil {
		return nil, false
	}
	return o.AvailabilityNotificationsEnabled, true
}

// HasAvailabilityNotificationsEnabled returns a boolean if a field has been set.
func (o *PrivateSyntheticLocationAllOf) HasAvailabilityNotificationsEnabled() bool {
	if o != nil && o.AvailabilityNotificationsEnabled != nil {
		return true
	}

	return false
}

// SetAvailabilityNotificationsEnabled gets a reference to the given bool and assigns it to the AvailabilityNotificationsEnabled field.
func (o *PrivateSyntheticLocationAllOf) SetAvailabilityNotificationsEnabled(v bool) {
	o.AvailabilityNotificationsEnabled = &v
}

func (o PrivateSyntheticLocationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}
	if o.AvailabilityLocationOutage != nil {
		toSerialize["availabilityLocationOutage"] = o.AvailabilityLocationOutage
	}
	if o.AvailabilityNodeOutage != nil {
		toSerialize["availabilityNodeOutage"] = o.AvailabilityNodeOutage
	}
	if o.LocationNodeOutageDelayInMinutes != nil {
		toSerialize["locationNodeOutageDelayInMinutes"] = o.LocationNodeOutageDelayInMinutes
	}
	if o.AvailabilityNotificationsEnabled != nil {
		toSerialize["availabilityNotificationsEnabled"] = o.AvailabilityNotificationsEnabled
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateSyntheticLocationAllOf struct {
	value *PrivateSyntheticLocationAllOf
	isSet bool
}

func (v NullablePrivateSyntheticLocationAllOf) Get() *PrivateSyntheticLocationAllOf {
	return v.value
}

func (v *NullablePrivateSyntheticLocationAllOf) Set(val *PrivateSyntheticLocationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateSyntheticLocationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateSyntheticLocationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateSyntheticLocationAllOf(val *PrivateSyntheticLocationAllOf) *NullablePrivateSyntheticLocationAllOf {
	return &NullablePrivateSyntheticLocationAllOf{value: val, isSet: true}
}

func (v NullablePrivateSyntheticLocationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateSyntheticLocationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


