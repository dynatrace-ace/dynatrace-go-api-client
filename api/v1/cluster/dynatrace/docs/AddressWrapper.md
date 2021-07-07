# AddressWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Address | 
**NewAddress** | Pointer to **string** | New address | [optional] 
**UsePublicIpForDomainGeneration** | Pointer to **bool** | If true, public IP will be used for domain generation. | [optional] 
**UsePublicIpForAgents** | Pointer to **bool** | If true, public IP address will be used for OneAgents. | [optional] 

## Methods

### NewAddressWrapper

`func NewAddressWrapper(address string, ) *AddressWrapper`

NewAddressWrapper instantiates a new AddressWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWrapperWithDefaults

`func NewAddressWrapperWithDefaults() *AddressWrapper`

NewAddressWrapperWithDefaults instantiates a new AddressWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AddressWrapper) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressWrapper) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressWrapper) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNewAddress

`func (o *AddressWrapper) GetNewAddress() string`

GetNewAddress returns the NewAddress field if non-nil, zero value otherwise.

### GetNewAddressOk

`func (o *AddressWrapper) GetNewAddressOk() (*string, bool)`

GetNewAddressOk returns a tuple with the NewAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAddress

`func (o *AddressWrapper) SetNewAddress(v string)`

SetNewAddress sets NewAddress field to given value.

### HasNewAddress

`func (o *AddressWrapper) HasNewAddress() bool`

HasNewAddress returns a boolean if a field has been set.

### GetUsePublicIpForDomainGeneration

`func (o *AddressWrapper) GetUsePublicIpForDomainGeneration() bool`

GetUsePublicIpForDomainGeneration returns the UsePublicIpForDomainGeneration field if non-nil, zero value otherwise.

### GetUsePublicIpForDomainGenerationOk

`func (o *AddressWrapper) GetUsePublicIpForDomainGenerationOk() (*bool, bool)`

GetUsePublicIpForDomainGenerationOk returns a tuple with the UsePublicIpForDomainGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePublicIpForDomainGeneration

`func (o *AddressWrapper) SetUsePublicIpForDomainGeneration(v bool)`

SetUsePublicIpForDomainGeneration sets UsePublicIpForDomainGeneration field to given value.

### HasUsePublicIpForDomainGeneration

`func (o *AddressWrapper) HasUsePublicIpForDomainGeneration() bool`

HasUsePublicIpForDomainGeneration returns a boolean if a field has been set.

### GetUsePublicIpForAgents

`func (o *AddressWrapper) GetUsePublicIpForAgents() bool`

GetUsePublicIpForAgents returns the UsePublicIpForAgents field if non-nil, zero value otherwise.

### GetUsePublicIpForAgentsOk

`func (o *AddressWrapper) GetUsePublicIpForAgentsOk() (*bool, bool)`

GetUsePublicIpForAgentsOk returns a tuple with the UsePublicIpForAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePublicIpForAgents

`func (o *AddressWrapper) SetUsePublicIpForAgents(v bool)`

SetUsePublicIpForAgents sets UsePublicIpForAgents field to given value.

### HasUsePublicIpForAgents

`func (o *AddressWrapper) HasUsePublicIpForAgents() bool`

HasUsePublicIpForAgents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


