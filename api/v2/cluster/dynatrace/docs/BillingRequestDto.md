# BillingRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterUuid** | Pointer to **string** |  | [optional] 
**TimeFrameStart** | Pointer to **time.Time** |  | [optional] 
**TimeFrameEnd** | Pointer to **time.Time** |  | [optional] 
**EnvironmentBillingEntries** | Pointer to [**[]EnvironmentUsageDto**](EnvironmentUsageDto.md) |  | [optional] 

## Methods

### NewBillingRequestDto

`func NewBillingRequestDto() *BillingRequestDto`

NewBillingRequestDto instantiates a new BillingRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingRequestDtoWithDefaults

`func NewBillingRequestDtoWithDefaults() *BillingRequestDto`

NewBillingRequestDtoWithDefaults instantiates a new BillingRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterUuid

`func (o *BillingRequestDto) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *BillingRequestDto) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *BillingRequestDto) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *BillingRequestDto) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetTimeFrameStart

`func (o *BillingRequestDto) GetTimeFrameStart() time.Time`

GetTimeFrameStart returns the TimeFrameStart field if non-nil, zero value otherwise.

### GetTimeFrameStartOk

`func (o *BillingRequestDto) GetTimeFrameStartOk() (*time.Time, bool)`

GetTimeFrameStartOk returns a tuple with the TimeFrameStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrameStart

`func (o *BillingRequestDto) SetTimeFrameStart(v time.Time)`

SetTimeFrameStart sets TimeFrameStart field to given value.

### HasTimeFrameStart

`func (o *BillingRequestDto) HasTimeFrameStart() bool`

HasTimeFrameStart returns a boolean if a field has been set.

### GetTimeFrameEnd

`func (o *BillingRequestDto) GetTimeFrameEnd() time.Time`

GetTimeFrameEnd returns the TimeFrameEnd field if non-nil, zero value otherwise.

### GetTimeFrameEndOk

`func (o *BillingRequestDto) GetTimeFrameEndOk() (*time.Time, bool)`

GetTimeFrameEndOk returns a tuple with the TimeFrameEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrameEnd

`func (o *BillingRequestDto) SetTimeFrameEnd(v time.Time)`

SetTimeFrameEnd sets TimeFrameEnd field to given value.

### HasTimeFrameEnd

`func (o *BillingRequestDto) HasTimeFrameEnd() bool`

HasTimeFrameEnd returns a boolean if a field has been set.

### GetEnvironmentBillingEntries

`func (o *BillingRequestDto) GetEnvironmentBillingEntries() []EnvironmentUsageDto`

GetEnvironmentBillingEntries returns the EnvironmentBillingEntries field if non-nil, zero value otherwise.

### GetEnvironmentBillingEntriesOk

`func (o *BillingRequestDto) GetEnvironmentBillingEntriesOk() (*[]EnvironmentUsageDto, bool)`

GetEnvironmentBillingEntriesOk returns a tuple with the EnvironmentBillingEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentBillingEntries

`func (o *BillingRequestDto) SetEnvironmentBillingEntries(v []EnvironmentUsageDto)`

SetEnvironmentBillingEntries sets EnvironmentBillingEntries field to given value.

### HasEnvironmentBillingEntries

`func (o *BillingRequestDto) HasEnvironmentBillingEntries() bool`

HasEnvironmentBillingEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


