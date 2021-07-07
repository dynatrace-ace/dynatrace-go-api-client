# MaintenanceWindowEvidenceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenanceWindowConfigId** | Pointer to **string** | The ID of the related maintenance window. | [optional] 
**EndTime** | Pointer to **int64** | The end time of the evidence, in UTC milliseconds. | [optional] 

## Methods

### NewMaintenanceWindowEvidenceAllOf

`func NewMaintenanceWindowEvidenceAllOf() *MaintenanceWindowEvidenceAllOf`

NewMaintenanceWindowEvidenceAllOf instantiates a new MaintenanceWindowEvidenceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceWindowEvidenceAllOfWithDefaults

`func NewMaintenanceWindowEvidenceAllOfWithDefaults() *MaintenanceWindowEvidenceAllOf`

NewMaintenanceWindowEvidenceAllOfWithDefaults instantiates a new MaintenanceWindowEvidenceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenanceWindowConfigId

`func (o *MaintenanceWindowEvidenceAllOf) GetMaintenanceWindowConfigId() string`

GetMaintenanceWindowConfigId returns the MaintenanceWindowConfigId field if non-nil, zero value otherwise.

### GetMaintenanceWindowConfigIdOk

`func (o *MaintenanceWindowEvidenceAllOf) GetMaintenanceWindowConfigIdOk() (*string, bool)`

GetMaintenanceWindowConfigIdOk returns a tuple with the MaintenanceWindowConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindowConfigId

`func (o *MaintenanceWindowEvidenceAllOf) SetMaintenanceWindowConfigId(v string)`

SetMaintenanceWindowConfigId sets MaintenanceWindowConfigId field to given value.

### HasMaintenanceWindowConfigId

`func (o *MaintenanceWindowEvidenceAllOf) HasMaintenanceWindowConfigId() bool`

HasMaintenanceWindowConfigId returns a boolean if a field has been set.

### GetEndTime

`func (o *MaintenanceWindowEvidenceAllOf) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MaintenanceWindowEvidenceAllOf) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MaintenanceWindowEvidenceAllOf) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MaintenanceWindowEvidenceAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


