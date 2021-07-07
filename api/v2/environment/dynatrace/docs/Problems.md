# Problems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int64** | The total number of entries in the result. | 
**PageSize** | Pointer to **int32** | The number of entries per page. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 
**Problems** | Pointer to [**[]Problem**](Problem.md) | The result entries. | [optional] 
**Warnings** | Pointer to **[]string** | A list of warnings | [optional] 

## Methods

### NewProblems

`func NewProblems(totalCount int64, ) *Problems`

NewProblems instantiates a new Problems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemsWithDefaults

`func NewProblemsWithDefaults() *Problems`

NewProblemsWithDefaults instantiates a new Problems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *Problems) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Problems) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Problems) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetPageSize

`func (o *Problems) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Problems) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *Problems) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *Problems) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetNextPageKey

`func (o *Problems) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *Problems) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *Problems) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *Problems) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.

### GetProblems

`func (o *Problems) GetProblems() []Problem`

GetProblems returns the Problems field if non-nil, zero value otherwise.

### GetProblemsOk

`func (o *Problems) GetProblemsOk() (*[]Problem, bool)`

GetProblemsOk returns a tuple with the Problems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblems

`func (o *Problems) SetProblems(v []Problem)`

SetProblems sets Problems field to given value.

### HasProblems

`func (o *Problems) HasProblems() bool`

HasProblems returns a boolean if a field has been set.

### GetWarnings

`func (o *Problems) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Problems) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Problems) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *Problems) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


