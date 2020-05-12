# DiskEventAnomalyDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | **string** | The ID of the disk event rule. | [optional] [readonly] 
**Name** | **string** | The name of the disk event rule. | 
**Enabled** | **bool** | Disk event rule enabled/disabled. | 
**Metric** | **string** | The metric to monitor. | 
**Threshold** | **float64** | The threshold to trigger disk event.    * A percentage for &#x60;LowDiskSpace&#x60; or &#x60;LowInodes&#x60; metrics.   * In milliseconds for &#x60;ReadTimeExceeding&#x60; or &#x60;WriteTimeExceeding&#x60; metrics. | 
**Samples** | **int32** | The number of samples to evaluate. | 
**ViolatingSamples** | **int32** | The number of samples that must violate the threshold to trigger an event. Must not exceed the number of evaluated samples. | 
**DiskNameFilter** | [**DiskNameFilter**](DiskNameFilter.md) |  | [optional] 
**TagFilters** | [**[]TagFilter**](TagFilter.md) | Narrows the rule usage down to the hosts matching the specified tags. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


