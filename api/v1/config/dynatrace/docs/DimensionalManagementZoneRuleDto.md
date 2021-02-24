# DimensionalManagementZoneRuleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether or not this rule is enabled. | 
**AppliesTo** | **string** | Target of the rule. | 
**Conditions** | [**[]DimensionalManagementZoneConditionDto**](DimensionalManagementZoneConditionDto.md) | List of the conditions that have to match. | 

## Methods

### NewDimensionalManagementZoneRuleDto

`func NewDimensionalManagementZoneRuleDto(enabled bool, appliesTo string, conditions []DimensionalManagementZoneConditionDto, ) *DimensionalManagementZoneRuleDto`

NewDimensionalManagementZoneRuleDto instantiates a new DimensionalManagementZoneRuleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDimensionalManagementZoneRuleDtoWithDefaults

`func NewDimensionalManagementZoneRuleDtoWithDefaults() *DimensionalManagementZoneRuleDto`

NewDimensionalManagementZoneRuleDtoWithDefaults instantiates a new DimensionalManagementZoneRuleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DimensionalManagementZoneRuleDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DimensionalManagementZoneRuleDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DimensionalManagementZoneRuleDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAppliesTo

`func (o *DimensionalManagementZoneRuleDto) GetAppliesTo() string`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *DimensionalManagementZoneRuleDto) GetAppliesToOk() (*string, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *DimensionalManagementZoneRuleDto) SetAppliesTo(v string)`

SetAppliesTo sets AppliesTo field to given value.


### GetConditions

`func (o *DimensionalManagementZoneRuleDto) GetConditions() []DimensionalManagementZoneConditionDto`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *DimensionalManagementZoneRuleDto) GetConditionsOk() (*[]DimensionalManagementZoneConditionDto, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *DimensionalManagementZoneRuleDto) SetConditions(v []DimensionalManagementZoneConditionDto)`

SetConditions sets Conditions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


