# CreateEnvironmentTokenManagementToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the token. | 
**ExpiresIn** | Pointer to [**Duration**](Duration.md) |  | [optional] 

## Methods

### NewCreateEnvironmentTokenManagementToken

`func NewCreateEnvironmentTokenManagementToken(name string, ) *CreateEnvironmentTokenManagementToken`

NewCreateEnvironmentTokenManagementToken instantiates a new CreateEnvironmentTokenManagementToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEnvironmentTokenManagementTokenWithDefaults

`func NewCreateEnvironmentTokenManagementTokenWithDefaults() *CreateEnvironmentTokenManagementToken`

NewCreateEnvironmentTokenManagementTokenWithDefaults instantiates a new CreateEnvironmentTokenManagementToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEnvironmentTokenManagementToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEnvironmentTokenManagementToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEnvironmentTokenManagementToken) SetName(v string)`

SetName sets Name field to given value.


### GetExpiresIn

`func (o *CreateEnvironmentTokenManagementToken) GetExpiresIn() Duration`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CreateEnvironmentTokenManagementToken) GetExpiresInOk() (*Duration, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CreateEnvironmentTokenManagementToken) SetExpiresIn(v Duration)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *CreateEnvironmentTokenManagementToken) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


