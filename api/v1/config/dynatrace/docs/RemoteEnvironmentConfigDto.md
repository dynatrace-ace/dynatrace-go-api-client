# RemoteEnvironmentConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the configuration. | [optional] 
**DisplayName** | **string** | The display name of the remote environment. | 
**Uri** | **string** | The URI of the remote environment. | 
**Token** | Pointer to **string** | The API token granting access to the remote environment.   The token must have the **Fetch data from a remote environment** (&#x60;RestRequestForwarding&#x60;) scope.   For security reasons, GET requests return this field as &#x60;null&#x60;. | [optional] 
**NetworkScope** | Pointer to **string** | The network scope of the remote environment: * &#x60;EXTERNAL&#x60;: The remote environment is located in an another network.  * &#x60;INTERNAL&#x60;: The remote environment is located in the same network.  * &#x60;CLUSTER&#x60;: The remote environment is located in the same cluster.   Dynatrace SaaS can only use &#x60;EXTERNAL&#x60;.  If not set, &#x60;EXTERNAL&#x60; is used. | [optional] 

## Methods

### NewRemoteEnvironmentConfigDto

`func NewRemoteEnvironmentConfigDto(displayName string, uri string, ) *RemoteEnvironmentConfigDto`

NewRemoteEnvironmentConfigDto instantiates a new RemoteEnvironmentConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteEnvironmentConfigDtoWithDefaults

`func NewRemoteEnvironmentConfigDtoWithDefaults() *RemoteEnvironmentConfigDto`

NewRemoteEnvironmentConfigDtoWithDefaults instantiates a new RemoteEnvironmentConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemoteEnvironmentConfigDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteEnvironmentConfigDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteEnvironmentConfigDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteEnvironmentConfigDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *RemoteEnvironmentConfigDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RemoteEnvironmentConfigDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RemoteEnvironmentConfigDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetUri

`func (o *RemoteEnvironmentConfigDto) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *RemoteEnvironmentConfigDto) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *RemoteEnvironmentConfigDto) SetUri(v string)`

SetUri sets Uri field to given value.


### GetToken

`func (o *RemoteEnvironmentConfigDto) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RemoteEnvironmentConfigDto) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RemoteEnvironmentConfigDto) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RemoteEnvironmentConfigDto) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetNetworkScope

`func (o *RemoteEnvironmentConfigDto) GetNetworkScope() string`

GetNetworkScope returns the NetworkScope field if non-nil, zero value otherwise.

### GetNetworkScopeOk

`func (o *RemoteEnvironmentConfigDto) GetNetworkScopeOk() (*string, bool)`

GetNetworkScopeOk returns a tuple with the NetworkScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkScope

`func (o *RemoteEnvironmentConfigDto) SetNetworkScope(v string)`

SetNetworkScope sets NetworkScope field to given value.

### HasNetworkScope

`func (o *RemoteEnvironmentConfigDto) HasNetworkScope() bool`

HasNetworkScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


