# Problem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImpactLevel** | **string** | The impact level of the problem. It shows what is affected by the problem. | 
**RecentComments** | Pointer to [**CommentsList**](CommentsList.md) |  | [optional] 
**RootCauseEntity** | Pointer to [**EntityStub**](EntityStub.md) |  | [optional] 
**ImpactAnalysis** | Pointer to [**ImpactAnalysis**](ImpactAnalysis.md) |  | [optional] 
**LinkedProblemInfo** | Pointer to [**LinkedProblem**](LinkedProblem.md) |  | [optional] 
**ProblemFilters** | Pointer to [**[]AlertingProfileStub**](AlertingProfileStub.md) | A list of alerting profiles that match the problem. | [optional] 
**EvidenceDetails** | Pointer to [**EvidenceDetails**](EvidenceDetails.md) |  | [optional] 
**ImpactedEntities** | Pointer to [**[]EntityStub**](EntityStub.md) | A list of all entities that are impacted by the problem. | [optional] 
**ManagementZones** | Pointer to [**[]ManagementZone**](ManagementZone.md) | A list of all management zones that the problem belongs to. | [optional] 
**EntityTags** | Pointer to [**[]METag**](METag.md) | A list of all entity tags of the problem. | [optional] 
**AffectedEntities** | Pointer to [**[]EntityStub**](EntityStub.md) | A list of all entities that are affected by the problem. | [optional] 
**DisplayId** | **string** | The display ID of the problem. | 
**ProblemId** | **string** | The ID of the problem. | 
**SeverityLevel** | **string** | The severity of the problem. | 
**Title** | **string** | The name of the problem, displayed in the UI. | 
**Status** | **string** | The status of the problem. | 
**StartTime** | **int64** | The start timestamp of the problem, in UTC milliseconds. | 
**EndTime** | **int64** | The end timestamp of the problem, in UTC milliseconds.    Has &#x60;-1&#x60; value, if the problem is still open. | 

## Methods

### NewProblem

`func NewProblem(impactLevel string, displayId string, problemId string, severityLevel string, title string, status string, startTime int64, endTime int64, ) *Problem`

NewProblem instantiates a new Problem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemWithDefaults

`func NewProblemWithDefaults() *Problem`

NewProblemWithDefaults instantiates a new Problem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetRecentComments

`func (o *Problem) GetRecentComments() CommentsList`

GetRecentComments returns the RecentComments field if non-nil, zero value otherwise.

### GetRecentCommentsOk

`func (o *Problem) GetRecentCommentsOk() (*CommentsList, bool)`

GetRecentCommentsOk returns a tuple with the RecentComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentComments

`func (o *Problem) SetRecentComments(v CommentsList)`

SetRecentComments sets RecentComments field to given value.

### HasRecentComments

`func (o *Problem) HasRecentComments() bool`

HasRecentComments returns a boolean if a field has been set.

### GetRootCauseEntity

`func (o *Problem) GetRootCauseEntity() EntityStub`

GetRootCauseEntity returns the RootCauseEntity field if non-nil, zero value otherwise.

### GetRootCauseEntityOk

`func (o *Problem) GetRootCauseEntityOk() (*EntityStub, bool)`

GetRootCauseEntityOk returns a tuple with the RootCauseEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseEntity

`func (o *Problem) SetRootCauseEntity(v EntityStub)`

SetRootCauseEntity sets RootCauseEntity field to given value.

### HasRootCauseEntity

`func (o *Problem) HasRootCauseEntity() bool`

HasRootCauseEntity returns a boolean if a field has been set.

### GetImpactAnalysis

`func (o *Problem) GetImpactAnalysis() ImpactAnalysis`

GetImpactAnalysis returns the ImpactAnalysis field if non-nil, zero value otherwise.

### GetImpactAnalysisOk

`func (o *Problem) GetImpactAnalysisOk() (*ImpactAnalysis, bool)`

GetImpactAnalysisOk returns a tuple with the ImpactAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactAnalysis

`func (o *Problem) SetImpactAnalysis(v ImpactAnalysis)`

SetImpactAnalysis sets ImpactAnalysis field to given value.

### HasImpactAnalysis

`func (o *Problem) HasImpactAnalysis() bool`

HasImpactAnalysis returns a boolean if a field has been set.

### GetLinkedProblemInfo

`func (o *Problem) GetLinkedProblemInfo() LinkedProblem`

GetLinkedProblemInfo returns the LinkedProblemInfo field if non-nil, zero value otherwise.

### GetLinkedProblemInfoOk

`func (o *Problem) GetLinkedProblemInfoOk() (*LinkedProblem, bool)`

GetLinkedProblemInfoOk returns a tuple with the LinkedProblemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedProblemInfo

`func (o *Problem) SetLinkedProblemInfo(v LinkedProblem)`

SetLinkedProblemInfo sets LinkedProblemInfo field to given value.

### HasLinkedProblemInfo

`func (o *Problem) HasLinkedProblemInfo() bool`

HasLinkedProblemInfo returns a boolean if a field has been set.

### GetProblemFilters

`func (o *Problem) GetProblemFilters() []AlertingProfileStub`

GetProblemFilters returns the ProblemFilters field if non-nil, zero value otherwise.

### GetProblemFiltersOk

`func (o *Problem) GetProblemFiltersOk() (*[]AlertingProfileStub, bool)`

GetProblemFiltersOk returns a tuple with the ProblemFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemFilters

`func (o *Problem) SetProblemFilters(v []AlertingProfileStub)`

SetProblemFilters sets ProblemFilters field to given value.

### HasProblemFilters

`func (o *Problem) HasProblemFilters() bool`

HasProblemFilters returns a boolean if a field has been set.

### GetEvidenceDetails

`func (o *Problem) GetEvidenceDetails() EvidenceDetails`

GetEvidenceDetails returns the EvidenceDetails field if non-nil, zero value otherwise.

### GetEvidenceDetailsOk

`func (o *Problem) GetEvidenceDetailsOk() (*EvidenceDetails, bool)`

GetEvidenceDetailsOk returns a tuple with the EvidenceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceDetails

`func (o *Problem) SetEvidenceDetails(v EvidenceDetails)`

SetEvidenceDetails sets EvidenceDetails field to given value.

### HasEvidenceDetails

`func (o *Problem) HasEvidenceDetails() bool`

HasEvidenceDetails returns a boolean if a field has been set.

### GetImpactedEntities

`func (o *Problem) GetImpactedEntities() []EntityStub`

GetImpactedEntities returns the ImpactedEntities field if non-nil, zero value otherwise.

### GetImpactedEntitiesOk

`func (o *Problem) GetImpactedEntitiesOk() (*[]EntityStub, bool)`

GetImpactedEntitiesOk returns a tuple with the ImpactedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactedEntities

`func (o *Problem) SetImpactedEntities(v []EntityStub)`

SetImpactedEntities sets ImpactedEntities field to given value.

### HasImpactedEntities

`func (o *Problem) HasImpactedEntities() bool`

HasImpactedEntities returns a boolean if a field has been set.

### GetManagementZones

`func (o *Problem) GetManagementZones() []ManagementZone`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *Problem) GetManagementZonesOk() (*[]ManagementZone, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *Problem) SetManagementZones(v []ManagementZone)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *Problem) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetEntityTags

`func (o *Problem) GetEntityTags() []METag`

GetEntityTags returns the EntityTags field if non-nil, zero value otherwise.

### GetEntityTagsOk

`func (o *Problem) GetEntityTagsOk() (*[]METag, bool)`

GetEntityTagsOk returns a tuple with the EntityTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTags

`func (o *Problem) SetEntityTags(v []METag)`

SetEntityTags sets EntityTags field to given value.

### HasEntityTags

`func (o *Problem) HasEntityTags() bool`

HasEntityTags returns a boolean if a field has been set.

### GetAffectedEntities

`func (o *Problem) GetAffectedEntities() []EntityStub`

GetAffectedEntities returns the AffectedEntities field if non-nil, zero value otherwise.

### GetAffectedEntitiesOk

`func (o *Problem) GetAffectedEntitiesOk() (*[]EntityStub, bool)`

GetAffectedEntitiesOk returns a tuple with the AffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEntities

`func (o *Problem) SetAffectedEntities(v []EntityStub)`

SetAffectedEntities sets AffectedEntities field to given value.

### HasAffectedEntities

`func (o *Problem) HasAffectedEntities() bool`

HasAffectedEntities returns a boolean if a field has been set.

### GetDisplayId

`func (o *Problem) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *Problem) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *Problem) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.


### GetProblemId

`func (o *Problem) GetProblemId() string`

GetProblemId returns the ProblemId field if non-nil, zero value otherwise.

### GetProblemIdOk

`func (o *Problem) GetProblemIdOk() (*string, bool)`

GetProblemIdOk returns a tuple with the ProblemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemId

`func (o *Problem) SetProblemId(v string)`

SetProblemId sets ProblemId field to given value.


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


### GetTitle

`func (o *Problem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Problem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Problem) SetTitle(v string)`

SetTitle sets Title field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


