# SecurityProblemEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **int64** | The timestamp when the event occurred. | [optional] [readonly] 
**Reason** | Pointer to **string** | The reason of the event creation. | [optional] [readonly] 
**RiskAssessmentSnapshot** | Pointer to [**RiskAssessmentSnapshot**](RiskAssessmentSnapshot.md) |  | [optional] 
**MuteState** | Pointer to [**MuteState**](MuteState.md) |  | [optional] 

## Methods

### NewSecurityProblemEvent

`func NewSecurityProblemEvent() *SecurityProblemEvent`

NewSecurityProblemEvent instantiates a new SecurityProblemEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityProblemEventWithDefaults

`func NewSecurityProblemEventWithDefaults() *SecurityProblemEvent`

NewSecurityProblemEventWithDefaults instantiates a new SecurityProblemEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *SecurityProblemEvent) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SecurityProblemEvent) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SecurityProblemEvent) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SecurityProblemEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetReason

`func (o *SecurityProblemEvent) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SecurityProblemEvent) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SecurityProblemEvent) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SecurityProblemEvent) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRiskAssessmentSnapshot

`func (o *SecurityProblemEvent) GetRiskAssessmentSnapshot() RiskAssessmentSnapshot`

GetRiskAssessmentSnapshot returns the RiskAssessmentSnapshot field if non-nil, zero value otherwise.

### GetRiskAssessmentSnapshotOk

`func (o *SecurityProblemEvent) GetRiskAssessmentSnapshotOk() (*RiskAssessmentSnapshot, bool)`

GetRiskAssessmentSnapshotOk returns a tuple with the RiskAssessmentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAssessmentSnapshot

`func (o *SecurityProblemEvent) SetRiskAssessmentSnapshot(v RiskAssessmentSnapshot)`

SetRiskAssessmentSnapshot sets RiskAssessmentSnapshot field to given value.

### HasRiskAssessmentSnapshot

`func (o *SecurityProblemEvent) HasRiskAssessmentSnapshot() bool`

HasRiskAssessmentSnapshot returns a boolean if a field has been set.

### GetMuteState

`func (o *SecurityProblemEvent) GetMuteState() MuteState`

GetMuteState returns the MuteState field if non-nil, zero value otherwise.

### GetMuteStateOk

`func (o *SecurityProblemEvent) GetMuteStateOk() (*MuteState, bool)`

GetMuteStateOk returns a tuple with the MuteState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteState

`func (o *SecurityProblemEvent) SetMuteState(v MuteState)`

SetMuteState sets MuteState field to given value.

### HasMuteState

`func (o *SecurityProblemEvent) HasMuteState() bool`

HasMuteState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


