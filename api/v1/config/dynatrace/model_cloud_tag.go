/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// CloudTag A cloud tag.
type CloudTag struct {
	// The name of the tag.
	Name string `json:"name,omitempty"`
	// The value of the tag.    If set to `null`, then resources with any value of the tag are monitored.
	Value string `json:"value,omitempty"`
}