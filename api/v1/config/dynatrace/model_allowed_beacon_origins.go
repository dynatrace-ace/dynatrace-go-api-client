/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// AllowedBeaconOrigins Configuration of the allowed beacon origins for CORS requests.
type AllowedBeaconOrigins struct {
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// A list of allowed beacon origins for CORS requests.
	AllowedBeaconOrigins *[]BeaconDomainPattern `json:"allowedBeaconOrigins,omitempty"`
	// Discard (`true`) or keep (`false`) beacons without the **Origin** HTTP header on the BeaconForwarder.   If not set, `false` is used.
	RejectBeaconsWithoutOriginHeader *bool `json:"rejectBeaconsWithoutOriginHeader,omitempty"`
}

// NewAllowedBeaconOrigins instantiates a new AllowedBeaconOrigins object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowedBeaconOrigins() *AllowedBeaconOrigins {
	this := AllowedBeaconOrigins{}
	return &this
}

// NewAllowedBeaconOriginsWithDefaults instantiates a new AllowedBeaconOrigins object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowedBeaconOriginsWithDefaults() *AllowedBeaconOrigins {
	this := AllowedBeaconOrigins{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AllowedBeaconOrigins) GetMetadata() ConfigurationMetadata {
	if o == nil || o.Metadata == nil {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedBeaconOrigins) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AllowedBeaconOrigins) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *AllowedBeaconOrigins) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetAllowedBeaconOrigins returns the AllowedBeaconOrigins field value if set, zero value otherwise.
func (o *AllowedBeaconOrigins) GetAllowedBeaconOrigins() []BeaconDomainPattern {
	if o == nil || o.AllowedBeaconOrigins == nil {
		var ret []BeaconDomainPattern
		return ret
	}
	return *o.AllowedBeaconOrigins
}

// GetAllowedBeaconOriginsOk returns a tuple with the AllowedBeaconOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedBeaconOrigins) GetAllowedBeaconOriginsOk() (*[]BeaconDomainPattern, bool) {
	if o == nil || o.AllowedBeaconOrigins == nil {
		return nil, false
	}
	return o.AllowedBeaconOrigins, true
}

// HasAllowedBeaconOrigins returns a boolean if a field has been set.
func (o *AllowedBeaconOrigins) HasAllowedBeaconOrigins() bool {
	if o != nil && o.AllowedBeaconOrigins != nil {
		return true
	}

	return false
}

// SetAllowedBeaconOrigins gets a reference to the given []BeaconDomainPattern and assigns it to the AllowedBeaconOrigins field.
func (o *AllowedBeaconOrigins) SetAllowedBeaconOrigins(v []BeaconDomainPattern) {
	o.AllowedBeaconOrigins = &v
}

// GetRejectBeaconsWithoutOriginHeader returns the RejectBeaconsWithoutOriginHeader field value if set, zero value otherwise.
func (o *AllowedBeaconOrigins) GetRejectBeaconsWithoutOriginHeader() bool {
	if o == nil || o.RejectBeaconsWithoutOriginHeader == nil {
		var ret bool
		return ret
	}
	return *o.RejectBeaconsWithoutOriginHeader
}

// GetRejectBeaconsWithoutOriginHeaderOk returns a tuple with the RejectBeaconsWithoutOriginHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedBeaconOrigins) GetRejectBeaconsWithoutOriginHeaderOk() (*bool, bool) {
	if o == nil || o.RejectBeaconsWithoutOriginHeader == nil {
		return nil, false
	}
	return o.RejectBeaconsWithoutOriginHeader, true
}

// HasRejectBeaconsWithoutOriginHeader returns a boolean if a field has been set.
func (o *AllowedBeaconOrigins) HasRejectBeaconsWithoutOriginHeader() bool {
	if o != nil && o.RejectBeaconsWithoutOriginHeader != nil {
		return true
	}

	return false
}

// SetRejectBeaconsWithoutOriginHeader gets a reference to the given bool and assigns it to the RejectBeaconsWithoutOriginHeader field.
func (o *AllowedBeaconOrigins) SetRejectBeaconsWithoutOriginHeader(v bool) {
	o.RejectBeaconsWithoutOriginHeader = &v
}

func (o AllowedBeaconOrigins) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.AllowedBeaconOrigins != nil {
		toSerialize["allowedBeaconOrigins"] = o.AllowedBeaconOrigins
	}
	if o.RejectBeaconsWithoutOriginHeader != nil {
		toSerialize["rejectBeaconsWithoutOriginHeader"] = o.RejectBeaconsWithoutOriginHeader
	}
	return json.Marshal(toSerialize)
}

type NullableAllowedBeaconOrigins struct {
	value *AllowedBeaconOrigins
	isSet bool
}

func (v NullableAllowedBeaconOrigins) Get() *AllowedBeaconOrigins {
	return v.value
}

func (v *NullableAllowedBeaconOrigins) Set(val *AllowedBeaconOrigins) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedBeaconOrigins) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedBeaconOrigins) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedBeaconOrigins(val *AllowedBeaconOrigins) *NullableAllowedBeaconOrigins {
	return &NullableAllowedBeaconOrigins{value: val, isSet: true}
}

func (v NullableAllowedBeaconOrigins) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedBeaconOrigins) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


