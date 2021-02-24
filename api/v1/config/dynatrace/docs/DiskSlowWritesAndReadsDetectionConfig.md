# DiskSlowWritesAndReadsDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**CustomThresholds** | Pointer to [**DiskSlowWriteAndReadsThresholds**](DiskSlowWriteAndReadsThresholds.md) |  | [optional] 

## Methods

### NewDiskSlowWritesAndReadsDetectionConfig

`func NewDiskSlowWritesAndReadsDetectionConfig(enabled bool, ) *DiskSlowWritesAndReadsDetectionConfig`

NewDiskSlowWritesAndReadsDetectionConfig instantiates a new DiskSlowWritesAndReadsDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskSlowWritesAndReadsDetectionConfigWithDefaults

`func NewDiskSlowWritesAndReadsDetectionConfigWithDefaults() *DiskSlowWritesAndReadsDetectionConfig`

NewDiskSlowWritesAndReadsDetectionConfigWithDefaults instantiates a new DiskSlowWritesAndReadsDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DiskSlowWritesAndReadsDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DiskSlowWritesAndReadsDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DiskSlowWritesAndReadsDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCustomThresholds

`func (o *DiskSlowWritesAndReadsDetectionConfig) GetCustomThresholds() DiskSlowWriteAndReadsThresholds`

GetCustomThresholds returns the CustomThresholds field if non-nil, zero value otherwise.

### GetCustomThresholdsOk

`func (o *DiskSlowWritesAndReadsDetectionConfig) GetCustomThresholdsOk() (*DiskSlowWriteAndReadsThresholds, bool)`

GetCustomThresholdsOk returns a tuple with the CustomThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomThresholds

`func (o *DiskSlowWritesAndReadsDetectionConfig) SetCustomThresholds(v DiskSlowWriteAndReadsThresholds)`

SetCustomThresholds sets CustomThresholds field to given value.

### HasCustomThresholds

`func (o *DiskSlowWritesAndReadsDetectionConfig) HasCustomThresholds() bool`

HasCustomThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


