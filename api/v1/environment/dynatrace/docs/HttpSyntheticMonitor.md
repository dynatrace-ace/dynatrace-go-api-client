# HttpSyntheticMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to [**[]RequestDto**](RequestDto.md) | A list of events for this monitor | [optional] 

## Methods

### NewHttpSyntheticMonitor

`func NewHttpSyntheticMonitor() *HttpSyntheticMonitor`

NewHttpSyntheticMonitor instantiates a new HttpSyntheticMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpSyntheticMonitorWithDefaults

`func NewHttpSyntheticMonitorWithDefaults() *HttpSyntheticMonitor`

NewHttpSyntheticMonitorWithDefaults instantiates a new HttpSyntheticMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *HttpSyntheticMonitor) GetRequests() []RequestDto`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *HttpSyntheticMonitor) GetRequestsOk() (*[]RequestDto, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *HttpSyntheticMonitor) SetRequests(v []RequestDto)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *HttpSyntheticMonitor) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


