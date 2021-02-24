# SplitSelectTransformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delimiter** | **string** | The delimiter for splitting the detected value. The delimiter itself is not kept. | 
**ItemIndex** | **int32** | The index of the element in the split array to be used. Indexing starts with &#x60;1&#x60;. | 

## Methods

### NewSplitSelectTransformation

`func NewSplitSelectTransformation(delimiter string, itemIndex int32, ) *SplitSelectTransformation`

NewSplitSelectTransformation instantiates a new SplitSelectTransformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitSelectTransformationWithDefaults

`func NewSplitSelectTransformationWithDefaults() *SplitSelectTransformation`

NewSplitSelectTransformationWithDefaults instantiates a new SplitSelectTransformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelimiter

`func (o *SplitSelectTransformation) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *SplitSelectTransformation) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *SplitSelectTransformation) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.


### GetItemIndex

`func (o *SplitSelectTransformation) GetItemIndex() int32`

GetItemIndex returns the ItemIndex field if non-nil, zero value otherwise.

### GetItemIndexOk

`func (o *SplitSelectTransformation) GetItemIndexOk() (*int32, bool)`

GetItemIndexOk returns a tuple with the ItemIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIndex

`func (o *SplitSelectTransformation) SetItemIndex(v int32)`

SetItemIndex sets ItemIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


