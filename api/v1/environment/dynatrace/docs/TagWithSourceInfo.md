# TagWithSourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | The source of the tag, such as USER, RULE_BASED or AUTO | [optional] 
**Context** | **string** | The origin of the tag, such as AWS or Cloud Foundry.    Custom tags use the &#x60;CONTEXTLESS&#x60; value. | 
**Key** | **string** | The key of the tag.    Custom tags have the tag value here. | 
**Value** | Pointer to **string** | The value of the tag.    Not applicable to custom tags. | [optional] 

## Methods

### NewTagWithSourceInfo

`func NewTagWithSourceInfo(context string, key string, ) *TagWithSourceInfo`

NewTagWithSourceInfo instantiates a new TagWithSourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagWithSourceInfoWithDefaults

`func NewTagWithSourceInfoWithDefaults() *TagWithSourceInfo`

NewTagWithSourceInfoWithDefaults instantiates a new TagWithSourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TagWithSourceInfo) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TagWithSourceInfo) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TagWithSourceInfo) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TagWithSourceInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetContext

`func (o *TagWithSourceInfo) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TagWithSourceInfo) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TagWithSourceInfo) SetContext(v string)`

SetContext sets Context field to given value.


### GetKey

`func (o *TagWithSourceInfo) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TagWithSourceInfo) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TagWithSourceInfo) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *TagWithSourceInfo) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TagWithSourceInfo) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TagWithSourceInfo) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TagWithSourceInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


