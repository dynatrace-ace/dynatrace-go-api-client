# InternetProxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scheme** | Pointer to **string** | Protocol which proxy server uses | [optional] 
**Server** | Pointer to **string** | Address (either IP or Hostname) of proxy server | [optional] 
**Port** | Pointer to **int32** | Port of proxy server | [optional] 
**NonProxyHosts** | Pointer to **[]string** | Definition of hosts for which proxy won&#39;t be used. | [optional] 
**UserOrPasswordDefined** | Pointer to **bool** | Indicates if user/password for proxy is configured | [optional] 

## Methods

### NewInternetProxy

`func NewInternetProxy() *InternetProxy`

NewInternetProxy instantiates a new InternetProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetProxyWithDefaults

`func NewInternetProxyWithDefaults() *InternetProxy`

NewInternetProxyWithDefaults instantiates a new InternetProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheme

`func (o *InternetProxy) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *InternetProxy) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *InternetProxy) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *InternetProxy) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetServer

`func (o *InternetProxy) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *InternetProxy) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *InternetProxy) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *InternetProxy) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetPort

`func (o *InternetProxy) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InternetProxy) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InternetProxy) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *InternetProxy) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetNonProxyHosts

`func (o *InternetProxy) GetNonProxyHosts() []string`

GetNonProxyHosts returns the NonProxyHosts field if non-nil, zero value otherwise.

### GetNonProxyHostsOk

`func (o *InternetProxy) GetNonProxyHostsOk() (*[]string, bool)`

GetNonProxyHostsOk returns a tuple with the NonProxyHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonProxyHosts

`func (o *InternetProxy) SetNonProxyHosts(v []string)`

SetNonProxyHosts sets NonProxyHosts field to given value.

### HasNonProxyHosts

`func (o *InternetProxy) HasNonProxyHosts() bool`

HasNonProxyHosts returns a boolean if a field has been set.

### GetUserOrPasswordDefined

`func (o *InternetProxy) GetUserOrPasswordDefined() bool`

GetUserOrPasswordDefined returns the UserOrPasswordDefined field if non-nil, zero value otherwise.

### GetUserOrPasswordDefinedOk

`func (o *InternetProxy) GetUserOrPasswordDefinedOk() (*bool, bool)`

GetUserOrPasswordDefinedOk returns a tuple with the UserOrPasswordDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOrPasswordDefined

`func (o *InternetProxy) SetUserOrPasswordDefined(v bool)`

SetUserOrPasswordDefined sets UserOrPasswordDefined field to given value.

### HasUserOrPasswordDefined

`func (o *InternetProxy) HasUserOrPasswordDefined() bool`

HasUserOrPasswordDefined returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


