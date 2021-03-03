# SyntheticTestStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the step. | 
**Title** | **string** | The name of the step, displayed in the UI. | 

## Methods

### NewSyntheticTestStep

`func NewSyntheticTestStep(id int64, title string, ) *SyntheticTestStep`

NewSyntheticTestStep instantiates a new SyntheticTestStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticTestStepWithDefaults

`func NewSyntheticTestStepWithDefaults() *SyntheticTestStep`

NewSyntheticTestStepWithDefaults instantiates a new SyntheticTestStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SyntheticTestStep) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyntheticTestStep) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyntheticTestStep) SetId(v int64)`

SetId sets Id field to given value.


### GetTitle

`func (o *SyntheticTestStep) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SyntheticTestStep) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SyntheticTestStep) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


