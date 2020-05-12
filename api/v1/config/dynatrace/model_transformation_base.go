/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// TransformationBase Configuration of transformation of the detected value.   If several transformations are specified, they are handled sequentially from top to bottom. Each transformation is applied to the result of the preceding transformation. For example, the second transformation is applied to the result of the first transformation.   The actual set of fields depends on the `type` of the transformation.
type TransformationBase struct {
	// Defines the actual set of fields depending on the value. See one of the following objects:   * `BEFORE` -> BeforeTransformation  * `AFTER` -> AfterTransformation  * `BETWEEN` -> BetweenTransformation  * `REPLACE_BETWEEN` -> ReplaceBetweenTransformation  * `REMOVE_NUMBERS` -> RemoveNumbersTransformation  * `REMOVE_CREDIT_CARDS` -> RemoveCreditCardNumbersTransformation  * `REMOVE_IBANS` -> RemoveIBANsTransformation  * `REMOVE_IPS` -> RemoveIPsTransformation  * `SPLIT_SELECT` -> SplitSelectTransformation  * `TAKE_SEGMENTS` -> TakeSegmentsTransformation  
	Type string `json:"type"`
}