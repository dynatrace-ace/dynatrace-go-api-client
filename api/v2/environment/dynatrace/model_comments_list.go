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

// CommentsList A list of comments.
type CommentsList struct {
	// The result entries.
	Comments *[]Comment `json:"comments,omitempty"`
	// The cursor for the next page of results. Has the value of `null` on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result.
	NextPageKey *string `json:"nextPageKey,omitempty"`
	// The total number of entries in the result.
	TotalCount int64 `json:"totalCount"`
	// The number of entries per page.
	PageSize *int32 `json:"pageSize,omitempty"`
}

// NewCommentsList instantiates a new CommentsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentsList(totalCount int64) *CommentsList {
	this := CommentsList{}
	this.TotalCount = totalCount
	return &this
}

// NewCommentsListWithDefaults instantiates a new CommentsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentsListWithDefaults() *CommentsList {
	this := CommentsList{}
	return &this
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *CommentsList) GetComments() []Comment {
	if o == nil || o.Comments == nil {
		var ret []Comment
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsList) GetCommentsOk() (*[]Comment, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *CommentsList) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given []Comment and assigns it to the Comments field.
func (o *CommentsList) SetComments(v []Comment) {
	o.Comments = &v
}

// GetNextPageKey returns the NextPageKey field value if set, zero value otherwise.
func (o *CommentsList) GetNextPageKey() string {
	if o == nil || o.NextPageKey == nil {
		var ret string
		return ret
	}
	return *o.NextPageKey
}

// GetNextPageKeyOk returns a tuple with the NextPageKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsList) GetNextPageKeyOk() (*string, bool) {
	if o == nil || o.NextPageKey == nil {
		return nil, false
	}
	return o.NextPageKey, true
}

// HasNextPageKey returns a boolean if a field has been set.
func (o *CommentsList) HasNextPageKey() bool {
	if o != nil && o.NextPageKey != nil {
		return true
	}

	return false
}

// SetNextPageKey gets a reference to the given string and assigns it to the NextPageKey field.
func (o *CommentsList) SetNextPageKey(v string) {
	o.NextPageKey = &v
}

// GetTotalCount returns the TotalCount field value
func (o *CommentsList) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *CommentsList) GetTotalCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *CommentsList) SetTotalCount(v int64) {
	o.TotalCount = v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *CommentsList) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsList) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *CommentsList) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *CommentsList) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o CommentsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.NextPageKey != nil {
		toSerialize["nextPageKey"] = o.NextPageKey
	}
	if true {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	return json.Marshal(toSerialize)
}

type NullableCommentsList struct {
	value *CommentsList
	isSet bool
}

func (v NullableCommentsList) Get() *CommentsList {
	return v.value
}

func (v *NullableCommentsList) Set(val *CommentsList) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentsList) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentsList(val *CommentsList) *NullableCommentsList {
	return &NullableCommentsList{value: val, isSet: true}
}

func (v NullableCommentsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

