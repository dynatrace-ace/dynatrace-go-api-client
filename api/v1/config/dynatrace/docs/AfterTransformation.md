# AfterTransformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delimiter** | **string** | The delimiter of the transformation. The transformation removes everything before this delimiter and keeps everything after it.     The delimiter itself is not kept.    If several delimiters appear in the initial value, only the first one is used. | 

## Methods

### NewAfterTransformation

`func NewAfterTransformation(delimiter string, ) *AfterTransformation`

NewAfterTransformation instantiates a new AfterTransformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfterTransformationWithDefaults

`func NewAfterTransformationWithDefaults() *AfterTransformation`

NewAfterTransformationWithDefaults instantiates a new AfterTransformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelimiter

`func (o *AfterTransformation) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *AfterTransformation) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *AfterTransformation) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


