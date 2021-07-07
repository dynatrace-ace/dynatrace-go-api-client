# MaintenanceWindowEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenanceWindowConfigId** | **string** | The ID of the related maintenance window. | 
**EndTime** | **int64** | The end time of the evidence, in UTC milliseconds. | 

## Methods

### NewMaintenanceWindowEvidence

`func NewMaintenanceWindowEvidence(maintenanceWindowConfigId string, endTime int64, ) *MaintenanceWindowEvidence`

NewMaintenanceWindowEvidence instantiates a new MaintenanceWindowEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceWindowEvidenceWithDefaults

`func NewMaintenanceWindowEvidenceWithDefaults() *MaintenanceWindowEvidence`

NewMaintenanceWindowEvidenceWithDefaults instantiates a new MaintenanceWindowEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenanceWindowConfigId

`func (o *MaintenanceWindowEvidence) GetMaintenanceWindowConfigId() string`

GetMaintenanceWindowConfigId returns the MaintenanceWindowConfigId field if non-nil, zero value otherwise.

### GetMaintenanceWindowConfigIdOk

`func (o *MaintenanceWindowEvidence) GetMaintenanceWindowConfigIdOk() (*string, bool)`

GetMaintenanceWindowConfigIdOk returns a tuple with the MaintenanceWindowConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindowConfigId

`func (o *MaintenanceWindowEvidence) SetMaintenanceWindowConfigId(v string)`

SetMaintenanceWindowConfigId sets MaintenanceWindowConfigId field to given value.


### GetEndTime

`func (o *MaintenanceWindowEvidence) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MaintenanceWindowEvidence) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MaintenanceWindowEvidence) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


