# RemotePluginEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the endpoint. | [optional] 
**PluginId** | Pointer to **string** | The ID of the plugin to which the endpoint belongs. | [optional] 
**Name** | Pointer to **string** | The name of the endpoint, displayed in Dynatrace. | [optional] 
**Enabled** | Pointer to **bool** | The endpoint is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 
**Properties** | Pointer to **map[string]string** | The list of endpoint parameters.    Each parameter is a property-value pair. | [optional] 
**ActiveGatePluginModule** | [**EntityShortRepresentation**](EntityShortRepresentation.md) |  | 

## Methods

### NewRemotePluginEndpoint

`func NewRemotePluginEndpoint(activeGatePluginModule EntityShortRepresentation, ) *RemotePluginEndpoint`

NewRemotePluginEndpoint instantiates a new RemotePluginEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemotePluginEndpointWithDefaults

`func NewRemotePluginEndpointWithDefaults() *RemotePluginEndpoint`

NewRemotePluginEndpointWithDefaults instantiates a new RemotePluginEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemotePluginEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemotePluginEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemotePluginEndpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemotePluginEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPluginId

`func (o *RemotePluginEndpoint) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *RemotePluginEndpoint) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *RemotePluginEndpoint) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.

### HasPluginId

`func (o *RemotePluginEndpoint) HasPluginId() bool`

HasPluginId returns a boolean if a field has been set.

### GetName

`func (o *RemotePluginEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemotePluginEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemotePluginEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemotePluginEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *RemotePluginEndpoint) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RemotePluginEndpoint) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RemotePluginEndpoint) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RemotePluginEndpoint) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProperties

`func (o *RemotePluginEndpoint) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RemotePluginEndpoint) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RemotePluginEndpoint) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *RemotePluginEndpoint) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetActiveGatePluginModule

`func (o *RemotePluginEndpoint) GetActiveGatePluginModule() EntityShortRepresentation`

GetActiveGatePluginModule returns the ActiveGatePluginModule field if non-nil, zero value otherwise.

### GetActiveGatePluginModuleOk

`func (o *RemotePluginEndpoint) GetActiveGatePluginModuleOk() (*EntityShortRepresentation, bool)`

GetActiveGatePluginModuleOk returns a tuple with the ActiveGatePluginModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGatePluginModule

`func (o *RemotePluginEndpoint) SetActiveGatePluginModule(v EntityShortRepresentation)`

SetActiveGatePluginModule sets ActiveGatePluginModule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


