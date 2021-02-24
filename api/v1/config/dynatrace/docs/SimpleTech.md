# SimpleTech

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Predefined technology, if technology is not predefined, then the verbatim type must be set | [optional] 
**VerbatimType** | Pointer to **string** | Non-predefined technology, use for custom technologies. | [optional] 

## Methods

### NewSimpleTech

`func NewSimpleTech() *SimpleTech`

NewSimpleTech instantiates a new SimpleTech object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleTechWithDefaults

`func NewSimpleTechWithDefaults() *SimpleTech`

NewSimpleTechWithDefaults instantiates a new SimpleTech object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SimpleTech) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimpleTech) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimpleTech) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SimpleTech) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVerbatimType

`func (o *SimpleTech) GetVerbatimType() string`

GetVerbatimType returns the VerbatimType field if non-nil, zero value otherwise.

### GetVerbatimTypeOk

`func (o *SimpleTech) GetVerbatimTypeOk() (*string, bool)`

GetVerbatimTypeOk returns a tuple with the VerbatimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbatimType

`func (o *SimpleTech) SetVerbatimType(v string)`

SetVerbatimType sets VerbatimType field to given value.

### HasVerbatimType

`func (o *SimpleTech) HasVerbatimType() bool`

HasVerbatimType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


