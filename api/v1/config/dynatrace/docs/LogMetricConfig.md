# LogMetricConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricKey** | **string** | The unique key of the metric.   The key must have the &#x60;calc:log.&#x60; prefix. | 
**Active** | Pointer to **bool** | The metric is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 
**DisplayName** | **string** | The name of the metric, displayed in the UI. | 
**Unit** | **string** | The unit of the metric. | 
**UnitDisplayName** | Pointer to **string** | The display name of the unit.    Only applicable if the **unit** is set to &#x60;UNSPECIFIED&#x60;. | [optional] 
**SearchString** | **string** | The pattern to look for in logs.    Use the [Dynatrace search query language](https://www.dynatrace.com/support/help/shortlink/log-viewer#search-for-text-patterns-in-log-files) to specify it. Quotes must be escaped.    To return all results, leave the field blank. | 
**MetricValueType** | **string** | The type of the metric data points calculation. For now the only allowed value is &#x60;OCCURRENCES&#x60;. | 
**ColumnDefiningValue** | Pointer to [**ColumnDefinition**](ColumnDefinition.md) |  | [optional] 
**LogSourceFilters** | [**[]LogSourceFilter**](LogSourceFilter.md) | A list of filters to define the logs to look into.    If several filters are specified, the OR logic applies. | 

## Methods

### NewLogMetricConfig

`func NewLogMetricConfig(metricKey string, displayName string, unit string, searchString string, metricValueType string, logSourceFilters []LogSourceFilter, ) *LogMetricConfig`

NewLogMetricConfig instantiates a new LogMetricConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogMetricConfigWithDefaults

`func NewLogMetricConfigWithDefaults() *LogMetricConfig`

NewLogMetricConfigWithDefaults instantiates a new LogMetricConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricKey

`func (o *LogMetricConfig) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *LogMetricConfig) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *LogMetricConfig) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetActive

`func (o *LogMetricConfig) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *LogMetricConfig) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *LogMetricConfig) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *LogMetricConfig) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDisplayName

`func (o *LogMetricConfig) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LogMetricConfig) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LogMetricConfig) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetUnit

`func (o *LogMetricConfig) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *LogMetricConfig) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *LogMetricConfig) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetUnitDisplayName

`func (o *LogMetricConfig) GetUnitDisplayName() string`

GetUnitDisplayName returns the UnitDisplayName field if non-nil, zero value otherwise.

### GetUnitDisplayNameOk

`func (o *LogMetricConfig) GetUnitDisplayNameOk() (*string, bool)`

GetUnitDisplayNameOk returns a tuple with the UnitDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitDisplayName

`func (o *LogMetricConfig) SetUnitDisplayName(v string)`

SetUnitDisplayName sets UnitDisplayName field to given value.

### HasUnitDisplayName

`func (o *LogMetricConfig) HasUnitDisplayName() bool`

HasUnitDisplayName returns a boolean if a field has been set.

### GetSearchString

`func (o *LogMetricConfig) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *LogMetricConfig) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *LogMetricConfig) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.


### GetMetricValueType

`func (o *LogMetricConfig) GetMetricValueType() string`

GetMetricValueType returns the MetricValueType field if non-nil, zero value otherwise.

### GetMetricValueTypeOk

`func (o *LogMetricConfig) GetMetricValueTypeOk() (*string, bool)`

GetMetricValueTypeOk returns a tuple with the MetricValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValueType

`func (o *LogMetricConfig) SetMetricValueType(v string)`

SetMetricValueType sets MetricValueType field to given value.


### GetColumnDefiningValue

`func (o *LogMetricConfig) GetColumnDefiningValue() ColumnDefinition`

GetColumnDefiningValue returns the ColumnDefiningValue field if non-nil, zero value otherwise.

### GetColumnDefiningValueOk

`func (o *LogMetricConfig) GetColumnDefiningValueOk() (*ColumnDefinition, bool)`

GetColumnDefiningValueOk returns a tuple with the ColumnDefiningValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDefiningValue

`func (o *LogMetricConfig) SetColumnDefiningValue(v ColumnDefinition)`

SetColumnDefiningValue sets ColumnDefiningValue field to given value.

### HasColumnDefiningValue

`func (o *LogMetricConfig) HasColumnDefiningValue() bool`

HasColumnDefiningValue returns a boolean if a field has been set.

### GetLogSourceFilters

`func (o *LogMetricConfig) GetLogSourceFilters() []LogSourceFilter`

GetLogSourceFilters returns the LogSourceFilters field if non-nil, zero value otherwise.

### GetLogSourceFiltersOk

`func (o *LogMetricConfig) GetLogSourceFiltersOk() (*[]LogSourceFilter, bool)`

GetLogSourceFiltersOk returns a tuple with the LogSourceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSourceFilters

`func (o *LogMetricConfig) SetLogSourceFilters(v []LogSourceFilter)`

SetLogSourceFilters sets LogSourceFilters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


