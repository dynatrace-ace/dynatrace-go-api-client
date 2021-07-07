# MetricData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageKey** | **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | 
**TotalCount** | **int64** | The total number of primary entities in the result.   Has the &#x60;0&#x60; value if none of the requested metrics is suitable for pagination. | 
**Warnings** | Pointer to **[]string** | A list of warnings | [optional] 
**Result** | Pointer to [**[]MetricSeriesCollection**](MetricSeriesCollection.md) | A list of metrics and their data points. | [optional] 

## Methods

### NewMetricData

`func NewMetricData(nextPageKey string, totalCount int64, ) *MetricData`

NewMetricData instantiates a new MetricData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDataWithDefaults

`func NewMetricDataWithDefaults() *MetricData`

NewMetricDataWithDefaults instantiates a new MetricData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageKey

`func (o *MetricData) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *MetricData) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *MetricData) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.


### GetTotalCount

`func (o *MetricData) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *MetricData) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *MetricData) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetWarnings

`func (o *MetricData) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *MetricData) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *MetricData) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *MetricData) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetResult

`func (o *MetricData) GetResult() []MetricSeriesCollection`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MetricData) GetResultOk() (*[]MetricSeriesCollection, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *MetricData) SetResult(v []MetricSeriesCollection)`

SetResult sets Result field to given value.

### HasResult

`func (o *MetricData) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


