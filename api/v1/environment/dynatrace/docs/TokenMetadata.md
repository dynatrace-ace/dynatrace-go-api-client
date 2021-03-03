# TokenMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the token. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the token. | [optional] 
**UserId** | Pointer to **string** | The owner of the token. | [optional] [readonly] 
**Revoked** | Pointer to **bool** | Revocation status of the token. Revoked tokens are disabled. | [optional] 
**Created** | Pointer to **int64** | The creation time as a unix timestamp in milliseconds. | [optional] [readonly] 
**Expires** | Pointer to **int64** | The expiration time as a unix timestamp in milliseconds. | [optional] 
**LastUse** | Pointer to **int64** | The unix timestamp in milliseconds when the token was last used. | [optional] [readonly] 
**Scopes** | Pointer to **[]string** | The list of permissions, assigned to the token. | [optional] 
**PersonalAccessToken** | Pointer to **bool** | Specifies if the token is a personal access token. Personal access tokens are tied to the permissions of their owner. | [optional] [readonly] 

## Methods

### NewTokenMetadata

`func NewTokenMetadata() *TokenMetadata`

NewTokenMetadata instantiates a new TokenMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenMetadataWithDefaults

`func NewTokenMetadataWithDefaults() *TokenMetadata`

NewTokenMetadataWithDefaults instantiates a new TokenMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TokenMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenMetadata) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TokenMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TokenMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokenMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserId

`func (o *TokenMetadata) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TokenMetadata) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TokenMetadata) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TokenMetadata) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRevoked

`func (o *TokenMetadata) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *TokenMetadata) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *TokenMetadata) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *TokenMetadata) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetCreated

`func (o *TokenMetadata) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TokenMetadata) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TokenMetadata) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TokenMetadata) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpires

`func (o *TokenMetadata) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TokenMetadata) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TokenMetadata) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TokenMetadata) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetLastUse

`func (o *TokenMetadata) GetLastUse() int64`

GetLastUse returns the LastUse field if non-nil, zero value otherwise.

### GetLastUseOk

`func (o *TokenMetadata) GetLastUseOk() (*int64, bool)`

GetLastUseOk returns a tuple with the LastUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUse

`func (o *TokenMetadata) SetLastUse(v int64)`

SetLastUse sets LastUse field to given value.

### HasLastUse

`func (o *TokenMetadata) HasLastUse() bool`

HasLastUse returns a boolean if a field has been set.

### GetScopes

`func (o *TokenMetadata) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenMetadata) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenMetadata) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenMetadata) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetPersonalAccessToken

`func (o *TokenMetadata) GetPersonalAccessToken() bool`

GetPersonalAccessToken returns the PersonalAccessToken field if non-nil, zero value otherwise.

### GetPersonalAccessTokenOk

`func (o *TokenMetadata) GetPersonalAccessTokenOk() (*bool, bool)`

GetPersonalAccessTokenOk returns a tuple with the PersonalAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalAccessToken

`func (o *TokenMetadata) SetPersonalAccessToken(v bool)`

SetPersonalAccessToken sets PersonalAccessToken field to given value.

### HasPersonalAccessToken

`func (o *TokenMetadata) HasPersonalAccessToken() bool`

HasPersonalAccessToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


