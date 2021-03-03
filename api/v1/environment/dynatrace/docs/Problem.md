# Problem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the problem. | [optional] 
**StartTime** | Pointer to **int64** | The start timestamp of the problem, in UTC milliseconds. | [optional] 
**EndTime** | Pointer to **int64** | The end timestamp of the problem, in UTC milliseconds.    Has the value &#x60;-1&#x60; if the problem is still open. | [optional] 
**DisplayName** | Pointer to **string** | The name of the problem, displayed in the UI. | [optional] 
**ImpactLevel** | Pointer to **string** | The impact level of the problem. It shows what is affected by the problem: infrastructure, service, or application. | [optional] 
**Status** | Pointer to **string** | The status of the problem. | [optional] 
**SeverityLevel** | Pointer to **string** | The severity of the problem. | [optional] 
**CommentCount** | Pointer to **int32** | The number of comments to the problem. | [optional] 
**TagsOfAffectedEntities** | Pointer to [**[]TagInfo**](TagInfo.md) | Tags of entities affected by the problem. | [optional] 
**RankedEvents** | Pointer to [**[]Event**](Event.md) | The list of events related to the problem. | [optional] 
**RankedImpacts** | Pointer to [**[]EventRestImpact**](EventRestImpact.md) | Provides impact information of the events in an aggregated form. For a more detailed impact analysis, see &#x60;rankedEvents&#x60;. | [optional] 
**AffectedCounts** | Pointer to [**ProblemAffectedCounts**](Problem_affectedCounts.md) |  | [optional] 
**RecoveredCounts** | Pointer to [**ProblemRecoveredCounts**](Problem_recoveredCounts.md) |  | [optional] 
**HasRootCause** | Pointer to **bool** | Indicates whether Dynatrace has found at least one possible root cause for the problem. | [optional] 

## Methods

### NewProblem

`func NewProblem() *Problem`

NewProblem instantiates a new Problem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemWithDefaults

`func NewProblemWithDefaults() *Problem`

NewProblemWithDefaults instantiates a new Problem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Problem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Problem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Problem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Problem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartTime

`func (o *Problem) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Problem) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Problem) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Problem) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *Problem) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Problem) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Problem) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Problem) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetDisplayName

`func (o *Problem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Problem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Problem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Problem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetImpactLevel

`func (o *Problem) GetImpactLevel() string`

GetImpactLevel returns the ImpactLevel field if non-nil, zero value otherwise.

### GetImpactLevelOk

`func (o *Problem) GetImpactLevelOk() (*string, bool)`

GetImpactLevelOk returns a tuple with the ImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactLevel

`func (o *Problem) SetImpactLevel(v string)`

SetImpactLevel sets ImpactLevel field to given value.

### HasImpactLevel

`func (o *Problem) HasImpactLevel() bool`

HasImpactLevel returns a boolean if a field has been set.

### GetStatus

`func (o *Problem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Problem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Problem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Problem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSeverityLevel

`func (o *Problem) GetSeverityLevel() string`

GetSeverityLevel returns the SeverityLevel field if non-nil, zero value otherwise.

### GetSeverityLevelOk

`func (o *Problem) GetSeverityLevelOk() (*string, bool)`

GetSeverityLevelOk returns a tuple with the SeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevel

`func (o *Problem) SetSeverityLevel(v string)`

SetSeverityLevel sets SeverityLevel field to given value.

### HasSeverityLevel

`func (o *Problem) HasSeverityLevel() bool`

HasSeverityLevel returns a boolean if a field has been set.

### GetCommentCount

`func (o *Problem) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *Problem) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *Problem) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *Problem) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.

### GetTagsOfAffectedEntities

`func (o *Problem) GetTagsOfAffectedEntities() []TagInfo`

GetTagsOfAffectedEntities returns the TagsOfAffectedEntities field if non-nil, zero value otherwise.

### GetTagsOfAffectedEntitiesOk

`func (o *Problem) GetTagsOfAffectedEntitiesOk() (*[]TagInfo, bool)`

GetTagsOfAffectedEntitiesOk returns a tuple with the TagsOfAffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsOfAffectedEntities

`func (o *Problem) SetTagsOfAffectedEntities(v []TagInfo)`

SetTagsOfAffectedEntities sets TagsOfAffectedEntities field to given value.

### HasTagsOfAffectedEntities

`func (o *Problem) HasTagsOfAffectedEntities() bool`

HasTagsOfAffectedEntities returns a boolean if a field has been set.

### GetRankedEvents

`func (o *Problem) GetRankedEvents() []Event`

GetRankedEvents returns the RankedEvents field if non-nil, zero value otherwise.

### GetRankedEventsOk

`func (o *Problem) GetRankedEventsOk() (*[]Event, bool)`

GetRankedEventsOk returns a tuple with the RankedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankedEvents

`func (o *Problem) SetRankedEvents(v []Event)`

SetRankedEvents sets RankedEvents field to given value.

### HasRankedEvents

`func (o *Problem) HasRankedEvents() bool`

HasRankedEvents returns a boolean if a field has been set.

### GetRankedImpacts

`func (o *Problem) GetRankedImpacts() []EventRestImpact`

GetRankedImpacts returns the RankedImpacts field if non-nil, zero value otherwise.

### GetRankedImpactsOk

`func (o *Problem) GetRankedImpactsOk() (*[]EventRestImpact, bool)`

GetRankedImpactsOk returns a tuple with the RankedImpacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankedImpacts

`func (o *Problem) SetRankedImpacts(v []EventRestImpact)`

SetRankedImpacts sets RankedImpacts field to given value.

### HasRankedImpacts

`func (o *Problem) HasRankedImpacts() bool`

HasRankedImpacts returns a boolean if a field has been set.

### GetAffectedCounts

`func (o *Problem) GetAffectedCounts() ProblemAffectedCounts`

GetAffectedCounts returns the AffectedCounts field if non-nil, zero value otherwise.

### GetAffectedCountsOk

`func (o *Problem) GetAffectedCountsOk() (*ProblemAffectedCounts, bool)`

GetAffectedCountsOk returns a tuple with the AffectedCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedCounts

`func (o *Problem) SetAffectedCounts(v ProblemAffectedCounts)`

SetAffectedCounts sets AffectedCounts field to given value.

### HasAffectedCounts

`func (o *Problem) HasAffectedCounts() bool`

HasAffectedCounts returns a boolean if a field has been set.

### GetRecoveredCounts

`func (o *Problem) GetRecoveredCounts() ProblemRecoveredCounts`

GetRecoveredCounts returns the RecoveredCounts field if non-nil, zero value otherwise.

### GetRecoveredCountsOk

`func (o *Problem) GetRecoveredCountsOk() (*ProblemRecoveredCounts, bool)`

GetRecoveredCountsOk returns a tuple with the RecoveredCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveredCounts

`func (o *Problem) SetRecoveredCounts(v ProblemRecoveredCounts)`

SetRecoveredCounts sets RecoveredCounts field to given value.

### HasRecoveredCounts

`func (o *Problem) HasRecoveredCounts() bool`

HasRecoveredCounts returns a boolean if a field has been set.

### GetHasRootCause

`func (o *Problem) GetHasRootCause() bool`

GetHasRootCause returns the HasRootCause field if non-nil, zero value otherwise.

### GetHasRootCauseOk

`func (o *Problem) GetHasRootCauseOk() (*bool, bool)`

GetHasRootCauseOk returns a tuple with the HasRootCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRootCause

`func (o *Problem) SetHasRootCause(v bool)`

SetHasRootCause sets HasRootCause field to given value.

### HasHasRootCause

`func (o *Problem) HasHasRootCause() bool`

HasHasRootCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


