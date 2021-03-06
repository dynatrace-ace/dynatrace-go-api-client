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

// CertificateStoreStatus struct for CertificateStoreStatus
type CertificateStoreStatus struct {
	DetailedError *string `json:"detailedError,omitempty"`
	CertificateStoreStatus *string `json:"certificateStoreStatus,omitempty"`
}

// NewCertificateStoreStatus instantiates a new CertificateStoreStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateStoreStatus() *CertificateStoreStatus {
	this := CertificateStoreStatus{}
	return &this
}

// NewCertificateStoreStatusWithDefaults instantiates a new CertificateStoreStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateStoreStatusWithDefaults() *CertificateStoreStatus {
	this := CertificateStoreStatus{}
	return &this
}

// GetDetailedError returns the DetailedError field value if set, zero value otherwise.
func (o *CertificateStoreStatus) GetDetailedError() string {
	if o == nil || o.DetailedError == nil {
		var ret string
		return ret
	}
	return *o.DetailedError
}

// GetDetailedErrorOk returns a tuple with the DetailedError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateStoreStatus) GetDetailedErrorOk() (*string, bool) {
	if o == nil || o.DetailedError == nil {
		return nil, false
	}
	return o.DetailedError, true
}

// HasDetailedError returns a boolean if a field has been set.
func (o *CertificateStoreStatus) HasDetailedError() bool {
	if o != nil && o.DetailedError != nil {
		return true
	}

	return false
}

// SetDetailedError gets a reference to the given string and assigns it to the DetailedError field.
func (o *CertificateStoreStatus) SetDetailedError(v string) {
	o.DetailedError = &v
}

// GetCertificateStoreStatus returns the CertificateStoreStatus field value if set, zero value otherwise.
func (o *CertificateStoreStatus) GetCertificateStoreStatus() string {
	if o == nil || o.CertificateStoreStatus == nil {
		var ret string
		return ret
	}
	return *o.CertificateStoreStatus
}

// GetCertificateStoreStatusOk returns a tuple with the CertificateStoreStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateStoreStatus) GetCertificateStoreStatusOk() (*string, bool) {
	if o == nil || o.CertificateStoreStatus == nil {
		return nil, false
	}
	return o.CertificateStoreStatus, true
}

// HasCertificateStoreStatus returns a boolean if a field has been set.
func (o *CertificateStoreStatus) HasCertificateStoreStatus() bool {
	if o != nil && o.CertificateStoreStatus != nil {
		return true
	}

	return false
}

// SetCertificateStoreStatus gets a reference to the given string and assigns it to the CertificateStoreStatus field.
func (o *CertificateStoreStatus) SetCertificateStoreStatus(v string) {
	o.CertificateStoreStatus = &v
}

func (o CertificateStoreStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DetailedError != nil {
		toSerialize["detailedError"] = o.DetailedError
	}
	if o.CertificateStoreStatus != nil {
		toSerialize["certificateStoreStatus"] = o.CertificateStoreStatus
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateStoreStatus struct {
	value *CertificateStoreStatus
	isSet bool
}

func (v NullableCertificateStoreStatus) Get() *CertificateStoreStatus {
	return v.value
}

func (v *NullableCertificateStoreStatus) Set(val *CertificateStoreStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateStoreStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateStoreStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateStoreStatus(val *CertificateStoreStatus) *NullableCertificateStoreStatus {
	return &NullableCertificateStoreStatus{value: val, isSet: true}
}

func (v NullableCertificateStoreStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateStoreStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


