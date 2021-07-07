# SloCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | The SLO is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 
**Name** | Pointer to **string** | The name of the SLO. | [optional] 
**CustomDescription** | Pointer to **string** | The custom description of the SLO (optional). | [optional] 
**MetricExpression** | Pointer to **string** | The percentage-based metric expression for the calculation of the SLO. This is not usable yet as it requires some server-side enablement to work. | [optional] 
**EvaluationType** | Pointer to **string** | The evaluation type of the SLO. | [optional] 
**Filter** | Pointer to **string** | The entity filter for the SLO evaluation. Use the [syntax of entity selector](https://dt-url.net/entityselector). | [optional] 
**Target** | Pointer to **float64** | The target value of the SLO. | [optional] 
**Warning** | Pointer to **float64** | The warning value of the SLO.    At warning state the SLO is still fulfilled but is getting close to failure. | [optional] 
**Timeframe** | Pointer to **string** | The timeframe for the SLO evaluation. Use the syntax of the global timeframe selector. | [optional] 
**UseRateMetric** | Pointer to **bool** | The type of the metric to use for SLO calculation:   * &#x60;true&#x60;: An existing percentage-based metric.  * &#x60;false&#x60;: A ratio of two metrics.   For a list of available metrics, see [Built-in metric page](https://dt-url.net/be03kow) or try the [GET metrics](https://dt-url.net/8e43kxf) API call. | [optional] 
**MetricRate** | Pointer to **string** | The percentage-based metric for the calculation of the SLO.   Required when the **useRateMetric** is set to &#x60;true&#x60;. | [optional] 
**MetricNumerator** | Pointer to **string** | The metric for the count of successes (the numerator in rate calculation).   Required when the **useRateMetric** is set to &#x60;false&#x60;. | [optional] 
**MetricDenominator** | Pointer to **string** | The total count metric (the denominator in rate calculation).   Required when the **useRateMetric** is set to &#x60;false&#x60;. | [optional] 

## Methods

### NewSloCreate

`func NewSloCreate() *SloCreate`

NewSloCreate instantiates a new SloCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloCreateWithDefaults

`func NewSloCreateWithDefaults() *SloCreate`

NewSloCreateWithDefaults instantiates a new SloCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SloCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SloCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SloCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SloCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *SloCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SloCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SloCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SloCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCustomDescription

`func (o *SloCreate) GetCustomDescription() string`

GetCustomDescription returns the CustomDescription field if non-nil, zero value otherwise.

### GetCustomDescriptionOk

`func (o *SloCreate) GetCustomDescriptionOk() (*string, bool)`

GetCustomDescriptionOk returns a tuple with the CustomDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDescription

`func (o *SloCreate) SetCustomDescription(v string)`

SetCustomDescription sets CustomDescription field to given value.

### HasCustomDescription

`func (o *SloCreate) HasCustomDescription() bool`

HasCustomDescription returns a boolean if a field has been set.

### GetMetricExpression

`func (o *SloCreate) GetMetricExpression() string`

GetMetricExpression returns the MetricExpression field if non-nil, zero value otherwise.

### GetMetricExpressionOk

`func (o *SloCreate) GetMetricExpressionOk() (*string, bool)`

GetMetricExpressionOk returns a tuple with the MetricExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricExpression

`func (o *SloCreate) SetMetricExpression(v string)`

SetMetricExpression sets MetricExpression field to given value.

### HasMetricExpression

`func (o *SloCreate) HasMetricExpression() bool`

HasMetricExpression returns a boolean if a field has been set.

### GetEvaluationType

`func (o *SloCreate) GetEvaluationType() string`

GetEvaluationType returns the EvaluationType field if non-nil, zero value otherwise.

### GetEvaluationTypeOk

`func (o *SloCreate) GetEvaluationTypeOk() (*string, bool)`

GetEvaluationTypeOk returns a tuple with the EvaluationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationType

`func (o *SloCreate) SetEvaluationType(v string)`

SetEvaluationType sets EvaluationType field to given value.

### HasEvaluationType

`func (o *SloCreate) HasEvaluationType() bool`

HasEvaluationType returns a boolean if a field has been set.

### GetFilter

`func (o *SloCreate) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SloCreate) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SloCreate) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *SloCreate) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetTarget

`func (o *SloCreate) GetTarget() float64`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SloCreate) GetTargetOk() (*float64, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SloCreate) SetTarget(v float64)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SloCreate) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetWarning

`func (o *SloCreate) GetWarning() float64`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *SloCreate) GetWarningOk() (*float64, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *SloCreate) SetWarning(v float64)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *SloCreate) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetTimeframe

`func (o *SloCreate) GetTimeframe() string`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *SloCreate) GetTimeframeOk() (*string, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *SloCreate) SetTimeframe(v string)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *SloCreate) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### GetUseRateMetric

`func (o *SloCreate) GetUseRateMetric() bool`

GetUseRateMetric returns the UseRateMetric field if non-nil, zero value otherwise.

### GetUseRateMetricOk

`func (o *SloCreate) GetUseRateMetricOk() (*bool, bool)`

GetUseRateMetricOk returns a tuple with the UseRateMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRateMetric

`func (o *SloCreate) SetUseRateMetric(v bool)`

SetUseRateMetric sets UseRateMetric field to given value.

### HasUseRateMetric

`func (o *SloCreate) HasUseRateMetric() bool`

HasUseRateMetric returns a boolean if a field has been set.

### GetMetricRate

`func (o *SloCreate) GetMetricRate() string`

GetMetricRate returns the MetricRate field if non-nil, zero value otherwise.

### GetMetricRateOk

`func (o *SloCreate) GetMetricRateOk() (*string, bool)`

GetMetricRateOk returns a tuple with the MetricRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricRate

`func (o *SloCreate) SetMetricRate(v string)`

SetMetricRate sets MetricRate field to given value.

### HasMetricRate

`func (o *SloCreate) HasMetricRate() bool`

HasMetricRate returns a boolean if a field has been set.

### GetMetricNumerator

`func (o *SloCreate) GetMetricNumerator() string`

GetMetricNumerator returns the MetricNumerator field if non-nil, zero value otherwise.

### GetMetricNumeratorOk

`func (o *SloCreate) GetMetricNumeratorOk() (*string, bool)`

GetMetricNumeratorOk returns a tuple with the MetricNumerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricNumerator

`func (o *SloCreate) SetMetricNumerator(v string)`

SetMetricNumerator sets MetricNumerator field to given value.

### HasMetricNumerator

`func (o *SloCreate) HasMetricNumerator() bool`

HasMetricNumerator returns a boolean if a field has been set.

### GetMetricDenominator

`func (o *SloCreate) GetMetricDenominator() string`

GetMetricDenominator returns the MetricDenominator field if non-nil, zero value otherwise.

### GetMetricDenominatorOk

`func (o *SloCreate) GetMetricDenominatorOk() (*string, bool)`

GetMetricDenominatorOk returns a tuple with the MetricDenominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDenominator

`func (o *SloCreate) SetMetricDenominator(v string)`

SetMetricDenominator sets MetricDenominator field to given value.

### HasMetricDenominator

`func (o *SloCreate) HasMetricDenominator() bool`

HasMetricDenominator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


