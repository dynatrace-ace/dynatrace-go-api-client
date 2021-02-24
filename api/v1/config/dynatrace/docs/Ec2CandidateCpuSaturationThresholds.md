# Ec2CandidateCpuSaturationThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUsagePercentage** | **int32** | Alert if CPU usage is higher than *X*% in 3 out of 5 samples. | 

## Methods

### NewEc2CandidateCpuSaturationThresholds

`func NewEc2CandidateCpuSaturationThresholds(cpuUsagePercentage int32, ) *Ec2CandidateCpuSaturationThresholds`

NewEc2CandidateCpuSaturationThresholds instantiates a new Ec2CandidateCpuSaturationThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEc2CandidateCpuSaturationThresholdsWithDefaults

`func NewEc2CandidateCpuSaturationThresholdsWithDefaults() *Ec2CandidateCpuSaturationThresholds`

NewEc2CandidateCpuSaturationThresholdsWithDefaults instantiates a new Ec2CandidateCpuSaturationThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUsagePercentage

`func (o *Ec2CandidateCpuSaturationThresholds) GetCpuUsagePercentage() int32`

GetCpuUsagePercentage returns the CpuUsagePercentage field if non-nil, zero value otherwise.

### GetCpuUsagePercentageOk

`func (o *Ec2CandidateCpuSaturationThresholds) GetCpuUsagePercentageOk() (*int32, bool)`

GetCpuUsagePercentageOk returns a tuple with the CpuUsagePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsagePercentage

`func (o *Ec2CandidateCpuSaturationThresholds) SetCpuUsagePercentage(v int32)`

SetCpuUsagePercentage sets CpuUsagePercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


