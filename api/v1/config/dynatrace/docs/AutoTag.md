# AutoTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the auto-tag. | [optional] 
**Name** | **string** | The name of the auto-tag, which is applied to entities.   Additionally you can specify a **valueFormat** in the tag rule. In that case the tag is used in the &#x60;name:valueFormat&#x60; format.   For example you can extend the &#x60;Infrastructure&#x60; tag to &#x60;Infrastructure:Windows&#x60; and &#x60;Infrastructure:Linux&#x60;. | 
**Rules** | Pointer to [**[]AutoTagRule**](AutoTagRule.md) | The list of rules for tag usage.   When there are multiple rules, the OR logic applies. | [optional] 

## Methods

### NewAutoTag

`func NewAutoTag(name string, ) *AutoTag`

NewAutoTag instantiates a new AutoTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTagWithDefaults

`func NewAutoTagWithDefaults() *AutoTag`

NewAutoTagWithDefaults instantiates a new AutoTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *AutoTag) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AutoTag) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AutoTag) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AutoTag) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *AutoTag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoTag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoTag) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoTag) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AutoTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoTag) SetName(v string)`

SetName sets Name field to given value.


### GetRules

`func (o *AutoTag) GetRules() []AutoTagRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AutoTag) GetRulesOk() (*[]AutoTagRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AutoTag) SetRules(v []AutoTagRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *AutoTag) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


