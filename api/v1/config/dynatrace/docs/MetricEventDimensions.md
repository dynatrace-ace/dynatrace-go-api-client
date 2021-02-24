# MetricEventDimensions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;ENTITY&#x60; -&gt; MetricEventEntityDimensions  * &#x60;STRING&#x60; -&gt; MetricEventStringDimensions   | 
**Key** | Pointer to **string** | The dimensions key on the metric. | [optional] 

## Methods

### NewMetricEventDimensions

`func NewMetricEventDimensions(filterType string, ) *MetricEventDimensions`

NewMetricEventDimensions instantiates a new MetricEventDimensions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricEventDimensionsWithDefaults

`func NewMetricEventDimensionsWithDefaults() *MetricEventDimensions`

NewMetricEventDimensionsWithDefaults instantiates a new MetricEventDimensions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *MetricEventDimensions) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *MetricEventDimensions) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *MetricEventDimensions) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetKey

`func (o *MetricEventDimensions) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricEventDimensions) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricEventDimensions) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MetricEventDimensions) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


