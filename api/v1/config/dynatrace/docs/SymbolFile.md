# SymbolFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | Pointer to **string** | The name of the application this file belongs to | [optional] 
**Size** | Pointer to **int32** | The size of the file in KB | [optional] 
**AppId** | Pointer to [**AppIdentifier**](AppIdentifier.md) |  | [optional] 
**UploadTimestamp** | Pointer to **int64** | The timestamp of the upload time of the file, in UTC milliseconds | [optional] 
**Pinned** | Pointer to **bool** | Is the file pinned and therefore cannot be deleted. | [optional] 

## Methods

### NewSymbolFile

`func NewSymbolFile() *SymbolFile`

NewSymbolFile instantiates a new SymbolFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolFileWithDefaults

`func NewSymbolFileWithDefaults() *SymbolFile`

NewSymbolFileWithDefaults instantiates a new SymbolFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationName

`func (o *SymbolFile) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *SymbolFile) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *SymbolFile) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *SymbolFile) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetSize

`func (o *SymbolFile) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SymbolFile) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SymbolFile) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *SymbolFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetAppId

`func (o *SymbolFile) GetAppId() AppIdentifier`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *SymbolFile) GetAppIdOk() (*AppIdentifier, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *SymbolFile) SetAppId(v AppIdentifier)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *SymbolFile) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetUploadTimestamp

`func (o *SymbolFile) GetUploadTimestamp() int64`

GetUploadTimestamp returns the UploadTimestamp field if non-nil, zero value otherwise.

### GetUploadTimestampOk

`func (o *SymbolFile) GetUploadTimestampOk() (*int64, bool)`

GetUploadTimestampOk returns a tuple with the UploadTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTimestamp

`func (o *SymbolFile) SetUploadTimestamp(v int64)`

SetUploadTimestamp sets UploadTimestamp field to given value.

### HasUploadTimestamp

`func (o *SymbolFile) HasUploadTimestamp() bool`

HasUploadTimestamp returns a boolean if a field has been set.

### GetPinned

`func (o *SymbolFile) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *SymbolFile) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *SymbolFile) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *SymbolFile) HasPinned() bool`

HasPinned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


