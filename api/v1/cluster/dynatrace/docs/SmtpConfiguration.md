# SmtpConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostName** | Pointer to **string** | Host name | [optional] 
**Port** | Pointer to **int32** | Integer value of port | [optional] 
**UserName** | Pointer to **string** | User name | [optional] 
**Password** | Pointer to **string** | Password | [optional] 
**IsPasswordConfigured** | Pointer to **bool** | If true, a password has been configured. | [optional] 
**ConnectionSecurity** | Pointer to **string** | Connection security | [optional] 
**SenderEmailAddress** | Pointer to **string** | Sender email address | [optional] 
**AllowFallbackViaMissionControl** | Pointer to **bool** | If true, we will send e-mails via Mission Control in case of problems with SMTP server. | [optional] 
**UseSmtpServer** | Pointer to **bool** | If true, we will send e-mails via SMTP server. | [optional] 

## Methods

### NewSmtpConfiguration

`func NewSmtpConfiguration() *SmtpConfiguration`

NewSmtpConfiguration instantiates a new SmtpConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpConfigurationWithDefaults

`func NewSmtpConfigurationWithDefaults() *SmtpConfiguration`

NewSmtpConfigurationWithDefaults instantiates a new SmtpConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostName

`func (o *SmtpConfiguration) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *SmtpConfiguration) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *SmtpConfiguration) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *SmtpConfiguration) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetPort

`func (o *SmtpConfiguration) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SmtpConfiguration) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SmtpConfiguration) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SmtpConfiguration) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUserName

`func (o *SmtpConfiguration) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SmtpConfiguration) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SmtpConfiguration) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SmtpConfiguration) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *SmtpConfiguration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SmtpConfiguration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SmtpConfiguration) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SmtpConfiguration) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetIsPasswordConfigured

`func (o *SmtpConfiguration) GetIsPasswordConfigured() bool`

GetIsPasswordConfigured returns the IsPasswordConfigured field if non-nil, zero value otherwise.

### GetIsPasswordConfiguredOk

`func (o *SmtpConfiguration) GetIsPasswordConfiguredOk() (*bool, bool)`

GetIsPasswordConfiguredOk returns a tuple with the IsPasswordConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordConfigured

`func (o *SmtpConfiguration) SetIsPasswordConfigured(v bool)`

SetIsPasswordConfigured sets IsPasswordConfigured field to given value.

### HasIsPasswordConfigured

`func (o *SmtpConfiguration) HasIsPasswordConfigured() bool`

HasIsPasswordConfigured returns a boolean if a field has been set.

### GetConnectionSecurity

`func (o *SmtpConfiguration) GetConnectionSecurity() string`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *SmtpConfiguration) GetConnectionSecurityOk() (*string, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *SmtpConfiguration) SetConnectionSecurity(v string)`

SetConnectionSecurity sets ConnectionSecurity field to given value.

### HasConnectionSecurity

`func (o *SmtpConfiguration) HasConnectionSecurity() bool`

HasConnectionSecurity returns a boolean if a field has been set.

### GetSenderEmailAddress

`func (o *SmtpConfiguration) GetSenderEmailAddress() string`

GetSenderEmailAddress returns the SenderEmailAddress field if non-nil, zero value otherwise.

### GetSenderEmailAddressOk

`func (o *SmtpConfiguration) GetSenderEmailAddressOk() (*string, bool)`

GetSenderEmailAddressOk returns a tuple with the SenderEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderEmailAddress

`func (o *SmtpConfiguration) SetSenderEmailAddress(v string)`

SetSenderEmailAddress sets SenderEmailAddress field to given value.

### HasSenderEmailAddress

`func (o *SmtpConfiguration) HasSenderEmailAddress() bool`

HasSenderEmailAddress returns a boolean if a field has been set.

### GetAllowFallbackViaMissionControl

`func (o *SmtpConfiguration) GetAllowFallbackViaMissionControl() bool`

GetAllowFallbackViaMissionControl returns the AllowFallbackViaMissionControl field if non-nil, zero value otherwise.

### GetAllowFallbackViaMissionControlOk

`func (o *SmtpConfiguration) GetAllowFallbackViaMissionControlOk() (*bool, bool)`

GetAllowFallbackViaMissionControlOk returns a tuple with the AllowFallbackViaMissionControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFallbackViaMissionControl

`func (o *SmtpConfiguration) SetAllowFallbackViaMissionControl(v bool)`

SetAllowFallbackViaMissionControl sets AllowFallbackViaMissionControl field to given value.

### HasAllowFallbackViaMissionControl

`func (o *SmtpConfiguration) HasAllowFallbackViaMissionControl() bool`

HasAllowFallbackViaMissionControl returns a boolean if a field has been set.

### GetUseSmtpServer

`func (o *SmtpConfiguration) GetUseSmtpServer() bool`

GetUseSmtpServer returns the UseSmtpServer field if non-nil, zero value otherwise.

### GetUseSmtpServerOk

`func (o *SmtpConfiguration) GetUseSmtpServerOk() (*bool, bool)`

GetUseSmtpServerOk returns a tuple with the UseSmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSmtpServer

`func (o *SmtpConfiguration) SetUseSmtpServer(v bool)`

SetUseSmtpServer sets UseSmtpServer field to given value.

### HasUseSmtpServer

`func (o *SmtpConfiguration) HasUseSmtpServer() bool`

HasUseSmtpServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


