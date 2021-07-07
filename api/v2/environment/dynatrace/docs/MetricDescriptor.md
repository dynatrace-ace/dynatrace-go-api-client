# MetricDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | Pointer to **[]string** | List of admissible primary entity types for this metric. Can be used for the &#x60;type&#x60; predicate in the &#x60;entitySelector&#x60;. | [optional] 
**MetricId** | **string** | The fully qualified key of the metric.   If a transformation has been used it is reflected in the metric key. | 
**DduBillable** | Pointer to **bool** | If &#x60;true&#x60; the usage of metric consumes [Davis data units](https://dt-url.net/ddu).    Metric expressions don&#39;t return this field. | [optional] 
**AggregationTypes** | Pointer to **[]string** | The list of allowed aggregations for this metric. | [optional] 
**DefaultAggregation** | Pointer to [**MetricDefaultAggregation**](MetricDefaultAggregation.md) |  | [optional] 
**DimensionDefinitions** | Pointer to [**[]MetricDimensionDefinition**](MetricDimensionDefinition.md) | The fine metric division (for example, process group and process ID for some process-related metric).   For [ingested metrics](https://dt-url.net/5d63ic1), dimensions that doesn&#39;t have have any data within the last 15 days are omitted.  | [optional] 
**MinimumValue** | Pointer to **float64** | The minimum value of the metric.   Metric expressions don&#39;t return this field. | [optional] 
**RootCauseRelevant** | Pointer to **bool** | The metric is (&#x60;true&#x60;) or is not (&#x60;false&#x60;) root cause relevant.   Metric expressions don&#39;t return this field. | [optional] 
**LastWritten** | Pointer to **int64** | The timestamp when the metric was last written.   Has the value of &#x60;null&#x60; for metric expressions or if the data has never been written. | [optional] 
**MetricValueType** | Pointer to [**MetricValueType**](MetricValueType.md) |  | [optional] 
**MaximumValue** | Pointer to **float64** | The maximum value of the metric.   Metric expressions don&#39;t return this field. | [optional] 
**Latency** | Pointer to **int64** | The latency (in minutes) to how long it takes before a new metric data point is available in Monitoring after it is written.   Metric expressions don&#39;t return this field. | [optional] 
**ImpactRelevant** | Pointer to **bool** | The metric is (&#x60;true&#x60;) or is not (&#x60;false&#x60;) impact relevant.   Metric expressions don&#39;t return this field. | [optional] 
**Created** | Pointer to **int64** | The timestamp of metric creation.   Built-in metrics and metric expressions have the value of &#x60;null&#x60;. | [optional] 
**Warnings** | Pointer to **[]string** | A list of potential warnings that affect this ID. For example deprecated feature usage etc. | [optional] 
**DisplayName** | Pointer to **string** | The name of the metric in the user interface. | [optional] 
**Description** | Pointer to **string** | A short description of the metric. | [optional] 
**Transformations** | Pointer to **[]string** | Transform operators that could be appended to the current transformation list. Must be enabled with the \&quot;fields\&quot; parameter on &#x60;/metrics&#x60; and is always present on &#x60;/metrics/{metricId}&#x60;. | [optional] 
**Unit** | Pointer to **string** | The unit of the metric. | [optional] 
**Tags** | Pointer to **[]string** | The tags applied to the metric.    Metric expressions don&#39;t return this field. | [optional] 

## Methods

### NewMetricDescriptor

`func NewMetricDescriptor(metricId string, ) *MetricDescriptor`

NewMetricDescriptor instantiates a new MetricDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDescriptorWithDefaults

`func NewMetricDescriptorWithDefaults() *MetricDescriptor`

NewMetricDescriptorWithDefaults instantiates a new MetricDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *MetricDescriptor) GetEntityType() []string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *MetricDescriptor) GetEntityTypeOk() (*[]string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *MetricDescriptor) SetEntityType(v []string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *MetricDescriptor) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetMetricId

`func (o *MetricDescriptor) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *MetricDescriptor) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *MetricDescriptor) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.


### GetDduBillable

`func (o *MetricDescriptor) GetDduBillable() bool`

GetDduBillable returns the DduBillable field if non-nil, zero value otherwise.

### GetDduBillableOk

`func (o *MetricDescriptor) GetDduBillableOk() (*bool, bool)`

GetDduBillableOk returns a tuple with the DduBillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDduBillable

`func (o *MetricDescriptor) SetDduBillable(v bool)`

SetDduBillable sets DduBillable field to given value.

### HasDduBillable

`func (o *MetricDescriptor) HasDduBillable() bool`

HasDduBillable returns a boolean if a field has been set.

### GetAggregationTypes

`func (o *MetricDescriptor) GetAggregationTypes() []string`

GetAggregationTypes returns the AggregationTypes field if non-nil, zero value otherwise.

### GetAggregationTypesOk

`func (o *MetricDescriptor) GetAggregationTypesOk() (*[]string, bool)`

GetAggregationTypesOk returns a tuple with the AggregationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationTypes

`func (o *MetricDescriptor) SetAggregationTypes(v []string)`

SetAggregationTypes sets AggregationTypes field to given value.

### HasAggregationTypes

`func (o *MetricDescriptor) HasAggregationTypes() bool`

HasAggregationTypes returns a boolean if a field has been set.

### GetDefaultAggregation

`func (o *MetricDescriptor) GetDefaultAggregation() MetricDefaultAggregation`

GetDefaultAggregation returns the DefaultAggregation field if non-nil, zero value otherwise.

### GetDefaultAggregationOk

`func (o *MetricDescriptor) GetDefaultAggregationOk() (*MetricDefaultAggregation, bool)`

GetDefaultAggregationOk returns a tuple with the DefaultAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAggregation

`func (o *MetricDescriptor) SetDefaultAggregation(v MetricDefaultAggregation)`

SetDefaultAggregation sets DefaultAggregation field to given value.

### HasDefaultAggregation

`func (o *MetricDescriptor) HasDefaultAggregation() bool`

HasDefaultAggregation returns a boolean if a field has been set.

### GetDimensionDefinitions

`func (o *MetricDescriptor) GetDimensionDefinitions() []MetricDimensionDefinition`

GetDimensionDefinitions returns the DimensionDefinitions field if non-nil, zero value otherwise.

### GetDimensionDefinitionsOk

`func (o *MetricDescriptor) GetDimensionDefinitionsOk() (*[]MetricDimensionDefinition, bool)`

GetDimensionDefinitionsOk returns a tuple with the DimensionDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionDefinitions

`func (o *MetricDescriptor) SetDimensionDefinitions(v []MetricDimensionDefinition)`

SetDimensionDefinitions sets DimensionDefinitions field to given value.

### HasDimensionDefinitions

`func (o *MetricDescriptor) HasDimensionDefinitions() bool`

HasDimensionDefinitions returns a boolean if a field has been set.

### GetMinimumValue

`func (o *MetricDescriptor) GetMinimumValue() float64`

GetMinimumValue returns the MinimumValue field if non-nil, zero value otherwise.

### GetMinimumValueOk

`func (o *MetricDescriptor) GetMinimumValueOk() (*float64, bool)`

GetMinimumValueOk returns a tuple with the MinimumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumValue

`func (o *MetricDescriptor) SetMinimumValue(v float64)`

SetMinimumValue sets MinimumValue field to given value.

### HasMinimumValue

`func (o *MetricDescriptor) HasMinimumValue() bool`

HasMinimumValue returns a boolean if a field has been set.

### GetRootCauseRelevant

`func (o *MetricDescriptor) GetRootCauseRelevant() bool`

GetRootCauseRelevant returns the RootCauseRelevant field if non-nil, zero value otherwise.

### GetRootCauseRelevantOk

`func (o *MetricDescriptor) GetRootCauseRelevantOk() (*bool, bool)`

GetRootCauseRelevantOk returns a tuple with the RootCauseRelevant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseRelevant

`func (o *MetricDescriptor) SetRootCauseRelevant(v bool)`

SetRootCauseRelevant sets RootCauseRelevant field to given value.

### HasRootCauseRelevant

`func (o *MetricDescriptor) HasRootCauseRelevant() bool`

HasRootCauseRelevant returns a boolean if a field has been set.

### GetLastWritten

`func (o *MetricDescriptor) GetLastWritten() int64`

GetLastWritten returns the LastWritten field if non-nil, zero value otherwise.

### GetLastWrittenOk

`func (o *MetricDescriptor) GetLastWrittenOk() (*int64, bool)`

GetLastWrittenOk returns a tuple with the LastWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWritten

`func (o *MetricDescriptor) SetLastWritten(v int64)`

SetLastWritten sets LastWritten field to given value.

### HasLastWritten

`func (o *MetricDescriptor) HasLastWritten() bool`

HasLastWritten returns a boolean if a field has been set.

### GetMetricValueType

`func (o *MetricDescriptor) GetMetricValueType() MetricValueType`

GetMetricValueType returns the MetricValueType field if non-nil, zero value otherwise.

### GetMetricValueTypeOk

`func (o *MetricDescriptor) GetMetricValueTypeOk() (*MetricValueType, bool)`

GetMetricValueTypeOk returns a tuple with the MetricValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValueType

`func (o *MetricDescriptor) SetMetricValueType(v MetricValueType)`

SetMetricValueType sets MetricValueType field to given value.

### HasMetricValueType

`func (o *MetricDescriptor) HasMetricValueType() bool`

HasMetricValueType returns a boolean if a field has been set.

### GetMaximumValue

`func (o *MetricDescriptor) GetMaximumValue() float64`

GetMaximumValue returns the MaximumValue field if non-nil, zero value otherwise.

### GetMaximumValueOk

`func (o *MetricDescriptor) GetMaximumValueOk() (*float64, bool)`

GetMaximumValueOk returns a tuple with the MaximumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValue

`func (o *MetricDescriptor) SetMaximumValue(v float64)`

SetMaximumValue sets MaximumValue field to given value.

### HasMaximumValue

`func (o *MetricDescriptor) HasMaximumValue() bool`

HasMaximumValue returns a boolean if a field has been set.

### GetLatency

`func (o *MetricDescriptor) GetLatency() int64`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *MetricDescriptor) GetLatencyOk() (*int64, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *MetricDescriptor) SetLatency(v int64)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *MetricDescriptor) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetImpactRelevant

`func (o *MetricDescriptor) GetImpactRelevant() bool`

GetImpactRelevant returns the ImpactRelevant field if non-nil, zero value otherwise.

### GetImpactRelevantOk

`func (o *MetricDescriptor) GetImpactRelevantOk() (*bool, bool)`

GetImpactRelevantOk returns a tuple with the ImpactRelevant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactRelevant

`func (o *MetricDescriptor) SetImpactRelevant(v bool)`

SetImpactRelevant sets ImpactRelevant field to given value.

### HasImpactRelevant

`func (o *MetricDescriptor) HasImpactRelevant() bool`

HasImpactRelevant returns a boolean if a field has been set.

### GetCreated

`func (o *MetricDescriptor) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MetricDescriptor) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MetricDescriptor) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *MetricDescriptor) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetWarnings

`func (o *MetricDescriptor) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *MetricDescriptor) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *MetricDescriptor) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *MetricDescriptor) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetDisplayName

`func (o *MetricDescriptor) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MetricDescriptor) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MetricDescriptor) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MetricDescriptor) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *MetricDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTransformations

`func (o *MetricDescriptor) GetTransformations() []string`

GetTransformations returns the Transformations field if non-nil, zero value otherwise.

### GetTransformationsOk

`func (o *MetricDescriptor) GetTransformationsOk() (*[]string, bool)`

GetTransformationsOk returns a tuple with the Transformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformations

`func (o *MetricDescriptor) SetTransformations(v []string)`

SetTransformations sets Transformations field to given value.

### HasTransformations

`func (o *MetricDescriptor) HasTransformations() bool`

HasTransformations returns a boolean if a field has been set.

### GetUnit

`func (o *MetricDescriptor) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricDescriptor) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricDescriptor) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricDescriptor) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetTags

`func (o *MetricDescriptor) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricDescriptor) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricDescriptor) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricDescriptor) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


