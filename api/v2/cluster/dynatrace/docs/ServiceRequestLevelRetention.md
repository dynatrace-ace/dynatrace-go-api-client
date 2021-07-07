# ServiceRequestLevelRetention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxLimitInDays** | Pointer to **int64** | Maximum retention limit [days] | [optional] 
**CurrentlyUsedInMillis** | Pointer to **int64** | Current data age [milliseconds] | [optional] [readonly] 
**CurrentlyUsedInDays** | Pointer to **int64** | Current data age [days] | [optional] [readonly] 

## Methods

### NewServiceRequestLevelRetention

`func NewServiceRequestLevelRetention() *ServiceRequestLevelRetention`

NewServiceRequestLevelRetention instantiates a new ServiceRequestLevelRetention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRequestLevelRetentionWithDefaults

`func NewServiceRequestLevelRetentionWithDefaults() *ServiceRequestLevelRetention`

NewServiceRequestLevelRetentionWithDefaults instantiates a new ServiceRequestLevelRetention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxLimitInDays

`func (o *ServiceRequestLevelRetention) GetMaxLimitInDays() int64`

GetMaxLimitInDays returns the MaxLimitInDays field if non-nil, zero value otherwise.

### GetMaxLimitInDaysOk

`func (o *ServiceRequestLevelRetention) GetMaxLimitInDaysOk() (*int64, bool)`

GetMaxLimitInDaysOk returns a tuple with the MaxLimitInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLimitInDays

`func (o *ServiceRequestLevelRetention) SetMaxLimitInDays(v int64)`

SetMaxLimitInDays sets MaxLimitInDays field to given value.

### HasMaxLimitInDays

`func (o *ServiceRequestLevelRetention) HasMaxLimitInDays() bool`

HasMaxLimitInDays returns a boolean if a field has been set.

### GetCurrentlyUsedInMillis

`func (o *ServiceRequestLevelRetention) GetCurrentlyUsedInMillis() int64`

GetCurrentlyUsedInMillis returns the CurrentlyUsedInMillis field if non-nil, zero value otherwise.

### GetCurrentlyUsedInMillisOk

`func (o *ServiceRequestLevelRetention) GetCurrentlyUsedInMillisOk() (*int64, bool)`

GetCurrentlyUsedInMillisOk returns a tuple with the CurrentlyUsedInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyUsedInMillis

`func (o *ServiceRequestLevelRetention) SetCurrentlyUsedInMillis(v int64)`

SetCurrentlyUsedInMillis sets CurrentlyUsedInMillis field to given value.

### HasCurrentlyUsedInMillis

`func (o *ServiceRequestLevelRetention) HasCurrentlyUsedInMillis() bool`

HasCurrentlyUsedInMillis returns a boolean if a field has been set.

### GetCurrentlyUsedInDays

`func (o *ServiceRequestLevelRetention) GetCurrentlyUsedInDays() int64`

GetCurrentlyUsedInDays returns the CurrentlyUsedInDays field if non-nil, zero value otherwise.

### GetCurrentlyUsedInDaysOk

`func (o *ServiceRequestLevelRetention) GetCurrentlyUsedInDaysOk() (*int64, bool)`

GetCurrentlyUsedInDaysOk returns a tuple with the CurrentlyUsedInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyUsedInDays

`func (o *ServiceRequestLevelRetention) SetCurrentlyUsedInDays(v int64)`

SetCurrentlyUsedInDays sets CurrentlyUsedInDays field to given value.

### HasCurrentlyUsedInDays

`func (o *ServiceRequestLevelRetention) HasCurrentlyUsedInDays() bool`

HasCurrentlyUsedInDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


