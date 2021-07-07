# NodeConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**WebUI** | Pointer to **bool** |  | [optional] 
**Agent** | Pointer to **bool** |  | [optional] 
**Datacenter** | Pointer to **string** |  | [optional] 
**KubernetesRole** | Pointer to **string** |  | [optional] 

## Methods

### NewNodeConfigDto

`func NewNodeConfigDto() *NodeConfigDto`

NewNodeConfigDto instantiates a new NodeConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeConfigDtoWithDefaults

`func NewNodeConfigDtoWithDefaults() *NodeConfigDto`

NewNodeConfigDtoWithDefaults instantiates a new NodeConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeConfigDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeConfigDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeConfigDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NodeConfigDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWebUI

`func (o *NodeConfigDto) GetWebUI() bool`

GetWebUI returns the WebUI field if non-nil, zero value otherwise.

### GetWebUIOk

`func (o *NodeConfigDto) GetWebUIOk() (*bool, bool)`

GetWebUIOk returns a tuple with the WebUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUI

`func (o *NodeConfigDto) SetWebUI(v bool)`

SetWebUI sets WebUI field to given value.

### HasWebUI

`func (o *NodeConfigDto) HasWebUI() bool`

HasWebUI returns a boolean if a field has been set.

### GetAgent

`func (o *NodeConfigDto) GetAgent() bool`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *NodeConfigDto) GetAgentOk() (*bool, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *NodeConfigDto) SetAgent(v bool)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *NodeConfigDto) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetDatacenter

`func (o *NodeConfigDto) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *NodeConfigDto) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *NodeConfigDto) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *NodeConfigDto) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetKubernetesRole

`func (o *NodeConfigDto) GetKubernetesRole() string`

GetKubernetesRole returns the KubernetesRole field if non-nil, zero value otherwise.

### GetKubernetesRoleOk

`func (o *NodeConfigDto) GetKubernetesRoleOk() (*string, bool)`

GetKubernetesRoleOk returns a tuple with the KubernetesRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesRole

`func (o *NodeConfigDto) SetKubernetesRole(v string)`

SetKubernetesRole sets KubernetesRole field to given value.

### HasKubernetesRole

`func (o *NodeConfigDto) HasKubernetesRole() bool`

HasKubernetesRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


