# ResourceProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | The name of the provider. | 
**ResourceType** | **string** | The type of the provider. | 
**BrandIconUrl** | **string** | The URL of the provider&#39;s icon. | 
**DomainNamePatterns** | **[]string** | A list of domain patterns of the provider. | 

## Methods

### NewResourceProvider

`func NewResourceProvider(resourceName string, resourceType string, brandIconUrl string, domainNamePatterns []string, ) *ResourceProvider`

NewResourceProvider instantiates a new ResourceProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceProviderWithDefaults

`func NewResourceProviderWithDefaults() *ResourceProvider`

NewResourceProviderWithDefaults instantiates a new ResourceProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *ResourceProvider) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceProvider) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ResourceProvider) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetResourceType

`func (o *ResourceProvider) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceProvider) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceProvider) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetBrandIconUrl

`func (o *ResourceProvider) GetBrandIconUrl() string`

GetBrandIconUrl returns the BrandIconUrl field if non-nil, zero value otherwise.

### GetBrandIconUrlOk

`func (o *ResourceProvider) GetBrandIconUrlOk() (*string, bool)`

GetBrandIconUrlOk returns a tuple with the BrandIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandIconUrl

`func (o *ResourceProvider) SetBrandIconUrl(v string)`

SetBrandIconUrl sets BrandIconUrl field to given value.


### GetDomainNamePatterns

`func (o *ResourceProvider) GetDomainNamePatterns() []string`

GetDomainNamePatterns returns the DomainNamePatterns field if non-nil, zero value otherwise.

### GetDomainNamePatternsOk

`func (o *ResourceProvider) GetDomainNamePatternsOk() (*[]string, bool)`

GetDomainNamePatternsOk returns a tuple with the DomainNamePatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNamePatterns

`func (o *ResourceProvider) SetDomainNamePatterns(v []string)`

SetDomainNamePatterns sets DomainNamePatterns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


