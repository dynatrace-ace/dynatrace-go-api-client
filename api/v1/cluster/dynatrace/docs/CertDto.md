# CertDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKeyEncoded** | **string** | Private key PKCS #8 standard, PEM base64-encoded format | 
**PublicKeyCertificateEncoded** | **string** | Certificate X.509 standard, PEM base64-encoded format | 

## Methods

### NewCertDto

`func NewCertDto(privateKeyEncoded string, publicKeyCertificateEncoded string, ) *CertDto`

NewCertDto instantiates a new CertDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertDtoWithDefaults

`func NewCertDtoWithDefaults() *CertDto`

NewCertDtoWithDefaults instantiates a new CertDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKeyEncoded

`func (o *CertDto) GetPrivateKeyEncoded() string`

GetPrivateKeyEncoded returns the PrivateKeyEncoded field if non-nil, zero value otherwise.

### GetPrivateKeyEncodedOk

`func (o *CertDto) GetPrivateKeyEncodedOk() (*string, bool)`

GetPrivateKeyEncodedOk returns a tuple with the PrivateKeyEncoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyEncoded

`func (o *CertDto) SetPrivateKeyEncoded(v string)`

SetPrivateKeyEncoded sets PrivateKeyEncoded field to given value.


### GetPublicKeyCertificateEncoded

`func (o *CertDto) GetPublicKeyCertificateEncoded() string`

GetPublicKeyCertificateEncoded returns the PublicKeyCertificateEncoded field if non-nil, zero value otherwise.

### GetPublicKeyCertificateEncodedOk

`func (o *CertDto) GetPublicKeyCertificateEncodedOk() (*string, bool)`

GetPublicKeyCertificateEncodedOk returns a tuple with the PublicKeyCertificateEncoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyCertificateEncoded

`func (o *CertDto) SetPublicKeyCertificateEncoded(v string)`

SetPublicKeyCertificateEncoded sets PublicKeyCertificateEncoded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


