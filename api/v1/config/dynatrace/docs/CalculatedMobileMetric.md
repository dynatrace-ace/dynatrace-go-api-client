# CalculatedMobileMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationIdentifier** | **string** | The Dynatrace entity ID of the application to which the metric belongs. | 
**Name** | **string** | The name of the metric, displayed in the UI. | 
**MetricKey** | **string** | The unique key of the metric.    The key must have the &#x60;calc:apps&#x60; prefix. | 
**Enabled** | **bool** | The metric is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**MetricType** | **string** | The type of the metric. | 
**Dimensions** | Pointer to [**[]CalculatedMobileMetricDimension**](CalculatedMobileMetricDimension.md) | A list of metric dimensions. | [optional] 
**UserActionFilter** | Pointer to [**CalculatedMobileMetricUserActionFilter**](CalculatedMobileMetricUserActionFilter.md) |  | [optional] 

## Methods

### NewCalculatedMobileMetric

`func NewCalculatedMobileMetric(applicationIdentifier string, name string, metricKey string, enabled bool, metricType string, ) *CalculatedMobileMetric`

NewCalculatedMobileMetric instantiates a new CalculatedMobileMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalculatedMobileMetricWithDefaults

`func NewCalculatedMobileMetricWithDefaults() *CalculatedMobileMetric`

NewCalculatedMobileMetricWithDefaults instantiates a new CalculatedMobileMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationIdentifier

`func (o *CalculatedMobileMetric) GetApplicationIdentifier() string`

GetApplicationIdentifier returns the ApplicationIdentifier field if non-nil, zero value otherwise.

### GetApplicationIdentifierOk

`func (o *CalculatedMobileMetric) GetApplicationIdentifierOk() (*string, bool)`

GetApplicationIdentifierOk returns a tuple with the ApplicationIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIdentifier

`func (o *CalculatedMobileMetric) SetApplicationIdentifier(v string)`

SetApplicationIdentifier sets ApplicationIdentifier field to given value.


### GetName

`func (o *CalculatedMobileMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CalculatedMobileMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CalculatedMobileMetric) SetName(v string)`

SetName sets Name field to given value.


### GetMetricKey

`func (o *CalculatedMobileMetric) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *CalculatedMobileMetric) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *CalculatedMobileMetric) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetEnabled

`func (o *CalculatedMobileMetric) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CalculatedMobileMetric) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CalculatedMobileMetric) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMetricType

`func (o *CalculatedMobileMetric) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *CalculatedMobileMetric) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *CalculatedMobileMetric) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.


### GetDimensions

`func (o *CalculatedMobileMetric) GetDimensions() []CalculatedMobileMetricDimension`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *CalculatedMobileMetric) GetDimensionsOk() (*[]CalculatedMobileMetricDimension, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *CalculatedMobileMetric) SetDimensions(v []CalculatedMobileMetricDimension)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *CalculatedMobileMetric) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetUserActionFilter

`func (o *CalculatedMobileMetric) GetUserActionFilter() CalculatedMobileMetricUserActionFilter`

GetUserActionFilter returns the UserActionFilter field if non-nil, zero value otherwise.

### GetUserActionFilterOk

`func (o *CalculatedMobileMetric) GetUserActionFilterOk() (*CalculatedMobileMetricUserActionFilter, bool)`

GetUserActionFilterOk returns a tuple with the UserActionFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionFilter

`func (o *CalculatedMobileMetric) SetUserActionFilter(v CalculatedMobileMetricUserActionFilter)`

SetUserActionFilter sets UserActionFilter field to given value.

### HasUserActionFilter

`func (o *CalculatedMobileMetric) HasUserActionFilter() bool`

HasUserActionFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


