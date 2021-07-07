# ComplexConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumPropertyCount** | Pointer to **int32** | The maximum number of properties that can be set. | [optional] 
**CustomValidatorId** | Pointer to **string** | The ID of a custom validator. | [optional] 
**MinimumPropertyCount** | Pointer to **int32** | The minimum number of properties that must be set. | [optional] 
**CustomMessage** | Pointer to **string** | A custom message for invalid values. | [optional] 
**Properties** | Pointer to **[]string** | A list of properties (defined by IDs) that are used to check the constraint. | [optional] 
**Type** | Pointer to **string** | The type of the constraint. | [optional] 

## Methods

### NewComplexConstraint

`func NewComplexConstraint() *ComplexConstraint`

NewComplexConstraint instantiates a new ComplexConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplexConstraintWithDefaults

`func NewComplexConstraintWithDefaults() *ComplexConstraint`

NewComplexConstraintWithDefaults instantiates a new ComplexConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumPropertyCount

`func (o *ComplexConstraint) GetMaximumPropertyCount() int32`

GetMaximumPropertyCount returns the MaximumPropertyCount field if non-nil, zero value otherwise.

### GetMaximumPropertyCountOk

`func (o *ComplexConstraint) GetMaximumPropertyCountOk() (*int32, bool)`

GetMaximumPropertyCountOk returns a tuple with the MaximumPropertyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPropertyCount

`func (o *ComplexConstraint) SetMaximumPropertyCount(v int32)`

SetMaximumPropertyCount sets MaximumPropertyCount field to given value.

### HasMaximumPropertyCount

`func (o *ComplexConstraint) HasMaximumPropertyCount() bool`

HasMaximumPropertyCount returns a boolean if a field has been set.

### GetCustomValidatorId

`func (o *ComplexConstraint) GetCustomValidatorId() string`

GetCustomValidatorId returns the CustomValidatorId field if non-nil, zero value otherwise.

### GetCustomValidatorIdOk

`func (o *ComplexConstraint) GetCustomValidatorIdOk() (*string, bool)`

GetCustomValidatorIdOk returns a tuple with the CustomValidatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValidatorId

`func (o *ComplexConstraint) SetCustomValidatorId(v string)`

SetCustomValidatorId sets CustomValidatorId field to given value.

### HasCustomValidatorId

`func (o *ComplexConstraint) HasCustomValidatorId() bool`

HasCustomValidatorId returns a boolean if a field has been set.

### GetMinimumPropertyCount

`func (o *ComplexConstraint) GetMinimumPropertyCount() int32`

GetMinimumPropertyCount returns the MinimumPropertyCount field if non-nil, zero value otherwise.

### GetMinimumPropertyCountOk

`func (o *ComplexConstraint) GetMinimumPropertyCountOk() (*int32, bool)`

GetMinimumPropertyCountOk returns a tuple with the MinimumPropertyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPropertyCount

`func (o *ComplexConstraint) SetMinimumPropertyCount(v int32)`

SetMinimumPropertyCount sets MinimumPropertyCount field to given value.

### HasMinimumPropertyCount

`func (o *ComplexConstraint) HasMinimumPropertyCount() bool`

HasMinimumPropertyCount returns a boolean if a field has been set.

### GetCustomMessage

`func (o *ComplexConstraint) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *ComplexConstraint) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *ComplexConstraint) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *ComplexConstraint) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.

### GetProperties

`func (o *ComplexConstraint) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ComplexConstraint) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ComplexConstraint) SetProperties(v []string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ComplexConstraint) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetType

`func (o *ComplexConstraint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComplexConstraint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComplexConstraint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ComplexConstraint) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


