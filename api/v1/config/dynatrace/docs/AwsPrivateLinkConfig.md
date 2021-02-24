# AwsPrivateLinkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Is AWS PrivateLink enabled | 
**VpcEndpointServiceName** | Pointer to **string** | The VirtualPrivateCluster-service name | [optional] [readonly] 

## Methods

### NewAwsPrivateLinkConfig

`func NewAwsPrivateLinkConfig(enabled bool, ) *AwsPrivateLinkConfig`

NewAwsPrivateLinkConfig instantiates a new AwsPrivateLinkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsPrivateLinkConfigWithDefaults

`func NewAwsPrivateLinkConfigWithDefaults() *AwsPrivateLinkConfig`

NewAwsPrivateLinkConfigWithDefaults instantiates a new AwsPrivateLinkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AwsPrivateLinkConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AwsPrivateLinkConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AwsPrivateLinkConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetVpcEndpointServiceName

`func (o *AwsPrivateLinkConfig) GetVpcEndpointServiceName() string`

GetVpcEndpointServiceName returns the VpcEndpointServiceName field if non-nil, zero value otherwise.

### GetVpcEndpointServiceNameOk

`func (o *AwsPrivateLinkConfig) GetVpcEndpointServiceNameOk() (*string, bool)`

GetVpcEndpointServiceNameOk returns a tuple with the VpcEndpointServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcEndpointServiceName

`func (o *AwsPrivateLinkConfig) SetVpcEndpointServiceName(v string)`

SetVpcEndpointServiceName sets VpcEndpointServiceName field to given value.

### HasVpcEndpointServiceName

`func (o *AwsPrivateLinkConfig) HasVpcEndpointServiceName() bool`

HasVpcEndpointServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


