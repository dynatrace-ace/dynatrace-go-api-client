# AutomaticLogoutConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogoutInactiveUsersEnabled** | **bool** | True if automatic logout is enabled | 
**UserInactivityTimeout** | **int32** | User inactivity timeout | 

## Methods

### NewAutomaticLogoutConfiguration

`func NewAutomaticLogoutConfiguration(logoutInactiveUsersEnabled bool, userInactivityTimeout int32, ) *AutomaticLogoutConfiguration`

NewAutomaticLogoutConfiguration instantiates a new AutomaticLogoutConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomaticLogoutConfigurationWithDefaults

`func NewAutomaticLogoutConfigurationWithDefaults() *AutomaticLogoutConfiguration`

NewAutomaticLogoutConfigurationWithDefaults instantiates a new AutomaticLogoutConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogoutInactiveUsersEnabled

`func (o *AutomaticLogoutConfiguration) GetLogoutInactiveUsersEnabled() bool`

GetLogoutInactiveUsersEnabled returns the LogoutInactiveUsersEnabled field if non-nil, zero value otherwise.

### GetLogoutInactiveUsersEnabledOk

`func (o *AutomaticLogoutConfiguration) GetLogoutInactiveUsersEnabledOk() (*bool, bool)`

GetLogoutInactiveUsersEnabledOk returns a tuple with the LogoutInactiveUsersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutInactiveUsersEnabled

`func (o *AutomaticLogoutConfiguration) SetLogoutInactiveUsersEnabled(v bool)`

SetLogoutInactiveUsersEnabled sets LogoutInactiveUsersEnabled field to given value.


### GetUserInactivityTimeout

`func (o *AutomaticLogoutConfiguration) GetUserInactivityTimeout() int32`

GetUserInactivityTimeout returns the UserInactivityTimeout field if non-nil, zero value otherwise.

### GetUserInactivityTimeoutOk

`func (o *AutomaticLogoutConfiguration) GetUserInactivityTimeoutOk() (*int32, bool)`

GetUserInactivityTimeoutOk returns a tuple with the UserInactivityTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInactivityTimeout

`func (o *AutomaticLogoutConfiguration) SetUserInactivityTimeout(v int32)`

SetUserInactivityTimeout sets UserInactivityTimeout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


