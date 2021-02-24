# LambdaHighErrorRateThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedInvocationsRate** | **int32** | Alert if failed invocations rate is higher than *X*% in 3 out of 5 samples. | 

## Methods

### NewLambdaHighErrorRateThresholds

`func NewLambdaHighErrorRateThresholds(failedInvocationsRate int32, ) *LambdaHighErrorRateThresholds`

NewLambdaHighErrorRateThresholds instantiates a new LambdaHighErrorRateThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLambdaHighErrorRateThresholdsWithDefaults

`func NewLambdaHighErrorRateThresholdsWithDefaults() *LambdaHighErrorRateThresholds`

NewLambdaHighErrorRateThresholdsWithDefaults instantiates a new LambdaHighErrorRateThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedInvocationsRate

`func (o *LambdaHighErrorRateThresholds) GetFailedInvocationsRate() int32`

GetFailedInvocationsRate returns the FailedInvocationsRate field if non-nil, zero value otherwise.

### GetFailedInvocationsRateOk

`func (o *LambdaHighErrorRateThresholds) GetFailedInvocationsRateOk() (*int32, bool)`

GetFailedInvocationsRateOk returns a tuple with the FailedInvocationsRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedInvocationsRate

`func (o *LambdaHighErrorRateThresholds) SetFailedInvocationsRate(v int32)`

SetFailedInvocationsRate sets FailedInvocationsRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


