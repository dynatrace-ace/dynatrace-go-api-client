# ConcurrentSessionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserLimit** | **int32** | Session limit for regular users (0 &#x3D; no limit) | 
**AdminLimit** | **int32** | Session limit for admin users (0 &#x3D; no limit) | 

## Methods

### NewConcurrentSessionPolicy

`func NewConcurrentSessionPolicy(userLimit int32, adminLimit int32, ) *ConcurrentSessionPolicy`

NewConcurrentSessionPolicy instantiates a new ConcurrentSessionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConcurrentSessionPolicyWithDefaults

`func NewConcurrentSessionPolicyWithDefaults() *ConcurrentSessionPolicy`

NewConcurrentSessionPolicyWithDefaults instantiates a new ConcurrentSessionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserLimit

`func (o *ConcurrentSessionPolicy) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *ConcurrentSessionPolicy) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *ConcurrentSessionPolicy) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.


### GetAdminLimit

`func (o *ConcurrentSessionPolicy) GetAdminLimit() int32`

GetAdminLimit returns the AdminLimit field if non-nil, zero value otherwise.

### GetAdminLimitOk

`func (o *ConcurrentSessionPolicy) GetAdminLimitOk() (*int32, bool)`

GetAdminLimitOk returns a tuple with the AdminLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLimit

`func (o *ConcurrentSessionPolicy) SetAdminLimit(v int32)`

SetAdminLimit sets AdminLimit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


