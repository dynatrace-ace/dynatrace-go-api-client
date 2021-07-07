# ExtensionStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **int64** | Timestamp of the latest status of given configuration. | [optional] 
**Status** | Pointer to **string** | Latest status of given configuration. | [optional] 

## Methods

### NewExtensionStatusDto

`func NewExtensionStatusDto() *ExtensionStatusDto`

NewExtensionStatusDto instantiates a new ExtensionStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionStatusDtoWithDefaults

`func NewExtensionStatusDtoWithDefaults() *ExtensionStatusDto`

NewExtensionStatusDtoWithDefaults instantiates a new ExtensionStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ExtensionStatusDto) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExtensionStatusDto) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExtensionStatusDto) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ExtensionStatusDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetStatus

`func (o *ExtensionStatusDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExtensionStatusDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExtensionStatusDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExtensionStatusDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


