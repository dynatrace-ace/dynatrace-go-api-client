# SecurityProblemDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityProblemId** | Pointer to **string** | The ID of the security problem. | [optional] [readonly] 
**DisplayId** | Pointer to **string** | The displayId of the security problem. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the security problem. | [optional] [readonly] 
**Muted** | Pointer to **bool** | Indicates if a security problem is muted. | [optional] [readonly] 
**ExternalVulnerabilityId** | Pointer to **string** | The external vulnerability ID of the security problem. | [optional] [readonly] 
**VulnerabilityType** | Pointer to **string** | The type of the vulnerability. | [optional] [readonly] 
**Title** | Pointer to **string** | The title of the security problem. | [optional] [readonly] 
**PackageName** | Pointer to **string** | The package name of the security problem. | [optional] [readonly] 
**Url** | Pointer to **string** | The URL to the security problem details page. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the security problem. | [optional] [readonly] 
**Technology** | Pointer to **string** | The technology of the security problem. | [optional] [readonly] 
**FirstSeenTimestamp** | Pointer to **int64** | The timestamp of the first occurrence of the security problem. | [optional] [readonly] 
**LastUpdatedTimestamp** | Pointer to **int64** | The timestamp of the most recent security problem change. | [optional] [readonly] 
**RiskAssessment** | Pointer to [**RiskAssessment**](RiskAssessment.md) |  | [optional] 
**ManagementZones** | Pointer to [**[]ManagementZone**](ManagementZone.md) | Management zones to which the affected entities belong. | [optional] [readonly] 
**CveIds** | Pointer to **[]string** | CVE IDs of the security problem. | [optional] [readonly] 
**Events** | Pointer to [**[]SecurityProblemEvent**](SecurityProblemEvent.md) | An ordered (newest first) list of events of the security problem. | [optional] 
**VulnerableComponents** | Pointer to [**[]VulnerableComponent**](VulnerableComponent.md) | A list of vulnerable components of the security problem.   A vulnerable component is what causes the security problem. | [optional] [readonly] 
**AffectedEntities** | Pointer to **[]string** | A list of affected entities of the security problem.   An affected entity is an entity where a vulnerable component runs. | [optional] [readonly] 
**ExposedEntities** | Pointer to **[]string** | A list of exposed entities of the security problem.   An exposed entity is an affected entity that is exposed to the internet. | [optional] [readonly] 
**ReachableDataAssets** | Pointer to **[]string** | A list of data assets reachable by affected entities of the security problem.   A data asset is a service that has database access. | [optional] [readonly] 
**RelatedEntities** | Pointer to [**RelatedEntitiesList**](RelatedEntitiesList.md) |  | [optional] 
**RelatedContainerImages** | Pointer to [**[]SecurityProblemDetailsRelatedContainerImages**](SecurityProblemDetailsRelatedContainerImages.md) | A list of related container images of the security problem.   A related container image is a container image that contains at least one *affected entity*. | [optional] [readonly] 

## Methods

### NewSecurityProblemDetails

`func NewSecurityProblemDetails() *SecurityProblemDetails`

NewSecurityProblemDetails instantiates a new SecurityProblemDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityProblemDetailsWithDefaults

`func NewSecurityProblemDetailsWithDefaults() *SecurityProblemDetails`

NewSecurityProblemDetailsWithDefaults instantiates a new SecurityProblemDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityProblemId

`func (o *SecurityProblemDetails) GetSecurityProblemId() string`

GetSecurityProblemId returns the SecurityProblemId field if non-nil, zero value otherwise.

### GetSecurityProblemIdOk

`func (o *SecurityProblemDetails) GetSecurityProblemIdOk() (*string, bool)`

GetSecurityProblemIdOk returns a tuple with the SecurityProblemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProblemId

`func (o *SecurityProblemDetails) SetSecurityProblemId(v string)`

SetSecurityProblemId sets SecurityProblemId field to given value.

### HasSecurityProblemId

`func (o *SecurityProblemDetails) HasSecurityProblemId() bool`

HasSecurityProblemId returns a boolean if a field has been set.

### GetDisplayId

`func (o *SecurityProblemDetails) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *SecurityProblemDetails) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *SecurityProblemDetails) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *SecurityProblemDetails) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityProblemDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityProblemDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityProblemDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityProblemDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMuted

`func (o *SecurityProblemDetails) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *SecurityProblemDetails) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *SecurityProblemDetails) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *SecurityProblemDetails) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetExternalVulnerabilityId

`func (o *SecurityProblemDetails) GetExternalVulnerabilityId() string`

GetExternalVulnerabilityId returns the ExternalVulnerabilityId field if non-nil, zero value otherwise.

### GetExternalVulnerabilityIdOk

`func (o *SecurityProblemDetails) GetExternalVulnerabilityIdOk() (*string, bool)`

GetExternalVulnerabilityIdOk returns a tuple with the ExternalVulnerabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVulnerabilityId

`func (o *SecurityProblemDetails) SetExternalVulnerabilityId(v string)`

SetExternalVulnerabilityId sets ExternalVulnerabilityId field to given value.

### HasExternalVulnerabilityId

`func (o *SecurityProblemDetails) HasExternalVulnerabilityId() bool`

HasExternalVulnerabilityId returns a boolean if a field has been set.

### GetVulnerabilityType

`func (o *SecurityProblemDetails) GetVulnerabilityType() string`

GetVulnerabilityType returns the VulnerabilityType field if non-nil, zero value otherwise.

### GetVulnerabilityTypeOk

`func (o *SecurityProblemDetails) GetVulnerabilityTypeOk() (*string, bool)`

GetVulnerabilityTypeOk returns a tuple with the VulnerabilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityType

`func (o *SecurityProblemDetails) SetVulnerabilityType(v string)`

SetVulnerabilityType sets VulnerabilityType field to given value.

### HasVulnerabilityType

`func (o *SecurityProblemDetails) HasVulnerabilityType() bool`

HasVulnerabilityType returns a boolean if a field has been set.

### GetTitle

`func (o *SecurityProblemDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SecurityProblemDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SecurityProblemDetails) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SecurityProblemDetails) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPackageName

`func (o *SecurityProblemDetails) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *SecurityProblemDetails) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *SecurityProblemDetails) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *SecurityProblemDetails) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetUrl

`func (o *SecurityProblemDetails) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SecurityProblemDetails) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SecurityProblemDetails) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SecurityProblemDetails) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDescription

`func (o *SecurityProblemDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityProblemDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityProblemDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityProblemDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTechnology

`func (o *SecurityProblemDetails) GetTechnology() string`

GetTechnology returns the Technology field if non-nil, zero value otherwise.

### GetTechnologyOk

`func (o *SecurityProblemDetails) GetTechnologyOk() (*string, bool)`

GetTechnologyOk returns a tuple with the Technology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnology

`func (o *SecurityProblemDetails) SetTechnology(v string)`

SetTechnology sets Technology field to given value.

### HasTechnology

`func (o *SecurityProblemDetails) HasTechnology() bool`

HasTechnology returns a boolean if a field has been set.

### GetFirstSeenTimestamp

`func (o *SecurityProblemDetails) GetFirstSeenTimestamp() int64`

GetFirstSeenTimestamp returns the FirstSeenTimestamp field if non-nil, zero value otherwise.

### GetFirstSeenTimestampOk

`func (o *SecurityProblemDetails) GetFirstSeenTimestampOk() (*int64, bool)`

GetFirstSeenTimestampOk returns a tuple with the FirstSeenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenTimestamp

`func (o *SecurityProblemDetails) SetFirstSeenTimestamp(v int64)`

SetFirstSeenTimestamp sets FirstSeenTimestamp field to given value.

### HasFirstSeenTimestamp

`func (o *SecurityProblemDetails) HasFirstSeenTimestamp() bool`

HasFirstSeenTimestamp returns a boolean if a field has been set.

### GetLastUpdatedTimestamp

`func (o *SecurityProblemDetails) GetLastUpdatedTimestamp() int64`

GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampOk

`func (o *SecurityProblemDetails) GetLastUpdatedTimestampOk() (*int64, bool)`

GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestamp

`func (o *SecurityProblemDetails) SetLastUpdatedTimestamp(v int64)`

SetLastUpdatedTimestamp sets LastUpdatedTimestamp field to given value.

### HasLastUpdatedTimestamp

`func (o *SecurityProblemDetails) HasLastUpdatedTimestamp() bool`

HasLastUpdatedTimestamp returns a boolean if a field has been set.

### GetRiskAssessment

`func (o *SecurityProblemDetails) GetRiskAssessment() RiskAssessment`

GetRiskAssessment returns the RiskAssessment field if non-nil, zero value otherwise.

### GetRiskAssessmentOk

`func (o *SecurityProblemDetails) GetRiskAssessmentOk() (*RiskAssessment, bool)`

GetRiskAssessmentOk returns a tuple with the RiskAssessment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAssessment

`func (o *SecurityProblemDetails) SetRiskAssessment(v RiskAssessment)`

SetRiskAssessment sets RiskAssessment field to given value.

### HasRiskAssessment

`func (o *SecurityProblemDetails) HasRiskAssessment() bool`

HasRiskAssessment returns a boolean if a field has been set.

### GetManagementZones

`func (o *SecurityProblemDetails) GetManagementZones() []ManagementZone`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *SecurityProblemDetails) GetManagementZonesOk() (*[]ManagementZone, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *SecurityProblemDetails) SetManagementZones(v []ManagementZone)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *SecurityProblemDetails) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetCveIds

`func (o *SecurityProblemDetails) GetCveIds() []string`

GetCveIds returns the CveIds field if non-nil, zero value otherwise.

### GetCveIdsOk

`func (o *SecurityProblemDetails) GetCveIdsOk() (*[]string, bool)`

GetCveIdsOk returns a tuple with the CveIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveIds

`func (o *SecurityProblemDetails) SetCveIds(v []string)`

SetCveIds sets CveIds field to given value.

### HasCveIds

`func (o *SecurityProblemDetails) HasCveIds() bool`

HasCveIds returns a boolean if a field has been set.

### GetEvents

`func (o *SecurityProblemDetails) GetEvents() []SecurityProblemEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SecurityProblemDetails) GetEventsOk() (*[]SecurityProblemEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SecurityProblemDetails) SetEvents(v []SecurityProblemEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SecurityProblemDetails) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetVulnerableComponents

`func (o *SecurityProblemDetails) GetVulnerableComponents() []VulnerableComponent`

GetVulnerableComponents returns the VulnerableComponents field if non-nil, zero value otherwise.

### GetVulnerableComponentsOk

`func (o *SecurityProblemDetails) GetVulnerableComponentsOk() (*[]VulnerableComponent, bool)`

GetVulnerableComponentsOk returns a tuple with the VulnerableComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableComponents

`func (o *SecurityProblemDetails) SetVulnerableComponents(v []VulnerableComponent)`

SetVulnerableComponents sets VulnerableComponents field to given value.

### HasVulnerableComponents

`func (o *SecurityProblemDetails) HasVulnerableComponents() bool`

HasVulnerableComponents returns a boolean if a field has been set.

### GetAffectedEntities

`func (o *SecurityProblemDetails) GetAffectedEntities() []string`

GetAffectedEntities returns the AffectedEntities field if non-nil, zero value otherwise.

### GetAffectedEntitiesOk

`func (o *SecurityProblemDetails) GetAffectedEntitiesOk() (*[]string, bool)`

GetAffectedEntitiesOk returns a tuple with the AffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEntities

`func (o *SecurityProblemDetails) SetAffectedEntities(v []string)`

SetAffectedEntities sets AffectedEntities field to given value.

### HasAffectedEntities

`func (o *SecurityProblemDetails) HasAffectedEntities() bool`

HasAffectedEntities returns a boolean if a field has been set.

### GetExposedEntities

`func (o *SecurityProblemDetails) GetExposedEntities() []string`

GetExposedEntities returns the ExposedEntities field if non-nil, zero value otherwise.

### GetExposedEntitiesOk

`func (o *SecurityProblemDetails) GetExposedEntitiesOk() (*[]string, bool)`

GetExposedEntitiesOk returns a tuple with the ExposedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedEntities

`func (o *SecurityProblemDetails) SetExposedEntities(v []string)`

SetExposedEntities sets ExposedEntities field to given value.

### HasExposedEntities

`func (o *SecurityProblemDetails) HasExposedEntities() bool`

HasExposedEntities returns a boolean if a field has been set.

### GetReachableDataAssets

`func (o *SecurityProblemDetails) GetReachableDataAssets() []string`

GetReachableDataAssets returns the ReachableDataAssets field if non-nil, zero value otherwise.

### GetReachableDataAssetsOk

`func (o *SecurityProblemDetails) GetReachableDataAssetsOk() (*[]string, bool)`

GetReachableDataAssetsOk returns a tuple with the ReachableDataAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachableDataAssets

`func (o *SecurityProblemDetails) SetReachableDataAssets(v []string)`

SetReachableDataAssets sets ReachableDataAssets field to given value.

### HasReachableDataAssets

`func (o *SecurityProblemDetails) HasReachableDataAssets() bool`

HasReachableDataAssets returns a boolean if a field has been set.

### GetRelatedEntities

`func (o *SecurityProblemDetails) GetRelatedEntities() RelatedEntitiesList`

GetRelatedEntities returns the RelatedEntities field if non-nil, zero value otherwise.

### GetRelatedEntitiesOk

`func (o *SecurityProblemDetails) GetRelatedEntitiesOk() (*RelatedEntitiesList, bool)`

GetRelatedEntitiesOk returns a tuple with the RelatedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedEntities

`func (o *SecurityProblemDetails) SetRelatedEntities(v RelatedEntitiesList)`

SetRelatedEntities sets RelatedEntities field to given value.

### HasRelatedEntities

`func (o *SecurityProblemDetails) HasRelatedEntities() bool`

HasRelatedEntities returns a boolean if a field has been set.

### GetRelatedContainerImages

`func (o *SecurityProblemDetails) GetRelatedContainerImages() []SecurityProblemDetailsRelatedContainerImages`

GetRelatedContainerImages returns the RelatedContainerImages field if non-nil, zero value otherwise.

### GetRelatedContainerImagesOk

`func (o *SecurityProblemDetails) GetRelatedContainerImagesOk() (*[]SecurityProblemDetailsRelatedContainerImages, bool)`

GetRelatedContainerImagesOk returns a tuple with the RelatedContainerImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedContainerImages

`func (o *SecurityProblemDetails) SetRelatedContainerImages(v []SecurityProblemDetailsRelatedContainerImages)`

SetRelatedContainerImages sets RelatedContainerImages field to given value.

### HasRelatedContainerImages

`func (o *SecurityProblemDetails) HasRelatedContainerImages() bool`

HasRelatedContainerImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


