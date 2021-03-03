# ModuleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModuleType** | Pointer to **string** | The type of the code module. | [optional] 
**Instances** | Pointer to [**[]ModuleInstance**](ModuleInstance.md) | A list of instances of the code module. | [optional] 

## Methods

### NewModuleInfo

`func NewModuleInfo() *ModuleInfo`

NewModuleInfo instantiates a new ModuleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleInfoWithDefaults

`func NewModuleInfoWithDefaults() *ModuleInfo`

NewModuleInfoWithDefaults instantiates a new ModuleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModuleType

`func (o *ModuleInfo) GetModuleType() string`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *ModuleInfo) GetModuleTypeOk() (*string, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *ModuleInfo) SetModuleType(v string)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *ModuleInfo) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### GetInstances

`func (o *ModuleInfo) GetInstances() []ModuleInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ModuleInfo) GetInstancesOk() (*[]ModuleInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ModuleInfo) SetInstances(v []ModuleInstance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ModuleInfo) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


