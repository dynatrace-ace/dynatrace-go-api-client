# OutageHandlingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalOutage** | **bool** | When enabled (&#x60;true&#x60;), generate a problem and send an alert when the monitor is unavailable at all configured locations. | 
**LocalOutage** | **bool** | When enabled (&#x60;true&#x60;), generate a problem and send an alert when the monitor is unavailable for one or more consecutive runs at any location. | 
**LocalOutagePolicy** | [**LocalOutagePolicy**](LocalOutagePolicy.md) |  | 
**RetryOnError** | Pointer to **bool** | Schedule retry if browser monitor execution results in a fail. For HTTP monitors this property is ignored. | [optional] [default to true]

## Methods

### NewOutageHandlingPolicy

`func NewOutageHandlingPolicy(globalOutage bool, localOutage bool, localOutagePolicy LocalOutagePolicy, ) *OutageHandlingPolicy`

NewOutageHandlingPolicy instantiates a new OutageHandlingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutageHandlingPolicyWithDefaults

`func NewOutageHandlingPolicyWithDefaults() *OutageHandlingPolicy`

NewOutageHandlingPolicyWithDefaults instantiates a new OutageHandlingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalOutage

`func (o *OutageHandlingPolicy) GetGlobalOutage() bool`

GetGlobalOutage returns the GlobalOutage field if non-nil, zero value otherwise.

### GetGlobalOutageOk

`func (o *OutageHandlingPolicy) GetGlobalOutageOk() (*bool, bool)`

GetGlobalOutageOk returns a tuple with the GlobalOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalOutage

`func (o *OutageHandlingPolicy) SetGlobalOutage(v bool)`

SetGlobalOutage sets GlobalOutage field to given value.


### GetLocalOutage

`func (o *OutageHandlingPolicy) GetLocalOutage() bool`

GetLocalOutage returns the LocalOutage field if non-nil, zero value otherwise.

### GetLocalOutageOk

`func (o *OutageHandlingPolicy) GetLocalOutageOk() (*bool, bool)`

GetLocalOutageOk returns a tuple with the LocalOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalOutage

`func (o *OutageHandlingPolicy) SetLocalOutage(v bool)`

SetLocalOutage sets LocalOutage field to given value.


### GetLocalOutagePolicy

`func (o *OutageHandlingPolicy) GetLocalOutagePolicy() LocalOutagePolicy`

GetLocalOutagePolicy returns the LocalOutagePolicy field if non-nil, zero value otherwise.

### GetLocalOutagePolicyOk

`func (o *OutageHandlingPolicy) GetLocalOutagePolicyOk() (*LocalOutagePolicy, bool)`

GetLocalOutagePolicyOk returns a tuple with the LocalOutagePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalOutagePolicy

`func (o *OutageHandlingPolicy) SetLocalOutagePolicy(v LocalOutagePolicy)`

SetLocalOutagePolicy sets LocalOutagePolicy field to given value.


### GetRetryOnError

`func (o *OutageHandlingPolicy) GetRetryOnError() bool`

GetRetryOnError returns the RetryOnError field if non-nil, zero value otherwise.

### GetRetryOnErrorOk

`func (o *OutageHandlingPolicy) GetRetryOnErrorOk() (*bool, bool)`

GetRetryOnErrorOk returns a tuple with the RetryOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryOnError

`func (o *OutageHandlingPolicy) SetRetryOnError(v bool)`

SetRetryOnError sets RetryOnError field to given value.

### HasRetryOnError

`func (o *OutageHandlingPolicy) HasRetryOnError() bool`

HasRetryOnError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


