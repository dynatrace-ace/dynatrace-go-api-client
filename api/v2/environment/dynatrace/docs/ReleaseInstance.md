# ReleaseInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** | The entity id of the instance. | [optional] 
**BuildVersion** | Pointer to **string** | The build version | [optional] 
**SecurityVulnerabilities** | Pointer to **[]string** | List of Security vulnerabilities Ids | [optional] 
**Problems** | Pointer to **[]string** | List of event Ids of open problems | [optional] 

## Methods

### NewReleaseInstance

`func NewReleaseInstance() *ReleaseInstance`

NewReleaseInstance instantiates a new ReleaseInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseInstanceWithDefaults

`func NewReleaseInstanceWithDefaults() *ReleaseInstance`

NewReleaseInstanceWithDefaults instantiates a new ReleaseInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *ReleaseInstance) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ReleaseInstance) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ReleaseInstance) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ReleaseInstance) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetBuildVersion

`func (o *ReleaseInstance) GetBuildVersion() string`

GetBuildVersion returns the BuildVersion field if non-nil, zero value otherwise.

### GetBuildVersionOk

`func (o *ReleaseInstance) GetBuildVersionOk() (*string, bool)`

GetBuildVersionOk returns a tuple with the BuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersion

`func (o *ReleaseInstance) SetBuildVersion(v string)`

SetBuildVersion sets BuildVersion field to given value.

### HasBuildVersion

`func (o *ReleaseInstance) HasBuildVersion() bool`

HasBuildVersion returns a boolean if a field has been set.

### GetSecurityVulnerabilities

`func (o *ReleaseInstance) GetSecurityVulnerabilities() []string`

GetSecurityVulnerabilities returns the SecurityVulnerabilities field if non-nil, zero value otherwise.

### GetSecurityVulnerabilitiesOk

`func (o *ReleaseInstance) GetSecurityVulnerabilitiesOk() (*[]string, bool)`

GetSecurityVulnerabilitiesOk returns a tuple with the SecurityVulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityVulnerabilities

`func (o *ReleaseInstance) SetSecurityVulnerabilities(v []string)`

SetSecurityVulnerabilities sets SecurityVulnerabilities field to given value.

### HasSecurityVulnerabilities

`func (o *ReleaseInstance) HasSecurityVulnerabilities() bool`

HasSecurityVulnerabilities returns a boolean if a field has been set.

### GetProblems

`func (o *ReleaseInstance) GetProblems() []string`

GetProblems returns the Problems field if non-nil, zero value otherwise.

### GetProblemsOk

`func (o *ReleaseInstance) GetProblemsOk() (*[]string, bool)`

GetProblemsOk returns a tuple with the Problems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblems

`func (o *ReleaseInstance) SetProblems(v []string)`

SetProblems sets Problems field to given value.

### HasProblems

`func (o *ReleaseInstance) HasProblems() bool`

HasProblems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


