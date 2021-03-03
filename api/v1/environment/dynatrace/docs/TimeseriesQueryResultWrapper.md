# TimeseriesQueryResultWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**TimeseriesDataPointQueryResult**](TimeseriesDataPointQueryResult.md) |  | [optional] 

## Methods

### NewTimeseriesQueryResultWrapper

`func NewTimeseriesQueryResultWrapper() *TimeseriesQueryResultWrapper`

NewTimeseriesQueryResultWrapper instantiates a new TimeseriesQueryResultWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesQueryResultWrapperWithDefaults

`func NewTimeseriesQueryResultWrapperWithDefaults() *TimeseriesQueryResultWrapper`

NewTimeseriesQueryResultWrapperWithDefaults instantiates a new TimeseriesQueryResultWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *TimeseriesQueryResultWrapper) GetResult() TimeseriesDataPointQueryResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TimeseriesQueryResultWrapper) GetResultOk() (*TimeseriesDataPointQueryResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TimeseriesQueryResultWrapper) SetResult(v TimeseriesDataPointQueryResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *TimeseriesQueryResultWrapper) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


