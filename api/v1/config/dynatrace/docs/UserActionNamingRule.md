# UserActionNamingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | **string** | Naming pattern. Use Curly brackets &#x60;{}&#x60; to select placeholders. | 
**Conditions** | Pointer to [**[]UserActionNamingRuleCondition**](UserActionNamingRuleCondition.md) | Defines the conditions when the naming rule should apply. | [optional] 

## Methods

### NewUserActionNamingRule

`func NewUserActionNamingRule(template string, ) *UserActionNamingRule`

NewUserActionNamingRule instantiates a new UserActionNamingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActionNamingRuleWithDefaults

`func NewUserActionNamingRuleWithDefaults() *UserActionNamingRule`

NewUserActionNamingRuleWithDefaults instantiates a new UserActionNamingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *UserActionNamingRule) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *UserActionNamingRule) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *UserActionNamingRule) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetConditions

`func (o *UserActionNamingRule) GetConditions() []UserActionNamingRuleCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *UserActionNamingRule) GetConditionsOk() (*[]UserActionNamingRuleCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *UserActionNamingRule) SetConditions(v []UserActionNamingRuleCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *UserActionNamingRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


