# TimeseriesQueryMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeseriesId** | Pointer to **string** | The case-sensitive ID of the metric, where you want to read data points.   You can execute a GET timeseries request, to obtain the list of available metrics. | [optional] 
**AggregationType** | Pointer to **string** | The aggregation type for the resulting data points.   If the requested metric doesn&#39;t support the specified aggregation, the request will result in an error. | [optional] 
**StartTimestamp** | Pointer to **int64** | The start timestamp of the timeframe, in UTC milliseconds. | [optional] 
**EndTimestamp** | Pointer to **int64** | The start timestamp of the timeframe, in UTC milliseconds.   If later than the current time, Dynatrace automatically uses current time instead.   The timeframe must not exceed 6 months. | [optional] 
**Predict** | Pointer to **bool** | The flag to predict future data points. | [optional] 
**RelativeTime** | Pointer to **string** | The relative timeframe, back from the current time. | [optional] 
**QueryMode** | Pointer to **string** | Defines the type of result that the call should return. Valid result modes are:  &#x60;series&#x60;: returns all the data points of the metric in the specified timeframe.  &#x60;total&#x60;: returns one scalar value for the specified timeframe.   By default, the &#x60;series&#x60; mode is used. | [optional] 
**Entities** | Pointer to **[]string** | Filters requested data points by entities which should deliver them. You can specify several entities at once.   Allowed values are Dynatrace entity IDs.   If the selected entity doesn&#39;t support the requested metric, the request will result in an error. | [optional] 
**Tags** | Pointer to **[]string** | Filters requiested data points by entity which should deliver them. Only data from entities with the specified tag is delivered.   You can specify several tags in the following format: &#x60;tags&#x3D;tag1&amp;tags&#x3D;tag2&#x60;. The entity has to match *all* the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags use following format: &#x60;[context]key:value&#x60;. | [optional] 
**Filters** | Pointer to **map[string]string** | A filter is an object, containing map of filter keys and its values. Valid filter keys are:   &#x60;processType&#x60;: Filters by process type. See Process types for allowed values. &#x60;osType&#x60;: Filters by operating system. See OS types for allowed values. &#x60;serviceType&#x60;: Filters by service type. See Service types for allowed values. &#x60;technology&#x60;: Filters by technology type. See Technology types for allowed values. &#x60;webServiceName&#x60;: Filters by web service name. &#x60;webServiceNamespace&#x60;: Filters by the web service namespace. &#x60;host&#x60;: Filters by entity ID of the host, for example HOST-007. | [optional] 
**Percentile** | Pointer to **int32** | Specifies which percentile of the selected response time metric should be delivered.  Only applicable to the &#x60;PERCENTILE&#x60; aggregation type.   Valid values for percentile are between 1 and 99.   Please keep in mind that percentile export is only possible for response-time based metrics such as application and service response times. | [optional] 
**IncludeParentIds** | Pointer to **bool** | Specifies whether the results should exposes dimension mappings between parent entities and their children.  For instance: SERVICE-0000000000000001, SERVICE_METHOD-0000000000000001 | [optional] 
**ConsiderMaintenanceWindowsForAvailability** | Pointer to **bool** | Exclude (&#x60;true&#x60;) or include (&#x60;false&#x60;) data points from any [maintenance window](https://dt-url.net/b2123rg0), defined in your environment. | [optional] 

## Methods

### NewTimeseriesQueryMessage

`func NewTimeseriesQueryMessage() *TimeseriesQueryMessage`

NewTimeseriesQueryMessage instantiates a new TimeseriesQueryMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesQueryMessageWithDefaults

`func NewTimeseriesQueryMessageWithDefaults() *TimeseriesQueryMessage`

NewTimeseriesQueryMessageWithDefaults instantiates a new TimeseriesQueryMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeseriesId

`func (o *TimeseriesQueryMessage) GetTimeseriesId() string`

GetTimeseriesId returns the TimeseriesId field if non-nil, zero value otherwise.

### GetTimeseriesIdOk

`func (o *TimeseriesQueryMessage) GetTimeseriesIdOk() (*string, bool)`

GetTimeseriesIdOk returns a tuple with the TimeseriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseriesId

`func (o *TimeseriesQueryMessage) SetTimeseriesId(v string)`

SetTimeseriesId sets TimeseriesId field to given value.

### HasTimeseriesId

`func (o *TimeseriesQueryMessage) HasTimeseriesId() bool`

HasTimeseriesId returns a boolean if a field has been set.

### GetAggregationType

`func (o *TimeseriesQueryMessage) GetAggregationType() string`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *TimeseriesQueryMessage) GetAggregationTypeOk() (*string, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *TimeseriesQueryMessage) SetAggregationType(v string)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *TimeseriesQueryMessage) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *TimeseriesQueryMessage) GetStartTimestamp() int64`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *TimeseriesQueryMessage) GetStartTimestampOk() (*int64, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *TimeseriesQueryMessage) SetStartTimestamp(v int64)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *TimeseriesQueryMessage) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetEndTimestamp

`func (o *TimeseriesQueryMessage) GetEndTimestamp() int64`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *TimeseriesQueryMessage) GetEndTimestampOk() (*int64, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *TimeseriesQueryMessage) SetEndTimestamp(v int64)`

SetEndTimestamp sets EndTimestamp field to given value.

### HasEndTimestamp

`func (o *TimeseriesQueryMessage) HasEndTimestamp() bool`

HasEndTimestamp returns a boolean if a field has been set.

### GetPredict

`func (o *TimeseriesQueryMessage) GetPredict() bool`

GetPredict returns the Predict field if non-nil, zero value otherwise.

### GetPredictOk

`func (o *TimeseriesQueryMessage) GetPredictOk() (*bool, bool)`

GetPredictOk returns a tuple with the Predict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredict

`func (o *TimeseriesQueryMessage) SetPredict(v bool)`

SetPredict sets Predict field to given value.

### HasPredict

`func (o *TimeseriesQueryMessage) HasPredict() bool`

HasPredict returns a boolean if a field has been set.

### GetRelativeTime

`func (o *TimeseriesQueryMessage) GetRelativeTime() string`

GetRelativeTime returns the RelativeTime field if non-nil, zero value otherwise.

### GetRelativeTimeOk

`func (o *TimeseriesQueryMessage) GetRelativeTimeOk() (*string, bool)`

GetRelativeTimeOk returns a tuple with the RelativeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeTime

`func (o *TimeseriesQueryMessage) SetRelativeTime(v string)`

SetRelativeTime sets RelativeTime field to given value.

### HasRelativeTime

`func (o *TimeseriesQueryMessage) HasRelativeTime() bool`

HasRelativeTime returns a boolean if a field has been set.

### GetQueryMode

`func (o *TimeseriesQueryMessage) GetQueryMode() string`

GetQueryMode returns the QueryMode field if non-nil, zero value otherwise.

### GetQueryModeOk

`func (o *TimeseriesQueryMessage) GetQueryModeOk() (*string, bool)`

GetQueryModeOk returns a tuple with the QueryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryMode

`func (o *TimeseriesQueryMessage) SetQueryMode(v string)`

SetQueryMode sets QueryMode field to given value.

### HasQueryMode

`func (o *TimeseriesQueryMessage) HasQueryMode() bool`

HasQueryMode returns a boolean if a field has been set.

### GetEntities

`func (o *TimeseriesQueryMessage) GetEntities() []string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *TimeseriesQueryMessage) GetEntitiesOk() (*[]string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *TimeseriesQueryMessage) SetEntities(v []string)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *TimeseriesQueryMessage) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetTags

`func (o *TimeseriesQueryMessage) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TimeseriesQueryMessage) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TimeseriesQueryMessage) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TimeseriesQueryMessage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFilters

`func (o *TimeseriesQueryMessage) GetFilters() map[string]string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *TimeseriesQueryMessage) GetFiltersOk() (*map[string]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *TimeseriesQueryMessage) SetFilters(v map[string]string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *TimeseriesQueryMessage) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetPercentile

`func (o *TimeseriesQueryMessage) GetPercentile() int32`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *TimeseriesQueryMessage) GetPercentileOk() (*int32, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *TimeseriesQueryMessage) SetPercentile(v int32)`

SetPercentile sets Percentile field to given value.

### HasPercentile

`func (o *TimeseriesQueryMessage) HasPercentile() bool`

HasPercentile returns a boolean if a field has been set.

### GetIncludeParentIds

`func (o *TimeseriesQueryMessage) GetIncludeParentIds() bool`

GetIncludeParentIds returns the IncludeParentIds field if non-nil, zero value otherwise.

### GetIncludeParentIdsOk

`func (o *TimeseriesQueryMessage) GetIncludeParentIdsOk() (*bool, bool)`

GetIncludeParentIdsOk returns a tuple with the IncludeParentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeParentIds

`func (o *TimeseriesQueryMessage) SetIncludeParentIds(v bool)`

SetIncludeParentIds sets IncludeParentIds field to given value.

### HasIncludeParentIds

`func (o *TimeseriesQueryMessage) HasIncludeParentIds() bool`

HasIncludeParentIds returns a boolean if a field has been set.

### GetConsiderMaintenanceWindowsForAvailability

`func (o *TimeseriesQueryMessage) GetConsiderMaintenanceWindowsForAvailability() bool`

GetConsiderMaintenanceWindowsForAvailability returns the ConsiderMaintenanceWindowsForAvailability field if non-nil, zero value otherwise.

### GetConsiderMaintenanceWindowsForAvailabilityOk

`func (o *TimeseriesQueryMessage) GetConsiderMaintenanceWindowsForAvailabilityOk() (*bool, bool)`

GetConsiderMaintenanceWindowsForAvailabilityOk returns a tuple with the ConsiderMaintenanceWindowsForAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsiderMaintenanceWindowsForAvailability

`func (o *TimeseriesQueryMessage) SetConsiderMaintenanceWindowsForAvailability(v bool)`

SetConsiderMaintenanceWindowsForAvailability sets ConsiderMaintenanceWindowsForAvailability field to given value.

### HasConsiderMaintenanceWindowsForAvailability

`func (o *TimeseriesQueryMessage) HasConsiderMaintenanceWindowsForAvailability() bool`

HasConsiderMaintenanceWindowsForAvailability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


