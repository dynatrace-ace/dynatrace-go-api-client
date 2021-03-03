# HostsListPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PercentageOfEnvironmentSearched** | Pointer to **float64** | The progress of the environment search, in percent. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results.    Has the value of &#x60;null&#x60; on the last page.   There might be another page of results even if the current page is empty. | [optional] 
**Hosts** | Pointer to [**[]HostAgentInfo**](HostAgentInfo.md) | A list of hosts with OneAgent deployment information for each host. | [optional] 

## Methods

### NewHostsListPage

`func NewHostsListPage() *HostsListPage`

NewHostsListPage instantiates a new HostsListPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostsListPageWithDefaults

`func NewHostsListPageWithDefaults() *HostsListPage`

NewHostsListPageWithDefaults instantiates a new HostsListPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentageOfEnvironmentSearched

`func (o *HostsListPage) GetPercentageOfEnvironmentSearched() float64`

GetPercentageOfEnvironmentSearched returns the PercentageOfEnvironmentSearched field if non-nil, zero value otherwise.

### GetPercentageOfEnvironmentSearchedOk

`func (o *HostsListPage) GetPercentageOfEnvironmentSearchedOk() (*float64, bool)`

GetPercentageOfEnvironmentSearchedOk returns a tuple with the PercentageOfEnvironmentSearched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOfEnvironmentSearched

`func (o *HostsListPage) SetPercentageOfEnvironmentSearched(v float64)`

SetPercentageOfEnvironmentSearched sets PercentageOfEnvironmentSearched field to given value.

### HasPercentageOfEnvironmentSearched

`func (o *HostsListPage) HasPercentageOfEnvironmentSearched() bool`

HasPercentageOfEnvironmentSearched returns a boolean if a field has been set.

### GetNextPageKey

`func (o *HostsListPage) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *HostsListPage) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *HostsListPage) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *HostsListPage) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.

### GetHosts

`func (o *HostsListPage) GetHosts() []HostAgentInfo`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *HostsListPage) GetHostsOk() (*[]HostAgentInfo, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *HostsListPage) SetHosts(v []HostAgentInfo)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *HostsListPage) HasHosts() bool`

HasHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


