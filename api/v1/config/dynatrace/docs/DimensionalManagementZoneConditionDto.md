# DimensionalManagementZoneConditionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionType** | **string** | The type of the condition. | 
**RuleMatcher** | **string** | In what way the actual values are compared to the expected ones. | 
**Key** | **string** | The main value that is compared. For dimensions, this is the key. | 
**Value** | Pointer to **string** | The dimension value. For types other than dimension this must be null. | [optional] 

## Methods

### NewDimensionalManagementZoneConditionDto

`func NewDimensionalManagementZoneConditionDto(conditionType string, ruleMatcher string, key string, ) *DimensionalManagementZoneConditionDto`

NewDimensionalManagementZoneConditionDto instantiates a new DimensionalManagementZoneConditionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDimensionalManagementZoneConditionDtoWithDefaults

`func NewDimensionalManagementZoneConditionDtoWithDefaults() *DimensionalManagementZoneConditionDto`

NewDimensionalManagementZoneConditionDtoWithDefaults instantiates a new DimensionalManagementZoneConditionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionType

`func (o *DimensionalManagementZoneConditionDto) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *DimensionalManagementZoneConditionDto) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *DimensionalManagementZoneConditionDto) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.


### GetRuleMatcher

`func (o *DimensionalManagementZoneConditionDto) GetRuleMatcher() string`

GetRuleMatcher returns the RuleMatcher field if non-nil, zero value otherwise.

### GetRuleMatcherOk

`func (o *DimensionalManagementZoneConditionDto) GetRuleMatcherOk() (*string, bool)`

GetRuleMatcherOk returns a tuple with the RuleMatcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleMatcher

`func (o *DimensionalManagementZoneConditionDto) SetRuleMatcher(v string)`

SetRuleMatcher sets RuleMatcher field to given value.


### GetKey

`func (o *DimensionalManagementZoneConditionDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DimensionalManagementZoneConditionDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DimensionalManagementZoneConditionDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *DimensionalManagementZoneConditionDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DimensionalManagementZoneConditionDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DimensionalManagementZoneConditionDto) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DimensionalManagementZoneConditionDto) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


