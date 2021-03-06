/*
 * Dynatrace Cluster API
 *
 * Dynatrace Managed exposes cluster management functionality via REST endpoints. This interactive documentation also acts as a REST client you can use to interact with Dynatrace Managed clusters.   To authorize, use the API Token generated in [Settings - API Tokens page](/cmc#cm/apiToken). The HTTP status code of the response shows the result of your request. The expected response code for a successful request is documented individually per REST endpoint. Additionally the following error response codes can occur in our REST interface:  * 400 - Bad Request: Some request parameters are not correct. See response body for details. * 401 - Unauthorized: A valid authorization header is required but is missing. * 403 - Forbidden: Execution of request is not allowed, e.g. api-token is invalid. * 404 - Not Found: Endpoint does not exist or some entities could not be found, e.g. User account. * 500 - Internal Server Error: See response body for details. * 556 - Upgrade in progress: Operation couldn't be performed during the upgrade. 
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// SslCertDto SSL certificate configuration.
type SslCertDto struct {
	// Private key PKCS #8 standard, PEM base64-encoded format
	PrivateKeyEncoded string `json:"privateKeyEncoded"`
	// Certificate X.509 standard, PEM base64-encoded format, server certificate
	PublicKeyCertificateEncoded string `json:"publicKeyCertificateEncoded"`
	// Certificate(s) X.509 standard, PEM base64-encoded format, intermediate and root certificates
	CertificateChainEncoded *string `json:"certificateChainEncoded,omitempty"`
}

// NewSslCertDto instantiates a new SslCertDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSslCertDto(privateKeyEncoded string, publicKeyCertificateEncoded string) *SslCertDto {
	this := SslCertDto{}
	this.PrivateKeyEncoded = privateKeyEncoded
	this.PublicKeyCertificateEncoded = publicKeyCertificateEncoded
	return &this
}

// NewSslCertDtoWithDefaults instantiates a new SslCertDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSslCertDtoWithDefaults() *SslCertDto {
	this := SslCertDto{}
	return &this
}

// GetPrivateKeyEncoded returns the PrivateKeyEncoded field value
func (o *SslCertDto) GetPrivateKeyEncoded() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKeyEncoded
}

// GetPrivateKeyEncodedOk returns a tuple with the PrivateKeyEncoded field value
// and a boolean to check if the value has been set.
func (o *SslCertDto) GetPrivateKeyEncodedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PrivateKeyEncoded, true
}

// SetPrivateKeyEncoded sets field value
func (o *SslCertDto) SetPrivateKeyEncoded(v string) {
	o.PrivateKeyEncoded = v
}

// GetPublicKeyCertificateEncoded returns the PublicKeyCertificateEncoded field value
func (o *SslCertDto) GetPublicKeyCertificateEncoded() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKeyCertificateEncoded
}

// GetPublicKeyCertificateEncodedOk returns a tuple with the PublicKeyCertificateEncoded field value
// and a boolean to check if the value has been set.
func (o *SslCertDto) GetPublicKeyCertificateEncodedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PublicKeyCertificateEncoded, true
}

// SetPublicKeyCertificateEncoded sets field value
func (o *SslCertDto) SetPublicKeyCertificateEncoded(v string) {
	o.PublicKeyCertificateEncoded = v
}

// GetCertificateChainEncoded returns the CertificateChainEncoded field value if set, zero value otherwise.
func (o *SslCertDto) GetCertificateChainEncoded() string {
	if o == nil || o.CertificateChainEncoded == nil {
		var ret string
		return ret
	}
	return *o.CertificateChainEncoded
}

// GetCertificateChainEncodedOk returns a tuple with the CertificateChainEncoded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SslCertDto) GetCertificateChainEncodedOk() (*string, bool) {
	if o == nil || o.CertificateChainEncoded == nil {
		return nil, false
	}
	return o.CertificateChainEncoded, true
}

// HasCertificateChainEncoded returns a boolean if a field has been set.
func (o *SslCertDto) HasCertificateChainEncoded() bool {
	if o != nil && o.CertificateChainEncoded != nil {
		return true
	}

	return false
}

// SetCertificateChainEncoded gets a reference to the given string and assigns it to the CertificateChainEncoded field.
func (o *SslCertDto) SetCertificateChainEncoded(v string) {
	o.CertificateChainEncoded = &v
}

func (o SslCertDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["privateKeyEncoded"] = o.PrivateKeyEncoded
	}
	if true {
		toSerialize["publicKeyCertificateEncoded"] = o.PublicKeyCertificateEncoded
	}
	if o.CertificateChainEncoded != nil {
		toSerialize["certificateChainEncoded"] = o.CertificateChainEncoded
	}
	return json.Marshal(toSerialize)
}

type NullableSslCertDto struct {
	value *SslCertDto
	isSet bool
}

func (v NullableSslCertDto) Get() *SslCertDto {
	return v.value
}

func (v *NullableSslCertDto) Set(val *SslCertDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSslCertDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSslCertDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSslCertDto(val *SslCertDto) *NullableSslCertDto {
	return &NullableSslCertDto{value: val, isSet: true}
}

func (v NullableSslCertDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSslCertDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


