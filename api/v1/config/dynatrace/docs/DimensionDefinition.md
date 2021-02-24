# DimensionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the dimension. | 
**Dimension** | **string** | The dimension value pattern.    You can define custom placeholders in the **placeholders** field and use them here. | 
**Placeholders** | Pointer to [**[]Placeholder**](Placeholder.md) | The list of custom placeholders to be used in a dimension value pattern. | [optional] 
**TopX** | **int32** | The number of top values to be calculated. | 
**TopXDirection** | **string** | How to calculate the **topX** values. | 
**TopXAggregation** | **string** | The aggregation of the dimension. | 

## Methods

### NewDimensionDefinition

`func NewDimensionDefinition(name string, dimension string, topX int32, topXDirection string, topXAggregation string, ) *DimensionDefinition`

NewDimensionDefinition instantiates a new DimensionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDimensionDefinitionWithDefaults

`func NewDimensionDefinitionWithDefaults() *DimensionDefinition`

NewDimensionDefinitionWithDefaults instantiates a new DimensionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DimensionDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DimensionDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DimensionDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetDimension

`func (o *DimensionDefinition) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *DimensionDefinition) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *DimensionDefinition) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetPlaceholders

`func (o *DimensionDefinition) GetPlaceholders() []Placeholder`

GetPlaceholders returns the Placeholders field if non-nil, zero value otherwise.

### GetPlaceholdersOk

`func (o *DimensionDefinition) GetPlaceholdersOk() (*[]Placeholder, bool)`

GetPlaceholdersOk returns a tuple with the Placeholders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholders

`func (o *DimensionDefinition) SetPlaceholders(v []Placeholder)`

SetPlaceholders sets Placeholders field to given value.

### HasPlaceholders

`func (o *DimensionDefinition) HasPlaceholders() bool`

HasPlaceholders returns a boolean if a field has been set.

### GetTopX

`func (o *DimensionDefinition) GetTopX() int32`

GetTopX returns the TopX field if non-nil, zero value otherwise.

### GetTopXOk

`func (o *DimensionDefinition) GetTopXOk() (*int32, bool)`

GetTopXOk returns a tuple with the TopX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopX

`func (o *DimensionDefinition) SetTopX(v int32)`

SetTopX sets TopX field to given value.


### GetTopXDirection

`func (o *DimensionDefinition) GetTopXDirection() string`

GetTopXDirection returns the TopXDirection field if non-nil, zero value otherwise.

### GetTopXDirectionOk

`func (o *DimensionDefinition) GetTopXDirectionOk() (*string, bool)`

GetTopXDirectionOk returns a tuple with the TopXDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopXDirection

`func (o *DimensionDefinition) SetTopXDirection(v string)`

SetTopXDirection sets TopXDirection field to given value.


### GetTopXAggregation

`func (o *DimensionDefinition) GetTopXAggregation() string`

GetTopXAggregation returns the TopXAggregation field if non-nil, zero value otherwise.

### GetTopXAggregationOk

`func (o *DimensionDefinition) GetTopXAggregationOk() (*string, bool)`

GetTopXAggregationOk returns a tuple with the TopXAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopXAggregation

`func (o *DimensionDefinition) SetTopXAggregation(v string)`

SetTopXAggregation sets TopXAggregation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


