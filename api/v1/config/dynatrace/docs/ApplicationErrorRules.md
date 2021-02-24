# ApplicationErrorRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoreJavaScriptErrorsInApdexCalculation** | **bool** | Exclude (&#x60;true&#x60;) or include (&#x60;false&#x60;) JavaScript errors in Apdex calculation. | 
**IgnoreHttpErrorsInApdexCalculation** | **bool** | Exclude (&#x60;true&#x60;) or include (&#x60;false&#x60;) HTTP errors listed in **httpErrorRules** in Apdex calculation. | 
**IgnoreCustomErrorsInApdexCalculation** | **bool** | Exclude (&#x60;true&#x60;) or include (&#x60;false&#x60;) custom errors listed in **customErrorRules** in Apdex calculation. | 
**HttpErrorRules** | [**[]HttpErrorRule**](HttpErrorRule.md) | An ordered list of HTTP errors.   Rules are evaluated from top to bottom; the first matching rule applies. | 
**CustomErrorRules** | [**[]CustomErrorRule**](CustomErrorRule.md) | An ordered list of custom errors.   Rules are evaluated from top to bottom; the first matching rule applies. | 

## Methods

### NewApplicationErrorRules

`func NewApplicationErrorRules(ignoreJavaScriptErrorsInApdexCalculation bool, ignoreHttpErrorsInApdexCalculation bool, ignoreCustomErrorsInApdexCalculation bool, httpErrorRules []HttpErrorRule, customErrorRules []CustomErrorRule, ) *ApplicationErrorRules`

NewApplicationErrorRules instantiates a new ApplicationErrorRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationErrorRulesWithDefaults

`func NewApplicationErrorRulesWithDefaults() *ApplicationErrorRules`

NewApplicationErrorRulesWithDefaults instantiates a new ApplicationErrorRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoreJavaScriptErrorsInApdexCalculation

`func (o *ApplicationErrorRules) GetIgnoreJavaScriptErrorsInApdexCalculation() bool`

GetIgnoreJavaScriptErrorsInApdexCalculation returns the IgnoreJavaScriptErrorsInApdexCalculation field if non-nil, zero value otherwise.

### GetIgnoreJavaScriptErrorsInApdexCalculationOk

`func (o *ApplicationErrorRules) GetIgnoreJavaScriptErrorsInApdexCalculationOk() (*bool, bool)`

GetIgnoreJavaScriptErrorsInApdexCalculationOk returns a tuple with the IgnoreJavaScriptErrorsInApdexCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreJavaScriptErrorsInApdexCalculation

`func (o *ApplicationErrorRules) SetIgnoreJavaScriptErrorsInApdexCalculation(v bool)`

SetIgnoreJavaScriptErrorsInApdexCalculation sets IgnoreJavaScriptErrorsInApdexCalculation field to given value.


### GetIgnoreHttpErrorsInApdexCalculation

`func (o *ApplicationErrorRules) GetIgnoreHttpErrorsInApdexCalculation() bool`

GetIgnoreHttpErrorsInApdexCalculation returns the IgnoreHttpErrorsInApdexCalculation field if non-nil, zero value otherwise.

### GetIgnoreHttpErrorsInApdexCalculationOk

`func (o *ApplicationErrorRules) GetIgnoreHttpErrorsInApdexCalculationOk() (*bool, bool)`

GetIgnoreHttpErrorsInApdexCalculationOk returns a tuple with the IgnoreHttpErrorsInApdexCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreHttpErrorsInApdexCalculation

`func (o *ApplicationErrorRules) SetIgnoreHttpErrorsInApdexCalculation(v bool)`

SetIgnoreHttpErrorsInApdexCalculation sets IgnoreHttpErrorsInApdexCalculation field to given value.


### GetIgnoreCustomErrorsInApdexCalculation

`func (o *ApplicationErrorRules) GetIgnoreCustomErrorsInApdexCalculation() bool`

GetIgnoreCustomErrorsInApdexCalculation returns the IgnoreCustomErrorsInApdexCalculation field if non-nil, zero value otherwise.

### GetIgnoreCustomErrorsInApdexCalculationOk

`func (o *ApplicationErrorRules) GetIgnoreCustomErrorsInApdexCalculationOk() (*bool, bool)`

GetIgnoreCustomErrorsInApdexCalculationOk returns a tuple with the IgnoreCustomErrorsInApdexCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCustomErrorsInApdexCalculation

`func (o *ApplicationErrorRules) SetIgnoreCustomErrorsInApdexCalculation(v bool)`

SetIgnoreCustomErrorsInApdexCalculation sets IgnoreCustomErrorsInApdexCalculation field to given value.


### GetHttpErrorRules

`func (o *ApplicationErrorRules) GetHttpErrorRules() []HttpErrorRule`

GetHttpErrorRules returns the HttpErrorRules field if non-nil, zero value otherwise.

### GetHttpErrorRulesOk

`func (o *ApplicationErrorRules) GetHttpErrorRulesOk() (*[]HttpErrorRule, bool)`

GetHttpErrorRulesOk returns a tuple with the HttpErrorRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpErrorRules

`func (o *ApplicationErrorRules) SetHttpErrorRules(v []HttpErrorRule)`

SetHttpErrorRules sets HttpErrorRules field to given value.


### GetCustomErrorRules

`func (o *ApplicationErrorRules) GetCustomErrorRules() []CustomErrorRule`

GetCustomErrorRules returns the CustomErrorRules field if non-nil, zero value otherwise.

### GetCustomErrorRulesOk

`func (o *ApplicationErrorRules) GetCustomErrorRulesOk() (*[]CustomErrorRule, bool)`

GetCustomErrorRulesOk returns a tuple with the CustomErrorRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomErrorRules

`func (o *ApplicationErrorRules) SetCustomErrorRules(v []CustomErrorRule)`

SetCustomErrorRules sets CustomErrorRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


