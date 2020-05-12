/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// GlobalEventCaptureSettings Global event capture settings.
type GlobalEventCaptureSettings struct {
	// MouseUp enabled/disabled.
	MouseUp bool `json:"mouseUp"`
	// MouseDown enabled/disabled.
	MouseDown bool `json:"mouseDown"`
	// Click enabled/disabled.
	Click bool `json:"click"`
	// DoubleClick enabled/disabled.
	DoubleClick bool `json:"doubleClick"`
	// KeyUp enabled/disabled.
	KeyUp bool `json:"keyUp"`
	// KeyDown enabled/disabled.
	KeyDown bool `json:"keyDown"`
	// Scroll enabled/disabled.
	Scroll bool `json:"scroll"`
	// Additional events to be captured globally as user input.   For example, DragStart or DragEnd.
	AdditionalEventCapturedAsUserInput string `json:"additionalEventCapturedAsUserInput"`
}