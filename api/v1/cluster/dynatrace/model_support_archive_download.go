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

// SupportArchiveDownload struct for SupportArchiveDownload
type SupportArchiveDownload struct {
	DownloadState *string `json:"downloadState,omitempty"`
	Report *SupportArchiveReport `json:"report,omitempty"`
}

// NewSupportArchiveDownload instantiates a new SupportArchiveDownload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportArchiveDownload() *SupportArchiveDownload {
	this := SupportArchiveDownload{}
	return &this
}

// NewSupportArchiveDownloadWithDefaults instantiates a new SupportArchiveDownload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportArchiveDownloadWithDefaults() *SupportArchiveDownload {
	this := SupportArchiveDownload{}
	return &this
}

// GetDownloadState returns the DownloadState field value if set, zero value otherwise.
func (o *SupportArchiveDownload) GetDownloadState() string {
	if o == nil || o.DownloadState == nil {
		var ret string
		return ret
	}
	return *o.DownloadState
}

// GetDownloadStateOk returns a tuple with the DownloadState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportArchiveDownload) GetDownloadStateOk() (*string, bool) {
	if o == nil || o.DownloadState == nil {
		return nil, false
	}
	return o.DownloadState, true
}

// HasDownloadState returns a boolean if a field has been set.
func (o *SupportArchiveDownload) HasDownloadState() bool {
	if o != nil && o.DownloadState != nil {
		return true
	}

	return false
}

// SetDownloadState gets a reference to the given string and assigns it to the DownloadState field.
func (o *SupportArchiveDownload) SetDownloadState(v string) {
	o.DownloadState = &v
}

// GetReport returns the Report field value if set, zero value otherwise.
func (o *SupportArchiveDownload) GetReport() SupportArchiveReport {
	if o == nil || o.Report == nil {
		var ret SupportArchiveReport
		return ret
	}
	return *o.Report
}

// GetReportOk returns a tuple with the Report field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportArchiveDownload) GetReportOk() (*SupportArchiveReport, bool) {
	if o == nil || o.Report == nil {
		return nil, false
	}
	return o.Report, true
}

// HasReport returns a boolean if a field has been set.
func (o *SupportArchiveDownload) HasReport() bool {
	if o != nil && o.Report != nil {
		return true
	}

	return false
}

// SetReport gets a reference to the given SupportArchiveReport and assigns it to the Report field.
func (o *SupportArchiveDownload) SetReport(v SupportArchiveReport) {
	o.Report = &v
}

func (o SupportArchiveDownload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DownloadState != nil {
		toSerialize["downloadState"] = o.DownloadState
	}
	if o.Report != nil {
		toSerialize["report"] = o.Report
	}
	return json.Marshal(toSerialize)
}

type NullableSupportArchiveDownload struct {
	value *SupportArchiveDownload
	isSet bool
}

func (v NullableSupportArchiveDownload) Get() *SupportArchiveDownload {
	return v.value
}

func (v *NullableSupportArchiveDownload) Set(val *SupportArchiveDownload) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportArchiveDownload) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportArchiveDownload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportArchiveDownload(val *SupportArchiveDownload) *NullableSupportArchiveDownload {
	return &NullableSupportArchiveDownload{value: val, isSet: true}
}

func (v NullableSupportArchiveDownload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportArchiveDownload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


