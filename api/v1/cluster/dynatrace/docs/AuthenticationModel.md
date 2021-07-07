# AuthenticationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationProvider** | **string** | Defines the authentication provider | 
**AuthorizationProvider** | Pointer to **string** | Defines the authorization provider | [optional] 
**SsoOnly** | Pointer to **bool** | Defines login page type as SSO | [optional] [readonly] 

## Methods

### NewAuthenticationModel

`func NewAuthenticationModel(authenticationProvider string, ) *AuthenticationModel`

NewAuthenticationModel instantiates a new AuthenticationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationModelWithDefaults

`func NewAuthenticationModelWithDefaults() *AuthenticationModel`

NewAuthenticationModelWithDefaults instantiates a new AuthenticationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationProvider

`func (o *AuthenticationModel) GetAuthenticationProvider() string`

GetAuthenticationProvider returns the AuthenticationProvider field if non-nil, zero value otherwise.

### GetAuthenticationProviderOk

`func (o *AuthenticationModel) GetAuthenticationProviderOk() (*string, bool)`

GetAuthenticationProviderOk returns a tuple with the AuthenticationProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProvider

`func (o *AuthenticationModel) SetAuthenticationProvider(v string)`

SetAuthenticationProvider sets AuthenticationProvider field to given value.


### GetAuthorizationProvider

`func (o *AuthenticationModel) GetAuthorizationProvider() string`

GetAuthorizationProvider returns the AuthorizationProvider field if non-nil, zero value otherwise.

### GetAuthorizationProviderOk

`func (o *AuthenticationModel) GetAuthorizationProviderOk() (*string, bool)`

GetAuthorizationProviderOk returns a tuple with the AuthorizationProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationProvider

`func (o *AuthenticationModel) SetAuthorizationProvider(v string)`

SetAuthorizationProvider sets AuthorizationProvider field to given value.

### HasAuthorizationProvider

`func (o *AuthenticationModel) HasAuthorizationProvider() bool`

HasAuthorizationProvider returns a boolean if a field has been set.

### GetSsoOnly

`func (o *AuthenticationModel) GetSsoOnly() bool`

GetSsoOnly returns the SsoOnly field if non-nil, zero value otherwise.

### GetSsoOnlyOk

`func (o *AuthenticationModel) GetSsoOnlyOk() (*bool, bool)`

GetSsoOnlyOk returns a tuple with the SsoOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoOnly

`func (o *AuthenticationModel) SetSsoOnly(v bool)`

SetSsoOnly sets SsoOnly field to given value.

### HasSsoOnly

`func (o *AuthenticationModel) HasSsoOnly() bool`

HasSsoOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


