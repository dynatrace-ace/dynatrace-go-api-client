# TimeseriesDefinition

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

## Methods

### NewTimeseriesDefinition

`func NewTimeseriesDefinition() *TimeseriesDefinition`

NewTimeseriesDefinition instantiates a new TimeseriesDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesDefinitionWithDefaults

`func NewTimeseriesDefinitionWithDefaults() *TimeseriesDefinition`

NewTimeseriesDefinitionWithDefaults instantiates a new TimeseriesDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeseriesId

`func (o *TimeseriesDefinition) GetTimeseriesId() string`

GetTimeseriesId returns the TimeseriesId field if non-nil, zero value otherwise.

### GetTimeseriesIdOk

`func (o *TimeseriesDefinition) GetTimeseriesIdOk() (*string, bool)`

GetTimeseriesIdOk returns a tuple with the TimeseriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseriesId

`func (o *TimeseriesDefinition) SetTimeseriesId(v string)`

SetTimeseriesId sets TimeseriesId field to given value.

### HasTimeseriesId

`func (o *TimeseriesDefinition) HasTimeseriesId() bool`

HasTimeseriesId returns a boolean if a field has been set.

### GetDisplayName

`func (o *TimeseriesDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TimeseriesDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TimeseriesDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TimeseriesDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDimensions

`func (o *TimeseriesDefinition) GetDimensions() []string`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *TimeseriesDefinition) GetDimensionsOk() (*[]string, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *TimeseriesDefinition) SetDimensions(v []string)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *TimeseriesDefinition) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetAggregationTypes

`func (o *TimeseriesDefinition) GetAggregationTypes() []string`

GetAggregationTypes returns the AggregationTypes field if non-nil, zero value otherwise.

### GetAggregationTypesOk

`func (o *TimeseriesDefinition) GetAggregationTypesOk() (*[]string, bool)`

GetAggregationTypesOk returns a tuple with the AggregationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationTypes

`func (o *TimeseriesDefinition) SetAggregationTypes(v []string)`

SetAggregationTypes sets AggregationTypes field to given value.

### HasAggregationTypes

`func (o *TimeseriesDefinition) HasAggregationTypes() bool`

HasAggregationTypes returns a boolean if a field has been set.

### GetUnit

`func (o *TimeseriesDefinition) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TimeseriesDefinition) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TimeseriesDefinition) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *TimeseriesDefinition) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetFilter

`func (o *TimeseriesDefinition) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TimeseriesDefinition) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TimeseriesDefinition) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TimeseriesDefinition) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetDetailedSource

`func (o *TimeseriesDefinition) GetDetailedSource() string`

GetDetailedSource returns the DetailedSource field if non-nil, zero value otherwise.

### GetDetailedSourceOk

`func (o *TimeseriesDefinition) GetDetailedSourceOk() (*string, bool)`

GetDetailedSourceOk returns a tuple with the DetailedSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedSource

`func (o *TimeseriesDefinition) SetDetailedSource(v string)`

SetDetailedSource sets DetailedSource field to given value.

### HasDetailedSource

`func (o *TimeseriesDefinition) HasDetailedSource() bool`

HasDetailedSource returns a boolean if a field has been set.

### GetPluginId

`func (o *TimeseriesDefinition) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *TimeseriesDefinition) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *TimeseriesDefinition) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.

### HasPluginId

`func (o *TimeseriesDefinition) HasPluginId() bool`

HasPluginId returns a boolean if a field has been set.

### GetTypes

`func (o *TimeseriesDefinition) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *TimeseriesDefinition) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *TimeseriesDefinition) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *TimeseriesDefinition) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


