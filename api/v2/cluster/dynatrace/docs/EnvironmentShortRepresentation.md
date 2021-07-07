# EnvironmentShortRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the Dynatrace entity. | 
**Name** | Pointer to **string** | The name of the Dynatrace entity. | [optional] [readonly] 
**Description** | Pointer to **string** | A short description of the Dynatrace entity. | [optional] [readonly] 
**TokenManagementToken** | Pointer to **string** | A token with the &#39;Token management&#39; permission. Can be used to within the newly created environment to create other tokens for configuring this environment. This value is only set if an environment was created with the query parameter &#39;createToken&#x3D;true&#39;. | [optional] [readonly] 

## Methods

### NewEnvironmentShortRepresentation

`func NewEnvironmentShortRepresentation(id string, ) *EnvironmentShortRepresentation`

NewEnvironmentShortRepresentation instantiates a new EnvironmentShortRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentShortRepresentationWithDefaults

`func NewEnvironmentShortRepresentationWithDefaults() *EnvironmentShortRepresentation`

NewEnvironmentShortRepresentationWithDefaults instantiates a new EnvironmentShortRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentShortRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentShortRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentShortRepresentation) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EnvironmentShortRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentShortRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentShortRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentShortRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EnvironmentShortRepresentation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentShortRepresentation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentShortRepresentation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentShortRepresentation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTokenManagementToken

`func (o *EnvironmentShortRepresentation) GetTokenManagementToken() string`

GetTokenManagementToken returns the TokenManagementToken field if non-nil, zero value otherwise.

### GetTokenManagementTokenOk

`func (o *EnvironmentShortRepresentation) GetTokenManagementTokenOk() (*string, bool)`

GetTokenManagementTokenOk returns a tuple with the TokenManagementToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenManagementToken

`func (o *EnvironmentShortRepresentation) SetTokenManagementToken(v string)`

SetTokenManagementToken sets TokenManagementToken field to given value.

### HasTokenManagementToken

`func (o *EnvironmentShortRepresentation) HasTokenManagementToken() bool`

HasTokenManagementToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


