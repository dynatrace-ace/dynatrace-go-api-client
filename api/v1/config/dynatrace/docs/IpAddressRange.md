# IpAddressRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetMask** | Pointer to **int32** | The subnet mask of the IP address range. | [optional] 
**Address** | **string** | The IP address to be mapped.   For an IP address range, this is the **from** address. | 
**AddressTo** | Pointer to **string** | The **to** address of the IP address range. | [optional] 

## Methods

### NewIpAddressRange

`func NewIpAddressRange(address string, ) *IpAddressRange`

NewIpAddressRange instantiates a new IpAddressRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAddressRangeWithDefaults

`func NewIpAddressRangeWithDefaults() *IpAddressRange`

NewIpAddressRangeWithDefaults instantiates a new IpAddressRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetMask

`func (o *IpAddressRange) GetSubnetMask() int32`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *IpAddressRange) GetSubnetMaskOk() (*int32, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *IpAddressRange) SetSubnetMask(v int32)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *IpAddressRange) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetAddress

`func (o *IpAddressRange) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IpAddressRange) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IpAddressRange) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddressTo

`func (o *IpAddressRange) GetAddressTo() string`

GetAddressTo returns the AddressTo field if non-nil, zero value otherwise.

### GetAddressToOk

`func (o *IpAddressRange) GetAddressToOk() (*string, bool)`

GetAddressToOk returns a tuple with the AddressTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressTo

`func (o *IpAddressRange) SetAddressTo(v string)`

SetAddressTo sets AddressTo field to given value.

### HasAddressTo

`func (o *IpAddressRange) HasAddressTo() bool`

HasAddressTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


