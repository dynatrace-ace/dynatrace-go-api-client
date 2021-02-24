# ExtensionListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extensions** | Pointer to [**[]ExtensionDto**](ExtensionDto.md) | A list of extensions. | [optional] 
**TotalResults** | Pointer to **int32** | The total number of entries in the result. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 

## Methods

### NewExtensionListDto

`func NewExtensionListDto() *ExtensionListDto`

NewExtensionListDto instantiates a new ExtensionListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionListDtoWithDefaults

`func NewExtensionListDtoWithDefaults() *ExtensionListDto`

NewExtensionListDtoWithDefaults instantiates a new ExtensionListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensions

`func (o *ExtensionListDto) GetExtensions() []ExtensionDto`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ExtensionListDto) GetExtensionsOk() (*[]ExtensionDto, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ExtensionListDto) SetExtensions(v []ExtensionDto)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ExtensionListDto) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetTotalResults

`func (o *ExtensionListDto) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ExtensionListDto) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ExtensionListDto) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *ExtensionListDto) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetNextPageKey

`func (o *ExtensionListDto) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *ExtensionListDto) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *ExtensionListDto) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *ExtensionListDto) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


