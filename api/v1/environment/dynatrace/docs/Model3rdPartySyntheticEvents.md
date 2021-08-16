# Model3rdPartySyntheticEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyntheticEngineName** | **string** | The type of the third-party synthetic monitor. | 
**Open** | Pointer to [**[]Model3rdPartyEventOpenNotification**](Model3rdPartyEventOpenNotification.md) | The list of open third-party synthetic events. | [optional] 
**Resolved** | Pointer to [**[]Model3rdPartyEventResolvedNotification**](Model3rdPartyEventResolvedNotification.md) | The list of closed third-party synthetic events. | [optional] 

## Methods

### NewModel3rdPartySyntheticEvents

`func NewModel3rdPartySyntheticEvents(syntheticEngineName string, ) *Model3rdPartySyntheticEvents`

NewModel3rdPartySyntheticEvents instantiates a new Model3rdPartySyntheticEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel3rdPartySyntheticEventsWithDefaults

`func NewModel3rdPartySyntheticEventsWithDefaults() *Model3rdPartySyntheticEvents`

NewModel3rdPartySyntheticEventsWithDefaults instantiates a new Model3rdPartySyntheticEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyntheticEngineName

`func (o *Model3rdPartySyntheticEvents) GetSyntheticEngineName() string`

GetSyntheticEngineName returns the SyntheticEngineName field if non-nil, zero value otherwise.

### GetSyntheticEngineNameOk

`func (o *Model3rdPartySyntheticEvents) GetSyntheticEngineNameOk() (*string, bool)`

GetSyntheticEngineNameOk returns a tuple with the SyntheticEngineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticEngineName

`func (o *Model3rdPartySyntheticEvents) SetSyntheticEngineName(v string)`

SetSyntheticEngineName sets SyntheticEngineName field to given value.


### GetOpen

`func (o *Model3rdPartySyntheticEvents) GetOpen() []Model3rdPartyEventOpenNotification`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *Model3rdPartySyntheticEvents) GetOpenOk() (*[]Model3rdPartyEventOpenNotification, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *Model3rdPartySyntheticEvents) SetOpen(v []Model3rdPartyEventOpenNotification)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *Model3rdPartySyntheticEvents) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetResolved

`func (o *Model3rdPartySyntheticEvents) GetResolved() []Model3rdPartyEventResolvedNotification`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *Model3rdPartySyntheticEvents) GetResolvedOk() (*[]Model3rdPartyEventResolvedNotification, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *Model3rdPartySyntheticEvents) SetResolved(v []Model3rdPartyEventResolvedNotification)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *Model3rdPartySyntheticEvents) HasResolved() bool`

HasResolved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


