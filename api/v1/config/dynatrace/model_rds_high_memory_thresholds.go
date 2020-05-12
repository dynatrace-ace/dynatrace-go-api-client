/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// RdsHighMemoryThresholds Custom thresholds for RDS running out of memory. If not set, automatic mode is used.    **All** conditions must be fulfilled to trigger an alert.
type RdsHighMemoryThresholds struct {
	// Freeable memory is lower than *X* Megabytes in 3 out of 5 samples.
	FreeMemory float32 `json:"freeMemory"`
	// Swap usage is higher than *X* Gigabytes in 3 out of 5 samples.
	SwapUsage float32 `json:"swapUsage"`
}