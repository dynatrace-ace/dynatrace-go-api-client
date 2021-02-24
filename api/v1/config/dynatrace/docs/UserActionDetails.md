# UserActionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The value to be matched to hit the conversion goal. | [optional] 
**CaseSensitive** | Pointer to **bool** | The match is case-sensitive (&#x60;true&#x60;) or (&#x60;false&#x60;). | [optional] 
**MatchType** | Pointer to **string** | The operator of the match. | [optional] 
**MatchEntity** | Pointer to **string** | The type of the entity to which the rule applies. | [optional] 
**ActionType** | Pointer to **string** | Type of the action to which the rule applies. | [optional] 

## Methods

### NewUserActionDetails

`func NewUserActionDetails() *UserActionDetails`

NewUserActionDetails instantiates a new UserActionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActionDetailsWithDefaults

`func NewUserActionDetailsWithDefaults() *UserActionDetails`

NewUserActionDetailsWithDefaults instantiates a new UserActionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *UserActionDetails) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserActionDetails) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserActionDetails) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UserActionDetails) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *UserActionDetails) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *UserActionDetails) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *UserActionDetails) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *UserActionDetails) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetMatchType

`func (o *UserActionDetails) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *UserActionDetails) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *UserActionDetails) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *UserActionDetails) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.

### GetMatchEntity

`func (o *UserActionDetails) GetMatchEntity() string`

GetMatchEntity returns the MatchEntity field if non-nil, zero value otherwise.

### GetMatchEntityOk

`func (o *UserActionDetails) GetMatchEntityOk() (*string, bool)`

GetMatchEntityOk returns a tuple with the MatchEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchEntity

`func (o *UserActionDetails) SetMatchEntity(v string)`

SetMatchEntity sets MatchEntity field to given value.

### HasMatchEntity

`func (o *UserActionDetails) HasMatchEntity() bool`

HasMatchEntity returns a boolean if a field has been set.

### GetActionType

`func (o *UserActionDetails) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *UserActionDetails) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *UserActionDetails) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *UserActionDetails) HasActionType() bool`

HasActionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


