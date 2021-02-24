# MonitoringConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The Dynatrace entity ID of the host where OneAgent is deployed. | [optional] [readonly] 
**MonitoringEnabled** | **bool** | The monitoring is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**MonitoringMode** | **string** | The monitoring mode for the host: full stack or infrastructure only. | 

## Methods

### NewMonitoringConfig

`func NewMonitoringConfig(monitoringEnabled bool, monitoringMode string, ) *MonitoringConfig`

NewMonitoringConfig instantiates a new MonitoringConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringConfigWithDefaults

`func NewMonitoringConfigWithDefaults() *MonitoringConfig`

NewMonitoringConfigWithDefaults instantiates a new MonitoringConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *MonitoringConfig) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MonitoringConfig) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MonitoringConfig) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MonitoringConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *MonitoringConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitoringConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitoringConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MonitoringConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMonitoringEnabled

`func (o *MonitoringConfig) GetMonitoringEnabled() bool`

GetMonitoringEnabled returns the MonitoringEnabled field if non-nil, zero value otherwise.

### GetMonitoringEnabledOk

`func (o *MonitoringConfig) GetMonitoringEnabledOk() (*bool, bool)`

GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEnabled

`func (o *MonitoringConfig) SetMonitoringEnabled(v bool)`

SetMonitoringEnabled sets MonitoringEnabled field to given value.


### GetMonitoringMode

`func (o *MonitoringConfig) GetMonitoringMode() string`

GetMonitoringMode returns the MonitoringMode field if non-nil, zero value otherwise.

### GetMonitoringModeOk

`func (o *MonitoringConfig) GetMonitoringModeOk() (*string, bool)`

GetMonitoringModeOk returns a tuple with the MonitoringMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringMode

`func (o *MonitoringConfig) SetMonitoringMode(v string)`

SetMonitoringMode sets MonitoringMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


