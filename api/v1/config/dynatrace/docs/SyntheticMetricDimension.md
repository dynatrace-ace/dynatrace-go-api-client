# SyntheticMetricDimension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopX** | Pointer to **int32** | The number of top values to be calculated. | [optional] 
**Dimension** | **string** | The dimension of the metric. | 

## Methods

### NewSyntheticMetricDimension

`func NewSyntheticMetricDimension(dimension string, ) *SyntheticMetricDimension`

NewSyntheticMetricDimension instantiates a new SyntheticMetricDimension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticMetricDimensionWithDefaults

`func NewSyntheticMetricDimensionWithDefaults() *SyntheticMetricDimension`

NewSyntheticMetricDimensionWithDefaults instantiates a new SyntheticMetricDimension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopX

`func (o *SyntheticMetricDimension) GetTopX() int32`

GetTopX returns the TopX field if non-nil, zero value otherwise.

### GetTopXOk

`func (o *SyntheticMetricDimension) GetTopXOk() (*int32, bool)`

GetTopXOk returns a tuple with the TopX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopX

`func (o *SyntheticMetricDimension) SetTopX(v int32)`

SetTopX sets TopX field to given value.

### HasTopX

`func (o *SyntheticMetricDimension) HasTopX() bool`

HasTopX returns a boolean if a field has been set.

### GetDimension

`func (o *SyntheticMetricDimension) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *SyntheticMetricDimension) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *SyntheticMetricDimension) SetDimension(v string)`

SetDimension sets Dimension field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


