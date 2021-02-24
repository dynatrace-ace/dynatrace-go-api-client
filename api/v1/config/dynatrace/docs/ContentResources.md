# ContentResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceProviders** | Pointer to [**[]ResourceProvider**](ResourceProvider.md) | An ordered list of manually added content providers.   Rules are evaluated from top to bottom; the first matching rules applies. | [optional] 
**ResourceUrlCleanupRules** | Pointer to [**[]ResourceUrlCleanupRule**](ResourceUrlCleanupRule.md) | An ordered list of resource URL cleanup rules.   Rules are evaluated from top to bottom; the first matching rules applies. | [optional] 
**ResourceTypes** | Pointer to [**[]ResourceType**](ResourceType.md) | An ordered list of manually defined resource types.   Rules are evaluated from top to bottom; the first matching rules applies. | [optional] 

## Methods

### NewContentResources

`func NewContentResources() *ContentResources`

NewContentResources instantiates a new ContentResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentResourcesWithDefaults

`func NewContentResourcesWithDefaults() *ContentResources`

NewContentResourcesWithDefaults instantiates a new ContentResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceProviders

`func (o *ContentResources) GetResourceProviders() []ResourceProvider`

GetResourceProviders returns the ResourceProviders field if non-nil, zero value otherwise.

### GetResourceProvidersOk

`func (o *ContentResources) GetResourceProvidersOk() (*[]ResourceProvider, bool)`

GetResourceProvidersOk returns a tuple with the ResourceProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProviders

`func (o *ContentResources) SetResourceProviders(v []ResourceProvider)`

SetResourceProviders sets ResourceProviders field to given value.

### HasResourceProviders

`func (o *ContentResources) HasResourceProviders() bool`

HasResourceProviders returns a boolean if a field has been set.

### GetResourceUrlCleanupRules

`func (o *ContentResources) GetResourceUrlCleanupRules() []ResourceUrlCleanupRule`

GetResourceUrlCleanupRules returns the ResourceUrlCleanupRules field if non-nil, zero value otherwise.

### GetResourceUrlCleanupRulesOk

`func (o *ContentResources) GetResourceUrlCleanupRulesOk() (*[]ResourceUrlCleanupRule, bool)`

GetResourceUrlCleanupRulesOk returns a tuple with the ResourceUrlCleanupRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUrlCleanupRules

`func (o *ContentResources) SetResourceUrlCleanupRules(v []ResourceUrlCleanupRule)`

SetResourceUrlCleanupRules sets ResourceUrlCleanupRules field to given value.

### HasResourceUrlCleanupRules

`func (o *ContentResources) HasResourceUrlCleanupRules() bool`

HasResourceUrlCleanupRules returns a boolean if a field has been set.

### GetResourceTypes

`func (o *ContentResources) GetResourceTypes() []ResourceType`

GetResourceTypes returns the ResourceTypes field if non-nil, zero value otherwise.

### GetResourceTypesOk

`func (o *ContentResources) GetResourceTypesOk() (*[]ResourceType, bool)`

GetResourceTypesOk returns a tuple with the ResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypes

`func (o *ContentResources) SetResourceTypes(v []ResourceType)`

SetResourceTypes sets ResourceTypes field to given value.

### HasResourceTypes

`func (o *ContentResources) HasResourceTypes() bool`

HasResourceTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


