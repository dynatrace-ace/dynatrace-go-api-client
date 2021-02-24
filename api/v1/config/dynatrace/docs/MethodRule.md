# MethodRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the method rule. | [optional] 
**MethodName** | **string** | The method to instrument. | 
**ArgumentTypes** | Pointer to **[]string** | Fully qualified types of argument the method expects. | [optional] 
**ReturnType** | **string** | Fully qualified type the method returns. | 

## Methods

### NewMethodRule

`func NewMethodRule(methodName string, returnType string, ) *MethodRule`

NewMethodRule instantiates a new MethodRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMethodRuleWithDefaults

`func NewMethodRuleWithDefaults() *MethodRule`

NewMethodRuleWithDefaults instantiates a new MethodRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MethodRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MethodRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MethodRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MethodRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethodName

`func (o *MethodRule) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *MethodRule) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *MethodRule) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.


### GetArgumentTypes

`func (o *MethodRule) GetArgumentTypes() []string`

GetArgumentTypes returns the ArgumentTypes field if non-nil, zero value otherwise.

### GetArgumentTypesOk

`func (o *MethodRule) GetArgumentTypesOk() (*[]string, bool)`

GetArgumentTypesOk returns a tuple with the ArgumentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgumentTypes

`func (o *MethodRule) SetArgumentTypes(v []string)`

SetArgumentTypes sets ArgumentTypes field to given value.

### HasArgumentTypes

`func (o *MethodRule) HasArgumentTypes() bool`

HasArgumentTypes returns a boolean if a field has been set.

### GetReturnType

`func (o *MethodRule) GetReturnType() string`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *MethodRule) GetReturnTypeOk() (*string, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *MethodRule) SetReturnType(v string)`

SetReturnType sets ReturnType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


