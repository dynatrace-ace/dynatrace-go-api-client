/*
 * Dynatrace Cluster API
 *
 * Dynatrace Managed exposes cluster management functionality via REST endpoints. This interactive documentation also acts as a REST client you can use to interact with Dynatrace Managed clusters.   To authorize, use the API Token generated in [Settings - API Tokens page](/cmc#cm/apiToken). The HTTP status code of the response shows the result of your request. The expected response code for a successful request is documented individually per REST endpoint. Additionally the following error response codes can occur in our REST interface:  * 400 - Bad Request: Some request parameters are not correct. See response body for details. * 401 - Unauthorized: A valid authorization header is required but is missing. * 403 - Forbidden: Execution of request is not allowed, e.g. api-token is invalid. * 404 - Not Found: Endpoint does not exist or some entities could not be found, e.g. User account. * 500 - Internal Server Error: See response body for details. * 556 - Upgrade in progress: Operation couldn't be performed during the upgrade. 
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// LdapGroupsQueryDtoImpl struct for LdapGroupsQueryDtoImpl
type LdapGroupsQueryDtoImpl struct {
	GroupsDn *string `json:"groupsDn,omitempty"`
	GroupsFilter *string `json:"groupsFilter,omitempty"`
	GroupsIdAttribute *string `json:"groupsIdAttribute,omitempty"`
	GroupsDisplayNameAttribute *string `json:"groupsDisplayNameAttribute,omitempty"`
	GroupsMemberAttribute *string `json:"groupsMemberAttribute,omitempty"`
}

// NewLdapGroupsQueryDtoImpl instantiates a new LdapGroupsQueryDtoImpl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapGroupsQueryDtoImpl() *LdapGroupsQueryDtoImpl {
	this := LdapGroupsQueryDtoImpl{}
	return &this
}

// NewLdapGroupsQueryDtoImplWithDefaults instantiates a new LdapGroupsQueryDtoImpl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapGroupsQueryDtoImplWithDefaults() *LdapGroupsQueryDtoImpl {
	this := LdapGroupsQueryDtoImpl{}
	return &this
}

// GetGroupsDn returns the GroupsDn field value if set, zero value otherwise.
func (o *LdapGroupsQueryDtoImpl) GetGroupsDn() string {
	if o == nil || o.GroupsDn == nil {
		var ret string
		return ret
	}
	return *o.GroupsDn
}

// GetGroupsDnOk returns a tuple with the GroupsDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapGroupsQueryDtoImpl) GetGroupsDnOk() (*string, bool) {
	if o == nil || o.GroupsDn == nil {
		return nil, false
	}
	return o.GroupsDn, true
}

// HasGroupsDn returns a boolean if a field has been set.
func (o *LdapGroupsQueryDtoImpl) HasGroupsDn() bool {
	if o != nil && o.GroupsDn != nil {
		return true
	}

	return false
}

// SetGroupsDn gets a reference to the given string and assigns it to the GroupsDn field.
func (o *LdapGroupsQueryDtoImpl) SetGroupsDn(v string) {
	o.GroupsDn = &v
}

// GetGroupsFilter returns the GroupsFilter field value if set, zero value otherwise.
func (o *LdapGroupsQueryDtoImpl) GetGroupsFilter() string {
	if o == nil || o.GroupsFilter == nil {
		var ret string
		return ret
	}
	return *o.GroupsFilter
}

// GetGroupsFilterOk returns a tuple with the GroupsFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapGroupsQueryDtoImpl) GetGroupsFilterOk() (*string, bool) {
	if o == nil || o.GroupsFilter == nil {
		return nil, false
	}
	return o.GroupsFilter, true
}

// HasGroupsFilter returns a boolean if a field has been set.
func (o *LdapGroupsQueryDtoImpl) HasGroupsFilter() bool {
	if o != nil && o.GroupsFilter != nil {
		return true
	}

	return false
}

// SetGroupsFilter gets a reference to the given string and assigns it to the GroupsFilter field.
func (o *LdapGroupsQueryDtoImpl) SetGroupsFilter(v string) {
	o.GroupsFilter = &v
}

// GetGroupsIdAttribute returns the GroupsIdAttribute field value if set, zero value otherwise.
func (o *LdapGroupsQueryDtoImpl) GetGroupsIdAttribute() string {
	if o == nil || o.GroupsIdAttribute == nil {
		var ret string
		return ret
	}
	return *o.GroupsIdAttribute
}

// GetGroupsIdAttributeOk returns a tuple with the GroupsIdAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapGroupsQueryDtoImpl) GetGroupsIdAttributeOk() (*string, bool) {
	if o == nil || o.GroupsIdAttribute == nil {
		return nil, false
	}
	return o.GroupsIdAttribute, true
}

// HasGroupsIdAttribute returns a boolean if a field has been set.
func (o *LdapGroupsQueryDtoImpl) HasGroupsIdAttribute() bool {
	if o != nil && o.GroupsIdAttribute != nil {
		return true
	}

	return false
}

// SetGroupsIdAttribute gets a reference to the given string and assigns it to the GroupsIdAttribute field.
func (o *LdapGroupsQueryDtoImpl) SetGroupsIdAttribute(v string) {
	o.GroupsIdAttribute = &v
}

// GetGroupsDisplayNameAttribute returns the GroupsDisplayNameAttribute field value if set, zero value otherwise.
func (o *LdapGroupsQueryDtoImpl) GetGroupsDisplayNameAttribute() string {
	if o == nil || o.GroupsDisplayNameAttribute == nil {
		var ret string
		return ret
	}
	return *o.GroupsDisplayNameAttribute
}

// GetGroupsDisplayNameAttributeOk returns a tuple with the GroupsDisplayNameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapGroupsQueryDtoImpl) GetGroupsDisplayNameAttributeOk() (*string, bool) {
	if o == nil || o.GroupsDisplayNameAttribute == nil {
		return nil, false
	}
	return o.GroupsDisplayNameAttribute, true
}

// HasGroupsDisplayNameAttribute returns a boolean if a field has been set.
func (o *LdapGroupsQueryDtoImpl) HasGroupsDisplayNameAttribute() bool {
	if o != nil && o.GroupsDisplayNameAttribute != nil {
		return true
	}

	return false
}

// SetGroupsDisplayNameAttribute gets a reference to the given string and assigns it to the GroupsDisplayNameAttribute field.
func (o *LdapGroupsQueryDtoImpl) SetGroupsDisplayNameAttribute(v string) {
	o.GroupsDisplayNameAttribute = &v
}

// GetGroupsMemberAttribute returns the GroupsMemberAttribute field value if set, zero value otherwise.
func (o *LdapGroupsQueryDtoImpl) GetGroupsMemberAttribute() string {
	if o == nil || o.GroupsMemberAttribute == nil {
		var ret string
		return ret
	}
	return *o.GroupsMemberAttribute
}

// GetGroupsMemberAttributeOk returns a tuple with the GroupsMemberAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapGroupsQueryDtoImpl) GetGroupsMemberAttributeOk() (*string, bool) {
	if o == nil || o.GroupsMemberAttribute == nil {
		return nil, false
	}
	return o.GroupsMemberAttribute, true
}

// HasGroupsMemberAttribute returns a boolean if a field has been set.
func (o *LdapGroupsQueryDtoImpl) HasGroupsMemberAttribute() bool {
	if o != nil && o.GroupsMemberAttribute != nil {
		return true
	}

	return false
}

// SetGroupsMemberAttribute gets a reference to the given string and assigns it to the GroupsMemberAttribute field.
func (o *LdapGroupsQueryDtoImpl) SetGroupsMemberAttribute(v string) {
	o.GroupsMemberAttribute = &v
}

func (o LdapGroupsQueryDtoImpl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupsDn != nil {
		toSerialize["groupsDn"] = o.GroupsDn
	}
	if o.GroupsFilter != nil {
		toSerialize["groupsFilter"] = o.GroupsFilter
	}
	if o.GroupsIdAttribute != nil {
		toSerialize["groupsIdAttribute"] = o.GroupsIdAttribute
	}
	if o.GroupsDisplayNameAttribute != nil {
		toSerialize["groupsDisplayNameAttribute"] = o.GroupsDisplayNameAttribute
	}
	if o.GroupsMemberAttribute != nil {
		toSerialize["groupsMemberAttribute"] = o.GroupsMemberAttribute
	}
	return json.Marshal(toSerialize)
}

type NullableLdapGroupsQueryDtoImpl struct {
	value *LdapGroupsQueryDtoImpl
	isSet bool
}

func (v NullableLdapGroupsQueryDtoImpl) Get() *LdapGroupsQueryDtoImpl {
	return v.value
}

func (v *NullableLdapGroupsQueryDtoImpl) Set(val *LdapGroupsQueryDtoImpl) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapGroupsQueryDtoImpl) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapGroupsQueryDtoImpl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapGroupsQueryDtoImpl(val *LdapGroupsQueryDtoImpl) *NullableLdapGroupsQueryDtoImpl {
	return &NullableLdapGroupsQueryDtoImpl{value: val, isSet: true}
}

func (v NullableLdapGroupsQueryDtoImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapGroupsQueryDtoImpl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


