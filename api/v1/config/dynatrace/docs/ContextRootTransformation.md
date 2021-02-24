# ContextRootTransformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;BEFORE&#x60; -&gt; BeforeTransformation  * &#x60;REPLACE_BETWEEN&#x60; -&gt; ReplaceBetweenTransformation  * &#x60;REMOVE_NUMBERS&#x60; -&gt; RemoveNumbersTransformation  * &#x60;REMOVE_CREDIT_CARDS&#x60; -&gt; RemoveCreditCardNumbersTransformation  * &#x60;REMOVE_IBANS&#x60; -&gt; RemoveIBANsTransformation  * &#x60;REMOVE_IPS&#x60; -&gt; RemoveIPsTransformation   | 

## Methods

### NewContextRootTransformation

`func NewContextRootTransformation(type_ string, ) *ContextRootTransformation`

NewContextRootTransformation instantiates a new ContextRootTransformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextRootTransformationWithDefaults

`func NewContextRootTransformationWithDefaults() *ContextRootTransformation`

NewContextRootTransformationWithDefaults instantiates a new ContextRootTransformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContextRootTransformation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContextRootTransformation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContextRootTransformation) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


