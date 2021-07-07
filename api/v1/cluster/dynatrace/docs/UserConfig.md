# UserConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User ID | 
**Email** | **string** | User&#39;s email address | 
**FirstName** | **string** | User&#39;s first name | 
**LastName** | **string** | User&#39;s last name | 
**PasswordClearText** | Pointer to **string** | User&#39;s password in a clear text; used only to set initial password | [optional] 
**Groups** | Pointer to **[]string** | List of user&#39;s user group IDs. | [optional] 

## Methods

### NewUserConfig

`func NewUserConfig(id string, email string, firstName string, lastName string, ) *UserConfig`

NewUserConfig instantiates a new UserConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserConfigWithDefaults

`func NewUserConfigWithDefaults() *UserConfig`

NewUserConfigWithDefaults instantiates a new UserConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserConfig) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *UserConfig) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserConfig) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserConfig) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *UserConfig) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserConfig) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserConfig) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UserConfig) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserConfig) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserConfig) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetPasswordClearText

`func (o *UserConfig) GetPasswordClearText() string`

GetPasswordClearText returns the PasswordClearText field if non-nil, zero value otherwise.

### GetPasswordClearTextOk

`func (o *UserConfig) GetPasswordClearTextOk() (*string, bool)`

GetPasswordClearTextOk returns a tuple with the PasswordClearText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordClearText

`func (o *UserConfig) SetPasswordClearText(v string)`

SetPasswordClearText sets PasswordClearText field to given value.

### HasPasswordClearText

`func (o *UserConfig) HasPasswordClearText() bool`

HasPasswordClearText returns a boolean if a field has been set.

### GetGroups

`func (o *UserConfig) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserConfig) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserConfig) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserConfig) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


