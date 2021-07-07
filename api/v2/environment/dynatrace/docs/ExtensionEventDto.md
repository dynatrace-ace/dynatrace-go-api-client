# ExtensionEventDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** | Timestamp of the latest event | [optional] 
**Severity** | Pointer to **string** | Severity of the latest event | [optional] 
**Content** | Pointer to **string** | Content | [optional] 

## Methods

### NewExtensionEventDto

`func NewExtensionEventDto() *ExtensionEventDto`

NewExtensionEventDto instantiates a new ExtensionEventDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionEventDtoWithDefaults

`func NewExtensionEventDtoWithDefaults() *ExtensionEventDto`

NewExtensionEventDtoWithDefaults instantiates a new ExtensionEventDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ExtensionEventDto) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExtensionEventDto) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExtensionEventDto) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ExtensionEventDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetSeverity

`func (o *ExtensionEventDto) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ExtensionEventDto) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ExtensionEventDto) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ExtensionEventDto) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetContent

`func (o *ExtensionEventDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ExtensionEventDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ExtensionEventDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ExtensionEventDto) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


