# OpaqueAndExternalWebRequestRule

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
**Conditions** | Pointer to [**[]ConditionsOpaqueAndExternalWebRequestAttributeTypeDto**](ConditionsOpaqueAndExternalWebRequestAttributeTypeDto.md) | A list of conditions of the rule.   If several conditions are specified, the AND logic applies. | [optional] 
**ApplicationId** | Pointer to [**ApplicationId**](ApplicationId.md) |  | [optional] 
**ContextRoot** | Pointer to [**ContextRoot**](ContextRoot.md) |  | [optional] 
**Port** | Pointer to [**Port**](Port.md) |  | [optional] 
**PublicDomainName** | Pointer to [**PublicDomainName**](PublicDomainName.md) |  | [optional] 

## Methods

### NewOpaqueAndExternalWebRequestRule

`func NewOpaqueAndExternalWebRequestRule(type_ string, name string, enabled bool, ) *OpaqueAndExternalWebRequestRule`

NewOpaqueAndExternalWebRequestRule instantiates a new OpaqueAndExternalWebRequestRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpaqueAndExternalWebRequestRuleWithDefaults

`func NewOpaqueAndExternalWebRequestRuleWithDefaults() *OpaqueAndExternalWebRequestRule`

NewOpaqueAndExternalWebRequestRuleWithDefaults instantiates a new OpaqueAndExternalWebRequestRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OpaqueAndExternalWebRequestRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpaqueAndExternalWebRequestRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpaqueAndExternalWebRequestRule) SetType(v string)`

SetType sets Type field to given value.


### GetMetadata

`func (o *OpaqueAndExternalWebRequestRule) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OpaqueAndExternalWebRequestRule) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OpaqueAndExternalWebRequestRule) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OpaqueAndExternalWebRequestRule) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetManagementZones

`func (o *OpaqueAndExternalWebRequestRule) GetManagementZones() []string`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *OpaqueAndExternalWebRequestRule) GetManagementZonesOk() (*[]string, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *OpaqueAndExternalWebRequestRule) SetManagementZones(v []string)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *OpaqueAndExternalWebRequestRule) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetId

`func (o *OpaqueAndExternalWebRequestRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpaqueAndExternalWebRequestRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpaqueAndExternalWebRequestRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OpaqueAndExternalWebRequestRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrder

`func (o *OpaqueAndExternalWebRequestRule) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OpaqueAndExternalWebRequestRule) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OpaqueAndExternalWebRequestRule) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *OpaqueAndExternalWebRequestRule) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetName

`func (o *OpaqueAndExternalWebRequestRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpaqueAndExternalWebRequestRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpaqueAndExternalWebRequestRule) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OpaqueAndExternalWebRequestRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpaqueAndExternalWebRequestRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpaqueAndExternalWebRequestRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpaqueAndExternalWebRequestRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *OpaqueAndExternalWebRequestRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OpaqueAndExternalWebRequestRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OpaqueAndExternalWebRequestRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetConditions

`func (o *OpaqueAndExternalWebRequestRule) GetConditions() []ConditionsOpaqueAndExternalWebRequestAttributeTypeDto`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *OpaqueAndExternalWebRequestRule) GetConditionsOk() (*[]ConditionsOpaqueAndExternalWebRequestAttributeTypeDto, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *OpaqueAndExternalWebRequestRule) SetConditions(v []ConditionsOpaqueAndExternalWebRequestAttributeTypeDto)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *OpaqueAndExternalWebRequestRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetApplicationId

`func (o *OpaqueAndExternalWebRequestRule) GetApplicationId() ApplicationId`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *OpaqueAndExternalWebRequestRule) GetApplicationIdOk() (*ApplicationId, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *OpaqueAndExternalWebRequestRule) SetApplicationId(v ApplicationId)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *OpaqueAndExternalWebRequestRule) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetContextRoot

`func (o *OpaqueAndExternalWebRequestRule) GetContextRoot() ContextRoot`

GetContextRoot returns the ContextRoot field if non-nil, zero value otherwise.

### GetContextRootOk

`func (o *OpaqueAndExternalWebRequestRule) GetContextRootOk() (*ContextRoot, bool)`

GetContextRootOk returns a tuple with the ContextRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextRoot

`func (o *OpaqueAndExternalWebRequestRule) SetContextRoot(v ContextRoot)`

SetContextRoot sets ContextRoot field to given value.

### HasContextRoot

`func (o *OpaqueAndExternalWebRequestRule) HasContextRoot() bool`

HasContextRoot returns a boolean if a field has been set.

### GetPort

`func (o *OpaqueAndExternalWebRequestRule) GetPort() Port`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OpaqueAndExternalWebRequestRule) GetPortOk() (*Port, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OpaqueAndExternalWebRequestRule) SetPort(v Port)`

SetPort sets Port field to given value.

### HasPort

`func (o *OpaqueAndExternalWebRequestRule) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPublicDomainName

`func (o *OpaqueAndExternalWebRequestRule) GetPublicDomainName() PublicDomainName`

GetPublicDomainName returns the PublicDomainName field if non-nil, zero value otherwise.

### GetPublicDomainNameOk

`func (o *OpaqueAndExternalWebRequestRule) GetPublicDomainNameOk() (*PublicDomainName, bool)`

GetPublicDomainNameOk returns a tuple with the PublicDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDomainName

`func (o *OpaqueAndExternalWebRequestRule) SetPublicDomainName(v PublicDomainName)`

SetPublicDomainName sets PublicDomainName field to given value.

### HasPublicDomainName

`func (o *OpaqueAndExternalWebRequestRule) HasPublicDomainName() bool`

HasPublicDomainName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


