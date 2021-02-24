# ManagementZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the management zone. | [optional] 
**Name** | **string** | The name of the management zone. | 
**Rules** | Pointer to [**[]ManagementZoneRule**](ManagementZoneRule.md) | A list of rules for management zone usage.   Each rule is evaluated independently of all other rules. | [optional] 
**DimensionalRules** | Pointer to [**[]DimensionalManagementZoneRuleDto**](DimensionalManagementZoneRuleDto.md) | A list of dimensional data rules for management zone usage.   Each rule is evaluated independently of all other rules. | [optional] 

## Methods

### NewManagementZone

`func NewManagementZone(name string, ) *ManagementZone`

NewManagementZone instantiates a new ManagementZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementZoneWithDefaults

`func NewManagementZoneWithDefaults() *ManagementZone`

NewManagementZoneWithDefaults instantiates a new ManagementZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ManagementZone) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ManagementZone) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ManagementZone) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ManagementZone) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *ManagementZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagementZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagementZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ManagementZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ManagementZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementZone) SetName(v string)`

SetName sets Name field to given value.


### GetRules

`func (o *ManagementZone) GetRules() []ManagementZoneRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ManagementZone) GetRulesOk() (*[]ManagementZoneRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ManagementZone) SetRules(v []ManagementZoneRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *ManagementZone) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetDimensionalRules

`func (o *ManagementZone) GetDimensionalRules() []DimensionalManagementZoneRuleDto`

GetDimensionalRules returns the DimensionalRules field if non-nil, zero value otherwise.

### GetDimensionalRulesOk

`func (o *ManagementZone) GetDimensionalRulesOk() (*[]DimensionalManagementZoneRuleDto, bool)`

GetDimensionalRulesOk returns a tuple with the DimensionalRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionalRules

`func (o *ManagementZone) SetDimensionalRules(v []DimensionalManagementZoneRuleDto)`

SetDimensionalRules sets DimensionalRules field to given value.

### HasDimensionalRules

`func (o *ManagementZone) HasDimensionalRules() bool`

HasDimensionalRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


