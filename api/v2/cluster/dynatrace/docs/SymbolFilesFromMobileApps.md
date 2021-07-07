# SymbolFilesFromMobileApps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentlyUsed** | Pointer to **int64** | Currently used storage [bytes] | [optional] [readonly] 
**MaxLimit** | Pointer to **int64** | Maximum storage limit [bytes] | [optional] 

## Methods

### NewSymbolFilesFromMobileApps

`func NewSymbolFilesFromMobileApps() *SymbolFilesFromMobileApps`

NewSymbolFilesFromMobileApps instantiates a new SymbolFilesFromMobileApps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolFilesFromMobileAppsWithDefaults

`func NewSymbolFilesFromMobileAppsWithDefaults() *SymbolFilesFromMobileApps`

NewSymbolFilesFromMobileAppsWithDefaults instantiates a new SymbolFilesFromMobileApps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentlyUsed

`func (o *SymbolFilesFromMobileApps) GetCurrentlyUsed() int64`

GetCurrentlyUsed returns the CurrentlyUsed field if non-nil, zero value otherwise.

### GetCurrentlyUsedOk

`func (o *SymbolFilesFromMobileApps) GetCurrentlyUsedOk() (*int64, bool)`

GetCurrentlyUsedOk returns a tuple with the CurrentlyUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyUsed

`func (o *SymbolFilesFromMobileApps) SetCurrentlyUsed(v int64)`

SetCurrentlyUsed sets CurrentlyUsed field to given value.

### HasCurrentlyUsed

`func (o *SymbolFilesFromMobileApps) HasCurrentlyUsed() bool`

HasCurrentlyUsed returns a boolean if a field has been set.

### GetMaxLimit

`func (o *SymbolFilesFromMobileApps) GetMaxLimit() int64`

GetMaxLimit returns the MaxLimit field if non-nil, zero value otherwise.

### GetMaxLimitOk

`func (o *SymbolFilesFromMobileApps) GetMaxLimitOk() (*int64, bool)`

GetMaxLimitOk returns a tuple with the MaxLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLimit

`func (o *SymbolFilesFromMobileApps) SetMaxLimit(v int64)`

SetMaxLimit sets MaxLimit field to given value.

### HasMaxLimit

`func (o *SymbolFilesFromMobileApps) HasMaxLimit() bool`

HasMaxLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


