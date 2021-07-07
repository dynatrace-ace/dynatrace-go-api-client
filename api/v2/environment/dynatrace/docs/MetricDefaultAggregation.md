# MetricDefaultAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of default aggregation. | [optional] 
**Parameter** | Pointer to **float64** | The percentile to be delivered. Valid values are between &#x60;0&#x60; and &#x60;100&#x60;.   Applicable only to the &#x60;percentile&#x60; aggregation type. | [optional] 

## Methods

### NewMetricDefaultAggregation

`func NewMetricDefaultAggregation() *MetricDefaultAggregation`

NewMetricDefaultAggregation instantiates a new MetricDefaultAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDefaultAggregationWithDefaults

`func NewMetricDefaultAggregationWithDefaults() *MetricDefaultAggregation`

NewMetricDefaultAggregationWithDefaults instantiates a new MetricDefaultAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetricDefaultAggregation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricDefaultAggregation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricDefaultAggregation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetricDefaultAggregation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetParameter

`func (o *MetricDefaultAggregation) GetParameter() float64`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *MetricDefaultAggregation) GetParameterOk() (*float64, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *MetricDefaultAggregation) SetParameter(v float64)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *MetricDefaultAggregation) HasParameter() bool`

HasParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


