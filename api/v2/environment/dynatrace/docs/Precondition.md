# Precondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Precondition** | Pointer to [**Precondition**](Precondition.md) |  | [optional] 
**ExpectedValues** | Pointer to **[]map[string]interface{}** | A list of valid values of the property.   Only applicable to properties of the &#x60;IN&#x60; type. | [optional] 
**Preconditions** | Pointer to [**[]Precondition**](Precondition.md) | A list of child preconditions to be evaluated.   Only applicable to properties of the &#x60;AND&#x60; and &#x60;OR&#x60; types. | [optional] 
**ExpectedValue** | Pointer to **map[string]interface{}** | The expected value of the property.   Only applicable to properties of the &#x60;EQUALS&#x60; type. | [optional] 
**Property** | Pointer to **string** | The property to be evaluated. | [optional] 
**Type** | Pointer to **string** | The type of the precondition. | [optional] 

## Methods

### NewPrecondition

`func NewPrecondition() *Precondition`

NewPrecondition instantiates a new Precondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreconditionWithDefaults

`func NewPreconditionWithDefaults() *Precondition`

NewPreconditionWithDefaults instantiates a new Precondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrecondition

`func (o *Precondition) GetPrecondition() Precondition`

GetPrecondition returns the Precondition field if non-nil, zero value otherwise.

### GetPreconditionOk

`func (o *Precondition) GetPreconditionOk() (*Precondition, bool)`

GetPreconditionOk returns a tuple with the Precondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecondition

`func (o *Precondition) SetPrecondition(v Precondition)`

SetPrecondition sets Precondition field to given value.

### HasPrecondition

`func (o *Precondition) HasPrecondition() bool`

HasPrecondition returns a boolean if a field has been set.

### GetExpectedValues

`func (o *Precondition) GetExpectedValues() []map[string]interface{}`

GetExpectedValues returns the ExpectedValues field if non-nil, zero value otherwise.

### GetExpectedValuesOk

`func (o *Precondition) GetExpectedValuesOk() (*[]map[string]interface{}, bool)`

GetExpectedValuesOk returns a tuple with the ExpectedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedValues

`func (o *Precondition) SetExpectedValues(v []map[string]interface{})`

SetExpectedValues sets ExpectedValues field to given value.

### HasExpectedValues

`func (o *Precondition) HasExpectedValues() bool`

HasExpectedValues returns a boolean if a field has been set.

### GetPreconditions

`func (o *Precondition) GetPreconditions() []Precondition`

GetPreconditions returns the Preconditions field if non-nil, zero value otherwise.

### GetPreconditionsOk

`func (o *Precondition) GetPreconditionsOk() (*[]Precondition, bool)`

GetPreconditionsOk returns a tuple with the Preconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditions

`func (o *Precondition) SetPreconditions(v []Precondition)`

SetPreconditions sets Preconditions field to given value.

### HasPreconditions

`func (o *Precondition) HasPreconditions() bool`

HasPreconditions returns a boolean if a field has been set.

### GetExpectedValue

`func (o *Precondition) GetExpectedValue() map[string]interface{}`

GetExpectedValue returns the ExpectedValue field if non-nil, zero value otherwise.

### GetExpectedValueOk

`func (o *Precondition) GetExpectedValueOk() (*map[string]interface{}, bool)`

GetExpectedValueOk returns a tuple with the ExpectedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedValue

`func (o *Precondition) SetExpectedValue(v map[string]interface{})`

SetExpectedValue sets ExpectedValue field to given value.

### HasExpectedValue

`func (o *Precondition) HasExpectedValue() bool`

HasExpectedValue returns a boolean if a field has been set.

### GetProperty

`func (o *Precondition) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *Precondition) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *Precondition) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *Precondition) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetType

`func (o *Precondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Precondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Precondition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Precondition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


