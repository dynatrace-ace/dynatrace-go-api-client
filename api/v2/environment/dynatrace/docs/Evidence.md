# Evidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvidenceType** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;EVENT&#x60; -&gt; EventEvidence  * &#x60;METRIC&#x60; -&gt; MetricEvidence  * &#x60;TRANSACTIONAL&#x60; -&gt; TransactionalEvidence  * &#x60;MAINTENANCE_WINDOW&#x60; -&gt; MaintenanceWindowEvidence  * &#x60;AVAILABILITY_EVIDENCE&#x60; -&gt; AvailabilityEvidence   | 
**DisplayName** | **string** | The display name of the evidence. | 
**Entity** | [**EntityStub**](EntityStub.md) |  | 
**GroupingEntity** | Pointer to [**EntityStub**](EntityStub.md) |  | [optional] 
**RootCauseRelevant** | **bool** | The evidence is (&#x60;true&#x60;) or is not (&#x60;false&#x60;) a part of the root cause. | 
**StartTime** | **int64** | The start time of the evidence, in UTC milliseconds. | 

## Methods

### NewEvidence

`func NewEvidence(evidenceType string, displayName string, entity EntityStub, rootCauseRelevant bool, startTime int64, ) *Evidence`

NewEvidence instantiates a new Evidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidenceWithDefaults

`func NewEvidenceWithDefaults() *Evidence`

NewEvidenceWithDefaults instantiates a new Evidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvidenceType

`func (o *Evidence) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *Evidence) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *Evidence) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.


### GetDisplayName

`func (o *Evidence) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Evidence) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Evidence) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEntity

`func (o *Evidence) GetEntity() EntityStub`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Evidence) GetEntityOk() (*EntityStub, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *Evidence) SetEntity(v EntityStub)`

SetEntity sets Entity field to given value.


### GetGroupingEntity

`func (o *Evidence) GetGroupingEntity() EntityStub`

GetGroupingEntity returns the GroupingEntity field if non-nil, zero value otherwise.

### GetGroupingEntityOk

`func (o *Evidence) GetGroupingEntityOk() (*EntityStub, bool)`

GetGroupingEntityOk returns a tuple with the GroupingEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingEntity

`func (o *Evidence) SetGroupingEntity(v EntityStub)`

SetGroupingEntity sets GroupingEntity field to given value.

### HasGroupingEntity

`func (o *Evidence) HasGroupingEntity() bool`

HasGroupingEntity returns a boolean if a field has been set.

### GetRootCauseRelevant

`func (o *Evidence) GetRootCauseRelevant() bool`

GetRootCauseRelevant returns the RootCauseRelevant field if non-nil, zero value otherwise.

### GetRootCauseRelevantOk

`func (o *Evidence) GetRootCauseRelevantOk() (*bool, bool)`

GetRootCauseRelevantOk returns a tuple with the RootCauseRelevant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseRelevant

`func (o *Evidence) SetRootCauseRelevant(v bool)`

SetRootCauseRelevant sets RootCauseRelevant field to given value.


### GetStartTime

`func (o *Evidence) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Evidence) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Evidence) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


