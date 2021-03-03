# TimeseriesQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeseriesId** | Pointer to **string** | The ID of the metric. | [optional] 
**DisplayName** | Pointer to **string** | The name of the metric in the user interface. | [optional] 
**Dimensions** | Pointer to **[]string** | The fine metric division, for example process group and process ID for some process-related metric. | [optional] 
**AggregationTypes** | Pointer to **[]string** | The list of allowed aggregations for this metric. | [optional] 
**Unit** | Pointer to **string** | The unit of the metric. | [optional] 
**Filter** | Pointer to **string** | The feature, where the metric originates. | [optional] 
**DetailedSource** | Pointer to **string** | The feature, where the metric originates. | [optional] 
**PluginId** | Pointer to **string** | The ID of the plugin, where the metric originates. | [optional] 
**Types** | Pointer to **[]string** | Technology type definition. Used to group metrics under a logical technology name. | [optional] 
**DataResult** | Pointer to [**TimeseriesDataPointQueryResult**](TimeseriesDataPointQueryResult.md) |  | [optional] 

## Methods

### NewTimeseriesQueryResult

`func NewTimeseriesQueryResult() *TimeseriesQueryResult`

NewTimeseriesQueryResult instantiates a new TimeseriesQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesQueryResultWithDefaults

`func NewTimeseriesQueryResultWithDefaults() *TimeseriesQueryResult`

NewTimeseriesQueryResultWithDefaults instantiates a new TimeseriesQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeseriesId

`func (o *TimeseriesQueryResult) GetTimeseriesId() string`

GetTimeseriesId returns the TimeseriesId field if non-nil, zero value otherwise.

### GetTimeseriesIdOk

`func (o *TimeseriesQueryResult) GetTimeseriesIdOk() (*string, bool)`

GetTimeseriesIdOk returns a tuple with the TimeseriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseriesId

`func (o *TimeseriesQueryResult) SetTimeseriesId(v string)`

SetTimeseriesId sets TimeseriesId field to given value.

### HasTimeseriesId

`func (o *TimeseriesQueryResult) HasTimeseriesId() bool`

HasTimeseriesId returns a boolean if a field has been set.

### GetDisplayName

`func (o *TimeseriesQueryResult) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TimeseriesQueryResult) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TimeseriesQueryResult) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TimeseriesQueryResult) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDimensions

`func (o *TimeseriesQueryResult) GetDimensions() []string`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *TimeseriesQueryResult) GetDimensionsOk() (*[]string, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *TimeseriesQueryResult) SetDimensions(v []string)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *TimeseriesQueryResult) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetAggregationTypes

`func (o *TimeseriesQueryResult) GetAggregationTypes() []string`

GetAggregationTypes returns the AggregationTypes field if non-nil, zero value otherwise.

### GetAggregationTypesOk

`func (o *TimeseriesQueryResult) GetAggregationTypesOk() (*[]string, bool)`

GetAggregationTypesOk returns a tuple with the AggregationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationTypes

`func (o *TimeseriesQueryResult) SetAggregationTypes(v []string)`

SetAggregationTypes sets AggregationTypes field to given value.

### HasAggregationTypes

`func (o *TimeseriesQueryResult) HasAggregationTypes() bool`

HasAggregationTypes returns a boolean if a field has been set.

### GetUnit

`func (o *TimeseriesQueryResult) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TimeseriesQueryResult) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TimeseriesQueryResult) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *TimeseriesQueryResult) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetFilter

`func (o *TimeseriesQueryResult) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TimeseriesQueryResult) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TimeseriesQueryResult) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TimeseriesQueryResult) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetDetailedSource

`func (o *TimeseriesQueryResult) GetDetailedSource() string`

GetDetailedSource returns the DetailedSource field if non-nil, zero value otherwise.

### GetDetailedSourceOk

`func (o *TimeseriesQueryResult) GetDetailedSourceOk() (*string, bool)`

GetDetailedSourceOk returns a tuple with the DetailedSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedSource

`func (o *TimeseriesQueryResult) SetDetailedSource(v string)`

SetDetailedSource sets DetailedSource field to given value.

### HasDetailedSource

`func (o *TimeseriesQueryResult) HasDetailedSource() bool`

HasDetailedSource returns a boolean if a field has been set.

### GetPluginId

`func (o *TimeseriesQueryResult) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *TimeseriesQueryResult) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *TimeseriesQueryResult) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.

### HasPluginId

`func (o *TimeseriesQueryResult) HasPluginId() bool`

HasPluginId returns a boolean if a field has been set.

### GetTypes

`func (o *TimeseriesQueryResult) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *TimeseriesQueryResult) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *TimeseriesQueryResult) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *TimeseriesQueryResult) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetDataResult

`func (o *TimeseriesQueryResult) GetDataResult() TimeseriesDataPointQueryResult`

GetDataResult returns the DataResult field if non-nil, zero value otherwise.

### GetDataResultOk

`func (o *TimeseriesQueryResult) GetDataResultOk() (*TimeseriesDataPointQueryResult, bool)`

GetDataResultOk returns a tuple with the DataResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataResult

`func (o *TimeseriesQueryResult) SetDataResult(v TimeseriesDataPointQueryResult)`

SetDataResult sets DataResult field to given value.

### HasDataResult

`func (o *TimeseriesQueryResult) HasDataResult() bool`

HasDataResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


