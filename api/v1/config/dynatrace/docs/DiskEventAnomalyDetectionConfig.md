# DiskEventAnomalyDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the disk event rule. | [optional] [readonly] 
**Name** | **string** | The name of the disk event rule. | 
**Enabled** | **bool** | Disk event rule enabled/disabled. | 
**Metric** | **string** | The metric to monitor. | 
**Threshold** | **float64** | The threshold to trigger disk event.    * A percentage for &#x60;LowDiskSpace&#x60; or &#x60;LowInodes&#x60; metrics.   * In milliseconds for &#x60;ReadTimeExceeding&#x60; or &#x60;WriteTimeExceeding&#x60; metrics. | 
**Samples** | **int32** | The number of samples to evaluate. | 
**ViolatingSamples** | **int32** | The number of samples that must violate the threshold to trigger an event. Must not exceed the number of evaluated samples. | 
**DiskNameFilter** | Pointer to [**DiskNameFilter**](DiskNameFilter.md) |  | [optional] 
**TagFilters** | Pointer to [**[]TagFilter**](TagFilter.md) | Narrows the rule usage down to the hosts matching the specified tags. | [optional] 
**HostGroupId** | Pointer to **string** | Narrows the rule usage down to disks that run on hosts that themselves run on the specified host group. | [optional] 

## Methods

### NewDiskEventAnomalyDetectionConfig

`func NewDiskEventAnomalyDetectionConfig(name string, enabled bool, metric string, threshold float64, samples int32, violatingSamples int32, ) *DiskEventAnomalyDetectionConfig`

NewDiskEventAnomalyDetectionConfig instantiates a new DiskEventAnomalyDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskEventAnomalyDetectionConfigWithDefaults

`func NewDiskEventAnomalyDetectionConfigWithDefaults() *DiskEventAnomalyDetectionConfig`

NewDiskEventAnomalyDetectionConfigWithDefaults instantiates a new DiskEventAnomalyDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *DiskEventAnomalyDetectionConfig) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DiskEventAnomalyDetectionConfig) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DiskEventAnomalyDetectionConfig) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DiskEventAnomalyDetectionConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *DiskEventAnomalyDetectionConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskEventAnomalyDetectionConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskEventAnomalyDetectionConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiskEventAnomalyDetectionConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DiskEventAnomalyDetectionConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiskEventAnomalyDetectionConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiskEventAnomalyDetectionConfig) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *DiskEventAnomalyDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DiskEventAnomalyDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DiskEventAnomalyDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMetric

`func (o *DiskEventAnomalyDetectionConfig) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *DiskEventAnomalyDetectionConfig) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *DiskEventAnomalyDetectionConfig) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetThreshold

`func (o *DiskEventAnomalyDetectionConfig) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *DiskEventAnomalyDetectionConfig) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *DiskEventAnomalyDetectionConfig) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.


### GetSamples

`func (o *DiskEventAnomalyDetectionConfig) GetSamples() int32`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *DiskEventAnomalyDetectionConfig) GetSamplesOk() (*int32, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *DiskEventAnomalyDetectionConfig) SetSamples(v int32)`

SetSamples sets Samples field to given value.


### GetViolatingSamples

`func (o *DiskEventAnomalyDetectionConfig) GetViolatingSamples() int32`

GetViolatingSamples returns the ViolatingSamples field if non-nil, zero value otherwise.

### GetViolatingSamplesOk

`func (o *DiskEventAnomalyDetectionConfig) GetViolatingSamplesOk() (*int32, bool)`

GetViolatingSamplesOk returns a tuple with the ViolatingSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolatingSamples

`func (o *DiskEventAnomalyDetectionConfig) SetViolatingSamples(v int32)`

SetViolatingSamples sets ViolatingSamples field to given value.


### GetDiskNameFilter

`func (o *DiskEventAnomalyDetectionConfig) GetDiskNameFilter() DiskNameFilter`

GetDiskNameFilter returns the DiskNameFilter field if non-nil, zero value otherwise.

### GetDiskNameFilterOk

`func (o *DiskEventAnomalyDetectionConfig) GetDiskNameFilterOk() (*DiskNameFilter, bool)`

GetDiskNameFilterOk returns a tuple with the DiskNameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskNameFilter

`func (o *DiskEventAnomalyDetectionConfig) SetDiskNameFilter(v DiskNameFilter)`

SetDiskNameFilter sets DiskNameFilter field to given value.

### HasDiskNameFilter

`func (o *DiskEventAnomalyDetectionConfig) HasDiskNameFilter() bool`

HasDiskNameFilter returns a boolean if a field has been set.

### GetTagFilters

`func (o *DiskEventAnomalyDetectionConfig) GetTagFilters() []TagFilter`

GetTagFilters returns the TagFilters field if non-nil, zero value otherwise.

### GetTagFiltersOk

`func (o *DiskEventAnomalyDetectionConfig) GetTagFiltersOk() (*[]TagFilter, bool)`

GetTagFiltersOk returns a tuple with the TagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilters

`func (o *DiskEventAnomalyDetectionConfig) SetTagFilters(v []TagFilter)`

SetTagFilters sets TagFilters field to given value.

### HasTagFilters

`func (o *DiskEventAnomalyDetectionConfig) HasTagFilters() bool`

HasTagFilters returns a boolean if a field has been set.

### GetHostGroupId

`func (o *DiskEventAnomalyDetectionConfig) GetHostGroupId() string`

GetHostGroupId returns the HostGroupId field if non-nil, zero value otherwise.

### GetHostGroupIdOk

`func (o *DiskEventAnomalyDetectionConfig) GetHostGroupIdOk() (*string, bool)`

GetHostGroupIdOk returns a tuple with the HostGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroupId

`func (o *DiskEventAnomalyDetectionConfig) SetHostGroupId(v string)`

SetHostGroupId sets HostGroupId field to given value.

### HasHostGroupId

`func (o *DiskEventAnomalyDetectionConfig) HasHostGroupId() bool`

HasHostGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


