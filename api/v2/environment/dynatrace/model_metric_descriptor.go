/*
 * Dynatrace Environment API
 *
 *  Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// MetricDescriptor The descriptor of a metric.
type MetricDescriptor struct {
	// List of admissible primary entity types for this metric. Can be used for the `type` predicate in the `entitySelector`.
	EntityType *[]string `json:"entityType,omitempty"`
	// The fully qualified key of the metric.   If a transformation has been used it is reflected in the metric key.
	MetricId string `json:"metricId"`
	// If `true` the usage of metric consumes [Davis data units](https://dt-url.net/ddu).    Metric expressions don't return this field.
	DduBillable *bool `json:"dduBillable,omitempty"`
	// The list of allowed aggregations for this metric.
	AggregationTypes *[]string `json:"aggregationTypes,omitempty"`
	DefaultAggregation *MetricDefaultAggregation `json:"defaultAggregation,omitempty"`
	// The fine metric division (for example, process group and process ID for some process-related metric).   For [ingested metrics](https://dt-url.net/5d63ic1), dimensions that doesn't have have any data within the last 15 days are omitted. 
	DimensionDefinitions *[]MetricDimensionDefinition `json:"dimensionDefinitions,omitempty"`
	// The minimum value of the metric.   Metric expressions don't return this field.
	MinimumValue *float64 `json:"minimumValue,omitempty"`
	// The metric is (`true`) or is not (`false`) root cause relevant.   Metric expressions don't return this field.
	RootCauseRelevant *bool `json:"rootCauseRelevant,omitempty"`
	// The timestamp when the metric was last written.   Has the value of `null` for metric expressions or if the data has never been written.
	LastWritten *int64 `json:"lastWritten,omitempty"`
	MetricValueType *MetricValueType `json:"metricValueType,omitempty"`
	// The maximum value of the metric.   Metric expressions don't return this field.
	MaximumValue *float64 `json:"maximumValue,omitempty"`
	// The latency (in minutes) to how long it takes before a new metric data point is available in Monitoring after it is written.   Metric expressions don't return this field.
	Latency *int64 `json:"latency,omitempty"`
	// The metric is (`true`) or is not (`false`) impact relevant.   Metric expressions don't return this field.
	ImpactRelevant *bool `json:"impactRelevant,omitempty"`
	// The timestamp of metric creation.   Built-in metrics and metric expressions have the value of `null`.
	Created *int64 `json:"created,omitempty"`
	// A list of potential warnings that affect this ID. For example deprecated feature usage etc.
	Warnings *[]string `json:"warnings,omitempty"`
	// The name of the metric in the user interface.
	DisplayName *string `json:"displayName,omitempty"`
	// A short description of the metric.
	Description *string `json:"description,omitempty"`
	// Transform operators that could be appended to the current transformation list. Must be enabled with the \"fields\" parameter on `/metrics` and is always present on `/metrics/{metricId}`.
	Transformations *[]string `json:"transformations,omitempty"`
	// The unit of the metric.
	Unit *string `json:"unit,omitempty"`
	// The tags applied to the metric.    Metric expressions don't return this field.
	Tags *[]string `json:"tags,omitempty"`
}

// NewMetricDescriptor instantiates a new MetricDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricDescriptor(metricId string) *MetricDescriptor {
	this := MetricDescriptor{}
	this.MetricId = metricId
	return &this
}

// NewMetricDescriptorWithDefaults instantiates a new MetricDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricDescriptorWithDefaults() *MetricDescriptor {
	this := MetricDescriptor{}
	return &this
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *MetricDescriptor) GetEntityType() []string {
	if o == nil || o.EntityType == nil {
		var ret []string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetEntityTypeOk() (*[]string, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *MetricDescriptor) HasEntityType() bool {
	if o != nil && o.EntityType != nil {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given []string and assigns it to the EntityType field.
func (o *MetricDescriptor) SetEntityType(v []string) {
	o.EntityType = &v
}

// GetMetricId returns the MetricId field value
func (o *MetricDescriptor) GetMetricId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetMetricIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MetricId, true
}

// SetMetricId sets field value
func (o *MetricDescriptor) SetMetricId(v string) {
	o.MetricId = v
}

// GetDduBillable returns the DduBillable field value if set, zero value otherwise.
func (o *MetricDescriptor) GetDduBillable() bool {
	if o == nil || o.DduBillable == nil {
		var ret bool
		return ret
	}
	return *o.DduBillable
}

// GetDduBillableOk returns a tuple with the DduBillable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetDduBillableOk() (*bool, bool) {
	if o == nil || o.DduBillable == nil {
		return nil, false
	}
	return o.DduBillable, true
}

// HasDduBillable returns a boolean if a field has been set.
func (o *MetricDescriptor) HasDduBillable() bool {
	if o != nil && o.DduBillable != nil {
		return true
	}

	return false
}

// SetDduBillable gets a reference to the given bool and assigns it to the DduBillable field.
func (o *MetricDescriptor) SetDduBillable(v bool) {
	o.DduBillable = &v
}

// GetAggregationTypes returns the AggregationTypes field value if set, zero value otherwise.
func (o *MetricDescriptor) GetAggregationTypes() []string {
	if o == nil || o.AggregationTypes == nil {
		var ret []string
		return ret
	}
	return *o.AggregationTypes
}

// GetAggregationTypesOk returns a tuple with the AggregationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetAggregationTypesOk() (*[]string, bool) {
	if o == nil || o.AggregationTypes == nil {
		return nil, false
	}
	return o.AggregationTypes, true
}

// HasAggregationTypes returns a boolean if a field has been set.
func (o *MetricDescriptor) HasAggregationTypes() bool {
	if o != nil && o.AggregationTypes != nil {
		return true
	}

	return false
}

// SetAggregationTypes gets a reference to the given []string and assigns it to the AggregationTypes field.
func (o *MetricDescriptor) SetAggregationTypes(v []string) {
	o.AggregationTypes = &v
}

// GetDefaultAggregation returns the DefaultAggregation field value if set, zero value otherwise.
func (o *MetricDescriptor) GetDefaultAggregation() MetricDefaultAggregation {
	if o == nil || o.DefaultAggregation == nil {
		var ret MetricDefaultAggregation
		return ret
	}
	return *o.DefaultAggregation
}

// GetDefaultAggregationOk returns a tuple with the DefaultAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetDefaultAggregationOk() (*MetricDefaultAggregation, bool) {
	if o == nil || o.DefaultAggregation == nil {
		return nil, false
	}
	return o.DefaultAggregation, true
}

// HasDefaultAggregation returns a boolean if a field has been set.
func (o *MetricDescriptor) HasDefaultAggregation() bool {
	if o != nil && o.DefaultAggregation != nil {
		return true
	}

	return false
}

// SetDefaultAggregation gets a reference to the given MetricDefaultAggregation and assigns it to the DefaultAggregation field.
func (o *MetricDescriptor) SetDefaultAggregation(v MetricDefaultAggregation) {
	o.DefaultAggregation = &v
}

// GetDimensionDefinitions returns the DimensionDefinitions field value if set, zero value otherwise.
func (o *MetricDescriptor) GetDimensionDefinitions() []MetricDimensionDefinition {
	if o == nil || o.DimensionDefinitions == nil {
		var ret []MetricDimensionDefinition
		return ret
	}
	return *o.DimensionDefinitions
}

// GetDimensionDefinitionsOk returns a tuple with the DimensionDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetDimensionDefinitionsOk() (*[]MetricDimensionDefinition, bool) {
	if o == nil || o.DimensionDefinitions == nil {
		return nil, false
	}
	return o.DimensionDefinitions, true
}

// HasDimensionDefinitions returns a boolean if a field has been set.
func (o *MetricDescriptor) HasDimensionDefinitions() bool {
	if o != nil && o.DimensionDefinitions != nil {
		return true
	}

	return false
}

// SetDimensionDefinitions gets a reference to the given []MetricDimensionDefinition and assigns it to the DimensionDefinitions field.
func (o *MetricDescriptor) SetDimensionDefinitions(v []MetricDimensionDefinition) {
	o.DimensionDefinitions = &v
}

// GetMinimumValue returns the MinimumValue field value if set, zero value otherwise.
func (o *MetricDescriptor) GetMinimumValue() float64 {
	if o == nil || o.MinimumValue == nil {
		var ret float64
		return ret
	}
	return *o.MinimumValue
}

// GetMinimumValueOk returns a tuple with the MinimumValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetMinimumValueOk() (*float64, bool) {
	if o == nil || o.MinimumValue == nil {
		return nil, false
	}
	return o.MinimumValue, true
}

// HasMinimumValue returns a boolean if a field has been set.
func (o *MetricDescriptor) HasMinimumValue() bool {
	if o != nil && o.MinimumValue != nil {
		return true
	}

	return false
}

// SetMinimumValue gets a reference to the given float64 and assigns it to the MinimumValue field.
func (o *MetricDescriptor) SetMinimumValue(v float64) {
	o.MinimumValue = &v
}

// GetRootCauseRelevant returns the RootCauseRelevant field value if set, zero value otherwise.
func (o *MetricDescriptor) GetRootCauseRelevant() bool {
	if o == nil || o.RootCauseRelevant == nil {
		var ret bool
		return ret
	}
	return *o.RootCauseRelevant
}

// GetRootCauseRelevantOk returns a tuple with the RootCauseRelevant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetRootCauseRelevantOk() (*bool, bool) {
	if o == nil || o.RootCauseRelevant == nil {
		return nil, false
	}
	return o.RootCauseRelevant, true
}

// HasRootCauseRelevant returns a boolean if a field has been set.
func (o *MetricDescriptor) HasRootCauseRelevant() bool {
	if o != nil && o.RootCauseRelevant != nil {
		return true
	}

	return false
}

// SetRootCauseRelevant gets a reference to the given bool and assigns it to the RootCauseRelevant field.
func (o *MetricDescriptor) SetRootCauseRelevant(v bool) {
	o.RootCauseRelevant = &v
}

// GetLastWritten returns the LastWritten field value if set, zero value otherwise.
func (o *MetricDescriptor) GetLastWritten() int64 {
	if o == nil || o.LastWritten == nil {
		var ret int64
		return ret
	}
	return *o.LastWritten
}

// GetLastWrittenOk returns a tuple with the LastWritten field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetLastWrittenOk() (*int64, bool) {
	if o == nil || o.LastWritten == nil {
		return nil, false
	}
	return o.LastWritten, true
}

// HasLastWritten returns a boolean if a field has been set.
func (o *MetricDescriptor) HasLastWritten() bool {
	if o != nil && o.LastWritten != nil {
		return true
	}

	return false
}

// SetLastWritten gets a reference to the given int64 and assigns it to the LastWritten field.
func (o *MetricDescriptor) SetLastWritten(v int64) {
	o.LastWritten = &v
}

// GetMetricValueType returns the MetricValueType field value if set, zero value otherwise.
func (o *MetricDescriptor) GetMetricValueType() MetricValueType {
	if o == nil || o.MetricValueType == nil {
		var ret MetricValueType
		return ret
	}
	return *o.MetricValueType
}

// GetMetricValueTypeOk returns a tuple with the MetricValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetMetricValueTypeOk() (*MetricValueType, bool) {
	if o == nil || o.MetricValueType == nil {
		return nil, false
	}
	return o.MetricValueType, true
}

// HasMetricValueType returns a boolean if a field has been set.
func (o *MetricDescriptor) HasMetricValueType() bool {
	if o != nil && o.MetricValueType != nil {
		return true
	}

	return false
}

// SetMetricValueType gets a reference to the given MetricValueType and assigns it to the MetricValueType field.
func (o *MetricDescriptor) SetMetricValueType(v MetricValueType) {
	o.MetricValueType = &v
}

// GetMaximumValue returns the MaximumValue field value if set, zero value otherwise.
func (o *MetricDescriptor) GetMaximumValue() float64 {
	if o == nil || o.MaximumValue == nil {
		var ret float64
		return ret
	}
	return *o.MaximumValue
}

// GetMaximumValueOk returns a tuple with the MaximumValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetMaximumValueOk() (*float64, bool) {
	if o == nil || o.MaximumValue == nil {
		return nil, false
	}
	return o.MaximumValue, true
}

// HasMaximumValue returns a boolean if a field has been set.
func (o *MetricDescriptor) HasMaximumValue() bool {
	if o != nil && o.MaximumValue != nil {
		return true
	}

	return false
}

// SetMaximumValue gets a reference to the given float64 and assigns it to the MaximumValue field.
func (o *MetricDescriptor) SetMaximumValue(v float64) {
	o.MaximumValue = &v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *MetricDescriptor) GetLatency() int64 {
	if o == nil || o.Latency == nil {
		var ret int64
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetLatencyOk() (*int64, bool) {
	if o == nil || o.Latency == nil {
		return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *MetricDescriptor) HasLatency() bool {
	if o != nil && o.Latency != nil {
		return true
	}

	return false
}

// SetLatency gets a reference to the given int64 and assigns it to the Latency field.
func (o *MetricDescriptor) SetLatency(v int64) {
	o.Latency = &v
}

// GetImpactRelevant returns the ImpactRelevant field value if set, zero value otherwise.
func (o *MetricDescriptor) GetImpactRelevant() bool {
	if o == nil || o.ImpactRelevant == nil {
		var ret bool
		return ret
	}
	return *o.ImpactRelevant
}

// GetImpactRelevantOk returns a tuple with the ImpactRelevant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetImpactRelevantOk() (*bool, bool) {
	if o == nil || o.ImpactRelevant == nil {
		return nil, false
	}
	return o.ImpactRelevant, true
}

// HasImpactRelevant returns a boolean if a field has been set.
func (o *MetricDescriptor) HasImpactRelevant() bool {
	if o != nil && o.ImpactRelevant != nil {
		return true
	}

	return false
}

// SetImpactRelevant gets a reference to the given bool and assigns it to the ImpactRelevant field.
func (o *MetricDescriptor) SetImpactRelevant(v bool) {
	o.ImpactRelevant = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *MetricDescriptor) GetCreated() int64 {
	if o == nil || o.Created == nil {
		var ret int64
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetCreatedOk() (*int64, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *MetricDescriptor) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int64 and assigns it to the Created field.
func (o *MetricDescriptor) SetCreated(v int64) {
	o.Created = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *MetricDescriptor) GetWarnings() []string {
	if o == nil || o.Warnings == nil {
		var ret []string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetWarningsOk() (*[]string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *MetricDescriptor) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *MetricDescriptor) SetWarnings(v []string) {
	o.Warnings = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MetricDescriptor) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MetricDescriptor) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MetricDescriptor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MetricDescriptor) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MetricDescriptor) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MetricDescriptor) SetDescription(v string) {
	o.Description = &v
}

// GetTransformations returns the Transformations field value if set, zero value otherwise.
func (o *MetricDescriptor) GetTransformations() []string {
	if o == nil || o.Transformations == nil {
		var ret []string
		return ret
	}
	return *o.Transformations
}

// GetTransformationsOk returns a tuple with the Transformations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetTransformationsOk() (*[]string, bool) {
	if o == nil || o.Transformations == nil {
		return nil, false
	}
	return o.Transformations, true
}

// HasTransformations returns a boolean if a field has been set.
func (o *MetricDescriptor) HasTransformations() bool {
	if o != nil && o.Transformations != nil {
		return true
	}

	return false
}

// SetTransformations gets a reference to the given []string and assigns it to the Transformations field.
func (o *MetricDescriptor) SetTransformations(v []string) {
	o.Transformations = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *MetricDescriptor) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *MetricDescriptor) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *MetricDescriptor) SetUnit(v string) {
	o.Unit = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MetricDescriptor) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MetricDescriptor) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *MetricDescriptor) SetTags(v []string) {
	o.Tags = &v
}

func (o MetricDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntityType != nil {
		toSerialize["entityType"] = o.EntityType
	}
	if true {
		toSerialize["metricId"] = o.MetricId
	}
	if o.DduBillable != nil {
		toSerialize["dduBillable"] = o.DduBillable
	}
	if o.AggregationTypes != nil {
		toSerialize["aggregationTypes"] = o.AggregationTypes
	}
	if o.DefaultAggregation != nil {
		toSerialize["defaultAggregation"] = o.DefaultAggregation
	}
	if o.DimensionDefinitions != nil {
		toSerialize["dimensionDefinitions"] = o.DimensionDefinitions
	}
	if o.MinimumValue != nil {
		toSerialize["minimumValue"] = o.MinimumValue
	}
	if o.RootCauseRelevant != nil {
		toSerialize["rootCauseRelevant"] = o.RootCauseRelevant
	}
	if o.LastWritten != nil {
		toSerialize["lastWritten"] = o.LastWritten
	}
	if o.MetricValueType != nil {
		toSerialize["metricValueType"] = o.MetricValueType
	}
	if o.MaximumValue != nil {
		toSerialize["maximumValue"] = o.MaximumValue
	}
	if o.Latency != nil {
		toSerialize["latency"] = o.Latency
	}
	if o.ImpactRelevant != nil {
		toSerialize["impactRelevant"] = o.ImpactRelevant
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Transformations != nil {
		toSerialize["transformations"] = o.Transformations
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableMetricDescriptor struct {
	value *MetricDescriptor
	isSet bool
}

func (v NullableMetricDescriptor) Get() *MetricDescriptor {
	return v.value
}

func (v *NullableMetricDescriptor) Set(val *MetricDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricDescriptor(val *MetricDescriptor) *NullableMetricDescriptor {
	return &NullableMetricDescriptor{value: val, isSet: true}
}

func (v NullableMetricDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


