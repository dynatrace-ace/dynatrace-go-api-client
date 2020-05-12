/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// RequestAttribute struct for RequestAttribute
type RequestAttribute struct {
	Metadata ConfigurationMetadata `json:"metadata,omitempty"`
	// The ID of the request attribute.
	Id string `json:"id,omitempty"`
	// The name of the request attribute.
	Name string `json:"name"`
	// The request attribute is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	// The data type of the request attribute.
	DataType string `json:"dataType"`
	// The list of data sources.
	DataSources []DataSource `json:"dataSources"`
	// String values transformation.    If the **dataType** is not `string`, set the `Original` here.
	Normalization string `json:"normalization"`
	// Aggregation type for the request values.
	Aggregation string `json:"aggregation"`
	// Confidential data flag. Set `true` to treat the captured data as confidential.
	Confidential bool `json:"confidential"`
	// Personal data masking flag. Set `true` to skip masking.    Warning: This will potentially access personalized data.
	SkipPersonalDataMasking bool `json:"skipPersonalDataMasking"`
}