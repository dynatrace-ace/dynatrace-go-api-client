# Host

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
**ToRelationships** | Pointer to [**HostToRelationships**](Host_toRelationships.md) |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**LogicalCpus** | Pointer to **int32** | The AIX instance logical CPU count. | [optional] 
**AzureVmName** | Pointer to **string** |  | [optional] 
**MonitoringMode** | Pointer to **string** |  | [optional] 
**LogicalCpuCores** | Pointer to **int32** |  | [optional] 
**OpenstackAvZone** | Pointer to **string** |  | [optional] 
**GceProjectId** | Pointer to **string** | The Google Compute Engine numeric project ID. | [optional] 
**AzureComputeModeName** | Pointer to **string** |  | [optional] 
**OsType** | Pointer to **string** |  | [optional] 
**AutoScalingGroup** | Pointer to **string** |  | [optional] 
**UserLevel** | Pointer to **string** |  | [optional] 
**AwsSecurityGroup** | Pointer to **[]string** |  | [optional] 
**IsMonitoringCandidate** | Pointer to **bool** |  | [optional] 
**GceProject** | Pointer to **string** | The Google Compute Engine project. | [optional] 
**AmiId** | Pointer to **string** |  | [optional] 
**ZosCPUModelNumber** | Pointer to **string** | The CPU model number. | [optional] 
**ZosTotalZiipProcessors** | Pointer to **int32** | Number of assigned support processors for this LPAR. | [optional] 
**LocalHostName** | Pointer to **string** |  | [optional] 
**EsxiHostName** | Pointer to **string** |  | [optional] 
**OpenstackProjectName** | Pointer to **string** |  | [optional] 
**NetworkZoneId** | Pointer to **string** | The ID of network zone the entity is in. | [optional] 
**OsArchitecture** | Pointer to **string** |  | [optional] 
**AgentVersion** | Pointer to [**AgentVersion**](AgentVersion.md) |  | [optional] 
**BoshDeploymentId** | Pointer to **string** | The Cloud Foundry BOSH deployment ID. | [optional] 
**AwsInstanceId** | Pointer to **string** |  | [optional] 
**PaasMemoryLimit** | Pointer to **int64** |  | [optional] 
**BoshInstanceName** | Pointer to **string** | The Cloud Foundry BOSH instance name. | [optional] 
**Bitness** | Pointer to **string** |  | [optional] 
**KubernetesLabels** | Pointer to **map[string]map[string]interface{}** | The kubernetes labels defined on the entity. | [optional] 
**AwsNameTag** | Pointer to **string** | The name inherited from AWS. | [optional] 
**GceInstanceId** | Pointer to **string** | The Google Compute Engine instance ID. | [optional] 
**BeanstalkEnvironmentName** | Pointer to **string** |  | [optional] 
**GcePublicIpAddresses** | Pointer to **[]string** | The public IP addresses of the Google Compute Engine. | [optional] 
**ScaleSetName** | Pointer to **string** |  | [optional] 
**OneAgentCustomHostName** | Pointer to **string** | The custom name defined in OneAgent config. | [optional] 
**AzureHostNames** | Pointer to **[]string** |  | [optional] 
**GceInstanceName** | Pointer to **string** | The Google Compute Engine instance name. | [optional] 
**PaasType** | Pointer to **string** |  | [optional] 
**BoshInstanceId** | Pointer to **string** | The Cloud Foundry BOSH instance ID. | [optional] 
**AzureSku** | Pointer to **string** |  | [optional] 
**CpuCores** | Pointer to **int32** |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**ManagementZones** | Pointer to [**[]EntityShortRepresentation**](EntityShortRepresentation.md) | The management zones that the entity is part of. | [optional] 
**BoshName** | Pointer to **string** | The Cloud Foundry BOSH name. | [optional] 
**ZosTotalPhysicalMemory** | Pointer to **int32** | Memory assigned to the host (Terabyte). | [optional] 
**OpenstackVmName** | Pointer to **string** |  | [optional] 
**ZosVirtualization** | Pointer to **string** | Type of virtualization on the mainframe. | [optional] 
**VirtualCpus** | Pointer to **int32** | The AIX instance virtual CPU count. | [optional] 
**BoshAvailabilityZone** | Pointer to **string** | The Cloud Foundry BOSH availability zone. | [optional] 
**CloudType** | Pointer to **string** |  | [optional] 
**PublicHostName** | Pointer to **string** |  | [optional] 
**ConsumedHostUnits** | Pointer to **string** |  | [optional] 
**HostGroup** | Pointer to [**HostGroup**](HostGroup.md) |  | [optional] 
**SimultaneousMultithreading** | Pointer to **int32** | The AIX instance simultaneous threads count. | [optional] 
**BoshStemcellVersion** | Pointer to **string** | The Cloud Foundry BOSH stemcell version. | [optional] 
**ZosTotalGeneralPurposeProcessors** | Pointer to **int32** | Number of assigned processors for this LPAR. | [optional] 
**OpenstackComputeNodeName** | Pointer to **string** |  | [optional] 
**KubernetesCluster** | Pointer to **string** | The kubernetes cluster the entity is in. | [optional] 
**SoftwareTechnologies** | Pointer to [**[]TechnologyInfo**](TechnologyInfo.md) |  | [optional] 
**GceMachineType** | Pointer to **string** | The Google Compute Engine machine type. | [optional] 
**ZosCPUSerialNumber** | Pointer to **string** | The CPU serial number. | [optional] 
**PaasAgentVersions** | Pointer to [**[]AgentVersion**](AgentVersion.md) | The versions of the PaaS agents currently running on the entity. | [optional] 
**OpenStackInstaceType** | Pointer to **string** |  | [optional] 
**AwsInstanceType** | Pointer to **string** |  | [optional] 
**AzureSiteNames** | Pointer to **[]string** |  | [optional] 
**OpenstackSecurityGroups** | Pointer to **[]string** |  | [optional] 
**VmwareName** | Pointer to **string** |  | [optional] 
**CloudPlatformVendorVersion** | Pointer to **string** | Defines the cloud platform vendor version. | [optional] 
**KubernetesNode** | Pointer to **string** | The kubernetes node the entity is in. | [optional] 
**GcpZone** | Pointer to **string** | The Google Cloud Platform Zone. | [optional] 
**HypervisorType** | Pointer to **string** |  | [optional] 
**IpAddresses** | Pointer to **[]string** |  | [optional] 
**LocalIp** | Pointer to **string** |  | [optional] 
**FromRelationships** | Pointer to [**HostFromRelationships**](Host_fromRelationships.md) |  | [optional] 

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

### GetEntityId

`func (o *Host) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Host) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Host) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *Host) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetDisplayName

`func (o *Host) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Host) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Host) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Host) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCustomizedName

`func (o *Host) GetCustomizedName() string`

GetCustomizedName returns the CustomizedName field if non-nil, zero value otherwise.

### GetCustomizedNameOk

`func (o *Host) GetCustomizedNameOk() (*string, bool)`

GetCustomizedNameOk returns a tuple with the CustomizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizedName

`func (o *Host) SetCustomizedName(v string)`

SetCustomizedName sets CustomizedName field to given value.

### HasCustomizedName

`func (o *Host) HasCustomizedName() bool`

HasCustomizedName returns a boolean if a field has been set.

### GetDiscoveredName

`func (o *Host) GetDiscoveredName() string`

GetDiscoveredName returns the DiscoveredName field if non-nil, zero value otherwise.

### GetDiscoveredNameOk

`func (o *Host) GetDiscoveredNameOk() (*string, bool)`

GetDiscoveredNameOk returns a tuple with the DiscoveredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredName

`func (o *Host) SetDiscoveredName(v string)`

SetDiscoveredName sets DiscoveredName field to given value.

### HasDiscoveredName

`func (o *Host) HasDiscoveredName() bool`

HasDiscoveredName returns a boolean if a field has been set.

### GetFirstSeenTimestamp

`func (o *Host) GetFirstSeenTimestamp() int64`

GetFirstSeenTimestamp returns the FirstSeenTimestamp field if non-nil, zero value otherwise.

### GetFirstSeenTimestampOk

`func (o *Host) GetFirstSeenTimestampOk() (*int64, bool)`

GetFirstSeenTimestampOk returns a tuple with the FirstSeenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenTimestamp

`func (o *Host) SetFirstSeenTimestamp(v int64)`

SetFirstSeenTimestamp sets FirstSeenTimestamp field to given value.

### HasFirstSeenTimestamp

`func (o *Host) HasFirstSeenTimestamp() bool`

HasFirstSeenTimestamp returns a boolean if a field has been set.

### GetLastSeenTimestamp

`func (o *Host) GetLastSeenTimestamp() int64`

GetLastSeenTimestamp returns the LastSeenTimestamp field if non-nil, zero value otherwise.

### GetLastSeenTimestampOk

`func (o *Host) GetLastSeenTimestampOk() (*int64, bool)`

GetLastSeenTimestampOk returns a tuple with the LastSeenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenTimestamp

`func (o *Host) SetLastSeenTimestamp(v int64)`

SetLastSeenTimestamp sets LastSeenTimestamp field to given value.

### HasLastSeenTimestamp

`func (o *Host) HasLastSeenTimestamp() bool`

HasLastSeenTimestamp returns a boolean if a field has been set.

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

### GetToRelationships

`func (o *Host) GetToRelationships() HostToRelationships`

GetToRelationships returns the ToRelationships field if non-nil, zero value otherwise.

### GetToRelationshipsOk

`func (o *Host) GetToRelationshipsOk() (*HostToRelationships, bool)`

GetToRelationshipsOk returns a tuple with the ToRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRelationships

`func (o *Host) SetToRelationships(v HostToRelationships)`

SetToRelationships sets ToRelationships field to given value.

### HasToRelationships

`func (o *Host) HasToRelationships() bool`

HasToRelationships returns a boolean if a field has been set.

### GetOsVersion

`func (o *Host) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *Host) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *Host) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *Host) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetLogicalCpus

`func (o *Host) GetLogicalCpus() int32`

GetLogicalCpus returns the LogicalCpus field if non-nil, zero value otherwise.

### GetLogicalCpusOk

`func (o *Host) GetLogicalCpusOk() (*int32, bool)`

GetLogicalCpusOk returns a tuple with the LogicalCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalCpus

`func (o *Host) SetLogicalCpus(v int32)`

SetLogicalCpus sets LogicalCpus field to given value.

### HasLogicalCpus

`func (o *Host) HasLogicalCpus() bool`

HasLogicalCpus returns a boolean if a field has been set.

### GetAzureVmName

`func (o *Host) GetAzureVmName() string`

GetAzureVmName returns the AzureVmName field if non-nil, zero value otherwise.

### GetAzureVmNameOk

`func (o *Host) GetAzureVmNameOk() (*string, bool)`

GetAzureVmNameOk returns a tuple with the AzureVmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVmName

`func (o *Host) SetAzureVmName(v string)`

SetAzureVmName sets AzureVmName field to given value.

### HasAzureVmName

`func (o *Host) HasAzureVmName() bool`

HasAzureVmName returns a boolean if a field has been set.

### GetMonitoringMode

`func (o *Host) GetMonitoringMode() string`

GetMonitoringMode returns the MonitoringMode field if non-nil, zero value otherwise.

### GetMonitoringModeOk

`func (o *Host) GetMonitoringModeOk() (*string, bool)`

GetMonitoringModeOk returns a tuple with the MonitoringMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringMode

`func (o *Host) SetMonitoringMode(v string)`

SetMonitoringMode sets MonitoringMode field to given value.

### HasMonitoringMode

`func (o *Host) HasMonitoringMode() bool`

HasMonitoringMode returns a boolean if a field has been set.

### GetLogicalCpuCores

`func (o *Host) GetLogicalCpuCores() int32`

GetLogicalCpuCores returns the LogicalCpuCores field if non-nil, zero value otherwise.

### GetLogicalCpuCoresOk

`func (o *Host) GetLogicalCpuCoresOk() (*int32, bool)`

GetLogicalCpuCoresOk returns a tuple with the LogicalCpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalCpuCores

`func (o *Host) SetLogicalCpuCores(v int32)`

SetLogicalCpuCores sets LogicalCpuCores field to given value.

### HasLogicalCpuCores

`func (o *Host) HasLogicalCpuCores() bool`

HasLogicalCpuCores returns a boolean if a field has been set.

### GetOpenstackAvZone

`func (o *Host) GetOpenstackAvZone() string`

GetOpenstackAvZone returns the OpenstackAvZone field if non-nil, zero value otherwise.

### GetOpenstackAvZoneOk

`func (o *Host) GetOpenstackAvZoneOk() (*string, bool)`

GetOpenstackAvZoneOk returns a tuple with the OpenstackAvZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackAvZone

`func (o *Host) SetOpenstackAvZone(v string)`

SetOpenstackAvZone sets OpenstackAvZone field to given value.

### HasOpenstackAvZone

`func (o *Host) HasOpenstackAvZone() bool`

HasOpenstackAvZone returns a boolean if a field has been set.

### GetGceProjectId

`func (o *Host) GetGceProjectId() string`

GetGceProjectId returns the GceProjectId field if non-nil, zero value otherwise.

### GetGceProjectIdOk

`func (o *Host) GetGceProjectIdOk() (*string, bool)`

GetGceProjectIdOk returns a tuple with the GceProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGceProjectId

`func (o *Host) SetGceProjectId(v string)`

SetGceProjectId sets GceProjectId field to given value.

### HasGceProjectId

`func (o *Host) HasGceProjectId() bool`

HasGceProjectId returns a boolean if a field has been set.

### GetAzureComputeModeName

`func (o *Host) GetAzureComputeModeName() string`

GetAzureComputeModeName returns the AzureComputeModeName field if non-nil, zero value otherwise.

### GetAzureComputeModeNameOk

`func (o *Host) GetAzureComputeModeNameOk() (*string, bool)`

GetAzureComputeModeNameOk returns a tuple with the AzureComputeModeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureComputeModeName

`func (o *Host) SetAzureComputeModeName(v string)`

SetAzureComputeModeName sets AzureComputeModeName field to given value.

### HasAzureComputeModeName

`func (o *Host) HasAzureComputeModeName() bool`

HasAzureComputeModeName returns a boolean if a field has been set.

### GetOsType

`func (o *Host) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *Host) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *Host) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *Host) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetAutoScalingGroup

`func (o *Host) GetAutoScalingGroup() string`

GetAutoScalingGroup returns the AutoScalingGroup field if non-nil, zero value otherwise.

### GetAutoScalingGroupOk

`func (o *Host) GetAutoScalingGroupOk() (*string, bool)`

GetAutoScalingGroupOk returns a tuple with the AutoScalingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalingGroup

`func (o *Host) SetAutoScalingGroup(v string)`

SetAutoScalingGroup sets AutoScalingGroup field to given value.

### HasAutoScalingGroup

`func (o *Host) HasAutoScalingGroup() bool`

HasAutoScalingGroup returns a boolean if a field has been set.

### GetUserLevel

`func (o *Host) GetUserLevel() string`

GetUserLevel returns the UserLevel field if non-nil, zero value otherwise.

### GetUserLevelOk

`func (o *Host) GetUserLevelOk() (*string, bool)`

GetUserLevelOk returns a tuple with the UserLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLevel

`func (o *Host) SetUserLevel(v string)`

SetUserLevel sets UserLevel field to given value.

### HasUserLevel

`func (o *Host) HasUserLevel() bool`

HasUserLevel returns a boolean if a field has been set.

### GetAwsSecurityGroup

`func (o *Host) GetAwsSecurityGroup() []string`

GetAwsSecurityGroup returns the AwsSecurityGroup field if non-nil, zero value otherwise.

### GetAwsSecurityGroupOk

`func (o *Host) GetAwsSecurityGroupOk() (*[]string, bool)`

GetAwsSecurityGroupOk returns a tuple with the AwsSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecurityGroup

`func (o *Host) SetAwsSecurityGroup(v []string)`

SetAwsSecurityGroup sets AwsSecurityGroup field to given value.

### HasAwsSecurityGroup

`func (o *Host) HasAwsSecurityGroup() bool`

HasAwsSecurityGroup returns a boolean if a field has been set.

### GetIsMonitoringCandidate

`func (o *Host) GetIsMonitoringCandidate() bool`

GetIsMonitoringCandidate returns the IsMonitoringCandidate field if non-nil, zero value otherwise.

### GetIsMonitoringCandidateOk

`func (o *Host) GetIsMonitoringCandidateOk() (*bool, bool)`

GetIsMonitoringCandidateOk returns a tuple with the IsMonitoringCandidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonitoringCandidate

`func (o *Host) SetIsMonitoringCandidate(v bool)`

SetIsMonitoringCandidate sets IsMonitoringCandidate field to given value.

### HasIsMonitoringCandidate

`func (o *Host) HasIsMonitoringCandidate() bool`

HasIsMonitoringCandidate returns a boolean if a field has been set.

### GetGceProject

`func (o *Host) GetGceProject() string`

GetGceProject returns the GceProject field if non-nil, zero value otherwise.

### GetGceProjectOk

`func (o *Host) GetGceProjectOk() (*string, bool)`

GetGceProjectOk returns a tuple with the GceProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGceProject

`func (o *Host) SetGceProject(v string)`

SetGceProject sets GceProject field to given value.

### HasGceProject

`func (o *Host) HasGceProject() bool`

HasGceProject returns a boolean if a field has been set.

### GetAmiId

`func (o *Host) GetAmiId() string`

GetAmiId returns the AmiId field if non-nil, zero value otherwise.

### GetAmiIdOk

`func (o *Host) GetAmiIdOk() (*string, bool)`

GetAmiIdOk returns a tuple with the AmiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmiId

`func (o *Host) SetAmiId(v string)`

SetAmiId sets AmiId field to given value.

### HasAmiId

`func (o *Host) HasAmiId() bool`

HasAmiId returns a boolean if a field has been set.

### GetZosCPUModelNumber

`func (o *Host) GetZosCPUModelNumber() string`

GetZosCPUModelNumber returns the ZosCPUModelNumber field if non-nil, zero value otherwise.

### GetZosCPUModelNumberOk

`func (o *Host) GetZosCPUModelNumberOk() (*string, bool)`

GetZosCPUModelNumberOk returns a tuple with the ZosCPUModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZosCPUModelNumber

`func (o *Host) SetZosCPUModelNumber(v string)`

SetZosCPUModelNumber sets ZosCPUModelNumber field to given value.

### HasZosCPUModelNumber

`func (o *Host) HasZosCPUModelNumber() bool`

HasZosCPUModelNumber returns a boolean if a field has been set.

### GetZosTotalZiipProcessors

`func (o *Host) GetZosTotalZiipProcessors() int32`

GetZosTotalZiipProcessors returns the ZosTotalZiipProcessors field if non-nil, zero value otherwise.

### GetZosTotalZiipProcessorsOk

`func (o *Host) GetZosTotalZiipProcessorsOk() (*int32, bool)`

GetZosTotalZiipProcessorsOk returns a tuple with the ZosTotalZiipProcessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZosTotalZiipProcessors

`func (o *Host) SetZosTotalZiipProcessors(v int32)`

SetZosTotalZiipProcessors sets ZosTotalZiipProcessors field to given value.

### HasZosTotalZiipProcessors

`func (o *Host) HasZosTotalZiipProcessors() bool`

HasZosTotalZiipProcessors returns a boolean if a field has been set.

### GetLocalHostName

`func (o *Host) GetLocalHostName() string`

GetLocalHostName returns the LocalHostName field if non-nil, zero value otherwise.

### GetLocalHostNameOk

`func (o *Host) GetLocalHostNameOk() (*string, bool)`

GetLocalHostNameOk returns a tuple with the LocalHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalHostName

`func (o *Host) SetLocalHostName(v string)`

SetLocalHostName sets LocalHostName field to given value.

### HasLocalHostName

`func (o *Host) HasLocalHostName() bool`

HasLocalHostName returns a boolean if a field has been set.

### GetEsxiHostName

`func (o *Host) GetEsxiHostName() string`

GetEsxiHostName returns the EsxiHostName field if non-nil, zero value otherwise.

### GetEsxiHostNameOk

`func (o *Host) GetEsxiHostNameOk() (*string, bool)`

GetEsxiHostNameOk returns a tuple with the EsxiHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsxiHostName

`func (o *Host) SetEsxiHostName(v string)`

SetEsxiHostName sets EsxiHostName field to given value.

### HasEsxiHostName

`func (o *Host) HasEsxiHostName() bool`

HasEsxiHostName returns a boolean if a field has been set.

### GetOpenstackProjectName

`func (o *Host) GetOpenstackProjectName() string`

GetOpenstackProjectName returns the OpenstackProjectName field if non-nil, zero value otherwise.

### GetOpenstackProjectNameOk

`func (o *Host) GetOpenstackProjectNameOk() (*string, bool)`

GetOpenstackProjectNameOk returns a tuple with the OpenstackProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackProjectName

`func (o *Host) SetOpenstackProjectName(v string)`

SetOpenstackProjectName sets OpenstackProjectName field to given value.

### HasOpenstackProjectName

`func (o *Host) HasOpenstackProjectName() bool`

HasOpenstackProjectName returns a boolean if a field has been set.

### GetNetworkZoneId

`func (o *Host) GetNetworkZoneId() string`

GetNetworkZoneId returns the NetworkZoneId field if non-nil, zero value otherwise.

### GetNetworkZoneIdOk

`func (o *Host) GetNetworkZoneIdOk() (*string, bool)`

GetNetworkZoneIdOk returns a tuple with the NetworkZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkZoneId

`func (o *Host) SetNetworkZoneId(v string)`

SetNetworkZoneId sets NetworkZoneId field to given value.

### HasNetworkZoneId

`func (o *Host) HasNetworkZoneId() bool`

HasNetworkZoneId returns a boolean if a field has been set.

### GetOsArchitecture

`func (o *Host) GetOsArchitecture() string`

GetOsArchitecture returns the OsArchitecture field if non-nil, zero value otherwise.

### GetOsArchitectureOk

`func (o *Host) GetOsArchitectureOk() (*string, bool)`

GetOsArchitectureOk returns a tuple with the OsArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsArchitecture

`func (o *Host) SetOsArchitecture(v string)`

SetOsArchitecture sets OsArchitecture field to given value.

### HasOsArchitecture

`func (o *Host) HasOsArchitecture() bool`

HasOsArchitecture returns a boolean if a field has been set.

### GetAgentVersion

`func (o *Host) GetAgentVersion() AgentVersion`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *Host) GetAgentVersionOk() (*AgentVersion, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *Host) SetAgentVersion(v AgentVersion)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *Host) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetBoshDeploymentId

`func (o *Host) GetBoshDeploymentId() string`

GetBoshDeploymentId returns the BoshDeploymentId field if non-nil, zero value otherwise.

### GetBoshDeploymentIdOk

`func (o *Host) GetBoshDeploymentIdOk() (*string, bool)`

GetBoshDeploymentIdOk returns a tuple with the BoshDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoshDeploymentId

`func (o *Host) SetBoshDeploymentId(v string)`

SetBoshDeploymentId sets BoshDeploymentId field to given value.

### HasBoshDeploymentId

`func (o *Host) HasBoshDeploymentId() bool`

HasBoshDeploymentId returns a boolean if a field has been set.

### GetAwsInstanceId

`func (o *Host) GetAwsInstanceId() string`

GetAwsInstanceId returns the AwsInstanceId field if non-nil, zero value otherwise.

### GetAwsInstanceIdOk

`func (o *Host) GetAwsInstanceIdOk() (*string, bool)`

GetAwsInstanceIdOk returns a tuple with the AwsInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsInstanceId

`func (o *Host) SetAwsInstanceId(v string)`

SetAwsInstanceId sets AwsInstanceId field to given value.

### HasAwsInstanceId

`func (o *Host) HasAwsInstanceId() bool`

HasAwsInstanceId returns a boolean if a field has been set.

### GetPaasMemoryLimit

`func (o *Host) GetPaasMemoryLimit() int64`

GetPaasMemoryLimit returns the PaasMemoryLimit field if non-nil, zero value otherwise.

### GetPaasMemoryLimitOk

`func (o *Host) GetPaasMemoryLimitOk() (*int64, bool)`

GetPaasMemoryLimitOk returns a tuple with the PaasMemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaasMemoryLimit

`func (o *Host) SetPaasMemoryLimit(v int64)`

SetPaasMemoryLimit sets PaasMemoryLimit field to given value.

### HasPaasMemoryLimit

`func (o *Host) HasPaasMemoryLimit() bool`

HasPaasMemoryLimit returns a boolean if a field has been set.

### GetBoshInstanceName

`func (o *Host) GetBoshInstanceName() string`

GetBoshInstanceName returns the BoshInstanceName field if non-nil, zero value otherwise.

### GetBoshInstanceNameOk

`func (o *Host) GetBoshInstanceNameOk() (*string, bool)`

GetBoshInstanceNameOk returns a tuple with the BoshInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoshInstanceName

`func (o *Host) SetBoshInstanceName(v string)`

SetBoshInstanceName sets BoshInstanceName field to given value.

### HasBoshInstanceName

`func (o *Host) HasBoshInstanceName() bool`

HasBoshInstanceName returns a boolean if a field has been set.

### GetBitness

`func (o *Host) GetBitness() string`

GetBitness returns the Bitness field if non-nil, zero value otherwise.

### GetBitnessOk

`func (o *Host) GetBitnessOk() (*string, bool)`

GetBitnessOk returns a tuple with the Bitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitness

`func (o *Host) SetBitness(v string)`

SetBitness sets Bitness field to given value.

### HasBitness

`func (o *Host) HasBitness() bool`

HasBitness returns a boolean if a field has been set.

### GetKubernetesLabels

`func (o *Host) GetKubernetesLabels() map[string]map[string]interface{}`

GetKubernetesLabels returns the KubernetesLabels field if non-nil, zero value otherwise.

### GetKubernetesLabelsOk

`func (o *Host) GetKubernetesLabelsOk() (*map[string]map[string]interface{}, bool)`

GetKubernetesLabelsOk returns a tuple with the KubernetesLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesLabels

`func (o *Host) SetKubernetesLabels(v map[string]map[string]interface{})`

SetKubernetesLabels sets KubernetesLabels field to given value.

### HasKubernetesLabels

`func (o *Host) HasKubernetesLabels() bool`

HasKubernetesLabels returns a boolean if a field has been set.

### GetAwsNameTag

`func (o *Host) GetAwsNameTag() string`

GetAwsNameTag returns the AwsNameTag field if non-nil, zero value otherwise.

### GetAwsNameTagOk

`func (o *Host) GetAwsNameTagOk() (*string, bool)`

GetAwsNameTagOk returns a tuple with the AwsNameTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsNameTag

`func (o *Host) SetAwsNameTag(v string)`

SetAwsNameTag sets AwsNameTag field to given value.

### HasAwsNameTag

`func (o *Host) HasAwsNameTag() bool`

HasAwsNameTag returns a boolean if a field has been set.

### GetGceInstanceId

`func (o *Host) GetGceInstanceId() string`

GetGceInstanceId returns the GceInstanceId field if non-nil, zero value otherwise.

### GetGceInstanceIdOk

`func (o *Host) GetGceInstanceIdOk() (*string, bool)`

GetGceInstanceIdOk returns a tuple with the GceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGceInstanceId

`func (o *Host) SetGceInstanceId(v string)`

SetGceInstanceId sets GceInstanceId field to given value.

### HasGceInstanceId

`func (o *Host) HasGceInstanceId() bool`

HasGceInstanceId returns a boolean if a field has been set.

### GetBeanstalkEnvironmentName

`func (o *Host) GetBeanstalkEnvironmentName() string`

GetBeanstalkEnvironmentName returns the BeanstalkEnvironmentName field if non-nil, zero value otherwise.

### GetBeanstalkEnvironmentNameOk

`func (o *Host) GetBeanstalkEnvironmentNameOk() (*string, bool)`

GetBeanstalkEnvironmentNameOk returns a tuple with the BeanstalkEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeanstalkEnvironmentName

`func (o *Host) SetBeanstalkEnvironmentName(v string)`

SetBeanstalkEnvironmentName sets BeanstalkEnvironmentName field to given value.

### HasBeanstalkEnvironmentName

`func (o *Host) HasBeanstalkEnvironmentName() bool`

HasBeanstalkEnvironmentName returns a boolean if a field has been set.

### GetGcePublicIpAddresses

`func (o *Host) GetGcePublicIpAddresses() []string`

GetGcePublicIpAddresses returns the GcePublicIpAddresses field if non-nil, zero value otherwise.

### GetGcePublicIpAddressesOk

`func (o *Host) GetGcePublicIpAddressesOk() (*[]string, bool)`

GetGcePublicIpAddressesOk returns a tuple with the GcePublicIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcePublicIpAddresses

`func (o *Host) SetGcePublicIpAddresses(v []string)`

SetGcePublicIpAddresses sets GcePublicIpAddresses field to given value.

### HasGcePublicIpAddresses

`func (o *Host) HasGcePublicIpAddresses() bool`

HasGcePublicIpAddresses returns a boolean if a field has been set.

### GetScaleSetName

`func (o *Host) GetScaleSetName() string`

GetScaleSetName returns the ScaleSetName field if non-nil, zero value otherwise.

### GetScaleSetNameOk

`func (o *Host) GetScaleSetNameOk() (*string, bool)`

GetScaleSetNameOk returns a tuple with the ScaleSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleSetName

`func (o *Host) SetScaleSetName(v string)`

SetScaleSetName sets ScaleSetName field to given value.

### HasScaleSetName

`func (o *Host) HasScaleSetName() bool`

HasScaleSetName returns a boolean if a field has been set.

### GetOneAgentCustomHostName

`func (o *Host) GetOneAgentCustomHostName() string`

GetOneAgentCustomHostName returns the OneAgentCustomHostName field if non-nil, zero value otherwise.

### GetOneAgentCustomHostNameOk

`func (o *Host) GetOneAgentCustomHostNameOk() (*string, bool)`

GetOneAgentCustomHostNameOk returns a tuple with the OneAgentCustomHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneAgentCustomHostName

`func (o *Host) SetOneAgentCustomHostName(v string)`

SetOneAgentCustomHostName sets OneAgentCustomHostName field to given value.

### HasOneAgentCustomHostName

`func (o *Host) HasOneAgentCustomHostName() bool`

HasOneAgentCustomHostName returns a boolean if a field has been set.

### GetAzureHostNames

`func (o *Host) GetAzureHostNames() []string`

GetAzureHostNames returns the AzureHostNames field if non-nil, zero value otherwise.

### GetAzureHostNamesOk

`func (o *Host) GetAzureHostNamesOk() (*[]string, bool)`

GetAzureHostNamesOk returns a tuple with the AzureHostNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureHostNames

`func (o *Host) SetAzureHostNames(v []string)`

SetAzureHostNames sets AzureHostNames field to given value.

### HasAzureHostNames

`func (o *Host) HasAzureHostNames() bool`

HasAzureHostNames returns a boolean if a field has been set.

### GetGceInstanceName

`func (o *Host) GetGceInstanceName() string`

GetGceInstanceName returns the GceInstanceName field if non-nil, zero value otherwise.

### GetGceInstanceNameOk

`func (o *Host) GetGceInstanceNameOk() (*string, bool)`

GetGceInstanceNameOk returns a tuple with the GceInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGceInstanceName

`func (o *Host) SetGceInstanceName(v string)`

SetGceInstanceName sets GceInstanceName field to given value.

### HasGceInstanceName

`func (o *Host) HasGceInstanceName() bool`

HasGceInstanceName returns a boolean if a field has been set.

### GetPaasType

`func (o *Host) GetPaasType() string`

GetPaasType returns the PaasType field if non-nil, zero value otherwise.

### GetPaasTypeOk

`func (o *Host) GetPaasTypeOk() (*string, bool)`

GetPaasTypeOk returns a tuple with the PaasType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaasType

`func (o *Host) SetPaasType(v string)`

SetPaasType sets PaasType field to given value.

### HasPaasType

`func (o *Host) HasPaasType() bool`

HasPaasType returns a boolean if a field has been set.

### GetBoshInstanceId

`func (o *Host) GetBoshInstanceId() string`

GetBoshInstanceId returns the BoshInstanceId field if non-nil, zero value otherwise.

### GetBoshInstanceIdOk

`func (o *Host) GetBoshInstanceIdOk() (*string, bool)`

GetBoshInstanceIdOk returns a tuple with the BoshInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoshInstanceId

`func (o *Host) SetBoshInstanceId(v string)`

SetBoshInstanceId sets BoshInstanceId field to given value.

### HasBoshInstanceId

`func (o *Host) HasBoshInstanceId() bool`

HasBoshInstanceId returns a boolean if a field has been set.

### GetAzureSku

`func (o *Host) GetAzureSku() string`

GetAzureSku returns the AzureSku field if non-nil, zero value otherwise.

### GetAzureSkuOk

`func (o *Host) GetAzureSkuOk() (*string, bool)`

GetAzureSkuOk returns a tuple with the AzureSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSku

`func (o *Host) SetAzureSku(v string)`

SetAzureSku sets AzureSku field to given value.

### HasAzureSku

`func (o *Host) HasAzureSku() bool`

HasAzureSku returns a boolean if a field has been set.

### GetCpuCores

`func (o *Host) GetCpuCores() int32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *Host) GetCpuCoresOk() (*int32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *Host) SetCpuCores(v int32)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *Host) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetPublicIp

`func (o *Host) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *Host) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *Host) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *Host) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

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

### GetBoshName

`func (o *Host) GetBoshName() string`

GetBoshName returns the BoshName field if non-nil, zero value otherwise.

### GetBoshNameOk

`func (o *Host) GetBoshNameOk() (*string, bool)`

GetBoshNameOk returns a tuple with the BoshName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoshName

`func (o *Host) SetBoshName(v string)`

SetBoshName sets BoshName field to given value.

### HasBoshName

`func (o *Host) HasBoshName() bool`

HasBoshName returns a boolean if a field has been set.

### GetZosTotalPhysicalMemory

`func (o *Host) GetZosTotalPhysicalMemory() int32`

GetZosTotalPhysicalMemory returns the ZosTotalPhysicalMemory field if non-nil, zero value otherwise.

### GetZosTotalPhysicalMemoryOk

`func (o *Host) GetZosTotalPhysicalMemoryOk() (*int32, bool)`

GetZosTotalPhysicalMemoryOk returns a tuple with the ZosTotalPhysicalMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZosTotalPhysicalMemory

`func (o *Host) SetZosTotalPhysicalMemory(v int32)`

SetZosTotalPhysicalMemory sets ZosTotalPhysicalMemory field to given value.

### HasZosTotalPhysicalMemory

`func (o *Host) HasZosTotalPhysicalMemory() bool`

HasZosTotalPhysicalMemory returns a boolean if a field has been set.

### GetOpenstackVmName

`func (o *Host) GetOpenstackVmName() string`

GetOpenstackVmName returns the OpenstackVmName field if non-nil, zero value otherwise.

### GetOpenstackVmNameOk

`func (o *Host) GetOpenstackVmNameOk() (*string, bool)`

GetOpenstackVmNameOk returns a tuple with the OpenstackVmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackVmName

`func (o *Host) SetOpenstackVmName(v string)`

SetOpenstackVmName sets OpenstackVmName field to given value.

### HasOpenstackVmName

`func (o *Host) HasOpenstackVmName() bool`

HasOpenstackVmName returns a boolean if a field has been set.

### GetZosVirtualization

`func (o *Host) GetZosVirtualization() string`

GetZosVirtualization returns the ZosVirtualization field if non-nil, zero value otherwise.

### GetZosVirtualizationOk

`func (o *Host) GetZosVirtualizationOk() (*string, bool)`

GetZosVirtualizationOk returns a tuple with the ZosVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZosVirtualization

`func (o *Host) SetZosVirtualization(v string)`

SetZosVirtualization sets ZosVirtualization field to given value.

### HasZosVirtualization

`func (o *Host) HasZosVirtualization() bool`

HasZosVirtualization returns a boolean if a field has been set.

### GetVirtualCpus

`func (o *Host) GetVirtualCpus() int32`

GetVirtualCpus returns the VirtualCpus field if non-nil, zero value otherwise.

### GetVirtualCpusOk

`func (o *Host) GetVirtualCpusOk() (*int32, bool)`

GetVirtualCpusOk returns a tuple with the VirtualCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCpus

`func (o *Host) SetVirtualCpus(v int32)`

SetVirtualCpus sets VirtualCpus field to given value.

### HasVirtualCpus

`func (o *Host) HasVirtualCpus() bool`

HasVirtualCpus returns a boolean if a field has been set.

### GetBoshAvailabilityZone

`func (o *Host) GetBoshAvailabilityZone() string`

GetBoshAvailabilityZone returns the BoshAvailabilityZone field if non-nil, zero value otherwise.

### GetBoshAvailabilityZoneOk

`func (o *Host) GetBoshAvailabilityZoneOk() (*string, bool)`

GetBoshAvailabilityZoneOk returns a tuple with the BoshAvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoshAvailabilityZone

`func (o *Host) SetBoshAvailabilityZone(v string)`

SetBoshAvailabilityZone sets BoshAvailabilityZone field to given value.

### HasBoshAvailabilityZone

`func (o *Host) HasBoshAvailabilityZone() bool`

HasBoshAvailabilityZone returns a boolean if a field has been set.

### GetCloudType

`func (o *Host) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *Host) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *Host) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *Host) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetPublicHostName

`func (o *Host) GetPublicHostName() string`

GetPublicHostName returns the PublicHostName field if non-nil, zero value otherwise.

### GetPublicHostNameOk

`func (o *Host) GetPublicHostNameOk() (*string, bool)`

GetPublicHostNameOk returns a tuple with the PublicHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicHostName

`func (o *Host) SetPublicHostName(v string)`

SetPublicHostName sets PublicHostName field to given value.

### HasPublicHostName

`func (o *Host) HasPublicHostName() bool`

HasPublicHostName returns a boolean if a field has been set.

### GetConsumedHostUnits

`func (o *Host) GetConsumedHostUnits() string`

GetConsumedHostUnits returns the ConsumedHostUnits field if non-nil, zero value otherwise.

### GetConsumedHostUnitsOk

`func (o *Host) GetConsumedHostUnitsOk() (*string, bool)`

GetConsumedHostUnitsOk returns a tuple with the ConsumedHostUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedHostUnits

`func (o *Host) SetConsumedHostUnits(v string)`

SetConsumedHostUnits sets ConsumedHostUnits field to given value.

### HasConsumedHostUnits

`func (o *Host) HasConsumedHostUnits() bool`

HasConsumedHostUnits returns a boolean if a field has been set.

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

### GetSimultaneousMultithreading

`func (o *Host) GetSimultaneousMultithreading() int32`

GetSimultaneousMultithreading returns the SimultaneousMultithreading field if non-nil, zero value otherwise.

### GetSimultaneousMultithreadingOk

`func (o *Host) GetSimultaneousMultithreadingOk() (*int32, bool)`

GetSimultaneousMultithreadingOk returns a tuple with the SimultaneousMultithreading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimultaneousMultithreading

`func (o *Host) SetSimultaneousMultithreading(v int32)`

SetSimultaneousMultithreading sets SimultaneousMultithreading field to given value.

### HasSimultaneousMultithreading

`func (o *Host) HasSimultaneousMultithreading() bool`

HasSimultaneousMultithreading returns a boolean if a field has been set.

### GetBoshStemcellVersion

`func (o *Host) GetBoshStemcellVersion() string`

GetBoshStemcellVersion returns the BoshStemcellVersion field if non-nil, zero value otherwise.

### GetBoshStemcellVersionOk

`func (o *Host) GetBoshStemcellVersionOk() (*string, bool)`

GetBoshStemcellVersionOk returns a tuple with the BoshStemcellVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoshStemcellVersion

`func (o *Host) SetBoshStemcellVersion(v string)`

SetBoshStemcellVersion sets BoshStemcellVersion field to given value.

### HasBoshStemcellVersion

`func (o *Host) HasBoshStemcellVersion() bool`

HasBoshStemcellVersion returns a boolean if a field has been set.

### GetZosTotalGeneralPurposeProcessors

`func (o *Host) GetZosTotalGeneralPurposeProcessors() int32`

GetZosTotalGeneralPurposeProcessors returns the ZosTotalGeneralPurposeProcessors field if non-nil, zero value otherwise.

### GetZosTotalGeneralPurposeProcessorsOk

`func (o *Host) GetZosTotalGeneralPurposeProcessorsOk() (*int32, bool)`

GetZosTotalGeneralPurposeProcessorsOk returns a tuple with the ZosTotalGeneralPurposeProcessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZosTotalGeneralPurposeProcessors

`func (o *Host) SetZosTotalGeneralPurposeProcessors(v int32)`

SetZosTotalGeneralPurposeProcessors sets ZosTotalGeneralPurposeProcessors field to given value.

### HasZosTotalGeneralPurposeProcessors

`func (o *Host) HasZosTotalGeneralPurposeProcessors() bool`

HasZosTotalGeneralPurposeProcessors returns a boolean if a field has been set.

### GetOpenstackComputeNodeName

`func (o *Host) GetOpenstackComputeNodeName() string`

GetOpenstackComputeNodeName returns the OpenstackComputeNodeName field if non-nil, zero value otherwise.

### GetOpenstackComputeNodeNameOk

`func (o *Host) GetOpenstackComputeNodeNameOk() (*string, bool)`

GetOpenstackComputeNodeNameOk returns a tuple with the OpenstackComputeNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackComputeNodeName

`func (o *Host) SetOpenstackComputeNodeName(v string)`

SetOpenstackComputeNodeName sets OpenstackComputeNodeName field to given value.

### HasOpenstackComputeNodeName

`func (o *Host) HasOpenstackComputeNodeName() bool`

HasOpenstackComputeNodeName returns a boolean if a field has been set.

### GetKubernetesCluster

`func (o *Host) GetKubernetesCluster() string`

GetKubernetesCluster returns the KubernetesCluster field if non-nil, zero value otherwise.

### GetKubernetesClusterOk

`func (o *Host) GetKubernetesClusterOk() (*string, bool)`

GetKubernetesClusterOk returns a tuple with the KubernetesCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCluster

`func (o *Host) SetKubernetesCluster(v string)`

SetKubernetesCluster sets KubernetesCluster field to given value.

### HasKubernetesCluster

`func (o *Host) HasKubernetesCluster() bool`

HasKubernetesCluster returns a boolean if a field has been set.

### GetSoftwareTechnologies

`func (o *Host) GetSoftwareTechnologies() []TechnologyInfo`

GetSoftwareTechnologies returns the SoftwareTechnologies field if non-nil, zero value otherwise.

### GetSoftwareTechnologiesOk

`func (o *Host) GetSoftwareTechnologiesOk() (*[]TechnologyInfo, bool)`

GetSoftwareTechnologiesOk returns a tuple with the SoftwareTechnologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTechnologies

`func (o *Host) SetSoftwareTechnologies(v []TechnologyInfo)`

SetSoftwareTechnologies sets SoftwareTechnologies field to given value.

### HasSoftwareTechnologies

`func (o *Host) HasSoftwareTechnologies() bool`

HasSoftwareTechnologies returns a boolean if a field has been set.

### GetGceMachineType

`func (o *Host) GetGceMachineType() string`

GetGceMachineType returns the GceMachineType field if non-nil, zero value otherwise.

### GetGceMachineTypeOk

`func (o *Host) GetGceMachineTypeOk() (*string, bool)`

GetGceMachineTypeOk returns a tuple with the GceMachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGceMachineType

`func (o *Host) SetGceMachineType(v string)`

SetGceMachineType sets GceMachineType field to given value.

### HasGceMachineType

`func (o *Host) HasGceMachineType() bool`

HasGceMachineType returns a boolean if a field has been set.

### GetZosCPUSerialNumber

`func (o *Host) GetZosCPUSerialNumber() string`

GetZosCPUSerialNumber returns the ZosCPUSerialNumber field if non-nil, zero value otherwise.

### GetZosCPUSerialNumberOk

`func (o *Host) GetZosCPUSerialNumberOk() (*string, bool)`

GetZosCPUSerialNumberOk returns a tuple with the ZosCPUSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZosCPUSerialNumber

`func (o *Host) SetZosCPUSerialNumber(v string)`

SetZosCPUSerialNumber sets ZosCPUSerialNumber field to given value.

### HasZosCPUSerialNumber

`func (o *Host) HasZosCPUSerialNumber() bool`

HasZosCPUSerialNumber returns a boolean if a field has been set.

### GetPaasAgentVersions

`func (o *Host) GetPaasAgentVersions() []AgentVersion`

GetPaasAgentVersions returns the PaasAgentVersions field if non-nil, zero value otherwise.

### GetPaasAgentVersionsOk

`func (o *Host) GetPaasAgentVersionsOk() (*[]AgentVersion, bool)`

GetPaasAgentVersionsOk returns a tuple with the PaasAgentVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaasAgentVersions

`func (o *Host) SetPaasAgentVersions(v []AgentVersion)`

SetPaasAgentVersions sets PaasAgentVersions field to given value.

### HasPaasAgentVersions

`func (o *Host) HasPaasAgentVersions() bool`

HasPaasAgentVersions returns a boolean if a field has been set.

### GetOpenStackInstaceType

`func (o *Host) GetOpenStackInstaceType() string`

GetOpenStackInstaceType returns the OpenStackInstaceType field if non-nil, zero value otherwise.

### GetOpenStackInstaceTypeOk

`func (o *Host) GetOpenStackInstaceTypeOk() (*string, bool)`

GetOpenStackInstaceTypeOk returns a tuple with the OpenStackInstaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackInstaceType

`func (o *Host) SetOpenStackInstaceType(v string)`

SetOpenStackInstaceType sets OpenStackInstaceType field to given value.

### HasOpenStackInstaceType

`func (o *Host) HasOpenStackInstaceType() bool`

HasOpenStackInstaceType returns a boolean if a field has been set.

### GetAwsInstanceType

`func (o *Host) GetAwsInstanceType() string`

GetAwsInstanceType returns the AwsInstanceType field if non-nil, zero value otherwise.

### GetAwsInstanceTypeOk

`func (o *Host) GetAwsInstanceTypeOk() (*string, bool)`

GetAwsInstanceTypeOk returns a tuple with the AwsInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsInstanceType

`func (o *Host) SetAwsInstanceType(v string)`

SetAwsInstanceType sets AwsInstanceType field to given value.

### HasAwsInstanceType

`func (o *Host) HasAwsInstanceType() bool`

HasAwsInstanceType returns a boolean if a field has been set.

### GetAzureSiteNames

`func (o *Host) GetAzureSiteNames() []string`

GetAzureSiteNames returns the AzureSiteNames field if non-nil, zero value otherwise.

### GetAzureSiteNamesOk

`func (o *Host) GetAzureSiteNamesOk() (*[]string, bool)`

GetAzureSiteNamesOk returns a tuple with the AzureSiteNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSiteNames

`func (o *Host) SetAzureSiteNames(v []string)`

SetAzureSiteNames sets AzureSiteNames field to given value.

### HasAzureSiteNames

`func (o *Host) HasAzureSiteNames() bool`

HasAzureSiteNames returns a boolean if a field has been set.

### GetOpenstackSecurityGroups

`func (o *Host) GetOpenstackSecurityGroups() []string`

GetOpenstackSecurityGroups returns the OpenstackSecurityGroups field if non-nil, zero value otherwise.

### GetOpenstackSecurityGroupsOk

`func (o *Host) GetOpenstackSecurityGroupsOk() (*[]string, bool)`

GetOpenstackSecurityGroupsOk returns a tuple with the OpenstackSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackSecurityGroups

`func (o *Host) SetOpenstackSecurityGroups(v []string)`

SetOpenstackSecurityGroups sets OpenstackSecurityGroups field to given value.

### HasOpenstackSecurityGroups

`func (o *Host) HasOpenstackSecurityGroups() bool`

HasOpenstackSecurityGroups returns a boolean if a field has been set.

### GetVmwareName

`func (o *Host) GetVmwareName() string`

GetVmwareName returns the VmwareName field if non-nil, zero value otherwise.

### GetVmwareNameOk

`func (o *Host) GetVmwareNameOk() (*string, bool)`

GetVmwareNameOk returns a tuple with the VmwareName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareName

`func (o *Host) SetVmwareName(v string)`

SetVmwareName sets VmwareName field to given value.

### HasVmwareName

`func (o *Host) HasVmwareName() bool`

HasVmwareName returns a boolean if a field has been set.

### GetCloudPlatformVendorVersion

`func (o *Host) GetCloudPlatformVendorVersion() string`

GetCloudPlatformVendorVersion returns the CloudPlatformVendorVersion field if non-nil, zero value otherwise.

### GetCloudPlatformVendorVersionOk

`func (o *Host) GetCloudPlatformVendorVersionOk() (*string, bool)`

GetCloudPlatformVendorVersionOk returns a tuple with the CloudPlatformVendorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPlatformVendorVersion

`func (o *Host) SetCloudPlatformVendorVersion(v string)`

SetCloudPlatformVendorVersion sets CloudPlatformVendorVersion field to given value.

### HasCloudPlatformVendorVersion

`func (o *Host) HasCloudPlatformVendorVersion() bool`

HasCloudPlatformVendorVersion returns a boolean if a field has been set.

### GetKubernetesNode

`func (o *Host) GetKubernetesNode() string`

GetKubernetesNode returns the KubernetesNode field if non-nil, zero value otherwise.

### GetKubernetesNodeOk

`func (o *Host) GetKubernetesNodeOk() (*string, bool)`

GetKubernetesNodeOk returns a tuple with the KubernetesNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesNode

`func (o *Host) SetKubernetesNode(v string)`

SetKubernetesNode sets KubernetesNode field to given value.

### HasKubernetesNode

`func (o *Host) HasKubernetesNode() bool`

HasKubernetesNode returns a boolean if a field has been set.

### GetGcpZone

`func (o *Host) GetGcpZone() string`

GetGcpZone returns the GcpZone field if non-nil, zero value otherwise.

### GetGcpZoneOk

`func (o *Host) GetGcpZoneOk() (*string, bool)`

GetGcpZoneOk returns a tuple with the GcpZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpZone

`func (o *Host) SetGcpZone(v string)`

SetGcpZone sets GcpZone field to given value.

### HasGcpZone

`func (o *Host) HasGcpZone() bool`

HasGcpZone returns a boolean if a field has been set.

### GetHypervisorType

`func (o *Host) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *Host) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *Host) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *Host) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetIpAddresses

`func (o *Host) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *Host) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *Host) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *Host) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetLocalIp

`func (o *Host) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *Host) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *Host) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *Host) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetFromRelationships

`func (o *Host) GetFromRelationships() HostFromRelationships`

GetFromRelationships returns the FromRelationships field if non-nil, zero value otherwise.

### GetFromRelationshipsOk

`func (o *Host) GetFromRelationshipsOk() (*HostFromRelationships, bool)`

GetFromRelationshipsOk returns a tuple with the FromRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromRelationships

`func (o *Host) SetFromRelationships(v HostFromRelationships)`

SetFromRelationships sets FromRelationships field to given value.

### HasFromRelationships

`func (o *Host) HasFromRelationships() bool`

HasFromRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


