# ValueProcessing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValueCondition** | Pointer to [**ValueCondition**](ValueCondition.md) |  | [optional] 
**ValueExtractorRegex** | Pointer to **string** | Extract value from captured data per regex. | [optional] 
**SplitAt** | Pointer to **string** | Split (preprocessed) string values at this separator. | [optional] 
**Trim** | **bool** | Prune Whitespaces. Defaults to false. | 
**ExtractSubstring** | Pointer to [**ExtractSubstring**](ExtractSubstring.md) |  | [optional] 

## Methods

### NewValueProcessing

`func NewValueProcessing(trim bool, ) *ValueProcessing`

NewValueProcessing instantiates a new ValueProcessing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueProcessingWithDefaults

`func NewValueProcessingWithDefaults() *ValueProcessing`

NewValueProcessingWithDefaults instantiates a new ValueProcessing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValueCondition

`func (o *ValueProcessing) GetValueCondition() ValueCondition`

GetValueCondition returns the ValueCondition field if non-nil, zero value otherwise.

### GetValueConditionOk

`func (o *ValueProcessing) GetValueConditionOk() (*ValueCondition, bool)`

GetValueConditionOk returns a tuple with the ValueCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCondition

`func (o *ValueProcessing) SetValueCondition(v ValueCondition)`

SetValueCondition sets ValueCondition field to given value.

### HasValueCondition

`func (o *ValueProcessing) HasValueCondition() bool`

HasValueCondition returns a boolean if a field has been set.

### GetValueExtractorRegex

`func (o *ValueProcessing) GetValueExtractorRegex() string`

GetValueExtractorRegex returns the ValueExtractorRegex field if non-nil, zero value otherwise.

### GetValueExtractorRegexOk

`func (o *ValueProcessing) GetValueExtractorRegexOk() (*string, bool)`

GetValueExtractorRegexOk returns a tuple with the ValueExtractorRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueExtractorRegex

`func (o *ValueProcessing) SetValueExtractorRegex(v string)`

SetValueExtractorRegex sets ValueExtractorRegex field to given value.

### HasValueExtractorRegex

`func (o *ValueProcessing) HasValueExtractorRegex() bool`

HasValueExtractorRegex returns a boolean if a field has been set.

### GetSplitAt

`func (o *ValueProcessing) GetSplitAt() string`

GetSplitAt returns the SplitAt field if non-nil, zero value otherwise.

### GetSplitAtOk

`func (o *ValueProcessing) GetSplitAtOk() (*string, bool)`

GetSplitAtOk returns a tuple with the SplitAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitAt

`func (o *ValueProcessing) SetSplitAt(v string)`

SetSplitAt sets SplitAt field to given value.

### HasSplitAt

`func (o *ValueProcessing) HasSplitAt() bool`

HasSplitAt returns a boolean if a field has been set.

### GetTrim

`func (o *ValueProcessing) GetTrim() bool`

GetTrim returns the Trim field if non-nil, zero value otherwise.

### GetTrimOk

`func (o *ValueProcessing) GetTrimOk() (*bool, bool)`

GetTrimOk returns a tuple with the Trim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrim

`func (o *ValueProcessing) SetTrim(v bool)`

SetTrim sets Trim field to given value.


### GetExtractSubstring

`func (o *ValueProcessing) GetExtractSubstring() ExtractSubstring`

GetExtractSubstring returns the ExtractSubstring field if non-nil, zero value otherwise.

### GetExtractSubstringOk

`func (o *ValueProcessing) GetExtractSubstringOk() (*ExtractSubstring, bool)`

GetExtractSubstringOk returns a tuple with the ExtractSubstring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractSubstring

`func (o *ValueProcessing) SetExtractSubstring(v ExtractSubstring)`

SetExtractSubstring sets ExtractSubstring field to given value.

### HasExtractSubstring

`func (o *ValueProcessing) HasExtractSubstring() bool`

HasExtractSubstring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


