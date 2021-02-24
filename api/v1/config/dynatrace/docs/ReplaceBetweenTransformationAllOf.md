# ReplaceBetweenTransformationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **string** | The starting delimiter. The transformation removes everything before it. The delimiter itself is not kept. | [optional] 
**Before** | Pointer to **string** | The ending delimiter. The transformation removes everything after it. The delimiter itself is not kept. | [optional] 
**Replacement** | Pointer to **string** | The value to be used instead of the detected value. | [optional] 

## Methods

### NewReplaceBetweenTransformationAllOf

`func NewReplaceBetweenTransformationAllOf() *ReplaceBetweenTransformationAllOf`

NewReplaceBetweenTransformationAllOf instantiates a new ReplaceBetweenTransformationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceBetweenTransformationAllOfWithDefaults

`func NewReplaceBetweenTransformationAllOfWithDefaults() *ReplaceBetweenTransformationAllOf`

NewReplaceBetweenTransformationAllOfWithDefaults instantiates a new ReplaceBetweenTransformationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *ReplaceBetweenTransformationAllOf) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ReplaceBetweenTransformationAllOf) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ReplaceBetweenTransformationAllOf) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *ReplaceBetweenTransformationAllOf) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetBefore

`func (o *ReplaceBetweenTransformationAllOf) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *ReplaceBetweenTransformationAllOf) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *ReplaceBetweenTransformationAllOf) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *ReplaceBetweenTransformationAllOf) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetReplacement

`func (o *ReplaceBetweenTransformationAllOf) GetReplacement() string`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *ReplaceBetweenTransformationAllOf) GetReplacementOk() (*string, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *ReplaceBetweenTransformationAllOf) SetReplacement(v string)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *ReplaceBetweenTransformationAllOf) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


