# RemoveNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **int32** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 

## Methods

### NewRemoveNode

`func NewRemoveNode() *RemoveNode`

NewRemoveNode instantiates a new RemoveNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveNodeWithDefaults

`func NewRemoveNodeWithDefaults() *RemoveNode`

NewRemoveNodeWithDefaults instantiates a new RemoveNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *RemoveNode) GetNodeId() int32`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *RemoveNode) GetNodeIdOk() (*int32, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *RemoveNode) SetNodeId(v int32)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *RemoveNode) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetIpAddress

`func (o *RemoveNode) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *RemoveNode) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *RemoveNode) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *RemoveNode) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


