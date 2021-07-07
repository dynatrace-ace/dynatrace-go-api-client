# ApiTokenList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiTokens** | Pointer to [**[]ApiToken**](ApiToken.md) | A list of API tokens. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 
**TotalCount** | **int64** | The total number of entries in the result. | 
**PageSize** | Pointer to **int32** | The number of entries per page. | [optional] 

## Methods

### NewApiTokenList

`func NewApiTokenList(totalCount int64, ) *ApiTokenList`

NewApiTokenList instantiates a new ApiTokenList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokenListWithDefaults

`func NewApiTokenListWithDefaults() *ApiTokenList`

NewApiTokenListWithDefaults instantiates a new ApiTokenList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiTokens

`func (o *ApiTokenList) GetApiTokens() []ApiToken`

GetApiTokens returns the ApiTokens field if non-nil, zero value otherwise.

### GetApiTokensOk

`func (o *ApiTokenList) GetApiTokensOk() (*[]ApiToken, bool)`

GetApiTokensOk returns a tuple with the ApiTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokens

`func (o *ApiTokenList) SetApiTokens(v []ApiToken)`

SetApiTokens sets ApiTokens field to given value.

### HasApiTokens

`func (o *ApiTokenList) HasApiTokens() bool`

HasApiTokens returns a boolean if a field has been set.

### GetNextPageKey

`func (o *ApiTokenList) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *ApiTokenList) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *ApiTokenList) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *ApiTokenList) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.

### GetTotalCount

`func (o *ApiTokenList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ApiTokenList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ApiTokenList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetPageSize

`func (o *ApiTokenList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ApiTokenList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ApiTokenList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ApiTokenList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


