# Releases

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Releases** | Pointer to [**[]Release**](Release.md) | A list of releases. | [optional] 
**ReleasesWithProblems** | Pointer to **int64** | Number of releases with problems. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 
**TotalCount** | **int64** | The total number of entries in the result. | 
**PageSize** | Pointer to **int32** | The number of entries per page. | [optional] 

## Methods

### NewReleases

`func NewReleases(totalCount int64, ) *Releases`

NewReleases instantiates a new Releases object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleasesWithDefaults

`func NewReleasesWithDefaults() *Releases`

NewReleasesWithDefaults instantiates a new Releases object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleases

`func (o *Releases) GetReleases() []Release`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *Releases) GetReleasesOk() (*[]Release, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *Releases) SetReleases(v []Release)`

SetReleases sets Releases field to given value.

### HasReleases

`func (o *Releases) HasReleases() bool`

HasReleases returns a boolean if a field has been set.

### GetReleasesWithProblems

`func (o *Releases) GetReleasesWithProblems() int64`

GetReleasesWithProblems returns the ReleasesWithProblems field if non-nil, zero value otherwise.

### GetReleasesWithProblemsOk

`func (o *Releases) GetReleasesWithProblemsOk() (*int64, bool)`

GetReleasesWithProblemsOk returns a tuple with the ReleasesWithProblems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasesWithProblems

`func (o *Releases) SetReleasesWithProblems(v int64)`

SetReleasesWithProblems sets ReleasesWithProblems field to given value.

### HasReleasesWithProblems

`func (o *Releases) HasReleasesWithProblems() bool`

HasReleasesWithProblems returns a boolean if a field has been set.

### GetNextPageKey

`func (o *Releases) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *Releases) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *Releases) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *Releases) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.

### GetTotalCount

`func (o *Releases) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Releases) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Releases) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetPageSize

`func (o *Releases) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Releases) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *Releases) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *Releases) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


