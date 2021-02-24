# ExtensionConfigurationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationsList** | Pointer to [**[]EntityShortRepresentation**](EntityShortRepresentation.md) | List of configurations. | [optional] 
**TotalResults** | Pointer to **int32** | The total number of entries in the result. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 

## Methods

### NewExtensionConfigurationList

`func NewExtensionConfigurationList() *ExtensionConfigurationList`

NewExtensionConfigurationList instantiates a new ExtensionConfigurationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionConfigurationListWithDefaults

`func NewExtensionConfigurationListWithDefaults() *ExtensionConfigurationList`

NewExtensionConfigurationListWithDefaults instantiates a new ExtensionConfigurationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationsList

`func (o *ExtensionConfigurationList) GetConfigurationsList() []EntityShortRepresentation`

GetConfigurationsList returns the ConfigurationsList field if non-nil, zero value otherwise.

### GetConfigurationsListOk

`func (o *ExtensionConfigurationList) GetConfigurationsListOk() (*[]EntityShortRepresentation, bool)`

GetConfigurationsListOk returns a tuple with the ConfigurationsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationsList

`func (o *ExtensionConfigurationList) SetConfigurationsList(v []EntityShortRepresentation)`

SetConfigurationsList sets ConfigurationsList field to given value.

### HasConfigurationsList

`func (o *ExtensionConfigurationList) HasConfigurationsList() bool`

HasConfigurationsList returns a boolean if a field has been set.

### GetTotalResults

`func (o *ExtensionConfigurationList) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ExtensionConfigurationList) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ExtensionConfigurationList) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *ExtensionConfigurationList) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetNextPageKey

`func (o *ExtensionConfigurationList) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *ExtensionConfigurationList) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *ExtensionConfigurationList) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *ExtensionConfigurationList) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


