# EntityBaselineData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The ID of the Dynatrace entity. | 
**DisplayName** | Pointer to **string** | The display name of the entity. | [optional] 
**ErrorRate** | Pointer to **float32** | The error rate baseline. | [optional] 
**HasLoadBaseline** | Pointer to **bool** | The entity has a load baseline (&#x60;true&#x60;) or doesn&#39;t (&#x60;false&#x60;). | [optional] 
**MicrosMedian** | Pointer to **float32** | The median baseline, in microseconds. | [optional] 
**Micros90thPercentile** | Pointer to **float32** | The 90th percentile baseline, in microseconds. | [optional] 
**ChildBaselines** | Pointer to [**[]EntityBaselineData**](EntityBaselineData.md) | The baseline data for child entities of this entity, for example a &#x60;SERVICE_METHOD&#x60; of a &#x60;SERVICE_METHOD_GROUP&#x60;. | [optional] 

## Methods

### NewEntityBaselineData

`func NewEntityBaselineData(entityId string, ) *EntityBaselineData`

NewEntityBaselineData instantiates a new EntityBaselineData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityBaselineDataWithDefaults

`func NewEntityBaselineDataWithDefaults() *EntityBaselineData`

NewEntityBaselineDataWithDefaults instantiates a new EntityBaselineData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *EntityBaselineData) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *EntityBaselineData) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *EntityBaselineData) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetDisplayName

`func (o *EntityBaselineData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EntityBaselineData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EntityBaselineData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EntityBaselineData) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetErrorRate

`func (o *EntityBaselineData) GetErrorRate() float32`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *EntityBaselineData) GetErrorRateOk() (*float32, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *EntityBaselineData) SetErrorRate(v float32)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *EntityBaselineData) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetHasLoadBaseline

`func (o *EntityBaselineData) GetHasLoadBaseline() bool`

GetHasLoadBaseline returns the HasLoadBaseline field if non-nil, zero value otherwise.

### GetHasLoadBaselineOk

`func (o *EntityBaselineData) GetHasLoadBaselineOk() (*bool, bool)`

GetHasLoadBaselineOk returns a tuple with the HasLoadBaseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLoadBaseline

`func (o *EntityBaselineData) SetHasLoadBaseline(v bool)`

SetHasLoadBaseline sets HasLoadBaseline field to given value.

### HasHasLoadBaseline

`func (o *EntityBaselineData) HasHasLoadBaseline() bool`

HasHasLoadBaseline returns a boolean if a field has been set.

### GetMicrosMedian

`func (o *EntityBaselineData) GetMicrosMedian() float32`

GetMicrosMedian returns the MicrosMedian field if non-nil, zero value otherwise.

### GetMicrosMedianOk

`func (o *EntityBaselineData) GetMicrosMedianOk() (*float32, bool)`

GetMicrosMedianOk returns a tuple with the MicrosMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosMedian

`func (o *EntityBaselineData) SetMicrosMedian(v float32)`

SetMicrosMedian sets MicrosMedian field to given value.

### HasMicrosMedian

`func (o *EntityBaselineData) HasMicrosMedian() bool`

HasMicrosMedian returns a boolean if a field has been set.

### GetMicros90thPercentile

`func (o *EntityBaselineData) GetMicros90thPercentile() float32`

GetMicros90thPercentile returns the Micros90thPercentile field if non-nil, zero value otherwise.

### GetMicros90thPercentileOk

`func (o *EntityBaselineData) GetMicros90thPercentileOk() (*float32, bool)`

GetMicros90thPercentileOk returns a tuple with the Micros90thPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicros90thPercentile

`func (o *EntityBaselineData) SetMicros90thPercentile(v float32)`

SetMicros90thPercentile sets Micros90thPercentile field to given value.

### HasMicros90thPercentile

`func (o *EntityBaselineData) HasMicros90thPercentile() bool`

HasMicros90thPercentile returns a boolean if a field has been set.

### GetChildBaselines

`func (o *EntityBaselineData) GetChildBaselines() []EntityBaselineData`

GetChildBaselines returns the ChildBaselines field if non-nil, zero value otherwise.

### GetChildBaselinesOk

`func (o *EntityBaselineData) GetChildBaselinesOk() (*[]EntityBaselineData, bool)`

GetChildBaselinesOk returns a tuple with the ChildBaselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildBaselines

`func (o *EntityBaselineData) SetChildBaselines(v []EntityBaselineData)`

SetChildBaselines sets ChildBaselines field to given value.

### HasChildBaselines

`func (o *EntityBaselineData) HasChildBaselines() bool`

HasChildBaselines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


