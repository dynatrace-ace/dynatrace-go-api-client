# AwsAnomalyDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**RdsHighCpuDetection** | [**RdsHighCpuDetectionConfig**](RdsHighCpuDetectionConfig.md) |  | 
**RdsHighWriteReadLatencyDetection** | [**RdsHighWriteReadLatencyDetectionConfig**](RdsHighWriteReadLatencyDetectionConfig.md) |  | 
**RdsLowStorageDetection** | [**RdsLowStorageDetectionConfig**](RdsLowStorageDetectionConfig.md) |  | 
**RdsHighMemoryDetection** | [**RdsHighMemoryDetectionConfig**](RdsHighMemoryDetectionConfig.md) |  | 
**ElbHighConnectionErrorsDetection** | [**ElbHighConnectionErrorsDetectionConfig**](ElbHighConnectionErrorsDetectionConfig.md) |  | 
**RdsRestartsSequenceDetection** | [**RdsRestartsSequenceDetectionConfig**](RdsRestartsSequenceDetectionConfig.md) |  | 
**LambdaHighErrorRateDetection** | [**LambdaHighErrorRateDetectionConfig**](LambdaHighErrorRateDetectionConfig.md) |  | 
**Ec2CandidateCpuSaturationDetection** | Pointer to [**Ec2CandidateCpuSaturationDetectionConfig**](Ec2CandidateCpuSaturationDetectionConfig.md) |  | [optional] 

## Methods

### NewAwsAnomalyDetectionConfig

`func NewAwsAnomalyDetectionConfig(rdsHighCpuDetection RdsHighCpuDetectionConfig, rdsHighWriteReadLatencyDetection RdsHighWriteReadLatencyDetectionConfig, rdsLowStorageDetection RdsLowStorageDetectionConfig, rdsHighMemoryDetection RdsHighMemoryDetectionConfig, elbHighConnectionErrorsDetection ElbHighConnectionErrorsDetectionConfig, rdsRestartsSequenceDetection RdsRestartsSequenceDetectionConfig, lambdaHighErrorRateDetection LambdaHighErrorRateDetectionConfig, ) *AwsAnomalyDetectionConfig`

NewAwsAnomalyDetectionConfig instantiates a new AwsAnomalyDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsAnomalyDetectionConfigWithDefaults

`func NewAwsAnomalyDetectionConfigWithDefaults() *AwsAnomalyDetectionConfig`

NewAwsAnomalyDetectionConfigWithDefaults instantiates a new AwsAnomalyDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *AwsAnomalyDetectionConfig) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsAnomalyDetectionConfig) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsAnomalyDetectionConfig) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AwsAnomalyDetectionConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRdsHighCpuDetection

`func (o *AwsAnomalyDetectionConfig) GetRdsHighCpuDetection() RdsHighCpuDetectionConfig`

GetRdsHighCpuDetection returns the RdsHighCpuDetection field if non-nil, zero value otherwise.

### GetRdsHighCpuDetectionOk

`func (o *AwsAnomalyDetectionConfig) GetRdsHighCpuDetectionOk() (*RdsHighCpuDetectionConfig, bool)`

GetRdsHighCpuDetectionOk returns a tuple with the RdsHighCpuDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsHighCpuDetection

`func (o *AwsAnomalyDetectionConfig) SetRdsHighCpuDetection(v RdsHighCpuDetectionConfig)`

SetRdsHighCpuDetection sets RdsHighCpuDetection field to given value.


### GetRdsHighWriteReadLatencyDetection

`func (o *AwsAnomalyDetectionConfig) GetRdsHighWriteReadLatencyDetection() RdsHighWriteReadLatencyDetectionConfig`

GetRdsHighWriteReadLatencyDetection returns the RdsHighWriteReadLatencyDetection field if non-nil, zero value otherwise.

### GetRdsHighWriteReadLatencyDetectionOk

`func (o *AwsAnomalyDetectionConfig) GetRdsHighWriteReadLatencyDetectionOk() (*RdsHighWriteReadLatencyDetectionConfig, bool)`

GetRdsHighWriteReadLatencyDetectionOk returns a tuple with the RdsHighWriteReadLatencyDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsHighWriteReadLatencyDetection

`func (o *AwsAnomalyDetectionConfig) SetRdsHighWriteReadLatencyDetection(v RdsHighWriteReadLatencyDetectionConfig)`

SetRdsHighWriteReadLatencyDetection sets RdsHighWriteReadLatencyDetection field to given value.


### GetRdsLowStorageDetection

`func (o *AwsAnomalyDetectionConfig) GetRdsLowStorageDetection() RdsLowStorageDetectionConfig`

GetRdsLowStorageDetection returns the RdsLowStorageDetection field if non-nil, zero value otherwise.

### GetRdsLowStorageDetectionOk

`func (o *AwsAnomalyDetectionConfig) GetRdsLowStorageDetectionOk() (*RdsLowStorageDetectionConfig, bool)`

GetRdsLowStorageDetectionOk returns a tuple with the RdsLowStorageDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsLowStorageDetection

`func (o *AwsAnomalyDetectionConfig) SetRdsLowStorageDetection(v RdsLowStorageDetectionConfig)`

SetRdsLowStorageDetection sets RdsLowStorageDetection field to given value.


### GetRdsHighMemoryDetection

`func (o *AwsAnomalyDetectionConfig) GetRdsHighMemoryDetection() RdsHighMemoryDetectionConfig`

GetRdsHighMemoryDetection returns the RdsHighMemoryDetection field if non-nil, zero value otherwise.

### GetRdsHighMemoryDetectionOk

`func (o *AwsAnomalyDetectionConfig) GetRdsHighMemoryDetectionOk() (*RdsHighMemoryDetectionConfig, bool)`

GetRdsHighMemoryDetectionOk returns a tuple with the RdsHighMemoryDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsHighMemoryDetection

`func (o *AwsAnomalyDetectionConfig) SetRdsHighMemoryDetection(v RdsHighMemoryDetectionConfig)`

SetRdsHighMemoryDetection sets RdsHighMemoryDetection field to given value.


### GetElbHighConnectionErrorsDetection

`func (o *AwsAnomalyDetectionConfig) GetElbHighConnectionErrorsDetection() ElbHighConnectionErrorsDetectionConfig`

GetElbHighConnectionErrorsDetection returns the ElbHighConnectionErrorsDetection field if non-nil, zero value otherwise.

### GetElbHighConnectionErrorsDetectionOk

`func (o *AwsAnomalyDetectionConfig) GetElbHighConnectionErrorsDetectionOk() (*ElbHighConnectionErrorsDetectionConfig, bool)`

GetElbHighConnectionErrorsDetectionOk returns a tuple with the ElbHighConnectionErrorsDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElbHighConnectionErrorsDetection

`func (o *AwsAnomalyDetectionConfig) SetElbHighConnectionErrorsDetection(v ElbHighConnectionErrorsDetectionConfig)`

SetElbHighConnectionErrorsDetection sets ElbHighConnectionErrorsDetection field to given value.


### GetRdsRestartsSequenceDetection

`func (o *AwsAnomalyDetectionConfig) GetRdsRestartsSequenceDetection() RdsRestartsSequenceDetectionConfig`

GetRdsRestartsSequenceDetection returns the RdsRestartsSequenceDetection field if non-nil, zero value otherwise.

### GetRdsRestartsSequenceDetectionOk

`func (o *AwsAnomalyDetectionConfig) GetRdsRestartsSequenceDetectionOk() (*RdsRestartsSequenceDetectionConfig, bool)`

GetRdsRestartsSequenceDetectionOk returns a tuple with the RdsRestartsSequenceDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsRestartsSequenceDetection

`func (o *AwsAnomalyDetectionConfig) SetRdsRestartsSequenceDetection(v RdsRestartsSequenceDetectionConfig)`

SetRdsRestartsSequenceDetection sets RdsRestartsSequenceDetection field to given value.


### GetLambdaHighErrorRateDetection

`func (o *AwsAnomalyDetectionConfig) GetLambdaHighErrorRateDetection() LambdaHighErrorRateDetectionConfig`

GetLambdaHighErrorRateDetection returns the LambdaHighErrorRateDetection field if non-nil, zero value otherwise.

### GetLambdaHighErrorRateDetectionOk

`func (o *AwsAnomalyDetectionConfig) GetLambdaHighErrorRateDetectionOk() (*LambdaHighErrorRateDetectionConfig, bool)`

GetLambdaHighErrorRateDetectionOk returns a tuple with the LambdaHighErrorRateDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLambdaHighErrorRateDetection

`func (o *AwsAnomalyDetectionConfig) SetLambdaHighErrorRateDetection(v LambdaHighErrorRateDetectionConfig)`

SetLambdaHighErrorRateDetection sets LambdaHighErrorRateDetection field to given value.


### GetEc2CandidateCpuSaturationDetection

`func (o *AwsAnomalyDetectionConfig) GetEc2CandidateCpuSaturationDetection() Ec2CandidateCpuSaturationDetectionConfig`

GetEc2CandidateCpuSaturationDetection returns the Ec2CandidateCpuSaturationDetection field if non-nil, zero value otherwise.

### GetEc2CandidateCpuSaturationDetectionOk

`func (o *AwsAnomalyDetectionConfig) GetEc2CandidateCpuSaturationDetectionOk() (*Ec2CandidateCpuSaturationDetectionConfig, bool)`

GetEc2CandidateCpuSaturationDetectionOk returns a tuple with the Ec2CandidateCpuSaturationDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEc2CandidateCpuSaturationDetection

`func (o *AwsAnomalyDetectionConfig) SetEc2CandidateCpuSaturationDetection(v Ec2CandidateCpuSaturationDetectionConfig)`

SetEc2CandidateCpuSaturationDetection sets Ec2CandidateCpuSaturationDetection field to given value.

### HasEc2CandidateCpuSaturationDetection

`func (o *AwsAnomalyDetectionConfig) HasEc2CandidateCpuSaturationDetection() bool`

HasEc2CandidateCpuSaturationDetection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


