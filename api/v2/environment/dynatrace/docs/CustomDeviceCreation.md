# CustomDeviceCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDeviceId** | **string** | The internal ID of the custom device.    If you use the ID of an existing device, the respective parameters will be updated. | 
**DisplayName** | **string** | The name of the custom device to be displayed in the user interface. | 
**Group** | Pointer to **string** | User defined group ID of entity.   The group ID helps to keep a consistent picture of device-group relations. One of many cases where a proper group is important is service detection: you can define which custom devices should lead to the same service by defining the same group ID for them.   If you set a group ID, it will be hashed into the Dynatrace entity ID of the custom device. In that case the custom device can only be part of one custom device group.   If you don&#39;t set the group ID, Dynatrace will create it based on the ID or type of the custom device. Also, the group will not be hashed into the device ID which means the device may switch groups. | [optional] 
**IpAddresses** | Pointer to **[]string** | The list of IP addresses that belong to the custom device.   These addresses are used to automatically discover the horizontal communication relationship between this component and all other observed components within Smartscape. Once a connection is discovered, it is automatically mapped and shown within Smartscape.   If you send a value (including an empty value), the existing values will be overwritten.   If you send &#x60;null&#x60; or omit this field, the existing values will be kept. | [optional] 
**ListenPorts** | Pointer to **[]int32** | The list of ports the custom devices listens to.   These ports are used to discover the horizontal communication relationship between this component and all other observed components within Smartscape.   Once a connection is discovered, it is automatically mapped and shown within Smartscape.   If ports are specified, you should also add at least one IP address or a DNS name for the custom device.   If you send a value, the existing values will be overwritten.   If you send &#x60;null&#x60;, or an empty value, or omit this field, the existing values will be kept. | [optional] 
**Type** | Pointer to **string** | The technology type definition of the custom device.   It must be the same technology type of the metric you&#39;re reporting.   If you send a value, the existing value will be overwritten.   If you send &#x60;null&#x60;, empty or omit this field, the existing value will be kept. | [optional] 
**FaviconUrl** | Pointer to **string** | The icon to be displayed for your custom component within Smartscape. Provide the full URL of the icon file. | [optional] 
**ConfigUrl** | Pointer to **string** | The URL of a configuration web page for the custom device, such as a login page for a firewall or router. | [optional] 
**Properties** | Pointer to **map[string]string** | The list of key-value pair properties that will be shown beneath the infographics of your custom device. | [optional] 
**DnsNames** | Pointer to **[]string** | The list of DNS names related to the custom device.   These names are used to automatically discover the horizontal communication relationship between this component and all other observed components within Smartscape. Once a connection is discovered, it is automatically mapped and shown within Smartscape.   If you send a value, the existing values will be overwritten.   If you send &#x60;null&#x60; or an empty value; or omit this field, the existing values will be kept. | [optional] 
**MessageType** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomDeviceCreation

`func NewCustomDeviceCreation(customDeviceId string, displayName string, ) *CustomDeviceCreation`

NewCustomDeviceCreation instantiates a new CustomDeviceCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDeviceCreationWithDefaults

`func NewCustomDeviceCreationWithDefaults() *CustomDeviceCreation`

NewCustomDeviceCreationWithDefaults instantiates a new CustomDeviceCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDeviceId

`func (o *CustomDeviceCreation) GetCustomDeviceId() string`

GetCustomDeviceId returns the CustomDeviceId field if non-nil, zero value otherwise.

### GetCustomDeviceIdOk

`func (o *CustomDeviceCreation) GetCustomDeviceIdOk() (*string, bool)`

GetCustomDeviceIdOk returns a tuple with the CustomDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDeviceId

`func (o *CustomDeviceCreation) SetCustomDeviceId(v string)`

SetCustomDeviceId sets CustomDeviceId field to given value.


### GetDisplayName

`func (o *CustomDeviceCreation) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CustomDeviceCreation) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CustomDeviceCreation) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetGroup

`func (o *CustomDeviceCreation) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CustomDeviceCreation) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CustomDeviceCreation) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CustomDeviceCreation) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetIpAddresses

`func (o *CustomDeviceCreation) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *CustomDeviceCreation) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *CustomDeviceCreation) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *CustomDeviceCreation) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetListenPorts

`func (o *CustomDeviceCreation) GetListenPorts() []int32`

GetListenPorts returns the ListenPorts field if non-nil, zero value otherwise.

### GetListenPortsOk

`func (o *CustomDeviceCreation) GetListenPortsOk() (*[]int32, bool)`

GetListenPortsOk returns a tuple with the ListenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPorts

`func (o *CustomDeviceCreation) SetListenPorts(v []int32)`

SetListenPorts sets ListenPorts field to given value.

### HasListenPorts

`func (o *CustomDeviceCreation) HasListenPorts() bool`

HasListenPorts returns a boolean if a field has been set.

### GetType

`func (o *CustomDeviceCreation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomDeviceCreation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomDeviceCreation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomDeviceCreation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFaviconUrl

`func (o *CustomDeviceCreation) GetFaviconUrl() string`

GetFaviconUrl returns the FaviconUrl field if non-nil, zero value otherwise.

### GetFaviconUrlOk

`func (o *CustomDeviceCreation) GetFaviconUrlOk() (*string, bool)`

GetFaviconUrlOk returns a tuple with the FaviconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaviconUrl

`func (o *CustomDeviceCreation) SetFaviconUrl(v string)`

SetFaviconUrl sets FaviconUrl field to given value.

### HasFaviconUrl

`func (o *CustomDeviceCreation) HasFaviconUrl() bool`

HasFaviconUrl returns a boolean if a field has been set.

### GetConfigUrl

`func (o *CustomDeviceCreation) GetConfigUrl() string`

GetConfigUrl returns the ConfigUrl field if non-nil, zero value otherwise.

### GetConfigUrlOk

`func (o *CustomDeviceCreation) GetConfigUrlOk() (*string, bool)`

GetConfigUrlOk returns a tuple with the ConfigUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigUrl

`func (o *CustomDeviceCreation) SetConfigUrl(v string)`

SetConfigUrl sets ConfigUrl field to given value.

### HasConfigUrl

`func (o *CustomDeviceCreation) HasConfigUrl() bool`

HasConfigUrl returns a boolean if a field has been set.

### GetProperties

`func (o *CustomDeviceCreation) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CustomDeviceCreation) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CustomDeviceCreation) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CustomDeviceCreation) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetDnsNames

`func (o *CustomDeviceCreation) GetDnsNames() []string`

GetDnsNames returns the DnsNames field if non-nil, zero value otherwise.

### GetDnsNamesOk

`func (o *CustomDeviceCreation) GetDnsNamesOk() (*[]string, bool)`

GetDnsNamesOk returns a tuple with the DnsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNames

`func (o *CustomDeviceCreation) SetDnsNames(v []string)`

SetDnsNames sets DnsNames field to given value.

### HasDnsNames

`func (o *CustomDeviceCreation) HasDnsNames() bool`

HasDnsNames returns a boolean if a field has been set.

### GetMessageType

`func (o *CustomDeviceCreation) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *CustomDeviceCreation) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *CustomDeviceCreation) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *CustomDeviceCreation) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


