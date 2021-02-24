# GlobalExtensionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionId** | Pointer to **string** | The ID of the extension. | [optional] 
**Enabled** | **bool** | The extension is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**InfraOnlyEnabled** | Pointer to **bool** | The plugin is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;) globally for hosts in infrastructure-only monitoring mode | [optional] 
**Properties** | Pointer to **map[string]string** | The list of configuration parameters.    Each parameter is a key-value pair. | [optional] 

## Methods

### NewGlobalExtensionConfiguration

`func NewGlobalExtensionConfiguration(enabled bool, ) *GlobalExtensionConfiguration`

NewGlobalExtensionConfiguration instantiates a new GlobalExtensionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalExtensionConfigurationWithDefaults

`func NewGlobalExtensionConfigurationWithDefaults() *GlobalExtensionConfiguration`

NewGlobalExtensionConfigurationWithDefaults instantiates a new GlobalExtensionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionId

`func (o *GlobalExtensionConfiguration) GetExtensionId() string`

GetExtensionId returns the ExtensionId field if non-nil, zero value otherwise.

### GetExtensionIdOk

`func (o *GlobalExtensionConfiguration) GetExtensionIdOk() (*string, bool)`

GetExtensionIdOk returns a tuple with the ExtensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionId

`func (o *GlobalExtensionConfiguration) SetExtensionId(v string)`

SetExtensionId sets ExtensionId field to given value.

### HasExtensionId

`func (o *GlobalExtensionConfiguration) HasExtensionId() bool`

HasExtensionId returns a boolean if a field has been set.

### GetEnabled

`func (o *GlobalExtensionConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GlobalExtensionConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GlobalExtensionConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInfraOnlyEnabled

`func (o *GlobalExtensionConfiguration) GetInfraOnlyEnabled() bool`

GetInfraOnlyEnabled returns the InfraOnlyEnabled field if non-nil, zero value otherwise.

### GetInfraOnlyEnabledOk

`func (o *GlobalExtensionConfiguration) GetInfraOnlyEnabledOk() (*bool, bool)`

GetInfraOnlyEnabledOk returns a tuple with the InfraOnlyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraOnlyEnabled

`func (o *GlobalExtensionConfiguration) SetInfraOnlyEnabled(v bool)`

SetInfraOnlyEnabled sets InfraOnlyEnabled field to given value.

### HasInfraOnlyEnabled

`func (o *GlobalExtensionConfiguration) HasInfraOnlyEnabled() bool`

HasInfraOnlyEnabled returns a boolean if a field has been set.

### GetProperties

`func (o *GlobalExtensionConfiguration) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GlobalExtensionConfiguration) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GlobalExtensionConfiguration) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *GlobalExtensionConfiguration) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


