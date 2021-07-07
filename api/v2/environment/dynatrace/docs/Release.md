# Release

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseEntityId** | Pointer to **string** | The entity id of correlating release. | [optional] 
**SecurityVulnerabilitiesCount** | Pointer to **int32** | The number of security vulnerabilities of the entity | [optional] 
**AffectedByProblems** | Pointer to **bool** | The entity has one or more problems | [optional] 
**Instances** | Pointer to [**[]ReleaseInstance**](ReleaseInstance.md) | The instances entityIds included in this release | [optional] 
**AffectedBySecurityVulnerabilities** | Pointer to **bool** | The entity has one or more security vulnerabilities | [optional] 
**Throughput** | Pointer to **float64** | The count of bytes per second of the entity | [optional] 
**SoftwareTechs** | Pointer to [**[]SoftwareTechs**](SoftwareTechs.md) | The software technologies of the release | [optional] 
**Product** | Pointer to **string** | The product name | [optional] 
**Name** | Pointer to **string** | The entity name | [optional] 
**Version** | Pointer to **string** | The identified release version | [optional] 
**Running** | Pointer to **bool** | The related PGI is still running/monitored | [optional] 
**ProblemCount** | Pointer to **int32** | The number of problems of the entity | [optional] 
**Stage** | Pointer to **string** | The stage name | [optional] 

## Methods

### NewRelease

`func NewRelease() *Release`

NewRelease instantiates a new Release object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseWithDefaults

`func NewReleaseWithDefaults() *Release`

NewReleaseWithDefaults instantiates a new Release object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseEntityId

`func (o *Release) GetReleaseEntityId() string`

GetReleaseEntityId returns the ReleaseEntityId field if non-nil, zero value otherwise.

### GetReleaseEntityIdOk

`func (o *Release) GetReleaseEntityIdOk() (*string, bool)`

GetReleaseEntityIdOk returns a tuple with the ReleaseEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseEntityId

`func (o *Release) SetReleaseEntityId(v string)`

SetReleaseEntityId sets ReleaseEntityId field to given value.

### HasReleaseEntityId

`func (o *Release) HasReleaseEntityId() bool`

HasReleaseEntityId returns a boolean if a field has been set.

### GetSecurityVulnerabilitiesCount

`func (o *Release) GetSecurityVulnerabilitiesCount() int32`

GetSecurityVulnerabilitiesCount returns the SecurityVulnerabilitiesCount field if non-nil, zero value otherwise.

### GetSecurityVulnerabilitiesCountOk

`func (o *Release) GetSecurityVulnerabilitiesCountOk() (*int32, bool)`

GetSecurityVulnerabilitiesCountOk returns a tuple with the SecurityVulnerabilitiesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityVulnerabilitiesCount

`func (o *Release) SetSecurityVulnerabilitiesCount(v int32)`

SetSecurityVulnerabilitiesCount sets SecurityVulnerabilitiesCount field to given value.

### HasSecurityVulnerabilitiesCount

`func (o *Release) HasSecurityVulnerabilitiesCount() bool`

HasSecurityVulnerabilitiesCount returns a boolean if a field has been set.

### GetAffectedByProblems

`func (o *Release) GetAffectedByProblems() bool`

GetAffectedByProblems returns the AffectedByProblems field if non-nil, zero value otherwise.

### GetAffectedByProblemsOk

`func (o *Release) GetAffectedByProblemsOk() (*bool, bool)`

GetAffectedByProblemsOk returns a tuple with the AffectedByProblems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedByProblems

`func (o *Release) SetAffectedByProblems(v bool)`

SetAffectedByProblems sets AffectedByProblems field to given value.

### HasAffectedByProblems

`func (o *Release) HasAffectedByProblems() bool`

HasAffectedByProblems returns a boolean if a field has been set.

### GetInstances

`func (o *Release) GetInstances() []ReleaseInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *Release) GetInstancesOk() (*[]ReleaseInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *Release) SetInstances(v []ReleaseInstance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *Release) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetAffectedBySecurityVulnerabilities

`func (o *Release) GetAffectedBySecurityVulnerabilities() bool`

GetAffectedBySecurityVulnerabilities returns the AffectedBySecurityVulnerabilities field if non-nil, zero value otherwise.

### GetAffectedBySecurityVulnerabilitiesOk

`func (o *Release) GetAffectedBySecurityVulnerabilitiesOk() (*bool, bool)`

GetAffectedBySecurityVulnerabilitiesOk returns a tuple with the AffectedBySecurityVulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedBySecurityVulnerabilities

`func (o *Release) SetAffectedBySecurityVulnerabilities(v bool)`

SetAffectedBySecurityVulnerabilities sets AffectedBySecurityVulnerabilities field to given value.

### HasAffectedBySecurityVulnerabilities

`func (o *Release) HasAffectedBySecurityVulnerabilities() bool`

HasAffectedBySecurityVulnerabilities returns a boolean if a field has been set.

### GetThroughput

`func (o *Release) GetThroughput() float64`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *Release) GetThroughputOk() (*float64, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *Release) SetThroughput(v float64)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *Release) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetSoftwareTechs

`func (o *Release) GetSoftwareTechs() []SoftwareTechs`

GetSoftwareTechs returns the SoftwareTechs field if non-nil, zero value otherwise.

### GetSoftwareTechsOk

`func (o *Release) GetSoftwareTechsOk() (*[]SoftwareTechs, bool)`

GetSoftwareTechsOk returns a tuple with the SoftwareTechs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTechs

`func (o *Release) SetSoftwareTechs(v []SoftwareTechs)`

SetSoftwareTechs sets SoftwareTechs field to given value.

### HasSoftwareTechs

`func (o *Release) HasSoftwareTechs() bool`

HasSoftwareTechs returns a boolean if a field has been set.

### GetProduct

`func (o *Release) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Release) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Release) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *Release) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetName

`func (o *Release) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Release) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Release) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Release) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *Release) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Release) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Release) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Release) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRunning

`func (o *Release) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *Release) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *Release) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *Release) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetProblemCount

`func (o *Release) GetProblemCount() int32`

GetProblemCount returns the ProblemCount field if non-nil, zero value otherwise.

### GetProblemCountOk

`func (o *Release) GetProblemCountOk() (*int32, bool)`

GetProblemCountOk returns a tuple with the ProblemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemCount

`func (o *Release) SetProblemCount(v int32)`

SetProblemCount sets ProblemCount field to given value.

### HasProblemCount

`func (o *Release) HasProblemCount() bool`

HasProblemCount returns a boolean if a field has been set.

### GetStage

`func (o *Release) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *Release) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *Release) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *Release) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


