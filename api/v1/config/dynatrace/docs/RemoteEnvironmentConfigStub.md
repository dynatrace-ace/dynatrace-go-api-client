# RemoteEnvironmentConfigStub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Uri** | Pointer to **string** | The display name of the remote environment. | [optional] 
**NetworkScope** | Pointer to **string** | The network scope of the remote environment: * &#x60;EXTERNAL&#x60;: The remote environment is located in an another network.  * &#x60;INTERNAL&#x60;: The remote environment is located in the same network.  * &#x60;CLUSTER&#x60;: The remote environment is located in the same cluster.   Dynatrace SaaS can only use &#x60;EXTERNAL&#x60;.  If not set, &#x60;EXTERNAL&#x60; is used. | [optional] 

## Methods

### NewRemoteEnvironmentConfigStub

`func NewRemoteEnvironmentConfigStub(id string, ) *RemoteEnvironmentConfigStub`

NewRemoteEnvironmentConfigStub instantiates a new RemoteEnvironmentConfigStub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteEnvironmentConfigStubWithDefaults

`func NewRemoteEnvironmentConfigStubWithDefaults() *RemoteEnvironmentConfigStub`

NewRemoteEnvironmentConfigStubWithDefaults instantiates a new RemoteEnvironmentConfigStub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RemoteEnvironmentConfigStub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteEnvironmentConfigStub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteEnvironmentConfigStub) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemoteEnvironmentConfigStub) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *RemoteEnvironmentConfigStub) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteEnvironmentConfigStub) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteEnvironmentConfigStub) SetId(v string)`

SetId sets Id field to given value.


### GetUri

`func (o *RemoteEnvironmentConfigStub) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *RemoteEnvironmentConfigStub) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *RemoteEnvironmentConfigStub) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *RemoteEnvironmentConfigStub) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetNetworkScope

`func (o *RemoteEnvironmentConfigStub) GetNetworkScope() string`

GetNetworkScope returns the NetworkScope field if non-nil, zero value otherwise.

### GetNetworkScopeOk

`func (o *RemoteEnvironmentConfigStub) GetNetworkScopeOk() (*string, bool)`

GetNetworkScopeOk returns a tuple with the NetworkScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkScope

`func (o *RemoteEnvironmentConfigStub) SetNetworkScope(v string)`

SetNetworkScope sets NetworkScope field to given value.

### HasNetworkScope

`func (o *RemoteEnvironmentConfigStub) HasNetworkScope() bool`

HasNetworkScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


