# WebApplicationConfigIpAddressRestrictionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | The mode of the list of ip address restrictions. | 
**IpAddressRestrictions** | Pointer to [**[]IpAddressRange**](IpAddressRange.md) |  | [optional] 

## Methods

### NewWebApplicationConfigIpAddressRestrictionSettings

`func NewWebApplicationConfigIpAddressRestrictionSettings(mode string, ) *WebApplicationConfigIpAddressRestrictionSettings`

NewWebApplicationConfigIpAddressRestrictionSettings instantiates a new WebApplicationConfigIpAddressRestrictionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebApplicationConfigIpAddressRestrictionSettingsWithDefaults

`func NewWebApplicationConfigIpAddressRestrictionSettingsWithDefaults() *WebApplicationConfigIpAddressRestrictionSettings`

NewWebApplicationConfigIpAddressRestrictionSettingsWithDefaults instantiates a new WebApplicationConfigIpAddressRestrictionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *WebApplicationConfigIpAddressRestrictionSettings) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *WebApplicationConfigIpAddressRestrictionSettings) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *WebApplicationConfigIpAddressRestrictionSettings) SetMode(v string)`

SetMode sets Mode field to given value.


### GetIpAddressRestrictions

`func (o *WebApplicationConfigIpAddressRestrictionSettings) GetIpAddressRestrictions() []IpAddressRange`

GetIpAddressRestrictions returns the IpAddressRestrictions field if non-nil, zero value otherwise.

### GetIpAddressRestrictionsOk

`func (o *WebApplicationConfigIpAddressRestrictionSettings) GetIpAddressRestrictionsOk() (*[]IpAddressRange, bool)`

GetIpAddressRestrictionsOk returns a tuple with the IpAddressRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressRestrictions

`func (o *WebApplicationConfigIpAddressRestrictionSettings) SetIpAddressRestrictions(v []IpAddressRange)`

SetIpAddressRestrictions sets IpAddressRestrictions field to given value.

### HasIpAddressRestrictions

`func (o *WebApplicationConfigIpAddressRestrictionSettings) HasIpAddressRestrictions() bool`

HasIpAddressRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


