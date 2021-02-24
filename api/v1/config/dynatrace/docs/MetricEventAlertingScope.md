# MetricEventAlertingScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;ENTITY_ID&#x60; -&gt; EntityIdAlertingScope  * &#x60;MANAGEMENT_ZONE&#x60; -&gt; ManagementZoneAlertingScope  * &#x60;TAG&#x60; -&gt; TagFilterAlertingScope  * &#x60;NAME&#x60; -&gt; NameAlertingScope  * &#x60;CUSTOM_DEVICE_GROUP_NAME&#x60; -&gt; CustomDeviceGroupNameAlertingScope  * &#x60;HOST_GROUP_NAME&#x60; -&gt; HostGroupNameAlertingScope  * &#x60;HOST_NAME&#x60; -&gt; HostNameAlertingScope  * &#x60;PROCESS_GROUP_ID&#x60; -&gt; ProcessGroupIdAlertingScope  * &#x60;PROCESS_GROUP_NAME&#x60; -&gt; ProcessGroupNameAlertingScope   | 

## Methods

### NewMetricEventAlertingScope

`func NewMetricEventAlertingScope(filterType string, ) *MetricEventAlertingScope`

NewMetricEventAlertingScope instantiates a new MetricEventAlertingScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricEventAlertingScopeWithDefaults

`func NewMetricEventAlertingScopeWithDefaults() *MetricEventAlertingScope`

NewMetricEventAlertingScopeWithDefaults instantiates a new MetricEventAlertingScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *MetricEventAlertingScope) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *MetricEventAlertingScope) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *MetricEventAlertingScope) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


