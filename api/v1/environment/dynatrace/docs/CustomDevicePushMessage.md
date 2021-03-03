# CustomDevicePushMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the custom device that will appear in the user interface. | [optional] 
**Group** | Pointer to **string** | User defined group ID of entity.   The group ID helps to keep a consistent picture of device-group relations. One of many cases where a proper group is important is service detection: you can define which custom devices should lead to the same service by defining the same group ID for them.   If you set a group ID, it will be hashed into the Dynatrace entity ID of the custom device. In that case the custom device can only be part of one custom device group.   If you don&#39;t set the group ID, Dynatrace will create it based on the ID or type of the custom device. Also, the group will not be hashed into the device ID which means the device may switch groups. | [optional] 
**IpAddresses** | Pointer to **[]string** | The list of IP addresses that belong to the custom device.   These addresses are used to automatically discover the horizontal communication relationship between this component and all other observed components within Smartscape. Once a connection is discovered, it is automatically mapped and shown within Smartscape.   If you send a value (including an empty value), the existing values will be overwritten.   If you send &#x60;null&#x60; or omit this field, the existing values will be kept. | [optional] 
**ListenPorts** | Pointer to **[]int32** | The list of ports the custom devices listens to.   These ports are used to discover the horizontal communication relationship between this component and all other observed components within Smartscape.   Once a connection is discovered, it is automatically mapped and shown within Smartscape.   If ports are specified, you should also add at least one IP address or a host name for the custom device.   If you send a value, the existing values will be overwritten.   If you send &#x60;null&#x60;, or an empty value, or omit this field, the existing values will be kept. | [optional] 
**Type** | Pointer to **string** | The technology type definition of the custom device.   It must be the same technology type of the metric you&#39;re reporting.   If you send a value, the existing value will be overwritten.   If you send &#x60;null&#x60;, empty or omit this field, the existing value will be kept. | [optional] 
**Favicon** | Pointer to **string** | The icon to be displayed for your custom component within Smartscape. Provide the full URL of the icon file. | [optional] 
**ConfigUrl** | Pointer to **string** | The URL of a configuration web page for the custom device, such as a login page for a firewall or router. | [optional] 
**Properties** | Pointer to **map[string]string** | The list of key-value pair properties that will be shown beneath the infographics of your custom device. | [optional] 
**Tags** | Pointer to **[]string** | List of custom tags that you want to attach to your custom device. | [optional] 
**Series** | Pointer to [**[]EntityTimeseriesData**](EntityTimeseriesData.md) | The list of metric values that are reported for the custom device.   The metric you&#39;re reporting must already exist in Dynatrace. To learn how to create a custom metric, see [Timeseries API v1 - PUT a custom metric](https://dt-url.net/5k143rzb).   Dynatrace aggregates all the values you report for a custom device.   If you send a value (including an empty value), it will be added to the set of existing values.   If you send &#x60;null&#x60; or omit this field, the set of existing values won&#39;t change. | [optional] 
**HostNames** | Pointer to **[]string** | The list of host names related to the custom device.   These names are used to automatically discover the horizontal communication relationship between this component and all other observed components within Smartscape. Once a connection is discovered, it is automatically mapped and shown within Smartscape.   If you send a value, the existing values will be overwritten.   If you send &#x60;null&#x60; or an empty value; or omit this field, the existing values will be kept. | [optional] 

## Methods

### NewCustomDevicePushMessage

`func NewCustomDevicePushMessage() *CustomDevicePushMessage`

NewCustomDevicePushMessage instantiates a new CustomDevicePushMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDevicePushMessageWithDefaults

`func NewCustomDevicePushMessageWithDefaults() *CustomDevicePushMessage`

NewCustomDevicePushMessageWithDefaults instantiates a new CustomDevicePushMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CustomDevicePushMessage) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CustomDevicePushMessage) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CustomDevicePushMessage) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CustomDevicePushMessage) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGroup

`func (o *CustomDevicePushMessage) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CustomDevicePushMessage) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CustomDevicePushMessage) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CustomDevicePushMessage) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetIpAddresses

`func (o *CustomDevicePushMessage) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *CustomDevicePushMessage) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *CustomDevicePushMessage) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *CustomDevicePushMessage) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetListenPorts

`func (o *CustomDevicePushMessage) GetListenPorts() []int32`

GetListenPorts returns the ListenPorts field if non-nil, zero value otherwise.

### GetListenPortsOk

`func (o *CustomDevicePushMessage) GetListenPortsOk() (*[]int32, bool)`

GetListenPortsOk returns a tuple with the ListenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPorts

`func (o *CustomDevicePushMessage) SetListenPorts(v []int32)`

SetListenPorts sets ListenPorts field to given value.

### HasListenPorts

`func (o *CustomDevicePushMessage) HasListenPorts() bool`

HasListenPorts returns a boolean if a field has been set.

### GetType

`func (o *CustomDevicePushMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomDevicePushMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomDevicePushMessage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomDevicePushMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFavicon

`func (o *CustomDevicePushMessage) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *CustomDevicePushMessage) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *CustomDevicePushMessage) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *CustomDevicePushMessage) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### GetConfigUrl

`func (o *CustomDevicePushMessage) GetConfigUrl() string`

GetConfigUrl returns the ConfigUrl field if non-nil, zero value otherwise.

### GetConfigUrlOk

`func (o *CustomDevicePushMessage) GetConfigUrlOk() (*string, bool)`

GetConfigUrlOk returns a tuple with the ConfigUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigUrl

`func (o *CustomDevicePushMessage) SetConfigUrl(v string)`

SetConfigUrl sets ConfigUrl field to given value.

### HasConfigUrl

`func (o *CustomDevicePushMessage) HasConfigUrl() bool`

HasConfigUrl returns a boolean if a field has been set.

### GetProperties

`func (o *CustomDevicePushMessage) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CustomDevicePushMessage) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CustomDevicePushMessage) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CustomDevicePushMessage) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTags

`func (o *CustomDevicePushMessage) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CustomDevicePushMessage) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CustomDevicePushMessage) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CustomDevicePushMessage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSeries

`func (o *CustomDevicePushMessage) GetSeries() []EntityTimeseriesData`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *CustomDevicePushMessage) GetSeriesOk() (*[]EntityTimeseriesData, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *CustomDevicePushMessage) SetSeries(v []EntityTimeseriesData)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *CustomDevicePushMessage) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetHostNames

`func (o *CustomDevicePushMessage) GetHostNames() []string`

GetHostNames returns the HostNames field if non-nil, zero value otherwise.

### GetHostNamesOk

`func (o *CustomDevicePushMessage) GetHostNamesOk() (*[]string, bool)`

GetHostNamesOk returns a tuple with the HostNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNames

`func (o *CustomDevicePushMessage) SetHostNames(v []string)`

SetHostNames sets HostNames field to given value.

### HasHostNames

`func (o *CustomDevicePushMessage) HasHostNames() bool`

HasHostNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


