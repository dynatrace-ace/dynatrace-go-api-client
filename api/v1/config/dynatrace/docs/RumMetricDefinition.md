# RumMetricDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | The type of the RUM metric. | 
**PropertyKey** | Pointer to **string** | The key of the user action property.    Only applicable for &#x60;DoubleProperty&#x60; and &#x60;LongProperty&#x60; metrics. | [optional] 

## Methods

### NewRumMetricDefinition

`func NewRumMetricDefinition(metric string, ) *RumMetricDefinition`

NewRumMetricDefinition instantiates a new RumMetricDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRumMetricDefinitionWithDefaults

`func NewRumMetricDefinitionWithDefaults() *RumMetricDefinition`

NewRumMetricDefinitionWithDefaults instantiates a new RumMetricDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *RumMetricDefinition) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *RumMetricDefinition) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *RumMetricDefinition) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetPropertyKey

`func (o *RumMetricDefinition) GetPropertyKey() string`

GetPropertyKey returns the PropertyKey field if non-nil, zero value otherwise.

### GetPropertyKeyOk

`func (o *RumMetricDefinition) GetPropertyKeyOk() (*string, bool)`

GetPropertyKeyOk returns a tuple with the PropertyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyKey

`func (o *RumMetricDefinition) SetPropertyKey(v string)`

SetPropertyKey sets PropertyKey field to given value.

### HasPropertyKey

`func (o *RumMetricDefinition) HasPropertyKey() bool`

HasPropertyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


