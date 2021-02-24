# CredentialsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**[]CredentialsResponseElement**](CredentialsResponseElement.md) | A list of credentials sets for Synthetic monitors. | 

## Methods

### NewCredentialsList

`func NewCredentialsList(credentials []CredentialsResponseElement, ) *CredentialsList`

NewCredentialsList instantiates a new CredentialsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsListWithDefaults

`func NewCredentialsListWithDefaults() *CredentialsList`

NewCredentialsListWithDefaults instantiates a new CredentialsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *CredentialsList) GetCredentials() []CredentialsResponseElement`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CredentialsList) GetCredentialsOk() (*[]CredentialsResponseElement, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CredentialsList) SetCredentials(v []CredentialsResponseElement)`

SetCredentials sets Credentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


