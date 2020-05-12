/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// ExistsCompareOperation The condition of the `EXISTS` type.   The condition checks whether the specified attribute exists.
type ExistsCompareOperation struct {
	// Defines the actual set of fields depending on the value. See one of the following objects:   * `EQUALS` -> EqualsCompareOperation  * `STRING_CONTAINS` -> StringContainsCompareOperation  * `STARTS_WITH` -> StartsWithCompareOperation  * `ENDS_WITH` -> EndsWithCompareOperation  * `EXISTS` -> ExistsCompareOperation  * `IP_IN_RANGE` -> IpInRangeCompareOperation  * `LESS_THAN` -> LessThanCompareOperation  * `GREATER_THAN` -> GreaterThanCompareOperation  * `INT_EQUALS` -> IntEqualsCompareOperation  * `STRING_EQUALS` -> StringEqualsCompareOperation  * `TAG` -> TagCompareOperation  
	Type string `json:"type"`
	// Inverts the operation of the condition. Set to `true` to turn **exists** into **does not exist**.    If not set, then `false` is used.
	Negate bool `json:"negate,omitempty"`
}