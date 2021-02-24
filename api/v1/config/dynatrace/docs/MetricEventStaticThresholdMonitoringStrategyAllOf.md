# MetricEventStaticThresholdMonitoringStrategyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | Pointer to **int32** | The number of one-minute samples that form the sliding evaluation window. | [optional] 
**ViolatingSamples** | Pointer to **int32** | The number of one-minute samples within the evaluation window that must violate the threshold to trigger an event. | [optional] 
**DealertingSamples** | Pointer to **int32** | The number of one-minute samples within the evaluation window that must go back to normal to close the event. | [optional] 
**AlertCondition** | Pointer to **string** | The condition for the **threshold** value check: above or below. | [optional] 
**Threshold** | Pointer to **float64** | The value of the static threshold based on the specified unit. | [optional] 
**Unit** | Pointer to **string** | The unit of the threshold, matching the metric definition. | [optional] 

## Methods

### NewMetricEventStaticThresholdMonitoringStrategyAllOf

`func NewMetricEventStaticThresholdMonitoringStrategyAllOf() *MetricEventStaticThresholdMonitoringStrategyAllOf`

NewMetricEventStaticThresholdMonitoringStrategyAllOf instantiates a new MetricEventStaticThresholdMonitoringStrategyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricEventStaticThresholdMonitoringStrategyAllOfWithDefaults

`func NewMetricEventStaticThresholdMonitoringStrategyAllOfWithDefaults() *MetricEventStaticThresholdMonitoringStrategyAllOf`

NewMetricEventStaticThresholdMonitoringStrategyAllOfWithDefaults instantiates a new MetricEventStaticThresholdMonitoringStrategyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamples

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetSamples() int32`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetSamplesOk() (*int32, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) SetSamples(v int32)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetViolatingSamples

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetViolatingSamples() int32`

GetViolatingSamples returns the ViolatingSamples field if non-nil, zero value otherwise.

### GetViolatingSamplesOk

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetViolatingSamplesOk() (*int32, bool)`

GetViolatingSamplesOk returns a tuple with the ViolatingSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolatingSamples

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) SetViolatingSamples(v int32)`

SetViolatingSamples sets ViolatingSamples field to given value.

### HasViolatingSamples

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) HasViolatingSamples() bool`

HasViolatingSamples returns a boolean if a field has been set.

### GetDealertingSamples

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetDealertingSamples() int32`

GetDealertingSamples returns the DealertingSamples field if non-nil, zero value otherwise.

### GetDealertingSamplesOk

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetDealertingSamplesOk() (*int32, bool)`

GetDealertingSamplesOk returns a tuple with the DealertingSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealertingSamples

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) SetDealertingSamples(v int32)`

SetDealertingSamples sets DealertingSamples field to given value.

### HasDealertingSamples

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) HasDealertingSamples() bool`

HasDealertingSamples returns a boolean if a field has been set.

### GetAlertCondition

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetAlertCondition() string`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetAlertConditionOk() (*string, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) SetAlertCondition(v string)`

SetAlertCondition sets AlertCondition field to given value.

### HasAlertCondition

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) HasAlertCondition() bool`

HasAlertCondition returns a boolean if a field has been set.

### GetThreshold

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUnit

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


