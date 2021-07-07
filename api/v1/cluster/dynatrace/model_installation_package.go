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

// InstallationPackage struct for InstallationPackage
type InstallationPackage struct {
	// If true, update package is excluded from download to save storage
	ExcludedFromDownload *bool `json:"excludedFromDownload,omitempty"`
	// If true, update package is removed from Dynatrace repository. Once you remove it from your cluster it will be gone permanently.
	DeletedFromMCStorage *bool `json:"deletedFromMCStorage,omitempty"`
	// IDs of nodes that have finished package processing (i.e, downloaded or removed depending on status)
	ReadyNodeIds *[]int32 `json:"readyNodeIds,omitempty"`
	// Total file size of a package in bytes
	FileSizeInBytes *int64 `json:"fileSizeInBytes,omitempty"`
	// If true, it's possible to remove this package via the REST API or the WebUi
	DeleteEnabled *bool `json:"deleteEnabled,omitempty"`
	// Count of tenants where this agent is configured as standard agent version. Applies to OneAgent type only.
	CountOfTenantsUsingThisAgentAsStandardVersion *int32 `json:"countOfTenantsUsingThisAgentAsStandardVersion,omitempty"`
	// Version
	Version *string `json:"version,omitempty"`
	// Cluster-wide status
	Status *string `json:"status,omitempty"`
	// Type
	Type *string `json:"type,omitempty"`
}

// NewInstallationPackage instantiates a new InstallationPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallationPackage() *InstallationPackage {
	this := InstallationPackage{}
	return &this
}

// NewInstallationPackageWithDefaults instantiates a new InstallationPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallationPackageWithDefaults() *InstallationPackage {
	this := InstallationPackage{}
	return &this
}

// GetExcludedFromDownload returns the ExcludedFromDownload field value if set, zero value otherwise.
func (o *InstallationPackage) GetExcludedFromDownload() bool {
	if o == nil || o.ExcludedFromDownload == nil {
		var ret bool
		return ret
	}
	return *o.ExcludedFromDownload
}

// GetExcludedFromDownloadOk returns a tuple with the ExcludedFromDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationPackage) GetExcludedFromDownloadOk() (*bool, bool) {
	if o == nil || o.ExcludedFromDownload == nil {
		return nil, false
	}
	return o.ExcludedFromDownload, true
}

// HasExcludedFromDownload returns a boolean if a field has been set.
func (o *InstallationPackage) HasExcludedFromDownload() bool {
	if o != nil && o.ExcludedFromDownload != nil {
		return true
	}

	return false
}

// SetExcludedFromDownload gets a reference to the given bool and assigns it to the ExcludedFromDownload field.
func (o *InstallationPackage) SetExcludedFromDownload(v bool) {
	o.ExcludedFromDownload = &v
}

// GetDeletedFromMCStorage returns the DeletedFromMCStorage field value if set, zero value otherwise.
func (o *InstallationPackage) GetDeletedFromMCStorage() bool {
	if o == nil || o.DeletedFromMCStorage == nil {
		var ret bool
		return ret
	}
	return *o.DeletedFromMCStorage
}

// GetDeletedFromMCStorageOk returns a tuple with the DeletedFromMCStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationPackage) GetDeletedFromMCStorageOk() (*bool, bool) {
	if o == nil || o.DeletedFromMCStorage == nil {
		return nil, false
	}
	return o.DeletedFromMCStorage, true
}

// HasDeletedFromMCStorage returns a boolean if a field has been set.
func (o *InstallationPackage) HasDeletedFromMCStorage() bool {
	if o != nil && o.DeletedFromMCStorage != nil {
		return true
	}

	return false
}

// SetDeletedFromMCStorage gets a reference to the given bool and assigns it to the DeletedFromMCStorage field.
func (o *InstallationPackage) SetDeletedFromMCStorage(v bool) {
	o.DeletedFromMCStorage = &v
}

// GetReadyNodeIds returns the ReadyNodeIds field value if set, zero value otherwise.
func (o *InstallationPackage) GetReadyNodeIds() []int32 {
	if o == nil || o.ReadyNodeIds == nil {
		var ret []int32
		return ret
	}
	return *o.ReadyNodeIds
}

// GetReadyNodeIdsOk returns a tuple with the ReadyNodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationPackage) GetReadyNodeIdsOk() (*[]int32, bool) {
	if o == nil || o.ReadyNodeIds == nil {
		return nil, false
	}
	return o.ReadyNodeIds, true
}

// HasReadyNodeIds returns a boolean if a field has been set.
func (o *InstallationPackage) HasReadyNodeIds() bool {
	if o != nil && o.ReadyNodeIds != nil {
		return true
	}

	return false
}

// SetReadyNodeIds gets a reference to the given []int32 and assigns it to the ReadyNodeIds field.
func (o *InstallationPackage) SetReadyNodeIds(v []int32) {
	o.ReadyNodeIds = &v
}

// GetFileSizeInBytes returns the FileSizeInBytes field value if set, zero value otherwise.
func (o *InstallationPackage) GetFileSizeInBytes() int64 {
	if o == nil || o.FileSizeInBytes == nil {
		var ret int64
		return ret
	}
	return *o.FileSizeInBytes
}

// GetFileSizeInBytesOk returns a tuple with the FileSizeInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationPackage) GetFileSizeInBytesOk() (*int64, bool) {
	if o == nil || o.FileSizeInBytes == nil {
		return nil, false
	}
	return o.FileSizeInBytes, true
}

// HasFileSizeInBytes returns a boolean if a field has been set.
func (o *InstallationPackage) HasFileSizeInBytes() bool {
	if o != nil && o.FileSizeInBytes != nil {
		return true
	}

	return false
}

// SetFileSizeInBytes gets a reference to the given int64 and assigns it to the FileSizeInBytes field.
func (o *InstallationPackage) SetFileSizeInBytes(v int64) {
	o.FileSizeInBytes = &v
}

// GetDeleteEnabled returns the DeleteEnabled field value if set, zero value otherwise.
func (o *InstallationPackage) GetDeleteEnabled() bool {
	if o == nil || o.DeleteEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DeleteEnabled
}

// GetDeleteEnabledOk returns a tuple with the DeleteEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationPackage) GetDeleteEnabledOk() (*bool, bool) {
	if o == nil || o.DeleteEnabled == nil {
		return nil, false
	}
	return o.DeleteEnabled, true
}

// HasDeleteEnabled returns a boolean if a field has been set.
func (o *InstallationPackage) HasDeleteEnabled() bool {
	if o != nil && o.DeleteEnabled != nil {
		return true
	}

	return false
}

// SetDeleteEnabled gets a reference to the given bool and assigns it to the DeleteEnabled field.
func (o *InstallationPackage) SetDeleteEnabled(v bool) {
	o.DeleteEnabled = &v
}

// GetCountOfTenantsUsingThisAgentAsStandardVersion returns the CountOfTenantsUsingThisAgentAsStandardVersion field value if set, zero value otherwise.
func (o *InstallationPackage) GetCountOfTenantsUsingThisAgentAsStandardVersion() int32 {
	if o == nil || o.CountOfTenantsUsingThisAgentAsStandardVersion == nil {
		var ret int32
		return ret
	}
	return *o.CountOfTenantsUsingThisAgentAsStandardVersion
}

// GetCountOfTenantsUsingThisAgentAsStandardVersionOk returns a tuple with the CountOfTenantsUsingThisAgentAsStandardVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationPackage) GetCountOfTenantsUsingThisAgentAsStandardVersionOk() (*int32, bool) {
	if o == nil || o.CountOfTenantsUsingThisAgentAsStandardVersion == nil {
		return nil, false
	}
	return o.CountOfTenantsUsingThisAgentAsStandardVersion, true
}

// HasCountOfTenantsUsingThisAgentAsStandardVersion returns a boolean if a field has been set.
func (o *InstallationPackage) HasCountOfTenantsUsingThisAgentAsStandardVersion() bool {
	if o != nil && o.CountOfTenantsUsingThisAgentAsStandardVersion != nil {
		return true
	}

	return false
}

// SetCountOfTenantsUsingThisAgentAsStandardVersion gets a reference to the given int32 and assigns it to the CountOfTenantsUsingThisAgentAsStandardVersion field.
func (o *InstallationPackage) SetCountOfTenantsUsingThisAgentAsStandardVersion(v int32) {
	o.CountOfTenantsUsingThisAgentAsStandardVersion = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InstallationPackage) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationPackage) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InstallationPackage) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *InstallationPackage) SetVersion(v string) {
	o.Version = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InstallationPackage) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationPackage) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InstallationPackage) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InstallationPackage) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InstallationPackage) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationPackage) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InstallationPackage) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InstallationPackage) SetType(v string) {
	o.Type = &v
}

func (o InstallationPackage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExcludedFromDownload != nil {
		toSerialize["excludedFromDownload"] = o.ExcludedFromDownload
	}
	if o.DeletedFromMCStorage != nil {
		toSerialize["deletedFromMCStorage"] = o.DeletedFromMCStorage
	}
	if o.ReadyNodeIds != nil {
		toSerialize["readyNodeIds"] = o.ReadyNodeIds
	}
	if o.FileSizeInBytes != nil {
		toSerialize["fileSizeInBytes"] = o.FileSizeInBytes
	}
	if o.DeleteEnabled != nil {
		toSerialize["deleteEnabled"] = o.DeleteEnabled
	}
	if o.CountOfTenantsUsingThisAgentAsStandardVersion != nil {
		toSerialize["countOfTenantsUsingThisAgentAsStandardVersion"] = o.CountOfTenantsUsingThisAgentAsStandardVersion
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInstallationPackage struct {
	value *InstallationPackage
	isSet bool
}

func (v NullableInstallationPackage) Get() *InstallationPackage {
	return v.value
}

func (v *NullableInstallationPackage) Set(val *InstallationPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallationPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallationPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallationPackage(val *InstallationPackage) *NullableInstallationPackage {
	return &NullableInstallationPackage{value: val, isSet: true}
}

func (v NullableInstallationPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallationPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

