# RiskAssessment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskLevel** | Pointer to **string** | The Davis risk level.    It is calculated by Dynatrace on the basis of CVSS score. | [optional] [readonly] 
**RiskScore** | Pointer to **float32** | The Davis risk score (1-10).    It is calculated by Dynatrace on the basis of CVSS score. | [optional] [readonly] 
**RiskVector** | Pointer to **string** | The attack vector calculated by DT based on the CVSS attack vector. | [optional] [readonly] 
**BaseRiskLevel** | Pointer to **string** | The risk level from the CVSS score. | [optional] [readonly] 
**BaseRiskScore** | Pointer to **float32** | The risk score (1-10) from the CVSS score. | [optional] [readonly] 
**BaseRiskVector** | Pointer to **string** | The original attack vector of the CVSS assessment. | [optional] [readonly] 
**Exposure** | Pointer to **string** | The level of exposure of affected entities. | [optional] [readonly] 
**DataAssets** | Pointer to **string** | The reachability of related data assets by affected entities. | [optional] [readonly] 
**PublicExploit** | Pointer to **string** | The availability status of public exploits. | [optional] [readonly] 

## Methods

### NewRiskAssessment

`func NewRiskAssessment() *RiskAssessment`

NewRiskAssessment instantiates a new RiskAssessment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAssessmentWithDefaults

`func NewRiskAssessmentWithDefaults() *RiskAssessment`

NewRiskAssessmentWithDefaults instantiates a new RiskAssessment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskLevel

`func (o *RiskAssessment) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *RiskAssessment) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *RiskAssessment) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *RiskAssessment) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.

### GetRiskScore

`func (o *RiskAssessment) GetRiskScore() float32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *RiskAssessment) GetRiskScoreOk() (*float32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *RiskAssessment) SetRiskScore(v float32)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *RiskAssessment) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetRiskVector

`func (o *RiskAssessment) GetRiskVector() string`

GetRiskVector returns the RiskVector field if non-nil, zero value otherwise.

### GetRiskVectorOk

`func (o *RiskAssessment) GetRiskVectorOk() (*string, bool)`

GetRiskVectorOk returns a tuple with the RiskVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskVector

`func (o *RiskAssessment) SetRiskVector(v string)`

SetRiskVector sets RiskVector field to given value.

### HasRiskVector

`func (o *RiskAssessment) HasRiskVector() bool`

HasRiskVector returns a boolean if a field has been set.

### GetBaseRiskLevel

`func (o *RiskAssessment) GetBaseRiskLevel() string`

GetBaseRiskLevel returns the BaseRiskLevel field if non-nil, zero value otherwise.

### GetBaseRiskLevelOk

`func (o *RiskAssessment) GetBaseRiskLevelOk() (*string, bool)`

GetBaseRiskLevelOk returns a tuple with the BaseRiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseRiskLevel

`func (o *RiskAssessment) SetBaseRiskLevel(v string)`

SetBaseRiskLevel sets BaseRiskLevel field to given value.

### HasBaseRiskLevel

`func (o *RiskAssessment) HasBaseRiskLevel() bool`

HasBaseRiskLevel returns a boolean if a field has been set.

### GetBaseRiskScore

`func (o *RiskAssessment) GetBaseRiskScore() float32`

GetBaseRiskScore returns the BaseRiskScore field if non-nil, zero value otherwise.

### GetBaseRiskScoreOk

`func (o *RiskAssessment) GetBaseRiskScoreOk() (*float32, bool)`

GetBaseRiskScoreOk returns a tuple with the BaseRiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseRiskScore

`func (o *RiskAssessment) SetBaseRiskScore(v float32)`

SetBaseRiskScore sets BaseRiskScore field to given value.

### HasBaseRiskScore

`func (o *RiskAssessment) HasBaseRiskScore() bool`

HasBaseRiskScore returns a boolean if a field has been set.

### GetBaseRiskVector

`func (o *RiskAssessment) GetBaseRiskVector() string`

GetBaseRiskVector returns the BaseRiskVector field if non-nil, zero value otherwise.

### GetBaseRiskVectorOk

`func (o *RiskAssessment) GetBaseRiskVectorOk() (*string, bool)`

GetBaseRiskVectorOk returns a tuple with the BaseRiskVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseRiskVector

`func (o *RiskAssessment) SetBaseRiskVector(v string)`

SetBaseRiskVector sets BaseRiskVector field to given value.

### HasBaseRiskVector

`func (o *RiskAssessment) HasBaseRiskVector() bool`

HasBaseRiskVector returns a boolean if a field has been set.

### GetExposure

`func (o *RiskAssessment) GetExposure() string`

GetExposure returns the Exposure field if non-nil, zero value otherwise.

### GetExposureOk

`func (o *RiskAssessment) GetExposureOk() (*string, bool)`

GetExposureOk returns a tuple with the Exposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposure

`func (o *RiskAssessment) SetExposure(v string)`

SetExposure sets Exposure field to given value.

### HasExposure

`func (o *RiskAssessment) HasExposure() bool`

HasExposure returns a boolean if a field has been set.

### GetDataAssets

`func (o *RiskAssessment) GetDataAssets() string`

GetDataAssets returns the DataAssets field if non-nil, zero value otherwise.

### GetDataAssetsOk

`func (o *RiskAssessment) GetDataAssetsOk() (*string, bool)`

GetDataAssetsOk returns a tuple with the DataAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAssets

`func (o *RiskAssessment) SetDataAssets(v string)`

SetDataAssets sets DataAssets field to given value.

### HasDataAssets

`func (o *RiskAssessment) HasDataAssets() bool`

HasDataAssets returns a boolean if a field has been set.

### GetPublicExploit

`func (o *RiskAssessment) GetPublicExploit() string`

GetPublicExploit returns the PublicExploit field if non-nil, zero value otherwise.

### GetPublicExploitOk

`func (o *RiskAssessment) GetPublicExploitOk() (*string, bool)`

GetPublicExploitOk returns a tuple with the PublicExploit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicExploit

`func (o *RiskAssessment) SetPublicExploit(v string)`

SetPublicExploit sets PublicExploit field to given value.

### HasPublicExploit

`func (o *RiskAssessment) HasPublicExploit() bool`

HasPublicExploit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


