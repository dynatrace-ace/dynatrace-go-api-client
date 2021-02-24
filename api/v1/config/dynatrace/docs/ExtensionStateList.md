# ExtensionStateList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**States** | Pointer to [**[]ExtensionState**](ExtensionState.md) | A list of extension states. | [optional] 
**TotalResults** | Pointer to **int32** | The total number of entries in the result. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 

## Methods

### NewExtensionStateList

`func NewExtensionStateList() *ExtensionStateList`

NewExtensionStateList instantiates a new ExtensionStateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionStateListWithDefaults

`func NewExtensionStateListWithDefaults() *ExtensionStateList`

NewExtensionStateListWithDefaults instantiates a new ExtensionStateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStates

`func (o *ExtensionStateList) GetStates() []ExtensionState`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *ExtensionStateList) GetStatesOk() (*[]ExtensionState, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *ExtensionStateList) SetStates(v []ExtensionState)`

SetStates sets States field to given value.

### HasStates

`func (o *ExtensionStateList) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetTotalResults

`func (o *ExtensionStateList) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ExtensionStateList) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ExtensionStateList) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *ExtensionStateList) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetNextPageKey

`func (o *ExtensionStateList) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *ExtensionStateList) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *ExtensionStateList) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *ExtensionStateList) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


