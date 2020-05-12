/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// MaintenanceWindow Configuration of a maintenance window.
type MaintenanceWindow struct {
	Metadata ConfigurationMetadata `json:"metadata,omitempty"`
	// The ID of the maintenance window.
	Id string `json:"id,omitempty"`
	// The name of the maintenance window, displayed in the UI.
	Name string `json:"name"`
	// A short description of the maintenance purpose.
	Description string `json:"description"`
	// The type of the maintenance: planned or unplanned.
	Type string `json:"type"`
	// The type of suppression of alerting and problem detection during the maintenance.
	Suppression string `json:"suppression"`
	Scope Scope `json:"scope,omitempty"`
	Schedule Schedule `json:"schedule"`
}