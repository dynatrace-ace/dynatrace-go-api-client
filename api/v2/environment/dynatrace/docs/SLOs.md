# SLOs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slo** | Pointer to [**[]SLO**](SLO.md) | A list of SLOs. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 
**TotalCount** | **int64** | The total number of entries in the result. | 
**PageSize** | Pointer to **int32** | The number of entries per page. | [optional] 

## Methods

### NewSLOs

`func NewSLOs(totalCount int64, ) *SLOs`

NewSLOs instantiates a new SLOs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLOsWithDefaults

`func NewSLOsWithDefaults() *SLOs`

NewSLOsWithDefaults instantiates a new SLOs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlo

`func (o *SLOs) GetSlo() []SLO`

GetSlo returns the Slo field if non-nil, zero value otherwise.

### GetSloOk

`func (o *SLOs) GetSloOk() (*[]SLO, bool)`

GetSloOk returns a tuple with the Slo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlo

`func (o *SLOs) SetSlo(v []SLO)`

SetSlo sets Slo field to given value.

### HasSlo

`func (o *SLOs) HasSlo() bool`

HasSlo returns a boolean if a field has been set.

### GetNextPageKey

`func (o *SLOs) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *SLOs) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *SLOs) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *SLOs) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.

### GetTotalCount

`func (o *SLOs) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SLOs) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SLOs) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetPageSize

`func (o *SLOs) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *SLOs) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *SLOs) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *SLOs) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


