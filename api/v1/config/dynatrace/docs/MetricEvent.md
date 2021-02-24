# MetricEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the metric event. | [optional] 
**MetricId** | **string** | The ID of the metric evaluated by the metric event. | 
**Name** | **string** | The name of the metric event displayed in the UI. | 
**Description** | **string** | The description of the metric event. | 
**AggregationType** | Pointer to **string** | How the metric data points are aggregated for the evaluation.    The timeseries must support this aggregation. | [optional] 
**Severity** | Pointer to **string** | The type of the event to trigger on the threshold violation.   The &#x60;CUSTOM_ALERT&#x60; type is not correlated with other alerts. The &#x60;INFO&#x60; type does not open a problem. | [optional] 
**Enabled** | **bool** | The metric event is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**DisabledReason** | Pointer to **string** | The reason of automatic disabling.   The &#x60;NONE&#x60; means config was not disabled automatically. | [optional] [readonly] 
**WarningReason** | Pointer to **string** | The reason of a warning set on the config.   The &#x60;NONE&#x60; means config has no warnings. | [optional] [readonly] 
**AlertingScope** | Pointer to [**[]MetricEventAlertingScope**](MetricEventAlertingScope.md) | Defines the scope of the metric event. Only one filter is allowed per filter type, except for tags, where up to 3 are allowed. The filters are combined by conjunction. | [optional] 
**MetricDimensions** | Pointer to [**[]MetricEventDimensions**](MetricEventDimensions.md) | Defines the dimensions of the metric to alert on. The filters are combined by conjunction. | [optional] 
**MonitoringStrategy** | [**MetricEventMonitoringStrategy**](MetricEventMonitoringStrategy.md) |  | 
**PrimaryDimensionKey** | Pointer to **string** | Defines which dimension key should be used for the **alertingScope**. | [optional] 

## Methods

### NewMetricEvent

`func NewMetricEvent(metricId string, name string, description string, enabled bool, monitoringStrategy MetricEventMonitoringStrategy, ) *MetricEvent`

NewMetricEvent instantiates a new MetricEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricEventWithDefaults

`func NewMetricEventWithDefaults() *MetricEvent`

NewMetricEventWithDefaults instantiates a new MetricEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *MetricEvent) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MetricEvent) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MetricEvent) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MetricEvent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *MetricEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetricEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetricId

`func (o *MetricEvent) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *MetricEvent) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *MetricEvent) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.


### GetName

`func (o *MetricEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricEvent) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MetricEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricEvent) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAggregationType

`func (o *MetricEvent) GetAggregationType() string`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *MetricEvent) GetAggregationTypeOk() (*string, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *MetricEvent) SetAggregationType(v string)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *MetricEvent) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetSeverity

`func (o *MetricEvent) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *MetricEvent) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *MetricEvent) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *MetricEvent) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetEnabled

`func (o *MetricEvent) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MetricEvent) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MetricEvent) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDisabledReason

`func (o *MetricEvent) GetDisabledReason() string`

GetDisabledReason returns the DisabledReason field if non-nil, zero value otherwise.

### GetDisabledReasonOk

`func (o *MetricEvent) GetDisabledReasonOk() (*string, bool)`

GetDisabledReasonOk returns a tuple with the DisabledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledReason

`func (o *MetricEvent) SetDisabledReason(v string)`

SetDisabledReason sets DisabledReason field to given value.

### HasDisabledReason

`func (o *MetricEvent) HasDisabledReason() bool`

HasDisabledReason returns a boolean if a field has been set.

### GetWarningReason

`func (o *MetricEvent) GetWarningReason() string`

GetWarningReason returns the WarningReason field if non-nil, zero value otherwise.

### GetWarningReasonOk

`func (o *MetricEvent) GetWarningReasonOk() (*string, bool)`

GetWarningReasonOk returns a tuple with the WarningReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningReason

`func (o *MetricEvent) SetWarningReason(v string)`

SetWarningReason sets WarningReason field to given value.

### HasWarningReason

`func (o *MetricEvent) HasWarningReason() bool`

HasWarningReason returns a boolean if a field has been set.

### GetAlertingScope

`func (o *MetricEvent) GetAlertingScope() []MetricEventAlertingScope`

GetAlertingScope returns the AlertingScope field if non-nil, zero value otherwise.

### GetAlertingScopeOk

`func (o *MetricEvent) GetAlertingScopeOk() (*[]MetricEventAlertingScope, bool)`

GetAlertingScopeOk returns a tuple with the AlertingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingScope

`func (o *MetricEvent) SetAlertingScope(v []MetricEventAlertingScope)`

SetAlertingScope sets AlertingScope field to given value.

### HasAlertingScope

`func (o *MetricEvent) HasAlertingScope() bool`

HasAlertingScope returns a boolean if a field has been set.

### GetMetricDimensions

`func (o *MetricEvent) GetMetricDimensions() []MetricEventDimensions`

GetMetricDimensions returns the MetricDimensions field if non-nil, zero value otherwise.

### GetMetricDimensionsOk

`func (o *MetricEvent) GetMetricDimensionsOk() (*[]MetricEventDimensions, bool)`

GetMetricDimensionsOk returns a tuple with the MetricDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDimensions

`func (o *MetricEvent) SetMetricDimensions(v []MetricEventDimensions)`

SetMetricDimensions sets MetricDimensions field to given value.

### HasMetricDimensions

`func (o *MetricEvent) HasMetricDimensions() bool`

HasMetricDimensions returns a boolean if a field has been set.

### GetMonitoringStrategy

`func (o *MetricEvent) GetMonitoringStrategy() MetricEventMonitoringStrategy`

GetMonitoringStrategy returns the MonitoringStrategy field if non-nil, zero value otherwise.

### GetMonitoringStrategyOk

`func (o *MetricEvent) GetMonitoringStrategyOk() (*MetricEventMonitoringStrategy, bool)`

GetMonitoringStrategyOk returns a tuple with the MonitoringStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringStrategy

`func (o *MetricEvent) SetMonitoringStrategy(v MetricEventMonitoringStrategy)`

SetMonitoringStrategy sets MonitoringStrategy field to given value.


### GetPrimaryDimensionKey

`func (o *MetricEvent) GetPrimaryDimensionKey() string`

GetPrimaryDimensionKey returns the PrimaryDimensionKey field if non-nil, zero value otherwise.

### GetPrimaryDimensionKeyOk

`func (o *MetricEvent) GetPrimaryDimensionKeyOk() (*string, bool)`

GetPrimaryDimensionKeyOk returns a tuple with the PrimaryDimensionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDimensionKey

`func (o *MetricEvent) SetPrimaryDimensionKey(v string)`

SetPrimaryDimensionKey sets PrimaryDimensionKey field to given value.

### HasPrimaryDimensionKey

`func (o *MetricEvent) HasPrimaryDimensionKey() bool`

HasPrimaryDimensionKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


