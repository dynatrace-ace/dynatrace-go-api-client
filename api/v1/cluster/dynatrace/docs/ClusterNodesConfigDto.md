# ClusterNodesConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterNodes** | Pointer to [**[]NodeConfigDto**](NodeConfigDto.md) |  | [optional] 

## Methods

### NewClusterNodesConfigDto

`func NewClusterNodesConfigDto() *ClusterNodesConfigDto`

NewClusterNodesConfigDto instantiates a new ClusterNodesConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodesConfigDtoWithDefaults

`func NewClusterNodesConfigDtoWithDefaults() *ClusterNodesConfigDto`

NewClusterNodesConfigDtoWithDefaults instantiates a new ClusterNodesConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterNodes

`func (o *ClusterNodesConfigDto) GetClusterNodes() []NodeConfigDto`

GetClusterNodes returns the ClusterNodes field if non-nil, zero value otherwise.

### GetClusterNodesOk

`func (o *ClusterNodesConfigDto) GetClusterNodesOk() (*[]NodeConfigDto, bool)`

GetClusterNodesOk returns a tuple with the ClusterNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodes

`func (o *ClusterNodesConfigDto) SetClusterNodes(v []NodeConfigDto)`

SetClusterNodes sets ClusterNodes field to given value.

### HasClusterNodes

`func (o *ClusterNodesConfigDto) HasClusterNodes() bool`

HasClusterNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


