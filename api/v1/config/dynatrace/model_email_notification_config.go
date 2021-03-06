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

// EmailNotificationConfig Configuration of the email notification.
type EmailNotificationConfig struct {
	NotificationConfig
	// The subject of the email notifications.
	Subject string `json:"subject"`
	// The template of the email notification.   You can use the following placeholders:  * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.  * `{ProblemDetailsJSON}`: All problem event details, including root cause, as a JSON object.  * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.  
	Body string `json:"body"`
	// The list of the email recipients.
	Receivers []string `json:"receivers"`
	// The list of the email CC-recipients.
	CcReceivers []string `json:"ccReceivers"`
	// The list of the email BCC-recipients.
	BccReceivers []string `json:"bccReceivers"`
}

// NewEmailNotificationConfig instantiates a new EmailNotificationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailNotificationConfig(subject string, body string, receivers []string, ccReceivers []string, bccReceivers []string, ) *EmailNotificationConfig {
	this := EmailNotificationConfig{}
	this.Subject = subject
	this.Body = body
	this.Receivers = receivers
	this.CcReceivers = ccReceivers
	this.BccReceivers = bccReceivers
	return &this
}

// NewEmailNotificationConfigWithDefaults instantiates a new EmailNotificationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailNotificationConfigWithDefaults() *EmailNotificationConfig {
	this := EmailNotificationConfig{}
	return &this
}

// GetSubject returns the Subject field value
func (o *EmailNotificationConfig) GetSubject() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationConfig) GetSubjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *EmailNotificationConfig) SetSubject(v string) {
	o.Subject = v
}

// GetBody returns the Body field value
func (o *EmailNotificationConfig) GetBody() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationConfig) GetBodyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *EmailNotificationConfig) SetBody(v string) {
	o.Body = v
}

// GetReceivers returns the Receivers field value
func (o *EmailNotificationConfig) GetReceivers() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Receivers
}

// GetReceiversOk returns a tuple with the Receivers field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationConfig) GetReceiversOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Receivers, true
}

// SetReceivers sets field value
func (o *EmailNotificationConfig) SetReceivers(v []string) {
	o.Receivers = v
}

// GetCcReceivers returns the CcReceivers field value
func (o *EmailNotificationConfig) GetCcReceivers() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.CcReceivers
}

// GetCcReceiversOk returns a tuple with the CcReceivers field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationConfig) GetCcReceiversOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CcReceivers, true
}

// SetCcReceivers sets field value
func (o *EmailNotificationConfig) SetCcReceivers(v []string) {
	o.CcReceivers = v
}

// GetBccReceivers returns the BccReceivers field value
func (o *EmailNotificationConfig) GetBccReceivers() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.BccReceivers
}

// GetBccReceiversOk returns a tuple with the BccReceivers field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationConfig) GetBccReceiversOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BccReceivers, true
}

// SetBccReceivers sets field value
func (o *EmailNotificationConfig) SetBccReceivers(v []string) {
	o.BccReceivers = v
}

func (o EmailNotificationConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedNotificationConfig, errNotificationConfig := json.Marshal(o.NotificationConfig)
	if errNotificationConfig != nil {
		return []byte{}, errNotificationConfig
	}
	errNotificationConfig = json.Unmarshal([]byte(serializedNotificationConfig), &toSerialize)
	if errNotificationConfig != nil {
		return []byte{}, errNotificationConfig
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	if true {
		toSerialize["body"] = o.Body
	}
	if true {
		toSerialize["receivers"] = o.Receivers
	}
	if true {
		toSerialize["ccReceivers"] = o.CcReceivers
	}
	if true {
		toSerialize["bccReceivers"] = o.BccReceivers
	}
	return json.Marshal(toSerialize)
}

type NullableEmailNotificationConfig struct {
	value *EmailNotificationConfig
	isSet bool
}

func (v NullableEmailNotificationConfig) Get() *EmailNotificationConfig {
	return v.value
}

func (v *NullableEmailNotificationConfig) Set(val *EmailNotificationConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailNotificationConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailNotificationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailNotificationConfig(val *EmailNotificationConfig) *NullableEmailNotificationConfig {
	return &NullableEmailNotificationConfig{value: val, isSet: true}
}

func (v NullableEmailNotificationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailNotificationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


