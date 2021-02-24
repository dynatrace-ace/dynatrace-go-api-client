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

// UserSessionQueryTile Configuration of a User session query tile.
type UserSessionQueryTile struct {
	Tile
	// The name of the tile, set by user.
	CustomName string `json:"customName"`
	// A [user session query](https://www.dynatrace.com/support/help/shortlink/usql-info) executed by the tile.
	Query string `json:"query"`
	// The visualization of the tile.
	Type string `json:"type"`
	// The comparison timeframe of the query.    If specified, you additionally get the results of the same query with the specified time shift.
	TimeFrameShift *string `json:"timeFrameShift,omitempty"`
	VisualizationConfig *UserSessionQueryTileConfiguration `json:"visualizationConfig,omitempty"`
	// The limit of the results, if not set will use the default value of the system
	Limit *int32 `json:"limit,omitempty"`
}

// NewUserSessionQueryTile instantiates a new UserSessionQueryTile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSessionQueryTile(customName string, query string, type_ string, ) *UserSessionQueryTile {
	this := UserSessionQueryTile{}
	this.CustomName = customName
	this.Query = query
	this.Type = type_
	return &this
}

// NewUserSessionQueryTileWithDefaults instantiates a new UserSessionQueryTile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSessionQueryTileWithDefaults() *UserSessionQueryTile {
	this := UserSessionQueryTile{}
	return &this
}

// GetCustomName returns the CustomName field value
func (o *UserSessionQueryTile) GetCustomName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CustomName
}

// GetCustomNameOk returns a tuple with the CustomName field value
// and a boolean to check if the value has been set.
func (o *UserSessionQueryTile) GetCustomNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomName, true
}

// SetCustomName sets field value
func (o *UserSessionQueryTile) SetCustomName(v string) {
	o.CustomName = v
}

// GetQuery returns the Query field value
func (o *UserSessionQueryTile) GetQuery() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *UserSessionQueryTile) GetQueryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *UserSessionQueryTile) SetQuery(v string) {
	o.Query = v
}

// GetType returns the Type field value
func (o *UserSessionQueryTile) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserSessionQueryTile) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserSessionQueryTile) SetType(v string) {
	o.Type = v
}

// GetTimeFrameShift returns the TimeFrameShift field value if set, zero value otherwise.
func (o *UserSessionQueryTile) GetTimeFrameShift() string {
	if o == nil || o.TimeFrameShift == nil {
		var ret string
		return ret
	}
	return *o.TimeFrameShift
}

// GetTimeFrameShiftOk returns a tuple with the TimeFrameShift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionQueryTile) GetTimeFrameShiftOk() (*string, bool) {
	if o == nil || o.TimeFrameShift == nil {
		return nil, false
	}
	return o.TimeFrameShift, true
}

// HasTimeFrameShift returns a boolean if a field has been set.
func (o *UserSessionQueryTile) HasTimeFrameShift() bool {
	if o != nil && o.TimeFrameShift != nil {
		return true
	}

	return false
}

// SetTimeFrameShift gets a reference to the given string and assigns it to the TimeFrameShift field.
func (o *UserSessionQueryTile) SetTimeFrameShift(v string) {
	o.TimeFrameShift = &v
}

// GetVisualizationConfig returns the VisualizationConfig field value if set, zero value otherwise.
func (o *UserSessionQueryTile) GetVisualizationConfig() UserSessionQueryTileConfiguration {
	if o == nil || o.VisualizationConfig == nil {
		var ret UserSessionQueryTileConfiguration
		return ret
	}
	return *o.VisualizationConfig
}

// GetVisualizationConfigOk returns a tuple with the VisualizationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionQueryTile) GetVisualizationConfigOk() (*UserSessionQueryTileConfiguration, bool) {
	if o == nil || o.VisualizationConfig == nil {
		return nil, false
	}
	return o.VisualizationConfig, true
}

// HasVisualizationConfig returns a boolean if a field has been set.
func (o *UserSessionQueryTile) HasVisualizationConfig() bool {
	if o != nil && o.VisualizationConfig != nil {
		return true
	}

	return false
}

// SetVisualizationConfig gets a reference to the given UserSessionQueryTileConfiguration and assigns it to the VisualizationConfig field.
func (o *UserSessionQueryTile) SetVisualizationConfig(v UserSessionQueryTileConfiguration) {
	o.VisualizationConfig = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *UserSessionQueryTile) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionQueryTile) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *UserSessionQueryTile) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *UserSessionQueryTile) SetLimit(v int32) {
	o.Limit = &v
}

func (o UserSessionQueryTile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTile, errTile := json.Marshal(o.Tile)
	if errTile != nil {
		return []byte{}, errTile
	}
	errTile = json.Unmarshal([]byte(serializedTile), &toSerialize)
	if errTile != nil {
		return []byte{}, errTile
	}
	if true {
		toSerialize["customName"] = o.CustomName
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if true {
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

type NullableUserSessionQueryTile struct {
	value *UserSessionQueryTile
	isSet bool
}

func (v NullableUserSessionQueryTile) Get() *UserSessionQueryTile {
	return v.value
}

func (v *NullableUserSessionQueryTile) Set(val *UserSessionQueryTile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSessionQueryTile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSessionQueryTile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSessionQueryTile(val *UserSessionQueryTile) *NullableUserSessionQueryTile {
	return &NullableUserSessionQueryTile{value: val, isSet: true}
}

func (v NullableUserSessionQueryTile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSessionQueryTile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


