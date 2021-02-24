# DatabaseAnomalyDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**ResponseTimeDegradation** | [**ResponseTimeDegradationDetectionConfig**](ResponseTimeDegradationDetectionConfig.md) |  | 
**LoadDrop** | Pointer to [**LoadDropDetectionConfig**](LoadDropDetectionConfig.md) |  | [optional] 
**LoadSpike** | Pointer to [**LoadSpikeDetectionConfig**](LoadSpikeDetectionConfig.md) |  | [optional] 
**FailureRateIncrease** | [**FailureRateIncreaseDetectionConfig**](FailureRateIncreaseDetectionConfig.md) |  | 
**DatabaseConnectionFailureCount** | [**DatabaseConnectionFailureDetectionConfig**](DatabaseConnectionFailureDetectionConfig.md) |  | 

## Methods

### NewDatabaseAnomalyDetectionConfig

`func NewDatabaseAnomalyDetectionConfig(responseTimeDegradation ResponseTimeDegradationDetectionConfig, failureRateIncrease FailureRateIncreaseDetectionConfig, databaseConnectionFailureCount DatabaseConnectionFailureDetectionConfig, ) *DatabaseAnomalyDetectionConfig`

NewDatabaseAnomalyDetectionConfig instantiates a new DatabaseAnomalyDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseAnomalyDetectionConfigWithDefaults

`func NewDatabaseAnomalyDetectionConfigWithDefaults() *DatabaseAnomalyDetectionConfig`

NewDatabaseAnomalyDetectionConfigWithDefaults instantiates a new DatabaseAnomalyDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *DatabaseAnomalyDetectionConfig) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DatabaseAnomalyDetectionConfig) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DatabaseAnomalyDetectionConfig) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DatabaseAnomalyDetectionConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResponseTimeDegradation

`func (o *DatabaseAnomalyDetectionConfig) GetResponseTimeDegradation() ResponseTimeDegradationDetectionConfig`

GetResponseTimeDegradation returns the ResponseTimeDegradation field if non-nil, zero value otherwise.

### GetResponseTimeDegradationOk

`func (o *DatabaseAnomalyDetectionConfig) GetResponseTimeDegradationOk() (*ResponseTimeDegradationDetectionConfig, bool)`

GetResponseTimeDegradationOk returns a tuple with the ResponseTimeDegradation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeDegradation

`func (o *DatabaseAnomalyDetectionConfig) SetResponseTimeDegradation(v ResponseTimeDegradationDetectionConfig)`

SetResponseTimeDegradation sets ResponseTimeDegradation field to given value.


### GetLoadDrop

`func (o *DatabaseAnomalyDetectionConfig) GetLoadDrop() LoadDropDetectionConfig`

GetLoadDrop returns the LoadDrop field if non-nil, zero value otherwise.

### GetLoadDropOk

`func (o *DatabaseAnomalyDetectionConfig) GetLoadDropOk() (*LoadDropDetectionConfig, bool)`

GetLoadDropOk returns a tuple with the LoadDrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadDrop

`func (o *DatabaseAnomalyDetectionConfig) SetLoadDrop(v LoadDropDetectionConfig)`

SetLoadDrop sets LoadDrop field to given value.

### HasLoadDrop

`func (o *DatabaseAnomalyDetectionConfig) HasLoadDrop() bool`

HasLoadDrop returns a boolean if a field has been set.

### GetLoadSpike

`func (o *DatabaseAnomalyDetectionConfig) GetLoadSpike() LoadSpikeDetectionConfig`

GetLoadSpike returns the LoadSpike field if non-nil, zero value otherwise.

### GetLoadSpikeOk

`func (o *DatabaseAnomalyDetectionConfig) GetLoadSpikeOk() (*LoadSpikeDetectionConfig, bool)`

GetLoadSpikeOk returns a tuple with the LoadSpike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadSpike

`func (o *DatabaseAnomalyDetectionConfig) SetLoadSpike(v LoadSpikeDetectionConfig)`

SetLoadSpike sets LoadSpike field to given value.

### HasLoadSpike

`func (o *DatabaseAnomalyDetectionConfig) HasLoadSpike() bool`

HasLoadSpike returns a boolean if a field has been set.

### GetFailureRateIncrease

`func (o *DatabaseAnomalyDetectionConfig) GetFailureRateIncrease() FailureRateIncreaseDetectionConfig`

GetFailureRateIncrease returns the FailureRateIncrease field if non-nil, zero value otherwise.

### GetFailureRateIncreaseOk

`func (o *DatabaseAnomalyDetectionConfig) GetFailureRateIncreaseOk() (*FailureRateIncreaseDetectionConfig, bool)`

GetFailureRateIncreaseOk returns a tuple with the FailureRateIncrease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureRateIncrease

`func (o *DatabaseAnomalyDetectionConfig) SetFailureRateIncrease(v FailureRateIncreaseDetectionConfig)`

SetFailureRateIncrease sets FailureRateIncrease field to given value.


### GetDatabaseConnectionFailureCount

`func (o *DatabaseAnomalyDetectionConfig) GetDatabaseConnectionFailureCount() DatabaseConnectionFailureDetectionConfig`

GetDatabaseConnectionFailureCount returns the DatabaseConnectionFailureCount field if non-nil, zero value otherwise.

### GetDatabaseConnectionFailureCountOk

`func (o *DatabaseAnomalyDetectionConfig) GetDatabaseConnectionFailureCountOk() (*DatabaseConnectionFailureDetectionConfig, bool)`

GetDatabaseConnectionFailureCountOk returns a tuple with the DatabaseConnectionFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseConnectionFailureCount

`func (o *DatabaseAnomalyDetectionConfig) SetDatabaseConnectionFailureCount(v DatabaseConnectionFailureDetectionConfig)`

SetDatabaseConnectionFailureCount sets DatabaseConnectionFailureCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


