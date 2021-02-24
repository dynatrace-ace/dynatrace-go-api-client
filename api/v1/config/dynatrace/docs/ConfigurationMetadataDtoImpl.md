# ConfigurationMetadataDtoImpl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentConfigurationVersions** | Pointer to **[]string** | A Sorted list of string version numbers of the configuration. | [optional] 
**ConfigurationVersions** | Pointer to **[]int64** | A Sorted list of the version numbers of the configuration. | [optional] 
**ClusterVersion** | Pointer to **string** | Dynatrace server version. | [optional] 

## Methods

### NewConfigurationMetadataDtoImpl

`func NewConfigurationMetadataDtoImpl() *ConfigurationMetadataDtoImpl`

NewConfigurationMetadataDtoImpl instantiates a new ConfigurationMetadataDtoImpl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationMetadataDtoImplWithDefaults

`func NewConfigurationMetadataDtoImplWithDefaults() *ConfigurationMetadataDtoImpl`

NewConfigurationMetadataDtoImplWithDefaults instantiates a new ConfigurationMetadataDtoImpl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentConfigurationVersions

`func (o *ConfigurationMetadataDtoImpl) GetCurrentConfigurationVersions() []string`

GetCurrentConfigurationVersions returns the CurrentConfigurationVersions field if non-nil, zero value otherwise.

### GetCurrentConfigurationVersionsOk

`func (o *ConfigurationMetadataDtoImpl) GetCurrentConfigurationVersionsOk() (*[]string, bool)`

GetCurrentConfigurationVersionsOk returns a tuple with the CurrentConfigurationVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConfigurationVersions

`func (o *ConfigurationMetadataDtoImpl) SetCurrentConfigurationVersions(v []string)`

SetCurrentConfigurationVersions sets CurrentConfigurationVersions field to given value.

### HasCurrentConfigurationVersions

`func (o *ConfigurationMetadataDtoImpl) HasCurrentConfigurationVersions() bool`

HasCurrentConfigurationVersions returns a boolean if a field has been set.

### GetConfigurationVersions

`func (o *ConfigurationMetadataDtoImpl) GetConfigurationVersions() []int64`

GetConfigurationVersions returns the ConfigurationVersions field if non-nil, zero value otherwise.

### GetConfigurationVersionsOk

`func (o *ConfigurationMetadataDtoImpl) GetConfigurationVersionsOk() (*[]int64, bool)`

GetConfigurationVersionsOk returns a tuple with the ConfigurationVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationVersions

`func (o *ConfigurationMetadataDtoImpl) SetConfigurationVersions(v []int64)`

SetConfigurationVersions sets ConfigurationVersions field to given value.

### HasConfigurationVersions

`func (o *ConfigurationMetadataDtoImpl) HasConfigurationVersions() bool`

HasConfigurationVersions returns a boolean if a field has been set.

### GetClusterVersion

`func (o *ConfigurationMetadataDtoImpl) GetClusterVersion() string`

GetClusterVersion returns the ClusterVersion field if non-nil, zero value otherwise.

### GetClusterVersionOk

`func (o *ConfigurationMetadataDtoImpl) GetClusterVersionOk() (*string, bool)`

GetClusterVersionOk returns a tuple with the ClusterVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterVersion

`func (o *ConfigurationMetadataDtoImpl) SetClusterVersion(v string)`

SetClusterVersion sets ClusterVersion field to given value.

### HasClusterVersion

`func (o *ConfigurationMetadataDtoImpl) HasClusterVersion() bool`

HasClusterVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


