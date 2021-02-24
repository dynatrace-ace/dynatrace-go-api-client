# BeforeTransformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delimiter** | **string** | The delimiter of the transformation. The transformation keeps everything before this delimiter and removes everything after it.    The delimiter itself is not kept.   If several delimiters appear in the initial value, only the first one is used. | 

## Methods

### NewBeforeTransformation

`func NewBeforeTransformation(delimiter string, ) *BeforeTransformation`

NewBeforeTransformation instantiates a new BeforeTransformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeforeTransformationWithDefaults

`func NewBeforeTransformationWithDefaults() *BeforeTransformation`

NewBeforeTransformationWithDefaults instantiates a new BeforeTransformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelimiter

`func (o *BeforeTransformation) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *BeforeTransformation) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *BeforeTransformation) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


