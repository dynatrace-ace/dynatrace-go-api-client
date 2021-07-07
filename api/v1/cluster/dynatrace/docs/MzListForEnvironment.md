# MzListForEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentUuid** | Pointer to **string** | Environment UUID | [optional] 
**MzPermissions** | Pointer to [**[]MzPermissionsList**](MzPermissionsList.md) | List of management zone models with permissions | [optional] 

## Methods

### NewMzListForEnvironment

`func NewMzListForEnvironment() *MzListForEnvironment`

NewMzListForEnvironment instantiates a new MzListForEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMzListForEnvironmentWithDefaults

`func NewMzListForEnvironmentWithDefaults() *MzListForEnvironment`

NewMzListForEnvironmentWithDefaults instantiates a new MzListForEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentUuid

`func (o *MzListForEnvironment) GetEnvironmentUuid() string`

GetEnvironmentUuid returns the EnvironmentUuid field if non-nil, zero value otherwise.

### GetEnvironmentUuidOk

`func (o *MzListForEnvironment) GetEnvironmentUuidOk() (*string, bool)`

GetEnvironmentUuidOk returns a tuple with the EnvironmentUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUuid

`func (o *MzListForEnvironment) SetEnvironmentUuid(v string)`

SetEnvironmentUuid sets EnvironmentUuid field to given value.

### HasEnvironmentUuid

`func (o *MzListForEnvironment) HasEnvironmentUuid() bool`

HasEnvironmentUuid returns a boolean if a field has been set.

### GetMzPermissions

`func (o *MzListForEnvironment) GetMzPermissions() []MzPermissionsList`

GetMzPermissions returns the MzPermissions field if non-nil, zero value otherwise.

### GetMzPermissionsOk

`func (o *MzListForEnvironment) GetMzPermissionsOk() (*[]MzPermissionsList, bool)`

GetMzPermissionsOk returns a tuple with the MzPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMzPermissions

`func (o *MzListForEnvironment) SetMzPermissions(v []MzPermissionsList)`

SetMzPermissions sets MzPermissions field to given value.

### HasMzPermissions

`func (o *MzListForEnvironment) HasMzPermissions() bool`

HasMzPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


