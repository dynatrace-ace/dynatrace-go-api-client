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

// DashboardMetadata Parameters of a dashboard.
type DashboardMetadata struct {
	// The name of the dashboard.
	Name string `json:"name"`
	// The dashboard is shared (`true`) or private (`false`).
	Shared *bool `json:"shared,omitempty"`
	// The owner of the dashboard.
	Owner *string `json:"owner,omitempty"`
	SharingDetails *SharingInfo `json:"sharingDetails,omitempty"`
	DashboardFilter *DashboardFilter `json:"dashboardFilter,omitempty"`
	// A set of tags assigned to the dashboard.
	Tags *[]string `json:"tags,omitempty"`
	// The dashboard is a preset (`true`)
	Preset *bool `json:"preset,omitempty"`
	// A set of all possible global dashboard filters that can be applied to dashboard
	ValidFilterKeys *[]string `json:"validFilterKeys,omitempty"`
}

// NewDashboardMetadata instantiates a new DashboardMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardMetadata(name string, ) *DashboardMetadata {
	this := DashboardMetadata{}
	this.Name = name
	return &this
}

// NewDashboardMetadataWithDefaults instantiates a new DashboardMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardMetadataWithDefaults() *DashboardMetadata {
	this := DashboardMetadata{}
	return &this
}

// GetName returns the Name field value
func (o *DashboardMetadata) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DashboardMetadata) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DashboardMetadata) SetName(v string) {
	o.Name = v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *DashboardMetadata) GetShared() bool {
	if o == nil || o.Shared == nil {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardMetadata) GetSharedOk() (*bool, bool) {
	if o == nil || o.Shared == nil {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *DashboardMetadata) HasShared() bool {
	if o != nil && o.Shared != nil {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *DashboardMetadata) SetShared(v bool) {
	o.Shared = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *DashboardMetadata) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardMetadata) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *DashboardMetadata) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *DashboardMetadata) SetOwner(v string) {
	o.Owner = &v
}

// GetSharingDetails returns the SharingDetails field value if set, zero value otherwise.
func (o *DashboardMetadata) GetSharingDetails() SharingInfo {
	if o == nil || o.SharingDetails == nil {
		var ret SharingInfo
		return ret
	}
	return *o.SharingDetails
}

// GetSharingDetailsOk returns a tuple with the SharingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardMetadata) GetSharingDetailsOk() (*SharingInfo, bool) {
	if o == nil || o.SharingDetails == nil {
		return nil, false
	}
	return o.SharingDetails, true
}

// HasSharingDetails returns a boolean if a field has been set.
func (o *DashboardMetadata) HasSharingDetails() bool {
	if o != nil && o.SharingDetails != nil {
		return true
	}

	return false
}

// SetSharingDetails gets a reference to the given SharingInfo and assigns it to the SharingDetails field.
func (o *DashboardMetadata) SetSharingDetails(v SharingInfo) {
	o.SharingDetails = &v
}

// GetDashboardFilter returns the DashboardFilter field value if set, zero value otherwise.
func (o *DashboardMetadata) GetDashboardFilter() DashboardFilter {
	if o == nil || o.DashboardFilter == nil {
		var ret DashboardFilter
		return ret
	}
	return *o.DashboardFilter
}

// GetDashboardFilterOk returns a tuple with the DashboardFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardMetadata) GetDashboardFilterOk() (*DashboardFilter, bool) {
	if o == nil || o.DashboardFilter == nil {
		return nil, false
	}
	return o.DashboardFilter, true
}

// HasDashboardFilter returns a boolean if a field has been set.
func (o *DashboardMetadata) HasDashboardFilter() bool {
	if o != nil && o.DashboardFilter != nil {
		return true
	}

	return false
}

// SetDashboardFilter gets a reference to the given DashboardFilter and assigns it to the DashboardFilter field.
func (o *DashboardMetadata) SetDashboardFilter(v DashboardFilter) {
	o.DashboardFilter = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DashboardMetadata) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardMetadata) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DashboardMetadata) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DashboardMetadata) SetTags(v []string) {
	o.Tags = &v
}

// GetPreset returns the Preset field value if set, zero value otherwise.
func (o *DashboardMetadata) GetPreset() bool {
	if o == nil || o.Preset == nil {
		var ret bool
		return ret
	}
	return *o.Preset
}

// GetPresetOk returns a tuple with the Preset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardMetadata) GetPresetOk() (*bool, bool) {
	if o == nil || o.Preset == nil {
		return nil, false
	}
	return o.Preset, true
}

// HasPreset returns a boolean if a field has been set.
func (o *DashboardMetadata) HasPreset() bool {
	if o != nil && o.Preset != nil {
		return true
	}

	return false
}

// SetPreset gets a reference to the given bool and assigns it to the Preset field.
func (o *DashboardMetadata) SetPreset(v bool) {
	o.Preset = &v
}

// GetValidFilterKeys returns the ValidFilterKeys field value if set, zero value otherwise.
func (o *DashboardMetadata) GetValidFilterKeys() []string {
	if o == nil || o.ValidFilterKeys == nil {
		var ret []string
		return ret
	}
	return *o.ValidFilterKeys
}

// GetValidFilterKeysOk returns a tuple with the ValidFilterKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardMetadata) GetValidFilterKeysOk() (*[]string, bool) {
	if o == nil || o.ValidFilterKeys == nil {
		return nil, false
	}
	return o.ValidFilterKeys, true
}

// HasValidFilterKeys returns a boolean if a field has been set.
func (o *DashboardMetadata) HasValidFilterKeys() bool {
	if o != nil && o.ValidFilterKeys != nil {
		return true
	}

	return false
}

// SetValidFilterKeys gets a reference to the given []string and assigns it to the ValidFilterKeys field.
func (o *DashboardMetadata) SetValidFilterKeys(v []string) {
	o.ValidFilterKeys = &v
}

func (o DashboardMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Shared != nil {
		toSerialize["shared"] = o.Shared
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.SharingDetails != nil {
		toSerialize["sharingDetails"] = o.SharingDetails
	}
	if o.DashboardFilter != nil {
		toSerialize["dashboardFilter"] = o.DashboardFilter
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Preset != nil {
		toSerialize["preset"] = o.Preset
	}
	if o.ValidFilterKeys != nil {
		toSerialize["validFilterKeys"] = o.ValidFilterKeys
	}
	return json.Marshal(toSerialize)
}

type NullableDashboardMetadata struct {
	value *DashboardMetadata
	isSet bool
}

func (v NullableDashboardMetadata) Get() *DashboardMetadata {
	return v.value
}

func (v *NullableDashboardMetadata) Set(val *DashboardMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardMetadata(val *DashboardMetadata) *NullableDashboardMetadata {
	return &NullableDashboardMetadata{value: val, isSet: true}
}

func (v NullableDashboardMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


