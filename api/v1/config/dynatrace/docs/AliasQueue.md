# AliasQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliasQueue** | **string** | The name of the alias queue. | 
**BaseQueue** | **string** | The name of the base queue, which the alias queue should point to. | 
**ClusterVisibility** | **[]string** | The alias queue is visible in these [clusters](https://www.ibm.com/support/knowledgecenter/en/SSFKSJ_7.5.0/com.ibm.mq.pro.doc/q002750_.htm). The queue manager must be part of these clusters. | 

## Methods

### NewAliasQueue

`func NewAliasQueue(aliasQueue string, baseQueue string, clusterVisibility []string, ) *AliasQueue`

NewAliasQueue instantiates a new AliasQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAliasQueueWithDefaults

`func NewAliasQueueWithDefaults() *AliasQueue`

NewAliasQueueWithDefaults instantiates a new AliasQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliasQueue

`func (o *AliasQueue) GetAliasQueue() string`

GetAliasQueue returns the AliasQueue field if non-nil, zero value otherwise.

### GetAliasQueueOk

`func (o *AliasQueue) GetAliasQueueOk() (*string, bool)`

GetAliasQueueOk returns a tuple with the AliasQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasQueue

`func (o *AliasQueue) SetAliasQueue(v string)`

SetAliasQueue sets AliasQueue field to given value.


### GetBaseQueue

`func (o *AliasQueue) GetBaseQueue() string`

GetBaseQueue returns the BaseQueue field if non-nil, zero value otherwise.

### GetBaseQueueOk

`func (o *AliasQueue) GetBaseQueueOk() (*string, bool)`

GetBaseQueueOk returns a tuple with the BaseQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseQueue

`func (o *AliasQueue) SetBaseQueue(v string)`

SetBaseQueue sets BaseQueue field to given value.


### GetClusterVisibility

`func (o *AliasQueue) GetClusterVisibility() []string`

GetClusterVisibility returns the ClusterVisibility field if non-nil, zero value otherwise.

### GetClusterVisibilityOk

`func (o *AliasQueue) GetClusterVisibilityOk() (*[]string, bool)`

GetClusterVisibilityOk returns a tuple with the ClusterVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterVisibility

`func (o *AliasQueue) SetClusterVisibility(v []string)`

SetClusterVisibility sets ClusterVisibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


