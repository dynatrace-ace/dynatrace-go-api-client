# ResponseTimeDegradationAutodetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseTimeDegradationMilliseconds** | **int32** | Alert if the response time degrades by more than *X* milliseconds. | 
**ResponseTimeDegradationPercent** | **int32** | Alert if the response time degrades by more than *X* %. | 
**SlowestResponseTimeDegradationMilliseconds** | **int32** | Alert if the response time of the slowest 10% degrades by more than *X* milliseconds. | 
**SlowestResponseTimeDegradationPercent** | **int32** | Alert if the response time of the slowest 10% degrades by more than *X* %. | 
**LoadThreshold** | **string** | Minimal service load to detect response time degradation.    Response time degradation of services with smaller load won&#39;t trigger alerts. | 

## Methods

### NewResponseTimeDegradationAutodetectionConfig

`func NewResponseTimeDegradationAutodetectionConfig(responseTimeDegradationMilliseconds int32, responseTimeDegradationPercent int32, slowestResponseTimeDegradationMilliseconds int32, slowestResponseTimeDegradationPercent int32, loadThreshold string, ) *ResponseTimeDegradationAutodetectionConfig`

NewResponseTimeDegradationAutodetectionConfig instantiates a new ResponseTimeDegradationAutodetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTimeDegradationAutodetectionConfigWithDefaults

`func NewResponseTimeDegradationAutodetectionConfigWithDefaults() *ResponseTimeDegradationAutodetectionConfig`

NewResponseTimeDegradationAutodetectionConfigWithDefaults instantiates a new ResponseTimeDegradationAutodetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseTimeDegradationMilliseconds

`func (o *ResponseTimeDegradationAutodetectionConfig) GetResponseTimeDegradationMilliseconds() int32`

GetResponseTimeDegradationMilliseconds returns the ResponseTimeDegradationMilliseconds field if non-nil, zero value otherwise.

### GetResponseTimeDegradationMillisecondsOk

`func (o *ResponseTimeDegradationAutodetectionConfig) GetResponseTimeDegradationMillisecondsOk() (*int32, bool)`

GetResponseTimeDegradationMillisecondsOk returns a tuple with the ResponseTimeDegradationMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeDegradationMilliseconds

`func (o *ResponseTimeDegradationAutodetectionConfig) SetResponseTimeDegradationMilliseconds(v int32)`

SetResponseTimeDegradationMilliseconds sets ResponseTimeDegradationMilliseconds field to given value.


### GetResponseTimeDegradationPercent

`func (o *ResponseTimeDegradationAutodetectionConfig) GetResponseTimeDegradationPercent() int32`

GetResponseTimeDegradationPercent returns the ResponseTimeDegradationPercent field if non-nil, zero value otherwise.

### GetResponseTimeDegradationPercentOk

`func (o *ResponseTimeDegradationAutodetectionConfig) GetResponseTimeDegradationPercentOk() (*int32, bool)`

GetResponseTimeDegradationPercentOk returns a tuple with the ResponseTimeDegradationPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeDegradationPercent

`func (o *ResponseTimeDegradationAutodetectionConfig) SetResponseTimeDegradationPercent(v int32)`

SetResponseTimeDegradationPercent sets ResponseTimeDegradationPercent field to given value.


### GetSlowestResponseTimeDegradationMilliseconds

`func (o *ResponseTimeDegradationAutodetectionConfig) GetSlowestResponseTimeDegradationMilliseconds() int32`

GetSlowestResponseTimeDegradationMilliseconds returns the SlowestResponseTimeDegradationMilliseconds field if non-nil, zero value otherwise.

### GetSlowestResponseTimeDegradationMillisecondsOk

`func (o *ResponseTimeDegradationAutodetectionConfig) GetSlowestResponseTimeDegradationMillisecondsOk() (*int32, bool)`

GetSlowestResponseTimeDegradationMillisecondsOk returns a tuple with the SlowestResponseTimeDegradationMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowestResponseTimeDegradationMilliseconds

`func (o *ResponseTimeDegradationAutodetectionConfig) SetSlowestResponseTimeDegradationMilliseconds(v int32)`

SetSlowestResponseTimeDegradationMilliseconds sets SlowestResponseTimeDegradationMilliseconds field to given value.


### GetSlowestResponseTimeDegradationPercent

`func (o *ResponseTimeDegradationAutodetectionConfig) GetSlowestResponseTimeDegradationPercent() int32`

GetSlowestResponseTimeDegradationPercent returns the SlowestResponseTimeDegradationPercent field if non-nil, zero value otherwise.

### GetSlowestResponseTimeDegradationPercentOk

`func (o *ResponseTimeDegradationAutodetectionConfig) GetSlowestResponseTimeDegradationPercentOk() (*int32, bool)`

GetSlowestResponseTimeDegradationPercentOk returns a tuple with the SlowestResponseTimeDegradationPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowestResponseTimeDegradationPercent

`func (o *ResponseTimeDegradationAutodetectionConfig) SetSlowestResponseTimeDegradationPercent(v int32)`

SetSlowestResponseTimeDegradationPercent sets SlowestResponseTimeDegradationPercent field to given value.


### GetLoadThreshold

`func (o *ResponseTimeDegradationAutodetectionConfig) GetLoadThreshold() string`

GetLoadThreshold returns the LoadThreshold field if non-nil, zero value otherwise.

### GetLoadThresholdOk

`func (o *ResponseTimeDegradationAutodetectionConfig) GetLoadThresholdOk() (*string, bool)`

GetLoadThresholdOk returns a tuple with the LoadThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadThreshold

`func (o *ResponseTimeDegradationAutodetectionConfig) SetLoadThreshold(v string)`

SetLoadThreshold sets LoadThreshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


