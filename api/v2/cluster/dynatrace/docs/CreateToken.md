# CreateToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the token. | 
**ExpiresIn** | Pointer to [**Duration**](Duration.md) |  | [optional] 
**Scopes** | **[]string** | The list of scopes to be assigned to the token.  * &#x60;DiagnosticExport&#x60;: DiagnosticExport.  * &#x60;ControlManagement&#x60;: ControlManagement.  * &#x60;UnattendedInstall&#x60;: UnattendedInstall.  * &#x60;ServiceProviderAPI&#x60;: Service Provider API.  * &#x60;ExternalSyntheticIntegration&#x60;: Create and read synthetic monitors, locations, and nodes.  * &#x60;ClusterTokenManagement&#x60;: Cluster token management.  * &#x60;ReadSyntheticData&#x60;: Read synthetic monitors, locations, and nodes.  * &#x60;Nodekeeper&#x60;: Nodekeeper access for node management.  * &#x60;EnvironmentTokenManagement&#x60;: \&quot;Token Management\&quot; Token creation for existing environments.  * &#x60;settings.read&#x60;: Read settings.  * &#x60;settings.write&#x60;: Write settings.  * &#x60;apiTokens.read&#x60;: Read API tokens.  * &#x60;apiTokens.write&#x60;: Write API tokens.   | 

## Methods

### NewCreateToken

`func NewCreateToken(name string, scopes []string, ) *CreateToken`

NewCreateToken instantiates a new CreateToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokenWithDefaults

`func NewCreateTokenWithDefaults() *CreateToken`

NewCreateTokenWithDefaults instantiates a new CreateToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateToken) SetName(v string)`

SetName sets Name field to given value.


### GetExpiresIn

`func (o *CreateToken) GetExpiresIn() Duration`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CreateToken) GetExpiresInOk() (*Duration, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CreateToken) SetExpiresIn(v Duration)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *CreateToken) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetScopes

`func (o *CreateToken) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateToken) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateToken) SetScopes(v []string)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


