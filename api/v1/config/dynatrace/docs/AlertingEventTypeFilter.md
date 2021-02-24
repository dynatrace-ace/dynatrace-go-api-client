# AlertingEventTypeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PredefinedEventFilter** | Pointer to [**AlertingPredefinedEventFilter**](AlertingPredefinedEventFilter.md) |  | [optional] 
**CustomEventFilter** | Pointer to [**AlertingCustomEventFilter**](AlertingCustomEventFilter.md) |  | [optional] 

## Methods

### NewAlertingEventTypeFilter

`func NewAlertingEventTypeFilter() *AlertingEventTypeFilter`

NewAlertingEventTypeFilter instantiates a new AlertingEventTypeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingEventTypeFilterWithDefaults

`func NewAlertingEventTypeFilterWithDefaults() *AlertingEventTypeFilter`

NewAlertingEventTypeFilterWithDefaults instantiates a new AlertingEventTypeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPredefinedEventFilter

`func (o *AlertingEventTypeFilter) GetPredefinedEventFilter() AlertingPredefinedEventFilter`

GetPredefinedEventFilter returns the PredefinedEventFilter field if non-nil, zero value otherwise.

### GetPredefinedEventFilterOk

`func (o *AlertingEventTypeFilter) GetPredefinedEventFilterOk() (*AlertingPredefinedEventFilter, bool)`

GetPredefinedEventFilterOk returns a tuple with the PredefinedEventFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedEventFilter

`func (o *AlertingEventTypeFilter) SetPredefinedEventFilter(v AlertingPredefinedEventFilter)`

SetPredefinedEventFilter sets PredefinedEventFilter field to given value.

### HasPredefinedEventFilter

`func (o *AlertingEventTypeFilter) HasPredefinedEventFilter() bool`

HasPredefinedEventFilter returns a boolean if a field has been set.

### GetCustomEventFilter

`func (o *AlertingEventTypeFilter) GetCustomEventFilter() AlertingCustomEventFilter`

GetCustomEventFilter returns the CustomEventFilter field if non-nil, zero value otherwise.

### GetCustomEventFilterOk

`func (o *AlertingEventTypeFilter) GetCustomEventFilterOk() (*AlertingCustomEventFilter, bool)`

GetCustomEventFilterOk returns a tuple with the CustomEventFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEventFilter

`func (o *AlertingEventTypeFilter) SetCustomEventFilter(v AlertingCustomEventFilter)`

SetCustomEventFilter sets CustomEventFilter field to given value.

### HasCustomEventFilter

`func (o *AlertingEventTypeFilter) HasCustomEventFilter() bool`

HasCustomEventFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


