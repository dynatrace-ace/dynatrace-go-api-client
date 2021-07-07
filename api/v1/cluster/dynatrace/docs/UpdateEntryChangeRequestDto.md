# UpdateEntryChangeRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedFromDownload** | Pointer to **bool** | If true, update package will be excluded from download and removed from storage. If false, already excluded package will be re-downloaded. | [optional] 

## Methods

### NewUpdateEntryChangeRequestDto

`func NewUpdateEntryChangeRequestDto() *UpdateEntryChangeRequestDto`

NewUpdateEntryChangeRequestDto instantiates a new UpdateEntryChangeRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEntryChangeRequestDtoWithDefaults

`func NewUpdateEntryChangeRequestDtoWithDefaults() *UpdateEntryChangeRequestDto`

NewUpdateEntryChangeRequestDtoWithDefaults instantiates a new UpdateEntryChangeRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludedFromDownload

`func (o *UpdateEntryChangeRequestDto) GetExcludedFromDownload() bool`

GetExcludedFromDownload returns the ExcludedFromDownload field if non-nil, zero value otherwise.

### GetExcludedFromDownloadOk

`func (o *UpdateEntryChangeRequestDto) GetExcludedFromDownloadOk() (*bool, bool)`

GetExcludedFromDownloadOk returns a tuple with the ExcludedFromDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedFromDownload

`func (o *UpdateEntryChangeRequestDto) SetExcludedFromDownload(v bool)`

SetExcludedFromDownload sets ExcludedFromDownload field to given value.

### HasExcludedFromDownload

`func (o *UpdateEntryChangeRequestDto) HasExcludedFromDownload() bool`

HasExcludedFromDownload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


