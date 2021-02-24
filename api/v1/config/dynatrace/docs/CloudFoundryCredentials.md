# CloudFoundryCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the given credentials configuration. | [optional] [readonly] 
**Active** | Pointer to **bool** | The monitoring is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;) for given credentials configuration.   If not set on creation, the &#x60;true&#x60; value is used.   If the field is omitted during an update, the old value remains unaffected. | [optional] 
**EndpointStatus** | Pointer to **string** | The status of the configured endpoint.  ASSIGNED: The credentials are assigned to an ActiveGate which is responsible for processing. UNASSIGNED: The credentials are not yet assigned to an ActiveGate so there is currently no processing. DISABLED: The credentials have been disabled by the user. FASTCHECK_AUTH_ERROR: The credentials are invalid. FASTCHECK_TLS_ERROR: The endpoint TLS certificate is invalid. FASTCHECK_NO_RESPONSE: The endpoint did not return a result until the timeout was reached. FASTCHECK_INVALID_ENDPOINT: The endpoint URL was invalid. FASTCHECK_AUTH_LOCKED: The credentials seem to be locked. UNKNOWN: An unknown error occured.  | [optional] [readonly] 
**EndpointStatusInfo** | Pointer to **string** | The detailed status info of the configured endpoint. | [optional] [readonly] 
**Name** | **string** | The name of the Cloud Foundry foundation credentials.   Allowed characters are letters, numbers, whitespaces, and the following characters: &#x60;.+-_&#x60;. Leading or trailing whitespace is not allowed. | 
**ApiUrl** | **string** | The URL of the Cloud Foundry foundation credentials.   The URL must be valid according to RFC 2396.   Leading or trailing whitespaces are not allowed. | 
**LoginUrl** | **string** | The login URL of the Cloud Foundry foundation credentials.   The URL must be valid according to RFC 2396.   Leading or trailing whitespaces are not allowed. | 
**Username** | **string** | The username of the Cloud Foundry foundation credentials.   Leading and trailing whitespaces are not allowed. | 
**Password** | Pointer to **string** | The password of the Cloud Foundry foundation credentials. | [optional] 

## Methods

### NewCloudFoundryCredentials

`func NewCloudFoundryCredentials(name string, apiUrl string, loginUrl string, username string, ) *CloudFoundryCredentials`

NewCloudFoundryCredentials instantiates a new CloudFoundryCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFoundryCredentialsWithDefaults

`func NewCloudFoundryCredentialsWithDefaults() *CloudFoundryCredentials`

NewCloudFoundryCredentialsWithDefaults instantiates a new CloudFoundryCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CloudFoundryCredentials) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CloudFoundryCredentials) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CloudFoundryCredentials) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CloudFoundryCredentials) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *CloudFoundryCredentials) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudFoundryCredentials) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudFoundryCredentials) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudFoundryCredentials) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActive

`func (o *CloudFoundryCredentials) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CloudFoundryCredentials) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CloudFoundryCredentials) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CloudFoundryCredentials) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEndpointStatus

`func (o *CloudFoundryCredentials) GetEndpointStatus() string`

GetEndpointStatus returns the EndpointStatus field if non-nil, zero value otherwise.

### GetEndpointStatusOk

`func (o *CloudFoundryCredentials) GetEndpointStatusOk() (*string, bool)`

GetEndpointStatusOk returns a tuple with the EndpointStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointStatus

`func (o *CloudFoundryCredentials) SetEndpointStatus(v string)`

SetEndpointStatus sets EndpointStatus field to given value.

### HasEndpointStatus

`func (o *CloudFoundryCredentials) HasEndpointStatus() bool`

HasEndpointStatus returns a boolean if a field has been set.

### GetEndpointStatusInfo

`func (o *CloudFoundryCredentials) GetEndpointStatusInfo() string`

GetEndpointStatusInfo returns the EndpointStatusInfo field if non-nil, zero value otherwise.

### GetEndpointStatusInfoOk

`func (o *CloudFoundryCredentials) GetEndpointStatusInfoOk() (*string, bool)`

GetEndpointStatusInfoOk returns a tuple with the EndpointStatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointStatusInfo

`func (o *CloudFoundryCredentials) SetEndpointStatusInfo(v string)`

SetEndpointStatusInfo sets EndpointStatusInfo field to given value.

### HasEndpointStatusInfo

`func (o *CloudFoundryCredentials) HasEndpointStatusInfo() bool`

HasEndpointStatusInfo returns a boolean if a field has been set.

### GetName

`func (o *CloudFoundryCredentials) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudFoundryCredentials) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudFoundryCredentials) SetName(v string)`

SetName sets Name field to given value.


### GetApiUrl

`func (o *CloudFoundryCredentials) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *CloudFoundryCredentials) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *CloudFoundryCredentials) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.


### GetLoginUrl

`func (o *CloudFoundryCredentials) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *CloudFoundryCredentials) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *CloudFoundryCredentials) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.


### GetUsername

`func (o *CloudFoundryCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CloudFoundryCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CloudFoundryCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CloudFoundryCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CloudFoundryCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CloudFoundryCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CloudFoundryCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


