# UpgradeStartupState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | System precondition check state | [optional] 
**Error** | Pointer to **string** | Error | [optional] 

## Methods

### NewUpgradeStartupState

`func NewUpgradeStartupState() *UpgradeStartupState`

NewUpgradeStartupState instantiates a new UpgradeStartupState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeStartupStateWithDefaults

`func NewUpgradeStartupStateWithDefaults() *UpgradeStartupState`

NewUpgradeStartupStateWithDefaults instantiates a new UpgradeStartupState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *UpgradeStartupState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpgradeStartupState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpgradeStartupState) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpgradeStartupState) HasState() bool`

HasState returns a boolean if a field has been set.

### GetError

`func (o *UpgradeStartupState) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UpgradeStartupState) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UpgradeStartupState) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *UpgradeStartupState) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


