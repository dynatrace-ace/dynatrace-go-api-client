# WaterfallSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UncompressedResourcesThreshold** | **int32** | Warn about uncompressed resources larger than *X* bytes. | 
**ResourcesThreshold** | **int32** | Warn about resources larger than *X* bytes. | 
**ResourceBrowserCachingThreshold** | **int32** | Warn about resources with a lower browser cache rate above *X*%. | 
**SlowFirstPartyResourcesThreshold** | **int32** | Warn about slow 1st party resources with a response time above *X* ms. | 
**SlowThirdPartyResourcesThreshold** | **int32** | Warn about slow 3rd party resources with a response time above *X* ms. | 
**SlowCdnResourcesThreshold** | **int32** | Warn about slow CDN resources with a response time above *X* ms. | 
**SpeedIndexVisuallyCompleteRatioThreshold** | **int32** | Warn if Speed index exceeds *X* % of Visually complete. | 

## Methods

### NewWaterfallSettings

`func NewWaterfallSettings(uncompressedResourcesThreshold int32, resourcesThreshold int32, resourceBrowserCachingThreshold int32, slowFirstPartyResourcesThreshold int32, slowThirdPartyResourcesThreshold int32, slowCdnResourcesThreshold int32, speedIndexVisuallyCompleteRatioThreshold int32, ) *WaterfallSettings`

NewWaterfallSettings instantiates a new WaterfallSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaterfallSettingsWithDefaults

`func NewWaterfallSettingsWithDefaults() *WaterfallSettings`

NewWaterfallSettingsWithDefaults instantiates a new WaterfallSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUncompressedResourcesThreshold

`func (o *WaterfallSettings) GetUncompressedResourcesThreshold() int32`

GetUncompressedResourcesThreshold returns the UncompressedResourcesThreshold field if non-nil, zero value otherwise.

### GetUncompressedResourcesThresholdOk

`func (o *WaterfallSettings) GetUncompressedResourcesThresholdOk() (*int32, bool)`

GetUncompressedResourcesThresholdOk returns a tuple with the UncompressedResourcesThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncompressedResourcesThreshold

`func (o *WaterfallSettings) SetUncompressedResourcesThreshold(v int32)`

SetUncompressedResourcesThreshold sets UncompressedResourcesThreshold field to given value.


### GetResourcesThreshold

`func (o *WaterfallSettings) GetResourcesThreshold() int32`

GetResourcesThreshold returns the ResourcesThreshold field if non-nil, zero value otherwise.

### GetResourcesThresholdOk

`func (o *WaterfallSettings) GetResourcesThresholdOk() (*int32, bool)`

GetResourcesThresholdOk returns a tuple with the ResourcesThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesThreshold

`func (o *WaterfallSettings) SetResourcesThreshold(v int32)`

SetResourcesThreshold sets ResourcesThreshold field to given value.


### GetResourceBrowserCachingThreshold

`func (o *WaterfallSettings) GetResourceBrowserCachingThreshold() int32`

GetResourceBrowserCachingThreshold returns the ResourceBrowserCachingThreshold field if non-nil, zero value otherwise.

### GetResourceBrowserCachingThresholdOk

`func (o *WaterfallSettings) GetResourceBrowserCachingThresholdOk() (*int32, bool)`

GetResourceBrowserCachingThresholdOk returns a tuple with the ResourceBrowserCachingThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceBrowserCachingThreshold

`func (o *WaterfallSettings) SetResourceBrowserCachingThreshold(v int32)`

SetResourceBrowserCachingThreshold sets ResourceBrowserCachingThreshold field to given value.


### GetSlowFirstPartyResourcesThreshold

`func (o *WaterfallSettings) GetSlowFirstPartyResourcesThreshold() int32`

GetSlowFirstPartyResourcesThreshold returns the SlowFirstPartyResourcesThreshold field if non-nil, zero value otherwise.

### GetSlowFirstPartyResourcesThresholdOk

`func (o *WaterfallSettings) GetSlowFirstPartyResourcesThresholdOk() (*int32, bool)`

GetSlowFirstPartyResourcesThresholdOk returns a tuple with the SlowFirstPartyResourcesThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowFirstPartyResourcesThreshold

`func (o *WaterfallSettings) SetSlowFirstPartyResourcesThreshold(v int32)`

SetSlowFirstPartyResourcesThreshold sets SlowFirstPartyResourcesThreshold field to given value.


### GetSlowThirdPartyResourcesThreshold

`func (o *WaterfallSettings) GetSlowThirdPartyResourcesThreshold() int32`

GetSlowThirdPartyResourcesThreshold returns the SlowThirdPartyResourcesThreshold field if non-nil, zero value otherwise.

### GetSlowThirdPartyResourcesThresholdOk

`func (o *WaterfallSettings) GetSlowThirdPartyResourcesThresholdOk() (*int32, bool)`

GetSlowThirdPartyResourcesThresholdOk returns a tuple with the SlowThirdPartyResourcesThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowThirdPartyResourcesThreshold

`func (o *WaterfallSettings) SetSlowThirdPartyResourcesThreshold(v int32)`

SetSlowThirdPartyResourcesThreshold sets SlowThirdPartyResourcesThreshold field to given value.


### GetSlowCdnResourcesThreshold

`func (o *WaterfallSettings) GetSlowCdnResourcesThreshold() int32`

GetSlowCdnResourcesThreshold returns the SlowCdnResourcesThreshold field if non-nil, zero value otherwise.

### GetSlowCdnResourcesThresholdOk

`func (o *WaterfallSettings) GetSlowCdnResourcesThresholdOk() (*int32, bool)`

GetSlowCdnResourcesThresholdOk returns a tuple with the SlowCdnResourcesThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowCdnResourcesThreshold

`func (o *WaterfallSettings) SetSlowCdnResourcesThreshold(v int32)`

SetSlowCdnResourcesThreshold sets SlowCdnResourcesThreshold field to given value.


### GetSpeedIndexVisuallyCompleteRatioThreshold

`func (o *WaterfallSettings) GetSpeedIndexVisuallyCompleteRatioThreshold() int32`

GetSpeedIndexVisuallyCompleteRatioThreshold returns the SpeedIndexVisuallyCompleteRatioThreshold field if non-nil, zero value otherwise.

### GetSpeedIndexVisuallyCompleteRatioThresholdOk

`func (o *WaterfallSettings) GetSpeedIndexVisuallyCompleteRatioThresholdOk() (*int32, bool)`

GetSpeedIndexVisuallyCompleteRatioThresholdOk returns a tuple with the SpeedIndexVisuallyCompleteRatioThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedIndexVisuallyCompleteRatioThreshold

`func (o *WaterfallSettings) SetSpeedIndexVisuallyCompleteRatioThreshold(v int32)`

SetSpeedIndexVisuallyCompleteRatioThreshold sets SpeedIndexVisuallyCompleteRatioThreshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


