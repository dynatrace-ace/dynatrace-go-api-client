# AwsSupportingServiceMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the metric of the supporting service. | 
**Statistic** | **string** | The statistic (aggregation) to be used for the metric. AVG_MIN_MAX value is 3 statistics at once: AVERAGE, MINIMUM and MAXIMUM | 
**Dimensions** | **[]string** | A list of metric&#39;s dimensions names. | 

## Methods

### NewAwsSupportingServiceMetric

`func NewAwsSupportingServiceMetric(name string, statistic string, dimensions []string, ) *AwsSupportingServiceMetric`

NewAwsSupportingServiceMetric instantiates a new AwsSupportingServiceMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSupportingServiceMetricWithDefaults

`func NewAwsSupportingServiceMetricWithDefaults() *AwsSupportingServiceMetric`

NewAwsSupportingServiceMetricWithDefaults instantiates a new AwsSupportingServiceMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AwsSupportingServiceMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsSupportingServiceMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsSupportingServiceMetric) SetName(v string)`

SetName sets Name field to given value.


### GetStatistic

`func (o *AwsSupportingServiceMetric) GetStatistic() string`

GetStatistic returns the Statistic field if non-nil, zero value otherwise.

### GetStatisticOk

`func (o *AwsSupportingServiceMetric) GetStatisticOk() (*string, bool)`

GetStatisticOk returns a tuple with the Statistic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistic

`func (o *AwsSupportingServiceMetric) SetStatistic(v string)`

SetStatistic sets Statistic field to given value.


### GetDimensions

`func (o *AwsSupportingServiceMetric) GetDimensions() []string`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *AwsSupportingServiceMetric) GetDimensionsOk() (*[]string, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *AwsSupportingServiceMetric) SetDimensions(v []string)`

SetDimensions sets Dimensions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


