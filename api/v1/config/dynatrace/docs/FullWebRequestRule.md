# FullWebRequestRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the service detection rule. | 
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**ManagementZones** | Pointer to **[]string** | Specifies the management zones of the process group for which this service detection rule should be created. | [optional] 
**Id** | Pointer to **string** | The ID of the service detection rule. | [optional] 
**Order** | Pointer to **string** | The order of the rule in the rules list.    The rules are evaluated from top to bottom. The first matching rule applies. | [optional] 
**Name** | **string** | The name of the rule. | 
**Description** | Pointer to **string** | A short description of the rule. | [optional] 
**Enabled** | **bool** | The rule is enabled(&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**Conditions** | Pointer to [**[]ConditionsFullWebRequestAttributeTypeDto**](ConditionsFullWebRequestAttributeTypeDto.md) | A list of conditions of the rule.   If several conditions are specified, the AND logic applies. | [optional] 
**ApplicationId** | Pointer to [**ApplicationId**](ApplicationId.md) |  | [optional] 
**ContextRoot** | Pointer to [**ContextRoot**](ContextRoot.md) |  | [optional] 
**ServerName** | Pointer to [**ServerName**](ServerName.md) |  | [optional] 

## Methods

### NewFullWebRequestRule

`func NewFullWebRequestRule(type_ string, name string, enabled bool, ) *FullWebRequestRule`

NewFullWebRequestRule instantiates a new FullWebRequestRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullWebRequestRuleWithDefaults

`func NewFullWebRequestRuleWithDefaults() *FullWebRequestRule`

NewFullWebRequestRuleWithDefaults instantiates a new FullWebRequestRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FullWebRequestRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FullWebRequestRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FullWebRequestRule) SetType(v string)`

SetType sets Type field to given value.


### GetMetadata

`func (o *FullWebRequestRule) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FullWebRequestRule) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FullWebRequestRule) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FullWebRequestRule) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetManagementZones

`func (o *FullWebRequestRule) GetManagementZones() []string`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *FullWebRequestRule) GetManagementZonesOk() (*[]string, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *FullWebRequestRule) SetManagementZones(v []string)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *FullWebRequestRule) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetId

`func (o *FullWebRequestRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullWebRequestRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullWebRequestRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FullWebRequestRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrder

`func (o *FullWebRequestRule) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FullWebRequestRule) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FullWebRequestRule) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *FullWebRequestRule) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetName

`func (o *FullWebRequestRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullWebRequestRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullWebRequestRule) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FullWebRequestRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FullWebRequestRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FullWebRequestRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FullWebRequestRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *FullWebRequestRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FullWebRequestRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FullWebRequestRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetConditions

`func (o *FullWebRequestRule) GetConditions() []ConditionsFullWebRequestAttributeTypeDto`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *FullWebRequestRule) GetConditionsOk() (*[]ConditionsFullWebRequestAttributeTypeDto, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *FullWebRequestRule) SetConditions(v []ConditionsFullWebRequestAttributeTypeDto)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *FullWebRequestRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetApplicationId

`func (o *FullWebRequestRule) GetApplicationId() ApplicationId`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *FullWebRequestRule) GetApplicationIdOk() (*ApplicationId, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *FullWebRequestRule) SetApplicationId(v ApplicationId)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *FullWebRequestRule) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetContextRoot

`func (o *FullWebRequestRule) GetContextRoot() ContextRoot`

GetContextRoot returns the ContextRoot field if non-nil, zero value otherwise.

### GetContextRootOk

`func (o *FullWebRequestRule) GetContextRootOk() (*ContextRoot, bool)`

GetContextRootOk returns a tuple with the ContextRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextRoot

`func (o *FullWebRequestRule) SetContextRoot(v ContextRoot)`

SetContextRoot sets ContextRoot field to given value.

### HasContextRoot

`func (o *FullWebRequestRule) HasContextRoot() bool`

HasContextRoot returns a boolean if a field has been set.

### GetServerName

`func (o *FullWebRequestRule) GetServerName() ServerName`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *FullWebRequestRule) GetServerNameOk() (*ServerName, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *FullWebRequestRule) SetServerName(v ServerName)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *FullWebRequestRule) HasServerName() bool`

HasServerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


