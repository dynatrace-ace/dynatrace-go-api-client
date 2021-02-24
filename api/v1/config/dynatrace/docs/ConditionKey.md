# ConditionKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | **string** | The attribute to be used for comparision. | 
**Type** | Pointer to **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;PROCESS_CUSTOM_METADATA_KEY&#x60; -&gt; CustomProcessMetadataConditionKey  * &#x60;HOST_CUSTOM_METADATA_KEY&#x60; -&gt; CustomHostMetadataConditionKey  * &#x60;PROCESS_PREDEFINED_METADATA_KEY&#x60; -&gt; ProcessMetadataConditionKey  * &#x60;STRING&#x60; -&gt; StringConditionKey   | [optional] 
**DynamicKey** | Pointer to **map[string]interface{}** | Dynamic key generated based on selected type/attribute. | [optional] 

## Methods

### NewConditionKey

`func NewConditionKey(attribute string, ) *ConditionKey`

NewConditionKey instantiates a new ConditionKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionKeyWithDefaults

`func NewConditionKeyWithDefaults() *ConditionKey`

NewConditionKeyWithDefaults instantiates a new ConditionKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *ConditionKey) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *ConditionKey) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *ConditionKey) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetType

`func (o *ConditionKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConditionKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConditionKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConditionKey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDynamicKey

`func (o *ConditionKey) GetDynamicKey() map[string]interface{}`

GetDynamicKey returns the DynamicKey field if non-nil, zero value otherwise.

### GetDynamicKeyOk

`func (o *ConditionKey) GetDynamicKeyOk() (*map[string]interface{}, bool)`

GetDynamicKeyOk returns a tuple with the DynamicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicKey

`func (o *ConditionKey) SetDynamicKey(v map[string]interface{})`

SetDynamicKey sets DynamicKey field to given value.

### HasDynamicKey

`func (o *ConditionKey) HasDynamicKey() bool`

HasDynamicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


