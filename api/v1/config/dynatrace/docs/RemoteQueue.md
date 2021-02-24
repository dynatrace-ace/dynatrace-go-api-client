# RemoteQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalQueue** | **string** | The name of the local definition of the remote queue. | 
**RemoteQueue** | **string** | The name of the remote queue. | 
**RemoteQueueManager** | **string** | The name of the remote queue manager. | 
**ClusterVisibility** | **[]string** | The local definition of the remote queue is visible in these [clusters](https://www.ibm.com/support/knowledgecenter/en/SSFKSJ_7.5.0/com.ibm.mq.pro.doc/q002750_.htm). The queue manager must be part of these clusters. | 

## Methods

### NewRemoteQueue

`func NewRemoteQueue(localQueue string, remoteQueue string, remoteQueueManager string, clusterVisibility []string, ) *RemoteQueue`

NewRemoteQueue instantiates a new RemoteQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteQueueWithDefaults

`func NewRemoteQueueWithDefaults() *RemoteQueue`

NewRemoteQueueWithDefaults instantiates a new RemoteQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalQueue

`func (o *RemoteQueue) GetLocalQueue() string`

GetLocalQueue returns the LocalQueue field if non-nil, zero value otherwise.

### GetLocalQueueOk

`func (o *RemoteQueue) GetLocalQueueOk() (*string, bool)`

GetLocalQueueOk returns a tuple with the LocalQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalQueue

`func (o *RemoteQueue) SetLocalQueue(v string)`

SetLocalQueue sets LocalQueue field to given value.


### GetRemoteQueue

`func (o *RemoteQueue) GetRemoteQueue() string`

GetRemoteQueue returns the RemoteQueue field if non-nil, zero value otherwise.

### GetRemoteQueueOk

`func (o *RemoteQueue) GetRemoteQueueOk() (*string, bool)`

GetRemoteQueueOk returns a tuple with the RemoteQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteQueue

`func (o *RemoteQueue) SetRemoteQueue(v string)`

SetRemoteQueue sets RemoteQueue field to given value.


### GetRemoteQueueManager

`func (o *RemoteQueue) GetRemoteQueueManager() string`

GetRemoteQueueManager returns the RemoteQueueManager field if non-nil, zero value otherwise.

### GetRemoteQueueManagerOk

`func (o *RemoteQueue) GetRemoteQueueManagerOk() (*string, bool)`

GetRemoteQueueManagerOk returns a tuple with the RemoteQueueManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteQueueManager

`func (o *RemoteQueue) SetRemoteQueueManager(v string)`

SetRemoteQueueManager sets RemoteQueueManager field to given value.


### GetClusterVisibility

`func (o *RemoteQueue) GetClusterVisibility() []string`

GetClusterVisibility returns the ClusterVisibility field if non-nil, zero value otherwise.

### GetClusterVisibilityOk

`func (o *RemoteQueue) GetClusterVisibilityOk() (*[]string, bool)`

GetClusterVisibilityOk returns a tuple with the ClusterVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterVisibility

`func (o *RemoteQueue) SetClusterVisibility(v []string)`

SetClusterVisibility sets ClusterVisibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


