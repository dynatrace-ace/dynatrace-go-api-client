# ConfigurationMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationVersions** | Pointer to **[]int64** | A sorted list of the version numbers of the configuration. | [optional] 
**CurrentConfigurationVersions** | Pointer to **[]string** | A sorted list of version numbers of the configuration. | [optional] 
**ClusterVersion** | Pointer to **string** | Dynatrace version. | [optional] 

## Methods

### NewConfigurationMetadata

`func NewConfigurationMetadata() *ConfigurationMetadata`

NewConfigurationMetadata instantiates a new ConfigurationMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationMetadataWithDefaults

`func NewConfigurationMetadataWithDefaults() *ConfigurationMetadata`

NewConfigurationMetadataWithDefaults instantiates a new ConfigurationMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationVersions

`func (o *ConfigurationMetadata) GetConfigurationVersions() []int64`

GetConfigurationVersions returns the ConfigurationVersions field if non-nil, zero value otherwise.

### GetConfigurationVersionsOk

`func (o *ConfigurationMetadata) GetConfigurationVersionsOk() (*[]int64, bool)`

GetConfigurationVersionsOk returns a tuple with the ConfigurationVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationVersions

`func (o *ConfigurationMetadata) SetConfigurationVersions(v []int64)`

SetConfigurationVersions sets ConfigurationVersions field to given value.

### HasConfigurationVersions

`func (o *ConfigurationMetadata) HasConfigurationVersions() bool`

HasConfigurationVersions returns a boolean if a field has been set.

### GetCurrentConfigurationVersions

`func (o *ConfigurationMetadata) GetCurrentConfigurationVersions() []string`

GetCurrentConfigurationVersions returns the CurrentConfigurationVersions field if non-nil, zero value otherwise.

### GetCurrentConfigurationVersionsOk

`func (o *ConfigurationMetadata) GetCurrentConfigurationVersionsOk() (*[]string, bool)`

GetCurrentConfigurationVersionsOk returns a tuple with the CurrentConfigurationVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConfigurationVersions

`func (o *ConfigurationMetadata) SetCurrentConfigurationVersions(v []string)`

SetCurrentConfigurationVersions sets CurrentConfigurationVersions field to given value.

### HasCurrentConfigurationVersions

`func (o *ConfigurationMetadata) HasCurrentConfigurationVersions() bool`

HasCurrentConfigurationVersions returns a boolean if a field has been set.

### GetClusterVersion

`func (o *ConfigurationMetadata) GetClusterVersion() string`

GetClusterVersion returns the ClusterVersion field if non-nil, zero value otherwise.

### GetClusterVersionOk

`func (o *ConfigurationMetadata) GetClusterVersionOk() (*string, bool)`

GetClusterVersionOk returns a tuple with the ClusterVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterVersion

`func (o *ConfigurationMetadata) SetClusterVersion(v string)`

SetClusterVersion sets ClusterVersion field to given value.

### HasClusterVersion

`func (o *ConfigurationMetadata) HasClusterVersion() bool`

HasClusterVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


