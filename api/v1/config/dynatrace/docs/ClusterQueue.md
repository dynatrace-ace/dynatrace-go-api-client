# ClusterQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalQueue** | **string** | The name of the local queue. | 
**ClusterVisibility** | **[]string** | The local queue is visible in these [clusters](https://www.ibm.com/support/knowledgecenter/en/SSFKSJ_7.5.0/com.ibm.mq.pro.doc/q002750_.htm). The queue manager must be part of these clusters. | 

## Methods

### NewClusterQueue

`func NewClusterQueue(localQueue string, clusterVisibility []string, ) *ClusterQueue`

NewClusterQueue instantiates a new ClusterQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterQueueWithDefaults

`func NewClusterQueueWithDefaults() *ClusterQueue`

NewClusterQueueWithDefaults instantiates a new ClusterQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalQueue

`func (o *ClusterQueue) GetLocalQueue() string`

GetLocalQueue returns the LocalQueue field if non-nil, zero value otherwise.

### GetLocalQueueOk

`func (o *ClusterQueue) GetLocalQueueOk() (*string, bool)`

GetLocalQueueOk returns a tuple with the LocalQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalQueue

`func (o *ClusterQueue) SetLocalQueue(v string)`

SetLocalQueue sets LocalQueue field to given value.


### GetClusterVisibility

`func (o *ClusterQueue) GetClusterVisibility() []string`

GetClusterVisibility returns the ClusterVisibility field if non-nil, zero value otherwise.

### GetClusterVisibilityOk

`func (o *ClusterQueue) GetClusterVisibilityOk() (*[]string, bool)`

GetClusterVisibilityOk returns a tuple with the ClusterVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterVisibility

`func (o *ClusterQueue) SetClusterVisibility(v []string)`

SetClusterVisibility sets ClusterVisibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


