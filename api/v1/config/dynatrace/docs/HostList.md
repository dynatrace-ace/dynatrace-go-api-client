# HostList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResults** | Pointer to **int32** | Total number of results | [optional] 
**Hosts** | Pointer to [**[]Host**](Host.md) | The list of hosts | [optional] 
**NextPageKey** | Pointer to **string** | Next page key used for paging | [optional] 

## Methods

### NewHostList

`func NewHostList() *HostList`

NewHostList instantiates a new HostList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostListWithDefaults

`func NewHostListWithDefaults() *HostList`

NewHostListWithDefaults instantiates a new HostList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResults

`func (o *HostList) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *HostList) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *HostList) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *HostList) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetHosts

`func (o *HostList) GetHosts() []Host`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *HostList) GetHostsOk() (*[]Host, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *HostList) SetHosts(v []Host)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *HostList) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetNextPageKey

`func (o *HostList) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *HostList) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *HostList) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *HostList) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


