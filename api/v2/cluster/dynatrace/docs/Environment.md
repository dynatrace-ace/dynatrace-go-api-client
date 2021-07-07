# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The display name of the environment. | 
**Id** | Pointer to **string** | The ID of the environment. | [optional] 
**Trial** | Pointer to **bool** | Specifies whether the environment is a trial environment or a non-trial environment. Creating a trial environment is only possible if your license allows that. The default value is false (non-trial). | [optional] 
**State** | Pointer to **string** | Indicates whether the environment is enabled or disabled. The default value is ENABLED. | [optional] 
**Tags** | Pointer to **[]string** | A set of tags that are assigned to this environment. Every tag can have a maximum length of 100 characters. | [optional] 
**CreationDate** | Pointer to **string** | The creation date of the environment in ISO 8601 format (yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;) | [optional] [readonly] 
**Quotas** | Pointer to [**EnvironmentQuotas**](EnvironmentQuotas.md) |  | [optional] 
**Storage** | Pointer to [**EnvironmentStorage**](EnvironmentStorage.md) |  | [optional] 

## Methods

### NewEnvironment

`func NewEnvironment(name string, ) *Environment`

NewEnvironment instantiates a new Environment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentWithDefaults

`func NewEnvironmentWithDefaults() *Environment`

NewEnvironmentWithDefaults instantiates a new Environment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Environment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Environment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Environment) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *Environment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Environment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Environment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Environment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTrial

`func (o *Environment) GetTrial() bool`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *Environment) GetTrialOk() (*bool, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *Environment) SetTrial(v bool)`

SetTrial sets Trial field to given value.

### HasTrial

`func (o *Environment) HasTrial() bool`

HasTrial returns a boolean if a field has been set.

### GetState

`func (o *Environment) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Environment) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Environment) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Environment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *Environment) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Environment) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Environment) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Environment) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreationDate

`func (o *Environment) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Environment) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Environment) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Environment) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetQuotas

`func (o *Environment) GetQuotas() EnvironmentQuotas`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *Environment) GetQuotasOk() (*EnvironmentQuotas, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *Environment) SetQuotas(v EnvironmentQuotas)`

SetQuotas sets Quotas field to given value.

### HasQuotas

`func (o *Environment) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.

### GetStorage

`func (o *Environment) GetStorage() EnvironmentStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *Environment) GetStorageOk() (*EnvironmentStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *Environment) SetStorage(v EnvironmentStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *Environment) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


