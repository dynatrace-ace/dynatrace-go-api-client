# TileFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeframe** | Pointer to **string** | The default timeframe of the tile. | [optional] 
**ManagementZone** | Pointer to [**EntityShortRepresentation**](EntityShortRepresentation.md) |  | [optional] 

## Methods

### NewTileFilter

`func NewTileFilter() *TileFilter`

NewTileFilter instantiates a new TileFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTileFilterWithDefaults

`func NewTileFilterWithDefaults() *TileFilter`

NewTileFilterWithDefaults instantiates a new TileFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeframe

`func (o *TileFilter) GetTimeframe() string`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *TileFilter) GetTimeframeOk() (*string, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *TileFilter) SetTimeframe(v string)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *TileFilter) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### GetManagementZone

`func (o *TileFilter) GetManagementZone() EntityShortRepresentation`

GetManagementZone returns the ManagementZone field if non-nil, zero value otherwise.

### GetManagementZoneOk

`func (o *TileFilter) GetManagementZoneOk() (*EntityShortRepresentation, bool)`

GetManagementZoneOk returns a tuple with the ManagementZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZone

`func (o *TileFilter) SetManagementZone(v EntityShortRepresentation)`

SetManagementZone sets ManagementZone field to given value.

### HasManagementZone

`func (o *TileFilter) HasManagementZone() bool`

HasManagementZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


