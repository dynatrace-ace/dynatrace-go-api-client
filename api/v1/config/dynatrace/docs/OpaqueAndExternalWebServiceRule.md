# OpaqueAndExternalWebServiceRule

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
**Conditions** | Pointer to [**[]ConditionsOpaqueAndExternalWebServiceAttributeTypeDto**](ConditionsOpaqueAndExternalWebServiceAttributeTypeDto.md) | A list of conditions of the rule.   If several conditions are specified, the AND logic applies. | [optional] 
**DetectAsWebRequestService** | Pointer to **bool** | Detect the matching requests as web services (&#x60;false&#x60;) or web request services (&#x60;true&#x60;).   Setting this field to &#x60;true&#x60; prevents detecting of matching requests as opaque web services. An opaque web request service is created instead. If you need to further modify the resulting web request service, you need to create a separate rule of the &#x60;OPAQUE_AND_EXTERNAL_WEB_REQUEST&#x60; type.   Default is &#x60;false&#x60;, detecting matching requests as opaque web services. | [optional] 
**UrlPath** | Pointer to [**UrlPath**](UrlPath.md) |  | [optional] 
**Port** | Pointer to [**Port**](Port.md) |  | [optional] 

## Methods

### NewOpaqueAndExternalWebServiceRule

`func NewOpaqueAndExternalWebServiceRule(type_ string, name string, enabled bool, ) *OpaqueAndExternalWebServiceRule`

NewOpaqueAndExternalWebServiceRule instantiates a new OpaqueAndExternalWebServiceRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpaqueAndExternalWebServiceRuleWithDefaults

`func NewOpaqueAndExternalWebServiceRuleWithDefaults() *OpaqueAndExternalWebServiceRule`

NewOpaqueAndExternalWebServiceRuleWithDefaults instantiates a new OpaqueAndExternalWebServiceRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OpaqueAndExternalWebServiceRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpaqueAndExternalWebServiceRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpaqueAndExternalWebServiceRule) SetType(v string)`

SetType sets Type field to given value.


### GetMetadata

`func (o *OpaqueAndExternalWebServiceRule) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OpaqueAndExternalWebServiceRule) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OpaqueAndExternalWebServiceRule) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OpaqueAndExternalWebServiceRule) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetManagementZones

`func (o *OpaqueAndExternalWebServiceRule) GetManagementZones() []string`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *OpaqueAndExternalWebServiceRule) GetManagementZonesOk() (*[]string, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *OpaqueAndExternalWebServiceRule) SetManagementZones(v []string)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *OpaqueAndExternalWebServiceRule) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetId

`func (o *OpaqueAndExternalWebServiceRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpaqueAndExternalWebServiceRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpaqueAndExternalWebServiceRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OpaqueAndExternalWebServiceRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrder

`func (o *OpaqueAndExternalWebServiceRule) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OpaqueAndExternalWebServiceRule) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OpaqueAndExternalWebServiceRule) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *OpaqueAndExternalWebServiceRule) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetName

`func (o *OpaqueAndExternalWebServiceRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpaqueAndExternalWebServiceRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpaqueAndExternalWebServiceRule) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OpaqueAndExternalWebServiceRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpaqueAndExternalWebServiceRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpaqueAndExternalWebServiceRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpaqueAndExternalWebServiceRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *OpaqueAndExternalWebServiceRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OpaqueAndExternalWebServiceRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OpaqueAndExternalWebServiceRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetConditions

`func (o *OpaqueAndExternalWebServiceRule) GetConditions() []ConditionsOpaqueAndExternalWebServiceAttributeTypeDto`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *OpaqueAndExternalWebServiceRule) GetConditionsOk() (*[]ConditionsOpaqueAndExternalWebServiceAttributeTypeDto, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *OpaqueAndExternalWebServiceRule) SetConditions(v []ConditionsOpaqueAndExternalWebServiceAttributeTypeDto)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *OpaqueAndExternalWebServiceRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDetectAsWebRequestService

`func (o *OpaqueAndExternalWebServiceRule) GetDetectAsWebRequestService() bool`

GetDetectAsWebRequestService returns the DetectAsWebRequestService field if non-nil, zero value otherwise.

### GetDetectAsWebRequestServiceOk

`func (o *OpaqueAndExternalWebServiceRule) GetDetectAsWebRequestServiceOk() (*bool, bool)`

GetDetectAsWebRequestServiceOk returns a tuple with the DetectAsWebRequestService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectAsWebRequestService

`func (o *OpaqueAndExternalWebServiceRule) SetDetectAsWebRequestService(v bool)`

SetDetectAsWebRequestService sets DetectAsWebRequestService field to given value.

### HasDetectAsWebRequestService

`func (o *OpaqueAndExternalWebServiceRule) HasDetectAsWebRequestService() bool`

HasDetectAsWebRequestService returns a boolean if a field has been set.

### GetUrlPath

`func (o *OpaqueAndExternalWebServiceRule) GetUrlPath() UrlPath`

GetUrlPath returns the UrlPath field if non-nil, zero value otherwise.

### GetUrlPathOk

`func (o *OpaqueAndExternalWebServiceRule) GetUrlPathOk() (*UrlPath, bool)`

GetUrlPathOk returns a tuple with the UrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPath

`func (o *OpaqueAndExternalWebServiceRule) SetUrlPath(v UrlPath)`

SetUrlPath sets UrlPath field to given value.

### HasUrlPath

`func (o *OpaqueAndExternalWebServiceRule) HasUrlPath() bool`

HasUrlPath returns a boolean if a field has been set.

### GetPort

`func (o *OpaqueAndExternalWebServiceRule) GetPort() Port`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OpaqueAndExternalWebServiceRule) GetPortOk() (*Port, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OpaqueAndExternalWebServiceRule) SetPort(v Port)`

SetPort sets Port field to given value.

### HasPort

`func (o *OpaqueAndExternalWebServiceRule) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


