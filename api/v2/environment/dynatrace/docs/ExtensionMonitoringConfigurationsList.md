# ExtensionMonitoringConfigurationsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ExtensionMonitoringConfiguration**](ExtensionMonitoringConfiguration.md) | A list of extension monitoring configurations. | [optional] 
**TotalCount** | **int64** | The total number of entries in the result. | 
**PageSize** | Pointer to **int32** | The number of entries per page. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 

## Methods

### NewExtensionMonitoringConfigurationsList

`func NewExtensionMonitoringConfigurationsList(totalCount int64, ) *ExtensionMonitoringConfigurationsList`

NewExtensionMonitoringConfigurationsList instantiates a new ExtensionMonitoringConfigurationsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionMonitoringConfigurationsListWithDefaults

`func NewExtensionMonitoringConfigurationsListWithDefaults() *ExtensionMonitoringConfigurationsList`

NewExtensionMonitoringConfigurationsListWithDefaults instantiates a new ExtensionMonitoringConfigurationsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ExtensionMonitoringConfigurationsList) GetItems() []ExtensionMonitoringConfiguration`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ExtensionMonitoringConfigurationsList) GetItemsOk() (*[]ExtensionMonitoringConfiguration, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ExtensionMonitoringConfigurationsList) SetItems(v []ExtensionMonitoringConfiguration)`

SetItems sets Items field to given value.

### HasItems

`func (o *ExtensionMonitoringConfigurationsList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalCount

`func (o *ExtensionMonitoringConfigurationsList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ExtensionMonitoringConfigurationsList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ExtensionMonitoringConfigurationsList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetPageSize

`func (o *ExtensionMonitoringConfigurationsList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ExtensionMonitoringConfigurationsList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ExtensionMonitoringConfigurationsList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ExtensionMonitoringConfigurationsList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetNextPageKey

`func (o *ExtensionMonitoringConfigurationsList) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *ExtensionMonitoringConfigurationsList) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *ExtensionMonitoringConfigurationsList) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *ExtensionMonitoringConfigurationsList) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


