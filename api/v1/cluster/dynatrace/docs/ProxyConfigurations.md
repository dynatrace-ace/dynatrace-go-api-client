# ProxyConfigurations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configurations** | Pointer to [**map[string]InternetProxy**](InternetProxy.md) | Map of (Data Center name, Proxy) pairs | [optional] 

## Methods

### NewProxyConfigurations

`func NewProxyConfigurations() *ProxyConfigurations`

NewProxyConfigurations instantiates a new ProxyConfigurations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyConfigurationsWithDefaults

`func NewProxyConfigurationsWithDefaults() *ProxyConfigurations`

NewProxyConfigurationsWithDefaults instantiates a new ProxyConfigurations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurations

`func (o *ProxyConfigurations) GetConfigurations() map[string]InternetProxy`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *ProxyConfigurations) GetConfigurationsOk() (*map[string]InternetProxy, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *ProxyConfigurations) SetConfigurations(v map[string]InternetProxy)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *ProxyConfigurations) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


