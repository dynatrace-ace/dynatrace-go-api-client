# LogMonitoringStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentlyUsed** | Pointer to **int64** | Currently used storage [bytes] | [optional] [readonly] 
**MaxLimit** | Pointer to **int64** | Maximum storage limit [bytes] | [optional] 

## Methods

### NewLogMonitoringStorage

`func NewLogMonitoringStorage() *LogMonitoringStorage`

NewLogMonitoringStorage instantiates a new LogMonitoringStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogMonitoringStorageWithDefaults

`func NewLogMonitoringStorageWithDefaults() *LogMonitoringStorage`

NewLogMonitoringStorageWithDefaults instantiates a new LogMonitoringStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentlyUsed

`func (o *LogMonitoringStorage) GetCurrentlyUsed() int64`

GetCurrentlyUsed returns the CurrentlyUsed field if non-nil, zero value otherwise.

### GetCurrentlyUsedOk

`func (o *LogMonitoringStorage) GetCurrentlyUsedOk() (*int64, bool)`

GetCurrentlyUsedOk returns a tuple with the CurrentlyUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyUsed

`func (o *LogMonitoringStorage) SetCurrentlyUsed(v int64)`

SetCurrentlyUsed sets CurrentlyUsed field to given value.

### HasCurrentlyUsed

`func (o *LogMonitoringStorage) HasCurrentlyUsed() bool`

HasCurrentlyUsed returns a boolean if a field has been set.

### GetMaxLimit

`func (o *LogMonitoringStorage) GetMaxLimit() int64`

GetMaxLimit returns the MaxLimit field if non-nil, zero value otherwise.

### GetMaxLimitOk

`func (o *LogMonitoringStorage) GetMaxLimitOk() (*int64, bool)`

GetMaxLimitOk returns a tuple with the MaxLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLimit

`func (o *LogMonitoringStorage) SetMaxLimit(v int64)`

SetMaxLimit sets MaxLimit field to given value.

### HasMaxLimit

`func (o *LogMonitoringStorage) HasMaxLimit() bool`

HasMaxLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


