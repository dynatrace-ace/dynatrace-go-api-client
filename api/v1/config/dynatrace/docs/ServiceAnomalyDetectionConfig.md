# ServiceAnomalyDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**ResponseTimeDegradation** | [**ResponseTimeDegradationDetectionConfig**](ResponseTimeDegradationDetectionConfig.md) |  | 
**LoadDrop** | Pointer to [**LoadDropDetectionConfig**](LoadDropDetectionConfig.md) |  | [optional] 
**LoadSpike** | Pointer to [**LoadSpikeDetectionConfig**](LoadSpikeDetectionConfig.md) |  | [optional] 
**FailureRateIncrease** | [**FailureRateIncreaseDetectionConfig**](FailureRateIncreaseDetectionConfig.md) |  | 

## Methods

### NewServiceAnomalyDetectionConfig

`func NewServiceAnomalyDetectionConfig(responseTimeDegradation ResponseTimeDegradationDetectionConfig, failureRateIncrease FailureRateIncreaseDetectionConfig, ) *ServiceAnomalyDetectionConfig`

NewServiceAnomalyDetectionConfig instantiates a new ServiceAnomalyDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAnomalyDetectionConfigWithDefaults

`func NewServiceAnomalyDetectionConfigWithDefaults() *ServiceAnomalyDetectionConfig`

NewServiceAnomalyDetectionConfigWithDefaults instantiates a new ServiceAnomalyDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ServiceAnomalyDetectionConfig) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceAnomalyDetectionConfig) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceAnomalyDetectionConfig) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceAnomalyDetectionConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResponseTimeDegradation

`func (o *ServiceAnomalyDetectionConfig) GetResponseTimeDegradation() ResponseTimeDegradationDetectionConfig`

GetResponseTimeDegradation returns the ResponseTimeDegradation field if non-nil, zero value otherwise.

### GetResponseTimeDegradationOk

`func (o *ServiceAnomalyDetectionConfig) GetResponseTimeDegradationOk() (*ResponseTimeDegradationDetectionConfig, bool)`

GetResponseTimeDegradationOk returns a tuple with the ResponseTimeDegradation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeDegradation

`func (o *ServiceAnomalyDetectionConfig) SetResponseTimeDegradation(v ResponseTimeDegradationDetectionConfig)`

SetResponseTimeDegradation sets ResponseTimeDegradation field to given value.


### GetLoadDrop

`func (o *ServiceAnomalyDetectionConfig) GetLoadDrop() LoadDropDetectionConfig`

GetLoadDrop returns the LoadDrop field if non-nil, zero value otherwise.

### GetLoadDropOk

`func (o *ServiceAnomalyDetectionConfig) GetLoadDropOk() (*LoadDropDetectionConfig, bool)`

GetLoadDropOk returns a tuple with the LoadDrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadDrop

`func (o *ServiceAnomalyDetectionConfig) SetLoadDrop(v LoadDropDetectionConfig)`

SetLoadDrop sets LoadDrop field to given value.

### HasLoadDrop

`func (o *ServiceAnomalyDetectionConfig) HasLoadDrop() bool`

HasLoadDrop returns a boolean if a field has been set.

### GetLoadSpike

`func (o *ServiceAnomalyDetectionConfig) GetLoadSpike() LoadSpikeDetectionConfig`

GetLoadSpike returns the LoadSpike field if non-nil, zero value otherwise.

### GetLoadSpikeOk

`func (o *ServiceAnomalyDetectionConfig) GetLoadSpikeOk() (*LoadSpikeDetectionConfig, bool)`

GetLoadSpikeOk returns a tuple with the LoadSpike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadSpike

`func (o *ServiceAnomalyDetectionConfig) SetLoadSpike(v LoadSpikeDetectionConfig)`

SetLoadSpike sets LoadSpike field to given value.

### HasLoadSpike

`func (o *ServiceAnomalyDetectionConfig) HasLoadSpike() bool`

HasLoadSpike returns a boolean if a field has been set.

### GetFailureRateIncrease

`func (o *ServiceAnomalyDetectionConfig) GetFailureRateIncrease() FailureRateIncreaseDetectionConfig`

GetFailureRateIncrease returns the FailureRateIncrease field if non-nil, zero value otherwise.

### GetFailureRateIncreaseOk

`func (o *ServiceAnomalyDetectionConfig) GetFailureRateIncreaseOk() (*FailureRateIncreaseDetectionConfig, bool)`

GetFailureRateIncreaseOk returns a tuple with the FailureRateIncrease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureRateIncrease

`func (o *ServiceAnomalyDetectionConfig) SetFailureRateIncrease(v FailureRateIncreaseDetectionConfig)`

SetFailureRateIncrease sets FailureRateIncrease field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


