# HostAutoUpdateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The Dynatrace entity ID of the host where OneAgent is deployed. | [optional] [readonly] 
**Setting** | **string** | The auto-update state of OneAgents on the host:   * &#x60;ENABLED&#x60;: OneAgent automatically updates to the most recent version.  * &#x60;DISABLED&#x60;: OneAgent updates to the version specified in the **version** field. * &#x60;INHERITED&#x60;: The setting from the host group (if the host is a member of a host group) or the environment-wide configuration (if the host doesn&#39;t belong to a host group) is used. | 
**Version** | Pointer to **string** | The version to which the OneAgent must be updated.   Specify the version in the &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;.&lt;timestamp&gt;&#x60; format (for example &#x60;1.191.0.20200326-161115&#x60;). You can fetch the list of available versions with the [GET available versions](https://www.dynatrace.com/support/help/shortlink/api-deployment-get-versions) call.   If no suitable installer is found for the provided version or the value is set to &#x60;null&#x60;, OneAgent won&#39;t be updated.   Only applicable when the **effectiveSetting** value is &#x60;DISABLED&#x60;.   If the **setting** parameter is set to &#x60;INHERITED&#x60; but the **version** is still set, it will result in a one-time update: OneAgent will be updated to the specified version and the **version** value will be set to &#x60;null&#x60;. For further updates the parent setting will be used. | [optional] 
**EffectiveSetting** | Pointer to **string** | The actual state of the auto-update on the host.   Applicable only if the **setting** parameter is set to &#x60;INHERITED&#x60;. In that case the value is taken from the host group or the environment-wide configuration. | [optional] [readonly] 
**EffectiveVersion** | Pointer to **string** | The actual version to which the OneAgent must be updated.   Applicable only if the **setting** parameter is set to &#x60;INHERITED&#x60; and the **version** parameter is set to &#x60;null&#x60;. In that case the value is taken from the host group or the environment-wide configuration. | [optional] [readonly] 

## Methods

### NewHostAutoUpdateConfig

`func NewHostAutoUpdateConfig(setting string, ) *HostAutoUpdateConfig`

NewHostAutoUpdateConfig instantiates a new HostAutoUpdateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostAutoUpdateConfigWithDefaults

`func NewHostAutoUpdateConfigWithDefaults() *HostAutoUpdateConfig`

NewHostAutoUpdateConfigWithDefaults instantiates a new HostAutoUpdateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *HostAutoUpdateConfig) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HostAutoUpdateConfig) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HostAutoUpdateConfig) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HostAutoUpdateConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *HostAutoUpdateConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostAutoUpdateConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostAutoUpdateConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HostAutoUpdateConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSetting

`func (o *HostAutoUpdateConfig) GetSetting() string`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *HostAutoUpdateConfig) GetSettingOk() (*string, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *HostAutoUpdateConfig) SetSetting(v string)`

SetSetting sets Setting field to given value.


### GetVersion

`func (o *HostAutoUpdateConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HostAutoUpdateConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HostAutoUpdateConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HostAutoUpdateConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEffectiveSetting

`func (o *HostAutoUpdateConfig) GetEffectiveSetting() string`

GetEffectiveSetting returns the EffectiveSetting field if non-nil, zero value otherwise.

### GetEffectiveSettingOk

`func (o *HostAutoUpdateConfig) GetEffectiveSettingOk() (*string, bool)`

GetEffectiveSettingOk returns a tuple with the EffectiveSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveSetting

`func (o *HostAutoUpdateConfig) SetEffectiveSetting(v string)`

SetEffectiveSetting sets EffectiveSetting field to given value.

### HasEffectiveSetting

`func (o *HostAutoUpdateConfig) HasEffectiveSetting() bool`

HasEffectiveSetting returns a boolean if a field has been set.

### GetEffectiveVersion

`func (o *HostAutoUpdateConfig) GetEffectiveVersion() string`

GetEffectiveVersion returns the EffectiveVersion field if non-nil, zero value otherwise.

### GetEffectiveVersionOk

`func (o *HostAutoUpdateConfig) GetEffectiveVersionOk() (*string, bool)`

GetEffectiveVersionOk returns a tuple with the EffectiveVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveVersion

`func (o *HostAutoUpdateConfig) SetEffectiveVersion(v string)`

SetEffectiveVersion sets EffectiveVersion field to given value.

### HasEffectiveVersion

`func (o *HostAutoUpdateConfig) HasEffectiveVersion() bool`

HasEffectiveVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


