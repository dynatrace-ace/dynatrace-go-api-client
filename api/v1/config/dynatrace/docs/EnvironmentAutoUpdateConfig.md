# EnvironmentAutoUpdateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Setting** | **string** | The auto-update state of OneAgents connecting to the environment:   * &#x60;ENABLED&#x60;: OneAgents automatically update to the most recent version.  * &#x60;DISABLED&#x60;: OneAgents update to the version specified in the **version** field.  OneAgents that connect to the environment use this configuration only when their **setting** parameter is set to &#x60;INHERITED&#x60;. | 
**Version** | Pointer to **string** | The version to which the OneAgent must be updated.   Specify the version in the &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;&#x60; format (for example &#x60;1.181.0&#x60;). You can fetch the list of available versions with the [GET available versions](https://www.dynatrace.com/support/help/shortlink/api-deployment-get-versions) call. If no suitable installer is found for the provided version or the value is set to &#x60;null&#x60;, OneAgent won&#39;t be updated.   Only applicable when the **setting** parameter is set to &#x60;DISABLED&#x60;. | [optional] 

## Methods

### NewEnvironmentAutoUpdateConfig

`func NewEnvironmentAutoUpdateConfig(setting string, ) *EnvironmentAutoUpdateConfig`

NewEnvironmentAutoUpdateConfig instantiates a new EnvironmentAutoUpdateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentAutoUpdateConfigWithDefaults

`func NewEnvironmentAutoUpdateConfigWithDefaults() *EnvironmentAutoUpdateConfig`

NewEnvironmentAutoUpdateConfigWithDefaults instantiates a new EnvironmentAutoUpdateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *EnvironmentAutoUpdateConfig) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnvironmentAutoUpdateConfig) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnvironmentAutoUpdateConfig) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EnvironmentAutoUpdateConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSetting

`func (o *EnvironmentAutoUpdateConfig) GetSetting() string`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *EnvironmentAutoUpdateConfig) GetSettingOk() (*string, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *EnvironmentAutoUpdateConfig) SetSetting(v string)`

SetSetting sets Setting field to given value.


### GetVersion

`func (o *EnvironmentAutoUpdateConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EnvironmentAutoUpdateConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EnvironmentAutoUpdateConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EnvironmentAutoUpdateConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


