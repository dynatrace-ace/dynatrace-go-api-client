# ApplicationDetectionRulesHostDetectionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**HostDetectionHeaders** | **[]string** | An ordered list of host detection headers.   Headers are evaluated from top to bottom; the first matching header applies. | 

## Methods

### NewApplicationDetectionRulesHostDetectionSettings

`func NewApplicationDetectionRulesHostDetectionSettings(hostDetectionHeaders []string, ) *ApplicationDetectionRulesHostDetectionSettings`

NewApplicationDetectionRulesHostDetectionSettings instantiates a new ApplicationDetectionRulesHostDetectionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDetectionRulesHostDetectionSettingsWithDefaults

`func NewApplicationDetectionRulesHostDetectionSettingsWithDefaults() *ApplicationDetectionRulesHostDetectionSettings`

NewApplicationDetectionRulesHostDetectionSettingsWithDefaults instantiates a new ApplicationDetectionRulesHostDetectionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ApplicationDetectionRulesHostDetectionSettings) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationDetectionRulesHostDetectionSettings) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationDetectionRulesHostDetectionSettings) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationDetectionRulesHostDetectionSettings) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetHostDetectionHeaders

`func (o *ApplicationDetectionRulesHostDetectionSettings) GetHostDetectionHeaders() []string`

GetHostDetectionHeaders returns the HostDetectionHeaders field if non-nil, zero value otherwise.

### GetHostDetectionHeadersOk

`func (o *ApplicationDetectionRulesHostDetectionSettings) GetHostDetectionHeadersOk() (*[]string, bool)`

GetHostDetectionHeadersOk returns a tuple with the HostDetectionHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDetectionHeaders

`func (o *ApplicationDetectionRulesHostDetectionSettings) SetHostDetectionHeaders(v []string)`

SetHostDetectionHeaders sets HostDetectionHeaders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


