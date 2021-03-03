# ParsingFieldTopValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **string** | Top value parsing field name | [optional] 
**Occurrences** | Pointer to [**[]Occurrence**](Occurrence.md) | Top value parsing field occurrences | [optional] 

## Methods

### NewParsingFieldTopValue

`func NewParsingFieldTopValue() *ParsingFieldTopValue`

NewParsingFieldTopValue instantiates a new ParsingFieldTopValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParsingFieldTopValueWithDefaults

`func NewParsingFieldTopValueWithDefaults() *ParsingFieldTopValue`

NewParsingFieldTopValueWithDefaults instantiates a new ParsingFieldTopValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *ParsingFieldTopValue) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ParsingFieldTopValue) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ParsingFieldTopValue) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *ParsingFieldTopValue) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetOccurrences

`func (o *ParsingFieldTopValue) GetOccurrences() []Occurrence`

GetOccurrences returns the Occurrences field if non-nil, zero value otherwise.

### GetOccurrencesOk

`func (o *ParsingFieldTopValue) GetOccurrencesOk() (*[]Occurrence, bool)`

GetOccurrencesOk returns a tuple with the Occurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrences

`func (o *ParsingFieldTopValue) SetOccurrences(v []Occurrence)`

SetOccurrences sets Occurrences field to given value.

### HasOccurrences

`func (o *ParsingFieldTopValue) HasOccurrences() bool`

HasOccurrences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


