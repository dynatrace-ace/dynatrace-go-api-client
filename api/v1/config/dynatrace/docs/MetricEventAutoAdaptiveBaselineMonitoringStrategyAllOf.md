# MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | Pointer to **int32** | The number of one-minute samples that form the sliding evaluation window. | [optional] 
**ViolatingSamples** | Pointer to **int32** | The number of one-minute samples within the evaluation window that must violate the threshold to trigger an event. | [optional] 
**DealertingSamples** | Pointer to **int32** | The number of one-minute samples within the evaluation window that must go back to normal to close the event. | [optional] 
**AlertCondition** | Pointer to **string** | The condition for the **threshold** value check: above or below. | [optional] 
**NumberOfSignalFluctuations** | Pointer to **float64** | Defines the factor of how many signal fluctuations are valid. Values above the baseline plus the signal fluctuation times the number of tolerated signal fluctuations are alerted. | [optional] 

## Methods

### NewMetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf

`func NewMetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf() *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf`

NewMetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf instantiates a new MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricEventAutoAdaptiveBaselineMonitoringStrategyAllOfWithDefaults

`func NewMetricEventAutoAdaptiveBaselineMonitoringStrategyAllOfWithDefaults() *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf`

NewMetricEventAutoAdaptiveBaselineMonitoringStrategyAllOfWithDefaults instantiates a new MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamples

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) GetSamples() int32`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) GetSamplesOk() (*int32, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) SetSamples(v int32)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetViolatingSamples

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) GetViolatingSamples() int32`

GetViolatingSamples returns the ViolatingSamples field if non-nil, zero value otherwise.

### GetViolatingSamplesOk

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) GetViolatingSamplesOk() (*int32, bool)`

GetViolatingSamplesOk returns a tuple with the ViolatingSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolatingSamples

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) SetViolatingSamples(v int32)`

SetViolatingSamples sets ViolatingSamples field to given value.

### HasViolatingSamples

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) HasViolatingSamples() bool`

HasViolatingSamples returns a boolean if a field has been set.

### GetDealertingSamples

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) GetDealertingSamples() int32`

GetDealertingSamples returns the DealertingSamples field if non-nil, zero value otherwise.

### GetDealertingSamplesOk

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) GetDealertingSamplesOk() (*int32, bool)`

GetDealertingSamplesOk returns a tuple with the DealertingSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealertingSamples

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) SetDealertingSamples(v int32)`

SetDealertingSamples sets DealertingSamples field to given value.

### HasDealertingSamples

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) HasDealertingSamples() bool`

HasDealertingSamples returns a boolean if a field has been set.

### GetAlertCondition

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) GetAlertCondition() string`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) GetAlertConditionOk() (*string, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) SetAlertCondition(v string)`

SetAlertCondition sets AlertCondition field to given value.

### HasAlertCondition

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) HasAlertCondition() bool`

HasAlertCondition returns a boolean if a field has been set.

### GetNumberOfSignalFluctuations

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) GetNumberOfSignalFluctuations() float64`

GetNumberOfSignalFluctuations returns the NumberOfSignalFluctuations field if non-nil, zero value otherwise.

### GetNumberOfSignalFluctuationsOk

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) GetNumberOfSignalFluctuationsOk() (*float64, bool)`

GetNumberOfSignalFluctuationsOk returns a tuple with the NumberOfSignalFluctuations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSignalFluctuations

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) SetNumberOfSignalFluctuations(v float64)`

SetNumberOfSignalFluctuations sets NumberOfSignalFluctuations field to given value.

### HasNumberOfSignalFluctuations

`func (o *MetricEventAutoAdaptiveBaselineMonitoringStrategyAllOf) HasNumberOfSignalFluctuations() bool`

HasNumberOfSignalFluctuations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


