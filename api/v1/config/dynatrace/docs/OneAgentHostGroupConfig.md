# OneAgentHostGroupConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Dynatrace entity ID of the host group. | [optional] [readonly] 
**AutoUpdateConfig** | Pointer to [**HostGroupAutoUpdateConfig**](HostGroupAutoUpdateConfig.md) |  | [optional] 

## Methods

### NewOneAgentHostGroupConfig

`func NewOneAgentHostGroupConfig() *OneAgentHostGroupConfig`

NewOneAgentHostGroupConfig instantiates a new OneAgentHostGroupConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneAgentHostGroupConfigWithDefaults

`func NewOneAgentHostGroupConfigWithDefaults() *OneAgentHostGroupConfig`

NewOneAgentHostGroupConfigWithDefaults instantiates a new OneAgentHostGroupConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OneAgentHostGroupConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OneAgentHostGroupConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OneAgentHostGroupConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OneAgentHostGroupConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAutoUpdateConfig

`func (o *OneAgentHostGroupConfig) GetAutoUpdateConfig() HostGroupAutoUpdateConfig`

GetAutoUpdateConfig returns the AutoUpdateConfig field if non-nil, zero value otherwise.

### GetAutoUpdateConfigOk

`func (o *OneAgentHostGroupConfig) GetAutoUpdateConfigOk() (*HostGroupAutoUpdateConfig, bool)`

GetAutoUpdateConfigOk returns a tuple with the AutoUpdateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdateConfig

`func (o *OneAgentHostGroupConfig) SetAutoUpdateConfig(v HostGroupAutoUpdateConfig)`

SetAutoUpdateConfig sets AutoUpdateConfig field to given value.

### HasAutoUpdateConfig

`func (o *OneAgentHostGroupConfig) HasAutoUpdateConfig() bool`

HasAutoUpdateConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


