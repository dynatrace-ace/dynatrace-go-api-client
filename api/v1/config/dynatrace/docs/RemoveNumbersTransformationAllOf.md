# RemoveNumbersTransformationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinDigitCount** | Pointer to **int32** | Remove numbers that contain at least *X* digits. | [optional] 
**IncludeHexNumbers** | Pointer to **bool** | Remove (&#x60;true&#x60;) or keep (&#x60;false&#x60;) hexadecimal numbers.    If not set, then &#x60;false&#x60; is used, keeping hexadecimal numbers. | [optional] 

## Methods

### NewRemoveNumbersTransformationAllOf

`func NewRemoveNumbersTransformationAllOf() *RemoveNumbersTransformationAllOf`

NewRemoveNumbersTransformationAllOf instantiates a new RemoveNumbersTransformationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveNumbersTransformationAllOfWithDefaults

`func NewRemoveNumbersTransformationAllOfWithDefaults() *RemoveNumbersTransformationAllOf`

NewRemoveNumbersTransformationAllOfWithDefaults instantiates a new RemoveNumbersTransformationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinDigitCount

`func (o *RemoveNumbersTransformationAllOf) GetMinDigitCount() int32`

GetMinDigitCount returns the MinDigitCount field if non-nil, zero value otherwise.

### GetMinDigitCountOk

`func (o *RemoveNumbersTransformationAllOf) GetMinDigitCountOk() (*int32, bool)`

GetMinDigitCountOk returns a tuple with the MinDigitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDigitCount

`func (o *RemoveNumbersTransformationAllOf) SetMinDigitCount(v int32)`

SetMinDigitCount sets MinDigitCount field to given value.

### HasMinDigitCount

`func (o *RemoveNumbersTransformationAllOf) HasMinDigitCount() bool`

HasMinDigitCount returns a boolean if a field has been set.

### GetIncludeHexNumbers

`func (o *RemoveNumbersTransformationAllOf) GetIncludeHexNumbers() bool`

GetIncludeHexNumbers returns the IncludeHexNumbers field if non-nil, zero value otherwise.

### GetIncludeHexNumbersOk

`func (o *RemoveNumbersTransformationAllOf) GetIncludeHexNumbersOk() (*bool, bool)`

GetIncludeHexNumbersOk returns a tuple with the IncludeHexNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHexNumbers

`func (o *RemoveNumbersTransformationAllOf) SetIncludeHexNumbers(v bool)`

SetIncludeHexNumbers sets IncludeHexNumbers field to given value.

### HasIncludeHexNumbers

`func (o *RemoveNumbersTransformationAllOf) HasIncludeHexNumbers() bool`

HasIncludeHexNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


