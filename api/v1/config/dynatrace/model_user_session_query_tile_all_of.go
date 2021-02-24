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

// UserSessionQueryTileAllOf struct for UserSessionQueryTileAllOf
type UserSessionQueryTileAllOf struct {
	// The name of the tile, set by user.
	CustomName *string `json:"customName,omitempty"`
	// A [user session query](https://www.dynatrace.com/support/help/shortlink/usql-info) executed by the tile.
	Query *string `json:"query,omitempty"`
	// The visualization of the tile.
	Type *string `json:"type,omitempty"`
	// The comparison timeframe of the query.    If specified, you additionally get the results of the same query with the specified time shift.
	TimeFrameShift *string `json:"timeFrameShift,omitempty"`
	VisualizationConfig *UserSessionQueryTileConfiguration `json:"visualizationConfig,omitempty"`
	// The limit of the results, if not set will use the default value of the system
	Limit *int32 `json:"limit,omitempty"`
}

// NewUserSessionQueryTileAllOf instantiates a new UserSessionQueryTileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSessionQueryTileAllOf() *UserSessionQueryTileAllOf {
	this := UserSessionQueryTileAllOf{}
	return &this
}

// NewUserSessionQueryTileAllOfWithDefaults instantiates a new UserSessionQueryTileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSessionQueryTileAllOfWithDefaults() *UserSessionQueryTileAllOf {
	this := UserSessionQueryTileAllOf{}
	return &this
}

// GetCustomName returns the CustomName field value if set, zero value otherwise.
func (o *UserSessionQueryTileAllOf) GetCustomName() string {
	if o == nil || o.CustomName == nil {
		var ret string
		return ret
	}
	return *o.CustomName
}

// GetCustomNameOk returns a tuple with the CustomName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionQueryTileAllOf) GetCustomNameOk() (*string, bool) {
	if o == nil || o.CustomName == nil {
		return nil, false
	}
	return o.CustomName, true
}

// HasCustomName returns a boolean if a field has been set.
func (o *UserSessionQueryTileAllOf) HasCustomName() bool {
	if o != nil && o.CustomName != nil {
		return true
	}

	return false
}

// SetCustomName gets a reference to the given string and assigns it to the CustomName field.
func (o *UserSessionQueryTileAllOf) SetCustomName(v string) {
	o.CustomName = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *UserSessionQueryTileAllOf) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionQueryTileAllOf) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *UserSessionQueryTileAllOf) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *UserSessionQueryTileAllOf) SetQuery(v string) {
	o.Query = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserSessionQueryTileAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionQueryTileAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserSessionQueryTileAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserSessionQueryTileAllOf) SetType(v string) {
	o.Type = &v
}

// GetTimeFrameShift returns the TimeFrameShift field value if set, zero value otherwise.
func (o *UserSessionQueryTileAllOf) GetTimeFrameShift() string {
	if o == nil || o.TimeFrameShift == nil {
		var ret string
		return ret
	}
	return *o.TimeFrameShift
}

// GetTimeFrameShiftOk returns a tuple with the TimeFrameShift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionQueryTileAllOf) GetTimeFrameShiftOk() (*string, bool) {
	if o == nil || o.TimeFrameShift == nil {
		return nil, false
	}
	return o.TimeFrameShift, true
}

// HasTimeFrameShift returns a boolean if a field has been set.
func (o *UserSessionQueryTileAllOf) HasTimeFrameShift() bool {
	if o != nil && o.TimeFrameShift != nil {
		return true
	}

	return false
}

// SetTimeFrameShift gets a reference to the given string and assigns it to the TimeFrameShift field.
func (o *UserSessionQueryTileAllOf) SetTimeFrameShift(v string) {
	o.TimeFrameShift = &v
}

// GetVisualizationConfig returns the VisualizationConfig field value if set, zero value otherwise.
func (o *UserSessionQueryTileAllOf) GetVisualizationConfig() UserSessionQueryTileConfiguration {
	if o == nil || o.VisualizationConfig == nil {
		var ret UserSessionQueryTileConfiguration
		return ret
	}
	return *o.VisualizationConfig
}

// GetVisualizationConfigOk returns a tuple with the VisualizationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionQueryTileAllOf) GetVisualizationConfigOk() (*UserSessionQueryTileConfiguration, bool) {
	if o == nil || o.VisualizationConfig == nil {
		return nil, false
	}
	return o.VisualizationConfig, true
}

// HasVisualizationConfig returns a boolean if a field has been set.
func (o *UserSessionQueryTileAllOf) HasVisualizationConfig() bool {
	if o != nil && o.VisualizationConfig != nil {
		return true
	}

	return false
}

// SetVisualizationConfig gets a reference to the given UserSessionQueryTileConfiguration and assigns it to the VisualizationConfig field.
func (o *UserSessionQueryTileAllOf) SetVisualizationConfig(v UserSessionQueryTileConfiguration) {
	o.VisualizationConfig = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *UserSessionQueryTileAllOf) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionQueryTileAllOf) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *UserSessionQueryTileAllOf) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *UserSessionQueryTileAllOf) SetLimit(v int32) {
	o.Limit = &v
}

func (o UserSessionQueryTileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomName != nil {
		toSerialize["customName"] = o.CustomName
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.TimeFrameShift != nil {
		toSerialize["timeFrameShift"] = o.TimeFrameShift
	}
	if o.VisualizationConfig != nil {
		toSerialize["visualizationConfig"] = o.VisualizationConfig
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

type NullableUserSessionQueryTileAllOf struct {
	value *UserSessionQueryTileAllOf
	isSet bool
}

func (v NullableUserSessionQueryTileAllOf) Get() *UserSessionQueryTileAllOf {
	return v.value
}

func (v *NullableUserSessionQueryTileAllOf) Set(val *UserSessionQueryTileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSessionQueryTileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSessionQueryTileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSessionQueryTileAllOf(val *UserSessionQueryTileAllOf) *NullableUserSessionQueryTileAllOf {
	return &NullableUserSessionQueryTileAllOf{value: val, isSet: true}
}

func (v NullableUserSessionQueryTileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSessionQueryTileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


