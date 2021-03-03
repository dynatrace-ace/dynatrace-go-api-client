# LoadingTimeThresholdsPolicyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Performance threshold is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**Thresholds** | Pointer to [**[]LoadingTimeThreshold**](LoadingTimeThreshold.md) | The list of performance threshold rules. | [optional] 

## Methods

### NewLoadingTimeThresholdsPolicyDto

`func NewLoadingTimeThresholdsPolicyDto(enabled bool, ) *LoadingTimeThresholdsPolicyDto`

NewLoadingTimeThresholdsPolicyDto instantiates a new LoadingTimeThresholdsPolicyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadingTimeThresholdsPolicyDtoWithDefaults

`func NewLoadingTimeThresholdsPolicyDtoWithDefaults() *LoadingTimeThresholdsPolicyDto`

NewLoadingTimeThresholdsPolicyDtoWithDefaults instantiates a new LoadingTimeThresholdsPolicyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *LoadingTimeThresholdsPolicyDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LoadingTimeThresholdsPolicyDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LoadingTimeThresholdsPolicyDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetThresholds

`func (o *LoadingTimeThresholdsPolicyDto) GetThresholds() []LoadingTimeThreshold`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *LoadingTimeThresholdsPolicyDto) GetThresholdsOk() (*[]LoadingTimeThreshold, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *LoadingTimeThresholdsPolicyDto) SetThresholds(v []LoadingTimeThreshold)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *LoadingTimeThresholdsPolicyDto) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


