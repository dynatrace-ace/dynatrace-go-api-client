# UpdateToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revoked** | Pointer to **bool** | The token is revoked (&#x60;true&#x60;) or active (&#x60;false&#x60;). | [optional] 
**Name** | Pointer to **string** | The name of the token. | [optional] 
**Scopes** | Pointer to **[]string** | The list of permissions assigned to the token.   Apart from the new permissions, you need to submit the existing permissions you want to keep, too. Any existing permission, missing in the payload, is revoked.  * &#x60;DiagnosticExport&#x60;: DiagnosticExport.  * &#x60;ControlManagement&#x60;: ControlManagement.  * &#x60;UnattendedInstall&#x60;: UnattendedInstall.  * &#x60;ServiceProviderAPI&#x60;: Service Provider API.  * &#x60;ExternalSyntheticIntegration&#x60;: Create and read synthetic monitors, locations, and nodes.  * &#x60;ClusterTokenManagement&#x60;: Cluster token management.  * &#x60;ReadSyntheticData&#x60;: Read synthetic monitors, locations, and nodes.  * &#x60;Nodekeeper&#x60;: Nodekeeper access for node management.  * &#x60;EnvironmentTokenManagement&#x60;: \&quot;Token Management\&quot; Token creation for existing environments.  * &#x60;settings.read&#x60;: Read settings.  * &#x60;settings.write&#x60;: Write settings.  * &#x60;apiTokens.read&#x60;: Read API tokens.  * &#x60;apiTokens.write&#x60;: Write API tokens.   | [optional] 

## Methods

### NewUpdateToken

`func NewUpdateToken() *UpdateToken`

NewUpdateToken instantiates a new UpdateToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTokenWithDefaults

`func NewUpdateTokenWithDefaults() *UpdateToken`

NewUpdateTokenWithDefaults instantiates a new UpdateToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevoked

`func (o *UpdateToken) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *UpdateToken) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *UpdateToken) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *UpdateToken) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetName

`func (o *UpdateToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScopes

`func (o *UpdateToken) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *UpdateToken) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *UpdateToken) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *UpdateToken) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


