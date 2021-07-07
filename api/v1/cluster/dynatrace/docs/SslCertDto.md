# SslCertDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKeyEncoded** | **string** | Private key PKCS #8 standard, PEM base64-encoded format | 
**PublicKeyCertificateEncoded** | **string** | Certificate X.509 standard, PEM base64-encoded format, server certificate | 
**CertificateChainEncoded** | Pointer to **string** | Certificate(s) X.509 standard, PEM base64-encoded format, intermediate and root certificates | [optional] 

## Methods

### NewSslCertDto

`func NewSslCertDto(privateKeyEncoded string, publicKeyCertificateEncoded string, ) *SslCertDto`

NewSslCertDto instantiates a new SslCertDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslCertDtoWithDefaults

`func NewSslCertDtoWithDefaults() *SslCertDto`

NewSslCertDtoWithDefaults instantiates a new SslCertDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKeyEncoded

`func (o *SslCertDto) GetPrivateKeyEncoded() string`

GetPrivateKeyEncoded returns the PrivateKeyEncoded field if non-nil, zero value otherwise.

### GetPrivateKeyEncodedOk

`func (o *SslCertDto) GetPrivateKeyEncodedOk() (*string, bool)`

GetPrivateKeyEncodedOk returns a tuple with the PrivateKeyEncoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyEncoded

`func (o *SslCertDto) SetPrivateKeyEncoded(v string)`

SetPrivateKeyEncoded sets PrivateKeyEncoded field to given value.


### GetPublicKeyCertificateEncoded

`func (o *SslCertDto) GetPublicKeyCertificateEncoded() string`

GetPublicKeyCertificateEncoded returns the PublicKeyCertificateEncoded field if non-nil, zero value otherwise.

### GetPublicKeyCertificateEncodedOk

`func (o *SslCertDto) GetPublicKeyCertificateEncodedOk() (*string, bool)`

GetPublicKeyCertificateEncodedOk returns a tuple with the PublicKeyCertificateEncoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyCertificateEncoded

`func (o *SslCertDto) SetPublicKeyCertificateEncoded(v string)`

SetPublicKeyCertificateEncoded sets PublicKeyCertificateEncoded field to given value.


### GetCertificateChainEncoded

`func (o *SslCertDto) GetCertificateChainEncoded() string`

GetCertificateChainEncoded returns the CertificateChainEncoded field if non-nil, zero value otherwise.

### GetCertificateChainEncodedOk

`func (o *SslCertDto) GetCertificateChainEncodedOk() (*string, bool)`

GetCertificateChainEncodedOk returns a tuple with the CertificateChainEncoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChainEncoded

`func (o *SslCertDto) SetCertificateChainEncoded(v string)`

SetCertificateChainEncoded sets CertificateChainEncoded field to given value.

### HasCertificateChainEncoded

`func (o *SslCertDto) HasCertificateChainEncoded() bool`

HasCertificateChainEncoded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


