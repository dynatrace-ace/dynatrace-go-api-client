# SecurityProblemList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int64** | The total number of entries in the result. | 
**PageSize** | Pointer to **int32** | The number of entries per page. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 
**SecurityProblems** | Pointer to [**[]SecurityProblem**](SecurityProblem.md) | A list of security problems. | [optional] [readonly] 

## Methods

### NewSecurityProblemList

`func NewSecurityProblemList(totalCount int64, ) *SecurityProblemList`

NewSecurityProblemList instantiates a new SecurityProblemList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityProblemListWithDefaults

`func NewSecurityProblemListWithDefaults() *SecurityProblemList`

NewSecurityProblemListWithDefaults instantiates a new SecurityProblemList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *SecurityProblemList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SecurityProblemList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SecurityProblemList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetPageSize

`func (o *SecurityProblemList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *SecurityProblemList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *SecurityProblemList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *SecurityProblemList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetNextPageKey

`func (o *SecurityProblemList) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *SecurityProblemList) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *SecurityProblemList) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *SecurityProblemList) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.

### GetSecurityProblems

`func (o *SecurityProblemList) GetSecurityProblems() []SecurityProblem`

GetSecurityProblems returns the SecurityProblems field if non-nil, zero value otherwise.

### GetSecurityProblemsOk

`func (o *SecurityProblemList) GetSecurityProblemsOk() (*[]SecurityProblem, bool)`

GetSecurityProblemsOk returns a tuple with the SecurityProblems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProblems

`func (o *SecurityProblemList) SetSecurityProblems(v []SecurityProblem)`

SetSecurityProblems sets SecurityProblems field to given value.

### HasSecurityProblems

`func (o *SecurityProblemList) HasSecurityProblems() bool`

HasSecurityProblems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


