# ResponseTimeDegradationThresholdConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseTimeThresholdMilliseconds** | **int32** | Response time during any 5-minute period to trigger an alert, in milliseconds. | 
**SlowestResponseTimeThresholdMilliseconds** | **int32** | Response time of the 10% slowest during any 5-minute period to trigger an alert, in milliseconds. | 
**LoadThreshold** | **string** | Minimal service load to detect response time degradation.    Response time degradation of services with smaller load won&#39;t trigger alerts. | 
**Sensitivity** | **string** | Sensitivity of the threshold.   With &#x60;low&#x60; sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won&#39;t trigger alerts.   With &#x60;high&#x60; sensitivity, no statistical confidence is used. Each violation triggers an alert. | 

## Methods

### NewResponseTimeDegradationThresholdConfig

`func NewResponseTimeDegradationThresholdConfig(responseTimeThresholdMilliseconds int32, slowestResponseTimeThresholdMilliseconds int32, loadThreshold string, sensitivity string, ) *ResponseTimeDegradationThresholdConfig`

NewResponseTimeDegradationThresholdConfig instantiates a new ResponseTimeDegradationThresholdConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTimeDegradationThresholdConfigWithDefaults

`func NewResponseTimeDegradationThresholdConfigWithDefaults() *ResponseTimeDegradationThresholdConfig`

NewResponseTimeDegradationThresholdConfigWithDefaults instantiates a new ResponseTimeDegradationThresholdConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseTimeThresholdMilliseconds

`func (o *ResponseTimeDegradationThresholdConfig) GetResponseTimeThresholdMilliseconds() int32`

GetResponseTimeThresholdMilliseconds returns the ResponseTimeThresholdMilliseconds field if non-nil, zero value otherwise.

### GetResponseTimeThresholdMillisecondsOk

`func (o *ResponseTimeDegradationThresholdConfig) GetResponseTimeThresholdMillisecondsOk() (*int32, bool)`

GetResponseTimeThresholdMillisecondsOk returns a tuple with the ResponseTimeThresholdMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeThresholdMilliseconds

`func (o *ResponseTimeDegradationThresholdConfig) SetResponseTimeThresholdMilliseconds(v int32)`

SetResponseTimeThresholdMilliseconds sets ResponseTimeThresholdMilliseconds field to given value.


### GetSlowestResponseTimeThresholdMilliseconds

`func (o *ResponseTimeDegradationThresholdConfig) GetSlowestResponseTimeThresholdMilliseconds() int32`

GetSlowestResponseTimeThresholdMilliseconds returns the SlowestResponseTimeThresholdMilliseconds field if non-nil, zero value otherwise.

### GetSlowestResponseTimeThresholdMillisecondsOk

`func (o *ResponseTimeDegradationThresholdConfig) GetSlowestResponseTimeThresholdMillisecondsOk() (*int32, bool)`

GetSlowestResponseTimeThresholdMillisecondsOk returns a tuple with the SlowestResponseTimeThresholdMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowestResponseTimeThresholdMilliseconds

`func (o *ResponseTimeDegradationThresholdConfig) SetSlowestResponseTimeThresholdMilliseconds(v int32)`

SetSlowestResponseTimeThresholdMilliseconds sets SlowestResponseTimeThresholdMilliseconds field to given value.


### GetLoadThreshold

`func (o *ResponseTimeDegradationThresholdConfig) GetLoadThreshold() string`

GetLoadThreshold returns the LoadThreshold field if non-nil, zero value otherwise.

### GetLoadThresholdOk

`func (o *ResponseTimeDegradationThresholdConfig) GetLoadThresholdOk() (*string, bool)`

GetLoadThresholdOk returns a tuple with the LoadThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadThreshold

`func (o *ResponseTimeDegradationThresholdConfig) SetLoadThreshold(v string)`

SetLoadThreshold sets LoadThreshold field to given value.


### GetSensitivity

`func (o *ResponseTimeDegradationThresholdConfig) GetSensitivity() string`

GetSensitivity returns the Sensitivity field if non-nil, zero value otherwise.

### GetSensitivityOk

`func (o *ResponseTimeDegradationThresholdConfig) GetSensitivityOk() (*string, bool)`

GetSensitivityOk returns a tuple with the Sensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitivity

`func (o *ResponseTimeDegradationThresholdConfig) SetSensitivity(v string)`

SetSensitivity sets Sensitivity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


