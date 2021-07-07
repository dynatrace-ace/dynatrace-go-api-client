# Constraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxLength** | Pointer to **int32** | The maximum allowed length of string values. | [optional] 
**MinLength** | Pointer to **int32** | The minimum required length of string values. | [optional] 
**CustomValidatorId** | Pointer to **string** | The ID of a custom validator. | [optional] 
**CustomMessage** | Pointer to **string** | A custom message for invalid values. | [optional] 
**UniqueProperties** | Pointer to **[]string** | A list of properties for which the combination of values must be unique. | [optional] 
**Maximum** | Pointer to **float32** | The maximum allowed value. | [optional] 
**Minimum** | Pointer to **float32** | The minimum allowed value. | [optional] 
**Pattern** | Pointer to **string** | The regular expression pattern for valid string values. | [optional] 
**Type** | Pointer to **string** | The type of the constraint. | [optional] 

## Methods

### NewConstraint

`func NewConstraint() *Constraint`

NewConstraint instantiates a new Constraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstraintWithDefaults

`func NewConstraintWithDefaults() *Constraint`

NewConstraintWithDefaults instantiates a new Constraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxLength

`func (o *Constraint) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *Constraint) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *Constraint) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *Constraint) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMinLength

`func (o *Constraint) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *Constraint) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *Constraint) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *Constraint) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetCustomValidatorId

`func (o *Constraint) GetCustomValidatorId() string`

GetCustomValidatorId returns the CustomValidatorId field if non-nil, zero value otherwise.

### GetCustomValidatorIdOk

`func (o *Constraint) GetCustomValidatorIdOk() (*string, bool)`

GetCustomValidatorIdOk returns a tuple with the CustomValidatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValidatorId

`func (o *Constraint) SetCustomValidatorId(v string)`

SetCustomValidatorId sets CustomValidatorId field to given value.

### HasCustomValidatorId

`func (o *Constraint) HasCustomValidatorId() bool`

HasCustomValidatorId returns a boolean if a field has been set.

### GetCustomMessage

`func (o *Constraint) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *Constraint) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *Constraint) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *Constraint) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.

### GetUniqueProperties

`func (o *Constraint) GetUniqueProperties() []string`

GetUniqueProperties returns the UniqueProperties field if non-nil, zero value otherwise.

### GetUniquePropertiesOk

`func (o *Constraint) GetUniquePropertiesOk() (*[]string, bool)`

GetUniquePropertiesOk returns a tuple with the UniqueProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueProperties

`func (o *Constraint) SetUniqueProperties(v []string)`

SetUniqueProperties sets UniqueProperties field to given value.

### HasUniqueProperties

`func (o *Constraint) HasUniqueProperties() bool`

HasUniqueProperties returns a boolean if a field has been set.

### GetMaximum

`func (o *Constraint) GetMaximum() float32`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *Constraint) GetMaximumOk() (*float32, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *Constraint) SetMaximum(v float32)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *Constraint) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetMinimum

`func (o *Constraint) GetMinimum() float32`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *Constraint) GetMinimumOk() (*float32, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *Constraint) SetMinimum(v float32)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *Constraint) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetPattern

`func (o *Constraint) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *Constraint) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *Constraint) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *Constraint) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetType

`func (o *Constraint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Constraint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Constraint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Constraint) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


