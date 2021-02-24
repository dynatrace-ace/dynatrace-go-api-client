# WebServiceName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transformations** | Pointer to [**[]TransformationBase**](TransformationBase.md) | Transformations to be applied to the detected value. | [optional] 
**ValueOverride** | Pointer to **string** | The value to be used instead of the detected value. | [optional] 

## Methods

### NewWebServiceName

`func NewWebServiceName() *WebServiceName`

NewWebServiceName instantiates a new WebServiceName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebServiceNameWithDefaults

`func NewWebServiceNameWithDefaults() *WebServiceName`

NewWebServiceNameWithDefaults instantiates a new WebServiceName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransformations

`func (o *WebServiceName) GetTransformations() []TransformationBase`

GetTransformations returns the Transformations field if non-nil, zero value otherwise.

### GetTransformationsOk

`func (o *WebServiceName) GetTransformationsOk() (*[]TransformationBase, bool)`

GetTransformationsOk returns a tuple with the Transformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformations

`func (o *WebServiceName) SetTransformations(v []TransformationBase)`

SetTransformations sets Transformations field to given value.

### HasTransformations

`func (o *WebServiceName) HasTransformations() bool`

HasTransformations returns a boolean if a field has been set.

### GetValueOverride

`func (o *WebServiceName) GetValueOverride() string`

GetValueOverride returns the ValueOverride field if non-nil, zero value otherwise.

### GetValueOverrideOk

`func (o *WebServiceName) GetValueOverrideOk() (*string, bool)`

GetValueOverrideOk returns a tuple with the ValueOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueOverride

`func (o *WebServiceName) SetValueOverride(v string)`

SetValueOverride sets ValueOverride field to given value.

### HasValueOverride

`func (o *WebServiceName) HasValueOverride() bool`

HasValueOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


