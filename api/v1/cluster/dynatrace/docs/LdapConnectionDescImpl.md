# LdapConnectionDescImpl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShouldUseSecureConnection** | Pointer to **bool** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**BindDn** | Pointer to **string** |  | [optional] 
**BindPassword** | Pointer to **string** |  | [optional] 
**BindPasswordSet** | Pointer to **bool** |  | [optional] 
**ReferralHopLimit** | Pointer to **int32** |  | [optional] 

## Methods

### NewLdapConnectionDescImpl

`func NewLdapConnectionDescImpl() *LdapConnectionDescImpl`

NewLdapConnectionDescImpl instantiates a new LdapConnectionDescImpl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapConnectionDescImplWithDefaults

`func NewLdapConnectionDescImplWithDefaults() *LdapConnectionDescImpl`

NewLdapConnectionDescImplWithDefaults instantiates a new LdapConnectionDescImpl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShouldUseSecureConnection

`func (o *LdapConnectionDescImpl) GetShouldUseSecureConnection() bool`

GetShouldUseSecureConnection returns the ShouldUseSecureConnection field if non-nil, zero value otherwise.

### GetShouldUseSecureConnectionOk

`func (o *LdapConnectionDescImpl) GetShouldUseSecureConnectionOk() (*bool, bool)`

GetShouldUseSecureConnectionOk returns a tuple with the ShouldUseSecureConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldUseSecureConnection

`func (o *LdapConnectionDescImpl) SetShouldUseSecureConnection(v bool)`

SetShouldUseSecureConnection sets ShouldUseSecureConnection field to given value.

### HasShouldUseSecureConnection

`func (o *LdapConnectionDescImpl) HasShouldUseSecureConnection() bool`

HasShouldUseSecureConnection returns a boolean if a field has been set.

### GetHostName

`func (o *LdapConnectionDescImpl) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *LdapConnectionDescImpl) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *LdapConnectionDescImpl) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *LdapConnectionDescImpl) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetPort

`func (o *LdapConnectionDescImpl) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LdapConnectionDescImpl) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LdapConnectionDescImpl) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LdapConnectionDescImpl) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetBindDn

`func (o *LdapConnectionDescImpl) GetBindDn() string`

GetBindDn returns the BindDn field if non-nil, zero value otherwise.

### GetBindDnOk

`func (o *LdapConnectionDescImpl) GetBindDnOk() (*string, bool)`

GetBindDnOk returns a tuple with the BindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDn

`func (o *LdapConnectionDescImpl) SetBindDn(v string)`

SetBindDn sets BindDn field to given value.

### HasBindDn

`func (o *LdapConnectionDescImpl) HasBindDn() bool`

HasBindDn returns a boolean if a field has been set.

### GetBindPassword

`func (o *LdapConnectionDescImpl) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *LdapConnectionDescImpl) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *LdapConnectionDescImpl) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.

### HasBindPassword

`func (o *LdapConnectionDescImpl) HasBindPassword() bool`

HasBindPassword returns a boolean if a field has been set.

### GetBindPasswordSet

`func (o *LdapConnectionDescImpl) GetBindPasswordSet() bool`

GetBindPasswordSet returns the BindPasswordSet field if non-nil, zero value otherwise.

### GetBindPasswordSetOk

`func (o *LdapConnectionDescImpl) GetBindPasswordSetOk() (*bool, bool)`

GetBindPasswordSetOk returns a tuple with the BindPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPasswordSet

`func (o *LdapConnectionDescImpl) SetBindPasswordSet(v bool)`

SetBindPasswordSet sets BindPasswordSet field to given value.

### HasBindPasswordSet

`func (o *LdapConnectionDescImpl) HasBindPasswordSet() bool`

HasBindPasswordSet returns a boolean if a field has been set.

### GetReferralHopLimit

`func (o *LdapConnectionDescImpl) GetReferralHopLimit() int32`

GetReferralHopLimit returns the ReferralHopLimit field if non-nil, zero value otherwise.

### GetReferralHopLimitOk

`func (o *LdapConnectionDescImpl) GetReferralHopLimitOk() (*int32, bool)`

GetReferralHopLimitOk returns a tuple with the ReferralHopLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralHopLimit

`func (o *LdapConnectionDescImpl) SetReferralHopLimit(v int32)`

SetReferralHopLimit sets ReferralHopLimit field to given value.

### HasReferralHopLimit

`func (o *LdapConnectionDescImpl) HasReferralHopLimit() bool`

HasReferralHopLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


