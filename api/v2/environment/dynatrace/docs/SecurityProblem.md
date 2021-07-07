# SecurityProblem

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
**Technology** | Pointer to **string** | The technology of the security problem. | [optional] [readonly] 
**FirstSeenTimestamp** | Pointer to **int64** | The timestamp of the first occurrence of the security problem. | [optional] [readonly] 
**LastUpdatedTimestamp** | Pointer to **int64** | The timestamp of the most recent security problem change. | [optional] [readonly] 
**RiskAssessment** | Pointer to [**RiskAssessment**](RiskAssessment.md) |  | [optional] 
**ManagementZones** | Pointer to [**[]ManagementZone**](ManagementZone.md) | Management zones to which the affected entities belong. | [optional] [readonly] 
**CveIds** | Pointer to **[]string** | CVE IDs of the security problem. | [optional] [readonly] 

## Methods

### NewSecurityProblem

`func NewSecurityProblem() *SecurityProblem`

NewSecurityProblem instantiates a new SecurityProblem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityProblemWithDefaults

`func NewSecurityProblemWithDefaults() *SecurityProblem`

NewSecurityProblemWithDefaults instantiates a new SecurityProblem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityProblemId

`func (o *SecurityProblem) GetSecurityProblemId() string`

GetSecurityProblemId returns the SecurityProblemId field if non-nil, zero value otherwise.

### GetSecurityProblemIdOk

`func (o *SecurityProblem) GetSecurityProblemIdOk() (*string, bool)`

GetSecurityProblemIdOk returns a tuple with the SecurityProblemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProblemId

`func (o *SecurityProblem) SetSecurityProblemId(v string)`

SetSecurityProblemId sets SecurityProblemId field to given value.

### HasSecurityProblemId

`func (o *SecurityProblem) HasSecurityProblemId() bool`

HasSecurityProblemId returns a boolean if a field has been set.

### GetDisplayId

`func (o *SecurityProblem) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *SecurityProblem) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *SecurityProblem) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *SecurityProblem) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityProblem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityProblem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityProblem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityProblem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMuted

`func (o *SecurityProblem) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *SecurityProblem) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *SecurityProblem) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *SecurityProblem) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetExternalVulnerabilityId

`func (o *SecurityProblem) GetExternalVulnerabilityId() string`

GetExternalVulnerabilityId returns the ExternalVulnerabilityId field if non-nil, zero value otherwise.

### GetExternalVulnerabilityIdOk

`func (o *SecurityProblem) GetExternalVulnerabilityIdOk() (*string, bool)`

GetExternalVulnerabilityIdOk returns a tuple with the ExternalVulnerabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVulnerabilityId

`func (o *SecurityProblem) SetExternalVulnerabilityId(v string)`

SetExternalVulnerabilityId sets ExternalVulnerabilityId field to given value.

### HasExternalVulnerabilityId

`func (o *SecurityProblem) HasExternalVulnerabilityId() bool`

HasExternalVulnerabilityId returns a boolean if a field has been set.

### GetVulnerabilityType

`func (o *SecurityProblem) GetVulnerabilityType() string`

GetVulnerabilityType returns the VulnerabilityType field if non-nil, zero value otherwise.

### GetVulnerabilityTypeOk

`func (o *SecurityProblem) GetVulnerabilityTypeOk() (*string, bool)`

GetVulnerabilityTypeOk returns a tuple with the VulnerabilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityType

`func (o *SecurityProblem) SetVulnerabilityType(v string)`

SetVulnerabilityType sets VulnerabilityType field to given value.

### HasVulnerabilityType

`func (o *SecurityProblem) HasVulnerabilityType() bool`

HasVulnerabilityType returns a boolean if a field has been set.

### GetTitle

`func (o *SecurityProblem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SecurityProblem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SecurityProblem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SecurityProblem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPackageName

`func (o *SecurityProblem) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *SecurityProblem) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *SecurityProblem) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *SecurityProblem) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetUrl

`func (o *SecurityProblem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SecurityProblem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SecurityProblem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SecurityProblem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTechnology

`func (o *SecurityProblem) GetTechnology() string`

GetTechnology returns the Technology field if non-nil, zero value otherwise.

### GetTechnologyOk

`func (o *SecurityProblem) GetTechnologyOk() (*string, bool)`

GetTechnologyOk returns a tuple with the Technology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnology

`func (o *SecurityProblem) SetTechnology(v string)`

SetTechnology sets Technology field to given value.

### HasTechnology

`func (o *SecurityProblem) HasTechnology() bool`

HasTechnology returns a boolean if a field has been set.

### GetFirstSeenTimestamp

`func (o *SecurityProblem) GetFirstSeenTimestamp() int64`

GetFirstSeenTimestamp returns the FirstSeenTimestamp field if non-nil, zero value otherwise.

### GetFirstSeenTimestampOk

`func (o *SecurityProblem) GetFirstSeenTimestampOk() (*int64, bool)`

GetFirstSeenTimestampOk returns a tuple with the FirstSeenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenTimestamp

`func (o *SecurityProblem) SetFirstSeenTimestamp(v int64)`

SetFirstSeenTimestamp sets FirstSeenTimestamp field to given value.

### HasFirstSeenTimestamp

`func (o *SecurityProblem) HasFirstSeenTimestamp() bool`

HasFirstSeenTimestamp returns a boolean if a field has been set.

### GetLastUpdatedTimestamp

`func (o *SecurityProblem) GetLastUpdatedTimestamp() int64`

GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampOk

`func (o *SecurityProblem) GetLastUpdatedTimestampOk() (*int64, bool)`

GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestamp

`func (o *SecurityProblem) SetLastUpdatedTimestamp(v int64)`

SetLastUpdatedTimestamp sets LastUpdatedTimestamp field to given value.

### HasLastUpdatedTimestamp

`func (o *SecurityProblem) HasLastUpdatedTimestamp() bool`

HasLastUpdatedTimestamp returns a boolean if a field has been set.

### GetRiskAssessment

`func (o *SecurityProblem) GetRiskAssessment() RiskAssessment`

GetRiskAssessment returns the RiskAssessment field if non-nil, zero value otherwise.

### GetRiskAssessmentOk

`func (o *SecurityProblem) GetRiskAssessmentOk() (*RiskAssessment, bool)`

GetRiskAssessmentOk returns a tuple with the RiskAssessment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAssessment

`func (o *SecurityProblem) SetRiskAssessment(v RiskAssessment)`

SetRiskAssessment sets RiskAssessment field to given value.

### HasRiskAssessment

`func (o *SecurityProblem) HasRiskAssessment() bool`

HasRiskAssessment returns a boolean if a field has been set.

### GetManagementZones

`func (o *SecurityProblem) GetManagementZones() []ManagementZone`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *SecurityProblem) GetManagementZonesOk() (*[]ManagementZone, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *SecurityProblem) SetManagementZones(v []ManagementZone)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *SecurityProblem) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetCveIds

`func (o *SecurityProblem) GetCveIds() []string`

GetCveIds returns the CveIds field if non-nil, zero value otherwise.

### GetCveIdsOk

`func (o *SecurityProblem) GetCveIdsOk() (*[]string, bool)`

GetCveIdsOk returns a tuple with the CveIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveIds

`func (o *SecurityProblem) SetCveIds(v []string)`

SetCveIds sets CveIds field to given value.

### HasCveIds

`func (o *SecurityProblem) HasCveIds() bool`

HasCveIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


