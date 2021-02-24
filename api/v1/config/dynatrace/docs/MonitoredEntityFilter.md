# MonitoredEntityFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the Dynatrace entities (for example, hosts or services) you want to pick up by matching. | [optional] 
**MzId** | Pointer to **string** | The ID of a management zone to which the matched entities must belong. | [optional] 
**Tags** | [**[]TagInfo**](TagInfo.md) | The tag you want to use for matching.   You can use custom tags from the UI, AWS tags, Cloud Foundry tags, OpenShift/Kubernetes, and tags based on environment variables. | 
**TagCombination** | Pointer to **string** | The logic that applies when several tags are specified: AND/OR.   If not set, the OR logic is used. | [optional] 

## Methods

### NewMonitoredEntityFilter

`func NewMonitoredEntityFilter(tags []TagInfo, ) *MonitoredEntityFilter`

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

### GetMzId

`func (o *MonitoredEntityFilter) GetMzId() string`

GetMzId returns the MzId field if non-nil, zero value otherwise.

### GetMzIdOk

`func (o *MonitoredEntityFilter) GetMzIdOk() (*string, bool)`

GetMzIdOk returns a tuple with the MzId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMzId

`func (o *MonitoredEntityFilter) SetMzId(v string)`

SetMzId sets MzId field to given value.

### HasMzId

`func (o *MonitoredEntityFilter) HasMzId() bool`

HasMzId returns a boolean if a field has been set.

### GetTags

`func (o *MonitoredEntityFilter) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MonitoredEntityFilter) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MonitoredEntityFilter) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.


### GetTagCombination

`func (o *MonitoredEntityFilter) GetTagCombination() string`

GetTagCombination returns the TagCombination field if non-nil, zero value otherwise.

### GetTagCombinationOk

`func (o *MonitoredEntityFilter) GetTagCombinationOk() (*string, bool)`

GetTagCombinationOk returns a tuple with the TagCombination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCombination

`func (o *MonitoredEntityFilter) SetTagCombination(v string)`

SetTagCombination sets TagCombination field to given value.

### HasTagCombination

`func (o *MonitoredEntityFilter) HasTagCombination() bool`

HasTagCombination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


