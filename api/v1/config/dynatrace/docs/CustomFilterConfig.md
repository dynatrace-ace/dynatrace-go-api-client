# CustomFilterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the filter.    It shows to which entity the filter belongs.    Custom charts have the &#x60;MIXED&#x60; type. | 
**CustomName** | **string** | The name of the tile, set by user | 
**DefaultName** | **string** | The default name of the tile | 
**ChartConfig** | [**CustomFilterChartConfig**](CustomFilterChartConfig.md) |  | 
**FiltersPerEntityType** | **map[string]interface{}** | A list of filters, applied to specific entity types. | 

## Methods

### NewCustomFilterConfig

`func NewCustomFilterConfig(type_ string, customName string, defaultName string, chartConfig CustomFilterChartConfig, filtersPerEntityType map[string]interface{}, ) *CustomFilterConfig`

NewCustomFilterConfig instantiates a new CustomFilterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFilterConfigWithDefaults

`func NewCustomFilterConfigWithDefaults() *CustomFilterConfig`

NewCustomFilterConfigWithDefaults instantiates a new CustomFilterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomFilterConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFilterConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFilterConfig) SetType(v string)`

SetType sets Type field to given value.


### GetCustomName

`func (o *CustomFilterConfig) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *CustomFilterConfig) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *CustomFilterConfig) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.


### GetDefaultName

`func (o *CustomFilterConfig) GetDefaultName() string`

GetDefaultName returns the DefaultName field if non-nil, zero value otherwise.

### GetDefaultNameOk

`func (o *CustomFilterConfig) GetDefaultNameOk() (*string, bool)`

GetDefaultNameOk returns a tuple with the DefaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultName

`func (o *CustomFilterConfig) SetDefaultName(v string)`

SetDefaultName sets DefaultName field to given value.


### GetChartConfig

`func (o *CustomFilterConfig) GetChartConfig() CustomFilterChartConfig`

GetChartConfig returns the ChartConfig field if non-nil, zero value otherwise.

### GetChartConfigOk

`func (o *CustomFilterConfig) GetChartConfigOk() (*CustomFilterChartConfig, bool)`

GetChartConfigOk returns a tuple with the ChartConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartConfig

`func (o *CustomFilterConfig) SetChartConfig(v CustomFilterChartConfig)`

SetChartConfig sets ChartConfig field to given value.


### GetFiltersPerEntityType

`func (o *CustomFilterConfig) GetFiltersPerEntityType() map[string]interface{}`

GetFiltersPerEntityType returns the FiltersPerEntityType field if non-nil, zero value otherwise.

### GetFiltersPerEntityTypeOk

`func (o *CustomFilterConfig) GetFiltersPerEntityTypeOk() (*map[string]interface{}, bool)`

GetFiltersPerEntityTypeOk returns a tuple with the FiltersPerEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltersPerEntityType

`func (o *CustomFilterConfig) SetFiltersPerEntityType(v map[string]interface{})`

SetFiltersPerEntityType sets FiltersPerEntityType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


