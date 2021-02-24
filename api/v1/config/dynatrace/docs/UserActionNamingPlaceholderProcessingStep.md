# UserActionNamingPlaceholderProcessingStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | An action to be taken by the processing:   * &#x60;SUBSTRING&#x60;: Extracts the string between **patternBefore** and **patternAfter**.  * &#x60;REPLACEMENT&#x60;: Replaces the string between **patternBefore** and **patternAfter** with the specified **replacement**. * &#x60;REPLACE_WITH_PATTERN&#x60;: Replaces the **patternToReplace** with the specified **replacement**.  * &#x60;EXTRACT_BY_REGULAR_EXPRESSION&#x60;: Extracts the part of the string that matches the **regularExpression**.  * &#x60;REPLACE_WITH_REGULAR_EXPRESSION&#x60;: Replaces all occurrences that match **regularExpression** with the specified **replacement**.  * &#x60;REPLACE_IDS&#x60;: Replaces all IDs and UUIDs with the specified **replacement**. | 
**PatternBefore** | Pointer to **string** | The pattern before the required value. It will be removed. | [optional] 
**PatternBeforeSearchType** | Pointer to **string** | The required occurrence of **patternBefore**. | [optional] 
**PatternAfter** | Pointer to **string** | The pattern after the required value. It will be removed. | [optional] 
**PatternAfterSearchType** | Pointer to **string** | The required occurrence of **patternAfter**. | [optional] 
**Replacement** | Pointer to **string** | Replacement for the original value. | [optional] 
**PatternToReplace** | Pointer to **string** | The pattern to be replaced.    Only applicable if the **type** is &#x60;REPLACE_WITH_PATTERN&#x60;. | [optional] 
**RegularExpression** | Pointer to **string** | A regular expression for the string to be extracted or replaced.    Only applicable if the **type** is &#x60;EXTRACT_BY_REGULAR_EXPRESSION&#x60; or &#x60;REPLACE_WITH_REGULAR_EXPRESSION&#x60;. | [optional] 
**FallbackToInput** | Pointer to **bool** | If set to true: Returns the input if **patternBefore** or **patternAfter** cannot be found and the **type** is &#x60;SUBSTRING&#x60;.    Returns the input if **regularExpression** doesn&#39;t match and **type** is &#x60;EXTRACT_BY_REGULAR_EXPRESSION&#x60;.    Otherwise null is returned. | [optional] 

## Methods

### NewUserActionNamingPlaceholderProcessingStep

`func NewUserActionNamingPlaceholderProcessingStep(type_ string, ) *UserActionNamingPlaceholderProcessingStep`

NewUserActionNamingPlaceholderProcessingStep instantiates a new UserActionNamingPlaceholderProcessingStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActionNamingPlaceholderProcessingStepWithDefaults

`func NewUserActionNamingPlaceholderProcessingStepWithDefaults() *UserActionNamingPlaceholderProcessingStep`

NewUserActionNamingPlaceholderProcessingStepWithDefaults instantiates a new UserActionNamingPlaceholderProcessingStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserActionNamingPlaceholderProcessingStep) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserActionNamingPlaceholderProcessingStep) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserActionNamingPlaceholderProcessingStep) SetType(v string)`

SetType sets Type field to given value.


### GetPatternBefore

`func (o *UserActionNamingPlaceholderProcessingStep) GetPatternBefore() string`

GetPatternBefore returns the PatternBefore field if non-nil, zero value otherwise.

### GetPatternBeforeOk

`func (o *UserActionNamingPlaceholderProcessingStep) GetPatternBeforeOk() (*string, bool)`

GetPatternBeforeOk returns a tuple with the PatternBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternBefore

`func (o *UserActionNamingPlaceholderProcessingStep) SetPatternBefore(v string)`

SetPatternBefore sets PatternBefore field to given value.

### HasPatternBefore

`func (o *UserActionNamingPlaceholderProcessingStep) HasPatternBefore() bool`

HasPatternBefore returns a boolean if a field has been set.

### GetPatternBeforeSearchType

`func (o *UserActionNamingPlaceholderProcessingStep) GetPatternBeforeSearchType() string`

GetPatternBeforeSearchType returns the PatternBeforeSearchType field if non-nil, zero value otherwise.

### GetPatternBeforeSearchTypeOk

`func (o *UserActionNamingPlaceholderProcessingStep) GetPatternBeforeSearchTypeOk() (*string, bool)`

GetPatternBeforeSearchTypeOk returns a tuple with the PatternBeforeSearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternBeforeSearchType

`func (o *UserActionNamingPlaceholderProcessingStep) SetPatternBeforeSearchType(v string)`

SetPatternBeforeSearchType sets PatternBeforeSearchType field to given value.

### HasPatternBeforeSearchType

`func (o *UserActionNamingPlaceholderProcessingStep) HasPatternBeforeSearchType() bool`

HasPatternBeforeSearchType returns a boolean if a field has been set.

### GetPatternAfter

`func (o *UserActionNamingPlaceholderProcessingStep) GetPatternAfter() string`

GetPatternAfter returns the PatternAfter field if non-nil, zero value otherwise.

### GetPatternAfterOk

`func (o *UserActionNamingPlaceholderProcessingStep) GetPatternAfterOk() (*string, bool)`

GetPatternAfterOk returns a tuple with the PatternAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternAfter

`func (o *UserActionNamingPlaceholderProcessingStep) SetPatternAfter(v string)`

SetPatternAfter sets PatternAfter field to given value.

### HasPatternAfter

`func (o *UserActionNamingPlaceholderProcessingStep) HasPatternAfter() bool`

HasPatternAfter returns a boolean if a field has been set.

### GetPatternAfterSearchType

`func (o *UserActionNamingPlaceholderProcessingStep) GetPatternAfterSearchType() string`

GetPatternAfterSearchType returns the PatternAfterSearchType field if non-nil, zero value otherwise.

### GetPatternAfterSearchTypeOk

`func (o *UserActionNamingPlaceholderProcessingStep) GetPatternAfterSearchTypeOk() (*string, bool)`

GetPatternAfterSearchTypeOk returns a tuple with the PatternAfterSearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternAfterSearchType

`func (o *UserActionNamingPlaceholderProcessingStep) SetPatternAfterSearchType(v string)`

SetPatternAfterSearchType sets PatternAfterSearchType field to given value.

### HasPatternAfterSearchType

`func (o *UserActionNamingPlaceholderProcessingStep) HasPatternAfterSearchType() bool`

HasPatternAfterSearchType returns a boolean if a field has been set.

### GetReplacement

`func (o *UserActionNamingPlaceholderProcessingStep) GetReplacement() string`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *UserActionNamingPlaceholderProcessingStep) GetReplacementOk() (*string, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *UserActionNamingPlaceholderProcessingStep) SetReplacement(v string)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *UserActionNamingPlaceholderProcessingStep) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### GetPatternToReplace

`func (o *UserActionNamingPlaceholderProcessingStep) GetPatternToReplace() string`

GetPatternToReplace returns the PatternToReplace field if non-nil, zero value otherwise.

### GetPatternToReplaceOk

`func (o *UserActionNamingPlaceholderProcessingStep) GetPatternToReplaceOk() (*string, bool)`

GetPatternToReplaceOk returns a tuple with the PatternToReplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternToReplace

`func (o *UserActionNamingPlaceholderProcessingStep) SetPatternToReplace(v string)`

SetPatternToReplace sets PatternToReplace field to given value.

### HasPatternToReplace

`func (o *UserActionNamingPlaceholderProcessingStep) HasPatternToReplace() bool`

HasPatternToReplace returns a boolean if a field has been set.

### GetRegularExpression

`func (o *UserActionNamingPlaceholderProcessingStep) GetRegularExpression() string`

GetRegularExpression returns the RegularExpression field if non-nil, zero value otherwise.

### GetRegularExpressionOk

`func (o *UserActionNamingPlaceholderProcessingStep) GetRegularExpressionOk() (*string, bool)`

GetRegularExpressionOk returns a tuple with the RegularExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularExpression

`func (o *UserActionNamingPlaceholderProcessingStep) SetRegularExpression(v string)`

SetRegularExpression sets RegularExpression field to given value.

### HasRegularExpression

`func (o *UserActionNamingPlaceholderProcessingStep) HasRegularExpression() bool`

HasRegularExpression returns a boolean if a field has been set.

### GetFallbackToInput

`func (o *UserActionNamingPlaceholderProcessingStep) GetFallbackToInput() bool`

GetFallbackToInput returns the FallbackToInput field if non-nil, zero value otherwise.

### GetFallbackToInputOk

`func (o *UserActionNamingPlaceholderProcessingStep) GetFallbackToInputOk() (*bool, bool)`

GetFallbackToInputOk returns a tuple with the FallbackToInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToInput

`func (o *UserActionNamingPlaceholderProcessingStep) SetFallbackToInput(v bool)`

SetFallbackToInput sets FallbackToInput field to given value.

### HasFallbackToInput

`func (o *UserActionNamingPlaceholderProcessingStep) HasFallbackToInput() bool`

HasFallbackToInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


