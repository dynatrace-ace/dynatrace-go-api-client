# MetricDescriptorCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageKey** | **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | 
**TotalCount** | **int64** | The estimated number of metrics in the result. | 
**Warnings** | Pointer to **[]string** | A list of potential warnings about the query. For example deprecated feature usage etc. | [optional] 
**Metrics** | Pointer to [**[]MetricDescriptor**](MetricDescriptor.md) | A list of metric along with their descriptors | [optional] 

## Methods

### NewMetricDescriptorCollection

`func NewMetricDescriptorCollection(nextPageKey string, totalCount int64, ) *MetricDescriptorCollection`

NewMetricDescriptorCollection instantiates a new MetricDescriptorCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDescriptorCollectionWithDefaults

`func NewMetricDescriptorCollectionWithDefaults() *MetricDescriptorCollection`

NewMetricDescriptorCollectionWithDefaults instantiates a new MetricDescriptorCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageKey

`func (o *MetricDescriptorCollection) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *MetricDescriptorCollection) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *MetricDescriptorCollection) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.


### GetTotalCount

`func (o *MetricDescriptorCollection) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *MetricDescriptorCollection) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *MetricDescriptorCollection) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetWarnings

`func (o *MetricDescriptorCollection) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *MetricDescriptorCollection) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *MetricDescriptorCollection) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *MetricDescriptorCollection) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetMetrics

`func (o *MetricDescriptorCollection) GetMetrics() []MetricDescriptor`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MetricDescriptorCollection) GetMetricsOk() (*[]MetricDescriptor, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MetricDescriptorCollection) SetMetrics(v []MetricDescriptor)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *MetricDescriptorCollection) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


