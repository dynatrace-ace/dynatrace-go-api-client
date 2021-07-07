# MzPermissionsForGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Group ID | [optional] 
**MzPermissionsPerEnvironment** | Pointer to [**[]MzListForEnvironment**](MzListForEnvironment.md) | List of management zone permissions per environment | [optional] 

## Methods

### NewMzPermissionsForGroup

`func NewMzPermissionsForGroup() *MzPermissionsForGroup`

NewMzPermissionsForGroup instantiates a new MzPermissionsForGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMzPermissionsForGroupWithDefaults

`func NewMzPermissionsForGroupWithDefaults() *MzPermissionsForGroup`

NewMzPermissionsForGroupWithDefaults instantiates a new MzPermissionsForGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *MzPermissionsForGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *MzPermissionsForGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *MzPermissionsForGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *MzPermissionsForGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetMzPermissionsPerEnvironment

`func (o *MzPermissionsForGroup) GetMzPermissionsPerEnvironment() []MzListForEnvironment`

GetMzPermissionsPerEnvironment returns the MzPermissionsPerEnvironment field if non-nil, zero value otherwise.

### GetMzPermissionsPerEnvironmentOk

`func (o *MzPermissionsForGroup) GetMzPermissionsPerEnvironmentOk() (*[]MzListForEnvironment, bool)`

GetMzPermissionsPerEnvironmentOk returns a tuple with the MzPermissionsPerEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMzPermissionsPerEnvironment

`func (o *MzPermissionsForGroup) SetMzPermissionsPerEnvironment(v []MzListForEnvironment)`

SetMzPermissionsPerEnvironment sets MzPermissionsPerEnvironment field to given value.

### HasMzPermissionsPerEnvironment

`func (o *MzPermissionsForGroup) HasMzPermissionsPerEnvironment() bool`

HasMzPermissionsPerEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


