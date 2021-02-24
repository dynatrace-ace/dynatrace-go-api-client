# DashboardMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the dashboard. | 
**Shared** | Pointer to **bool** | The dashboard is shared (&#x60;true&#x60;) or private (&#x60;false&#x60;). | [optional] 
**Owner** | Pointer to **string** | The owner of the dashboard. | [optional] 
**SharingDetails** | Pointer to [**SharingInfo**](SharingInfo.md) |  | [optional] 
**DashboardFilter** | Pointer to [**DashboardFilter**](DashboardFilter.md) |  | [optional] 
**Tags** | Pointer to **[]string** | A set of tags assigned to the dashboard. | [optional] 
**Preset** | Pointer to **bool** | The dashboard is a preset (&#x60;true&#x60;) | [optional] [readonly] 
**ValidFilterKeys** | Pointer to **[]string** | A set of all possible global dashboard filters that can be applied to dashboard | [optional] 

## Methods

### NewDashboardMetadata

`func NewDashboardMetadata(name string, ) *DashboardMetadata`

NewDashboardMetadata instantiates a new DashboardMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardMetadataWithDefaults

`func NewDashboardMetadataWithDefaults() *DashboardMetadata`

NewDashboardMetadataWithDefaults instantiates a new DashboardMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DashboardMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetShared

`func (o *DashboardMetadata) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *DashboardMetadata) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *DashboardMetadata) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *DashboardMetadata) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetOwner

`func (o *DashboardMetadata) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DashboardMetadata) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DashboardMetadata) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DashboardMetadata) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSharingDetails

`func (o *DashboardMetadata) GetSharingDetails() SharingInfo`

GetSharingDetails returns the SharingDetails field if non-nil, zero value otherwise.

### GetSharingDetailsOk

`func (o *DashboardMetadata) GetSharingDetailsOk() (*SharingInfo, bool)`

GetSharingDetailsOk returns a tuple with the SharingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingDetails

`func (o *DashboardMetadata) SetSharingDetails(v SharingInfo)`

SetSharingDetails sets SharingDetails field to given value.

### HasSharingDetails

`func (o *DashboardMetadata) HasSharingDetails() bool`

HasSharingDetails returns a boolean if a field has been set.

### GetDashboardFilter

`func (o *DashboardMetadata) GetDashboardFilter() DashboardFilter`

GetDashboardFilter returns the DashboardFilter field if non-nil, zero value otherwise.

### GetDashboardFilterOk

`func (o *DashboardMetadata) GetDashboardFilterOk() (*DashboardFilter, bool)`

GetDashboardFilterOk returns a tuple with the DashboardFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardFilter

`func (o *DashboardMetadata) SetDashboardFilter(v DashboardFilter)`

SetDashboardFilter sets DashboardFilter field to given value.

### HasDashboardFilter

`func (o *DashboardMetadata) HasDashboardFilter() bool`

HasDashboardFilter returns a boolean if a field has been set.

### GetTags

`func (o *DashboardMetadata) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DashboardMetadata) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DashboardMetadata) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DashboardMetadata) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPreset

`func (o *DashboardMetadata) GetPreset() bool`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *DashboardMetadata) GetPresetOk() (*bool, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *DashboardMetadata) SetPreset(v bool)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *DashboardMetadata) HasPreset() bool`

HasPreset returns a boolean if a field has been set.

### GetValidFilterKeys

`func (o *DashboardMetadata) GetValidFilterKeys() []string`

GetValidFilterKeys returns the ValidFilterKeys field if non-nil, zero value otherwise.

### GetValidFilterKeysOk

`func (o *DashboardMetadata) GetValidFilterKeysOk() (*[]string, bool)`

GetValidFilterKeysOk returns a tuple with the ValidFilterKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFilterKeys

`func (o *DashboardMetadata) SetValidFilterKeys(v []string)`

SetValidFilterKeys sets ValidFilterKeys field to given value.

### HasValidFilterKeys

`func (o *DashboardMetadata) HasValidFilterKeys() bool`

HasValidFilterKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


