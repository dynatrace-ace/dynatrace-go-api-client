# Host

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the host | [optional] 
**Id** | Pointer to **string** | The ID of the host | [optional] 
**HostGroup** | Pointer to [**HostGroup**](HostGroup.md) |  | [optional] 
**Tags** | Pointer to [**[]TagInfo**](TagInfo.md) | A list of tags of the host. | [optional] 
**ManagementZones** | Pointer to [**[]EntityShortRepresentation**](EntityShortRepresentation.md) | A list of management zones to which the host belongs. | [optional] 

## Methods

### NewHost

`func NewHost() *Host`

NewHost instantiates a new Host object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostWithDefaults

`func NewHostWithDefaults() *Host`

NewHostWithDefaults instantiates a new Host object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Host) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Host) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Host) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Host) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *Host) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Host) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Host) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Host) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHostGroup

`func (o *Host) GetHostGroup() HostGroup`

GetHostGroup returns the HostGroup field if non-nil, zero value otherwise.

### GetHostGroupOk

`func (o *Host) GetHostGroupOk() (*HostGroup, bool)`

GetHostGroupOk returns a tuple with the HostGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroup

`func (o *Host) SetHostGroup(v HostGroup)`

SetHostGroup sets HostGroup field to given value.

### HasHostGroup

`func (o *Host) HasHostGroup() bool`

HasHostGroup returns a boolean if a field has been set.

### GetTags

`func (o *Host) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Host) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Host) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Host) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetManagementZones

`func (o *Host) GetManagementZones() []EntityShortRepresentation`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *Host) GetManagementZonesOk() (*[]EntityShortRepresentation, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *Host) SetManagementZones(v []EntityShortRepresentation)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *Host) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


