# ExtensionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extensions** | Pointer to [**[]MinimalExtension**](MinimalExtension.md) | A list of extensions. | [optional] 
**TotalCount** | **int64** | The total number of entries in the result. | 
**PageSize** | Pointer to **int32** | The number of entries per page. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 

## Methods

### NewExtensionList

`func NewExtensionList(totalCount int64, ) *ExtensionList`

NewExtensionList instantiates a new ExtensionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionListWithDefaults

`func NewExtensionListWithDefaults() *ExtensionList`

NewExtensionListWithDefaults instantiates a new ExtensionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensions

`func (o *ExtensionList) GetExtensions() []MinimalExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ExtensionList) GetExtensionsOk() (*[]MinimalExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ExtensionList) SetExtensions(v []MinimalExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ExtensionList) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetTotalCount

`func (o *ExtensionList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ExtensionList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ExtensionList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetPageSize

`func (o *ExtensionList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ExtensionList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ExtensionList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ExtensionList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetNextPageKey

`func (o *ExtensionList) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *ExtensionList) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *ExtensionList) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *ExtensionList) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


