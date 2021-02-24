# PublicDomainName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transformations** | Pointer to [**[]TransformationBase**](TransformationBase.md) | Transformations to be applied to the detected value. | [optional] 
**ValueOverride** | Pointer to **string** | The value to be used instead of the detected value. | [optional] 
**CopyFromHostName** | Pointer to **bool** | Use (&#x60;true&#x60;) or don&#39;t use (&#x60;false&#x60;) the detected host name as base for transformation.    Not applicable if the override is specified. | [optional] 

## Methods

### NewPublicDomainName

`func NewPublicDomainName() *PublicDomainName`

NewPublicDomainName instantiates a new PublicDomainName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicDomainNameWithDefaults

`func NewPublicDomainNameWithDefaults() *PublicDomainName`

NewPublicDomainNameWithDefaults instantiates a new PublicDomainName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransformations

`func (o *PublicDomainName) GetTransformations() []TransformationBase`

GetTransformations returns the Transformations field if non-nil, zero value otherwise.

### GetTransformationsOk

`func (o *PublicDomainName) GetTransformationsOk() (*[]TransformationBase, bool)`

GetTransformationsOk returns a tuple with the Transformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformations

`func (o *PublicDomainName) SetTransformations(v []TransformationBase)`

SetTransformations sets Transformations field to given value.

### HasTransformations

`func (o *PublicDomainName) HasTransformations() bool`

HasTransformations returns a boolean if a field has been set.

### GetValueOverride

`func (o *PublicDomainName) GetValueOverride() string`

GetValueOverride returns the ValueOverride field if non-nil, zero value otherwise.

### GetValueOverrideOk

`func (o *PublicDomainName) GetValueOverrideOk() (*string, bool)`

GetValueOverrideOk returns a tuple with the ValueOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueOverride

`func (o *PublicDomainName) SetValueOverride(v string)`

SetValueOverride sets ValueOverride field to given value.

### HasValueOverride

`func (o *PublicDomainName) HasValueOverride() bool`

HasValueOverride returns a boolean if a field has been set.

### GetCopyFromHostName

`func (o *PublicDomainName) GetCopyFromHostName() bool`

GetCopyFromHostName returns the CopyFromHostName field if non-nil, zero value otherwise.

### GetCopyFromHostNameOk

`func (o *PublicDomainName) GetCopyFromHostNameOk() (*bool, bool)`

GetCopyFromHostNameOk returns a tuple with the CopyFromHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyFromHostName

`func (o *PublicDomainName) SetCopyFromHostName(v bool)`

SetCopyFromHostName sets CopyFromHostName field to given value.

### HasCopyFromHostName

`func (o *PublicDomainName) HasCopyFromHostName() bool`

HasCopyFromHostName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


