# ProblemFeedQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Problems** | Pointer to [**[]Problem**](Problem.md) | The list of problems and their details.   Contains all problems within specified timeframe, open and closed. | [optional] 
**Monitored** | Pointer to [**ProblemFeedQueryResultMonitored**](ProblemFeedQueryResult_monitored.md) |  | [optional] 

## Methods

### NewProblemFeedQueryResult

`func NewProblemFeedQueryResult() *ProblemFeedQueryResult`

NewProblemFeedQueryResult instantiates a new ProblemFeedQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemFeedQueryResultWithDefaults

`func NewProblemFeedQueryResultWithDefaults() *ProblemFeedQueryResult`

NewProblemFeedQueryResultWithDefaults instantiates a new ProblemFeedQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProblems

`func (o *ProblemFeedQueryResult) GetProblems() []Problem`

GetProblems returns the Problems field if non-nil, zero value otherwise.

### GetProblemsOk

`func (o *ProblemFeedQueryResult) GetProblemsOk() (*[]Problem, bool)`

GetProblemsOk returns a tuple with the Problems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblems

`func (o *ProblemFeedQueryResult) SetProblems(v []Problem)`

SetProblems sets Problems field to given value.

### HasProblems

`func (o *ProblemFeedQueryResult) HasProblems() bool`

HasProblems returns a boolean if a field has been set.

### GetMonitored

`func (o *ProblemFeedQueryResult) GetMonitored() ProblemFeedQueryResultMonitored`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *ProblemFeedQueryResult) GetMonitoredOk() (*ProblemFeedQueryResultMonitored, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *ProblemFeedQueryResult) SetMonitored(v ProblemFeedQueryResultMonitored)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *ProblemFeedQueryResult) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


