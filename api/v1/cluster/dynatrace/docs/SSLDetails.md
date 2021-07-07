# SSLDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InProgress** | Pointer to **bool** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**RestartRequired** | Pointer to **bool** |  | [optional] 
**CustomKeyStore** | Pointer to **bool** |  | [optional] 
**CustomKeyStoreWritable** | Pointer to **bool** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 

## Methods

### NewSSLDetails

`func NewSSLDetails() *SSLDetails`

NewSSLDetails instantiates a new SSLDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSLDetailsWithDefaults

`func NewSSLDetailsWithDefaults() *SSLDetails`

NewSSLDetailsWithDefaults instantiates a new SSLDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInProgress

`func (o *SSLDetails) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *SSLDetails) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *SSLDetails) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *SSLDetails) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetExpirationDate

`func (o *SSLDetails) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *SSLDetails) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *SSLDetails) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *SSLDetails) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetRestartRequired

`func (o *SSLDetails) GetRestartRequired() bool`

GetRestartRequired returns the RestartRequired field if non-nil, zero value otherwise.

### GetRestartRequiredOk

`func (o *SSLDetails) GetRestartRequiredOk() (*bool, bool)`

GetRestartRequiredOk returns a tuple with the RestartRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartRequired

`func (o *SSLDetails) SetRestartRequired(v bool)`

SetRestartRequired sets RestartRequired field to given value.

### HasRestartRequired

`func (o *SSLDetails) HasRestartRequired() bool`

HasRestartRequired returns a boolean if a field has been set.

### GetCustomKeyStore

`func (o *SSLDetails) GetCustomKeyStore() bool`

GetCustomKeyStore returns the CustomKeyStore field if non-nil, zero value otherwise.

### GetCustomKeyStoreOk

`func (o *SSLDetails) GetCustomKeyStoreOk() (*bool, bool)`

GetCustomKeyStoreOk returns a tuple with the CustomKeyStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomKeyStore

`func (o *SSLDetails) SetCustomKeyStore(v bool)`

SetCustomKeyStore sets CustomKeyStore field to given value.

### HasCustomKeyStore

`func (o *SSLDetails) HasCustomKeyStore() bool`

HasCustomKeyStore returns a boolean if a field has been set.

### GetCustomKeyStoreWritable

`func (o *SSLDetails) GetCustomKeyStoreWritable() bool`

GetCustomKeyStoreWritable returns the CustomKeyStoreWritable field if non-nil, zero value otherwise.

### GetCustomKeyStoreWritableOk

`func (o *SSLDetails) GetCustomKeyStoreWritableOk() (*bool, bool)`

GetCustomKeyStoreWritableOk returns a tuple with the CustomKeyStoreWritable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomKeyStoreWritable

`func (o *SSLDetails) SetCustomKeyStoreWritable(v bool)`

SetCustomKeyStoreWritable sets CustomKeyStoreWritable field to given value.

### HasCustomKeyStoreWritable

`func (o *SSLDetails) HasCustomKeyStoreWritable() bool`

HasCustomKeyStoreWritable returns a boolean if a field has been set.

### GetIssuer

`func (o *SSLDetails) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SSLDetails) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SSLDetails) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SSLDetails) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSubject

`func (o *SSLDetails) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SSLDetails) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SSLDetails) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SSLDetails) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetDefault

`func (o *SSLDetails) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *SSLDetails) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *SSLDetails) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *SSLDetails) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


