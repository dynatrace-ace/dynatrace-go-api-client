# LoadingTimeThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the threshold: total loading time or action loading time. | 
**ValueMs** | **int32** | Notify if monitor takes longer than *X* milliseconds to load. | 
**RequestIndex** | Pointer to **int32** | Specify the request to which an ACTION threshold applies. | [optional] 
**EventIndex** | Pointer to **int32** | Specify the event to which an ACTION threshold applies. | [optional] 

## Methods

### NewLoadingTimeThreshold

`func NewLoadingTimeThreshold(type_ string, valueMs int32, ) *LoadingTimeThreshold`

NewLoadingTimeThreshold instantiates a new LoadingTimeThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadingTimeThresholdWithDefaults

`func NewLoadingTimeThresholdWithDefaults() *LoadingTimeThreshold`

NewLoadingTimeThresholdWithDefaults instantiates a new LoadingTimeThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LoadingTimeThreshold) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoadingTimeThreshold) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoadingTimeThreshold) SetType(v string)`

SetType sets Type field to given value.


### GetValueMs

`func (o *LoadingTimeThreshold) GetValueMs() int32`

GetValueMs returns the ValueMs field if non-nil, zero value otherwise.

### GetValueMsOk

`func (o *LoadingTimeThreshold) GetValueMsOk() (*int32, bool)`

GetValueMsOk returns a tuple with the ValueMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMs

`func (o *LoadingTimeThreshold) SetValueMs(v int32)`

SetValueMs sets ValueMs field to given value.


### GetRequestIndex

`func (o *LoadingTimeThreshold) GetRequestIndex() int32`

GetRequestIndex returns the RequestIndex field if non-nil, zero value otherwise.

### GetRequestIndexOk

`func (o *LoadingTimeThreshold) GetRequestIndexOk() (*int32, bool)`

GetRequestIndexOk returns a tuple with the RequestIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestIndex

`func (o *LoadingTimeThreshold) SetRequestIndex(v int32)`

SetRequestIndex sets RequestIndex field to given value.

### HasRequestIndex

`func (o *LoadingTimeThreshold) HasRequestIndex() bool`

HasRequestIndex returns a boolean if a field has been set.

### GetEventIndex

`func (o *LoadingTimeThreshold) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *LoadingTimeThreshold) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *LoadingTimeThreshold) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.

### HasEventIndex

`func (o *LoadingTimeThreshold) HasEventIndex() bool`

HasEventIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


