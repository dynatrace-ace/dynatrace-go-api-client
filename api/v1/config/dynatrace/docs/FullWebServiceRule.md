# FullWebServiceRule

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
**Conditions** | Pointer to [**[]ConditionsFullWebServiceAttributeTypeDto**](ConditionsFullWebServiceAttributeTypeDto.md) | A list of conditions of the rule.   If several conditions are specified, the AND logic applies. | [optional] 
**DetectAsWebRequestService** | Pointer to **bool** | Detect the matching requests as full web services (&#x60;false&#x60;) or web request services (&#x60;true&#x60;).   Setting this field to &#x60;true&#x60; prevents detecting of matching requests as full web services. A web request service is created instead. If you need to further modify the resulting web request service, you need to create a separate rule of the &#x60;FULL_WEB_REQUEST&#x60; type.   Default is &#x60;false&#x60;, detecting matching requests as full web services. | [optional] 
**WebServiceName** | Pointer to [**WebServiceName**](WebServiceName.md) |  | [optional] 
**WebServiceNameSpace** | Pointer to [**WebServiceNameSpace**](WebServiceNameSpace.md) |  | [optional] 
**ApplicationId** | Pointer to [**ApplicationId**](ApplicationId.md) |  | [optional] 
**ContextRoot** | Pointer to [**ContextRoot**](ContextRoot.md) |  | [optional] 
**ServerName** | Pointer to [**ServerName**](ServerName.md) |  | [optional] 

## Methods

### NewFullWebServiceRule

`func NewFullWebServiceRule(type_ string, name string, enabled bool, ) *FullWebServiceRule`

NewFullWebServiceRule instantiates a new FullWebServiceRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullWebServiceRuleWithDefaults

`func NewFullWebServiceRuleWithDefaults() *FullWebServiceRule`

NewFullWebServiceRuleWithDefaults instantiates a new FullWebServiceRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FullWebServiceRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FullWebServiceRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FullWebServiceRule) SetType(v string)`

SetType sets Type field to given value.


### GetMetadata

`func (o *FullWebServiceRule) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FullWebServiceRule) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FullWebServiceRule) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FullWebServiceRule) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetManagementZones

`func (o *FullWebServiceRule) GetManagementZones() []string`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *FullWebServiceRule) GetManagementZonesOk() (*[]string, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *FullWebServiceRule) SetManagementZones(v []string)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *FullWebServiceRule) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetId

`func (o *FullWebServiceRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullWebServiceRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullWebServiceRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FullWebServiceRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrder

`func (o *FullWebServiceRule) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FullWebServiceRule) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FullWebServiceRule) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *FullWebServiceRule) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetName

`func (o *FullWebServiceRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullWebServiceRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullWebServiceRule) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FullWebServiceRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FullWebServiceRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FullWebServiceRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FullWebServiceRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *FullWebServiceRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FullWebServiceRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FullWebServiceRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetConditions

`func (o *FullWebServiceRule) GetConditions() []ConditionsFullWebServiceAttributeTypeDto`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *FullWebServiceRule) GetConditionsOk() (*[]ConditionsFullWebServiceAttributeTypeDto, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *FullWebServiceRule) SetConditions(v []ConditionsFullWebServiceAttributeTypeDto)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *FullWebServiceRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDetectAsWebRequestService

`func (o *FullWebServiceRule) GetDetectAsWebRequestService() bool`

GetDetectAsWebRequestService returns the DetectAsWebRequestService field if non-nil, zero value otherwise.

### GetDetectAsWebRequestServiceOk

`func (o *FullWebServiceRule) GetDetectAsWebRequestServiceOk() (*bool, bool)`

GetDetectAsWebRequestServiceOk returns a tuple with the DetectAsWebRequestService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectAsWebRequestService

`func (o *FullWebServiceRule) SetDetectAsWebRequestService(v bool)`

SetDetectAsWebRequestService sets DetectAsWebRequestService field to given value.

### HasDetectAsWebRequestService

`func (o *FullWebServiceRule) HasDetectAsWebRequestService() bool`

HasDetectAsWebRequestService returns a boolean if a field has been set.

### GetWebServiceName

`func (o *FullWebServiceRule) GetWebServiceName() WebServiceName`

GetWebServiceName returns the WebServiceName field if non-nil, zero value otherwise.

### GetWebServiceNameOk

`func (o *FullWebServiceRule) GetWebServiceNameOk() (*WebServiceName, bool)`

GetWebServiceNameOk returns a tuple with the WebServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceName

`func (o *FullWebServiceRule) SetWebServiceName(v WebServiceName)`

SetWebServiceName sets WebServiceName field to given value.

### HasWebServiceName

`func (o *FullWebServiceRule) HasWebServiceName() bool`

HasWebServiceName returns a boolean if a field has been set.

### GetWebServiceNameSpace

`func (o *FullWebServiceRule) GetWebServiceNameSpace() WebServiceNameSpace`

GetWebServiceNameSpace returns the WebServiceNameSpace field if non-nil, zero value otherwise.

### GetWebServiceNameSpaceOk

`func (o *FullWebServiceRule) GetWebServiceNameSpaceOk() (*WebServiceNameSpace, bool)`

GetWebServiceNameSpaceOk returns a tuple with the WebServiceNameSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceNameSpace

`func (o *FullWebServiceRule) SetWebServiceNameSpace(v WebServiceNameSpace)`

SetWebServiceNameSpace sets WebServiceNameSpace field to given value.

### HasWebServiceNameSpace

`func (o *FullWebServiceRule) HasWebServiceNameSpace() bool`

HasWebServiceNameSpace returns a boolean if a field has been set.

### GetApplicationId

`func (o *FullWebServiceRule) GetApplicationId() ApplicationId`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *FullWebServiceRule) GetApplicationIdOk() (*ApplicationId, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *FullWebServiceRule) SetApplicationId(v ApplicationId)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *FullWebServiceRule) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetContextRoot

`func (o *FullWebServiceRule) GetContextRoot() ContextRoot`

GetContextRoot returns the ContextRoot field if non-nil, zero value otherwise.

### GetContextRootOk

`func (o *FullWebServiceRule) GetContextRootOk() (*ContextRoot, bool)`

GetContextRootOk returns a tuple with the ContextRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextRoot

`func (o *FullWebServiceRule) SetContextRoot(v ContextRoot)`

SetContextRoot sets ContextRoot field to given value.

### HasContextRoot

`func (o *FullWebServiceRule) HasContextRoot() bool`

HasContextRoot returns a boolean if a field has been set.

### GetServerName

`func (o *FullWebServiceRule) GetServerName() ServerName`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *FullWebServiceRule) GetServerNameOk() (*ServerName, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *FullWebServiceRule) SetServerName(v ServerName)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *FullWebServiceRule) HasServerName() bool`

HasServerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


