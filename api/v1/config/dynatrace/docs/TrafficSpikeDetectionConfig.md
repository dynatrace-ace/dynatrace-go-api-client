# TrafficSpikeDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**TrafficSpikePercent** | Pointer to **int32** | Alert if the observed traffic is more than *X* % of the expected value. | [optional] 

## Methods

### NewTrafficSpikeDetectionConfig

`func NewTrafficSpikeDetectionConfig(enabled bool, ) *TrafficSpikeDetectionConfig`

NewTrafficSpikeDetectionConfig instantiates a new TrafficSpikeDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficSpikeDetectionConfigWithDefaults

`func NewTrafficSpikeDetectionConfigWithDefaults() *TrafficSpikeDetectionConfig`

NewTrafficSpikeDetectionConfigWithDefaults instantiates a new TrafficSpikeDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *TrafficSpikeDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TrafficSpikeDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TrafficSpikeDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetTrafficSpikePercent

`func (o *TrafficSpikeDetectionConfig) GetTrafficSpikePercent() int32`

GetTrafficSpikePercent returns the TrafficSpikePercent field if non-nil, zero value otherwise.

### GetTrafficSpikePercentOk

`func (o *TrafficSpikeDetectionConfig) GetTrafficSpikePercentOk() (*int32, bool)`

GetTrafficSpikePercentOk returns a tuple with the TrafficSpikePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSpikePercent

`func (o *TrafficSpikeDetectionConfig) SetTrafficSpikePercent(v int32)`

SetTrafficSpikePercent sets TrafficSpikePercent field to given value.

### HasTrafficSpikePercent

`func (o *TrafficSpikeDetectionConfig) HasTrafficSpikePercent() bool`

HasTrafficSpikePercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


