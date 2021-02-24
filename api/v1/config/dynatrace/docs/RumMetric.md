# RumMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationIdentifier** | **string** | The Dynatrace entity ID of the application to which the metric belongs. | 
**Name** | **string** | The displayed name of the metric. | 
**MetricKey** | **string** | The unique key of the metric.    The key must have the &#x60;calc:apps&#x60; prefix. | 
**Enabled** | **bool** | The metric is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**MetricDefinition** | [**RumMetricDefinition**](RumMetricDefinition.md) |  | 
**Dimensions** | Pointer to [**[]RumDimensionDefinition**](RumDimensionDefinition.md) | A list of metric dimensions. | [optional] 
**UserActionFilter** | Pointer to [**UserActionFilter**](UserActionFilter.md) |  | [optional] 

## Methods

### NewRumMetric

`func NewRumMetric(applicationIdentifier string, name string, metricKey string, enabled bool, metricDefinition RumMetricDefinition, ) *RumMetric`

NewRumMetric instantiates a new RumMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRumMetricWithDefaults

`func NewRumMetricWithDefaults() *RumMetric`

NewRumMetricWithDefaults instantiates a new RumMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationIdentifier

`func (o *RumMetric) GetApplicationIdentifier() string`

GetApplicationIdentifier returns the ApplicationIdentifier field if non-nil, zero value otherwise.

### GetApplicationIdentifierOk

`func (o *RumMetric) GetApplicationIdentifierOk() (*string, bool)`

GetApplicationIdentifierOk returns a tuple with the ApplicationIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIdentifier

`func (o *RumMetric) SetApplicationIdentifier(v string)`

SetApplicationIdentifier sets ApplicationIdentifier field to given value.


### GetName

`func (o *RumMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RumMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RumMetric) SetName(v string)`

SetName sets Name field to given value.


### GetMetricKey

`func (o *RumMetric) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *RumMetric) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *RumMetric) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetEnabled

`func (o *RumMetric) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RumMetric) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RumMetric) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMetricDefinition

`func (o *RumMetric) GetMetricDefinition() RumMetricDefinition`

GetMetricDefinition returns the MetricDefinition field if non-nil, zero value otherwise.

### GetMetricDefinitionOk

`func (o *RumMetric) GetMetricDefinitionOk() (*RumMetricDefinition, bool)`

GetMetricDefinitionOk returns a tuple with the MetricDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDefinition

`func (o *RumMetric) SetMetricDefinition(v RumMetricDefinition)`

SetMetricDefinition sets MetricDefinition field to given value.


### GetDimensions

`func (o *RumMetric) GetDimensions() []RumDimensionDefinition`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *RumMetric) GetDimensionsOk() (*[]RumDimensionDefinition, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *RumMetric) SetDimensions(v []RumDimensionDefinition)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *RumMetric) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetUserActionFilter

`func (o *RumMetric) GetUserActionFilter() UserActionFilter`

GetUserActionFilter returns the UserActionFilter field if non-nil, zero value otherwise.

### GetUserActionFilterOk

`func (o *RumMetric) GetUserActionFilterOk() (*UserActionFilter, bool)`

GetUserActionFilterOk returns a tuple with the UserActionFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionFilter

`func (o *RumMetric) SetUserActionFilter(v UserActionFilter)`

SetUserActionFilter sets UserActionFilter field to given value.

### HasUserActionFilter

`func (o *RumMetric) HasUserActionFilter() bool`

HasUserActionFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


