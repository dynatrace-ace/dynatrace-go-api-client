# InternetProxyChangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scheme** | **string** | Protocol which proxy server uses | 
**Server** | **string** | Address (either IP or Hostname) of proxy server | 
**Port** | **int32** | Port of proxy server | 
**User** | Pointer to **string** | User of proxy server, null means do not change previous value | [optional] 
**Password** | Pointer to **string** | Password of proxy server, null means do not change previous value | [optional] 
**NonProxyHosts** | Pointer to **[]string** | Definition of hosts for which proxy won&#39;t be used. You can define multiple hosts. Each host can start or end with wildcard &#39;*&#39; for instance to match whole domain. | [optional] 

## Methods

### NewInternetProxyChangeRequest

`func NewInternetProxyChangeRequest(scheme string, server string, port int32, ) *InternetProxyChangeRequest`

NewInternetProxyChangeRequest instantiates a new InternetProxyChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetProxyChangeRequestWithDefaults

`func NewInternetProxyChangeRequestWithDefaults() *InternetProxyChangeRequest`

NewInternetProxyChangeRequestWithDefaults instantiates a new InternetProxyChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheme

`func (o *InternetProxyChangeRequest) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *InternetProxyChangeRequest) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *InternetProxyChangeRequest) SetScheme(v string)`

SetScheme sets Scheme field to given value.


### GetServer

`func (o *InternetProxyChangeRequest) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *InternetProxyChangeRequest) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *InternetProxyChangeRequest) SetServer(v string)`

SetServer sets Server field to given value.


### GetPort

`func (o *InternetProxyChangeRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InternetProxyChangeRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InternetProxyChangeRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUser

`func (o *InternetProxyChangeRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *InternetProxyChangeRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *InternetProxyChangeRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *InternetProxyChangeRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPassword

`func (o *InternetProxyChangeRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InternetProxyChangeRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InternetProxyChangeRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InternetProxyChangeRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetNonProxyHosts

`func (o *InternetProxyChangeRequest) GetNonProxyHosts() []string`

GetNonProxyHosts returns the NonProxyHosts field if non-nil, zero value otherwise.

### GetNonProxyHostsOk

`func (o *InternetProxyChangeRequest) GetNonProxyHostsOk() (*[]string, bool)`

GetNonProxyHostsOk returns a tuple with the NonProxyHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonProxyHosts

`func (o *InternetProxyChangeRequest) SetNonProxyHosts(v []string)`

SetNonProxyHosts sets NonProxyHosts field to given value.

### HasNonProxyHosts

`func (o *InternetProxyChangeRequest) HasNonProxyHosts() bool`

HasNonProxyHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


