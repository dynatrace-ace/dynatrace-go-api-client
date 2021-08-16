# Service

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
**ToRelationships** | Pointer to [**ServiceToRelationships**](ServiceToRelationships.md) |  | [optional] 
**AkkaActorSystem** | Pointer to **string** | The services of the akka actor system. | [optional] 
**WebServiceNamespace** | Pointer to **string** |  | [optional] 
**ClassName** | Pointer to **string** |  | [optional] 
**IibApplicationName** | Pointer to **string** | The IIB application name. | [optional] 
**RemoteEndpoint** | Pointer to **string** | The endpoint of the remote service. | [optional] 
**IbmCtgServerName** | Pointer to **string** | The IBM CICS Transaction Gateway name. | [optional] 
**SoftwareTechnologies** | Pointer to [**[]TechnologyInfo**](TechnologyInfo.md) |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**DatabaseName** | Pointer to **string** |  | [optional] 
**WebServerName** | Pointer to **string** |  | [optional] 
**ManagementZones** | Pointer to [**[]EntityShortRepresentation**](EntityShortRepresentation.md) | The management zones that the entity is part of. | [optional] 
**IbmCtgGatewayUrl** | Pointer to **string** | The IBM CTG gateway URL. | [optional] 
**ServiceTechnologyTypes** | Pointer to **[]string** |  | [optional] 
**EsbApplicationName** | Pointer to **string** | The ESB application name. | [optional] 
**IpAddresses** | Pointer to **[]string** |  | [optional] 
**RemoteServiceName** | Pointer to **string** | The name of the remote service. | [optional] 
**AgentTechnologyType** | Pointer to **string** |  | [optional] 
**DatabaseVendor** | Pointer to **string** |  | [optional] 
**ContextRoot** | Pointer to **string** |  | [optional] 
**WebApplicationId** | Pointer to **string** |  | [optional] 
**DatabaseHostNames** | Pointer to **[]string** |  | [optional] 
**WebServiceName** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**FromRelationships** | Pointer to [**ServiceFromRelationships**](ServiceFromRelationships.md) |  | [optional] 

## Methods

### NewService

`func NewService() *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *Service) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Service) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Service) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *Service) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetDisplayName

`func (o *Service) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Service) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Service) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Service) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCustomizedName

`func (o *Service) GetCustomizedName() string`

GetCustomizedName returns the CustomizedName field if non-nil, zero value otherwise.

### GetCustomizedNameOk

`func (o *Service) GetCustomizedNameOk() (*string, bool)`

GetCustomizedNameOk returns a tuple with the CustomizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizedName

`func (o *Service) SetCustomizedName(v string)`

SetCustomizedName sets CustomizedName field to given value.

### HasCustomizedName

`func (o *Service) HasCustomizedName() bool`

HasCustomizedName returns a boolean if a field has been set.

### GetDiscoveredName

`func (o *Service) GetDiscoveredName() string`

GetDiscoveredName returns the DiscoveredName field if non-nil, zero value otherwise.

### GetDiscoveredNameOk

`func (o *Service) GetDiscoveredNameOk() (*string, bool)`

GetDiscoveredNameOk returns a tuple with the DiscoveredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredName

`func (o *Service) SetDiscoveredName(v string)`

SetDiscoveredName sets DiscoveredName field to given value.

### HasDiscoveredName

`func (o *Service) HasDiscoveredName() bool`

HasDiscoveredName returns a boolean if a field has been set.

### GetFirstSeenTimestamp

`func (o *Service) GetFirstSeenTimestamp() int64`

GetFirstSeenTimestamp returns the FirstSeenTimestamp field if non-nil, zero value otherwise.

### GetFirstSeenTimestampOk

`func (o *Service) GetFirstSeenTimestampOk() (*int64, bool)`

GetFirstSeenTimestampOk returns a tuple with the FirstSeenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenTimestamp

`func (o *Service) SetFirstSeenTimestamp(v int64)`

SetFirstSeenTimestamp sets FirstSeenTimestamp field to given value.

### HasFirstSeenTimestamp

`func (o *Service) HasFirstSeenTimestamp() bool`

HasFirstSeenTimestamp returns a boolean if a field has been set.

### GetLastSeenTimestamp

`func (o *Service) GetLastSeenTimestamp() int64`

GetLastSeenTimestamp returns the LastSeenTimestamp field if non-nil, zero value otherwise.

### GetLastSeenTimestampOk

`func (o *Service) GetLastSeenTimestampOk() (*int64, bool)`

GetLastSeenTimestampOk returns a tuple with the LastSeenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenTimestamp

`func (o *Service) SetLastSeenTimestamp(v int64)`

SetLastSeenTimestamp sets LastSeenTimestamp field to given value.

### HasLastSeenTimestamp

`func (o *Service) HasLastSeenTimestamp() bool`

HasLastSeenTimestamp returns a boolean if a field has been set.

### GetTags

`func (o *Service) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Service) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Service) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Service) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToRelationships

`func (o *Service) GetToRelationships() ServiceToRelationships`

GetToRelationships returns the ToRelationships field if non-nil, zero value otherwise.

### GetToRelationshipsOk

`func (o *Service) GetToRelationshipsOk() (*ServiceToRelationships, bool)`

GetToRelationshipsOk returns a tuple with the ToRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRelationships

`func (o *Service) SetToRelationships(v ServiceToRelationships)`

SetToRelationships sets ToRelationships field to given value.

### HasToRelationships

`func (o *Service) HasToRelationships() bool`

HasToRelationships returns a boolean if a field has been set.

### GetAkkaActorSystem

`func (o *Service) GetAkkaActorSystem() string`

GetAkkaActorSystem returns the AkkaActorSystem field if non-nil, zero value otherwise.

### GetAkkaActorSystemOk

`func (o *Service) GetAkkaActorSystemOk() (*string, bool)`

GetAkkaActorSystemOk returns a tuple with the AkkaActorSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkkaActorSystem

`func (o *Service) SetAkkaActorSystem(v string)`

SetAkkaActorSystem sets AkkaActorSystem field to given value.

### HasAkkaActorSystem

`func (o *Service) HasAkkaActorSystem() bool`

HasAkkaActorSystem returns a boolean if a field has been set.

### GetWebServiceNamespace

`func (o *Service) GetWebServiceNamespace() string`

GetWebServiceNamespace returns the WebServiceNamespace field if non-nil, zero value otherwise.

### GetWebServiceNamespaceOk

`func (o *Service) GetWebServiceNamespaceOk() (*string, bool)`

GetWebServiceNamespaceOk returns a tuple with the WebServiceNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceNamespace

`func (o *Service) SetWebServiceNamespace(v string)`

SetWebServiceNamespace sets WebServiceNamespace field to given value.

### HasWebServiceNamespace

`func (o *Service) HasWebServiceNamespace() bool`

HasWebServiceNamespace returns a boolean if a field has been set.

### GetClassName

`func (o *Service) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *Service) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *Service) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *Service) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetIibApplicationName

`func (o *Service) GetIibApplicationName() string`

GetIibApplicationName returns the IibApplicationName field if non-nil, zero value otherwise.

### GetIibApplicationNameOk

`func (o *Service) GetIibApplicationNameOk() (*string, bool)`

GetIibApplicationNameOk returns a tuple with the IibApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIibApplicationName

`func (o *Service) SetIibApplicationName(v string)`

SetIibApplicationName sets IibApplicationName field to given value.

### HasIibApplicationName

`func (o *Service) HasIibApplicationName() bool`

HasIibApplicationName returns a boolean if a field has been set.

### GetRemoteEndpoint

`func (o *Service) GetRemoteEndpoint() string`

GetRemoteEndpoint returns the RemoteEndpoint field if non-nil, zero value otherwise.

### GetRemoteEndpointOk

`func (o *Service) GetRemoteEndpointOk() (*string, bool)`

GetRemoteEndpointOk returns a tuple with the RemoteEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteEndpoint

`func (o *Service) SetRemoteEndpoint(v string)`

SetRemoteEndpoint sets RemoteEndpoint field to given value.

### HasRemoteEndpoint

`func (o *Service) HasRemoteEndpoint() bool`

HasRemoteEndpoint returns a boolean if a field has been set.

### GetIbmCtgServerName

`func (o *Service) GetIbmCtgServerName() string`

GetIbmCtgServerName returns the IbmCtgServerName field if non-nil, zero value otherwise.

### GetIbmCtgServerNameOk

`func (o *Service) GetIbmCtgServerNameOk() (*string, bool)`

GetIbmCtgServerNameOk returns a tuple with the IbmCtgServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbmCtgServerName

`func (o *Service) SetIbmCtgServerName(v string)`

SetIbmCtgServerName sets IbmCtgServerName field to given value.

### HasIbmCtgServerName

`func (o *Service) HasIbmCtgServerName() bool`

HasIbmCtgServerName returns a boolean if a field has been set.

### GetSoftwareTechnologies

`func (o *Service) GetSoftwareTechnologies() []TechnologyInfo`

GetSoftwareTechnologies returns the SoftwareTechnologies field if non-nil, zero value otherwise.

### GetSoftwareTechnologiesOk

`func (o *Service) GetSoftwareTechnologiesOk() (*[]TechnologyInfo, bool)`

GetSoftwareTechnologiesOk returns a tuple with the SoftwareTechnologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTechnologies

`func (o *Service) SetSoftwareTechnologies(v []TechnologyInfo)`

SetSoftwareTechnologies sets SoftwareTechnologies field to given value.

### HasSoftwareTechnologies

`func (o *Service) HasSoftwareTechnologies() bool`

HasSoftwareTechnologies returns a boolean if a field has been set.

### GetPort

`func (o *Service) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Service) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Service) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Service) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetDatabaseName

`func (o *Service) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *Service) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *Service) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *Service) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetWebServerName

`func (o *Service) GetWebServerName() string`

GetWebServerName returns the WebServerName field if non-nil, zero value otherwise.

### GetWebServerNameOk

`func (o *Service) GetWebServerNameOk() (*string, bool)`

GetWebServerNameOk returns a tuple with the WebServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServerName

`func (o *Service) SetWebServerName(v string)`

SetWebServerName sets WebServerName field to given value.

### HasWebServerName

`func (o *Service) HasWebServerName() bool`

HasWebServerName returns a boolean if a field has been set.

### GetManagementZones

`func (o *Service) GetManagementZones() []EntityShortRepresentation`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *Service) GetManagementZonesOk() (*[]EntityShortRepresentation, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *Service) SetManagementZones(v []EntityShortRepresentation)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *Service) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetIbmCtgGatewayUrl

`func (o *Service) GetIbmCtgGatewayUrl() string`

GetIbmCtgGatewayUrl returns the IbmCtgGatewayUrl field if non-nil, zero value otherwise.

### GetIbmCtgGatewayUrlOk

`func (o *Service) GetIbmCtgGatewayUrlOk() (*string, bool)`

GetIbmCtgGatewayUrlOk returns a tuple with the IbmCtgGatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbmCtgGatewayUrl

`func (o *Service) SetIbmCtgGatewayUrl(v string)`

SetIbmCtgGatewayUrl sets IbmCtgGatewayUrl field to given value.

### HasIbmCtgGatewayUrl

`func (o *Service) HasIbmCtgGatewayUrl() bool`

HasIbmCtgGatewayUrl returns a boolean if a field has been set.

### GetServiceTechnologyTypes

`func (o *Service) GetServiceTechnologyTypes() []string`

GetServiceTechnologyTypes returns the ServiceTechnologyTypes field if non-nil, zero value otherwise.

### GetServiceTechnologyTypesOk

`func (o *Service) GetServiceTechnologyTypesOk() (*[]string, bool)`

GetServiceTechnologyTypesOk returns a tuple with the ServiceTechnologyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTechnologyTypes

`func (o *Service) SetServiceTechnologyTypes(v []string)`

SetServiceTechnologyTypes sets ServiceTechnologyTypes field to given value.

### HasServiceTechnologyTypes

`func (o *Service) HasServiceTechnologyTypes() bool`

HasServiceTechnologyTypes returns a boolean if a field has been set.

### GetEsbApplicationName

`func (o *Service) GetEsbApplicationName() string`

GetEsbApplicationName returns the EsbApplicationName field if non-nil, zero value otherwise.

### GetEsbApplicationNameOk

`func (o *Service) GetEsbApplicationNameOk() (*string, bool)`

GetEsbApplicationNameOk returns a tuple with the EsbApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsbApplicationName

`func (o *Service) SetEsbApplicationName(v string)`

SetEsbApplicationName sets EsbApplicationName field to given value.

### HasEsbApplicationName

`func (o *Service) HasEsbApplicationName() bool`

HasEsbApplicationName returns a boolean if a field has been set.

### GetIpAddresses

`func (o *Service) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *Service) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *Service) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *Service) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetRemoteServiceName

`func (o *Service) GetRemoteServiceName() string`

GetRemoteServiceName returns the RemoteServiceName field if non-nil, zero value otherwise.

### GetRemoteServiceNameOk

`func (o *Service) GetRemoteServiceNameOk() (*string, bool)`

GetRemoteServiceNameOk returns a tuple with the RemoteServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteServiceName

`func (o *Service) SetRemoteServiceName(v string)`

SetRemoteServiceName sets RemoteServiceName field to given value.

### HasRemoteServiceName

`func (o *Service) HasRemoteServiceName() bool`

HasRemoteServiceName returns a boolean if a field has been set.

### GetAgentTechnologyType

`func (o *Service) GetAgentTechnologyType() string`

GetAgentTechnologyType returns the AgentTechnologyType field if non-nil, zero value otherwise.

### GetAgentTechnologyTypeOk

`func (o *Service) GetAgentTechnologyTypeOk() (*string, bool)`

GetAgentTechnologyTypeOk returns a tuple with the AgentTechnologyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTechnologyType

`func (o *Service) SetAgentTechnologyType(v string)`

SetAgentTechnologyType sets AgentTechnologyType field to given value.

### HasAgentTechnologyType

`func (o *Service) HasAgentTechnologyType() bool`

HasAgentTechnologyType returns a boolean if a field has been set.

### GetDatabaseVendor

`func (o *Service) GetDatabaseVendor() string`

GetDatabaseVendor returns the DatabaseVendor field if non-nil, zero value otherwise.

### GetDatabaseVendorOk

`func (o *Service) GetDatabaseVendorOk() (*string, bool)`

GetDatabaseVendorOk returns a tuple with the DatabaseVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseVendor

`func (o *Service) SetDatabaseVendor(v string)`

SetDatabaseVendor sets DatabaseVendor field to given value.

### HasDatabaseVendor

`func (o *Service) HasDatabaseVendor() bool`

HasDatabaseVendor returns a boolean if a field has been set.

### GetContextRoot

`func (o *Service) GetContextRoot() string`

GetContextRoot returns the ContextRoot field if non-nil, zero value otherwise.

### GetContextRootOk

`func (o *Service) GetContextRootOk() (*string, bool)`

GetContextRootOk returns a tuple with the ContextRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextRoot

`func (o *Service) SetContextRoot(v string)`

SetContextRoot sets ContextRoot field to given value.

### HasContextRoot

`func (o *Service) HasContextRoot() bool`

HasContextRoot returns a boolean if a field has been set.

### GetWebApplicationId

`func (o *Service) GetWebApplicationId() string`

GetWebApplicationId returns the WebApplicationId field if non-nil, zero value otherwise.

### GetWebApplicationIdOk

`func (o *Service) GetWebApplicationIdOk() (*string, bool)`

GetWebApplicationIdOk returns a tuple with the WebApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebApplicationId

`func (o *Service) SetWebApplicationId(v string)`

SetWebApplicationId sets WebApplicationId field to given value.

### HasWebApplicationId

`func (o *Service) HasWebApplicationId() bool`

HasWebApplicationId returns a boolean if a field has been set.

### GetDatabaseHostNames

`func (o *Service) GetDatabaseHostNames() []string`

GetDatabaseHostNames returns the DatabaseHostNames field if non-nil, zero value otherwise.

### GetDatabaseHostNamesOk

`func (o *Service) GetDatabaseHostNamesOk() (*[]string, bool)`

GetDatabaseHostNamesOk returns a tuple with the DatabaseHostNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseHostNames

`func (o *Service) SetDatabaseHostNames(v []string)`

SetDatabaseHostNames sets DatabaseHostNames field to given value.

### HasDatabaseHostNames

`func (o *Service) HasDatabaseHostNames() bool`

HasDatabaseHostNames returns a boolean if a field has been set.

### GetWebServiceName

`func (o *Service) GetWebServiceName() string`

GetWebServiceName returns the WebServiceName field if non-nil, zero value otherwise.

### GetWebServiceNameOk

`func (o *Service) GetWebServiceNameOk() (*string, bool)`

GetWebServiceNameOk returns a tuple with the WebServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceName

`func (o *Service) SetWebServiceName(v string)`

SetWebServiceName sets WebServiceName field to given value.

### HasWebServiceName

`func (o *Service) HasWebServiceName() bool`

HasWebServiceName returns a boolean if a field has been set.

### GetServiceType

`func (o *Service) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *Service) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *Service) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *Service) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetPath

`func (o *Service) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Service) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Service) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Service) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetFromRelationships

`func (o *Service) GetFromRelationships() ServiceFromRelationships`

GetFromRelationships returns the FromRelationships field if non-nil, zero value otherwise.

### GetFromRelationshipsOk

`func (o *Service) GetFromRelationshipsOk() (*ServiceFromRelationships, bool)`

GetFromRelationshipsOk returns a tuple with the FromRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromRelationships

`func (o *Service) SetFromRelationships(v ServiceFromRelationships)`

SetFromRelationships sets FromRelationships field to given value.

### HasFromRelationships

`func (o *Service) HasFromRelationships() bool`

HasFromRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


