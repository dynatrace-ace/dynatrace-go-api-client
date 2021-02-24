# IbmMQImsEntryQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the IMS entry queue. | [optional] 
**QueueManagerName** | **string** | The name of the queue manager. | 
**QueueName** | **string** | The name of the queue. | 

## Methods

### NewIbmMQImsEntryQueue

`func NewIbmMQImsEntryQueue(queueManagerName string, queueName string, ) *IbmMQImsEntryQueue`

NewIbmMQImsEntryQueue instantiates a new IbmMQImsEntryQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIbmMQImsEntryQueueWithDefaults

`func NewIbmMQImsEntryQueueWithDefaults() *IbmMQImsEntryQueue`

NewIbmMQImsEntryQueueWithDefaults instantiates a new IbmMQImsEntryQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IbmMQImsEntryQueue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IbmMQImsEntryQueue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IbmMQImsEntryQueue) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IbmMQImsEntryQueue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueueManagerName

`func (o *IbmMQImsEntryQueue) GetQueueManagerName() string`

GetQueueManagerName returns the QueueManagerName field if non-nil, zero value otherwise.

### GetQueueManagerNameOk

`func (o *IbmMQImsEntryQueue) GetQueueManagerNameOk() (*string, bool)`

GetQueueManagerNameOk returns a tuple with the QueueManagerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueManagerName

`func (o *IbmMQImsEntryQueue) SetQueueManagerName(v string)`

SetQueueManagerName sets QueueManagerName field to given value.


### GetQueueName

`func (o *IbmMQImsEntryQueue) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *IbmMQImsEntryQueue) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *IbmMQImsEntryQueue) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


