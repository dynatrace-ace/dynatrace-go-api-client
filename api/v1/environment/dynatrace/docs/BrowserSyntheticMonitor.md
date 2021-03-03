# BrowserSyntheticMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyPerformanceMetrics** | Pointer to [**KeyPerformanceMetrics**](KeyPerformanceMetrics.md) |  | [optional] 
**Events** | Pointer to [**[]EventDto**](EventDto.md) | A list of events for this monitor | [optional] 

## Methods

### NewBrowserSyntheticMonitor

`func NewBrowserSyntheticMonitor() *BrowserSyntheticMonitor`

NewBrowserSyntheticMonitor instantiates a new BrowserSyntheticMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserSyntheticMonitorWithDefaults

`func NewBrowserSyntheticMonitorWithDefaults() *BrowserSyntheticMonitor`

NewBrowserSyntheticMonitorWithDefaults instantiates a new BrowserSyntheticMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyPerformanceMetrics

`func (o *BrowserSyntheticMonitor) GetKeyPerformanceMetrics() KeyPerformanceMetrics`

GetKeyPerformanceMetrics returns the KeyPerformanceMetrics field if non-nil, zero value otherwise.

### GetKeyPerformanceMetricsOk

`func (o *BrowserSyntheticMonitor) GetKeyPerformanceMetricsOk() (*KeyPerformanceMetrics, bool)`

GetKeyPerformanceMetricsOk returns a tuple with the KeyPerformanceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPerformanceMetrics

`func (o *BrowserSyntheticMonitor) SetKeyPerformanceMetrics(v KeyPerformanceMetrics)`

SetKeyPerformanceMetrics sets KeyPerformanceMetrics field to given value.

### HasKeyPerformanceMetrics

`func (o *BrowserSyntheticMonitor) HasKeyPerformanceMetrics() bool`

HasKeyPerformanceMetrics returns a boolean if a field has been set.

### GetEvents

`func (o *BrowserSyntheticMonitor) GetEvents() []EventDto`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *BrowserSyntheticMonitor) GetEventsOk() (*[]EventDto, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *BrowserSyntheticMonitor) SetEvents(v []EventDto)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *BrowserSyntheticMonitor) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


