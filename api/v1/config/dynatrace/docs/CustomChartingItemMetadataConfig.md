# CustomChartingItemMetadataConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastModified** | Pointer to **int64** | The timestamp of the last metadata modification, in UTC milliseconds. | [optional] 
**CustomColor** | **string** | The color of the metric in the chart, hex format. | 

## Methods

### NewCustomChartingItemMetadataConfig

`func NewCustomChartingItemMetadataConfig(customColor string, ) *CustomChartingItemMetadataConfig`

NewCustomChartingItemMetadataConfig instantiates a new CustomChartingItemMetadataConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomChartingItemMetadataConfigWithDefaults

`func NewCustomChartingItemMetadataConfigWithDefaults() *CustomChartingItemMetadataConfig`

NewCustomChartingItemMetadataConfigWithDefaults instantiates a new CustomChartingItemMetadataConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastModified

`func (o *CustomChartingItemMetadataConfig) GetLastModified() int64`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *CustomChartingItemMetadataConfig) GetLastModifiedOk() (*int64, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *CustomChartingItemMetadataConfig) SetLastModified(v int64)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *CustomChartingItemMetadataConfig) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetCustomColor

`func (o *CustomChartingItemMetadataConfig) GetCustomColor() string`

GetCustomColor returns the CustomColor field if non-nil, zero value otherwise.

### GetCustomColorOk

`func (o *CustomChartingItemMetadataConfig) GetCustomColorOk() (*string, bool)`

GetCustomColorOk returns a tuple with the CustomColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomColor

`func (o *CustomChartingItemMetadataConfig) SetCustomColor(v string)`

SetCustomColor sets CustomColor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


