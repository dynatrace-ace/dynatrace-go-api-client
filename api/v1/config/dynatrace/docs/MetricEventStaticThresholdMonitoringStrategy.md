# MetricEventStaticThresholdMonitoringStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | **int32** | The number of one-minute samples that form the sliding evaluation window. | 
**ViolatingSamples** | **int32** | The number of one-minute samples within the evaluation window that must violate the threshold to trigger an event. | 
**DealertingSamples** | **int32** | The number of one-minute samples within the evaluation window that must go back to normal to close the event. | 
**AlertCondition** | **string** | The condition for the **threshold** value check: above or below. | 
**Threshold** | **float64** | The value of the static threshold based on the specified unit. | 
**Unit** | **string** | The unit of the threshold, matching the metric definition. | 

## Methods

### NewMetricEventStaticThresholdMonitoringStrategy

`func NewMetricEventStaticThresholdMonitoringStrategy(samples int32, violatingSamples int32, dealertingSamples int32, alertCondition string, threshold float64, unit string, ) *MetricEventStaticThresholdMonitoringStrategy`

NewMetricEventStaticThresholdMonitoringStrategy instantiates a new MetricEventStaticThresholdMonitoringStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricEventStaticThresholdMonitoringStrategyWithDefaults

`func NewMetricEventStaticThresholdMonitoringStrategyWithDefaults() *MetricEventStaticThresholdMonitoringStrategy`

NewMetricEventStaticThresholdMonitoringStrategyWithDefaults instantiates a new MetricEventStaticThresholdMonitoringStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamples

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetSamples() int32`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetSamplesOk() (*int32, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *MetricEventStaticThresholdMonitoringStrategy) SetSamples(v int32)`

SetSamples sets Samples field to given value.


### GetViolatingSamples

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetViolatingSamples() int32`

GetViolatingSamples returns the ViolatingSamples field if non-nil, zero value otherwise.

### GetViolatingSamplesOk

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetViolatingSamplesOk() (*int32, bool)`

GetViolatingSamplesOk returns a tuple with the ViolatingSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolatingSamples

`func (o *MetricEventStaticThresholdMonitoringStrategy) SetViolatingSamples(v int32)`

SetViolatingSamples sets ViolatingSamples field to given value.


### GetDealertingSamples

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetDealertingSamples() int32`

GetDealertingSamples returns the DealertingSamples field if non-nil, zero value otherwise.

### GetDealertingSamplesOk

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetDealertingSamplesOk() (*int32, bool)`

GetDealertingSamplesOk returns a tuple with the DealertingSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealertingSamples

`func (o *MetricEventStaticThresholdMonitoringStrategy) SetDealertingSamples(v int32)`

SetDealertingSamples sets DealertingSamples field to given value.


### GetAlertCondition

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetAlertCondition() string`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetAlertConditionOk() (*string, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *MetricEventStaticThresholdMonitoringStrategy) SetAlertCondition(v string)`

SetAlertCondition sets AlertCondition field to given value.


### GetThreshold

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *MetricEventStaticThresholdMonitoringStrategy) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.


### GetUnit

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricEventStaticThresholdMonitoringStrategy) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


