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

// JiraNotificationConfig Configuration of the Jira notification.
type JiraNotificationConfig struct {
	NotificationConfig
	// The username of the Jira profile.
	Username string `json:"username"`
	// The password for the Jira profile.
	Password *string `json:"password,omitempty"`
	// The URL of the Jira API endpoint.
	Url string `json:"url"`
	// The project key of the Jira issue to be created by this notification.
	ProjectKey string `json:"projectKey"`
	// The type of the Jira issue to be created by this notification.
	IssueType string `json:"issueType"`
	// The summary of the Jira issue to be created by this notification.   You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.  
	Summary string `json:"summary"`
	// The description of the Jira issue to be created by this notification.    You can use same placeholders as in issue summary.
	Description string `json:"description"`
}

// NewJiraNotificationConfig instantiates a new JiraNotificationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJiraNotificationConfig(username string, url string, projectKey string, issueType string, summary string, description string, ) *JiraNotificationConfig {
	this := JiraNotificationConfig{}
	this.Username = username
	this.Url = url
	this.ProjectKey = projectKey
	this.IssueType = issueType
	this.Summary = summary
	this.Description = description
	return &this
}

// NewJiraNotificationConfigWithDefaults instantiates a new JiraNotificationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJiraNotificationConfigWithDefaults() *JiraNotificationConfig {
	this := JiraNotificationConfig{}
	return &this
}

// GetUsername returns the Username field value
func (o *JiraNotificationConfig) GetUsername() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *JiraNotificationConfig) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *JiraNotificationConfig) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *JiraNotificationConfig) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraNotificationConfig) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *JiraNotificationConfig) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *JiraNotificationConfig) SetPassword(v string) {
	o.Password = &v
}

// GetUrl returns the Url field value
func (o *JiraNotificationConfig) GetUrl() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *JiraNotificationConfig) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *JiraNotificationConfig) SetUrl(v string) {
	o.Url = v
}

// GetProjectKey returns the ProjectKey field value
func (o *JiraNotificationConfig) GetProjectKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ProjectKey
}

// GetProjectKeyOk returns a tuple with the ProjectKey field value
// and a boolean to check if the value has been set.
func (o *JiraNotificationConfig) GetProjectKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProjectKey, true
}

// SetProjectKey sets field value
func (o *JiraNotificationConfig) SetProjectKey(v string) {
	o.ProjectKey = v
}

// GetIssueType returns the IssueType field value
func (o *JiraNotificationConfig) GetIssueType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.IssueType
}

// GetIssueTypeOk returns a tuple with the IssueType field value
// and a boolean to check if the value has been set.
func (o *JiraNotificationConfig) GetIssueTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IssueType, true
}

// SetIssueType sets field value
func (o *JiraNotificationConfig) SetIssueType(v string) {
	o.IssueType = v
}

// GetSummary returns the Summary field value
func (o *JiraNotificationConfig) GetSummary() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *JiraNotificationConfig) GetSummaryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *JiraNotificationConfig) SetSummary(v string) {
	o.Summary = v
}

// GetDescription returns the Description field value
func (o *JiraNotificationConfig) GetDescription() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *JiraNotificationConfig) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *JiraNotificationConfig) SetDescription(v string) {
	o.Description = v
}

func (o JiraNotificationConfig) MarshalJSON() ([]byte, error) {
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
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["projectKey"] = o.ProjectKey
	}
	if true {
		toSerialize["issueType"] = o.IssueType
	}
	if true {
		toSerialize["summary"] = o.Summary
	}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableJiraNotificationConfig struct {
	value *JiraNotificationConfig
	isSet bool
}

func (v NullableJiraNotificationConfig) Get() *JiraNotificationConfig {
	return v.value
}

func (v *NullableJiraNotificationConfig) Set(val *JiraNotificationConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableJiraNotificationConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableJiraNotificationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJiraNotificationConfig(val *JiraNotificationConfig) *NullableJiraNotificationConfig {
	return &NullableJiraNotificationConfig{value: val, isSet: true}
}

func (v NullableJiraNotificationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJiraNotificationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


