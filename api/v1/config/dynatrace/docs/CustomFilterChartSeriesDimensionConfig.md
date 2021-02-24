# CustomFilterChartSeriesDimensionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the dimension by which the metric is split. | 
**Name** | Pointer to **string** | The name of the dimension by which the metric is split. | [optional] 
**Values** | **[]string** | The splitting value. | 
**EntityDimension** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustomFilterChartSeriesDimensionConfig

`func NewCustomFilterChartSeriesDimensionConfig(id string, values []string, ) *CustomFilterChartSeriesDimensionConfig`

NewCustomFilterChartSeriesDimensionConfig instantiates a new CustomFilterChartSeriesDimensionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFilterChartSeriesDimensionConfigWithDefaults

`func NewCustomFilterChartSeriesDimensionConfigWithDefaults() *CustomFilterChartSeriesDimensionConfig`

NewCustomFilterChartSeriesDimensionConfigWithDefaults instantiates a new CustomFilterChartSeriesDimensionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFilterChartSeriesDimensionConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFilterChartSeriesDimensionConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFilterChartSeriesDimensionConfig) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CustomFilterChartSeriesDimensionConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFilterChartSeriesDimensionConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFilterChartSeriesDimensionConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomFilterChartSeriesDimensionConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValues

`func (o *CustomFilterChartSeriesDimensionConfig) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CustomFilterChartSeriesDimensionConfig) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CustomFilterChartSeriesDimensionConfig) SetValues(v []string)`

SetValues sets Values field to given value.


### GetEntityDimension

`func (o *CustomFilterChartSeriesDimensionConfig) GetEntityDimension() bool`

GetEntityDimension returns the EntityDimension field if non-nil, zero value otherwise.

### GetEntityDimensionOk

`func (o *CustomFilterChartSeriesDimensionConfig) GetEntityDimensionOk() (*bool, bool)`

GetEntityDimensionOk returns a tuple with the EntityDimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityDimension

`func (o *CustomFilterChartSeriesDimensionConfig) SetEntityDimension(v bool)`

SetEntityDimension sets EntityDimension field to given value.

### HasEntityDimension

`func (o *CustomFilterChartSeriesDimensionConfig) HasEntityDimension() bool`

HasEntityDimension returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


