/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// TimeoutSettings Settings for timed action capture.
type TimeoutSettings struct {
	// Timed action support enabled/disabled.   Enable to detect actions that trigger sending of XHRs via *setTimout* methods.
	TimedActionSupport bool `json:"timedActionSupport"`
	// Defines how deep temporary actions may cascade. 0 disables temporary actions completely. Recommended value if enabled is 3.
	TemporaryActionLimit int32 `json:"temporaryActionLimit"`
	// The total timeout of all cascaded timeouts that should still be able to create a temporary action
	TemporaryActionTotalTimeout int32 `json:"temporaryActionTotalTimeout"`
}