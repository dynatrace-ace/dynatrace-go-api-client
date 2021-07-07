# VulnerableComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Dynatrace entity ID of the vulnerable component. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | The display name of the vulnerable component. | [optional] [readonly] 
**FileName** | Pointer to **string** | The file name of the vulnerable component. | [optional] [readonly] 
**NumberOfAffectedEntities** | Pointer to **int32** | The number of affected entities. | [optional] [readonly] 
**AffectedEntities** | Pointer to **[]string** | The list of affected entities. | [optional] [readonly] 

## Methods

### NewVulnerableComponent

`func NewVulnerableComponent() *VulnerableComponent`

NewVulnerableComponent instantiates a new VulnerableComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVulnerableComponentWithDefaults

`func NewVulnerableComponentWithDefaults() *VulnerableComponent`

NewVulnerableComponentWithDefaults instantiates a new VulnerableComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VulnerableComponent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VulnerableComponent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VulnerableComponent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VulnerableComponent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *VulnerableComponent) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VulnerableComponent) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VulnerableComponent) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *VulnerableComponent) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFileName

`func (o *VulnerableComponent) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *VulnerableComponent) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *VulnerableComponent) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *VulnerableComponent) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetNumberOfAffectedEntities

`func (o *VulnerableComponent) GetNumberOfAffectedEntities() int32`

GetNumberOfAffectedEntities returns the NumberOfAffectedEntities field if non-nil, zero value otherwise.

### GetNumberOfAffectedEntitiesOk

`func (o *VulnerableComponent) GetNumberOfAffectedEntitiesOk() (*int32, bool)`

GetNumberOfAffectedEntitiesOk returns a tuple with the NumberOfAffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAffectedEntities

`func (o *VulnerableComponent) SetNumberOfAffectedEntities(v int32)`

SetNumberOfAffectedEntities sets NumberOfAffectedEntities field to given value.

### HasNumberOfAffectedEntities

`func (o *VulnerableComponent) HasNumberOfAffectedEntities() bool`

HasNumberOfAffectedEntities returns a boolean if a field has been set.

### GetAffectedEntities

`func (o *VulnerableComponent) GetAffectedEntities() []string`

GetAffectedEntities returns the AffectedEntities field if non-nil, zero value otherwise.

### GetAffectedEntitiesOk

`func (o *VulnerableComponent) GetAffectedEntitiesOk() (*[]string, bool)`

GetAffectedEntitiesOk returns a tuple with the AffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEntities

`func (o *VulnerableComponent) SetAffectedEntities(v []string)`

SetAffectedEntities sets AffectedEntities field to given value.

### HasAffectedEntities

`func (o *VulnerableComponent) HasAffectedEntities() bool`

HasAffectedEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


