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

// SlackNotificationConfigAllOf struct for SlackNotificationConfigAllOf
type SlackNotificationConfigAllOf struct {
	// The URL of the Slack WebHook.   This is confidential information, therefore GET requests return this field with the `null` value, and it is optional for PUT requests.
	Url *string `json:"url,omitempty"`
	// The channel (for example, `#general`) or the user (for example, `@john.smith`) to send the message to.
	Channel *string `json:"channel,omitempty"`
	// The content of the message.   You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.  
	Title *string `json:"title,omitempty"`
}

// NewSlackNotificationConfigAllOf instantiates a new SlackNotificationConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlackNotificationConfigAllOf() *SlackNotificationConfigAllOf {
	this := SlackNotificationConfigAllOf{}
	return &this
}

// NewSlackNotificationConfigAllOfWithDefaults instantiates a new SlackNotificationConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackNotificationConfigAllOfWithDefaults() *SlackNotificationConfigAllOf {
	this := SlackNotificationConfigAllOf{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SlackNotificationConfigAllOf) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackNotificationConfigAllOf) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SlackNotificationConfigAllOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SlackNotificationConfigAllOf) SetUrl(v string) {
	o.Url = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *SlackNotificationConfigAllOf) GetChannel() string {
	if o == nil || o.Channel == nil {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackNotificationConfigAllOf) GetChannelOk() (*string, bool) {
	if o == nil || o.Channel == nil {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *SlackNotificationConfigAllOf) HasChannel() bool {
	if o != nil && o.Channel != nil {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *SlackNotificationConfigAllOf) SetChannel(v string) {
	o.Channel = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SlackNotificationConfigAllOf) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackNotificationConfigAllOf) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SlackNotificationConfigAllOf) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SlackNotificationConfigAllOf) SetTitle(v string) {
	o.Title = &v
}

func (o SlackNotificationConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Channel != nil {
		toSerialize["channel"] = o.Channel
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableSlackNotificationConfigAllOf struct {
	value *SlackNotificationConfigAllOf
	isSet bool
}

func (v NullableSlackNotificationConfigAllOf) Get() *SlackNotificationConfigAllOf {
	return v.value
}

func (v *NullableSlackNotificationConfigAllOf) Set(val *SlackNotificationConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSlackNotificationConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSlackNotificationConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlackNotificationConfigAllOf(val *SlackNotificationConfigAllOf) *NullableSlackNotificationConfigAllOf {
	return &NullableSlackNotificationConfigAllOf{value: val, isSet: true}
}

func (v NullableSlackNotificationConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlackNotificationConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


