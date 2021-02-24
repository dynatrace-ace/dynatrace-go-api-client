# UserSessionQueryTileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomName** | Pointer to **string** | The name of the tile, set by user. | [optional] 
**Query** | Pointer to **string** | A [user session query](https://www.dynatrace.com/support/help/shortlink/usql-info) executed by the tile. | [optional] 
**Type** | Pointer to **string** | The visualization of the tile. | [optional] 
**TimeFrameShift** | Pointer to **string** | The comparison timeframe of the query.    If specified, you additionally get the results of the same query with the specified time shift. | [optional] 
**VisualizationConfig** | Pointer to [**UserSessionQueryTileConfiguration**](UserSessionQueryTileConfiguration.md) |  | [optional] 
**Limit** | Pointer to **int32** | The limit of the results, if not set will use the default value of the system | [optional] 

## Methods

### NewUserSessionQueryTileAllOf

`func NewUserSessionQueryTileAllOf() *UserSessionQueryTileAllOf`

NewUserSessionQueryTileAllOf instantiates a new UserSessionQueryTileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionQueryTileAllOfWithDefaults

`func NewUserSessionQueryTileAllOfWithDefaults() *UserSessionQueryTileAllOf`

NewUserSessionQueryTileAllOfWithDefaults instantiates a new UserSessionQueryTileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomName

`func (o *UserSessionQueryTileAllOf) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *UserSessionQueryTileAllOf) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *UserSessionQueryTileAllOf) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *UserSessionQueryTileAllOf) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### GetQuery

`func (o *UserSessionQueryTileAllOf) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *UserSessionQueryTileAllOf) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *UserSessionQueryTileAllOf) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *UserSessionQueryTileAllOf) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetType

`func (o *UserSessionQueryTileAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSessionQueryTileAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSessionQueryTileAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserSessionQueryTileAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeFrameShift

`func (o *UserSessionQueryTileAllOf) GetTimeFrameShift() string`

GetTimeFrameShift returns the TimeFrameShift field if non-nil, zero value otherwise.

### GetTimeFrameShiftOk

`func (o *UserSessionQueryTileAllOf) GetTimeFrameShiftOk() (*string, bool)`

GetTimeFrameShiftOk returns a tuple with the TimeFrameShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrameShift

`func (o *UserSessionQueryTileAllOf) SetTimeFrameShift(v string)`

SetTimeFrameShift sets TimeFrameShift field to given value.

### HasTimeFrameShift

`func (o *UserSessionQueryTileAllOf) HasTimeFrameShift() bool`

HasTimeFrameShift returns a boolean if a field has been set.

### GetVisualizationConfig

`func (o *UserSessionQueryTileAllOf) GetVisualizationConfig() UserSessionQueryTileConfiguration`

GetVisualizationConfig returns the VisualizationConfig field if non-nil, zero value otherwise.

### GetVisualizationConfigOk

`func (o *UserSessionQueryTileAllOf) GetVisualizationConfigOk() (*UserSessionQueryTileConfiguration, bool)`

GetVisualizationConfigOk returns a tuple with the VisualizationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizationConfig

`func (o *UserSessionQueryTileAllOf) SetVisualizationConfig(v UserSessionQueryTileConfiguration)`

SetVisualizationConfig sets VisualizationConfig field to given value.

### HasVisualizationConfig

`func (o *UserSessionQueryTileAllOf) HasVisualizationConfig() bool`

HasVisualizationConfig returns a boolean if a field has been set.

### GetLimit

`func (o *UserSessionQueryTileAllOf) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *UserSessionQueryTileAllOf) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *UserSessionQueryTileAllOf) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *UserSessionQueryTileAllOf) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


