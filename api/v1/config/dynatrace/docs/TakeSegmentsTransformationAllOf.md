# TakeSegmentsTransformationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentCount** | Pointer to **int32** | The number of elements to be kept. | [optional] 
**Delimiter** | Pointer to **string** | The delimiter for splitting the detected value. The delimiter itself is not kept. | [optional] 
**TakeFromEnd** | Pointer to **bool** | Keeps the first (&#x60;false&#x60;) or last (&#x60;true&#x60;) elements.    If not set, then &#x60;false&#x60; is used, keeping the first elements. | [optional] 

## Methods

### NewTakeSegmentsTransformationAllOf

`func NewTakeSegmentsTransformationAllOf() *TakeSegmentsTransformationAllOf`

NewTakeSegmentsTransformationAllOf instantiates a new TakeSegmentsTransformationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTakeSegmentsTransformationAllOfWithDefaults

`func NewTakeSegmentsTransformationAllOfWithDefaults() *TakeSegmentsTransformationAllOf`

NewTakeSegmentsTransformationAllOfWithDefaults instantiates a new TakeSegmentsTransformationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentCount

`func (o *TakeSegmentsTransformationAllOf) GetSegmentCount() int32`

GetSegmentCount returns the SegmentCount field if non-nil, zero value otherwise.

### GetSegmentCountOk

`func (o *TakeSegmentsTransformationAllOf) GetSegmentCountOk() (*int32, bool)`

GetSegmentCountOk returns a tuple with the SegmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentCount

`func (o *TakeSegmentsTransformationAllOf) SetSegmentCount(v int32)`

SetSegmentCount sets SegmentCount field to given value.

### HasSegmentCount

`func (o *TakeSegmentsTransformationAllOf) HasSegmentCount() bool`

HasSegmentCount returns a boolean if a field has been set.

### GetDelimiter

`func (o *TakeSegmentsTransformationAllOf) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *TakeSegmentsTransformationAllOf) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *TakeSegmentsTransformationAllOf) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *TakeSegmentsTransformationAllOf) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.

### GetTakeFromEnd

`func (o *TakeSegmentsTransformationAllOf) GetTakeFromEnd() bool`

GetTakeFromEnd returns the TakeFromEnd field if non-nil, zero value otherwise.

### GetTakeFromEndOk

`func (o *TakeSegmentsTransformationAllOf) GetTakeFromEndOk() (*bool, bool)`

GetTakeFromEndOk returns a tuple with the TakeFromEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeFromEnd

`func (o *TakeSegmentsTransformationAllOf) SetTakeFromEnd(v bool)`

SetTakeFromEnd sets TakeFromEnd field to given value.

### HasTakeFromEnd

`func (o *TakeSegmentsTransformationAllOf) HasTakeFromEnd() bool`

HasTakeFromEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


