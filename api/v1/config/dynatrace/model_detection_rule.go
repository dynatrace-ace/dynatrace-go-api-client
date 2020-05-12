/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// DetectionRule struct for DetectionRule
type DetectionRule struct {
	// The ID of the detection rule.
	Id string `json:"id,omitempty"`
	// Rule enabled/disabled.
	Enabled bool `json:"enabled"`
	// The PHP file containing the class or methods to instrument.   Required for PHP custom service.    Not applicable to Java and .NET.
	FileName string `json:"fileName,omitempty"`
	// Matcher applying to the file name. Default value is `ENDS_WITH` (if applicable).
	FileNameMatcher string `json:"fileNameMatcher,omitempty"`
	// The fully qualified class or interface to instrument.   Required for Java and .NET custom services.    Not applicable to PHP.
	ClassName string `json:"className,omitempty"`
	// Matcher applying to the class name. `STARTS_WITH` can only be used if there is at least one annotation defined. Default value is `EQUALS`.
	Matcher string `json:"matcher,omitempty"`
	// List of methods to instrument.
	MethodRules []MethodRule `json:"methodRules"`
	// Additional annotations filter of the rule.   Only classes where all listed annotations are available in the class itself or any of its superclasses are instrumented.   Not applicable to PHP.
	Annotations []string `json:"annotations,omitempty"`
}