/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// OpsGenieNotificationConfigAllOf struct for OpsGenieNotificationConfigAllOf
type OpsGenieNotificationConfigAllOf struct {
	// The API key to access OpsGenie.
	ApiKey *string `json:"apiKey,omitempty"`
	// The region domain of the OpsGenie.
	Domain *string `json:"domain,omitempty"`
	// The content of the message.   You can use the following placeholders:  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  
	Message *string `json:"message,omitempty"`
}

// NewOpsGenieNotificationConfigAllOf instantiates a new OpsGenieNotificationConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsGenieNotificationConfigAllOf() *OpsGenieNotificationConfigAllOf {
	this := OpsGenieNotificationConfigAllOf{}
	return &this
}

// NewOpsGenieNotificationConfigAllOfWithDefaults instantiates a new OpsGenieNotificationConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsGenieNotificationConfigAllOfWithDefaults() *OpsGenieNotificationConfigAllOf {
	this := OpsGenieNotificationConfigAllOf{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *OpsGenieNotificationConfigAllOf) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsGenieNotificationConfigAllOf) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *OpsGenieNotificationConfigAllOf) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *OpsGenieNotificationConfigAllOf) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *OpsGenieNotificationConfigAllOf) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsGenieNotificationConfigAllOf) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *OpsGenieNotificationConfigAllOf) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *OpsGenieNotificationConfigAllOf) SetDomain(v string) {
	o.Domain = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OpsGenieNotificationConfigAllOf) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsGenieNotificationConfigAllOf) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OpsGenieNotificationConfigAllOf) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OpsGenieNotificationConfigAllOf) SetMessage(v string) {
	o.Message = &v
}

func (o OpsGenieNotificationConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiKey != nil {
		toSerialize["apiKey"] = o.ApiKey
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableOpsGenieNotificationConfigAllOf struct {
	value *OpsGenieNotificationConfigAllOf
	isSet bool
}

func (v NullableOpsGenieNotificationConfigAllOf) Get() *OpsGenieNotificationConfigAllOf {
	return v.value
}

func (v *NullableOpsGenieNotificationConfigAllOf) Set(val *OpsGenieNotificationConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsGenieNotificationConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsGenieNotificationConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsGenieNotificationConfigAllOf(val *OpsGenieNotificationConfigAllOf) *NullableOpsGenieNotificationConfigAllOf {
	return &NullableOpsGenieNotificationConfigAllOf{value: val, isSet: true}
}

func (v NullableOpsGenieNotificationConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsGenieNotificationConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


