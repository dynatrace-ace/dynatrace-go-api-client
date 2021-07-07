# SessionReplayStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetentionReductionPercentage** | Pointer to **string** | Percentage of truncation for new data. | [optional] [readonly] 
**RetentionReductionReason** | Pointer to **string** | Reason of truncation. | [optional] [readonly] 
**CurrentlyUsed** | Pointer to **int64** | Currently used storage [bytes] | [optional] [readonly] 
**MaxLimit** | Pointer to **int64** | Maximum storage limit [bytes] | [optional] 

## Methods

### NewSessionReplayStorage

`func NewSessionReplayStorage() *SessionReplayStorage`

NewSessionReplayStorage instantiates a new SessionReplayStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionReplayStorageWithDefaults

`func NewSessionReplayStorageWithDefaults() *SessionReplayStorage`

NewSessionReplayStorageWithDefaults instantiates a new SessionReplayStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetentionReductionPercentage

`func (o *SessionReplayStorage) GetRetentionReductionPercentage() string`

GetRetentionReductionPercentage returns the RetentionReductionPercentage field if non-nil, zero value otherwise.

### GetRetentionReductionPercentageOk

`func (o *SessionReplayStorage) GetRetentionReductionPercentageOk() (*string, bool)`

GetRetentionReductionPercentageOk returns a tuple with the RetentionReductionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionReductionPercentage

`func (o *SessionReplayStorage) SetRetentionReductionPercentage(v string)`

SetRetentionReductionPercentage sets RetentionReductionPercentage field to given value.

### HasRetentionReductionPercentage

`func (o *SessionReplayStorage) HasRetentionReductionPercentage() bool`

HasRetentionReductionPercentage returns a boolean if a field has been set.

### GetRetentionReductionReason

`func (o *SessionReplayStorage) GetRetentionReductionReason() string`

GetRetentionReductionReason returns the RetentionReductionReason field if non-nil, zero value otherwise.

### GetRetentionReductionReasonOk

`func (o *SessionReplayStorage) GetRetentionReductionReasonOk() (*string, bool)`

GetRetentionReductionReasonOk returns a tuple with the RetentionReductionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionReductionReason

`func (o *SessionReplayStorage) SetRetentionReductionReason(v string)`

SetRetentionReductionReason sets RetentionReductionReason field to given value.

### HasRetentionReductionReason

`func (o *SessionReplayStorage) HasRetentionReductionReason() bool`

HasRetentionReductionReason returns a boolean if a field has been set.

### GetCurrentlyUsed

`func (o *SessionReplayStorage) GetCurrentlyUsed() int64`

GetCurrentlyUsed returns the CurrentlyUsed field if non-nil, zero value otherwise.

### GetCurrentlyUsedOk

`func (o *SessionReplayStorage) GetCurrentlyUsedOk() (*int64, bool)`

GetCurrentlyUsedOk returns a tuple with the CurrentlyUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyUsed

`func (o *SessionReplayStorage) SetCurrentlyUsed(v int64)`

SetCurrentlyUsed sets CurrentlyUsed field to given value.

### HasCurrentlyUsed

`func (o *SessionReplayStorage) HasCurrentlyUsed() bool`

HasCurrentlyUsed returns a boolean if a field has been set.

### GetMaxLimit

`func (o *SessionReplayStorage) GetMaxLimit() int64`

GetMaxLimit returns the MaxLimit field if non-nil, zero value otherwise.

### GetMaxLimitOk

`func (o *SessionReplayStorage) GetMaxLimitOk() (*int64, bool)`

GetMaxLimitOk returns a tuple with the MaxLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLimit

`func (o *SessionReplayStorage) SetMaxLimit(v int64)`

SetMaxLimit sets MaxLimit field to given value.

### HasMaxLimit

`func (o *SessionReplayStorage) HasMaxLimit() bool`

HasMaxLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


