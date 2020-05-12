# QueueManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the queue manager. | 
**Clusters** | **[]string** | The queue manager will have access to all cluster visible queues in these [clusters](https://www.ibm.com/support/knowledgecenter/en/SSFKSJ_7.5.0/com.ibm.mq.pro.doc/q002750_.htm). | 
**AliasQueues** | [**[]AliasQueue**](AliasQueue.md) | The list of alias queues in the queue manager. | 
**RemoteQueues** | [**[]RemoteQueue**](RemoteQueue.md) | The list of remote queues in the queue manager. | 
**ClusterQueues** | [**[]ClusterQueue**](ClusterQueue.md) | The list of cluster queues in the queue manager. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


