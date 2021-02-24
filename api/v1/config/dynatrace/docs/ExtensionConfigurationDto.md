# ExtensionConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionId** | Pointer to **string** | The ID of the extension. | [optional] 
**Enabled** | **bool** | The extension is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**UseGlobal** | **bool** | Allows to skip current configuration and use global one. | 
**Properties** | Pointer to **map[string]string** | The list of extension parameters.    Each parameter is a key-value pair. | [optional] 
**HostId** | Pointer to **string** | The ID of the host on which the extension runs. | [optional] 
**ActiveGate** | Pointer to [**EntityShortRepresentation**](EntityShortRepresentation.md) |  | [optional] 
**EndpointId** | Pointer to **string** | The ID of the endpoint. | [optional] 
**EndpointName** | Pointer to **string** | The name of the endpoint, displayed in Dynatrace. | [optional] 

## Methods

### NewExtensionConfigurationDto

`func NewExtensionConfigurationDto(enabled bool, useGlobal bool, ) *ExtensionConfigurationDto`

NewExtensionConfigurationDto instantiates a new ExtensionConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionConfigurationDtoWithDefaults

`func NewExtensionConfigurationDtoWithDefaults() *ExtensionConfigurationDto`

NewExtensionConfigurationDtoWithDefaults instantiates a new ExtensionConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionId

`func (o *ExtensionConfigurationDto) GetExtensionId() string`

GetExtensionId returns the ExtensionId field if non-nil, zero value otherwise.

### GetExtensionIdOk

`func (o *ExtensionConfigurationDto) GetExtensionIdOk() (*string, bool)`

GetExtensionIdOk returns a tuple with the ExtensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionId

`func (o *ExtensionConfigurationDto) SetExtensionId(v string)`

SetExtensionId sets ExtensionId field to given value.

### HasExtensionId

`func (o *ExtensionConfigurationDto) HasExtensionId() bool`

HasExtensionId returns a boolean if a field has been set.

### GetEnabled

`func (o *ExtensionConfigurationDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExtensionConfigurationDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExtensionConfigurationDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetUseGlobal

`func (o *ExtensionConfigurationDto) GetUseGlobal() bool`

GetUseGlobal returns the UseGlobal field if non-nil, zero value otherwise.

### GetUseGlobalOk

`func (o *ExtensionConfigurationDto) GetUseGlobalOk() (*bool, bool)`

GetUseGlobalOk returns a tuple with the UseGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGlobal

`func (o *ExtensionConfigurationDto) SetUseGlobal(v bool)`

SetUseGlobal sets UseGlobal field to given value.


### GetProperties

`func (o *ExtensionConfigurationDto) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ExtensionConfigurationDto) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ExtensionConfigurationDto) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ExtensionConfigurationDto) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetHostId

`func (o *ExtensionConfigurationDto) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ExtensionConfigurationDto) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ExtensionConfigurationDto) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *ExtensionConfigurationDto) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetActiveGate

`func (o *ExtensionConfigurationDto) GetActiveGate() EntityShortRepresentation`

GetActiveGate returns the ActiveGate field if non-nil, zero value otherwise.

### GetActiveGateOk

`func (o *ExtensionConfigurationDto) GetActiveGateOk() (*EntityShortRepresentation, bool)`

GetActiveGateOk returns a tuple with the ActiveGate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGate

`func (o *ExtensionConfigurationDto) SetActiveGate(v EntityShortRepresentation)`

SetActiveGate sets ActiveGate field to given value.

### HasActiveGate

`func (o *ExtensionConfigurationDto) HasActiveGate() bool`

HasActiveGate returns a boolean if a field has been set.

### GetEndpointId

`func (o *ExtensionConfigurationDto) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *ExtensionConfigurationDto) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *ExtensionConfigurationDto) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *ExtensionConfigurationDto) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetEndpointName

`func (o *ExtensionConfigurationDto) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *ExtensionConfigurationDto) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *ExtensionConfigurationDto) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.

### HasEndpointName

`func (o *ExtensionConfigurationDto) HasEndpointName() bool`

HasEndpointName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


