# EntityTypeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int64** | The total number of entries in the result. | 
**PageSize** | Pointer to **int32** | The number of entries per page. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 
**Types** | Pointer to [**[]EntityType**](EntityType.md) | The list of meta information for all available entity-types | [optional] 

## Methods

### NewEntityTypeList

`func NewEntityTypeList(totalCount int64, ) *EntityTypeList`

NewEntityTypeList instantiates a new EntityTypeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityTypeListWithDefaults

`func NewEntityTypeListWithDefaults() *EntityTypeList`

NewEntityTypeListWithDefaults instantiates a new EntityTypeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *EntityTypeList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *EntityTypeList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *EntityTypeList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetPageSize

`func (o *EntityTypeList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *EntityTypeList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *EntityTypeList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *EntityTypeList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetNextPageKey

`func (o *EntityTypeList) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *EntityTypeList) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *EntityTypeList) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *EntityTypeList) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.

### GetTypes

`func (o *EntityTypeList) GetTypes() []EntityType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *EntityTypeList) GetTypesOk() (*[]EntityType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *EntityTypeList) SetTypes(v []EntityType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *EntityTypeList) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


