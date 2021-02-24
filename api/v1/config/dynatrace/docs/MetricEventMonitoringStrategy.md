# MetricEventMonitoringStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;STATIC_THRESHOLD&#x60; -&gt; MetricEventStaticThresholdMonitoringStrategy  * &#x60;AUTO_ADAPTIVE_BASELINE&#x60; -&gt; MetricEventAutoAdaptiveBaselineMonitoringStrategy   | 

## Methods

### NewMetricEventMonitoringStrategy

`func NewMetricEventMonitoringStrategy(type_ string, ) *MetricEventMonitoringStrategy`

NewMetricEventMonitoringStrategy instantiates a new MetricEventMonitoringStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricEventMonitoringStrategyWithDefaults

`func NewMetricEventMonitoringStrategyWithDefaults() *MetricEventMonitoringStrategy`

NewMetricEventMonitoringStrategyWithDefaults instantiates a new MetricEventMonitoringStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetricEventMonitoringStrategy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricEventMonitoringStrategy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricEventMonitoringStrategy) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


