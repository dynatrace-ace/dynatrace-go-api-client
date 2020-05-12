/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// ApplicationDetectionRuleConfig Application detection rule.
type ApplicationDetectionRuleConfig struct {
	Metadata ConfigurationMetadataDtoImpl `json:"metadata,omitempty"`
	// The ID of the rule.
	Id string `json:"id,omitempty"`
	// The order of the rule in the rules list.   The rules are evaluated from top to bottom. The first matching rule applies.
	Order string `json:"order,omitempty"`
	// The Dynatrace entity ID of the application, for example `APPLICATION-4A3B43`.    You must use an existing ID. If you need to create a rule for an application that doesn't exist yet, [create an application first](https://www.dynatrace.com/support/help/shortlink/api-config-web-app-post-web-app) and then configure detection rules for it.
	ApplicationIdentifier string `json:"applicationIdentifier"`
	FilterConfig ApplicationFilter `json:"filterConfig"`
}