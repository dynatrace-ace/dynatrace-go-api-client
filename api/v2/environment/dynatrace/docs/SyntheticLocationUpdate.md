# SyntheticLocationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;PUBLIC&#x60; -&gt; SyntheticPublicLocationUpdate  * &#x60;PRIVATE&#x60; -&gt; SyntheticPrivateLocationUpdate   | [optional] 

## Methods

### NewSyntheticLocationUpdate

`func NewSyntheticLocationUpdate() *SyntheticLocationUpdate`

NewSyntheticLocationUpdate instantiates a new SyntheticLocationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticLocationUpdateWithDefaults

`func NewSyntheticLocationUpdateWithDefaults() *SyntheticLocationUpdate`

NewSyntheticLocationUpdateWithDefaults instantiates a new SyntheticLocationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SyntheticLocationUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticLocationUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticLocationUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SyntheticLocationUpdate) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


