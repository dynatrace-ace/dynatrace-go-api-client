# ResourceTimingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**W3cResourceTimings** | **bool** | W3C resource timings for third party/CDN enabled/disabled. | 
**NonW3cResourceTimings** | **bool** | Timing for JavaScript files and images on non-W3C supported browsers enabled/disabled. | 
**NonW3cResourceTimingsInstrumentationDelay** | **int32** | Instrumentation delay for monitoring resource and image resource impact in browsers that don&#39;t offer W3C resource timings.   Valid values range from 0 to 9999.  Only effective if **nonW3cResourceTimings** is enabled. | 
**ResourceTimingCaptureType** | **string** | Defines how detailed resource timings are captured.  Only effective if **w3cResourceTimings** or **nonW3cResourceTimings** is enabled. | 
**ResourceTimingsDomainLimit** | **int32** | Limits the number of domains for which W3C resource timings are captured.  Only effective if **resourceTimingCaptureType** is &#x60;CAPTURE_LIMITED_SUMMARIES&#x60;. | 

## Methods

### NewResourceTimingSettings

`func NewResourceTimingSettings(w3cResourceTimings bool, nonW3cResourceTimings bool, nonW3cResourceTimingsInstrumentationDelay int32, resourceTimingCaptureType string, resourceTimingsDomainLimit int32, ) *ResourceTimingSettings`

NewResourceTimingSettings instantiates a new ResourceTimingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTimingSettingsWithDefaults

`func NewResourceTimingSettingsWithDefaults() *ResourceTimingSettings`

NewResourceTimingSettingsWithDefaults instantiates a new ResourceTimingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetW3cResourceTimings

`func (o *ResourceTimingSettings) GetW3cResourceTimings() bool`

GetW3cResourceTimings returns the W3cResourceTimings field if non-nil, zero value otherwise.

### GetW3cResourceTimingsOk

`func (o *ResourceTimingSettings) GetW3cResourceTimingsOk() (*bool, bool)`

GetW3cResourceTimingsOk returns a tuple with the W3cResourceTimings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW3cResourceTimings

`func (o *ResourceTimingSettings) SetW3cResourceTimings(v bool)`

SetW3cResourceTimings sets W3cResourceTimings field to given value.


### GetNonW3cResourceTimings

`func (o *ResourceTimingSettings) GetNonW3cResourceTimings() bool`

GetNonW3cResourceTimings returns the NonW3cResourceTimings field if non-nil, zero value otherwise.

### GetNonW3cResourceTimingsOk

`func (o *ResourceTimingSettings) GetNonW3cResourceTimingsOk() (*bool, bool)`

GetNonW3cResourceTimingsOk returns a tuple with the NonW3cResourceTimings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonW3cResourceTimings

`func (o *ResourceTimingSettings) SetNonW3cResourceTimings(v bool)`

SetNonW3cResourceTimings sets NonW3cResourceTimings field to given value.


### GetNonW3cResourceTimingsInstrumentationDelay

`func (o *ResourceTimingSettings) GetNonW3cResourceTimingsInstrumentationDelay() int32`

GetNonW3cResourceTimingsInstrumentationDelay returns the NonW3cResourceTimingsInstrumentationDelay field if non-nil, zero value otherwise.

### GetNonW3cResourceTimingsInstrumentationDelayOk

`func (o *ResourceTimingSettings) GetNonW3cResourceTimingsInstrumentationDelayOk() (*int32, bool)`

GetNonW3cResourceTimingsInstrumentationDelayOk returns a tuple with the NonW3cResourceTimingsInstrumentationDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonW3cResourceTimingsInstrumentationDelay

`func (o *ResourceTimingSettings) SetNonW3cResourceTimingsInstrumentationDelay(v int32)`

SetNonW3cResourceTimingsInstrumentationDelay sets NonW3cResourceTimingsInstrumentationDelay field to given value.


### GetResourceTimingCaptureType

`func (o *ResourceTimingSettings) GetResourceTimingCaptureType() string`

GetResourceTimingCaptureType returns the ResourceTimingCaptureType field if non-nil, zero value otherwise.

### GetResourceTimingCaptureTypeOk

`func (o *ResourceTimingSettings) GetResourceTimingCaptureTypeOk() (*string, bool)`

GetResourceTimingCaptureTypeOk returns a tuple with the ResourceTimingCaptureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTimingCaptureType

`func (o *ResourceTimingSettings) SetResourceTimingCaptureType(v string)`

SetResourceTimingCaptureType sets ResourceTimingCaptureType field to given value.


### GetResourceTimingsDomainLimit

`func (o *ResourceTimingSettings) GetResourceTimingsDomainLimit() int32`

GetResourceTimingsDomainLimit returns the ResourceTimingsDomainLimit field if non-nil, zero value otherwise.

### GetResourceTimingsDomainLimitOk

`func (o *ResourceTimingSettings) GetResourceTimingsDomainLimitOk() (*int32, bool)`

GetResourceTimingsDomainLimitOk returns a tuple with the ResourceTimingsDomainLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTimingsDomainLimit

`func (o *ResourceTimingSettings) SetResourceTimingsDomainLimit(v int32)`

SetResourceTimingsDomainLimit sets ResourceTimingsDomainLimit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


