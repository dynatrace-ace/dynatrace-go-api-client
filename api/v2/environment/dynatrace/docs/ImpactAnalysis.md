# ImpactAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Impacts** | Pointer to [**[]Impact**](Impact.md) | A list of all impacts of the problem. | [optional] 

## Methods

### NewImpactAnalysis

`func NewImpactAnalysis() *ImpactAnalysis`

NewImpactAnalysis instantiates a new ImpactAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImpactAnalysisWithDefaults

`func NewImpactAnalysisWithDefaults() *ImpactAnalysis`

NewImpactAnalysisWithDefaults instantiates a new ImpactAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpacts

`func (o *ImpactAnalysis) GetImpacts() []Impact`

GetImpacts returns the Impacts field if non-nil, zero value otherwise.

### GetImpactsOk

`func (o *ImpactAnalysis) GetImpactsOk() (*[]Impact, bool)`

GetImpactsOk returns a tuple with the Impacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpacts

`func (o *ImpactAnalysis) SetImpacts(v []Impact)`

SetImpacts sets Impacts field to given value.

### HasImpacts

`func (o *ImpactAnalysis) HasImpacts() bool`

HasImpacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


