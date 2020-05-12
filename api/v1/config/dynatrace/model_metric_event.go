/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// MetricEvent The configuration of the metric event.
type MetricEvent struct {
	Metadata ConfigurationMetadata `json:"metadata,omitempty"`
	// The ID of the metric event.
	Id string `json:"id,omitempty"`
	// The ID of the metric evaluated by the metric event.
	MetricId string `json:"metricId"`
	// The name of the metric event displayed in the UI.
	Name string `json:"name"`
	// The description of the metric event.
	Description string `json:"description"`
	// How the metric data points are aggregated for the evaluation.    The timeseries must support this aggregation.
	AggregationType string `json:"aggregationType,omitempty"`
	// The type of the event to trigger on the threshold violation.   The `CUSTOM_ALERT` type is not correlated with other alerts. The `INFO` type does not open a problem.
	Severity string `json:"severity,omitempty"`
	// The condition for the **threshold** value check: above or below.
	AlertCondition string `json:"alertCondition"`
	// The number of one-minute samples that form sliding evaluation window.
	Samples int32 `json:"samples"`
	// The number of one-minute samples within the evaluation window that must violate the threshold to trigger an event
	ViolatingSamples int32 `json:"violatingSamples"`
	// The number of one-minute samples within the evaluation window that must go back to normal to close the event.
	DealertingSamples int32 `json:"dealertingSamples"`
	// The value of the threshold.
	Threshold float64 `json:"threshold"`
	// The metric event is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	// Defines the scope of the metric event. Only one filter is allowed per filter type, except for tags, where up to 3 are allowed. The filters are combined by conjunction.
	AlertingScope []MetricEventAlertingScope `json:"alertingScope,omitempty"`
	// Defines the dimensions of the metric to alert on. The filters are combined by conjunction.
	MetricDimensions []MetricEventDimensions `json:"metricDimensions,omitempty"`
	// The unit of the threshold, matching the metric definition.  If not set the metrics unit is picked.
	Unit string `json:"unit,omitempty"`
}