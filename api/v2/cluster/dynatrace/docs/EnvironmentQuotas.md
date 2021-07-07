# EnvironmentQuotas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostUnits** | Pointer to [**HostUnitQuota**](HostUnitQuota.md) |  | [optional] 
**DemUnits** | Pointer to [**DemUnitsQuota**](DemUnitsQuota.md) |  | [optional] 
**UserSessions** | Pointer to [**UserSessionsQuota**](UserSessionsQuota.md) |  | [optional] 
**SessionProperties** | Pointer to [**SessionPropertiesQuota**](SessionPropertiesQuota.md) |  | [optional] 
**SyntheticMonitors** | Pointer to [**SyntheticQuota**](SyntheticQuota.md) |  | [optional] 
**CustomMetrics** | Pointer to [**CustomMetricsQuota**](CustomMetricsQuota.md) |  | [optional] 
**DavisDataUnits** | Pointer to [**DavisDataUnitsQuota**](DavisDataUnitsQuota.md) |  | [optional] 
**LogMonitoring** | Pointer to [**LogMonitoringQuota**](LogMonitoringQuota.md) |  | [optional] 

## Methods

### NewEnvironmentQuotas

`func NewEnvironmentQuotas() *EnvironmentQuotas`

NewEnvironmentQuotas instantiates a new EnvironmentQuotas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentQuotasWithDefaults

`func NewEnvironmentQuotasWithDefaults() *EnvironmentQuotas`

NewEnvironmentQuotasWithDefaults instantiates a new EnvironmentQuotas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostUnits

`func (o *EnvironmentQuotas) GetHostUnits() HostUnitQuota`

GetHostUnits returns the HostUnits field if non-nil, zero value otherwise.

### GetHostUnitsOk

`func (o *EnvironmentQuotas) GetHostUnitsOk() (*HostUnitQuota, bool)`

GetHostUnitsOk returns a tuple with the HostUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUnits

`func (o *EnvironmentQuotas) SetHostUnits(v HostUnitQuota)`

SetHostUnits sets HostUnits field to given value.

### HasHostUnits

`func (o *EnvironmentQuotas) HasHostUnits() bool`

HasHostUnits returns a boolean if a field has been set.

### GetDemUnits

`func (o *EnvironmentQuotas) GetDemUnits() DemUnitsQuota`

GetDemUnits returns the DemUnits field if non-nil, zero value otherwise.

### GetDemUnitsOk

`func (o *EnvironmentQuotas) GetDemUnitsOk() (*DemUnitsQuota, bool)`

GetDemUnitsOk returns a tuple with the DemUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemUnits

`func (o *EnvironmentQuotas) SetDemUnits(v DemUnitsQuota)`

SetDemUnits sets DemUnits field to given value.

### HasDemUnits

`func (o *EnvironmentQuotas) HasDemUnits() bool`

HasDemUnits returns a boolean if a field has been set.

### GetUserSessions

`func (o *EnvironmentQuotas) GetUserSessions() UserSessionsQuota`

GetUserSessions returns the UserSessions field if non-nil, zero value otherwise.

### GetUserSessionsOk

`func (o *EnvironmentQuotas) GetUserSessionsOk() (*UserSessionsQuota, bool)`

GetUserSessionsOk returns a tuple with the UserSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSessions

`func (o *EnvironmentQuotas) SetUserSessions(v UserSessionsQuota)`

SetUserSessions sets UserSessions field to given value.

### HasUserSessions

`func (o *EnvironmentQuotas) HasUserSessions() bool`

HasUserSessions returns a boolean if a field has been set.

### GetSessionProperties

`func (o *EnvironmentQuotas) GetSessionProperties() SessionPropertiesQuota`

GetSessionProperties returns the SessionProperties field if non-nil, zero value otherwise.

### GetSessionPropertiesOk

`func (o *EnvironmentQuotas) GetSessionPropertiesOk() (*SessionPropertiesQuota, bool)`

GetSessionPropertiesOk returns a tuple with the SessionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionProperties

`func (o *EnvironmentQuotas) SetSessionProperties(v SessionPropertiesQuota)`

SetSessionProperties sets SessionProperties field to given value.

### HasSessionProperties

`func (o *EnvironmentQuotas) HasSessionProperties() bool`

HasSessionProperties returns a boolean if a field has been set.

### GetSyntheticMonitors

`func (o *EnvironmentQuotas) GetSyntheticMonitors() SyntheticQuota`

GetSyntheticMonitors returns the SyntheticMonitors field if non-nil, zero value otherwise.

### GetSyntheticMonitorsOk

`func (o *EnvironmentQuotas) GetSyntheticMonitorsOk() (*SyntheticQuota, bool)`

GetSyntheticMonitorsOk returns a tuple with the SyntheticMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticMonitors

`func (o *EnvironmentQuotas) SetSyntheticMonitors(v SyntheticQuota)`

SetSyntheticMonitors sets SyntheticMonitors field to given value.

### HasSyntheticMonitors

`func (o *EnvironmentQuotas) HasSyntheticMonitors() bool`

HasSyntheticMonitors returns a boolean if a field has been set.

### GetCustomMetrics

`func (o *EnvironmentQuotas) GetCustomMetrics() CustomMetricsQuota`

GetCustomMetrics returns the CustomMetrics field if non-nil, zero value otherwise.

### GetCustomMetricsOk

`func (o *EnvironmentQuotas) GetCustomMetricsOk() (*CustomMetricsQuota, bool)`

GetCustomMetricsOk returns a tuple with the CustomMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetrics

`func (o *EnvironmentQuotas) SetCustomMetrics(v CustomMetricsQuota)`

SetCustomMetrics sets CustomMetrics field to given value.

### HasCustomMetrics

`func (o *EnvironmentQuotas) HasCustomMetrics() bool`

HasCustomMetrics returns a boolean if a field has been set.

### GetDavisDataUnits

`func (o *EnvironmentQuotas) GetDavisDataUnits() DavisDataUnitsQuota`

GetDavisDataUnits returns the DavisDataUnits field if non-nil, zero value otherwise.

### GetDavisDataUnitsOk

`func (o *EnvironmentQuotas) GetDavisDataUnitsOk() (*DavisDataUnitsQuota, bool)`

GetDavisDataUnitsOk returns a tuple with the DavisDataUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDavisDataUnits

`func (o *EnvironmentQuotas) SetDavisDataUnits(v DavisDataUnitsQuota)`

SetDavisDataUnits sets DavisDataUnits field to given value.

### HasDavisDataUnits

`func (o *EnvironmentQuotas) HasDavisDataUnits() bool`

HasDavisDataUnits returns a boolean if a field has been set.

### GetLogMonitoring

`func (o *EnvironmentQuotas) GetLogMonitoring() LogMonitoringQuota`

GetLogMonitoring returns the LogMonitoring field if non-nil, zero value otherwise.

### GetLogMonitoringOk

`func (o *EnvironmentQuotas) GetLogMonitoringOk() (*LogMonitoringQuota, bool)`

GetLogMonitoringOk returns a tuple with the LogMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMonitoring

`func (o *EnvironmentQuotas) SetLogMonitoring(v LogMonitoringQuota)`

SetLogMonitoring sets LogMonitoring field to given value.

### HasLogMonitoring

`func (o *EnvironmentQuotas) HasLogMonitoring() bool`

HasLogMonitoring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


