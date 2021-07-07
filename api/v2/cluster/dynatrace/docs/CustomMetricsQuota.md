# CustomMetricsQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentUsage** | Pointer to **float64** | Current environment usage. | [optional] [readonly] 
**MaxLimit** | Pointer to **int64** | Concurrent environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited. | [optional] 

## Methods

### NewCustomMetricsQuota

`func NewCustomMetricsQuota() *CustomMetricsQuota`

NewCustomMetricsQuota instantiates a new CustomMetricsQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomMetricsQuotaWithDefaults

`func NewCustomMetricsQuotaWithDefaults() *CustomMetricsQuota`

NewCustomMetricsQuotaWithDefaults instantiates a new CustomMetricsQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentUsage

`func (o *CustomMetricsQuota) GetCurrentUsage() float64`

GetCurrentUsage returns the CurrentUsage field if non-nil, zero value otherwise.

### GetCurrentUsageOk

`func (o *CustomMetricsQuota) GetCurrentUsageOk() (*float64, bool)`

GetCurrentUsageOk returns a tuple with the CurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsage

`func (o *CustomMetricsQuota) SetCurrentUsage(v float64)`

SetCurrentUsage sets CurrentUsage field to given value.

### HasCurrentUsage

`func (o *CustomMetricsQuota) HasCurrentUsage() bool`

HasCurrentUsage returns a boolean if a field has been set.

### GetMaxLimit

`func (o *CustomMetricsQuota) GetMaxLimit() int64`

GetMaxLimit returns the MaxLimit field if non-nil, zero value otherwise.

### GetMaxLimitOk

`func (o *CustomMetricsQuota) GetMaxLimitOk() (*int64, bool)`

GetMaxLimitOk returns a tuple with the MaxLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLimit

`func (o *CustomMetricsQuota) SetMaxLimit(v int64)`

SetMaxLimit sets MaxLimit field to given value.

### HasMaxLimit

`func (o *CustomMetricsQuota) HasMaxLimit() bool`

HasMaxLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


