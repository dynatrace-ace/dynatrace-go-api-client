# DashboardFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeframe** | Pointer to **string** | The default timeframe of the dashboard. | [optional] 
**ManagementZone** | Pointer to [**EntityShortRepresentation**](EntityShortRepresentation.md) |  | [optional] 

## Methods

### NewDashboardFilter

`func NewDashboardFilter() *DashboardFilter`

NewDashboardFilter instantiates a new DashboardFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardFilterWithDefaults

`func NewDashboardFilterWithDefaults() *DashboardFilter`

NewDashboardFilterWithDefaults instantiates a new DashboardFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeframe

`func (o *DashboardFilter) GetTimeframe() string`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *DashboardFilter) GetTimeframeOk() (*string, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *DashboardFilter) SetTimeframe(v string)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *DashboardFilter) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### GetManagementZone

`func (o *DashboardFilter) GetManagementZone() EntityShortRepresentation`

GetManagementZone returns the ManagementZone field if non-nil, zero value otherwise.

### GetManagementZoneOk

`func (o *DashboardFilter) GetManagementZoneOk() (*EntityShortRepresentation, bool)`

GetManagementZoneOk returns a tuple with the ManagementZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZone

`func (o *DashboardFilter) SetManagementZone(v EntityShortRepresentation)`

SetManagementZone sets ManagementZone field to given value.

### HasManagementZone

`func (o *DashboardFilter) HasManagementZone() bool`

HasManagementZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


