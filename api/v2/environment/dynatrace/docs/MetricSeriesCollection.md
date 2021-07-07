# MetricSeriesCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricId** | **string** | The key of the metric.   If any transformation is applied, it is included here. | 
**Data** | Pointer to [**[]MetricSeries**](MetricSeries.md) | Data points of the metric. | [optional] 
**Warnings** | Pointer to **[]string** | A list of potential warnings that affect this ID. For example deprecated feature usage etc. | [optional] 

## Methods

### NewMetricSeriesCollection

`func NewMetricSeriesCollection(metricId string, ) *MetricSeriesCollection`

NewMetricSeriesCollection instantiates a new MetricSeriesCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricSeriesCollectionWithDefaults

`func NewMetricSeriesCollectionWithDefaults() *MetricSeriesCollection`

NewMetricSeriesCollectionWithDefaults instantiates a new MetricSeriesCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricId

`func (o *MetricSeriesCollection) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *MetricSeriesCollection) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *MetricSeriesCollection) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.


### GetData

`func (o *MetricSeriesCollection) GetData() []MetricSeries`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MetricSeriesCollection) GetDataOk() (*[]MetricSeries, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MetricSeriesCollection) SetData(v []MetricSeries)`

SetData sets Data field to given value.

### HasData

`func (o *MetricSeriesCollection) HasData() bool`

HasData returns a boolean if a field has been set.

### GetWarnings

`func (o *MetricSeriesCollection) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *MetricSeriesCollection) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *MetricSeriesCollection) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *MetricSeriesCollection) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


