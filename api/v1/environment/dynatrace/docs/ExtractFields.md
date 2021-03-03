# ExtractFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParsingMode** | Pointer to **string** | The parsing mode for log analysis entries presentation. Possible values are: &#x60;json&#x60;, &#x60;disabled&#x60;, and &#x60;all&#x60;. | [optional] 
**CustomParsingPatterns** | Pointer to **string** | The query for content extraction.   See the [Search patterns in log data and parse results](https://dt-url.net/vv83rhp) help page for the syntax definition and examples. | [optional] 

## Methods

### NewExtractFields

`func NewExtractFields() *ExtractFields`

NewExtractFields instantiates a new ExtractFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractFieldsWithDefaults

`func NewExtractFieldsWithDefaults() *ExtractFields`

NewExtractFieldsWithDefaults instantiates a new ExtractFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParsingMode

`func (o *ExtractFields) GetParsingMode() string`

GetParsingMode returns the ParsingMode field if non-nil, zero value otherwise.

### GetParsingModeOk

`func (o *ExtractFields) GetParsingModeOk() (*string, bool)`

GetParsingModeOk returns a tuple with the ParsingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsingMode

`func (o *ExtractFields) SetParsingMode(v string)`

SetParsingMode sets ParsingMode field to given value.

### HasParsingMode

`func (o *ExtractFields) HasParsingMode() bool`

HasParsingMode returns a boolean if a field has been set.

### GetCustomParsingPatterns

`func (o *ExtractFields) GetCustomParsingPatterns() string`

GetCustomParsingPatterns returns the CustomParsingPatterns field if non-nil, zero value otherwise.

### GetCustomParsingPatternsOk

`func (o *ExtractFields) GetCustomParsingPatternsOk() (*string, bool)`

GetCustomParsingPatternsOk returns a tuple with the CustomParsingPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParsingPatterns

`func (o *ExtractFields) SetCustomParsingPatterns(v string)`

SetCustomParsingPatterns sets CustomParsingPatterns field to given value.

### HasCustomParsingPatterns

`func (o *ExtractFields) HasCustomParsingPatterns() bool`

HasCustomParsingPatterns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


