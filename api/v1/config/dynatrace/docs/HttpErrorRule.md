# HttpErrorRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsiderUnknownErrorCode** | **bool** | If &#x60;true&#x60;, match by errors that have unknown HTTP status code. | 
**ErrorCodes** | Pointer to **string** | The HTTP status code or status code range to match by.   This field is required if **considerUnknownErrorCode** is set to &#x60;false&#x60;. | [optional] 
**FilterByUrl** | **bool** | If &#x60;true&#x60;, filter errors by URL. | 
**Filter** | Pointer to **string** | The matching rule for the URL. | [optional] 
**Url** | Pointer to **string** | The URL to look for. | [optional] 
**Capture** | **bool** | Capture (&#x60;true&#x60;) or ignore (&#x60;false&#x60;) the error. | 
**ImpactApdex** | **bool** | Include (&#x60;true&#x60;) or exclude (&#x60;false&#x60;) the error in Apdex calculation. | 
**ConsiderForAi** | **bool** | Include (&#x60;true&#x60;) or exclude (&#x60;false&#x60;) the error in Davis AI [problem detection and analysis](https://www.dynatrace.com/support/help/shortlink/problems-hub). | 

## Methods

### NewHttpErrorRule

`func NewHttpErrorRule(considerUnknownErrorCode bool, filterByUrl bool, capture bool, impactApdex bool, considerForAi bool, ) *HttpErrorRule`

NewHttpErrorRule instantiates a new HttpErrorRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpErrorRuleWithDefaults

`func NewHttpErrorRuleWithDefaults() *HttpErrorRule`

NewHttpErrorRuleWithDefaults instantiates a new HttpErrorRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsiderUnknownErrorCode

`func (o *HttpErrorRule) GetConsiderUnknownErrorCode() bool`

GetConsiderUnknownErrorCode returns the ConsiderUnknownErrorCode field if non-nil, zero value otherwise.

### GetConsiderUnknownErrorCodeOk

`func (o *HttpErrorRule) GetConsiderUnknownErrorCodeOk() (*bool, bool)`

GetConsiderUnknownErrorCodeOk returns a tuple with the ConsiderUnknownErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsiderUnknownErrorCode

`func (o *HttpErrorRule) SetConsiderUnknownErrorCode(v bool)`

SetConsiderUnknownErrorCode sets ConsiderUnknownErrorCode field to given value.


### GetErrorCodes

`func (o *HttpErrorRule) GetErrorCodes() string`

GetErrorCodes returns the ErrorCodes field if non-nil, zero value otherwise.

### GetErrorCodesOk

`func (o *HttpErrorRule) GetErrorCodesOk() (*string, bool)`

GetErrorCodesOk returns a tuple with the ErrorCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCodes

`func (o *HttpErrorRule) SetErrorCodes(v string)`

SetErrorCodes sets ErrorCodes field to given value.

### HasErrorCodes

`func (o *HttpErrorRule) HasErrorCodes() bool`

HasErrorCodes returns a boolean if a field has been set.

### GetFilterByUrl

`func (o *HttpErrorRule) GetFilterByUrl() bool`

GetFilterByUrl returns the FilterByUrl field if non-nil, zero value otherwise.

### GetFilterByUrlOk

`func (o *HttpErrorRule) GetFilterByUrlOk() (*bool, bool)`

GetFilterByUrlOk returns a tuple with the FilterByUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterByUrl

`func (o *HttpErrorRule) SetFilterByUrl(v bool)`

SetFilterByUrl sets FilterByUrl field to given value.


### GetFilter

`func (o *HttpErrorRule) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *HttpErrorRule) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *HttpErrorRule) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *HttpErrorRule) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetUrl

`func (o *HttpErrorRule) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HttpErrorRule) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HttpErrorRule) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HttpErrorRule) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCapture

`func (o *HttpErrorRule) GetCapture() bool`

GetCapture returns the Capture field if non-nil, zero value otherwise.

### GetCaptureOk

`func (o *HttpErrorRule) GetCaptureOk() (*bool, bool)`

GetCaptureOk returns a tuple with the Capture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapture

`func (o *HttpErrorRule) SetCapture(v bool)`

SetCapture sets Capture field to given value.


### GetImpactApdex

`func (o *HttpErrorRule) GetImpactApdex() bool`

GetImpactApdex returns the ImpactApdex field if non-nil, zero value otherwise.

### GetImpactApdexOk

`func (o *HttpErrorRule) GetImpactApdexOk() (*bool, bool)`

GetImpactApdexOk returns a tuple with the ImpactApdex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactApdex

`func (o *HttpErrorRule) SetImpactApdex(v bool)`

SetImpactApdex sets ImpactApdex field to given value.


### GetConsiderForAi

`func (o *HttpErrorRule) GetConsiderForAi() bool`

GetConsiderForAi returns the ConsiderForAi field if non-nil, zero value otherwise.

### GetConsiderForAiOk

`func (o *HttpErrorRule) GetConsiderForAiOk() (*bool, bool)`

GetConsiderForAiOk returns a tuple with the ConsiderForAi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsiderForAi

`func (o *HttpErrorRule) SetConsiderForAi(v bool)`

SetConsiderForAi sets ConsiderForAi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


