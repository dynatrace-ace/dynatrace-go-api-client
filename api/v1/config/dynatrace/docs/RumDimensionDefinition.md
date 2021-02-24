# RumDimensionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopX** | **int32** | The number of top values to be calculated. | 
**Dimension** | **string** | The dimension of the metric. | 
**PropertyKey** | Pointer to **string** | The key of the user action property.    Only applicable for the &#x60;StringProperty&#x60; dimension. | [optional] 

## Methods

### NewRumDimensionDefinition

`func NewRumDimensionDefinition(topX int32, dimension string, ) *RumDimensionDefinition`

NewRumDimensionDefinition instantiates a new RumDimensionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRumDimensionDefinitionWithDefaults

`func NewRumDimensionDefinitionWithDefaults() *RumDimensionDefinition`

NewRumDimensionDefinitionWithDefaults instantiates a new RumDimensionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopX

`func (o *RumDimensionDefinition) GetTopX() int32`

GetTopX returns the TopX field if non-nil, zero value otherwise.

### GetTopXOk

`func (o *RumDimensionDefinition) GetTopXOk() (*int32, bool)`

GetTopXOk returns a tuple with the TopX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopX

`func (o *RumDimensionDefinition) SetTopX(v int32)`

SetTopX sets TopX field to given value.


### GetDimension

`func (o *RumDimensionDefinition) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *RumDimensionDefinition) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *RumDimensionDefinition) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetPropertyKey

`func (o *RumDimensionDefinition) GetPropertyKey() string`

GetPropertyKey returns the PropertyKey field if non-nil, zero value otherwise.

### GetPropertyKeyOk

`func (o *RumDimensionDefinition) GetPropertyKeyOk() (*string, bool)`

GetPropertyKeyOk returns a tuple with the PropertyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyKey

`func (o *RumDimensionDefinition) SetPropertyKey(v string)`

SetPropertyKey sets PropertyKey field to given value.

### HasPropertyKey

`func (o *RumDimensionDefinition) HasPropertyKey() bool`

HasPropertyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


