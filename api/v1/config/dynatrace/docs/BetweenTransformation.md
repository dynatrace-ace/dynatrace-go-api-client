# BetweenTransformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | **string** | The starting delimiter. The transformation removes everything before it. The delimiter itself is not kept.  | 
**Before** | **string** | The ending delimiter. The transformation removes everything after it. The delimiter itself is not kept. | 

## Methods

### NewBetweenTransformation

`func NewBetweenTransformation(after string, before string, ) *BetweenTransformation`

NewBetweenTransformation instantiates a new BetweenTransformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetweenTransformationWithDefaults

`func NewBetweenTransformationWithDefaults() *BetweenTransformation`

NewBetweenTransformationWithDefaults instantiates a new BetweenTransformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *BetweenTransformation) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *BetweenTransformation) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *BetweenTransformation) SetAfter(v string)`

SetAfter sets After field to given value.


### GetBefore

`func (o *BetweenTransformation) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *BetweenTransformation) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *BetweenTransformation) SetBefore(v string)`

SetBefore sets Before field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


