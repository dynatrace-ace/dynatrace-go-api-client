# DavisSecurityAdviceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int64** | The total number of entries in the result. | 
**PageSize** | Pointer to **int32** | The number of entries per page. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 
**Advices** | Pointer to [**[]DavisSecurityAdvice**](DavisSecurityAdvice.md) |  | [optional] 

## Methods

### NewDavisSecurityAdviceList

`func NewDavisSecurityAdviceList(totalCount int64, ) *DavisSecurityAdviceList`

NewDavisSecurityAdviceList instantiates a new DavisSecurityAdviceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDavisSecurityAdviceListWithDefaults

`func NewDavisSecurityAdviceListWithDefaults() *DavisSecurityAdviceList`

NewDavisSecurityAdviceListWithDefaults instantiates a new DavisSecurityAdviceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *DavisSecurityAdviceList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DavisSecurityAdviceList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DavisSecurityAdviceList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetPageSize

`func (o *DavisSecurityAdviceList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DavisSecurityAdviceList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DavisSecurityAdviceList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DavisSecurityAdviceList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetNextPageKey

`func (o *DavisSecurityAdviceList) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *DavisSecurityAdviceList) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *DavisSecurityAdviceList) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *DavisSecurityAdviceList) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.

### GetAdvices

`func (o *DavisSecurityAdviceList) GetAdvices() []DavisSecurityAdvice`

GetAdvices returns the Advices field if non-nil, zero value otherwise.

### GetAdvicesOk

`func (o *DavisSecurityAdviceList) GetAdvicesOk() (*[]DavisSecurityAdvice, bool)`

GetAdvicesOk returns a tuple with the Advices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvices

`func (o *DavisSecurityAdviceList) SetAdvices(v []DavisSecurityAdvice)`

SetAdvices sets Advices field to given value.

### HasAdvices

`func (o *DavisSecurityAdviceList) HasAdvices() bool`

HasAdvices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


