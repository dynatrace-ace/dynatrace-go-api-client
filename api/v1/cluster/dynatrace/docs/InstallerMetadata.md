# InstallerMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterBackupPath** | Pointer to **string** | Network attached storage path for the backup | [optional] 
**ElasticsearchBackupPath** | Pointer to **string** | Elasticsearch repository used for the backup | [optional] 
**DatacenterWithBackupEnabled** | Pointer to **string** | Datacenter for backup, used only in multidc setups | [optional] 

## Methods

### NewInstallerMetadata

`func NewInstallerMetadata() *InstallerMetadata`

NewInstallerMetadata instantiates a new InstallerMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallerMetadataWithDefaults

`func NewInstallerMetadataWithDefaults() *InstallerMetadata`

NewInstallerMetadataWithDefaults instantiates a new InstallerMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterBackupPath

`func (o *InstallerMetadata) GetClusterBackupPath() string`

GetClusterBackupPath returns the ClusterBackupPath field if non-nil, zero value otherwise.

### GetClusterBackupPathOk

`func (o *InstallerMetadata) GetClusterBackupPathOk() (*string, bool)`

GetClusterBackupPathOk returns a tuple with the ClusterBackupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterBackupPath

`func (o *InstallerMetadata) SetClusterBackupPath(v string)`

SetClusterBackupPath sets ClusterBackupPath field to given value.

### HasClusterBackupPath

`func (o *InstallerMetadata) HasClusterBackupPath() bool`

HasClusterBackupPath returns a boolean if a field has been set.

### GetElasticsearchBackupPath

`func (o *InstallerMetadata) GetElasticsearchBackupPath() string`

GetElasticsearchBackupPath returns the ElasticsearchBackupPath field if non-nil, zero value otherwise.

### GetElasticsearchBackupPathOk

`func (o *InstallerMetadata) GetElasticsearchBackupPathOk() (*string, bool)`

GetElasticsearchBackupPathOk returns a tuple with the ElasticsearchBackupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticsearchBackupPath

`func (o *InstallerMetadata) SetElasticsearchBackupPath(v string)`

SetElasticsearchBackupPath sets ElasticsearchBackupPath field to given value.

### HasElasticsearchBackupPath

`func (o *InstallerMetadata) HasElasticsearchBackupPath() bool`

HasElasticsearchBackupPath returns a boolean if a field has been set.

### GetDatacenterWithBackupEnabled

`func (o *InstallerMetadata) GetDatacenterWithBackupEnabled() string`

GetDatacenterWithBackupEnabled returns the DatacenterWithBackupEnabled field if non-nil, zero value otherwise.

### GetDatacenterWithBackupEnabledOk

`func (o *InstallerMetadata) GetDatacenterWithBackupEnabledOk() (*string, bool)`

GetDatacenterWithBackupEnabledOk returns a tuple with the DatacenterWithBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterWithBackupEnabled

`func (o *InstallerMetadata) SetDatacenterWithBackupEnabled(v string)`

SetDatacenterWithBackupEnabled sets DatacenterWithBackupEnabled field to given value.

### HasDatacenterWithBackupEnabled

`func (o *InstallerMetadata) HasDatacenterWithBackupEnabled() bool`

HasDatacenterWithBackupEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


