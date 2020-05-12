/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// ExtensionState The state of the extension.
type ExtensionState struct {
	// The ID of the extension.
	ExtensionId string `json:"extensionId,omitempty"`
	// The version of the extension (for example `1.0.0`).
	Version string `json:"version,omitempty"`
	// The ID of the endpoint where the state is detected - Active Gate only.
	EndpointId string `json:"endpointId,omitempty"`
	// The state of the extension.
	State string `json:"state,omitempty"`
	// A short description of the state.
	StateDescription string `json:"stateDescription,omitempty"`
	// The timestamp when the state was detected, in UTC milliseconds.
	Timestamp int64 `json:"timestamp,omitempty"`
	// The ID of the host on which the extension runs.
	HostId string `json:"hostId,omitempty"`
	// The ID of the entity on which the extension is active.
	ProcessId string `json:"processId,omitempty"`
}