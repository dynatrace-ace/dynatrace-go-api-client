# TechMonitoringList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Technologies** | [**[]Technology**](Technology.md) | A list of technology monitoring configurations. | 

## Methods

### NewTechMonitoringList

`func NewTechMonitoringList(technologies []Technology, ) *TechMonitoringList`

NewTechMonitoringList instantiates a new TechMonitoringList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechMonitoringListWithDefaults

`func NewTechMonitoringListWithDefaults() *TechMonitoringList`

NewTechMonitoringListWithDefaults instantiates a new TechMonitoringList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *TechMonitoringList) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TechMonitoringList) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TechMonitoringList) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TechMonitoringList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTechnologies

`func (o *TechMonitoringList) GetTechnologies() []Technology`

GetTechnologies returns the Technologies field if non-nil, zero value otherwise.

### GetTechnologiesOk

`func (o *TechMonitoringList) GetTechnologiesOk() (*[]Technology, bool)`

GetTechnologiesOk returns a tuple with the Technologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnologies

`func (o *TechMonitoringList) SetTechnologies(v []Technology)`

SetTechnologies sets Technologies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


