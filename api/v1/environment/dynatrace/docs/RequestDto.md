# RequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | Request identifier | 
**Name** | **string** | Request name | 
**SequenceNumber** | **int32** | Request sequence number | 

## Methods

### NewRequestDto

`func NewRequestDto(entityId string, name string, sequenceNumber int32, ) *RequestDto`

NewRequestDto instantiates a new RequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestDtoWithDefaults

`func NewRequestDtoWithDefaults() *RequestDto`

NewRequestDtoWithDefaults instantiates a new RequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *RequestDto) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *RequestDto) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *RequestDto) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetName

`func (o *RequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestDto) SetName(v string)`

SetName sets Name field to given value.


### GetSequenceNumber

`func (o *RequestDto) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *RequestDto) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *RequestDto) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


