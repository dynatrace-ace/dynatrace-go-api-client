# ApplicationDetectionRuleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadataDtoImpl**](ConfigurationMetadataDtoImpl.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the rule. | [optional] 
**Order** | Pointer to **string** | The order of the rule in the rules list.   The rules are evaluated from top to bottom. The first matching rule applies. | [optional] 
**ApplicationIdentifier** | **string** | The Dynatrace entity ID of the application, for example &#x60;APPLICATION-4A3B43&#x60;.    You must use an existing ID. If you need to create a rule for an application that doesn&#39;t exist yet, [create an application first](https://www.dynatrace.com/support/help/shortlink/api-config-web-app-post-web-app) and then configure detection rules for it. | 
**FilterConfig** | [**ApplicationFilter**](ApplicationFilter.md) |  | 
**Name** | Pointer to **string** | The unique name of the Application detection rule. | [optional] 

## Methods

### NewApplicationDetectionRuleConfig

`func NewApplicationDetectionRuleConfig(applicationIdentifier string, filterConfig ApplicationFilter, ) *ApplicationDetectionRuleConfig`

NewApplicationDetectionRuleConfig instantiates a new ApplicationDetectionRuleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDetectionRuleConfigWithDefaults

`func NewApplicationDetectionRuleConfigWithDefaults() *ApplicationDetectionRuleConfig`

NewApplicationDetectionRuleConfigWithDefaults instantiates a new ApplicationDetectionRuleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ApplicationDetectionRuleConfig) GetMetadata() ConfigurationMetadataDtoImpl`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationDetectionRuleConfig) GetMetadataOk() (*ConfigurationMetadataDtoImpl, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationDetectionRuleConfig) SetMetadata(v ConfigurationMetadataDtoImpl)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationDetectionRuleConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *ApplicationDetectionRuleConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationDetectionRuleConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationDetectionRuleConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationDetectionRuleConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrder

`func (o *ApplicationDetectionRuleConfig) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApplicationDetectionRuleConfig) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApplicationDetectionRuleConfig) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ApplicationDetectionRuleConfig) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetApplicationIdentifier

`func (o *ApplicationDetectionRuleConfig) GetApplicationIdentifier() string`

GetApplicationIdentifier returns the ApplicationIdentifier field if non-nil, zero value otherwise.

### GetApplicationIdentifierOk

`func (o *ApplicationDetectionRuleConfig) GetApplicationIdentifierOk() (*string, bool)`

GetApplicationIdentifierOk returns a tuple with the ApplicationIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIdentifier

`func (o *ApplicationDetectionRuleConfig) SetApplicationIdentifier(v string)`

SetApplicationIdentifier sets ApplicationIdentifier field to given value.


### GetFilterConfig

`func (o *ApplicationDetectionRuleConfig) GetFilterConfig() ApplicationFilter`

GetFilterConfig returns the FilterConfig field if non-nil, zero value otherwise.

### GetFilterConfigOk

`func (o *ApplicationDetectionRuleConfig) GetFilterConfigOk() (*ApplicationFilter, bool)`

GetFilterConfigOk returns a tuple with the FilterConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterConfig

`func (o *ApplicationDetectionRuleConfig) SetFilterConfig(v ApplicationFilter)`

SetFilterConfig sets FilterConfig field to given value.


### GetName

`func (o *ApplicationDetectionRuleConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationDetectionRuleConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationDetectionRuleConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationDetectionRuleConfig) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


