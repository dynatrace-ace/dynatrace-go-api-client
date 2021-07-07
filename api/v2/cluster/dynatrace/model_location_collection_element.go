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

// LocationCollectionElement A synthetic location.
type LocationCollectionElement struct {
	// The name of the location.
	Name string `json:"name"`
	// The Dynatrace entity ID of the location.
	EntityId string `json:"entityId"`
	// The type of the location.
	Type string `json:"type"`
	// The cloud provider where the location is hosted.    Only applicable to `PUBLIC` locations.
	CloudPlatform *string `json:"cloudPlatform,omitempty"`
	// The list of IP addresses assigned to the location.    Only applicable to `PUBLIC` locations.
	Ips *[]string `json:"ips,omitempty"`
	// The release stage of the location.
	Stage *string `json:"stage,omitempty"`
	// The status of the location.
	Status *string `json:"status,omitempty"`
	// The Dynatrace GeoLocation ID of the location.
	GeoLocationId string `json:"geoLocationId"`
}

// NewLocationCollectionElement instantiates a new LocationCollectionElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationCollectionElement(name string, entityId string, type_ string, geoLocationId string) *LocationCollectionElement {
	this := LocationCollectionElement{}
	this.Name = name
	this.EntityId = entityId
	this.Type = type_
	this.GeoLocationId = geoLocationId
	return &this
}

// NewLocationCollectionElementWithDefaults instantiates a new LocationCollectionElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationCollectionElementWithDefaults() *LocationCollectionElement {
	this := LocationCollectionElement{}
	return &this
}

// GetName returns the Name field value
func (o *LocationCollectionElement) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LocationCollectionElement) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LocationCollectionElement) SetName(v string) {
	o.Name = v
}

// GetEntityId returns the EntityId field value
func (o *LocationCollectionElement) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *LocationCollectionElement) GetEntityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *LocationCollectionElement) SetEntityId(v string) {
	o.EntityId = v
}

// GetType returns the Type field value
func (o *LocationCollectionElement) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LocationCollectionElement) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LocationCollectionElement) SetType(v string) {
	o.Type = v
}

// GetCloudPlatform returns the CloudPlatform field value if set, zero value otherwise.
func (o *LocationCollectionElement) GetCloudPlatform() string {
	if o == nil || o.CloudPlatform == nil {
		var ret string
		return ret
	}
	return *o.CloudPlatform
}

// GetCloudPlatformOk returns a tuple with the CloudPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationCollectionElement) GetCloudPlatformOk() (*string, bool) {
	if o == nil || o.CloudPlatform == nil {
		return nil, false
	}
	return o.CloudPlatform, true
}

// HasCloudPlatform returns a boolean if a field has been set.
func (o *LocationCollectionElement) HasCloudPlatform() bool {
	if o != nil && o.CloudPlatform != nil {
		return true
	}

	return false
}

// SetCloudPlatform gets a reference to the given string and assigns it to the CloudPlatform field.
func (o *LocationCollectionElement) SetCloudPlatform(v string) {
	o.CloudPlatform = &v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *LocationCollectionElement) GetIps() []string {
	if o == nil || o.Ips == nil {
		var ret []string
		return ret
	}
	return *o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationCollectionElement) GetIpsOk() (*[]string, bool) {
	if o == nil || o.Ips == nil {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *LocationCollectionElement) HasIps() bool {
	if o != nil && o.Ips != nil {
		return true
	}

	return false
}

// SetIps gets a reference to the given []string and assigns it to the Ips field.
func (o *LocationCollectionElement) SetIps(v []string) {
	o.Ips = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *LocationCollectionElement) GetStage() string {
	if o == nil || o.Stage == nil {
		var ret string
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationCollectionElement) GetStageOk() (*string, bool) {
	if o == nil || o.Stage == nil {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *LocationCollectionElement) HasStage() bool {
	if o != nil && o.Stage != nil {
		return true
	}

	return false
}

// SetStage gets a reference to the given string and assigns it to the Stage field.
func (o *LocationCollectionElement) SetStage(v string) {
	o.Stage = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LocationCollectionElement) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationCollectionElement) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LocationCollectionElement) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LocationCollectionElement) SetStatus(v string) {
	o.Status = &v
}

// GetGeoLocationId returns the GeoLocationId field value
func (o *LocationCollectionElement) GetGeoLocationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GeoLocationId
}

// GetGeoLocationIdOk returns a tuple with the GeoLocationId field value
// and a boolean to check if the value has been set.
func (o *LocationCollectionElement) GetGeoLocationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GeoLocationId, true
}

// SetGeoLocationId sets field value
func (o *LocationCollectionElement) SetGeoLocationId(v string) {
	o.GeoLocationId = v
}

func (o LocationCollectionElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["entityId"] = o.EntityId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.CloudPlatform != nil {
		toSerialize["cloudPlatform"] = o.CloudPlatform
	}
	if o.Ips != nil {
		toSerialize["ips"] = o.Ips
	}
	if o.Stage != nil {
		toSerialize["stage"] = o.Stage
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["geoLocationId"] = o.GeoLocationId
	}
	return json.Marshal(toSerialize)
}

type NullableLocationCollectionElement struct {
	value *LocationCollectionElement
	isSet bool
}

func (v NullableLocationCollectionElement) Get() *LocationCollectionElement {
	return v.value
}

func (v *NullableLocationCollectionElement) Set(val *LocationCollectionElement) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationCollectionElement) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationCollectionElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationCollectionElement(val *LocationCollectionElement) *NullableLocationCollectionElement {
	return &NullableLocationCollectionElement{value: val, isSet: true}
}

func (v NullableLocationCollectionElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationCollectionElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


