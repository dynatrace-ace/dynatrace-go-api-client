# Placeholder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the placeholder. Use it in the naming pattern as &#x60;{name}&#x60;. | 
**Attribute** | **string** | The attribute to extract from. You can only use attributes of the **string** type. | 
**Kind** | **string** | The type of extraction.    Defines either usage of regular expression (&#x60;regex&#x60;) or the position of request attribute value to be extracted.   When the **attribute** is &#x60;SERVICE_REQUEST_ATTRIBUTE&#x60; attribute and **aggregation** is &#x60;COUNT&#x60;, needs to be set to &#x60;ORIGINAL_TEXT&#x60; | 
**DelimiterOrRegex** | Pointer to **string** | Depending on the **type** value:    * &#x60;REGEX_EXTRACTION&#x60;: The regular expression.   * &#x60;BETWEEN_DELIMITER&#x60;: The opening delimiter string to look for.   * All other values: The delimiter string to look for. | [optional] 
**EndDelimiter** | Pointer to **string** | The closing delimiter string to look for.    Required if the **kind** value is &#x60;BETWEEN_DELIMITER&#x60;. Not applicable otherwise. | [optional] 
**RequestAttribute** | Pointer to **string** | The request attribute to extract from.    Required if the **kind** value is &#x60;SERVICE_REQUEST_ATTRIBUTE&#x60;. Not applicable otherwise. | [optional] 
**Normalization** | Pointer to **string** | The format of the extracted string. | [optional] 
**UseFromChildCalls** | Pointer to **bool** | If &#x60;true&#x60; request attribute will be taken from a child service call.    Only applicable for the &#x60;SERVICE_REQUEST_ATTRIBUTE&#x60; attribute. Defaults to &#x60;false&#x60;. | [optional] 
**Aggregation** | Pointer to **string** | Which value of the request attribute must be used when it occurs across multiple child requests.   Only applicable for the &#x60;SERVICE_REQUEST_ATTRIBUTE&#x60; attribute, when **useFromChildCalls** is &#x60;true&#x60;.   For the &#x60;COUNT&#x60; aggregation, the **kind** field is not applicable. | [optional] 
**Source** | Pointer to [**PropagationSource**](PropagationSource.md) |  | [optional] 

## Methods

### NewPlaceholder

`func NewPlaceholder(name string, attribute string, kind string, ) *Placeholder`

NewPlaceholder instantiates a new Placeholder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceholderWithDefaults

`func NewPlaceholderWithDefaults() *Placeholder`

NewPlaceholderWithDefaults instantiates a new Placeholder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Placeholder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Placeholder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Placeholder) SetName(v string)`

SetName sets Name field to given value.


### GetAttribute

`func (o *Placeholder) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *Placeholder) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *Placeholder) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetKind

`func (o *Placeholder) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Placeholder) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Placeholder) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDelimiterOrRegex

`func (o *Placeholder) GetDelimiterOrRegex() string`

GetDelimiterOrRegex returns the DelimiterOrRegex field if non-nil, zero value otherwise.

### GetDelimiterOrRegexOk

`func (o *Placeholder) GetDelimiterOrRegexOk() (*string, bool)`

GetDelimiterOrRegexOk returns a tuple with the DelimiterOrRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiterOrRegex

`func (o *Placeholder) SetDelimiterOrRegex(v string)`

SetDelimiterOrRegex sets DelimiterOrRegex field to given value.

### HasDelimiterOrRegex

`func (o *Placeholder) HasDelimiterOrRegex() bool`

HasDelimiterOrRegex returns a boolean if a field has been set.

### GetEndDelimiter

`func (o *Placeholder) GetEndDelimiter() string`

GetEndDelimiter returns the EndDelimiter field if non-nil, zero value otherwise.

### GetEndDelimiterOk

`func (o *Placeholder) GetEndDelimiterOk() (*string, bool)`

GetEndDelimiterOk returns a tuple with the EndDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDelimiter

`func (o *Placeholder) SetEndDelimiter(v string)`

SetEndDelimiter sets EndDelimiter field to given value.

### HasEndDelimiter

`func (o *Placeholder) HasEndDelimiter() bool`

HasEndDelimiter returns a boolean if a field has been set.

### GetRequestAttribute

`func (o *Placeholder) GetRequestAttribute() string`

GetRequestAttribute returns the RequestAttribute field if non-nil, zero value otherwise.

### GetRequestAttributeOk

`func (o *Placeholder) GetRequestAttributeOk() (*string, bool)`

GetRequestAttributeOk returns a tuple with the RequestAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAttribute

`func (o *Placeholder) SetRequestAttribute(v string)`

SetRequestAttribute sets RequestAttribute field to given value.

### HasRequestAttribute

`func (o *Placeholder) HasRequestAttribute() bool`

HasRequestAttribute returns a boolean if a field has been set.

### GetNormalization

`func (o *Placeholder) GetNormalization() string`

GetNormalization returns the Normalization field if non-nil, zero value otherwise.

### GetNormalizationOk

`func (o *Placeholder) GetNormalizationOk() (*string, bool)`

GetNormalizationOk returns a tuple with the Normalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalization

`func (o *Placeholder) SetNormalization(v string)`

SetNormalization sets Normalization field to given value.

### HasNormalization

`func (o *Placeholder) HasNormalization() bool`

HasNormalization returns a boolean if a field has been set.

### GetUseFromChildCalls

`func (o *Placeholder) GetUseFromChildCalls() bool`

GetUseFromChildCalls returns the UseFromChildCalls field if non-nil, zero value otherwise.

### GetUseFromChildCallsOk

`func (o *Placeholder) GetUseFromChildCallsOk() (*bool, bool)`

GetUseFromChildCallsOk returns a tuple with the UseFromChildCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFromChildCalls

`func (o *Placeholder) SetUseFromChildCalls(v bool)`

SetUseFromChildCalls sets UseFromChildCalls field to given value.

### HasUseFromChildCalls

`func (o *Placeholder) HasUseFromChildCalls() bool`

HasUseFromChildCalls returns a boolean if a field has been set.

### GetAggregation

`func (o *Placeholder) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *Placeholder) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *Placeholder) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *Placeholder) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetSource

`func (o *Placeholder) GetSource() PropagationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Placeholder) GetSourceOk() (*PropagationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Placeholder) SetSource(v PropagationSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Placeholder) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


