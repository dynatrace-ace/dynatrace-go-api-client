# CalculatedMetricDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | The metric to be captured. | 
**RequestAttribute** | Pointer to **string** | The request attribute to be captured.    Only applicable when the **metric** parameter is set to &#x60;REQUEST_ATTRIBUTE&#x60;. | [optional] 

## Methods

### NewCalculatedMetricDefinition

`func NewCalculatedMetricDefinition(metric string, ) *CalculatedMetricDefinition`

NewCalculatedMetricDefinition instantiates a new CalculatedMetricDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalculatedMetricDefinitionWithDefaults

`func NewCalculatedMetricDefinitionWithDefaults() *CalculatedMetricDefinition`

NewCalculatedMetricDefinitionWithDefaults instantiates a new CalculatedMetricDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *CalculatedMetricDefinition) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *CalculatedMetricDefinition) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *CalculatedMetricDefinition) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetRequestAttribute

`func (o *CalculatedMetricDefinition) GetRequestAttribute() string`

GetRequestAttribute returns the RequestAttribute field if non-nil, zero value otherwise.

### GetRequestAttributeOk

`func (o *CalculatedMetricDefinition) GetRequestAttributeOk() (*string, bool)`

GetRequestAttributeOk returns a tuple with the RequestAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAttribute

`func (o *CalculatedMetricDefinition) SetRequestAttribute(v string)`

SetRequestAttribute sets RequestAttribute field to given value.

### HasRequestAttribute

`func (o *CalculatedMetricDefinition) HasRequestAttribute() bool`

HasRequestAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


