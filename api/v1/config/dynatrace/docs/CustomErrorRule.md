# CustomErrorRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyPattern** | **string** | The key of the error to look for. | 
**KeyMatcher** | **string** | The matching operation for the **keyPattern**. | 
**ValuePattern** | Pointer to **string** | The value of the error to look for. | [optional] 
**ValueMatcher** | Pointer to **string** | The matching operation for the **valuePattern**. | [optional] 
**Capture** | **bool** | Capture (&#x60;true&#x60;) or ignore (&#x60;false&#x60;) the error. | 
**ImpactApdex** | **bool** | Include (&#x60;true&#x60;) or exclude (&#x60;false&#x60;) the error in Apdex calculation. | 
**CustomAlerting** | **bool** | Include (&#x60;true&#x60;) or exclude (&#x60;false&#x60;) the error in Davis AI [problem detection and analysis](https://www.dynatrace.com/support/help/shortlink/problems-hub). | 

## Methods

### NewCustomErrorRule

`func NewCustomErrorRule(keyPattern string, keyMatcher string, capture bool, impactApdex bool, customAlerting bool, ) *CustomErrorRule`

NewCustomErrorRule instantiates a new CustomErrorRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomErrorRuleWithDefaults

`func NewCustomErrorRuleWithDefaults() *CustomErrorRule`

NewCustomErrorRuleWithDefaults instantiates a new CustomErrorRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyPattern

`func (o *CustomErrorRule) GetKeyPattern() string`

GetKeyPattern returns the KeyPattern field if non-nil, zero value otherwise.

### GetKeyPatternOk

`func (o *CustomErrorRule) GetKeyPatternOk() (*string, bool)`

GetKeyPatternOk returns a tuple with the KeyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPattern

`func (o *CustomErrorRule) SetKeyPattern(v string)`

SetKeyPattern sets KeyPattern field to given value.


### GetKeyMatcher

`func (o *CustomErrorRule) GetKeyMatcher() string`

GetKeyMatcher returns the KeyMatcher field if non-nil, zero value otherwise.

### GetKeyMatcherOk

`func (o *CustomErrorRule) GetKeyMatcherOk() (*string, bool)`

GetKeyMatcherOk returns a tuple with the KeyMatcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyMatcher

`func (o *CustomErrorRule) SetKeyMatcher(v string)`

SetKeyMatcher sets KeyMatcher field to given value.


### GetValuePattern

`func (o *CustomErrorRule) GetValuePattern() string`

GetValuePattern returns the ValuePattern field if non-nil, zero value otherwise.

### GetValuePatternOk

`func (o *CustomErrorRule) GetValuePatternOk() (*string, bool)`

GetValuePatternOk returns a tuple with the ValuePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuePattern

`func (o *CustomErrorRule) SetValuePattern(v string)`

SetValuePattern sets ValuePattern field to given value.

### HasValuePattern

`func (o *CustomErrorRule) HasValuePattern() bool`

HasValuePattern returns a boolean if a field has been set.

### GetValueMatcher

`func (o *CustomErrorRule) GetValueMatcher() string`

GetValueMatcher returns the ValueMatcher field if non-nil, zero value otherwise.

### GetValueMatcherOk

`func (o *CustomErrorRule) GetValueMatcherOk() (*string, bool)`

GetValueMatcherOk returns a tuple with the ValueMatcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMatcher

`func (o *CustomErrorRule) SetValueMatcher(v string)`

SetValueMatcher sets ValueMatcher field to given value.

### HasValueMatcher

`func (o *CustomErrorRule) HasValueMatcher() bool`

HasValueMatcher returns a boolean if a field has been set.

### GetCapture

`func (o *CustomErrorRule) GetCapture() bool`

GetCapture returns the Capture field if non-nil, zero value otherwise.

### GetCaptureOk

`func (o *CustomErrorRule) GetCaptureOk() (*bool, bool)`

GetCaptureOk returns a tuple with the Capture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapture

`func (o *CustomErrorRule) SetCapture(v bool)`

SetCapture sets Capture field to given value.


### GetImpactApdex

`func (o *CustomErrorRule) GetImpactApdex() bool`

GetImpactApdex returns the ImpactApdex field if non-nil, zero value otherwise.

### GetImpactApdexOk

`func (o *CustomErrorRule) GetImpactApdexOk() (*bool, bool)`

GetImpactApdexOk returns a tuple with the ImpactApdex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactApdex

`func (o *CustomErrorRule) SetImpactApdex(v bool)`

SetImpactApdex sets ImpactApdex field to given value.


### GetCustomAlerting

`func (o *CustomErrorRule) GetCustomAlerting() bool`

GetCustomAlerting returns the CustomAlerting field if non-nil, zero value otherwise.

### GetCustomAlertingOk

`func (o *CustomErrorRule) GetCustomAlertingOk() (*bool, bool)`

GetCustomAlertingOk returns a tuple with the CustomAlerting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAlerting

`func (o *CustomErrorRule) SetCustomAlerting(v bool)`

SetCustomAlerting sets CustomAlerting field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


