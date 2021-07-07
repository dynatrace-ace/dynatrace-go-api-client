/*
 * Dynatrace Environment API
 *
 *  Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// Release Contains data related to a single release of a component. A Release is a combination of a component and a version. A Component can be any form of deployable that can be associated with a version. In the first draft, a Component is always a Service.  The tuple <name, product, stage, version> is always unique.
type Release struct {
	// The entity id of correlating release.
	ReleaseEntityId *string `json:"releaseEntityId,omitempty"`
	// The number of security vulnerabilities of the entity
	SecurityVulnerabilitiesCount *int32 `json:"securityVulnerabilitiesCount,omitempty"`
	// The entity has one or more problems
	AffectedByProblems *bool `json:"affectedByProblems,omitempty"`
	// The instances entityIds included in this release
	Instances *[]ReleaseInstance `json:"instances,omitempty"`
	// The entity has one or more security vulnerabilities
	AffectedBySecurityVulnerabilities *bool `json:"affectedBySecurityVulnerabilities,omitempty"`
	// The count of bytes per second of the entity
	Throughput *float64 `json:"throughput,omitempty"`
	// The software technologies of the release
	SoftwareTechs *[]SoftwareTechs `json:"softwareTechs,omitempty"`
	// The product name
	Product *string `json:"product,omitempty"`
	// The entity name
	Name *string `json:"name,omitempty"`
	// The identified release version
	Version *string `json:"version,omitempty"`
	// The related PGI is still running/monitored
	Running *bool `json:"running,omitempty"`
	// The number of problems of the entity
	ProblemCount *int32 `json:"problemCount,omitempty"`
	// The stage name
	Stage *string `json:"stage,omitempty"`
}

// NewRelease instantiates a new Release object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelease() *Release {
	this := Release{}
	return &this
}

// NewReleaseWithDefaults instantiates a new Release object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseWithDefaults() *Release {
	this := Release{}
	return &this
}

// GetReleaseEntityId returns the ReleaseEntityId field value if set, zero value otherwise.
func (o *Release) GetReleaseEntityId() string {
	if o == nil || o.ReleaseEntityId == nil {
		var ret string
		return ret
	}
	return *o.ReleaseEntityId
}

// GetReleaseEntityIdOk returns a tuple with the ReleaseEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Release) GetReleaseEntityIdOk() (*string, bool) {
	if o == nil || o.ReleaseEntityId == nil {
		return nil, false
	}
	return o.ReleaseEntityId, true
}

// HasReleaseEntityId returns a boolean if a field has been set.
func (o *Release) HasReleaseEntityId() bool {
	if o != nil && o.ReleaseEntityId != nil {
		return true
	}

	return false
}

// SetReleaseEntityId gets a reference to the given string and assigns it to the ReleaseEntityId field.
func (o *Release) SetReleaseEntityId(v string) {
	o.ReleaseEntityId = &v
}

// GetSecurityVulnerabilitiesCount returns the SecurityVulnerabilitiesCount field value if set, zero value otherwise.
func (o *Release) GetSecurityVulnerabilitiesCount() int32 {
	if o == nil || o.SecurityVulnerabilitiesCount == nil {
		var ret int32
		return ret
	}
	return *o.SecurityVulnerabilitiesCount
}

// GetSecurityVulnerabilitiesCountOk returns a tuple with the SecurityVulnerabilitiesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Release) GetSecurityVulnerabilitiesCountOk() (*int32, bool) {
	if o == nil || o.SecurityVulnerabilitiesCount == nil {
		return nil, false
	}
	return o.SecurityVulnerabilitiesCount, true
}

// HasSecurityVulnerabilitiesCount returns a boolean if a field has been set.
func (o *Release) HasSecurityVulnerabilitiesCount() bool {
	if o != nil && o.SecurityVulnerabilitiesCount != nil {
		return true
	}

	return false
}

// SetSecurityVulnerabilitiesCount gets a reference to the given int32 and assigns it to the SecurityVulnerabilitiesCount field.
func (o *Release) SetSecurityVulnerabilitiesCount(v int32) {
	o.SecurityVulnerabilitiesCount = &v
}

// GetAffectedByProblems returns the AffectedByProblems field value if set, zero value otherwise.
func (o *Release) GetAffectedByProblems() bool {
	if o == nil || o.AffectedByProblems == nil {
		var ret bool
		return ret
	}
	return *o.AffectedByProblems
}

// GetAffectedByProblemsOk returns a tuple with the AffectedByProblems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Release) GetAffectedByProblemsOk() (*bool, bool) {
	if o == nil || o.AffectedByProblems == nil {
		return nil, false
	}
	return o.AffectedByProblems, true
}

// HasAffectedByProblems returns a boolean if a field has been set.
func (o *Release) HasAffectedByProblems() bool {
	if o != nil && o.AffectedByProblems != nil {
		return true
	}

	return false
}

// SetAffectedByProblems gets a reference to the given bool and assigns it to the AffectedByProblems field.
func (o *Release) SetAffectedByProblems(v bool) {
	o.AffectedByProblems = &v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *Release) GetInstances() []ReleaseInstance {
	if o == nil || o.Instances == nil {
		var ret []ReleaseInstance
		return ret
	}
	return *o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Release) GetInstancesOk() (*[]ReleaseInstance, bool) {
	if o == nil || o.Instances == nil {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *Release) HasInstances() bool {
	if o != nil && o.Instances != nil {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []ReleaseInstance and assigns it to the Instances field.
func (o *Release) SetInstances(v []ReleaseInstance) {
	o.Instances = &v
}

// GetAffectedBySecurityVulnerabilities returns the AffectedBySecurityVulnerabilities field value if set, zero value otherwise.
func (o *Release) GetAffectedBySecurityVulnerabilities() bool {
	if o == nil || o.AffectedBySecurityVulnerabilities == nil {
		var ret bool
		return ret
	}
	return *o.AffectedBySecurityVulnerabilities
}

// GetAffectedBySecurityVulnerabilitiesOk returns a tuple with the AffectedBySecurityVulnerabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Release) GetAffectedBySecurityVulnerabilitiesOk() (*bool, bool) {
	if o == nil || o.AffectedBySecurityVulnerabilities == nil {
		return nil, false
	}
	return o.AffectedBySecurityVulnerabilities, true
}

// HasAffectedBySecurityVulnerabilities returns a boolean if a field has been set.
func (o *Release) HasAffectedBySecurityVulnerabilities() bool {
	if o != nil && o.AffectedBySecurityVulnerabilities != nil {
		return true
	}

	return false
}

// SetAffectedBySecurityVulnerabilities gets a reference to the given bool and assigns it to the AffectedBySecurityVulnerabilities field.
func (o *Release) SetAffectedBySecurityVulnerabilities(v bool) {
	o.AffectedBySecurityVulnerabilities = &v
}

// GetThroughput returns the Throughput field value if set, zero value otherwise.
func (o *Release) GetThroughput() float64 {
	if o == nil || o.Throughput == nil {
		var ret float64
		return ret
	}
	return *o.Throughput
}

// GetThroughputOk returns a tuple with the Throughput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Release) GetThroughputOk() (*float64, bool) {
	if o == nil || o.Throughput == nil {
		return nil, false
	}
	return o.Throughput, true
}

// HasThroughput returns a boolean if a field has been set.
func (o *Release) HasThroughput() bool {
	if o != nil && o.Throughput != nil {
		return true
	}

	return false
}

// SetThroughput gets a reference to the given float64 and assigns it to the Throughput field.
func (o *Release) SetThroughput(v float64) {
	o.Throughput = &v
}

// GetSoftwareTechs returns the SoftwareTechs field value if set, zero value otherwise.
func (o *Release) GetSoftwareTechs() []SoftwareTechs {
	if o == nil || o.SoftwareTechs == nil {
		var ret []SoftwareTechs
		return ret
	}
	return *o.SoftwareTechs
}

// GetSoftwareTechsOk returns a tuple with the SoftwareTechs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Release) GetSoftwareTechsOk() (*[]SoftwareTechs, bool) {
	if o == nil || o.SoftwareTechs == nil {
		return nil, false
	}
	return o.SoftwareTechs, true
}

// HasSoftwareTechs returns a boolean if a field has been set.
func (o *Release) HasSoftwareTechs() bool {
	if o != nil && o.SoftwareTechs != nil {
		return true
	}

	return false
}

// SetSoftwareTechs gets a reference to the given []SoftwareTechs and assigns it to the SoftwareTechs field.
func (o *Release) SetSoftwareTechs(v []SoftwareTechs) {
	o.SoftwareTechs = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *Release) GetProduct() string {
	if o == nil || o.Product == nil {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Release) GetProductOk() (*string, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *Release) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *Release) SetProduct(v string) {
	o.Product = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Release) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Release) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Release) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Release) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Release) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Release) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Release) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Release) SetVersion(v string) {
	o.Version = &v
}

// GetRunning returns the Running field value if set, zero value otherwise.
func (o *Release) GetRunning() bool {
	if o == nil || o.Running == nil {
		var ret bool
		return ret
	}
	return *o.Running
}

// GetRunningOk returns a tuple with the Running field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Release) GetRunningOk() (*bool, bool) {
	if o == nil || o.Running == nil {
		return nil, false
	}
	return o.Running, true
}

// HasRunning returns a boolean if a field has been set.
func (o *Release) HasRunning() bool {
	if o != nil && o.Running != nil {
		return true
	}

	return false
}

// SetRunning gets a reference to the given bool and assigns it to the Running field.
func (o *Release) SetRunning(v bool) {
	o.Running = &v
}

// GetProblemCount returns the ProblemCount field value if set, zero value otherwise.
func (o *Release) GetProblemCount() int32 {
	if o == nil || o.ProblemCount == nil {
		var ret int32
		return ret
	}
	return *o.ProblemCount
}

// GetProblemCountOk returns a tuple with the ProblemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Release) GetProblemCountOk() (*int32, bool) {
	if o == nil || o.ProblemCount == nil {
		return nil, false
	}
	return o.ProblemCount, true
}

// HasProblemCount returns a boolean if a field has been set.
func (o *Release) HasProblemCount() bool {
	if o != nil && o.ProblemCount != nil {
		return true
	}

	return false
}

// SetProblemCount gets a reference to the given int32 and assigns it to the ProblemCount field.
func (o *Release) SetProblemCount(v int32) {
	o.ProblemCount = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *Release) GetStage() string {
	if o == nil || o.Stage == nil {
		var ret string
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Release) GetStageOk() (*string, bool) {
	if o == nil || o.Stage == nil {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *Release) HasStage() bool {
	if o != nil && o.Stage != nil {
		return true
	}

	return false
}

// SetStage gets a reference to the given string and assigns it to the Stage field.
func (o *Release) SetStage(v string) {
	o.Stage = &v
}

func (o Release) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReleaseEntityId != nil {
		toSerialize["releaseEntityId"] = o.ReleaseEntityId
	}
	if o.SecurityVulnerabilitiesCount != nil {
		toSerialize["securityVulnerabilitiesCount"] = o.SecurityVulnerabilitiesCount
	}
	if o.AffectedByProblems != nil {
		toSerialize["affectedByProblems"] = o.AffectedByProblems
	}
	if o.Instances != nil {
		toSerialize["instances"] = o.Instances
	}
	if o.AffectedBySecurityVulnerabilities != nil {
		toSerialize["affectedBySecurityVulnerabilities"] = o.AffectedBySecurityVulnerabilities
	}
	if o.Throughput != nil {
		toSerialize["throughput"] = o.Throughput
	}
	if o.SoftwareTechs != nil {
		toSerialize["softwareTechs"] = o.SoftwareTechs
	}
	if o.Product != nil {
		toSerialize["product"] = o.Product
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Running != nil {
		toSerialize["running"] = o.Running
	}
	if o.ProblemCount != nil {
		toSerialize["problemCount"] = o.ProblemCount
	}
	if o.Stage != nil {
		toSerialize["stage"] = o.Stage
	}
	return json.Marshal(toSerialize)
}

type NullableRelease struct {
	value *Release
	isSet bool
}

func (v NullableRelease) Get() *Release {
	return v.value
}

func (v *NullableRelease) Set(val *Release) {
	v.value = val
	v.isSet = true
}

func (v NullableRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelease(val *Release) *NullableRelease {
	return &NullableRelease{value: val, isSet: true}
}

func (v NullableRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

