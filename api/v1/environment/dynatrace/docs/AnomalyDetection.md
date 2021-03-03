# AnomalyDetection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutageHandling** | Pointer to [**OutageHandlingPolicy**](OutageHandlingPolicy.md) |  | [optional] 
**LoadingTimeThresholds** | Pointer to [**LoadingTimeThresholdsPolicyDto**](LoadingTimeThresholdsPolicyDto.md) |  | [optional] 

## Methods

### NewAnomalyDetection

`func NewAnomalyDetection() *AnomalyDetection`

NewAnomalyDetection instantiates a new AnomalyDetection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyDetectionWithDefaults

`func NewAnomalyDetectionWithDefaults() *AnomalyDetection`

NewAnomalyDetectionWithDefaults instantiates a new AnomalyDetection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutageHandling

`func (o *AnomalyDetection) GetOutageHandling() OutageHandlingPolicy`

GetOutageHandling returns the OutageHandling field if non-nil, zero value otherwise.

### GetOutageHandlingOk

`func (o *AnomalyDetection) GetOutageHandlingOk() (*OutageHandlingPolicy, bool)`

GetOutageHandlingOk returns a tuple with the OutageHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutageHandling

`func (o *AnomalyDetection) SetOutageHandling(v OutageHandlingPolicy)`

SetOutageHandling sets OutageHandling field to given value.

### HasOutageHandling

`func (o *AnomalyDetection) HasOutageHandling() bool`

HasOutageHandling returns a boolean if a field has been set.

### GetLoadingTimeThresholds

`func (o *AnomalyDetection) GetLoadingTimeThresholds() LoadingTimeThresholdsPolicyDto`

GetLoadingTimeThresholds returns the LoadingTimeThresholds field if non-nil, zero value otherwise.

### GetLoadingTimeThresholdsOk

`func (o *AnomalyDetection) GetLoadingTimeThresholdsOk() (*LoadingTimeThresholdsPolicyDto, bool)`

GetLoadingTimeThresholdsOk returns a tuple with the LoadingTimeThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadingTimeThresholds

`func (o *AnomalyDetection) SetLoadingTimeThresholds(v LoadingTimeThresholdsPolicyDto)`

SetLoadingTimeThresholds sets LoadingTimeThresholds field to given value.

### HasLoadingTimeThresholds

`func (o *AnomalyDetection) HasLoadingTimeThresholds() bool`

HasLoadingTimeThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


