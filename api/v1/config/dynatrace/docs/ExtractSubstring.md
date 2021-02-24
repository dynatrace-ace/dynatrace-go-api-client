# ExtractSubstring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | **string** | The position of the extracted string relative to delimiters. | 
**Delimiter** | **string** | The delimiter string. | 
**EndDelimiter** | Pointer to **string** | The end-delimiter string.    Required if the **position** value is &#x60;BETWEEN&#x60;. Otherwise not allowed. | [optional] 

## Methods

### NewExtractSubstring

`func NewExtractSubstring(position string, delimiter string, ) *ExtractSubstring`

NewExtractSubstring instantiates a new ExtractSubstring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractSubstringWithDefaults

`func NewExtractSubstringWithDefaults() *ExtractSubstring`

NewExtractSubstringWithDefaults instantiates a new ExtractSubstring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *ExtractSubstring) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ExtractSubstring) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ExtractSubstring) SetPosition(v string)`

SetPosition sets Position field to given value.


### GetDelimiter

`func (o *ExtractSubstring) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *ExtractSubstring) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *ExtractSubstring) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.


### GetEndDelimiter

`func (o *ExtractSubstring) GetEndDelimiter() string`

GetEndDelimiter returns the EndDelimiter field if non-nil, zero value otherwise.

### GetEndDelimiterOk

`func (o *ExtractSubstring) GetEndDelimiterOk() (*string, bool)`

GetEndDelimiterOk returns a tuple with the EndDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDelimiter

`func (o *ExtractSubstring) SetEndDelimiter(v string)`

SetEndDelimiter sets EndDelimiter field to given value.

### HasEndDelimiter

`func (o *ExtractSubstring) HasEndDelimiter() bool`

HasEndDelimiter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


