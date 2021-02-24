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

// LogMetricConfig Custom log metric definition.
type LogMetricConfig struct {
	// The unique key of the metric.   The key must have the `calc:log.` prefix.
	MetricKey string `json:"metricKey"`
	// The metric is enabled (`true`) or disabled (`false`).
	Active *bool `json:"active,omitempty"`
	// The name of the metric, displayed in the UI.
	DisplayName string `json:"displayName"`
	// The unit of the metric.
	Unit string `json:"unit"`
	// The display name of the unit.    Only applicable if the **unit** is set to `UNSPECIFIED`.
	UnitDisplayName *string `json:"unitDisplayName,omitempty"`
	// The pattern to look for in logs.    Use the [Dynatrace search query language](https://www.dynatrace.com/support/help/shortlink/log-viewer#search-for-text-patterns-in-log-files) to specify it. Quotes must be escaped.    To return all results, leave the field blank.
	SearchString string `json:"searchString"`
	// The type of the metric data points calculation. For now the only allowed value is `OCCURRENCES`.
	MetricValueType string `json:"metricValueType"`
	ColumnDefiningValue *ColumnDefinition `json:"columnDefiningValue,omitempty"`
	// A list of filters to define the logs to look into.    If several filters are specified, the OR logic applies.
	LogSourceFilters []LogSourceFilter `json:"logSourceFilters"`
}

// NewLogMetricConfig instantiates a new LogMetricConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogMetricConfig(metricKey string, displayName string, unit string, searchString string, metricValueType string, logSourceFilters []LogSourceFilter, ) *LogMetricConfig {
	this := LogMetricConfig{}
	this.MetricKey = metricKey
	this.DisplayName = displayName
	this.Unit = unit
	this.SearchString = searchString
	this.MetricValueType = metricValueType
	this.LogSourceFilters = logSourceFilters
	return &this
}

// NewLogMetricConfigWithDefaults instantiates a new LogMetricConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogMetricConfigWithDefaults() *LogMetricConfig {
	this := LogMetricConfig{}
	return &this
}

// GetMetricKey returns the MetricKey field value
func (o *LogMetricConfig) GetMetricKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.MetricKey
}

// GetMetricKeyOk returns a tuple with the MetricKey field value
// and a boolean to check if the value has been set.
func (o *LogMetricConfig) GetMetricKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MetricKey, true
}

// SetMetricKey sets field value
func (o *LogMetricConfig) SetMetricKey(v string) {
	o.MetricKey = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *LogMetricConfig) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogMetricConfig) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *LogMetricConfig) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *LogMetricConfig) SetActive(v bool) {
	o.Active = &v
}

// GetDisplayName returns the DisplayName field value
func (o *LogMetricConfig) GetDisplayName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *LogMetricConfig) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *LogMetricConfig) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetUnit returns the Unit field value
func (o *LogMetricConfig) GetUnit() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *LogMetricConfig) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *LogMetricConfig) SetUnit(v string) {
	o.Unit = v
}

// GetUnitDisplayName returns the UnitDisplayName field value if set, zero value otherwise.
func (o *LogMetricConfig) GetUnitDisplayName() string {
	if o == nil || o.UnitDisplayName == nil {
		var ret string
		return ret
	}
	return *o.UnitDisplayName
}

// GetUnitDisplayNameOk returns a tuple with the UnitDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogMetricConfig) GetUnitDisplayNameOk() (*string, bool) {
	if o == nil || o.UnitDisplayName == nil {
		return nil, false
	}
	return o.UnitDisplayName, true
}

// HasUnitDisplayName returns a boolean if a field has been set.
func (o *LogMetricConfig) HasUnitDisplayName() bool {
	if o != nil && o.UnitDisplayName != nil {
		return true
	}

	return false
}

// SetUnitDisplayName gets a reference to the given string and assigns it to the UnitDisplayName field.
func (o *LogMetricConfig) SetUnitDisplayName(v string) {
	o.UnitDisplayName = &v
}

// GetSearchString returns the SearchString field value
func (o *LogMetricConfig) GetSearchString() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SearchString
}

// GetSearchStringOk returns a tuple with the SearchString field value
// and a boolean to check if the value has been set.
func (o *LogMetricConfig) GetSearchStringOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SearchString, true
}

// SetSearchString sets field value
func (o *LogMetricConfig) SetSearchString(v string) {
	o.SearchString = v
}

// GetMetricValueType returns the MetricValueType field value
func (o *LogMetricConfig) GetMetricValueType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.MetricValueType
}

// GetMetricValueTypeOk returns a tuple with the MetricValueType field value
// and a boolean to check if the value has been set.
func (o *LogMetricConfig) GetMetricValueTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MetricValueType, true
}

// SetMetricValueType sets field value
func (o *LogMetricConfig) SetMetricValueType(v string) {
	o.MetricValueType = v
}

// GetColumnDefiningValue returns the ColumnDefiningValue field value if set, zero value otherwise.
func (o *LogMetricConfig) GetColumnDefiningValue() ColumnDefinition {
	if o == nil || o.ColumnDefiningValue == nil {
		var ret ColumnDefinition
		return ret
	}
	return *o.ColumnDefiningValue
}

// GetColumnDefiningValueOk returns a tuple with the ColumnDefiningValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogMetricConfig) GetColumnDefiningValueOk() (*ColumnDefinition, bool) {
	if o == nil || o.ColumnDefiningValue == nil {
		return nil, false
	}
	return o.ColumnDefiningValue, true
}

// HasColumnDefiningValue returns a boolean if a field has been set.
func (o *LogMetricConfig) HasColumnDefiningValue() bool {
	if o != nil && o.ColumnDefiningValue != nil {
		return true
	}

	return false
}

// SetColumnDefiningValue gets a reference to the given ColumnDefinition and assigns it to the ColumnDefiningValue field.
func (o *LogMetricConfig) SetColumnDefiningValue(v ColumnDefinition) {
	o.ColumnDefiningValue = &v
}

// GetLogSourceFilters returns the LogSourceFilters field value
func (o *LogMetricConfig) GetLogSourceFilters() []LogSourceFilter {
	if o == nil  {
		var ret []LogSourceFilter
		return ret
	}

	return o.LogSourceFilters
}

// GetLogSourceFiltersOk returns a tuple with the LogSourceFilters field value
// and a boolean to check if the value has been set.
func (o *LogMetricConfig) GetLogSourceFiltersOk() (*[]LogSourceFilter, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogSourceFilters, true
}

// SetLogSourceFilters sets field value
func (o *LogMetricConfig) SetLogSourceFilters(v []LogSourceFilter) {
	o.LogSourceFilters = v
}

func (o LogMetricConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metricKey"] = o.MetricKey
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	if o.UnitDisplayName != nil {
		toSerialize["unitDisplayName"] = o.UnitDisplayName
	}
	if true {
		toSerialize["searchString"] = o.SearchString
	}
	if true {
		toSerialize["metricValueType"] = o.MetricValueType
	}
	if o.ColumnDefiningValue != nil {
		toSerialize["columnDefiningValue"] = o.ColumnDefiningValue
	}
	if true {
		toSerialize["logSourceFilters"] = o.LogSourceFilters
	}
	return json.Marshal(toSerialize)
}

type NullableLogMetricConfig struct {
	value *LogMetricConfig
	isSet bool
}

func (v NullableLogMetricConfig) Get() *LogMetricConfig {
	return v.value
}

func (v *NullableLogMetricConfig) Set(val *LogMetricConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableLogMetricConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableLogMetricConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogMetricConfig(val *LogMetricConfig) *NullableLogMetricConfig {
	return &NullableLogMetricConfig{value: val, isSet: true}
}

func (v NullableLogMetricConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogMetricConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


