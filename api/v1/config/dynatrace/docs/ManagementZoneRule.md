# ManagementZoneRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of Dynatrace entities the management zone can be applied to. | 
**Enabled** | **bool** | The rule is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**PropagationTypes** | Pointer to **[]string** | How to apply the management zone to underlying entities:   * &#x60;SERVICE_TO_HOST_LIKE&#x60;: Apply to underlying hosts of matching services.  * &#x60;SERVICE_TO_PROCESS_GROUP_LIKE&#x60;: Apply to underlying process groups of matching services.  * &#x60;PROCESS_GROUP_TO_HOST&#x60;: Apply to underlying hosts of matching process groups.  * &#x60;PROCESS_GROUP_TO_SERVICE&#x60;: Apply to all services provided by matching process groups.  * &#x60;HOST_TO_PROCESS_GROUP_INSTANCE&#x60;: Apply to processes running on matching hosts.  * &#x60;CUSTOM_DEVICE_GROUP_TO_CUSTOM_DEVICE&#x60;: Apply to custom devices in matching custom device groups.  * &#x60;AZURE_TO_PG&#x60;: Apply to process groups connected to matching Azure entities.  * &#x60;AZURE_TO_SERVICE&#x60;: Apply to services provided by matching Azure entities. | [optional] 
**Conditions** | [**[]EntityRuleEngineCondition**](EntityRuleEngineCondition.md) | A list of matching rules for the management zone.   The management zone applies only if **all** conditions are fulfilled. | 

## Methods

### NewManagementZoneRule

`func NewManagementZoneRule(type_ string, enabled bool, conditions []EntityRuleEngineCondition, ) *ManagementZoneRule`

NewManagementZoneRule instantiates a new ManagementZoneRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementZoneRuleWithDefaults

`func NewManagementZoneRuleWithDefaults() *ManagementZoneRule`

NewManagementZoneRuleWithDefaults instantiates a new ManagementZoneRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManagementZoneRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagementZoneRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagementZoneRule) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *ManagementZoneRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ManagementZoneRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ManagementZoneRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPropagationTypes

`func (o *ManagementZoneRule) GetPropagationTypes() []string`

GetPropagationTypes returns the PropagationTypes field if non-nil, zero value otherwise.

### GetPropagationTypesOk

`func (o *ManagementZoneRule) GetPropagationTypesOk() (*[]string, bool)`

GetPropagationTypesOk returns a tuple with the PropagationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationTypes

`func (o *ManagementZoneRule) SetPropagationTypes(v []string)`

SetPropagationTypes sets PropagationTypes field to given value.

### HasPropagationTypes

`func (o *ManagementZoneRule) HasPropagationTypes() bool`

HasPropagationTypes returns a boolean if a field has been set.

### GetConditions

`func (o *ManagementZoneRule) GetConditions() []EntityRuleEngineCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ManagementZoneRule) GetConditionsOk() (*[]EntityRuleEngineCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ManagementZoneRule) SetConditions(v []EntityRuleEngineCondition)`

SetConditions sets Conditions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


