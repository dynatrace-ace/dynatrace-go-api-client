# MonitoringConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | **string** | The scope this monitoring configuration will be defined for | 
**Value** | Pointer to **map[string]interface{}** | The monitoring configuration | [optional] 

## Methods

### NewMonitoringConfigurationDto

`func NewMonitoringConfigurationDto(scope string, ) *MonitoringConfigurationDto`

NewMonitoringConfigurationDto instantiates a new MonitoringConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringConfigurationDtoWithDefaults

`func NewMonitoringConfigurationDtoWithDefaults() *MonitoringConfigurationDto`

NewMonitoringConfigurationDtoWithDefaults instantiates a new MonitoringConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *MonitoringConfigurationDto) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MonitoringConfigurationDto) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MonitoringConfigurationDto) SetScope(v string)`

SetScope sets Scope field to given value.


### GetValue

`func (o *MonitoringConfigurationDto) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MonitoringConfigurationDto) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MonitoringConfigurationDto) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *MonitoringConfigurationDto) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


