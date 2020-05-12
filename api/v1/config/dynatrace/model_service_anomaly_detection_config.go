/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// ServiceAnomalyDetectionConfig Dynatrace automatically detects service-related performance anomalies such as response time degradations and failure rate increases. Use these settings to configure detection sensitivity, set alert thresholds, or disable alerting for certain services.
type ServiceAnomalyDetectionConfig struct {
	Metadata ConfigurationMetadata `json:"metadata,omitempty"`
	ResponseTimeDegradation ResponseTimeDegradationDetectionConfig `json:"responseTimeDegradation"`
	LoadDrop LoadDropDetectionConfig `json:"loadDrop,omitempty"`
	LoadSpike LoadSpikeDetectionConfig `json:"loadSpike,omitempty"`
	FailureRateIncrease FailureRateIncreaseDetectionConfig `json:"failureRateIncrease"`
}