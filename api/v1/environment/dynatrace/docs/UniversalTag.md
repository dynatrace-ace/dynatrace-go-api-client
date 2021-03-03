# UniversalTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagKey** | Pointer to [**UniversalTagKey**](UniversalTagKey.md) |  | [optional] 
**Value** | Pointer to **string** | The value of the tag. Not applicable to custom tags.   If a tag does have a separate key and value (in the textual representation they are split by the colon ‘:’), this field is set with the actual value. Key-value pairs can occur for automatically imported tags and tags set by rules if extractors are used. | [optional] 
**Key** | **string** | The key of the tag. For custom tags, put the tag value here.  The key allows categorization of multiple tags. It is possible that there are multiple values for a single key which will all be represented as standalone tags. Therefore, the key does not have the semantic of a map key but is more like a key of a key-value tuple. In some cases, for example custom tags, the key represents the actual tag value and the value field is not set – those are called valueless tags. | 
**Context** | Pointer to **string** | The origin of the tag, such as AWS or Cloud Foundry. For custom tags use the &#x60;CONTEXTLESS&#x60; value.   The context is set for tags that are automatically imported by OneAgent (for example, from the AWS console or environment variables). It’s useful for determining the origin of tags when not manually defined, and it also helps to prevent clashes with other existing tags. If the tag is not automatically imported, &#x60;CONTEXTLESS&#x60; set. | [optional] 

## Methods

### NewUniversalTag

`func NewUniversalTag(key string, ) *UniversalTag`

NewUniversalTag instantiates a new UniversalTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniversalTagWithDefaults

`func NewUniversalTagWithDefaults() *UniversalTag`

NewUniversalTagWithDefaults instantiates a new UniversalTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagKey

`func (o *UniversalTag) GetTagKey() UniversalTagKey`

GetTagKey returns the TagKey field if non-nil, zero value otherwise.

### GetTagKeyOk

`func (o *UniversalTag) GetTagKeyOk() (*UniversalTagKey, bool)`

GetTagKeyOk returns a tuple with the TagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKey

`func (o *UniversalTag) SetTagKey(v UniversalTagKey)`

SetTagKey sets TagKey field to given value.

### HasTagKey

`func (o *UniversalTag) HasTagKey() bool`

HasTagKey returns a boolean if a field has been set.

### GetValue

`func (o *UniversalTag) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UniversalTag) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UniversalTag) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UniversalTag) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetKey

`func (o *UniversalTag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UniversalTag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UniversalTag) SetKey(v string)`

SetKey sets Key field to given value.


### GetContext

`func (o *UniversalTag) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UniversalTag) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UniversalTag) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *UniversalTag) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


