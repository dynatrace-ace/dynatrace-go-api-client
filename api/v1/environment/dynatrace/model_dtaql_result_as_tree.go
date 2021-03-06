/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace Environment API v1. To read about use cases and examples, refer to the [help page](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// DtaqlResultAsTree The user session query result as a tree.
type DtaqlResultAsTree struct {
	// The extrapolation level of the result.   To improve performance, some results may be calculated from a subset of actual data. The extrapolation level indicates the share of actual data in the result.   The number is the denominator of a fraction and indicates the amount of actual data. The value `1` means that the result contains only the actual data. The value `4` means that result is calculated using 1/4 of the actual data.   If you need the analysis to be based on the actual data, reduce the timeframe of your query. For example, in case of extrapolation level of `4`, try to use 1/4 of the original timeframe.
	ExtrapolationLevel *int32 `json:"extrapolationLevel,omitempty"`
	// A list of columns in the additionalValues table.    Only present if the endpoint was called with `deepLinkFields=true` parameter.
	AdditionalColumnNames *[]string `json:"additionalColumnNames,omitempty"`
	// A list of data rows.    Each array element represents a row in the table of additionally linked fields.   The size of each data row and the order of the elements correspond to the **additionalColumnNames** content.   Only present if the endpoint was called with `deepLinkFields=true` parameter.
	AdditionalValues *[][]map[string]interface{} `json:"additionalValues,omitempty"`
	// Additional information about the query and the result, that helps to understand the query and how the result was calculated.   Only appears when the **explain** parameter is set to `true`.   Example: The number of results was limited to the default of 50. Use the `LIMIT` clause to increase or decrease this limit.
	Explanations *[]string `json:"explanations,omitempty"`
	// A list of branches of the tree.    Typically, these are fields from the `SELECT` clause, that have been used in the `GROUP BY` clause.
	BranchNames *[]string `json:"branchNames,omitempty"`
	// A list of leaves on each tree branch.    Typically, these are fields from the `SELECT` clause, that have not been used in the `GROUP BY` clause.
	LeafNames *[]string `json:"leafNames,omitempty"`
	// The user session query result as a tree.
	Values *map[string]interface{} `json:"values,omitempty"`
}

// NewDtaqlResultAsTree instantiates a new DtaqlResultAsTree object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtaqlResultAsTree() *DtaqlResultAsTree {
	this := DtaqlResultAsTree{}
	return &this
}

// NewDtaqlResultAsTreeWithDefaults instantiates a new DtaqlResultAsTree object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtaqlResultAsTreeWithDefaults() *DtaqlResultAsTree {
	this := DtaqlResultAsTree{}
	return &this
}

// GetExtrapolationLevel returns the ExtrapolationLevel field value if set, zero value otherwise.
func (o *DtaqlResultAsTree) GetExtrapolationLevel() int32 {
	if o == nil || o.ExtrapolationLevel == nil {
		var ret int32
		return ret
	}
	return *o.ExtrapolationLevel
}

// GetExtrapolationLevelOk returns a tuple with the ExtrapolationLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtaqlResultAsTree) GetExtrapolationLevelOk() (*int32, bool) {
	if o == nil || o.ExtrapolationLevel == nil {
		return nil, false
	}
	return o.ExtrapolationLevel, true
}

// HasExtrapolationLevel returns a boolean if a field has been set.
func (o *DtaqlResultAsTree) HasExtrapolationLevel() bool {
	if o != nil && o.ExtrapolationLevel != nil {
		return true
	}

	return false
}

// SetExtrapolationLevel gets a reference to the given int32 and assigns it to the ExtrapolationLevel field.
func (o *DtaqlResultAsTree) SetExtrapolationLevel(v int32) {
	o.ExtrapolationLevel = &v
}

// GetAdditionalColumnNames returns the AdditionalColumnNames field value if set, zero value otherwise.
func (o *DtaqlResultAsTree) GetAdditionalColumnNames() []string {
	if o == nil || o.AdditionalColumnNames == nil {
		var ret []string
		return ret
	}
	return *o.AdditionalColumnNames
}

// GetAdditionalColumnNamesOk returns a tuple with the AdditionalColumnNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtaqlResultAsTree) GetAdditionalColumnNamesOk() (*[]string, bool) {
	if o == nil || o.AdditionalColumnNames == nil {
		return nil, false
	}
	return o.AdditionalColumnNames, true
}

// HasAdditionalColumnNames returns a boolean if a field has been set.
func (o *DtaqlResultAsTree) HasAdditionalColumnNames() bool {
	if o != nil && o.AdditionalColumnNames != nil {
		return true
	}

	return false
}

// SetAdditionalColumnNames gets a reference to the given []string and assigns it to the AdditionalColumnNames field.
func (o *DtaqlResultAsTree) SetAdditionalColumnNames(v []string) {
	o.AdditionalColumnNames = &v
}

// GetAdditionalValues returns the AdditionalValues field value if set, zero value otherwise.
func (o *DtaqlResultAsTree) GetAdditionalValues() [][]map[string]interface{} {
	if o == nil || o.AdditionalValues == nil {
		var ret [][]map[string]interface{}
		return ret
	}
	return *o.AdditionalValues
}

// GetAdditionalValuesOk returns a tuple with the AdditionalValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtaqlResultAsTree) GetAdditionalValuesOk() (*[][]map[string]interface{}, bool) {
	if o == nil || o.AdditionalValues == nil {
		return nil, false
	}
	return o.AdditionalValues, true
}

// HasAdditionalValues returns a boolean if a field has been set.
func (o *DtaqlResultAsTree) HasAdditionalValues() bool {
	if o != nil && o.AdditionalValues != nil {
		return true
	}

	return false
}

// SetAdditionalValues gets a reference to the given [][]map[string]interface{} and assigns it to the AdditionalValues field.
func (o *DtaqlResultAsTree) SetAdditionalValues(v [][]map[string]interface{}) {
	o.AdditionalValues = &v
}

// GetExplanations returns the Explanations field value if set, zero value otherwise.
func (o *DtaqlResultAsTree) GetExplanations() []string {
	if o == nil || o.Explanations == nil {
		var ret []string
		return ret
	}
	return *o.Explanations
}

// GetExplanationsOk returns a tuple with the Explanations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtaqlResultAsTree) GetExplanationsOk() (*[]string, bool) {
	if o == nil || o.Explanations == nil {
		return nil, false
	}
	return o.Explanations, true
}

// HasExplanations returns a boolean if a field has been set.
func (o *DtaqlResultAsTree) HasExplanations() bool {
	if o != nil && o.Explanations != nil {
		return true
	}

	return false
}

// SetExplanations gets a reference to the given []string and assigns it to the Explanations field.
func (o *DtaqlResultAsTree) SetExplanations(v []string) {
	o.Explanations = &v
}

// GetBranchNames returns the BranchNames field value if set, zero value otherwise.
func (o *DtaqlResultAsTree) GetBranchNames() []string {
	if o == nil || o.BranchNames == nil {
		var ret []string
		return ret
	}
	return *o.BranchNames
}

// GetBranchNamesOk returns a tuple with the BranchNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtaqlResultAsTree) GetBranchNamesOk() (*[]string, bool) {
	if o == nil || o.BranchNames == nil {
		return nil, false
	}
	return o.BranchNames, true
}

// HasBranchNames returns a boolean if a field has been set.
func (o *DtaqlResultAsTree) HasBranchNames() bool {
	if o != nil && o.BranchNames != nil {
		return true
	}

	return false
}

// SetBranchNames gets a reference to the given []string and assigns it to the BranchNames field.
func (o *DtaqlResultAsTree) SetBranchNames(v []string) {
	o.BranchNames = &v
}

// GetLeafNames returns the LeafNames field value if set, zero value otherwise.
func (o *DtaqlResultAsTree) GetLeafNames() []string {
	if o == nil || o.LeafNames == nil {
		var ret []string
		return ret
	}
	return *o.LeafNames
}

// GetLeafNamesOk returns a tuple with the LeafNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtaqlResultAsTree) GetLeafNamesOk() (*[]string, bool) {
	if o == nil || o.LeafNames == nil {
		return nil, false
	}
	return o.LeafNames, true
}

// HasLeafNames returns a boolean if a field has been set.
func (o *DtaqlResultAsTree) HasLeafNames() bool {
	if o != nil && o.LeafNames != nil {
		return true
	}

	return false
}

// SetLeafNames gets a reference to the given []string and assigns it to the LeafNames field.
func (o *DtaqlResultAsTree) SetLeafNames(v []string) {
	o.LeafNames = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *DtaqlResultAsTree) GetValues() map[string]interface{} {
	if o == nil || o.Values == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtaqlResultAsTree) GetValuesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *DtaqlResultAsTree) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given map[string]interface{} and assigns it to the Values field.
func (o *DtaqlResultAsTree) SetValues(v map[string]interface{}) {
	o.Values = &v
}

func (o DtaqlResultAsTree) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExtrapolationLevel != nil {
		toSerialize["extrapolationLevel"] = o.ExtrapolationLevel
	}
	if o.AdditionalColumnNames != nil {
		toSerialize["additionalColumnNames"] = o.AdditionalColumnNames
	}
	if o.AdditionalValues != nil {
		toSerialize["additionalValues"] = o.AdditionalValues
	}
	if o.Explanations != nil {
		toSerialize["explanations"] = o.Explanations
	}
	if o.BranchNames != nil {
		toSerialize["branchNames"] = o.BranchNames
	}
	if o.LeafNames != nil {
		toSerialize["leafNames"] = o.LeafNames
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableDtaqlResultAsTree struct {
	value *DtaqlResultAsTree
	isSet bool
}

func (v NullableDtaqlResultAsTree) Get() *DtaqlResultAsTree {
	return v.value
}

func (v *NullableDtaqlResultAsTree) Set(val *DtaqlResultAsTree) {
	v.value = val
	v.isSet = true
}

func (v NullableDtaqlResultAsTree) IsSet() bool {
	return v.isSet
}

func (v *NullableDtaqlResultAsTree) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtaqlResultAsTree(val *DtaqlResultAsTree) *NullableDtaqlResultAsTree {
	return &NullableDtaqlResultAsTree{value: val, isSet: true}
}

func (v NullableDtaqlResultAsTree) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtaqlResultAsTree) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


