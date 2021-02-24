# CalculatedSyntheticMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorIdentifier** | **string** | The Dynatrace entity ID of the synthetic monitor to which the metric belongs. | 
**Name** | **string** | The name of the metric, displayed in the UI. | 
**MetricKey** | **string** | The unique key of the metric.    The key must have the &#x60;calc:synthetic&#x60; prefix. | 
**Enabled** | **bool** | The metric is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**Metric** | **string** | The type of the synthetic metric. | 
**Dimensions** | Pointer to [**[]SyntheticMetricDimension**](SyntheticMetricDimension.md) | A list of metric dimensions. | [optional] 
**Filter** | Pointer to [**SyntheticMetricFilter**](SyntheticMetricFilter.md) |  | [optional] 

## Methods

### NewCalculatedSyntheticMetric

`func NewCalculatedSyntheticMetric(monitorIdentifier string, name string, metricKey string, enabled bool, metric string, ) *CalculatedSyntheticMetric`

NewCalculatedSyntheticMetric instantiates a new CalculatedSyntheticMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalculatedSyntheticMetricWithDefaults

`func NewCalculatedSyntheticMetricWithDefaults() *CalculatedSyntheticMetric`

NewCalculatedSyntheticMetricWithDefaults instantiates a new CalculatedSyntheticMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorIdentifier

`func (o *CalculatedSyntheticMetric) GetMonitorIdentifier() string`

GetMonitorIdentifier returns the MonitorIdentifier field if non-nil, zero value otherwise.

### GetMonitorIdentifierOk

`func (o *CalculatedSyntheticMetric) GetMonitorIdentifierOk() (*string, bool)`

GetMonitorIdentifierOk returns a tuple with the MonitorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorIdentifier

`func (o *CalculatedSyntheticMetric) SetMonitorIdentifier(v string)`

SetMonitorIdentifier sets MonitorIdentifier field to given value.


### GetName

`func (o *CalculatedSyntheticMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CalculatedSyntheticMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CalculatedSyntheticMetric) SetName(v string)`

SetName sets Name field to given value.


### GetMetricKey

`func (o *CalculatedSyntheticMetric) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *CalculatedSyntheticMetric) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *CalculatedSyntheticMetric) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetEnabled

`func (o *CalculatedSyntheticMetric) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CalculatedSyntheticMetric) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CalculatedSyntheticMetric) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMetric

`func (o *CalculatedSyntheticMetric) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *CalculatedSyntheticMetric) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *CalculatedSyntheticMetric) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetDimensions

`func (o *CalculatedSyntheticMetric) GetDimensions() []SyntheticMetricDimension`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *CalculatedSyntheticMetric) GetDimensionsOk() (*[]SyntheticMetricDimension, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *CalculatedSyntheticMetric) SetDimensions(v []SyntheticMetricDimension)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *CalculatedSyntheticMetric) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetFilter

`func (o *CalculatedSyntheticMetric) GetFilter() SyntheticMetricFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CalculatedSyntheticMetric) GetFilterOk() (*SyntheticMetricFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CalculatedSyntheticMetric) SetFilter(v SyntheticMetricFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CalculatedSyntheticMetric) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


