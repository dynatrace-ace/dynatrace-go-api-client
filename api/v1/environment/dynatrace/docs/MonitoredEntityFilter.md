# MonitoredEntityFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the Dynatrace entities (for example, hosts or services) you want to pick up by matching. | [optional] 
**Tags** | Pointer to [**[]UniversalTag**](UniversalTag.md) | The tag you want to use for matching.   You can use custom tags from the UI, AWS tags, Cloud Foundry tags, OpenShift/Kubernetes, and tags based on environment variables. | [optional] 

## Methods

### NewMonitoredEntityFilter

`func NewMonitoredEntityFilter() *MonitoredEntityFilter`

NewMonitoredEntityFilter instantiates a new MonitoredEntityFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoredEntityFilterWithDefaults

`func NewMonitoredEntityFilterWithDefaults() *MonitoredEntityFilter`

NewMonitoredEntityFilterWithDefaults instantiates a new MonitoredEntityFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MonitoredEntityFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitoredEntityFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitoredEntityFilter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MonitoredEntityFilter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTags

`func (o *MonitoredEntityFilter) GetTags() []UniversalTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MonitoredEntityFilter) GetTagsOk() (*[]UniversalTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MonitoredEntityFilter) SetTags(v []UniversalTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MonitoredEntityFilter) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


