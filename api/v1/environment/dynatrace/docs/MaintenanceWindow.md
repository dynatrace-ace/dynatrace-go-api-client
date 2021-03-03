# MaintenanceWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the maintenance window. | [optional] 
**Type** | **string** | The type of the maintenance: planned or unplanned. | 
**Description** | Pointer to **string** | A short description of the maintenance purpose. | [optional] 
**SuppressAlerts** | Pointer to **bool** | Alerting during maintenance is enabled (&#x60;false&#x60;) or disabled (&#x60;true&#x60;). | [optional] 
**SuppressProblems** | Pointer to **bool** | Problem detection during maintenance is enabled (&#x60;false&#x60;) or disabled (&#x60;true&#x60;). | [optional] 
**Scope** | Pointer to [**MaintenanceWindowScope**](MaintenanceWindowScope.md) |  | [optional] 
**Schedule** | [**MaintenanceWindowSchedule**](MaintenanceWindowSchedule.md) |  | 

## Methods

### NewMaintenanceWindow

`func NewMaintenanceWindow(type_ string, schedule MaintenanceWindowSchedule, ) *MaintenanceWindow`

NewMaintenanceWindow instantiates a new MaintenanceWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceWindowWithDefaults

`func NewMaintenanceWindowWithDefaults() *MaintenanceWindow`

NewMaintenanceWindowWithDefaults instantiates a new MaintenanceWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MaintenanceWindow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MaintenanceWindow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MaintenanceWindow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MaintenanceWindow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *MaintenanceWindow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MaintenanceWindow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MaintenanceWindow) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *MaintenanceWindow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MaintenanceWindow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MaintenanceWindow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MaintenanceWindow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSuppressAlerts

`func (o *MaintenanceWindow) GetSuppressAlerts() bool`

GetSuppressAlerts returns the SuppressAlerts field if non-nil, zero value otherwise.

### GetSuppressAlertsOk

`func (o *MaintenanceWindow) GetSuppressAlertsOk() (*bool, bool)`

GetSuppressAlertsOk returns a tuple with the SuppressAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressAlerts

`func (o *MaintenanceWindow) SetSuppressAlerts(v bool)`

SetSuppressAlerts sets SuppressAlerts field to given value.

### HasSuppressAlerts

`func (o *MaintenanceWindow) HasSuppressAlerts() bool`

HasSuppressAlerts returns a boolean if a field has been set.

### GetSuppressProblems

`func (o *MaintenanceWindow) GetSuppressProblems() bool`

GetSuppressProblems returns the SuppressProblems field if non-nil, zero value otherwise.

### GetSuppressProblemsOk

`func (o *MaintenanceWindow) GetSuppressProblemsOk() (*bool, bool)`

GetSuppressProblemsOk returns a tuple with the SuppressProblems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressProblems

`func (o *MaintenanceWindow) SetSuppressProblems(v bool)`

SetSuppressProblems sets SuppressProblems field to given value.

### HasSuppressProblems

`func (o *MaintenanceWindow) HasSuppressProblems() bool`

HasSuppressProblems returns a boolean if a field has been set.

### GetScope

`func (o *MaintenanceWindow) GetScope() MaintenanceWindowScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MaintenanceWindow) GetScopeOk() (*MaintenanceWindowScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MaintenanceWindow) SetScope(v MaintenanceWindowScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MaintenanceWindow) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSchedule

`func (o *MaintenanceWindow) GetSchedule() MaintenanceWindowSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *MaintenanceWindow) GetScheduleOk() (*MaintenanceWindowSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *MaintenanceWindow) SetSchedule(v MaintenanceWindowSchedule)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


