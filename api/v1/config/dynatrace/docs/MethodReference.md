# MethodReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visibility** | **string** | The visibility of the method to capture. | 
**Modifiers** | **[]string** | The modifiers of the method to capture. | 
**ClassName** | Pointer to **string** | The class name where the method to capture resides.    Either this or the **fileName** must be set. | [optional] 
**FileName** | Pointer to **string** | The file name where the method to capture resides.    Either this or **className** must be set. | [optional] 
**FileNameMatcher** | Pointer to **string** | The operator of the comparison.    If not set, &#x60;EQUALS&#x60; is used. | [optional] 
**MethodName** | **string** | The name of the method to capture. | 
**ArgumentTypes** | **[]string** | The list of argument types. | 
**ReturnType** | **string** | The return type. | 

## Methods

### NewMethodReference

`func NewMethodReference(visibility string, modifiers []string, methodName string, argumentTypes []string, returnType string, ) *MethodReference`

NewMethodReference instantiates a new MethodReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMethodReferenceWithDefaults

`func NewMethodReferenceWithDefaults() *MethodReference`

NewMethodReferenceWithDefaults instantiates a new MethodReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisibility

`func (o *MethodReference) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *MethodReference) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *MethodReference) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetModifiers

`func (o *MethodReference) GetModifiers() []string`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *MethodReference) GetModifiersOk() (*[]string, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *MethodReference) SetModifiers(v []string)`

SetModifiers sets Modifiers field to given value.


### GetClassName

`func (o *MethodReference) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *MethodReference) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *MethodReference) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *MethodReference) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetFileName

`func (o *MethodReference) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MethodReference) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *MethodReference) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *MethodReference) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileNameMatcher

`func (o *MethodReference) GetFileNameMatcher() string`

GetFileNameMatcher returns the FileNameMatcher field if non-nil, zero value otherwise.

### GetFileNameMatcherOk

`func (o *MethodReference) GetFileNameMatcherOk() (*string, bool)`

GetFileNameMatcherOk returns a tuple with the FileNameMatcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNameMatcher

`func (o *MethodReference) SetFileNameMatcher(v string)`

SetFileNameMatcher sets FileNameMatcher field to given value.

### HasFileNameMatcher

`func (o *MethodReference) HasFileNameMatcher() bool`

HasFileNameMatcher returns a boolean if a field has been set.

### GetMethodName

`func (o *MethodReference) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *MethodReference) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *MethodReference) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.


### GetArgumentTypes

`func (o *MethodReference) GetArgumentTypes() []string`

GetArgumentTypes returns the ArgumentTypes field if non-nil, zero value otherwise.

### GetArgumentTypesOk

`func (o *MethodReference) GetArgumentTypesOk() (*[]string, bool)`

GetArgumentTypesOk returns a tuple with the ArgumentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgumentTypes

`func (o *MethodReference) SetArgumentTypes(v []string)`

SetArgumentTypes sets ArgumentTypes field to given value.


### GetReturnType

`func (o *MethodReference) GetReturnType() string`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *MethodReference) GetReturnTypeOk() (*string, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *MethodReference) SetReturnType(v string)`

SetReturnType sets ReturnType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


