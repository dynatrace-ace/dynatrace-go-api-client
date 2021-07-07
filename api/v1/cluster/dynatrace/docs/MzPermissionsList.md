# MzPermissionsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MzId** | Pointer to **string** | The ID of the required management zone | [optional] 
**Permissions** | Pointer to **[]string** | The list of permissions for the required management zone | [optional] 

## Methods

### NewMzPermissionsList

`func NewMzPermissionsList() *MzPermissionsList`

NewMzPermissionsList instantiates a new MzPermissionsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMzPermissionsListWithDefaults

`func NewMzPermissionsListWithDefaults() *MzPermissionsList`

NewMzPermissionsListWithDefaults instantiates a new MzPermissionsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMzId

`func (o *MzPermissionsList) GetMzId() string`

GetMzId returns the MzId field if non-nil, zero value otherwise.

### GetMzIdOk

`func (o *MzPermissionsList) GetMzIdOk() (*string, bool)`

GetMzIdOk returns a tuple with the MzId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMzId

`func (o *MzPermissionsList) SetMzId(v string)`

SetMzId sets MzId field to given value.

### HasMzId

`func (o *MzPermissionsList) HasMzId() bool`

HasMzId returns a boolean if a field has been set.

### GetPermissions

`func (o *MzPermissionsList) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *MzPermissionsList) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *MzPermissionsList) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *MzPermissionsList) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


