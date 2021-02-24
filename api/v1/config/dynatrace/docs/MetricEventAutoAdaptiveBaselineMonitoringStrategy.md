# MetricEventAutoAdaptiveBaselineMonitoringStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | **int32** | The number of one-minute samples that form the sliding evaluation window. | 
**ViolatingSamples** | **int32** | The number of one-minute samples within the evaluation window that must violate the threshold to trigger an event. | 
**DealertingSamples** | **int32** | The number of one-minute samples within the evaluation window that must go back to normal to close the event. | 
**AlertCondition** | **string** | The condition for the **threshold** value check: above or below. | 
**NumberOfSignalFluctuations** | **float64** | Defines the factor of how many signal fluctuations are valid. Values above the baseline plus the signal fluctuation times the number of tolerated signal fluctuations are alerted. | 

## Methods

### NewMetricEventAutoAdaptiveBaselineMonitoringStrategy

`func NewMetricEventAutoAdaptiveBaselineMonitoringStrategy(samples int32, violatingSamples int32, dealertingSamples int32, alertCondition string, numberOfSignalFluctuations float64, ) *MetricEventAutoAdaptiveBaselineMonitoringStrategy`

NewMetricEventAutoAdaptiveBaselineMonitoringStrategy instantiates a new MetricEventAutoAdaptiveBaselineMonitoringStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricEventAutoAdaptiveBaselineMonitoringStrategyWithDefaults

`func NewMetricEventAutoAdaptiveBaselineMonitoringStrategyWithDefaults() *MetricEventAutoAdaptiveBaselineMonitoringStrategy`

NewMetricEventAutoAdaptiveBaselineMonitoringStrategyWithDefaults instantiates a new MetricEventAutoAdaptiveBaselineMonitoringStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamples

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategy) GetSamples() int32`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategy) GetSamplesOk() (*int32, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategy) SetSamples(v int32)`

SetSamples sets Samples field to given value.


### GetViolatingSamples

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategy) GetViolatingSamples() int32`

GetViolatingSamples returns the ViolatingSamples field if non-nil, zero value otherwise.

### GetViolatingSamplesOk

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategy) GetViolatingSamplesOk() (*int32, bool)`

GetViolatingSamplesOk returns a tuple with the ViolatingSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolatingSamples

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategy) SetViolatingSamples(v int32)`

SetViolatingSamples sets ViolatingSamples field to given value.


### GetDealertingSamples

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategy) GetDealertingSamples() int32`

GetDealertingSamples returns the DealertingSamples field if non-nil, zero value otherwise.

### GetDealertingSamplesOk

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategy) GetDealertingSamplesOk() (*int32, bool)`

GetDealertingSamplesOk returns a tuple with the DealertingSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealertingSamples

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategy) SetDealertingSamples(v int32)`

SetDealertingSamples sets DealertingSamples field to given value.


### GetAlertCondition

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategy) GetAlertCondition() string`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategy) GetAlertConditionOk() (*string, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategy) SetAlertCondition(v string)`

SetAlertCondition sets AlertCondition field to given value.


### GetNumberOfSignalFluctuations

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategy) GetNumberOfSignalFluctuations() float64`

GetNumberOfSignalFluctuations returns the NumberOfSignalFluctuations field if non-nil, zero value otherwise.

### GetNumberOfSignalFluctuationsOk

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategy) GetNumberOfSignalFluctuationsOk() (*float64, bool)`

GetNumberOfSignalFluctuationsOk returns a tuple with the NumberOfSignalFluctuations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSignalFluctuations

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategy) SetNumberOfSignalFluctuations(v float64)`

SetNumberOfSignalFluctuations sets NumberOfSignalFluctuations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


