# DatabaseConnectionFailureDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**ConnectionFailsCount** | Pointer to **int32** | Number of failed database connections during any **timePeriodMinutes** minutes period to trigger an alert. | [optional] 
**TimePeriodMinutes** | Pointer to **int32** | The *X* minutes time period during which the **connectionFailsCount** is evaluated. | [optional] 

## Methods

### NewDatabaseConnectionFailureDetectionConfig

`func NewDatabaseConnectionFailureDetectionConfig(enabled bool, ) *DatabaseConnectionFailureDetectionConfig`

NewDatabaseConnectionFailureDetectionConfig instantiates a new DatabaseConnectionFailureDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseConnectionFailureDetectionConfigWithDefaults

`func NewDatabaseConnectionFailureDetectionConfigWithDefaults() *DatabaseConnectionFailureDetectionConfig`

NewDatabaseConnectionFailureDetectionConfigWithDefaults instantiates a new DatabaseConnectionFailureDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DatabaseConnectionFailureDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DatabaseConnectionFailureDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DatabaseConnectionFailureDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetConnectionFailsCount

`func (o *DatabaseConnectionFailureDetectionConfig) GetConnectionFailsCount() int32`

GetConnectionFailsCount returns the ConnectionFailsCount field if non-nil, zero value otherwise.

### GetConnectionFailsCountOk

`func (o *DatabaseConnectionFailureDetectionConfig) GetConnectionFailsCountOk() (*int32, bool)`

GetConnectionFailsCountOk returns a tuple with the ConnectionFailsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionFailsCount

`func (o *DatabaseConnectionFailureDetectionConfig) SetConnectionFailsCount(v int32)`

SetConnectionFailsCount sets ConnectionFailsCount field to given value.

### HasConnectionFailsCount

`func (o *DatabaseConnectionFailureDetectionConfig) HasConnectionFailsCount() bool`

HasConnectionFailsCount returns a boolean if a field has been set.

### GetTimePeriodMinutes

`func (o *DatabaseConnectionFailureDetectionConfig) GetTimePeriodMinutes() int32`

GetTimePeriodMinutes returns the TimePeriodMinutes field if non-nil, zero value otherwise.

### GetTimePeriodMinutesOk

`func (o *DatabaseConnectionFailureDetectionConfig) GetTimePeriodMinutesOk() (*int32, bool)`

GetTimePeriodMinutesOk returns a tuple with the TimePeriodMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriodMinutes

`func (o *DatabaseConnectionFailureDetectionConfig) SetTimePeriodMinutes(v int32)`

SetTimePeriodMinutes sets TimePeriodMinutes field to given value.

### HasTimePeriodMinutes

`func (o *DatabaseConnectionFailureDetectionConfig) HasTimePeriodMinutes() bool`

HasTimePeriodMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


