# CalculatedServiceMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**TsmMetricKey** | **string** | The key of the calculated service metric. | 
**Name** | **string** | The displayed name of the metric. | 
**Enabled** | **bool** | The metric is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**MetricDefinition** | [**CalculatedMetricDefinition**](CalculatedMetricDefinition.md) |  | 
**Unit** | **string** | The unit of the metric. | 
**UnitDisplayName** | Pointer to **string** | The display name of the metric&#39;s unit.    Only applicable when the **unit** parameter is set to &#x60;UNSPECIFIED&#x60;. | [optional] 
**EntityId** | Pointer to **string** | Restricts the metric usage to the specified service.    This field is mutually exclusive with the **managementZones** field. | [optional] 
**ManagementZones** | Pointer to **[]string** | Restricts the metric usage to specified management zones.    This field is mutually exclusive with the **entityId** field. | [optional] 
**Conditions** | Pointer to [**[]Condition**](Condition.md) | The set of conditions for the metric usage.    **All** the specified conditions must be fulfilled to use the metric. | [optional] 
**DimensionDefinition** | Pointer to [**DimensionDefinition**](DimensionDefinition.md) |  | [optional] 

## Methods

### NewCalculatedServiceMetric

`func NewCalculatedServiceMetric(tsmMetricKey string, name string, enabled bool, metricDefinition CalculatedMetricDefinition, unit string, ) *CalculatedServiceMetric`

NewCalculatedServiceMetric instantiates a new CalculatedServiceMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalculatedServiceMetricWithDefaults

`func NewCalculatedServiceMetricWithDefaults() *CalculatedServiceMetric`

NewCalculatedServiceMetricWithDefaults instantiates a new CalculatedServiceMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CalculatedServiceMetric) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CalculatedServiceMetric) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CalculatedServiceMetric) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CalculatedServiceMetric) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTsmMetricKey

`func (o *CalculatedServiceMetric) GetTsmMetricKey() string`

GetTsmMetricKey returns the TsmMetricKey field if non-nil, zero value otherwise.

### GetTsmMetricKeyOk

`func (o *CalculatedServiceMetric) GetTsmMetricKeyOk() (*string, bool)`

GetTsmMetricKeyOk returns a tuple with the TsmMetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsmMetricKey

`func (o *CalculatedServiceMetric) SetTsmMetricKey(v string)`

SetTsmMetricKey sets TsmMetricKey field to given value.


### GetName

`func (o *CalculatedServiceMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CalculatedServiceMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CalculatedServiceMetric) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *CalculatedServiceMetric) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CalculatedServiceMetric) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CalculatedServiceMetric) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMetricDefinition

`func (o *CalculatedServiceMetric) GetMetricDefinition() CalculatedMetricDefinition`

GetMetricDefinition returns the MetricDefinition field if non-nil, zero value otherwise.

### GetMetricDefinitionOk

`func (o *CalculatedServiceMetric) GetMetricDefinitionOk() (*CalculatedMetricDefinition, bool)`

GetMetricDefinitionOk returns a tuple with the MetricDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDefinition

`func (o *CalculatedServiceMetric) SetMetricDefinition(v CalculatedMetricDefinition)`

SetMetricDefinition sets MetricDefinition field to given value.


### GetUnit

`func (o *CalculatedServiceMetric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CalculatedServiceMetric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CalculatedServiceMetric) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetUnitDisplayName

`func (o *CalculatedServiceMetric) GetUnitDisplayName() string`

GetUnitDisplayName returns the UnitDisplayName field if non-nil, zero value otherwise.

### GetUnitDisplayNameOk

`func (o *CalculatedServiceMetric) GetUnitDisplayNameOk() (*string, bool)`

GetUnitDisplayNameOk returns a tuple with the UnitDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitDisplayName

`func (o *CalculatedServiceMetric) SetUnitDisplayName(v string)`

SetUnitDisplayName sets UnitDisplayName field to given value.

### HasUnitDisplayName

`func (o *CalculatedServiceMetric) HasUnitDisplayName() bool`

HasUnitDisplayName returns a boolean if a field has been set.

### GetEntityId

`func (o *CalculatedServiceMetric) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *CalculatedServiceMetric) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *CalculatedServiceMetric) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *CalculatedServiceMetric) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetManagementZones

`func (o *CalculatedServiceMetric) GetManagementZones() []string`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *CalculatedServiceMetric) GetManagementZonesOk() (*[]string, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *CalculatedServiceMetric) SetManagementZones(v []string)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *CalculatedServiceMetric) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetConditions

`func (o *CalculatedServiceMetric) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CalculatedServiceMetric) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CalculatedServiceMetric) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CalculatedServiceMetric) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDimensionDefinition

`func (o *CalculatedServiceMetric) GetDimensionDefinition() DimensionDefinition`

GetDimensionDefinition returns the DimensionDefinition field if non-nil, zero value otherwise.

### GetDimensionDefinitionOk

`func (o *CalculatedServiceMetric) GetDimensionDefinitionOk() (*DimensionDefinition, bool)`

GetDimensionDefinitionOk returns a tuple with the DimensionDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionDefinition

`func (o *CalculatedServiceMetric) SetDimensionDefinition(v DimensionDefinition)`

SetDimensionDefinition sets DimensionDefinition field to given value.

### HasDimensionDefinition

`func (o *CalculatedServiceMetric) HasDimensionDefinition() bool`

HasDimensionDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


