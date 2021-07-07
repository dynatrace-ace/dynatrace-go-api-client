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

// ProblemCloseResult The result of closing a problem.
type ProblemCloseResult struct {
	// The timestamp when the user triggered the closing.
	CloseTimestamp int64 `json:"closeTimestamp"`
	// The ID of the problem.
	ProblemId string `json:"problemId"`
	Comment *Comment `json:"comment,omitempty"`
	// True, if the problem is being closed.
	Closing bool `json:"closing"`
}

// NewProblemCloseResult instantiates a new ProblemCloseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemCloseResult(closeTimestamp int64, problemId string, closing bool) *ProblemCloseResult {
	this := ProblemCloseResult{}
	this.CloseTimestamp = closeTimestamp
	this.ProblemId = problemId
	this.Closing = closing
	return &this
}

// NewProblemCloseResultWithDefaults instantiates a new ProblemCloseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemCloseResultWithDefaults() *ProblemCloseResult {
	this := ProblemCloseResult{}
	return &this
}

// GetCloseTimestamp returns the CloseTimestamp field value
func (o *ProblemCloseResult) GetCloseTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CloseTimestamp
}

// GetCloseTimestampOk returns a tuple with the CloseTimestamp field value
// and a boolean to check if the value has been set.
func (o *ProblemCloseResult) GetCloseTimestampOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CloseTimestamp, true
}

// SetCloseTimestamp sets field value
func (o *ProblemCloseResult) SetCloseTimestamp(v int64) {
	o.CloseTimestamp = v
}

// GetProblemId returns the ProblemId field value
func (o *ProblemCloseResult) GetProblemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProblemId
}

// GetProblemIdOk returns a tuple with the ProblemId field value
// and a boolean to check if the value has been set.
func (o *ProblemCloseResult) GetProblemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProblemId, true
}

// SetProblemId sets field value
func (o *ProblemCloseResult) SetProblemId(v string) {
	o.ProblemId = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ProblemCloseResult) GetComment() Comment {
	if o == nil || o.Comment == nil {
		var ret Comment
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemCloseResult) GetCommentOk() (*Comment, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ProblemCloseResult) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given Comment and assigns it to the Comment field.
func (o *ProblemCloseResult) SetComment(v Comment) {
	o.Comment = &v
}

// GetClosing returns the Closing field value
func (o *ProblemCloseResult) GetClosing() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Closing
}

// GetClosingOk returns a tuple with the Closing field value
// and a boolean to check if the value has been set.
func (o *ProblemCloseResult) GetClosingOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Closing, true
}

// SetClosing sets field value
func (o *ProblemCloseResult) SetClosing(v bool) {
	o.Closing = v
}

func (o ProblemCloseResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["closeTimestamp"] = o.CloseTimestamp
	}
	if true {
		toSerialize["problemId"] = o.ProblemId
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if true {
		toSerialize["closing"] = o.Closing
	}
	return json.Marshal(toSerialize)
}

type NullableProblemCloseResult struct {
	value *ProblemCloseResult
	isSet bool
}

func (v NullableProblemCloseResult) Get() *ProblemCloseResult {
	return v.value
}

func (v *NullableProblemCloseResult) Set(val *ProblemCloseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemCloseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemCloseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemCloseResult(val *ProblemCloseResult) *NullableProblemCloseResult {
	return &NullableProblemCloseResult{value: val, isSet: true}
}

func (v NullableProblemCloseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemCloseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

