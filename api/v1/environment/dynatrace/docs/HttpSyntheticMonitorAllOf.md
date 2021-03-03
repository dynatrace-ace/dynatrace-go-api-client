# HttpSyntheticMonitorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to [**[]RequestDto**](RequestDto.md) | A list of events for this monitor | [optional] 

## Methods

### NewHttpSyntheticMonitorAllOf

`func NewHttpSyntheticMonitorAllOf() *HttpSyntheticMonitorAllOf`

NewHttpSyntheticMonitorAllOf instantiates a new HttpSyntheticMonitorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpSyntheticMonitorAllOfWithDefaults

`func NewHttpSyntheticMonitorAllOfWithDefaults() *HttpSyntheticMonitorAllOf`

NewHttpSyntheticMonitorAllOfWithDefaults instantiates a new HttpSyntheticMonitorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *HttpSyntheticMonitorAllOf) GetRequests() []RequestDto`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *HttpSyntheticMonitorAllOf) GetRequestsOk() (*[]RequestDto, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *HttpSyntheticMonitorAllOf) SetRequests(v []RequestDto)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *HttpSyntheticMonitorAllOf) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


