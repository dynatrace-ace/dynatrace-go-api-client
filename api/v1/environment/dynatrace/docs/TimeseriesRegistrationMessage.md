# TimeseriesRegistrationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the metric that will appear in the user interface. It is limited to 256 characters. | [optional] 
**Unit** | Pointer to **string** | The unit the metric will use. If the parameter is not specified or the wrong value is specified, the &#x60;Count&#x60; value will be assigned. | [optional] 
**Dimensions** | Pointer to **[]string** | The metric dimension key that will be used to report multiple dimensions. For example, a dimension key to report the metric for different network cards for the same firewall.   You can use alphanumeric characters and the following punctuation marks: periods (&#x60;.&#x60;), hyphens (&#x60;-&#x60;), and underscores (&#x60;_&#x60;).   The CUSTOM_DEVICE dimension is added to each new custom metric automatically. | [optional] 
**Types** | Pointer to **[]string** | The definition of the technology type. It is used to group metrics under a logical technology name in the UI.   Metrics must be assigned a software technology type that is identical to the technology type of the custom device you are sending the metric to.   For example, if you define your custom device using type &#x60;F5-Firewall&#x60; you must also register all related custom metrics as type &#x60;F5-Firewall&#x60;. | [optional] 

## Methods

### NewTimeseriesRegistrationMessage

`func NewTimeseriesRegistrationMessage() *TimeseriesRegistrationMessage`

NewTimeseriesRegistrationMessage instantiates a new TimeseriesRegistrationMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesRegistrationMessageWithDefaults

`func NewTimeseriesRegistrationMessageWithDefaults() *TimeseriesRegistrationMessage`

NewTimeseriesRegistrationMessageWithDefaults instantiates a new TimeseriesRegistrationMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *TimeseriesRegistrationMessage) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TimeseriesRegistrationMessage) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TimeseriesRegistrationMessage) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TimeseriesRegistrationMessage) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUnit

`func (o *TimeseriesRegistrationMessage) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TimeseriesRegistrationMessage) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TimeseriesRegistrationMessage) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *TimeseriesRegistrationMessage) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetDimensions

`func (o *TimeseriesRegistrationMessage) GetDimensions() []string`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *TimeseriesRegistrationMessage) GetDimensionsOk() (*[]string, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *TimeseriesRegistrationMessage) SetDimensions(v []string)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *TimeseriesRegistrationMessage) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetTypes

`func (o *TimeseriesRegistrationMessage) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *TimeseriesRegistrationMessage) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *TimeseriesRegistrationMessage) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *TimeseriesRegistrationMessage) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


