# MetricDimensionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | The unique 0-based index of the dimension.    Appending transformations such as :names or :parents may change the indexes of dimensions. &#x60;null&#x60; is used for the dimensions of a metric with flexible dimensions, which can be referenced with their dimension key, but do not have an intrinsic order that could be used for the index. | 
**Name** | **string** | The name of the dimension. | 
**Key** | **string** | The key of the dimension.    It must be unique within the metric. | 
**Type** | **string** | The type of the dimension. | 
**DisplayName** | **string** | The display name of the dimension. | 

## Methods

### NewMetricDimensionDefinition

`func NewMetricDimensionDefinition(index int32, name string, key string, type_ string, displayName string, ) *MetricDimensionDefinition`

NewMetricDimensionDefinition instantiates a new MetricDimensionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDimensionDefinitionWithDefaults

`func NewMetricDimensionDefinitionWithDefaults() *MetricDimensionDefinition`

NewMetricDimensionDefinitionWithDefaults instantiates a new MetricDimensionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *MetricDimensionDefinition) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MetricDimensionDefinition) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *MetricDimensionDefinition) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetName

`func (o *MetricDimensionDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricDimensionDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricDimensionDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *MetricDimensionDefinition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricDimensionDefinition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricDimensionDefinition) SetKey(v string)`

SetKey sets Key field to given value.


### GetType

`func (o *MetricDimensionDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricDimensionDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricDimensionDefinition) SetType(v string)`

SetType sets Type field to given value.


### GetDisplayName

`func (o *MetricDimensionDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MetricDimensionDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MetricDimensionDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


