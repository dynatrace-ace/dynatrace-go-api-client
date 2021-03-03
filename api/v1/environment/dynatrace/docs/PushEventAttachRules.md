# PushEventAttachRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityIds** | Pointer to **[]string** | A list of entity IDs to which the event should be attached. | [optional] 
**TagRule** | Pointer to [**[]TagMatchRule**](TagMatchRule.md) | A set of matching rules to dynamically pick up entities based on tags.   Only entities seen within the last **24 hours** are taken into account for tag-based matching rules. | [optional] 

## Methods

### NewPushEventAttachRules

`func NewPushEventAttachRules() *PushEventAttachRules`

NewPushEventAttachRules instantiates a new PushEventAttachRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushEventAttachRulesWithDefaults

`func NewPushEventAttachRulesWithDefaults() *PushEventAttachRules`

NewPushEventAttachRulesWithDefaults instantiates a new PushEventAttachRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityIds

`func (o *PushEventAttachRules) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *PushEventAttachRules) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *PushEventAttachRules) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.

### HasEntityIds

`func (o *PushEventAttachRules) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.

### GetTagRule

`func (o *PushEventAttachRules) GetTagRule() []TagMatchRule`

GetTagRule returns the TagRule field if non-nil, zero value otherwise.

### GetTagRuleOk

`func (o *PushEventAttachRules) GetTagRuleOk() (*[]TagMatchRule, bool)`

GetTagRuleOk returns a tuple with the TagRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagRule

`func (o *PushEventAttachRules) SetTagRule(v []TagMatchRule)`

SetTagRule sets TagRule field to given value.

### HasTagRule

`func (o *PushEventAttachRules) HasTagRule() bool`

HasTagRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


