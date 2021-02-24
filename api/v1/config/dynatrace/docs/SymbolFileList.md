# SymbolFileList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SymbolFiles** | Pointer to [**[]SymbolFile**](SymbolFile.md) | A list of symbolication files. | [optional] 

## Methods

### NewSymbolFileList

`func NewSymbolFileList() *SymbolFileList`

NewSymbolFileList instantiates a new SymbolFileList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolFileListWithDefaults

`func NewSymbolFileListWithDefaults() *SymbolFileList`

NewSymbolFileListWithDefaults instantiates a new SymbolFileList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbolFiles

`func (o *SymbolFileList) GetSymbolFiles() []SymbolFile`

GetSymbolFiles returns the SymbolFiles field if non-nil, zero value otherwise.

### GetSymbolFilesOk

`func (o *SymbolFileList) GetSymbolFilesOk() (*[]SymbolFile, bool)`

GetSymbolFilesOk returns a tuple with the SymbolFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolFiles

`func (o *SymbolFileList) SetSymbolFiles(v []SymbolFile)`

SetSymbolFiles sets SymbolFiles field to given value.

### HasSymbolFiles

`func (o *SymbolFileList) HasSymbolFiles() bool`

HasSymbolFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


