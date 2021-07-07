# AuthenticationMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationProvider** | Pointer to **string** |  | [optional] 
**AuthorizationProvider** | Pointer to **string** |  | [optional] 
**SsoOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuthenticationMode

`func NewAuthenticationMode() *AuthenticationMode`

NewAuthenticationMode instantiates a new AuthenticationMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationModeWithDefaults

`func NewAuthenticationModeWithDefaults() *AuthenticationMode`

NewAuthenticationModeWithDefaults instantiates a new AuthenticationMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationProvider

`func (o *AuthenticationMode) GetAuthenticationProvider() string`

GetAuthenticationProvider returns the AuthenticationProvider field if non-nil, zero value otherwise.

### GetAuthenticationProviderOk

`func (o *AuthenticationMode) GetAuthenticationProviderOk() (*string, bool)`

GetAuthenticationProviderOk returns a tuple with the AuthenticationProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProvider

`func (o *AuthenticationMode) SetAuthenticationProvider(v string)`

SetAuthenticationProvider sets AuthenticationProvider field to given value.

### HasAuthenticationProvider

`func (o *AuthenticationMode) HasAuthenticationProvider() bool`

HasAuthenticationProvider returns a boolean if a field has been set.

### GetAuthorizationProvider

`func (o *AuthenticationMode) GetAuthorizationProvider() string`

GetAuthorizationProvider returns the AuthorizationProvider field if non-nil, zero value otherwise.

### GetAuthorizationProviderOk

`func (o *AuthenticationMode) GetAuthorizationProviderOk() (*string, bool)`

GetAuthorizationProviderOk returns a tuple with the AuthorizationProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationProvider

`func (o *AuthenticationMode) SetAuthorizationProvider(v string)`

SetAuthorizationProvider sets AuthorizationProvider field to given value.

### HasAuthorizationProvider

`func (o *AuthenticationMode) HasAuthorizationProvider() bool`

HasAuthorizationProvider returns a boolean if a field has been set.

### GetSsoOnly

`func (o *AuthenticationMode) GetSsoOnly() bool`

GetSsoOnly returns the SsoOnly field if non-nil, zero value otherwise.

### GetSsoOnlyOk

`func (o *AuthenticationMode) GetSsoOnlyOk() (*bool, bool)`

GetSsoOnlyOk returns a tuple with the SsoOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoOnly

`func (o *AuthenticationMode) SetSsoOnly(v bool)`

SetSsoOnly sets SsoOnly field to given value.

### HasSsoOnly

`func (o *AuthenticationMode) HasSsoOnly() bool`

HasSsoOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


