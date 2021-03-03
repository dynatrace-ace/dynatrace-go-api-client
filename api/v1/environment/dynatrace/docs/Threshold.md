# Threshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThresholdId** | Pointer to **string** | The ID of the threshold. | [optional] 
**TimeseriesId** | Pointer to **string** | The case-sensitive ID of the metric evaluated by threshold.   You can use it to find the metric via the &#x60;timeseries&#x60; endpoint of the API. | [optional] 
**Threshold** | Pointer to **float64** | The value of the threshold. | [optional] 
**AlertCondition** | Pointer to **string** | The condition for the threshold value check: above or below. | [optional] 
**Samples** | Pointer to **int32** | The number of one-minute samples that form sliding evaluation window. | [optional] 
**ViolatingSamples** | Pointer to **int32** | How many one-minute samples within the evaluation window should violate the threshold to trigger an event. | [optional] 
**EventType** | Pointer to **string** | The type of the event to trigger on the threshold violation. | [optional] 
**EventName** | Pointer to **string** | The name of the event to trigger on the threshold violation. | [optional] 
**Filter** | Pointer to **string** | The source of the threshold. | [optional] 
**Enabled** | Pointer to **bool** | The threshold is enabled/disabled. | [optional] 
**DealertingSamples** | Pointer to **int32** | How many one-minute samples within the evaluation window should go back to normal to close the event. | [optional] 
**Description** | Pointer to **string** | A description of the event, triggered by a threshold violation.    You can use the following placeholders:  {severity}: the actual metric value,  {alert_condition}: above or below threshold condition,  {threshold}: the threshold value,{violating_samples}: the number of samples, violating the threshold,  {dimensions}: metric&#39;s dimension that violated the threshold. | [optional] 

## Methods

### NewThreshold

`func NewThreshold() *Threshold`

NewThreshold instantiates a new Threshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdWithDefaults

`func NewThresholdWithDefaults() *Threshold`

NewThresholdWithDefaults instantiates a new Threshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThresholdId

`func (o *Threshold) GetThresholdId() string`

GetThresholdId returns the ThresholdId field if non-nil, zero value otherwise.

### GetThresholdIdOk

`func (o *Threshold) GetThresholdIdOk() (*string, bool)`

GetThresholdIdOk returns a tuple with the ThresholdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdId

`func (o *Threshold) SetThresholdId(v string)`

SetThresholdId sets ThresholdId field to given value.

### HasThresholdId

`func (o *Threshold) HasThresholdId() bool`

HasThresholdId returns a boolean if a field has been set.

### GetTimeseriesId

`func (o *Threshold) GetTimeseriesId() string`

GetTimeseriesId returns the TimeseriesId field if non-nil, zero value otherwise.

### GetTimeseriesIdOk

`func (o *Threshold) GetTimeseriesIdOk() (*string, bool)`

GetTimeseriesIdOk returns a tuple with the TimeseriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseriesId

`func (o *Threshold) SetTimeseriesId(v string)`

SetTimeseriesId sets TimeseriesId field to given value.

### HasTimeseriesId

`func (o *Threshold) HasTimeseriesId() bool`

HasTimeseriesId returns a boolean if a field has been set.

### GetThreshold

`func (o *Threshold) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Threshold) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *Threshold) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *Threshold) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetAlertCondition

`func (o *Threshold) GetAlertCondition() string`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *Threshold) GetAlertConditionOk() (*string, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *Threshold) SetAlertCondition(v string)`

SetAlertCondition sets AlertCondition field to given value.

### HasAlertCondition

`func (o *Threshold) HasAlertCondition() bool`

HasAlertCondition returns a boolean if a field has been set.

### GetSamples

`func (o *Threshold) GetSamples() int32`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *Threshold) GetSamplesOk() (*int32, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *Threshold) SetSamples(v int32)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *Threshold) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetViolatingSamples

`func (o *Threshold) GetViolatingSamples() int32`

GetViolatingSamples returns the ViolatingSamples field if non-nil, zero value otherwise.

### GetViolatingSamplesOk

`func (o *Threshold) GetViolatingSamplesOk() (*int32, bool)`

GetViolatingSamplesOk returns a tuple with the ViolatingSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolatingSamples

`func (o *Threshold) SetViolatingSamples(v int32)`

SetViolatingSamples sets ViolatingSamples field to given value.

### HasViolatingSamples

`func (o *Threshold) HasViolatingSamples() bool`

HasViolatingSamples returns a boolean if a field has been set.

### GetEventType

`func (o *Threshold) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Threshold) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Threshold) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *Threshold) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventName

`func (o *Threshold) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *Threshold) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *Threshold) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *Threshold) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetFilter

`func (o *Threshold) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Threshold) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Threshold) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Threshold) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetEnabled

`func (o *Threshold) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Threshold) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Threshold) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Threshold) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDealertingSamples

`func (o *Threshold) GetDealertingSamples() int32`

GetDealertingSamples returns the DealertingSamples field if non-nil, zero value otherwise.

### GetDealertingSamplesOk

`func (o *Threshold) GetDealertingSamplesOk() (*int32, bool)`

GetDealertingSamplesOk returns a tuple with the DealertingSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealertingSamples

`func (o *Threshold) SetDealertingSamples(v int32)`

SetDealertingSamples sets DealertingSamples field to given value.

### HasDealertingSamples

`func (o *Threshold) HasDealertingSamples() bool`

HasDealertingSamples returns a boolean if a field has been set.

### GetDescription

`func (o *Threshold) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Threshold) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Threshold) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Threshold) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


