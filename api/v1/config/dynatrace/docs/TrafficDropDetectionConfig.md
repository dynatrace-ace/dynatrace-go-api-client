# TrafficDropDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**TrafficDropPercent** | Pointer to **int32** | Alert if the observed traffic is less than *X* % of the expected value. | [optional] 

## Methods

### NewTrafficDropDetectionConfig

`func NewTrafficDropDetectionConfig(enabled bool, ) *TrafficDropDetectionConfig`

NewTrafficDropDetectionConfig instantiates a new TrafficDropDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficDropDetectionConfigWithDefaults

`func NewTrafficDropDetectionConfigWithDefaults() *TrafficDropDetectionConfig`

NewTrafficDropDetectionConfigWithDefaults instantiates a new TrafficDropDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *TrafficDropDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TrafficDropDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TrafficDropDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetTrafficDropPercent

`func (o *TrafficDropDetectionConfig) GetTrafficDropPercent() int32`

GetTrafficDropPercent returns the TrafficDropPercent field if non-nil, zero value otherwise.

### GetTrafficDropPercentOk

`func (o *TrafficDropDetectionConfig) GetTrafficDropPercentOk() (*int32, bool)`

GetTrafficDropPercentOk returns a tuple with the TrafficDropPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDropPercent

`func (o *TrafficDropDetectionConfig) SetTrafficDropPercent(v int32)`

SetTrafficDropPercent sets TrafficDropPercent field to given value.

### HasTrafficDropPercent

`func (o *TrafficDropDetectionConfig) HasTrafficDropPercent() bool`

HasTrafficDropPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


