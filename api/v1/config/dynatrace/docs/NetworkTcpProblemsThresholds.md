# NetworkTcpProblemsThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewConnectionFailuresPercentage** | **int32** | Percentage of new connection failures is higher than *X*% in 3 out of 5 samples. | 
**FailedConnectionsNumberPerMinute** | **int32** | Number of failed connections is higher than *X* connections per minute in 3 out of 5 samples. | 

## Methods

### NewNetworkTcpProblemsThresholds

`func NewNetworkTcpProblemsThresholds(newConnectionFailuresPercentage int32, failedConnectionsNumberPerMinute int32, ) *NetworkTcpProblemsThresholds`

NewNetworkTcpProblemsThresholds instantiates a new NetworkTcpProblemsThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTcpProblemsThresholdsWithDefaults

`func NewNetworkTcpProblemsThresholdsWithDefaults() *NetworkTcpProblemsThresholds`

NewNetworkTcpProblemsThresholdsWithDefaults instantiates a new NetworkTcpProblemsThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewConnectionFailuresPercentage

`func (o *NetworkTcpProblemsThresholds) GetNewConnectionFailuresPercentage() int32`

GetNewConnectionFailuresPercentage returns the NewConnectionFailuresPercentage field if non-nil, zero value otherwise.

### GetNewConnectionFailuresPercentageOk

`func (o *NetworkTcpProblemsThresholds) GetNewConnectionFailuresPercentageOk() (*int32, bool)`

GetNewConnectionFailuresPercentageOk returns a tuple with the NewConnectionFailuresPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewConnectionFailuresPercentage

`func (o *NetworkTcpProblemsThresholds) SetNewConnectionFailuresPercentage(v int32)`

SetNewConnectionFailuresPercentage sets NewConnectionFailuresPercentage field to given value.


### GetFailedConnectionsNumberPerMinute

`func (o *NetworkTcpProblemsThresholds) GetFailedConnectionsNumberPerMinute() int32`

GetFailedConnectionsNumberPerMinute returns the FailedConnectionsNumberPerMinute field if non-nil, zero value otherwise.

### GetFailedConnectionsNumberPerMinuteOk

`func (o *NetworkTcpProblemsThresholds) GetFailedConnectionsNumberPerMinuteOk() (*int32, bool)`

GetFailedConnectionsNumberPerMinuteOk returns a tuple with the FailedConnectionsNumberPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedConnectionsNumberPerMinute

`func (o *NetworkTcpProblemsThresholds) SetFailedConnectionsNumberPerMinute(v int32)`

SetFailedConnectionsNumberPerMinute sets FailedConnectionsNumberPerMinute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


