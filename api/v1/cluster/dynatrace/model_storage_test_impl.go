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

// StorageTestImpl struct for StorageTestImpl
type StorageTestImpl struct {
	StoragePath *string `json:"storagePath,omitempty"`
	StorageError *string `json:"storageError,omitempty"`
	InProgress *bool `json:"inProgress,omitempty"`
}

// NewStorageTestImpl instantiates a new StorageTestImpl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageTestImpl() *StorageTestImpl {
	this := StorageTestImpl{}
	return &this
}

// NewStorageTestImplWithDefaults instantiates a new StorageTestImpl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageTestImplWithDefaults() *StorageTestImpl {
	this := StorageTestImpl{}
	return &this
}

// GetStoragePath returns the StoragePath field value if set, zero value otherwise.
func (o *StorageTestImpl) GetStoragePath() string {
	if o == nil || o.StoragePath == nil {
		var ret string
		return ret
	}
	return *o.StoragePath
}

// GetStoragePathOk returns a tuple with the StoragePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageTestImpl) GetStoragePathOk() (*string, bool) {
	if o == nil || o.StoragePath == nil {
		return nil, false
	}
	return o.StoragePath, true
}

// HasStoragePath returns a boolean if a field has been set.
func (o *StorageTestImpl) HasStoragePath() bool {
	if o != nil && o.StoragePath != nil {
		return true
	}

	return false
}

// SetStoragePath gets a reference to the given string and assigns it to the StoragePath field.
func (o *StorageTestImpl) SetStoragePath(v string) {
	o.StoragePath = &v
}

// GetStorageError returns the StorageError field value if set, zero value otherwise.
func (o *StorageTestImpl) GetStorageError() string {
	if o == nil || o.StorageError == nil {
		var ret string
		return ret
	}
	return *o.StorageError
}

// GetStorageErrorOk returns a tuple with the StorageError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageTestImpl) GetStorageErrorOk() (*string, bool) {
	if o == nil || o.StorageError == nil {
		return nil, false
	}
	return o.StorageError, true
}

// HasStorageError returns a boolean if a field has been set.
func (o *StorageTestImpl) HasStorageError() bool {
	if o != nil && o.StorageError != nil {
		return true
	}

	return false
}

// SetStorageError gets a reference to the given string and assigns it to the StorageError field.
func (o *StorageTestImpl) SetStorageError(v string) {
	o.StorageError = &v
}

// GetInProgress returns the InProgress field value if set, zero value otherwise.
func (o *StorageTestImpl) GetInProgress() bool {
	if o == nil || o.InProgress == nil {
		var ret bool
		return ret
	}
	return *o.InProgress
}

// GetInProgressOk returns a tuple with the InProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageTestImpl) GetInProgressOk() (*bool, bool) {
	if o == nil || o.InProgress == nil {
		return nil, false
	}
	return o.InProgress, true
}

// HasInProgress returns a boolean if a field has been set.
func (o *StorageTestImpl) HasInProgress() bool {
	if o != nil && o.InProgress != nil {
		return true
	}

	return false
}

// SetInProgress gets a reference to the given bool and assigns it to the InProgress field.
func (o *StorageTestImpl) SetInProgress(v bool) {
	o.InProgress = &v
}

func (o StorageTestImpl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StoragePath != nil {
		toSerialize["storagePath"] = o.StoragePath
	}
	if o.StorageError != nil {
		toSerialize["storageError"] = o.StorageError
	}
	if o.InProgress != nil {
		toSerialize["inProgress"] = o.InProgress
	}
	return json.Marshal(toSerialize)
}

type NullableStorageTestImpl struct {
	value *StorageTestImpl
	isSet bool
}

func (v NullableStorageTestImpl) Get() *StorageTestImpl {
	return v.value
}

func (v *NullableStorageTestImpl) Set(val *StorageTestImpl) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageTestImpl) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageTestImpl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageTestImpl(val *StorageTestImpl) *NullableStorageTestImpl {
	return &NullableStorageTestImpl{value: val, isSet: true}
}

func (v NullableStorageTestImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageTestImpl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


