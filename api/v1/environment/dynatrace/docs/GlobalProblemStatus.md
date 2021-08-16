# GlobalProblemStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalOpenProblemsCount** | Pointer to **int32** | The total number of open problems in your environment. | [optional] 
**OpenProblemCounts** | Pointer to [**GlobalProblemStatusOpenProblemCounts**](GlobalProblemStatusOpenProblemCounts.md) |  | [optional] 

## Methods

### NewGlobalProblemStatus

`func NewGlobalProblemStatus() *GlobalProblemStatus`

NewGlobalProblemStatus instantiates a new GlobalProblemStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalProblemStatusWithDefaults

`func NewGlobalProblemStatusWithDefaults() *GlobalProblemStatus`

NewGlobalProblemStatusWithDefaults instantiates a new GlobalProblemStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalOpenProblemsCount

`func (o *GlobalProblemStatus) GetTotalOpenProblemsCount() int32`

GetTotalOpenProblemsCount returns the TotalOpenProblemsCount field if non-nil, zero value otherwise.

### GetTotalOpenProblemsCountOk

`func (o *GlobalProblemStatus) GetTotalOpenProblemsCountOk() (*int32, bool)`

GetTotalOpenProblemsCountOk returns a tuple with the TotalOpenProblemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenProblemsCount

`func (o *GlobalProblemStatus) SetTotalOpenProblemsCount(v int32)`

SetTotalOpenProblemsCount sets TotalOpenProblemsCount field to given value.

### HasTotalOpenProblemsCount

`func (o *GlobalProblemStatus) HasTotalOpenProblemsCount() bool`

HasTotalOpenProblemsCount returns a boolean if a field has been set.

### GetOpenProblemCounts

`func (o *GlobalProblemStatus) GetOpenProblemCounts() GlobalProblemStatusOpenProblemCounts`

GetOpenProblemCounts returns the OpenProblemCounts field if non-nil, zero value otherwise.

### GetOpenProblemCountsOk

`func (o *GlobalProblemStatus) GetOpenProblemCountsOk() (*GlobalProblemStatusOpenProblemCounts, bool)`

GetOpenProblemCountsOk returns a tuple with the OpenProblemCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenProblemCounts

`func (o *GlobalProblemStatus) SetOpenProblemCounts(v GlobalProblemStatusOpenProblemCounts)`

SetOpenProblemCounts sets OpenProblemCounts field to given value.

### HasOpenProblemCounts

`func (o *GlobalProblemStatus) HasOpenProblemCounts() bool`

HasOpenProblemCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


