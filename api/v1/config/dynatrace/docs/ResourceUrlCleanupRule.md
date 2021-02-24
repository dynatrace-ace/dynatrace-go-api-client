# ResourceUrlCleanupRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | The name of the rule. | 
**RegularExpression** | **string** | The pattern (regular expression) to look for. | 
**ReplaceWith** | **string** | The text to replace the found pattern with. | 

## Methods

### NewResourceUrlCleanupRule

`func NewResourceUrlCleanupRule(resourceName string, regularExpression string, replaceWith string, ) *ResourceUrlCleanupRule`

NewResourceUrlCleanupRule instantiates a new ResourceUrlCleanupRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUrlCleanupRuleWithDefaults

`func NewResourceUrlCleanupRuleWithDefaults() *ResourceUrlCleanupRule`

NewResourceUrlCleanupRuleWithDefaults instantiates a new ResourceUrlCleanupRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *ResourceUrlCleanupRule) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceUrlCleanupRule) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ResourceUrlCleanupRule) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetRegularExpression

`func (o *ResourceUrlCleanupRule) GetRegularExpression() string`

GetRegularExpression returns the RegularExpression field if non-nil, zero value otherwise.

### GetRegularExpressionOk

`func (o *ResourceUrlCleanupRule) GetRegularExpressionOk() (*string, bool)`

GetRegularExpressionOk returns a tuple with the RegularExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularExpression

`func (o *ResourceUrlCleanupRule) SetRegularExpression(v string)`

SetRegularExpression sets RegularExpression field to given value.


### GetReplaceWith

`func (o *ResourceUrlCleanupRule) GetReplaceWith() string`

GetReplaceWith returns the ReplaceWith field if non-nil, zero value otherwise.

### GetReplaceWithOk

`func (o *ResourceUrlCleanupRule) GetReplaceWithOk() (*string, bool)`

GetReplaceWithOk returns a tuple with the ReplaceWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceWith

`func (o *ResourceUrlCleanupRule) SetReplaceWith(v string)`

SetReplaceWith sets ReplaceWith field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


