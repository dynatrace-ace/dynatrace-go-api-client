# SimpleHostTech

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Predefined technology, if technology is not predefined, then the verbatim type must be set | [optional] 
**VerbatimType** | Pointer to **string** | Non-predefined technology, use for custom technologies. | [optional] 

## Methods

### NewSimpleHostTech

`func NewSimpleHostTech() *SimpleHostTech`

NewSimpleHostTech instantiates a new SimpleHostTech object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleHostTechWithDefaults

`func NewSimpleHostTechWithDefaults() *SimpleHostTech`

NewSimpleHostTechWithDefaults instantiates a new SimpleHostTech object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SimpleHostTech) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimpleHostTech) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimpleHostTech) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SimpleHostTech) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVerbatimType

`func (o *SimpleHostTech) GetVerbatimType() string`

GetVerbatimType returns the VerbatimType field if non-nil, zero value otherwise.

### GetVerbatimTypeOk

`func (o *SimpleHostTech) GetVerbatimTypeOk() (*string, bool)`

GetVerbatimTypeOk returns a tuple with the VerbatimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbatimType

`func (o *SimpleHostTech) SetVerbatimType(v string)`

SetVerbatimType sets VerbatimType field to given value.

### HasVerbatimType

`func (o *SimpleHostTech) HasVerbatimType() bool`

HasVerbatimType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


