# GlobalEventCaptureSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MouseUp** | **bool** | MouseUp enabled/disabled. | 
**MouseDown** | **bool** | MouseDown enabled/disabled. | 
**Click** | **bool** | Click enabled/disabled. | 
**DoubleClick** | **bool** | DoubleClick enabled/disabled. | 
**KeyUp** | **bool** | KeyUp enabled/disabled. | 
**KeyDown** | **bool** | KeyDown enabled/disabled. | 
**Scroll** | **bool** | Scroll enabled/disabled. | 
**AdditionalEventCapturedAsUserInput** | **string** | Additional events to be captured globally as user input.   For example, DragStart or DragEnd. | 

## Methods

### NewGlobalEventCaptureSettings

`func NewGlobalEventCaptureSettings(mouseUp bool, mouseDown bool, click bool, doubleClick bool, keyUp bool, keyDown bool, scroll bool, additionalEventCapturedAsUserInput string, ) *GlobalEventCaptureSettings`

NewGlobalEventCaptureSettings instantiates a new GlobalEventCaptureSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalEventCaptureSettingsWithDefaults

`func NewGlobalEventCaptureSettingsWithDefaults() *GlobalEventCaptureSettings`

NewGlobalEventCaptureSettingsWithDefaults instantiates a new GlobalEventCaptureSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMouseUp

`func (o *GlobalEventCaptureSettings) GetMouseUp() bool`

GetMouseUp returns the MouseUp field if non-nil, zero value otherwise.

### GetMouseUpOk

`func (o *GlobalEventCaptureSettings) GetMouseUpOk() (*bool, bool)`

GetMouseUpOk returns a tuple with the MouseUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMouseUp

`func (o *GlobalEventCaptureSettings) SetMouseUp(v bool)`

SetMouseUp sets MouseUp field to given value.


### GetMouseDown

`func (o *GlobalEventCaptureSettings) GetMouseDown() bool`

GetMouseDown returns the MouseDown field if non-nil, zero value otherwise.

### GetMouseDownOk

`func (o *GlobalEventCaptureSettings) GetMouseDownOk() (*bool, bool)`

GetMouseDownOk returns a tuple with the MouseDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMouseDown

`func (o *GlobalEventCaptureSettings) SetMouseDown(v bool)`

SetMouseDown sets MouseDown field to given value.


### GetClick

`func (o *GlobalEventCaptureSettings) GetClick() bool`

GetClick returns the Click field if non-nil, zero value otherwise.

### GetClickOk

`func (o *GlobalEventCaptureSettings) GetClickOk() (*bool, bool)`

GetClickOk returns a tuple with the Click field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClick

`func (o *GlobalEventCaptureSettings) SetClick(v bool)`

SetClick sets Click field to given value.


### GetDoubleClick

`func (o *GlobalEventCaptureSettings) GetDoubleClick() bool`

GetDoubleClick returns the DoubleClick field if non-nil, zero value otherwise.

### GetDoubleClickOk

`func (o *GlobalEventCaptureSettings) GetDoubleClickOk() (*bool, bool)`

GetDoubleClickOk returns a tuple with the DoubleClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleClick

`func (o *GlobalEventCaptureSettings) SetDoubleClick(v bool)`

SetDoubleClick sets DoubleClick field to given value.


### GetKeyUp

`func (o *GlobalEventCaptureSettings) GetKeyUp() bool`

GetKeyUp returns the KeyUp field if non-nil, zero value otherwise.

### GetKeyUpOk

`func (o *GlobalEventCaptureSettings) GetKeyUpOk() (*bool, bool)`

GetKeyUpOk returns a tuple with the KeyUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUp

`func (o *GlobalEventCaptureSettings) SetKeyUp(v bool)`

SetKeyUp sets KeyUp field to given value.


### GetKeyDown

`func (o *GlobalEventCaptureSettings) GetKeyDown() bool`

GetKeyDown returns the KeyDown field if non-nil, zero value otherwise.

### GetKeyDownOk

`func (o *GlobalEventCaptureSettings) GetKeyDownOk() (*bool, bool)`

GetKeyDownOk returns a tuple with the KeyDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyDown

`func (o *GlobalEventCaptureSettings) SetKeyDown(v bool)`

SetKeyDown sets KeyDown field to given value.


### GetScroll

`func (o *GlobalEventCaptureSettings) GetScroll() bool`

GetScroll returns the Scroll field if non-nil, zero value otherwise.

### GetScrollOk

`func (o *GlobalEventCaptureSettings) GetScrollOk() (*bool, bool)`

GetScrollOk returns a tuple with the Scroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScroll

`func (o *GlobalEventCaptureSettings) SetScroll(v bool)`

SetScroll sets Scroll field to given value.


### GetAdditionalEventCapturedAsUserInput

`func (o *GlobalEventCaptureSettings) GetAdditionalEventCapturedAsUserInput() string`

GetAdditionalEventCapturedAsUserInput returns the AdditionalEventCapturedAsUserInput field if non-nil, zero value otherwise.

### GetAdditionalEventCapturedAsUserInputOk

`func (o *GlobalEventCaptureSettings) GetAdditionalEventCapturedAsUserInputOk() (*string, bool)`

GetAdditionalEventCapturedAsUserInputOk returns a tuple with the AdditionalEventCapturedAsUserInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEventCapturedAsUserInput

`func (o *GlobalEventCaptureSettings) SetAdditionalEventCapturedAsUserInput(v string)`

SetAdditionalEventCapturedAsUserInput sets AdditionalEventCapturedAsUserInput field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


