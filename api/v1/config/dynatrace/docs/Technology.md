# Technology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the technology. | 
**MonitoringEnabled** | **bool** | The monitoring of the technology is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**Scope** | Pointer to **string** | The validity of the configuration:   * &#x60;HOST&#x60;: The setting is valid for OneAgent on host only. Other OneAgents, connected to the same Dynatrace server may have different setting.  * &#x60;ENVIRONMENT&#x60;: The setting is valid for all OneAgents, connected to the Dynatrace server. | [optional] 

## Methods

### NewTechnology

`func NewTechnology(type_ string, monitoringEnabled bool, ) *Technology`

NewTechnology instantiates a new Technology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechnologyWithDefaults

`func NewTechnologyWithDefaults() *Technology`

NewTechnologyWithDefaults instantiates a new Technology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Technology) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Technology) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Technology) SetType(v string)`

SetType sets Type field to given value.


### GetMonitoringEnabled

`func (o *Technology) GetMonitoringEnabled() bool`

GetMonitoringEnabled returns the MonitoringEnabled field if non-nil, zero value otherwise.

### GetMonitoringEnabledOk

`func (o *Technology) GetMonitoringEnabledOk() (*bool, bool)`

GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEnabled

`func (o *Technology) SetMonitoringEnabled(v bool)`

SetMonitoringEnabled sets MonitoringEnabled field to given value.


### GetScope

`func (o *Technology) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Technology) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Technology) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Technology) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


