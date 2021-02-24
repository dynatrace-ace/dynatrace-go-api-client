# SymbolFilePinning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pinned** | **bool** | New setting for file pinning. True to pin the file, false to unpin the file | 

## Methods

### NewSymbolFilePinning

`func NewSymbolFilePinning(pinned bool, ) *SymbolFilePinning`

NewSymbolFilePinning instantiates a new SymbolFilePinning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolFilePinningWithDefaults

`func NewSymbolFilePinningWithDefaults() *SymbolFilePinning`

NewSymbolFilePinningWithDefaults instantiates a new SymbolFilePinning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPinned

`func (o *SymbolFilePinning) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *SymbolFilePinning) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *SymbolFilePinning) SetPinned(v bool)`

SetPinned sets Pinned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


