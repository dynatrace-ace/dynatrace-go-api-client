# JavaScriptInjectionRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The enable or disable rule of the java script injection. | 
**UrlOperator** | **string** | The url operator of the java script injection. | 
**UrlPattern** | Pointer to **string** | The url pattern of the java script injection. | [optional] 
**Rule** | **string** | The url rule of the java script injection. | 
**HtmlPattern** | Pointer to **string** | The html pattern of the java script injection. | [optional] 

## Methods

### NewJavaScriptInjectionRules

`func NewJavaScriptInjectionRules(enabled bool, urlOperator string, rule string, ) *JavaScriptInjectionRules`

NewJavaScriptInjectionRules instantiates a new JavaScriptInjectionRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJavaScriptInjectionRulesWithDefaults

`func NewJavaScriptInjectionRulesWithDefaults() *JavaScriptInjectionRules`

NewJavaScriptInjectionRulesWithDefaults instantiates a new JavaScriptInjectionRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *JavaScriptInjectionRules) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JavaScriptInjectionRules) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JavaScriptInjectionRules) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetUrlOperator

`func (o *JavaScriptInjectionRules) GetUrlOperator() string`

GetUrlOperator returns the UrlOperator field if non-nil, zero value otherwise.

### GetUrlOperatorOk

`func (o *JavaScriptInjectionRules) GetUrlOperatorOk() (*string, bool)`

GetUrlOperatorOk returns a tuple with the UrlOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlOperator

`func (o *JavaScriptInjectionRules) SetUrlOperator(v string)`

SetUrlOperator sets UrlOperator field to given value.


### GetUrlPattern

`func (o *JavaScriptInjectionRules) GetUrlPattern() string`

GetUrlPattern returns the UrlPattern field if non-nil, zero value otherwise.

### GetUrlPatternOk

`func (o *JavaScriptInjectionRules) GetUrlPatternOk() (*string, bool)`

GetUrlPatternOk returns a tuple with the UrlPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPattern

`func (o *JavaScriptInjectionRules) SetUrlPattern(v string)`

SetUrlPattern sets UrlPattern field to given value.

### HasUrlPattern

`func (o *JavaScriptInjectionRules) HasUrlPattern() bool`

HasUrlPattern returns a boolean if a field has been set.

### GetRule

`func (o *JavaScriptInjectionRules) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *JavaScriptInjectionRules) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *JavaScriptInjectionRules) SetRule(v string)`

SetRule sets Rule field to given value.


### GetHtmlPattern

`func (o *JavaScriptInjectionRules) GetHtmlPattern() string`

GetHtmlPattern returns the HtmlPattern field if non-nil, zero value otherwise.

### GetHtmlPatternOk

`func (o *JavaScriptInjectionRules) GetHtmlPatternOk() (*string, bool)`

GetHtmlPatternOk returns a tuple with the HtmlPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlPattern

`func (o *JavaScriptInjectionRules) SetHtmlPattern(v string)`

SetHtmlPattern sets HtmlPattern field to given value.

### HasHtmlPattern

`func (o *JavaScriptInjectionRules) HasHtmlPattern() bool`

HasHtmlPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


