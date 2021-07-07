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

// CommentRequestDtoImpl struct for CommentRequestDtoImpl
type CommentRequestDtoImpl struct {
	// The text of the comment.
	Message string `json:"message"`
	// The context of the comment.
	Context *string `json:"context,omitempty"`
}

// NewCommentRequestDtoImpl instantiates a new CommentRequestDtoImpl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentRequestDtoImpl(message string) *CommentRequestDtoImpl {
	this := CommentRequestDtoImpl{}
	this.Message = message
	return &this
}

// NewCommentRequestDtoImplWithDefaults instantiates a new CommentRequestDtoImpl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentRequestDtoImplWithDefaults() *CommentRequestDtoImpl {
	this := CommentRequestDtoImpl{}
	return &this
}

// GetMessage returns the Message field value
func (o *CommentRequestDtoImpl) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CommentRequestDtoImpl) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CommentRequestDtoImpl) SetMessage(v string) {
	o.Message = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *CommentRequestDtoImpl) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentRequestDtoImpl) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *CommentRequestDtoImpl) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *CommentRequestDtoImpl) SetContext(v string) {
	o.Context = &v
}

func (o CommentRequestDtoImpl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	return json.Marshal(toSerialize)
}

type NullableCommentRequestDtoImpl struct {
	value *CommentRequestDtoImpl
	isSet bool
}

func (v NullableCommentRequestDtoImpl) Get() *CommentRequestDtoImpl {
	return v.value
}

func (v *NullableCommentRequestDtoImpl) Set(val *CommentRequestDtoImpl) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentRequestDtoImpl) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentRequestDtoImpl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentRequestDtoImpl(val *CommentRequestDtoImpl) *NullableCommentRequestDtoImpl {
	return &NullableCommentRequestDtoImpl{value: val, isSet: true}
}

func (v NullableCommentRequestDtoImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentRequestDtoImpl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

