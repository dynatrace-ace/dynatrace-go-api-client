# SyntheticPublicLocationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of the location:   * &#x60;ENABLED&#x60;: The location is displayed as active in the UI. You can assign monitors to the location.  * &#x60;DISABLED&#x60;: The location is displayed as inactive in the UI. You can&#39;t assign monitors to the location. Monitors already assigned to the location will stay there and will be executed from the location.  * &#x60;HIDDEN&#x60;: The location is not displayed in the UI. You can&#39;t assign monitors to the location. You can only set location as &#x60;HIDDEN&#x60; when no monitor is assigned to it. | 

## Methods

### NewSyntheticPublicLocationUpdate

`func NewSyntheticPublicLocationUpdate(status string, ) *SyntheticPublicLocationUpdate`

NewSyntheticPublicLocationUpdate instantiates a new SyntheticPublicLocationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticPublicLocationUpdateWithDefaults

`func NewSyntheticPublicLocationUpdateWithDefaults() *SyntheticPublicLocationUpdate`

NewSyntheticPublicLocationUpdateWithDefaults instantiates a new SyntheticPublicLocationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SyntheticPublicLocationUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticPublicLocationUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticPublicLocationUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


