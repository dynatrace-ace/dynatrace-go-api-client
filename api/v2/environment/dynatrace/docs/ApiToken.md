# ApiToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersonalAccessToken** | Pointer to **bool** | The token is a [personal access token](https://dt-url.net/wm03sop) (&#x60;true&#x60;) or an API token (&#x60;false&#x60;). | [optional] 
**LastUsedDate** | Pointer to **string** | Token last used date in ISO 8601 format (&#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;&#x60;) | [optional] 
**LastUsedIpAddress** | Pointer to **string** | Token last used IP address. | [optional] 
**ExpirationDate** | Pointer to **string** | Token expiration date in ISO 8601 format (&#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;&#x60;) | [optional] 
**CreationDate** | Pointer to **string** | Token creation date in ISO 8601 format (&#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;&#x60;) | [optional] 
**Scopes** | Pointer to **[]string** | A list of scopes assigned to the token. | [optional] 
**Name** | Pointer to **string** | The name of the token. | [optional] 
**Id** | Pointer to **string** | The ID of the token, consisting of prefix and public part of the token. | [optional] 
**Owner** | Pointer to **string** | The owner of the token. | [optional] 
**Enabled** | Pointer to **bool** | The token is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 

## Methods

### NewApiToken

`func NewApiToken() *ApiToken`

NewApiToken instantiates a new ApiToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokenWithDefaults

`func NewApiTokenWithDefaults() *ApiToken`

NewApiTokenWithDefaults instantiates a new ApiToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersonalAccessToken

`func (o *ApiToken) GetPersonalAccessToken() bool`

GetPersonalAccessToken returns the PersonalAccessToken field if non-nil, zero value otherwise.

### GetPersonalAccessTokenOk

`func (o *ApiToken) GetPersonalAccessTokenOk() (*bool, bool)`

GetPersonalAccessTokenOk returns a tuple with the PersonalAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalAccessToken

`func (o *ApiToken) SetPersonalAccessToken(v bool)`

SetPersonalAccessToken sets PersonalAccessToken field to given value.

### HasPersonalAccessToken

`func (o *ApiToken) HasPersonalAccessToken() bool`

HasPersonalAccessToken returns a boolean if a field has been set.

### GetLastUsedDate

`func (o *ApiToken) GetLastUsedDate() string`

GetLastUsedDate returns the LastUsedDate field if non-nil, zero value otherwise.

### GetLastUsedDateOk

`func (o *ApiToken) GetLastUsedDateOk() (*string, bool)`

GetLastUsedDateOk returns a tuple with the LastUsedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedDate

`func (o *ApiToken) SetLastUsedDate(v string)`

SetLastUsedDate sets LastUsedDate field to given value.

### HasLastUsedDate

`func (o *ApiToken) HasLastUsedDate() bool`

HasLastUsedDate returns a boolean if a field has been set.

### GetLastUsedIpAddress

`func (o *ApiToken) GetLastUsedIpAddress() string`

GetLastUsedIpAddress returns the LastUsedIpAddress field if non-nil, zero value otherwise.

### GetLastUsedIpAddressOk

`func (o *ApiToken) GetLastUsedIpAddressOk() (*string, bool)`

GetLastUsedIpAddressOk returns a tuple with the LastUsedIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedIpAddress

`func (o *ApiToken) SetLastUsedIpAddress(v string)`

SetLastUsedIpAddress sets LastUsedIpAddress field to given value.

### HasLastUsedIpAddress

`func (o *ApiToken) HasLastUsedIpAddress() bool`

HasLastUsedIpAddress returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ApiToken) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ApiToken) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ApiToken) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ApiToken) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetCreationDate

`func (o *ApiToken) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ApiToken) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ApiToken) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ApiToken) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetScopes

`func (o *ApiToken) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApiToken) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApiToken) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApiToken) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetName

`func (o *ApiToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ApiToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwner

`func (o *ApiToken) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ApiToken) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ApiToken) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ApiToken) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiToken) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiToken) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiToken) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiToken) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


