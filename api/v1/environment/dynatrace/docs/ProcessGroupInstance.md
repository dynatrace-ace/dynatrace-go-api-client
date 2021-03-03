# ProcessGroupInstance

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
**ToRelationships** | Pointer to [**ProcessGroupInstanceToRelationships**](ProcessGroupInstance_toRelationships.md) |  | [optional] 
**Metadata** | Pointer to [**ProcessGroupMetadata**](ProcessGroup_metadata.md) |  | [optional] 
**AzureHostName** | Pointer to **string** |  | [optional] 
**MonitoringState** | Pointer to [**MonitoringState**](MonitoringState.md) |  | [optional] 
**Bitness** | Pointer to **string** |  | [optional] 
**SoftwareTechnologies** | Pointer to [**[]TechnologyInfo**](TechnologyInfo.md) |  | [optional] 
**Modules** | Pointer to **[]string** |  | [optional] 
**VersionedModules** | Pointer to [**[]ProcessGroupInstanceModule**](ProcessGroupInstanceModule.md) |  | [optional] 
**ListenPorts** | Pointer to **[]int32** |  | [optional] 
**AgentVersions** | Pointer to [**[]AgentVersion**](AgentVersion.md) | Versions of OneAgents currently running on the entity. | [optional] 
**AzureSiteName** | Pointer to **string** |  | [optional] 
**ManagementZones** | Pointer to [**[]EntityShortRepresentation**](EntityShortRepresentation.md) | The management zones that the entity is part of. | [optional] 
**FromRelationships** | Pointer to [**ProcessGroupInstanceFromRelationships**](ProcessGroupInstance_fromRelationships.md) |  | [optional] 

## Methods

### NewProcessGroupInstance

`func NewProcessGroupInstance() *ProcessGroupInstance`

NewProcessGroupInstance instantiates a new ProcessGroupInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessGroupInstanceWithDefaults

`func NewProcessGroupInstanceWithDefaults() *ProcessGroupInstance`

NewProcessGroupInstanceWithDefaults instantiates a new ProcessGroupInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *ProcessGroupInstance) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ProcessGroupInstance) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ProcessGroupInstance) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ProcessGroupInstance) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetDisplayName

`func (o *ProcessGroupInstance) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ProcessGroupInstance) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ProcessGroupInstance) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ProcessGroupInstance) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCustomizedName

`func (o *ProcessGroupInstance) GetCustomizedName() string`

GetCustomizedName returns the CustomizedName field if non-nil, zero value otherwise.

### GetCustomizedNameOk

`func (o *ProcessGroupInstance) GetCustomizedNameOk() (*string, bool)`

GetCustomizedNameOk returns a tuple with the CustomizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizedName

`func (o *ProcessGroupInstance) SetCustomizedName(v string)`

SetCustomizedName sets CustomizedName field to given value.

### HasCustomizedName

`func (o *ProcessGroupInstance) HasCustomizedName() bool`

HasCustomizedName returns a boolean if a field has been set.

### GetDiscoveredName

`func (o *ProcessGroupInstance) GetDiscoveredName() string`

GetDiscoveredName returns the DiscoveredName field if non-nil, zero value otherwise.

### GetDiscoveredNameOk

`func (o *ProcessGroupInstance) GetDiscoveredNameOk() (*string, bool)`

GetDiscoveredNameOk returns a tuple with the DiscoveredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredName

`func (o *ProcessGroupInstance) SetDiscoveredName(v string)`

SetDiscoveredName sets DiscoveredName field to given value.

### HasDiscoveredName

`func (o *ProcessGroupInstance) HasDiscoveredName() bool`

HasDiscoveredName returns a boolean if a field has been set.

### GetFirstSeenTimestamp

`func (o *ProcessGroupInstance) GetFirstSeenTimestamp() int64`

GetFirstSeenTimestamp returns the FirstSeenTimestamp field if non-nil, zero value otherwise.

### GetFirstSeenTimestampOk

`func (o *ProcessGroupInstance) GetFirstSeenTimestampOk() (*int64, bool)`

GetFirstSeenTimestampOk returns a tuple with the FirstSeenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenTimestamp

`func (o *ProcessGroupInstance) SetFirstSeenTimestamp(v int64)`

SetFirstSeenTimestamp sets FirstSeenTimestamp field to given value.

### HasFirstSeenTimestamp

`func (o *ProcessGroupInstance) HasFirstSeenTimestamp() bool`

HasFirstSeenTimestamp returns a boolean if a field has been set.

### GetLastSeenTimestamp

`func (o *ProcessGroupInstance) GetLastSeenTimestamp() int64`

GetLastSeenTimestamp returns the LastSeenTimestamp field if non-nil, zero value otherwise.

### GetLastSeenTimestampOk

`func (o *ProcessGroupInstance) GetLastSeenTimestampOk() (*int64, bool)`

GetLastSeenTimestampOk returns a tuple with the LastSeenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenTimestamp

`func (o *ProcessGroupInstance) SetLastSeenTimestamp(v int64)`

SetLastSeenTimestamp sets LastSeenTimestamp field to given value.

### HasLastSeenTimestamp

`func (o *ProcessGroupInstance) HasLastSeenTimestamp() bool`

HasLastSeenTimestamp returns a boolean if a field has been set.

### GetTags

`func (o *ProcessGroupInstance) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProcessGroupInstance) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProcessGroupInstance) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProcessGroupInstance) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToRelationships

`func (o *ProcessGroupInstance) GetToRelationships() ProcessGroupInstanceToRelationships`

GetToRelationships returns the ToRelationships field if non-nil, zero value otherwise.

### GetToRelationshipsOk

`func (o *ProcessGroupInstance) GetToRelationshipsOk() (*ProcessGroupInstanceToRelationships, bool)`

GetToRelationshipsOk returns a tuple with the ToRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRelationships

`func (o *ProcessGroupInstance) SetToRelationships(v ProcessGroupInstanceToRelationships)`

SetToRelationships sets ToRelationships field to given value.

### HasToRelationships

`func (o *ProcessGroupInstance) HasToRelationships() bool`

HasToRelationships returns a boolean if a field has been set.

### GetMetadata

`func (o *ProcessGroupInstance) GetMetadata() ProcessGroupMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProcessGroupInstance) GetMetadataOk() (*ProcessGroupMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProcessGroupInstance) SetMetadata(v ProcessGroupMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProcessGroupInstance) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAzureHostName

`func (o *ProcessGroupInstance) GetAzureHostName() string`

GetAzureHostName returns the AzureHostName field if non-nil, zero value otherwise.

### GetAzureHostNameOk

`func (o *ProcessGroupInstance) GetAzureHostNameOk() (*string, bool)`

GetAzureHostNameOk returns a tuple with the AzureHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureHostName

`func (o *ProcessGroupInstance) SetAzureHostName(v string)`

SetAzureHostName sets AzureHostName field to given value.

### HasAzureHostName

`func (o *ProcessGroupInstance) HasAzureHostName() bool`

HasAzureHostName returns a boolean if a field has been set.

### GetMonitoringState

`func (o *ProcessGroupInstance) GetMonitoringState() MonitoringState`

GetMonitoringState returns the MonitoringState field if non-nil, zero value otherwise.

### GetMonitoringStateOk

`func (o *ProcessGroupInstance) GetMonitoringStateOk() (*MonitoringState, bool)`

GetMonitoringStateOk returns a tuple with the MonitoringState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringState

`func (o *ProcessGroupInstance) SetMonitoringState(v MonitoringState)`

SetMonitoringState sets MonitoringState field to given value.

### HasMonitoringState

`func (o *ProcessGroupInstance) HasMonitoringState() bool`

HasMonitoringState returns a boolean if a field has been set.

### GetBitness

`func (o *ProcessGroupInstance) GetBitness() string`

GetBitness returns the Bitness field if non-nil, zero value otherwise.

### GetBitnessOk

`func (o *ProcessGroupInstance) GetBitnessOk() (*string, bool)`

GetBitnessOk returns a tuple with the Bitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitness

`func (o *ProcessGroupInstance) SetBitness(v string)`

SetBitness sets Bitness field to given value.

### HasBitness

`func (o *ProcessGroupInstance) HasBitness() bool`

HasBitness returns a boolean if a field has been set.

### GetSoftwareTechnologies

`func (o *ProcessGroupInstance) GetSoftwareTechnologies() []TechnologyInfo`

GetSoftwareTechnologies returns the SoftwareTechnologies field if non-nil, zero value otherwise.

### GetSoftwareTechnologiesOk

`func (o *ProcessGroupInstance) GetSoftwareTechnologiesOk() (*[]TechnologyInfo, bool)`

GetSoftwareTechnologiesOk returns a tuple with the SoftwareTechnologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTechnologies

`func (o *ProcessGroupInstance) SetSoftwareTechnologies(v []TechnologyInfo)`

SetSoftwareTechnologies sets SoftwareTechnologies field to given value.

### HasSoftwareTechnologies

`func (o *ProcessGroupInstance) HasSoftwareTechnologies() bool`

HasSoftwareTechnologies returns a boolean if a field has been set.

### GetModules

`func (o *ProcessGroupInstance) GetModules() []string`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ProcessGroupInstance) GetModulesOk() (*[]string, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ProcessGroupInstance) SetModules(v []string)`

SetModules sets Modules field to given value.

### HasModules

`func (o *ProcessGroupInstance) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetVersionedModules

`func (o *ProcessGroupInstance) GetVersionedModules() []ProcessGroupInstanceModule`

GetVersionedModules returns the VersionedModules field if non-nil, zero value otherwise.

### GetVersionedModulesOk

`func (o *ProcessGroupInstance) GetVersionedModulesOk() (*[]ProcessGroupInstanceModule, bool)`

GetVersionedModulesOk returns a tuple with the VersionedModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedModules

`func (o *ProcessGroupInstance) SetVersionedModules(v []ProcessGroupInstanceModule)`

SetVersionedModules sets VersionedModules field to given value.

### HasVersionedModules

`func (o *ProcessGroupInstance) HasVersionedModules() bool`

HasVersionedModules returns a boolean if a field has been set.

### GetListenPorts

`func (o *ProcessGroupInstance) GetListenPorts() []int32`

GetListenPorts returns the ListenPorts field if non-nil, zero value otherwise.

### GetListenPortsOk

`func (o *ProcessGroupInstance) GetListenPortsOk() (*[]int32, bool)`

GetListenPortsOk returns a tuple with the ListenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPorts

`func (o *ProcessGroupInstance) SetListenPorts(v []int32)`

SetListenPorts sets ListenPorts field to given value.

### HasListenPorts

`func (o *ProcessGroupInstance) HasListenPorts() bool`

HasListenPorts returns a boolean if a field has been set.

### GetAgentVersions

`func (o *ProcessGroupInstance) GetAgentVersions() []AgentVersion`

GetAgentVersions returns the AgentVersions field if non-nil, zero value otherwise.

### GetAgentVersionsOk

`func (o *ProcessGroupInstance) GetAgentVersionsOk() (*[]AgentVersion, bool)`

GetAgentVersionsOk returns a tuple with the AgentVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersions

`func (o *ProcessGroupInstance) SetAgentVersions(v []AgentVersion)`

SetAgentVersions sets AgentVersions field to given value.

### HasAgentVersions

`func (o *ProcessGroupInstance) HasAgentVersions() bool`

HasAgentVersions returns a boolean if a field has been set.

### GetAzureSiteName

`func (o *ProcessGroupInstance) GetAzureSiteName() string`

GetAzureSiteName returns the AzureSiteName field if non-nil, zero value otherwise.

### GetAzureSiteNameOk

`func (o *ProcessGroupInstance) GetAzureSiteNameOk() (*string, bool)`

GetAzureSiteNameOk returns a tuple with the AzureSiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSiteName

`func (o *ProcessGroupInstance) SetAzureSiteName(v string)`

SetAzureSiteName sets AzureSiteName field to given value.

### HasAzureSiteName

`func (o *ProcessGroupInstance) HasAzureSiteName() bool`

HasAzureSiteName returns a boolean if a field has been set.

### GetManagementZones

`func (o *ProcessGroupInstance) GetManagementZones() []EntityShortRepresentation`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *ProcessGroupInstance) GetManagementZonesOk() (*[]EntityShortRepresentation, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *ProcessGroupInstance) SetManagementZones(v []EntityShortRepresentation)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *ProcessGroupInstance) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetFromRelationships

`func (o *ProcessGroupInstance) GetFromRelationships() ProcessGroupInstanceFromRelationships`

GetFromRelationships returns the FromRelationships field if non-nil, zero value otherwise.

### GetFromRelationshipsOk

`func (o *ProcessGroupInstance) GetFromRelationshipsOk() (*ProcessGroupInstanceFromRelationships, bool)`

GetFromRelationshipsOk returns a tuple with the FromRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromRelationships

`func (o *ProcessGroupInstance) SetFromRelationships(v ProcessGroupInstanceFromRelationships)`

SetFromRelationships sets FromRelationships field to given value.

### HasFromRelationships

`func (o *ProcessGroupInstance) HasFromRelationships() bool`

HasFromRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


