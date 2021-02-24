# AvailabilityMonitoringPG

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | How to monitor the availability of the process group:   * &#x60;PROCESS_IMPACT&#x60;: Alert if any process of the group becomes unavailable.  * &#x60;MINIMUM_THRESHOLD&#x60;: Alert if the number of active processes in the group falls below the specified threshold.  * &#x60;OFF&#x60;: Availability monitoring is disabled. | 
**MinimumThreshold** | Pointer to **int32** | Alert if the number of active processes in the group is lower than this value. | [optional] 

## Methods

### NewAvailabilityMonitoringPG

`func NewAvailabilityMonitoringPG(method string, ) *AvailabilityMonitoringPG`

NewAvailabilityMonitoringPG instantiates a new AvailabilityMonitoringPG object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityMonitoringPGWithDefaults

`func NewAvailabilityMonitoringPGWithDefaults() *AvailabilityMonitoringPG`

NewAvailabilityMonitoringPGWithDefaults instantiates a new AvailabilityMonitoringPG object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *AvailabilityMonitoringPG) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AvailabilityMonitoringPG) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AvailabilityMonitoringPG) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetMinimumThreshold

`func (o *AvailabilityMonitoringPG) GetMinimumThreshold() int32`

GetMinimumThreshold returns the MinimumThreshold field if non-nil, zero value otherwise.

### GetMinimumThresholdOk

`func (o *AvailabilityMonitoringPG) GetMinimumThresholdOk() (*int32, bool)`

GetMinimumThresholdOk returns a tuple with the MinimumThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumThreshold

`func (o *AvailabilityMonitoringPG) SetMinimumThreshold(v int32)`

SetMinimumThreshold sets MinimumThreshold field to given value.

### HasMinimumThreshold

`func (o *AvailabilityMonitoringPG) HasMinimumThreshold() bool`

HasMinimumThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


