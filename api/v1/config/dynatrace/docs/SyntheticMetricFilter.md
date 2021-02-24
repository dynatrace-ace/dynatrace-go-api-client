# SyntheticMetricFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | Pointer to **string** | Only user actions of the specified type are included in the metric calculation. | [optional] 
**HasError** | Pointer to **bool** | The execution status of the monitors to be included in the metric calculation:   * &#x60;true&#x60;: Only failed executions are included.    * &#x60;false&#x60;: All executions are included. | [optional] 
**ErrorCode** | Pointer to **int32** | Only executions finished with the specified error code are included in the metric calculation. | [optional] 
**Event** | Pointer to **string** | Only the specified browser clickpath event is included in the metric calculation.    Specify the Dynatrace entity ID of the event here. You can fetch the list of clickpath events of the monitor with the [GET a synthetic monitor](https://www.dynatrace.com/support/help/shortlink/api-synthetic-monitors-get-monitor) request from the Environment API | [optional] 
**Location** | Pointer to **string** | Only executions from the specified location are included in the metric calculation.    Specify the Dynatrace entity ID of the location here. You can fetch the list of locations the monitor is running from with the [GET a synthetic monitor](https://www.dynatrace.com/support/help/shortlink/api-synthetic-monitors-get-monitor) request from the Environment API. | [optional] 

## Methods

### NewSyntheticMetricFilter

`func NewSyntheticMetricFilter() *SyntheticMetricFilter`

NewSyntheticMetricFilter instantiates a new SyntheticMetricFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticMetricFilterWithDefaults

`func NewSyntheticMetricFilterWithDefaults() *SyntheticMetricFilter`

NewSyntheticMetricFilterWithDefaults instantiates a new SyntheticMetricFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *SyntheticMetricFilter) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *SyntheticMetricFilter) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *SyntheticMetricFilter) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *SyntheticMetricFilter) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetHasError

`func (o *SyntheticMetricFilter) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *SyntheticMetricFilter) GetHasErrorOk() (*bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasError

`func (o *SyntheticMetricFilter) SetHasError(v bool)`

SetHasError sets HasError field to given value.

### HasHasError

`func (o *SyntheticMetricFilter) HasHasError() bool`

HasHasError returns a boolean if a field has been set.

### GetErrorCode

`func (o *SyntheticMetricFilter) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SyntheticMetricFilter) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *SyntheticMetricFilter) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *SyntheticMetricFilter) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetEvent

`func (o *SyntheticMetricFilter) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *SyntheticMetricFilter) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *SyntheticMetricFilter) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *SyntheticMetricFilter) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetLocation

`func (o *SyntheticMetricFilter) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SyntheticMetricFilter) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SyntheticMetricFilter) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SyntheticMetricFilter) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


