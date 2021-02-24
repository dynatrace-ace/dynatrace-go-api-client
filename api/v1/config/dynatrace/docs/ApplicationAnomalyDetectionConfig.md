# ApplicationAnomalyDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**ResponseTimeDegradation** | [**ResponseTimeDegradationDetectionConfig**](ResponseTimeDegradationDetectionConfig.md) |  | 
**TrafficDrop** | [**TrafficDropDetectionConfig**](TrafficDropDetectionConfig.md) |  | 
**TrafficSpike** | [**TrafficSpikeDetectionConfig**](TrafficSpikeDetectionConfig.md) |  | 
**FailureRateIncrease** | [**FailureRateIncreaseDetectionConfig**](FailureRateIncreaseDetectionConfig.md) |  | 

## Methods

### NewApplicationAnomalyDetectionConfig

`func NewApplicationAnomalyDetectionConfig(responseTimeDegradation ResponseTimeDegradationDetectionConfig, trafficDrop TrafficDropDetectionConfig, trafficSpike TrafficSpikeDetectionConfig, failureRateIncrease FailureRateIncreaseDetectionConfig, ) *ApplicationAnomalyDetectionConfig`

NewApplicationAnomalyDetectionConfig instantiates a new ApplicationAnomalyDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAnomalyDetectionConfigWithDefaults

`func NewApplicationAnomalyDetectionConfigWithDefaults() *ApplicationAnomalyDetectionConfig`

NewApplicationAnomalyDetectionConfigWithDefaults instantiates a new ApplicationAnomalyDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ApplicationAnomalyDetectionConfig) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationAnomalyDetectionConfig) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationAnomalyDetectionConfig) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationAnomalyDetectionConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResponseTimeDegradation

`func (o *ApplicationAnomalyDetectionConfig) GetResponseTimeDegradation() ResponseTimeDegradationDetectionConfig`

GetResponseTimeDegradation returns the ResponseTimeDegradation field if non-nil, zero value otherwise.

### GetResponseTimeDegradationOk

`func (o *ApplicationAnomalyDetectionConfig) GetResponseTimeDegradationOk() (*ResponseTimeDegradationDetectionConfig, bool)`

GetResponseTimeDegradationOk returns a tuple with the ResponseTimeDegradation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeDegradation

`func (o *ApplicationAnomalyDetectionConfig) SetResponseTimeDegradation(v ResponseTimeDegradationDetectionConfig)`

SetResponseTimeDegradation sets ResponseTimeDegradation field to given value.


### GetTrafficDrop

`func (o *ApplicationAnomalyDetectionConfig) GetTrafficDrop() TrafficDropDetectionConfig`

GetTrafficDrop returns the TrafficDrop field if non-nil, zero value otherwise.

### GetTrafficDropOk

`func (o *ApplicationAnomalyDetectionConfig) GetTrafficDropOk() (*TrafficDropDetectionConfig, bool)`

GetTrafficDropOk returns a tuple with the TrafficDrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDrop

`func (o *ApplicationAnomalyDetectionConfig) SetTrafficDrop(v TrafficDropDetectionConfig)`

SetTrafficDrop sets TrafficDrop field to given value.


### GetTrafficSpike

`func (o *ApplicationAnomalyDetectionConfig) GetTrafficSpike() TrafficSpikeDetectionConfig`

GetTrafficSpike returns the TrafficSpike field if non-nil, zero value otherwise.

### GetTrafficSpikeOk

`func (o *ApplicationAnomalyDetectionConfig) GetTrafficSpikeOk() (*TrafficSpikeDetectionConfig, bool)`

GetTrafficSpikeOk returns a tuple with the TrafficSpike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSpike

`func (o *ApplicationAnomalyDetectionConfig) SetTrafficSpike(v TrafficSpikeDetectionConfig)`

SetTrafficSpike sets TrafficSpike field to given value.


### GetFailureRateIncrease

`func (o *ApplicationAnomalyDetectionConfig) GetFailureRateIncrease() FailureRateIncreaseDetectionConfig`

GetFailureRateIncrease returns the FailureRateIncrease field if non-nil, zero value otherwise.

### GetFailureRateIncreaseOk

`func (o *ApplicationAnomalyDetectionConfig) GetFailureRateIncreaseOk() (*FailureRateIncreaseDetectionConfig, bool)`

GetFailureRateIncreaseOk returns a tuple with the FailureRateIncrease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureRateIncrease

`func (o *ApplicationAnomalyDetectionConfig) SetFailureRateIncrease(v FailureRateIncreaseDetectionConfig)`

SetFailureRateIncrease sets FailureRateIncrease field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


