# EntityTimeseriesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeseriesId** | **string** | The ID of the metric, where you want to post data points. | 
**Dimensions** | Pointer to **map[string]string** | Dimensions of the data points you&#39;re posting.   The key of the metric dimension must be defined earlier in the metric definition. | [optional] 
**DataPoints** | **[][]float32** | List of data points.   Each data point is an array, containing the timestamp and the value.   Timestamp is UTC milliseconds reported as a number, for example: &#x60;1520523365557&#x60;.   You have the guaranteed timeframe of **30 minutes** into the past.   A custom metric must be registered **before** you can report a metric value. Therefore, the timestamp for reporting a value must be after the registration time of the metric. | 

## Methods

### NewEntityTimeseriesData

`func NewEntityTimeseriesData(timeseriesId string, dataPoints [][]float32, ) *EntityTimeseriesData`

NewEntityTimeseriesData instantiates a new EntityTimeseriesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityTimeseriesDataWithDefaults

`func NewEntityTimeseriesDataWithDefaults() *EntityTimeseriesData`

NewEntityTimeseriesDataWithDefaults instantiates a new EntityTimeseriesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeseriesId

`func (o *EntityTimeseriesData) GetTimeseriesId() string`

GetTimeseriesId returns the TimeseriesId field if non-nil, zero value otherwise.

### GetTimeseriesIdOk

`func (o *EntityTimeseriesData) GetTimeseriesIdOk() (*string, bool)`

GetTimeseriesIdOk returns a tuple with the TimeseriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseriesId

`func (o *EntityTimeseriesData) SetTimeseriesId(v string)`

SetTimeseriesId sets TimeseriesId field to given value.


### GetDimensions

`func (o *EntityTimeseriesData) GetDimensions() map[string]string`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *EntityTimeseriesData) GetDimensionsOk() (*map[string]string, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *EntityTimeseriesData) SetDimensions(v map[string]string)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *EntityTimeseriesData) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetDataPoints

`func (o *EntityTimeseriesData) GetDataPoints() [][]float32`

GetDataPoints returns the DataPoints field if non-nil, zero value otherwise.

### GetDataPointsOk

`func (o *EntityTimeseriesData) GetDataPointsOk() (*[][]float32, bool)`

GetDataPointsOk returns a tuple with the DataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPoints

`func (o *EntityTimeseriesData) SetDataPoints(v [][]float32)`

SetDataPoints sets DataPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


