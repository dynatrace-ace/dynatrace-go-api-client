# ProcessGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** | The Dynatrace entity ID of the required entity. | [optional] 
**DisplayName** | Pointer to **string** | The name of the Dynatrace entity as displayed in the UI. | [optional] 
**CustomizedName** | Pointer to **string** | The customized name of the entity | [optional] 
**DiscoveredName** | Pointer to **string** | The discovered name of the entity | [optional] 
**FirstSeenTimestamp** | Pointer to **int64** | The timestamp of when the entity was first detected, in UTC milliseconds | [optional] 
**LastSeenTimestamp** | Pointer to **int64** | The timestamp of when the entity was last detected, in UTC milliseconds | [optional] 
**Tags** | Pointer to [**[]TagInfo**](TagInfo.md) | The list of entity tags. | [optional] 
**ToRelationships** | Pointer to [**ProcessGroupToRelationships**](ProcessGroup_toRelationships.md) |  | [optional] 
**Metadata** | Pointer to [**ProcessGroupMetadata**](ProcessGroup_metadata.md) |  | [optional] 
**AzureHostName** | Pointer to **string** |  | [optional] 
**SoftwareTechnologies** | Pointer to [**[]TechnologyInfo**](TechnologyInfo.md) |  | [optional] 
**ListenPorts** | Pointer to **[]int32** |  | [optional] 
**AzureSiteName** | Pointer to **string** |  | [optional] 
**ManagementZones** | Pointer to [**[]EntityShortRepresentation**](EntityShortRepresentation.md) | The management zones that the entity is part of. | [optional] 
**FromRelationships** | Pointer to [**ProcessGroupFromRelationships**](ProcessGroup_fromRelationships.md) |  | [optional] 

## Methods

### NewProcessGroup

`func NewProcessGroup() *ProcessGroup`

NewProcessGroup instantiates a new ProcessGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessGroupWithDefaults

`func NewProcessGroupWithDefaults() *ProcessGroup`

NewProcessGroupWithDefaults instantiates a new ProcessGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *ProcessGroup) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ProcessGroup) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ProcessGroup) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ProcessGroup) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetDisplayName

`func (o *ProcessGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ProcessGroup) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ProcessGroup) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ProcessGroup) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCustomizedName

`func (o *ProcessGroup) GetCustomizedName() string`

GetCustomizedName returns the CustomizedName field if non-nil, zero value otherwise.

### GetCustomizedNameOk

`func (o *ProcessGroup) GetCustomizedNameOk() (*string, bool)`

GetCustomizedNameOk returns a tuple with the CustomizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizedName

`func (o *ProcessGroup) SetCustomizedName(v string)`

SetCustomizedName sets CustomizedName field to given value.

### HasCustomizedName

`func (o *ProcessGroup) HasCustomizedName() bool`

HasCustomizedName returns a boolean if a field has been set.

### GetDiscoveredName

`func (o *ProcessGroup) GetDiscoveredName() string`

GetDiscoveredName returns the DiscoveredName field if non-nil, zero value otherwise.

### GetDiscoveredNameOk

`func (o *ProcessGroup) GetDiscoveredNameOk() (*string, bool)`

GetDiscoveredNameOk returns a tuple with the DiscoveredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredName

`func (o *ProcessGroup) SetDiscoveredName(v string)`

SetDiscoveredName sets DiscoveredName field to given value.

### HasDiscoveredName

`func (o *ProcessGroup) HasDiscoveredName() bool`

HasDiscoveredName returns a boolean if a field has been set.

### GetFirstSeenTimestamp

`func (o *ProcessGroup) GetFirstSeenTimestamp() int64`

GetFirstSeenTimestamp returns the FirstSeenTimestamp field if non-nil, zero value otherwise.

### GetFirstSeenTimestampOk

`func (o *ProcessGroup) GetFirstSeenTimestampOk() (*int64, bool)`

GetFirstSeenTimestampOk returns a tuple with the FirstSeenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenTimestamp

`func (o *ProcessGroup) SetFirstSeenTimestamp(v int64)`

SetFirstSeenTimestamp sets FirstSeenTimestamp field to given value.

### HasFirstSeenTimestamp

`func (o *ProcessGroup) HasFirstSeenTimestamp() bool`

HasFirstSeenTimestamp returns a boolean if a field has been set.

### GetLastSeenTimestamp

`func (o *ProcessGroup) GetLastSeenTimestamp() int64`

GetLastSeenTimestamp returns the LastSeenTimestamp field if non-nil, zero value otherwise.

### GetLastSeenTimestampOk

`func (o *ProcessGroup) GetLastSeenTimestampOk() (*int64, bool)`

GetLastSeenTimestampOk returns a tuple with the LastSeenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenTimestamp

`func (o *ProcessGroup) SetLastSeenTimestamp(v int64)`

SetLastSeenTimestamp sets LastSeenTimestamp field to given value.

### HasLastSeenTimestamp

`func (o *ProcessGroup) HasLastSeenTimestamp() bool`

HasLastSeenTimestamp returns a boolean if a field has been set.

### GetTags

`func (o *ProcessGroup) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProcessGroup) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProcessGroup) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProcessGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToRelationships

`func (o *ProcessGroup) GetToRelationships() ProcessGroupToRelationships`

GetToRelationships returns the ToRelationships field if non-nil, zero value otherwise.

### GetToRelationshipsOk

`func (o *ProcessGroup) GetToRelationshipsOk() (*ProcessGroupToRelationships, bool)`

GetToRelationshipsOk returns a tuple with the ToRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRelationships

`func (o *ProcessGroup) SetToRelationships(v ProcessGroupToRelationships)`

SetToRelationships sets ToRelationships field to given value.

### HasToRelationships

`func (o *ProcessGroup) HasToRelationships() bool`

HasToRelationships returns a boolean if a field has been set.

### GetMetadata

`func (o *ProcessGroup) GetMetadata() ProcessGroupMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProcessGroup) GetMetadataOk() (*ProcessGroupMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProcessGroup) SetMetadata(v ProcessGroupMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProcessGroup) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAzureHostName

`func (o *ProcessGroup) GetAzureHostName() string`

GetAzureHostName returns the AzureHostName field if non-nil, zero value otherwise.

### GetAzureHostNameOk

`func (o *ProcessGroup) GetAzureHostNameOk() (*string, bool)`

GetAzureHostNameOk returns a tuple with the AzureHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureHostName

`func (o *ProcessGroup) SetAzureHostName(v string)`

SetAzureHostName sets AzureHostName field to given value.

### HasAzureHostName

`func (o *ProcessGroup) HasAzureHostName() bool`

HasAzureHostName returns a boolean if a field has been set.

### GetSoftwareTechnologies

`func (o *ProcessGroup) GetSoftwareTechnologies() []TechnologyInfo`

GetSoftwareTechnologies returns the SoftwareTechnologies field if non-nil, zero value otherwise.

### GetSoftwareTechnologiesOk

`func (o *ProcessGroup) GetSoftwareTechnologiesOk() (*[]TechnologyInfo, bool)`

GetSoftwareTechnologiesOk returns a tuple with the SoftwareTechnologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTechnologies

`func (o *ProcessGroup) SetSoftwareTechnologies(v []TechnologyInfo)`

SetSoftwareTechnologies sets SoftwareTechnologies field to given value.

### HasSoftwareTechnologies

`func (o *ProcessGroup) HasSoftwareTechnologies() bool`

HasSoftwareTechnologies returns a boolean if a field has been set.

### GetListenPorts

`func (o *ProcessGroup) GetListenPorts() []int32`

GetListenPorts returns the ListenPorts field if non-nil, zero value otherwise.

### GetListenPortsOk

`func (o *ProcessGroup) GetListenPortsOk() (*[]int32, bool)`

GetListenPortsOk returns a tuple with the ListenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPorts

`func (o *ProcessGroup) SetListenPorts(v []int32)`

SetListenPorts sets ListenPorts field to given value.

### HasListenPorts

`func (o *ProcessGroup) HasListenPorts() bool`

HasListenPorts returns a boolean if a field has been set.

### GetAzureSiteName

`func (o *ProcessGroup) GetAzureSiteName() string`

GetAzureSiteName returns the AzureSiteName field if non-nil, zero value otherwise.

### GetAzureSiteNameOk

`func (o *ProcessGroup) GetAzureSiteNameOk() (*string, bool)`

GetAzureSiteNameOk returns a tuple with the AzureSiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSiteName

`func (o *ProcessGroup) SetAzureSiteName(v string)`

SetAzureSiteName sets AzureSiteName field to given value.

### HasAzureSiteName

`func (o *ProcessGroup) HasAzureSiteName() bool`

HasAzureSiteName returns a boolean if a field has been set.

### GetManagementZones

`func (o *ProcessGroup) GetManagementZones() []EntityShortRepresentation`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *ProcessGroup) GetManagementZonesOk() (*[]EntityShortRepresentation, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *ProcessGroup) SetManagementZones(v []EntityShortRepresentation)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *ProcessGroup) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetFromRelationships

`func (o *ProcessGroup) GetFromRelationships() ProcessGroupFromRelationships`

GetFromRelationships returns the FromRelationships field if non-nil, zero value otherwise.

### GetFromRelationshipsOk

`func (o *ProcessGroup) GetFromRelationshipsOk() (*ProcessGroupFromRelationships, bool)`

GetFromRelationshipsOk returns a tuple with the FromRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromRelationships

`func (o *ProcessGroup) SetFromRelationships(v ProcessGroupFromRelationships)`

SetFromRelationships sets FromRelationships field to given value.

### HasFromRelationships

`func (o *ProcessGroup) HasFromRelationships() bool`

HasFromRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


