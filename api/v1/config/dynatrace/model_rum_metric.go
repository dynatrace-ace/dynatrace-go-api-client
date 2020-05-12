/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// RumMetric struct for RumMetric
type RumMetric struct {
	// The application identifier of the application where this RumMetric should belong to
	ApplicationIdentifier string `json:"applicationIdentifier"`
	// The name of the RumMetric config
	Name string `json:"name"`
	// The unique key of this RumMetric. Has to start with \"calc:apps.\"
	MetricKey string `json:"metricKey"`
	// RumMetric enabled/disabled
	Enabled bool `json:"enabled"`
	MetricDefinition RumMetricDefinition `json:"metricDefinition"`
	// Specifies the optional dimensions for this RumMetric.
	Dimensions []RumDimensionDefinition `json:"dimensions,omitempty"`
	UserActionFilter UserActionFilter `json:"userActionFilter,omitempty"`
}