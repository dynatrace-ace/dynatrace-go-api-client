# HostGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeId** | Pointer to **string** | The Dynatrace entity ID of the host group. | [optional] 
**Name** | Pointer to **string** | The name of the Dynatrace entity, displayed in the UI. | [optional] 

## Methods

### NewHostGroup

`func NewHostGroup() *HostGroup`

NewHostGroup instantiates a new HostGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostGroupWithDefaults

`func NewHostGroupWithDefaults() *HostGroup`

NewHostGroupWithDefaults instantiates a new HostGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeId

`func (o *HostGroup) GetMeId() string`

GetMeId returns the MeId field if non-nil, zero value otherwise.

### GetMeIdOk

`func (o *HostGroup) GetMeIdOk() (*string, bool)`

GetMeIdOk returns a tuple with the MeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeId

`func (o *HostGroup) SetMeId(v string)`

SetMeId sets MeId field to given value.

### HasMeId

`func (o *HostGroup) HasMeId() bool`

HasMeId returns a boolean if a field has been set.

### GetName

`func (o *HostGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HostGroup) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


