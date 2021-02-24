# AfterTransformationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delimiter** | Pointer to **string** | The delimiter of the transformation. The transformation removes everything before this delimiter and keeps everything after it.     The delimiter itself is not kept.    If several delimiters appear in the initial value, only the first one is used. | [optional] 

## Methods

### NewAfterTransformationAllOf

`func NewAfterTransformationAllOf() *AfterTransformationAllOf`

NewAfterTransformationAllOf instantiates a new AfterTransformationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfterTransformationAllOfWithDefaults

`func NewAfterTransformationAllOfWithDefaults() *AfterTransformationAllOf`

NewAfterTransformationAllOfWithDefaults instantiates a new AfterTransformationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelimiter

`func (o *AfterTransformationAllOf) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *AfterTransformationAllOf) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *AfterTransformationAllOf) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *AfterTransformationAllOf) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


