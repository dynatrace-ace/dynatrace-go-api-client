# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the dashboard. | [optional] 
**DashboardMetadata** | [**DashboardMetadata**](DashboardMetadata.md) |  | 
**Tiles** | [**[]Tile**](Tile.md) | The list of tiles on the dashboard. | 

## Methods

### NewDashboard

`func NewDashboard(dashboardMetadata DashboardMetadata, tiles []Tile, ) *Dashboard`

NewDashboard instantiates a new Dashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWithDefaults

`func NewDashboardWithDefaults() *Dashboard`

NewDashboardWithDefaults instantiates a new Dashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *Dashboard) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Dashboard) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Dashboard) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Dashboard) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *Dashboard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dashboard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dashboard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Dashboard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDashboardMetadata

`func (o *Dashboard) GetDashboardMetadata() DashboardMetadata`

GetDashboardMetadata returns the DashboardMetadata field if non-nil, zero value otherwise.

### GetDashboardMetadataOk

`func (o *Dashboard) GetDashboardMetadataOk() (*DashboardMetadata, bool)`

GetDashboardMetadataOk returns a tuple with the DashboardMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardMetadata

`func (o *Dashboard) SetDashboardMetadata(v DashboardMetadata)`

SetDashboardMetadata sets DashboardMetadata field to given value.


### GetTiles

`func (o *Dashboard) GetTiles() []Tile`

GetTiles returns the Tiles field if non-nil, zero value otherwise.

### GetTilesOk

`func (o *Dashboard) GetTilesOk() (*[]Tile, bool)`

GetTilesOk returns a tuple with the Tiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiles

`func (o *Dashboard) SetTiles(v []Tile)`

SetTiles sets Tiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


