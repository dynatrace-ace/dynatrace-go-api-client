# AnonymizationProgressResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Progress** | Pointer to **int32** | The progress of the anonymization job, percent.   -1 if the job is waiting for execution. | [optional] 

## Methods

### NewAnonymizationProgressResult

`func NewAnonymizationProgressResult() *AnonymizationProgressResult`

NewAnonymizationProgressResult instantiates a new AnonymizationProgressResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnonymizationProgressResultWithDefaults

`func NewAnonymizationProgressResultWithDefaults() *AnonymizationProgressResult`

NewAnonymizationProgressResultWithDefaults instantiates a new AnonymizationProgressResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgress

`func (o *AnonymizationProgressResult) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *AnonymizationProgressResult) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *AnonymizationProgressResult) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *AnonymizationProgressResult) HasProgress() bool`

HasProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


