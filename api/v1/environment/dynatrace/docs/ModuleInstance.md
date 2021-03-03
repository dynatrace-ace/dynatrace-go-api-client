# ModuleInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceName** | Pointer to **string** | The name of the instance. | [optional] 
**ModuleVersion** | Pointer to **string** | The version of the code module. | [optional] 
**FaultyVersion** | Pointer to **bool** | The code module version is faulty (&#x60;true&#x60;) or not (&#x60;false&#x60;). | [optional] 
**Active** | Pointer to **bool** | The code module instance is active (&#x60;true&#x60;) or inactive (&#x60;false&#x60;). | [optional] 

## Methods

### NewModuleInstance

`func NewModuleInstance() *ModuleInstance`

NewModuleInstance instantiates a new ModuleInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleInstanceWithDefaults

`func NewModuleInstanceWithDefaults() *ModuleInstance`

NewModuleInstanceWithDefaults instantiates a new ModuleInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceName

`func (o *ModuleInstance) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *ModuleInstance) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *ModuleInstance) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *ModuleInstance) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetModuleVersion

`func (o *ModuleInstance) GetModuleVersion() string`

GetModuleVersion returns the ModuleVersion field if non-nil, zero value otherwise.

### GetModuleVersionOk

`func (o *ModuleInstance) GetModuleVersionOk() (*string, bool)`

GetModuleVersionOk returns a tuple with the ModuleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleVersion

`func (o *ModuleInstance) SetModuleVersion(v string)`

SetModuleVersion sets ModuleVersion field to given value.

### HasModuleVersion

`func (o *ModuleInstance) HasModuleVersion() bool`

HasModuleVersion returns a boolean if a field has been set.

### GetFaultyVersion

`func (o *ModuleInstance) GetFaultyVersion() bool`

GetFaultyVersion returns the FaultyVersion field if non-nil, zero value otherwise.

### GetFaultyVersionOk

`func (o *ModuleInstance) GetFaultyVersionOk() (*bool, bool)`

GetFaultyVersionOk returns a tuple with the FaultyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultyVersion

`func (o *ModuleInstance) SetFaultyVersion(v bool)`

SetFaultyVersion sets FaultyVersion field to given value.

### HasFaultyVersion

`func (o *ModuleInstance) HasFaultyVersion() bool`

HasFaultyVersion returns a boolean if a field has been set.

### GetActive

`func (o *ModuleInstance) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ModuleInstance) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ModuleInstance) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ModuleInstance) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


