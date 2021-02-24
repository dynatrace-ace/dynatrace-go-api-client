# IpAddressMappings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddressMappingRules** | Pointer to [**[]IpAddressMappingRule**](IpAddressMappingRule.md) | A list of IP address mapping rules.   Rules are evaluated from top to bottom; the first matching rule applies. | [optional] 

## Methods

### NewIpAddressMappings

`func NewIpAddressMappings() *IpAddressMappings`

NewIpAddressMappings instantiates a new IpAddressMappings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAddressMappingsWithDefaults

`func NewIpAddressMappingsWithDefaults() *IpAddressMappings`

NewIpAddressMappingsWithDefaults instantiates a new IpAddressMappings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddressMappingRules

`func (o *IpAddressMappings) GetIpAddressMappingRules() []IpAddressMappingRule`

GetIpAddressMappingRules returns the IpAddressMappingRules field if non-nil, zero value otherwise.

### GetIpAddressMappingRulesOk

`func (o *IpAddressMappings) GetIpAddressMappingRulesOk() (*[]IpAddressMappingRule, bool)`

GetIpAddressMappingRulesOk returns a tuple with the IpAddressMappingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressMappingRules

`func (o *IpAddressMappings) SetIpAddressMappingRules(v []IpAddressMappingRule)`

SetIpAddressMappingRules sets IpAddressMappingRules field to given value.

### HasIpAddressMappingRules

`func (o *IpAddressMappings) HasIpAddressMappingRules() bool`

HasIpAddressMappingRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


