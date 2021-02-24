# ConstraintViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**ParameterLocation** | Pointer to **string** |  | [optional] 

## Methods

### NewConstraintViolation

`func NewConstraintViolation() *ConstraintViolation`

NewConstraintViolation instantiates a new ConstraintViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstraintViolationWithDefaults

`func NewConstraintViolationWithDefaults() *ConstraintViolation`

NewConstraintViolationWithDefaults instantiates a new ConstraintViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *ConstraintViolation) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ConstraintViolation) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ConstraintViolation) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ConstraintViolation) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMessage

`func (o *ConstraintViolation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConstraintViolation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConstraintViolation) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConstraintViolation) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPath

`func (o *ConstraintViolation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ConstraintViolation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ConstraintViolation) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ConstraintViolation) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetParameterLocation

`func (o *ConstraintViolation) GetParameterLocation() string`

GetParameterLocation returns the ParameterLocation field if non-nil, zero value otherwise.

### GetParameterLocationOk

`func (o *ConstraintViolation) GetParameterLocationOk() (*string, bool)`

GetParameterLocationOk returns a tuple with the ParameterLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterLocation

`func (o *ConstraintViolation) SetParameterLocation(v string)`

SetParameterLocation sets ParameterLocation field to given value.

### HasParameterLocation

`func (o *ConstraintViolation) HasParameterLocation() bool`

HasParameterLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


