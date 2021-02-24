# FrequentIssueDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**FrequentIssueDetectionApplicationEnabled** | **bool** | The detection for applications is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**FrequentIssueDetectionServiceEnabled** | **bool** | The detection for services is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**FrequentIssueDetectionInfrastructureEnabled** | **bool** | The detection for infrastructure is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 

## Methods

### NewFrequentIssueDetectionConfig

`func NewFrequentIssueDetectionConfig(frequentIssueDetectionApplicationEnabled bool, frequentIssueDetectionServiceEnabled bool, frequentIssueDetectionInfrastructureEnabled bool, ) *FrequentIssueDetectionConfig`

NewFrequentIssueDetectionConfig instantiates a new FrequentIssueDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrequentIssueDetectionConfigWithDefaults

`func NewFrequentIssueDetectionConfigWithDefaults() *FrequentIssueDetectionConfig`

NewFrequentIssueDetectionConfigWithDefaults instantiates a new FrequentIssueDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *FrequentIssueDetectionConfig) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FrequentIssueDetectionConfig) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FrequentIssueDetectionConfig) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FrequentIssueDetectionConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFrequentIssueDetectionApplicationEnabled

`func (o *FrequentIssueDetectionConfig) GetFrequentIssueDetectionApplicationEnabled() bool`

GetFrequentIssueDetectionApplicationEnabled returns the FrequentIssueDetectionApplicationEnabled field if non-nil, zero value otherwise.

### GetFrequentIssueDetectionApplicationEnabledOk

`func (o *FrequentIssueDetectionConfig) GetFrequentIssueDetectionApplicationEnabledOk() (*bool, bool)`

GetFrequentIssueDetectionApplicationEnabledOk returns a tuple with the FrequentIssueDetectionApplicationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequentIssueDetectionApplicationEnabled

`func (o *FrequentIssueDetectionConfig) SetFrequentIssueDetectionApplicationEnabled(v bool)`

SetFrequentIssueDetectionApplicationEnabled sets FrequentIssueDetectionApplicationEnabled field to given value.


### GetFrequentIssueDetectionServiceEnabled

`func (o *FrequentIssueDetectionConfig) GetFrequentIssueDetectionServiceEnabled() bool`

GetFrequentIssueDetectionServiceEnabled returns the FrequentIssueDetectionServiceEnabled field if non-nil, zero value otherwise.

### GetFrequentIssueDetectionServiceEnabledOk

`func (o *FrequentIssueDetectionConfig) GetFrequentIssueDetectionServiceEnabledOk() (*bool, bool)`

GetFrequentIssueDetectionServiceEnabledOk returns a tuple with the FrequentIssueDetectionServiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequentIssueDetectionServiceEnabled

`func (o *FrequentIssueDetectionConfig) SetFrequentIssueDetectionServiceEnabled(v bool)`

SetFrequentIssueDetectionServiceEnabled sets FrequentIssueDetectionServiceEnabled field to given value.


### GetFrequentIssueDetectionInfrastructureEnabled

`func (o *FrequentIssueDetectionConfig) GetFrequentIssueDetectionInfrastructureEnabled() bool`

GetFrequentIssueDetectionInfrastructureEnabled returns the FrequentIssueDetectionInfrastructureEnabled field if non-nil, zero value otherwise.

### GetFrequentIssueDetectionInfrastructureEnabledOk

`func (o *FrequentIssueDetectionConfig) GetFrequentIssueDetectionInfrastructureEnabledOk() (*bool, bool)`

GetFrequentIssueDetectionInfrastructureEnabledOk returns a tuple with the FrequentIssueDetectionInfrastructureEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequentIssueDetectionInfrastructureEnabled

`func (o *FrequentIssueDetectionConfig) SetFrequentIssueDetectionInfrastructureEnabled(v bool)`

SetFrequentIssueDetectionInfrastructureEnabled sets FrequentIssueDetectionInfrastructureEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


