/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// FullWebRequestRule The service detection rule of the `FULL_WEB_REQUEST` type.
type FullWebRequestRule struct {
	Type string `json:"type,omitempty"`
	Metadata ConfigurationMetadata `json:"metadata,omitempty"`
	// Specifies the management zones of the process group for which this service detection rule should be created.
	ManagementZones []string `json:"managementZones,omitempty"`
	// The ID of the service detection rule.
	Id string `json:"id,omitempty"`
	// The order of the rule in the rules list.    The rules are evaluated from top to bottom. The first matching rule applies.
	Order string `json:"order,omitempty"`
	// The name of the rule.
	Name string `json:"name"`
	// A short description of the rule.
	Description string `json:"description,omitempty"`
	// The rule is enabled(`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	// A list of conditions of the rule.   If several conditions are specified, the AND logic applies.
	Conditions []ConditionsFullWebRequestAttributeTypeDto `json:"conditions,omitempty"`
	ApplicationId ApplicationId `json:"applicationId,omitempty"`
	ContextRoot ContextRoot `json:"contextRoot,omitempty"`
	ServerName ServerName `json:"serverName,omitempty"`
}