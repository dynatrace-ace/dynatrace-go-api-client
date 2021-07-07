# Extension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionName** | Pointer to **string** | Extension name | [optional] 
**Version** | Pointer to **string** | Extension version | [optional] 
**Author** | Pointer to [**AuthorDto**](AuthorDto.md) |  | [optional] 
**DataSources** | Pointer to **[]string** | Data sources that extension uses to gather data | [optional] 
**Variables** | Pointer to **[]string** | Custom variables used in extension configuration | [optional] 
**FeatureSets** | Pointer to **[]string** | Available feature sets | [optional] 
**MinDynatraceVersion** | Pointer to **string** | Minimal Dynatrace version that works with the extension | [optional] 
**FileHash** | Pointer to **string** | SHA-256 hash of uploaded Extension file | [optional] 

## Methods

### NewExtension

`func NewExtension() *Extension`

NewExtension instantiates a new Extension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionWithDefaults

`func NewExtensionWithDefaults() *Extension`

NewExtensionWithDefaults instantiates a new Extension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionName

`func (o *Extension) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *Extension) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *Extension) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.

### HasExtensionName

`func (o *Extension) HasExtensionName() bool`

HasExtensionName returns a boolean if a field has been set.

### GetVersion

`func (o *Extension) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Extension) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Extension) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Extension) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAuthor

`func (o *Extension) GetAuthor() AuthorDto`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Extension) GetAuthorOk() (*AuthorDto, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Extension) SetAuthor(v AuthorDto)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Extension) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetDataSources

`func (o *Extension) GetDataSources() []string`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *Extension) GetDataSourcesOk() (*[]string, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *Extension) SetDataSources(v []string)`

SetDataSources sets DataSources field to given value.

### HasDataSources

`func (o *Extension) HasDataSources() bool`

HasDataSources returns a boolean if a field has been set.

### GetVariables

`func (o *Extension) GetVariables() []string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Extension) GetVariablesOk() (*[]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Extension) SetVariables(v []string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Extension) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetFeatureSets

`func (o *Extension) GetFeatureSets() []string`

GetFeatureSets returns the FeatureSets field if non-nil, zero value otherwise.

### GetFeatureSetsOk

`func (o *Extension) GetFeatureSetsOk() (*[]string, bool)`

GetFeatureSetsOk returns a tuple with the FeatureSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSets

`func (o *Extension) SetFeatureSets(v []string)`

SetFeatureSets sets FeatureSets field to given value.

### HasFeatureSets

`func (o *Extension) HasFeatureSets() bool`

HasFeatureSets returns a boolean if a field has been set.

### GetMinDynatraceVersion

`func (o *Extension) GetMinDynatraceVersion() string`

GetMinDynatraceVersion returns the MinDynatraceVersion field if non-nil, zero value otherwise.

### GetMinDynatraceVersionOk

`func (o *Extension) GetMinDynatraceVersionOk() (*string, bool)`

GetMinDynatraceVersionOk returns a tuple with the MinDynatraceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDynatraceVersion

`func (o *Extension) SetMinDynatraceVersion(v string)`

SetMinDynatraceVersion sets MinDynatraceVersion field to given value.

### HasMinDynatraceVersion

`func (o *Extension) HasMinDynatraceVersion() bool`

HasMinDynatraceVersion returns a boolean if a field has been set.

### GetFileHash

`func (o *Extension) GetFileHash() string`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *Extension) GetFileHashOk() (*string, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHash

`func (o *Extension) SetFileHash(v string)`

SetFileHash sets FileHash field to given value.

### HasFileHash

`func (o *Extension) HasFileHash() bool`

HasFileHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


