# MobileCustomApdex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToleratedThreshold** | **int32** | Apdex **tolerable** threshold, in milliseconds: a duration greater than or equal to this value is considered tolerable. | 
**FrustratingThreshold** | **int32** | Apdex **frustrated** threshold, in milliseconds: a duration greater than or equal to this value is considered frustrated. | 
**FrustratedOnError** | **bool** | Apdex error condition: if &#x60;true&#x60; the user session is considered frustrated when an error is reported. | 

## Methods

### NewMobileCustomApdex

`func NewMobileCustomApdex(toleratedThreshold int32, frustratingThreshold int32, frustratedOnError bool, ) *MobileCustomApdex`

NewMobileCustomApdex instantiates a new MobileCustomApdex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileCustomApdexWithDefaults

`func NewMobileCustomApdexWithDefaults() *MobileCustomApdex`

NewMobileCustomApdexWithDefaults instantiates a new MobileCustomApdex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToleratedThreshold

`func (o *MobileCustomApdex) GetToleratedThreshold() int32`

GetToleratedThreshold returns the ToleratedThreshold field if non-nil, zero value otherwise.

### GetToleratedThresholdOk

`func (o *MobileCustomApdex) GetToleratedThresholdOk() (*int32, bool)`

GetToleratedThresholdOk returns a tuple with the ToleratedThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToleratedThreshold

`func (o *MobileCustomApdex) SetToleratedThreshold(v int32)`

SetToleratedThreshold sets ToleratedThreshold field to given value.


### GetFrustratingThreshold

`func (o *MobileCustomApdex) GetFrustratingThreshold() int32`

GetFrustratingThreshold returns the FrustratingThreshold field if non-nil, zero value otherwise.

### GetFrustratingThresholdOk

`func (o *MobileCustomApdex) GetFrustratingThresholdOk() (*int32, bool)`

GetFrustratingThresholdOk returns a tuple with the FrustratingThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrustratingThreshold

`func (o *MobileCustomApdex) SetFrustratingThreshold(v int32)`

SetFrustratingThreshold sets FrustratingThreshold field to given value.


### GetFrustratedOnError

`func (o *MobileCustomApdex) GetFrustratedOnError() bool`

GetFrustratedOnError returns the FrustratedOnError field if non-nil, zero value otherwise.

### GetFrustratedOnErrorOk

`func (o *MobileCustomApdex) GetFrustratedOnErrorOk() (*bool, bool)`

GetFrustratedOnErrorOk returns a tuple with the FrustratedOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrustratedOnError

`func (o *MobileCustomApdex) SetFrustratedOnError(v bool)`

SetFrustratedOnError sets FrustratedOnError field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


