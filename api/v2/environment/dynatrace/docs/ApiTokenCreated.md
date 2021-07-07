# ApiTokenCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | The secret of the token. | [optional] 
**ExpirationDate** | Pointer to **string** | The token expiration date in ISO 8601 format (&#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;&#x60;). | [optional] 
**Id** | Pointer to **string** | The ID of the token, consisting of prefix and public part of the token. | [optional] 

## Methods

### NewApiTokenCreated

`func NewApiTokenCreated() *ApiTokenCreated`

NewApiTokenCreated instantiates a new ApiTokenCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokenCreatedWithDefaults

`func NewApiTokenCreatedWithDefaults() *ApiTokenCreated`

NewApiTokenCreatedWithDefaults instantiates a new ApiTokenCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ApiTokenCreated) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApiTokenCreated) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApiTokenCreated) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ApiTokenCreated) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ApiTokenCreated) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ApiTokenCreated) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ApiTokenCreated) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ApiTokenCreated) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetId

`func (o *ApiTokenCreated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiTokenCreated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiTokenCreated) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiTokenCreated) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


