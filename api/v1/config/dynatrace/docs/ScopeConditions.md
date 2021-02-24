# ScopeConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceTechnology** | Pointer to **string** | Only applies to this service technology. | [optional] 
**ProcessGroup** | Pointer to **string** | Only applies to this process group. Note that this can&#39;t be transferred between different clusters or environments. | [optional] 
**HostGroup** | Pointer to **string** | Only applies to this host group. | [optional] 
**TagOfProcessGroup** | Pointer to **string** | Only apply to process groups matching this tag. | [optional] 

## Methods

### NewScopeConditions

`func NewScopeConditions() *ScopeConditions`

NewScopeConditions instantiates a new ScopeConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeConditionsWithDefaults

`func NewScopeConditionsWithDefaults() *ScopeConditions`

NewScopeConditionsWithDefaults instantiates a new ScopeConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceTechnology

`func (o *ScopeConditions) GetServiceTechnology() string`

GetServiceTechnology returns the ServiceTechnology field if non-nil, zero value otherwise.

### GetServiceTechnologyOk

`func (o *ScopeConditions) GetServiceTechnologyOk() (*string, bool)`

GetServiceTechnologyOk returns a tuple with the ServiceTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTechnology

`func (o *ScopeConditions) SetServiceTechnology(v string)`

SetServiceTechnology sets ServiceTechnology field to given value.

### HasServiceTechnology

`func (o *ScopeConditions) HasServiceTechnology() bool`

HasServiceTechnology returns a boolean if a field has been set.

### GetProcessGroup

`func (o *ScopeConditions) GetProcessGroup() string`

GetProcessGroup returns the ProcessGroup field if non-nil, zero value otherwise.

### GetProcessGroupOk

`func (o *ScopeConditions) GetProcessGroupOk() (*string, bool)`

GetProcessGroupOk returns a tuple with the ProcessGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroup

`func (o *ScopeConditions) SetProcessGroup(v string)`

SetProcessGroup sets ProcessGroup field to given value.

### HasProcessGroup

`func (o *ScopeConditions) HasProcessGroup() bool`

HasProcessGroup returns a boolean if a field has been set.

### GetHostGroup

`func (o *ScopeConditions) GetHostGroup() string`

GetHostGroup returns the HostGroup field if non-nil, zero value otherwise.

### GetHostGroupOk

`func (o *ScopeConditions) GetHostGroupOk() (*string, bool)`

GetHostGroupOk returns a tuple with the HostGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroup

`func (o *ScopeConditions) SetHostGroup(v string)`

SetHostGroup sets HostGroup field to given value.

### HasHostGroup

`func (o *ScopeConditions) HasHostGroup() bool`

HasHostGroup returns a boolean if a field has been set.

### GetTagOfProcessGroup

`func (o *ScopeConditions) GetTagOfProcessGroup() string`

GetTagOfProcessGroup returns the TagOfProcessGroup field if non-nil, zero value otherwise.

### GetTagOfProcessGroupOk

`func (o *ScopeConditions) GetTagOfProcessGroupOk() (*string, bool)`

GetTagOfProcessGroupOk returns a tuple with the TagOfProcessGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagOfProcessGroup

`func (o *ScopeConditions) SetTagOfProcessGroup(v string)`

SetTagOfProcessGroup sets TagOfProcessGroup field to given value.

### HasTagOfProcessGroup

`func (o *ScopeConditions) HasTagOfProcessGroup() bool`

HasTagOfProcessGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


