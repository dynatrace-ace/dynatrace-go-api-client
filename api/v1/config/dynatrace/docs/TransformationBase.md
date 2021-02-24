# TransformationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;BEFORE&#x60; -&gt; BeforeTransformation  * &#x60;AFTER&#x60; -&gt; AfterTransformation  * &#x60;BETWEEN&#x60; -&gt; BetweenTransformation  * &#x60;REPLACE_BETWEEN&#x60; -&gt; ReplaceBetweenTransformation  * &#x60;REMOVE_NUMBERS&#x60; -&gt; RemoveNumbersTransformation  * &#x60;REMOVE_CREDIT_CARDS&#x60; -&gt; RemoveCreditCardNumbersTransformation  * &#x60;REMOVE_IBANS&#x60; -&gt; RemoveIBANsTransformation  * &#x60;REMOVE_IPS&#x60; -&gt; RemoveIPsTransformation  * &#x60;SPLIT_SELECT&#x60; -&gt; SplitSelectTransformation  * &#x60;TAKE_SEGMENTS&#x60; -&gt; TakeSegmentsTransformation   | 

## Methods

### NewTransformationBase

`func NewTransformationBase(type_ string, ) *TransformationBase`

NewTransformationBase instantiates a new TransformationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformationBaseWithDefaults

`func NewTransformationBaseWithDefaults() *TransformationBase`

NewTransformationBaseWithDefaults instantiates a new TransformationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransformationBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransformationBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransformationBase) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


