# SyntheticTestLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the location. | 
**Enabled** | Pointer to **bool** | The location is enabled/disabled. Default is &#x60;true&#x60;, enabling the location. | [optional] 

## Methods

### NewSyntheticTestLocation

`func NewSyntheticTestLocation(id string, ) *SyntheticTestLocation`

NewSyntheticTestLocation instantiates a new SyntheticTestLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticTestLocationWithDefaults

`func NewSyntheticTestLocationWithDefaults() *SyntheticTestLocation`

NewSyntheticTestLocationWithDefaults instantiates a new SyntheticTestLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SyntheticTestLocation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyntheticTestLocation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyntheticTestLocation) SetId(v string)`

SetId sets Id field to given value.


### GetEnabled

`func (o *SyntheticTestLocation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SyntheticTestLocation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SyntheticTestLocation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SyntheticTestLocation) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


