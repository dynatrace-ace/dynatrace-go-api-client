# DetectionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the detection rule. | [optional] 
**Enabled** | **bool** | Rule enabled/disabled. | 
**FileName** | Pointer to **string** | The PHP file containing the class or methods to instrument.   Required for PHP custom service.    Not applicable to Java and .NET. | [optional] 
**FileNameMatcher** | Pointer to **string** | Matcher applying to the file name. Default value is &#x60;ENDS_WITH&#x60; (if applicable). | [optional] 
**ClassName** | Pointer to **string** | The fully qualified class or interface to instrument.   Required for Java and .NET custom services.    Not applicable to PHP. | [optional] 
**Matcher** | Pointer to **string** | Matcher applying to the class name. &#x60;STARTS_WITH&#x60; can only be used if there is at least one annotation defined. Default value is &#x60;EQUALS&#x60;. | [optional] 
**MethodRules** | [**[]MethodRule**](MethodRule.md) | List of methods to instrument. | 
**Annotations** | Pointer to **[]string** | Additional annotations filter of the rule.   Only classes where all listed annotations are available in the class itself or any of its superclasses are instrumented.   Not applicable to PHP. | [optional] 

## Methods

### NewDetectionRule

`func NewDetectionRule(enabled bool, methodRules []MethodRule, ) *DetectionRule`

NewDetectionRule instantiates a new DetectionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectionRuleWithDefaults

`func NewDetectionRuleWithDefaults() *DetectionRule`

NewDetectionRuleWithDefaults instantiates a new DetectionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DetectionRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetectionRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetectionRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DetectionRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnabled

`func (o *DetectionRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DetectionRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DetectionRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFileName

`func (o *DetectionRule) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DetectionRule) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DetectionRule) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DetectionRule) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileNameMatcher

`func (o *DetectionRule) GetFileNameMatcher() string`

GetFileNameMatcher returns the FileNameMatcher field if non-nil, zero value otherwise.

### GetFileNameMatcherOk

`func (o *DetectionRule) GetFileNameMatcherOk() (*string, bool)`

GetFileNameMatcherOk returns a tuple with the FileNameMatcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNameMatcher

`func (o *DetectionRule) SetFileNameMatcher(v string)`

SetFileNameMatcher sets FileNameMatcher field to given value.

### HasFileNameMatcher

`func (o *DetectionRule) HasFileNameMatcher() bool`

HasFileNameMatcher returns a boolean if a field has been set.

### GetClassName

`func (o *DetectionRule) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *DetectionRule) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *DetectionRule) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *DetectionRule) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetMatcher

`func (o *DetectionRule) GetMatcher() string`

GetMatcher returns the Matcher field if non-nil, zero value otherwise.

### GetMatcherOk

`func (o *DetectionRule) GetMatcherOk() (*string, bool)`

GetMatcherOk returns a tuple with the Matcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcher

`func (o *DetectionRule) SetMatcher(v string)`

SetMatcher sets Matcher field to given value.

### HasMatcher

`func (o *DetectionRule) HasMatcher() bool`

HasMatcher returns a boolean if a field has been set.

### GetMethodRules

`func (o *DetectionRule) GetMethodRules() []MethodRule`

GetMethodRules returns the MethodRules field if non-nil, zero value otherwise.

### GetMethodRulesOk

`func (o *DetectionRule) GetMethodRulesOk() (*[]MethodRule, bool)`

GetMethodRulesOk returns a tuple with the MethodRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodRules

`func (o *DetectionRule) SetMethodRules(v []MethodRule)`

SetMethodRules sets MethodRules field to given value.


### GetAnnotations

`func (o *DetectionRule) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *DetectionRule) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *DetectionRule) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *DetectionRule) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


