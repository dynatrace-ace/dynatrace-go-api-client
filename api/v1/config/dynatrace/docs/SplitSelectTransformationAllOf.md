# SplitSelectTransformationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delimiter** | Pointer to **string** | The delimiter for splitting the detected value. The delimiter itself is not kept. | [optional] 
**ItemIndex** | Pointer to **int32** | The index of the element in the split array to be used. Indexing starts with &#x60;1&#x60;. | [optional] 

## Methods

### NewSplitSelectTransformationAllOf

`func NewSplitSelectTransformationAllOf() *SplitSelectTransformationAllOf`

NewSplitSelectTransformationAllOf instantiates a new SplitSelectTransformationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitSelectTransformationAllOfWithDefaults

`func NewSplitSelectTransformationAllOfWithDefaults() *SplitSelectTransformationAllOf`

NewSplitSelectTransformationAllOfWithDefaults instantiates a new SplitSelectTransformationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelimiter

`func (o *SplitSelectTransformationAllOf) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *SplitSelectTransformationAllOf) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *SplitSelectTransformationAllOf) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *SplitSelectTransformationAllOf) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.

### GetItemIndex

`func (o *SplitSelectTransformationAllOf) GetItemIndex() int32`

GetItemIndex returns the ItemIndex field if non-nil, zero value otherwise.

### GetItemIndexOk

`func (o *SplitSelectTransformationAllOf) GetItemIndexOk() (*int32, bool)`

GetItemIndexOk returns a tuple with the ItemIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIndex

`func (o *SplitSelectTransformationAllOf) SetItemIndex(v int32)`

SetItemIndex sets ItemIndex field to given value.

### HasItemIndex

`func (o *SplitSelectTransformationAllOf) HasItemIndex() bool`

HasItemIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


