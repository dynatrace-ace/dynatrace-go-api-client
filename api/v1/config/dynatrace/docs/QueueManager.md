# QueueManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the queue manager. | 
**Clusters** | **[]string** | The queue manager will have access to all cluster visible queues in these [clusters](https://www.ibm.com/support/knowledgecenter/en/SSFKSJ_7.5.0/com.ibm.mq.pro.doc/q002750_.htm). | 
**AliasQueues** | [**[]AliasQueue**](AliasQueue.md) | The list of alias queues in the queue manager. | 
**RemoteQueues** | [**[]RemoteQueue**](RemoteQueue.md) | The list of remote queues in the queue manager. | 
**ClusterQueues** | [**[]ClusterQueue**](ClusterQueue.md) | The list of cluster queues in the queue manager. | 

## Methods

### NewQueueManager

`func NewQueueManager(name string, clusters []string, aliasQueues []AliasQueue, remoteQueues []RemoteQueue, clusterQueues []ClusterQueue, ) *QueueManager`

NewQueueManager instantiates a new QueueManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueManagerWithDefaults

`func NewQueueManagerWithDefaults() *QueueManager`

NewQueueManagerWithDefaults instantiates a new QueueManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QueueManager) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueueManager) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueueManager) SetName(v string)`

SetName sets Name field to given value.


### GetClusters

`func (o *QueueManager) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *QueueManager) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *QueueManager) SetClusters(v []string)`

SetClusters sets Clusters field to given value.


### GetAliasQueues

`func (o *QueueManager) GetAliasQueues() []AliasQueue`

GetAliasQueues returns the AliasQueues field if non-nil, zero value otherwise.

### GetAliasQueuesOk

`func (o *QueueManager) GetAliasQueuesOk() (*[]AliasQueue, bool)`

GetAliasQueuesOk returns a tuple with the AliasQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasQueues

`func (o *QueueManager) SetAliasQueues(v []AliasQueue)`

SetAliasQueues sets AliasQueues field to given value.


### GetRemoteQueues

`func (o *QueueManager) GetRemoteQueues() []RemoteQueue`

GetRemoteQueues returns the RemoteQueues field if non-nil, zero value otherwise.

### GetRemoteQueuesOk

`func (o *QueueManager) GetRemoteQueuesOk() (*[]RemoteQueue, bool)`

GetRemoteQueuesOk returns a tuple with the RemoteQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteQueues

`func (o *QueueManager) SetRemoteQueues(v []RemoteQueue)`

SetRemoteQueues sets RemoteQueues field to given value.


### GetClusterQueues

`func (o *QueueManager) GetClusterQueues() []ClusterQueue`

GetClusterQueues returns the ClusterQueues field if non-nil, zero value otherwise.

### GetClusterQueuesOk

`func (o *QueueManager) GetClusterQueuesOk() (*[]ClusterQueue, bool)`

GetClusterQueuesOk returns a tuple with the ClusterQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterQueues

`func (o *QueueManager) SetClusterQueues(v []ClusterQueue)`

SetClusterQueues sets ClusterQueues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


