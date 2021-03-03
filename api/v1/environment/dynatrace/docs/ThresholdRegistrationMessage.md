# ThresholdRegistrationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThresholdId** | Pointer to **string** | The ID of the threshold. | [optional] 
**TimeseriesId** | Pointer to **string** | The case-sensitive ID of the metric evaluated by threshold.    You can find metric IDs, by performing the GET request to the &#x60;timeseries&#x60; endpoint of the API. | [optional] 
**Threshold** | Pointer to **float64** | The value of the threshold. | [optional] 
**AlertCondition** | Pointer to **string** | The condition for the threshold value check: above or below. | [optional] 
**Samples** | Pointer to **int32** | The number of one-minute samples to form the sliding evaluation window. | [optional] 
**ViolatingSamples** | Pointer to **int32** | How many one-minute samples within the evaluation window should violate the threshold to trigger an event. | [optional] 
**DealertingSamples** | Pointer to **int32** | How many one-minute samples within the evaluation window should go back to normal to close the event. | [optional] 
**EventType** | Pointer to **string** | The type of the event to trigger on the threshold violation. | [optional] 
**EventName** | Pointer to **string** | The name of the event, displayed in the UI. | [optional] 
**Description** | Pointer to **string** | A description of the event, triggered by a threshold violation.    You can use the following placeholders:  {severity}: the actual metric value,  {alert_condition}: above or below threshold condition,  {threshold}: the threshold value,{violating_samples}: the number of samples, violating the threshold,  {dimensions}: metric&#39;s dimension that violated the threshold. | [optional] 
**Enabled** | Pointer to **bool** | The threshold is enabled/disabled. | [optional] 

## Methods

### NewThresholdRegistrationMessage

`func NewThresholdRegistrationMessage() *ThresholdRegistrationMessage`

NewThresholdRegistrationMessage instantiates a new ThresholdRegistrationMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdRegistrationMessageWithDefaults

`func NewThresholdRegistrationMessageWithDefaults() *ThresholdRegistrationMessage`

NewThresholdRegistrationMessageWithDefaults instantiates a new ThresholdRegistrationMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThresholdId

`func (o *ThresholdRegistrationMessage) GetThresholdId() string`

GetThresholdId returns the ThresholdId field if non-nil, zero value otherwise.

### GetThresholdIdOk

`func (o *ThresholdRegistrationMessage) GetThresholdIdOk() (*string, bool)`

GetThresholdIdOk returns a tuple with the ThresholdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdId

`func (o *ThresholdRegistrationMessage) SetThresholdId(v string)`

SetThresholdId sets ThresholdId field to given value.

### HasThresholdId

`func (o *ThresholdRegistrationMessage) HasThresholdId() bool`

HasThresholdId returns a boolean if a field has been set.

### GetTimeseriesId

`func (o *ThresholdRegistrationMessage) GetTimeseriesId() string`

GetTimeseriesId returns the TimeseriesId field if non-nil, zero value otherwise.

### GetTimeseriesIdOk

`func (o *ThresholdRegistrationMessage) GetTimeseriesIdOk() (*string, bool)`

GetTimeseriesIdOk returns a tuple with the TimeseriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseriesId

`func (o *ThresholdRegistrationMessage) SetTimeseriesId(v string)`

SetTimeseriesId sets TimeseriesId field to given value.

### HasTimeseriesId

`func (o *ThresholdRegistrationMessage) HasTimeseriesId() bool`

HasTimeseriesId returns a boolean if a field has been set.

### GetThreshold

`func (o *ThresholdRegistrationMessage) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ThresholdRegistrationMessage) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ThresholdRegistrationMessage) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ThresholdRegistrationMessage) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetAlertCondition

`func (o *ThresholdRegistrationMessage) GetAlertCondition() string`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *ThresholdRegistrationMessage) GetAlertConditionOk() (*string, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *ThresholdRegistrationMessage) SetAlertCondition(v string)`

SetAlertCondition sets AlertCondition field to given value.

### HasAlertCondition

`func (o *ThresholdRegistrationMessage) HasAlertCondition() bool`

HasAlertCondition returns a boolean if a field has been set.

### GetSamples

`func (o *ThresholdRegistrationMessage) GetSamples() int32`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *ThresholdRegistrationMessage) GetSamplesOk() (*int32, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *ThresholdRegistrationMessage) SetSamples(v int32)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *ThresholdRegistrationMessage) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetViolatingSamples

`func (o *ThresholdRegistrationMessage) GetViolatingSamples() int32`

GetViolatingSamples returns the ViolatingSamples field if non-nil, zero value otherwise.

### GetViolatingSamplesOk

`func (o *ThresholdRegistrationMessage) GetViolatingSamplesOk() (*int32, bool)`

GetViolatingSamplesOk returns a tuple with the ViolatingSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolatingSamples

`func (o *ThresholdRegistrationMessage) SetViolatingSamples(v int32)`

SetViolatingSamples sets ViolatingSamples field to given value.

### HasViolatingSamples

`func (o *ThresholdRegistrationMessage) HasViolatingSamples() bool`

HasViolatingSamples returns a boolean if a field has been set.

### GetDealertingSamples

`func (o *ThresholdRegistrationMessage) GetDealertingSamples() int32`

GetDealertingSamples returns the DealertingSamples field if non-nil, zero value otherwise.

### GetDealertingSamplesOk

`func (o *ThresholdRegistrationMessage) GetDealertingSamplesOk() (*int32, bool)`

GetDealertingSamplesOk returns a tuple with the DealertingSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealertingSamples

`func (o *ThresholdRegistrationMessage) SetDealertingSamples(v int32)`

SetDealertingSamples sets DealertingSamples field to given value.

### HasDealertingSamples

`func (o *ThresholdRegistrationMessage) HasDealertingSamples() bool`

HasDealertingSamples returns a boolean if a field has been set.

### GetEventType

`func (o *ThresholdRegistrationMessage) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ThresholdRegistrationMessage) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ThresholdRegistrationMessage) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ThresholdRegistrationMessage) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventName

`func (o *ThresholdRegistrationMessage) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *ThresholdRegistrationMessage) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *ThresholdRegistrationMessage) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *ThresholdRegistrationMessage) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetDescription

`func (o *ThresholdRegistrationMessage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThresholdRegistrationMessage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThresholdRegistrationMessage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThresholdRegistrationMessage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ThresholdRegistrationMessage) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThresholdRegistrationMessage) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThresholdRegistrationMessage) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ThresholdRegistrationMessage) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


