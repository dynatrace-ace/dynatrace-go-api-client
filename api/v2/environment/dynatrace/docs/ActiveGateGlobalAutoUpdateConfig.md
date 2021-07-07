# ActiveGateGlobalAutoUpdateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalSetting** | **string** | The state of auto-updates for all ActiveGates connected to the environment or Managed cluster.   This setting is inherited by all ActiveGates that have the &#x60;INHERITED&#x60; setting. | 
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 

## Methods

### NewActiveGateGlobalAutoUpdateConfig

`func NewActiveGateGlobalAutoUpdateConfig(globalSetting string, ) *ActiveGateGlobalAutoUpdateConfig`

NewActiveGateGlobalAutoUpdateConfig instantiates a new ActiveGateGlobalAutoUpdateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveGateGlobalAutoUpdateConfigWithDefaults

`func NewActiveGateGlobalAutoUpdateConfigWithDefaults() *ActiveGateGlobalAutoUpdateConfig`

NewActiveGateGlobalAutoUpdateConfigWithDefaults instantiates a new ActiveGateGlobalAutoUpdateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalSetting

`func (o *ActiveGateGlobalAutoUpdateConfig) GetGlobalSetting() string`

GetGlobalSetting returns the GlobalSetting field if non-nil, zero value otherwise.

### GetGlobalSettingOk

`func (o *ActiveGateGlobalAutoUpdateConfig) GetGlobalSettingOk() (*string, bool)`

GetGlobalSettingOk returns a tuple with the GlobalSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSetting

`func (o *ActiveGateGlobalAutoUpdateConfig) SetGlobalSetting(v string)`

SetGlobalSetting sets GlobalSetting field to given value.


### GetMetadata

`func (o *ActiveGateGlobalAutoUpdateConfig) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActiveGateGlobalAutoUpdateConfig) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActiveGateGlobalAutoUpdateConfig) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ActiveGateGlobalAutoUpdateConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


