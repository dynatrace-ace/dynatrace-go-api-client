# DownloadsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**DownloadCount** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**FirstDownloadTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDownloadsDto

`func NewDownloadsDto() *DownloadsDto`

NewDownloadsDto instantiates a new DownloadsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadsDtoWithDefaults

`func NewDownloadsDtoWithDefaults() *DownloadsDto`

NewDownloadsDtoWithDefaults instantiates a new DownloadsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DownloadsDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DownloadsDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DownloadsDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DownloadsDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDownloadCount

`func (o *DownloadsDto) GetDownloadCount() int32`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *DownloadsDto) GetDownloadCountOk() (*int32, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *DownloadsDto) SetDownloadCount(v int32)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *DownloadsDto) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### GetVersion

`func (o *DownloadsDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DownloadsDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DownloadsDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DownloadsDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFirstDownloadTime

`func (o *DownloadsDto) GetFirstDownloadTime() time.Time`

GetFirstDownloadTime returns the FirstDownloadTime field if non-nil, zero value otherwise.

### GetFirstDownloadTimeOk

`func (o *DownloadsDto) GetFirstDownloadTimeOk() (*time.Time, bool)`

GetFirstDownloadTimeOk returns a tuple with the FirstDownloadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDownloadTime

`func (o *DownloadsDto) SetFirstDownloadTime(v time.Time)`

SetFirstDownloadTime sets FirstDownloadTime field to given value.

### HasFirstDownloadTime

`func (o *DownloadsDto) HasFirstDownloadTime() bool`

HasFirstDownloadTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


