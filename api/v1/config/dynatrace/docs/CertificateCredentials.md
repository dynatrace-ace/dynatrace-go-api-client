# CertificateCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | The certificate in the string format. | 
**Password** | **string** | The password of the credential. | 
**CertificateFormat** | **string** | The certificate format. | 

## Methods

### NewCertificateCredentials

`func NewCertificateCredentials(certificate string, password string, certificateFormat string, ) *CertificateCredentials`

NewCertificateCredentials instantiates a new CertificateCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateCredentialsWithDefaults

`func NewCertificateCredentialsWithDefaults() *CertificateCredentials`

NewCertificateCredentialsWithDefaults instantiates a new CertificateCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *CertificateCredentials) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateCredentials) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateCredentials) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetPassword

`func (o *CertificateCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CertificateCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CertificateCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetCertificateFormat

`func (o *CertificateCredentials) GetCertificateFormat() string`

GetCertificateFormat returns the CertificateFormat field if non-nil, zero value otherwise.

### GetCertificateFormatOk

`func (o *CertificateCredentials) GetCertificateFormatOk() (*string, bool)`

GetCertificateFormatOk returns a tuple with the CertificateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFormat

`func (o *CertificateCredentials) SetCertificateFormat(v string)`

SetCertificateFormat sets CertificateFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


