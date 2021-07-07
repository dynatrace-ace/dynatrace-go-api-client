# UserSessionsQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumedMobileSessionsThisMonth** | Pointer to **float64** | Monthly Mobile user sessions environment consumption. Resets each calendar month. | [optional] [readonly] 
**TotalConsumedThisYear** | Pointer to **float64** | Yearly total User sessions environment consumption. Resets each year on license creation date anniversary. | [optional] [readonly] 
**ConsumedUserSessionsWithWebSessionReplayThisYear** | Pointer to **float64** | Yearly Web user sessions with replay environment consumption. Resets each year on license creation date anniversary. | [optional] [readonly] 
**TotalMonthlyLimit** | Pointer to **int64** | Monthly total User sessions environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited. | [optional] 
**TotalAnnualLimit** | Pointer to **int64** | Annual total User sessions environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited. | [optional] 
**TotalConsumedThisMonth** | Pointer to **float64** | Monthly total User sessions environment consumption. Resets each calendar month. | [optional] [readonly] 
**ConsumedUserSessionsWithWebSessionReplayThisMonth** | Pointer to **float64** | Monthly Web user sessions with replay environment consumption. Resets each calendar month. | [optional] [readonly] 
**ConsumedMobileSessionsThisYear** | Pointer to **float64** | Yearly Mobile user sessions environment consumption. Resets each year on license creation date anniversary. | [optional] [readonly] 
**ConsumedUserSessionsWithMobileSessionReplayThisMonth** | Pointer to **float64** | Monthly Mobile user sessions with replay environment consumption. Resets each calendar month. | [optional] [readonly] 
**ConsumedUserSessionsWithMobileSessionReplayThisYear** | Pointer to **float64** | Yearly Mobile user sessions with replay environment consumption. Resets each year on license creation date anniversary. | [optional] [readonly] 

## Methods

### NewUserSessionsQuota

`func NewUserSessionsQuota() *UserSessionsQuota`

NewUserSessionsQuota instantiates a new UserSessionsQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionsQuotaWithDefaults

`func NewUserSessionsQuotaWithDefaults() *UserSessionsQuota`

NewUserSessionsQuotaWithDefaults instantiates a new UserSessionsQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumedMobileSessionsThisMonth

`func (o *UserSessionsQuota) GetConsumedMobileSessionsThisMonth() float64`

GetConsumedMobileSessionsThisMonth returns the ConsumedMobileSessionsThisMonth field if non-nil, zero value otherwise.

### GetConsumedMobileSessionsThisMonthOk

`func (o *UserSessionsQuota) GetConsumedMobileSessionsThisMonthOk() (*float64, bool)`

GetConsumedMobileSessionsThisMonthOk returns a tuple with the ConsumedMobileSessionsThisMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedMobileSessionsThisMonth

`func (o *UserSessionsQuota) SetConsumedMobileSessionsThisMonth(v float64)`

SetConsumedMobileSessionsThisMonth sets ConsumedMobileSessionsThisMonth field to given value.

### HasConsumedMobileSessionsThisMonth

`func (o *UserSessionsQuota) HasConsumedMobileSessionsThisMonth() bool`

HasConsumedMobileSessionsThisMonth returns a boolean if a field has been set.

### GetTotalConsumedThisYear

`func (o *UserSessionsQuota) GetTotalConsumedThisYear() float64`

GetTotalConsumedThisYear returns the TotalConsumedThisYear field if non-nil, zero value otherwise.

### GetTotalConsumedThisYearOk

`func (o *UserSessionsQuota) GetTotalConsumedThisYearOk() (*float64, bool)`

GetTotalConsumedThisYearOk returns a tuple with the TotalConsumedThisYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalConsumedThisYear

`func (o *UserSessionsQuota) SetTotalConsumedThisYear(v float64)`

SetTotalConsumedThisYear sets TotalConsumedThisYear field to given value.

### HasTotalConsumedThisYear

`func (o *UserSessionsQuota) HasTotalConsumedThisYear() bool`

HasTotalConsumedThisYear returns a boolean if a field has been set.

### GetConsumedUserSessionsWithWebSessionReplayThisYear

`func (o *UserSessionsQuota) GetConsumedUserSessionsWithWebSessionReplayThisYear() float64`

GetConsumedUserSessionsWithWebSessionReplayThisYear returns the ConsumedUserSessionsWithWebSessionReplayThisYear field if non-nil, zero value otherwise.

### GetConsumedUserSessionsWithWebSessionReplayThisYearOk

`func (o *UserSessionsQuota) GetConsumedUserSessionsWithWebSessionReplayThisYearOk() (*float64, bool)`

GetConsumedUserSessionsWithWebSessionReplayThisYearOk returns a tuple with the ConsumedUserSessionsWithWebSessionReplayThisYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedUserSessionsWithWebSessionReplayThisYear

`func (o *UserSessionsQuota) SetConsumedUserSessionsWithWebSessionReplayThisYear(v float64)`

SetConsumedUserSessionsWithWebSessionReplayThisYear sets ConsumedUserSessionsWithWebSessionReplayThisYear field to given value.

### HasConsumedUserSessionsWithWebSessionReplayThisYear

`func (o *UserSessionsQuota) HasConsumedUserSessionsWithWebSessionReplayThisYear() bool`

HasConsumedUserSessionsWithWebSessionReplayThisYear returns a boolean if a field has been set.

### GetTotalMonthlyLimit

`func (o *UserSessionsQuota) GetTotalMonthlyLimit() int64`

GetTotalMonthlyLimit returns the TotalMonthlyLimit field if non-nil, zero value otherwise.

### GetTotalMonthlyLimitOk

`func (o *UserSessionsQuota) GetTotalMonthlyLimitOk() (*int64, bool)`

GetTotalMonthlyLimitOk returns a tuple with the TotalMonthlyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMonthlyLimit

`func (o *UserSessionsQuota) SetTotalMonthlyLimit(v int64)`

SetTotalMonthlyLimit sets TotalMonthlyLimit field to given value.

### HasTotalMonthlyLimit

`func (o *UserSessionsQuota) HasTotalMonthlyLimit() bool`

HasTotalMonthlyLimit returns a boolean if a field has been set.

### GetTotalAnnualLimit

`func (o *UserSessionsQuota) GetTotalAnnualLimit() int64`

GetTotalAnnualLimit returns the TotalAnnualLimit field if non-nil, zero value otherwise.

### GetTotalAnnualLimitOk

`func (o *UserSessionsQuota) GetTotalAnnualLimitOk() (*int64, bool)`

GetTotalAnnualLimitOk returns a tuple with the TotalAnnualLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAnnualLimit

`func (o *UserSessionsQuota) SetTotalAnnualLimit(v int64)`

SetTotalAnnualLimit sets TotalAnnualLimit field to given value.

### HasTotalAnnualLimit

`func (o *UserSessionsQuota) HasTotalAnnualLimit() bool`

HasTotalAnnualLimit returns a boolean if a field has been set.

### GetTotalConsumedThisMonth

`func (o *UserSessionsQuota) GetTotalConsumedThisMonth() float64`

GetTotalConsumedThisMonth returns the TotalConsumedThisMonth field if non-nil, zero value otherwise.

### GetTotalConsumedThisMonthOk

`func (o *UserSessionsQuota) GetTotalConsumedThisMonthOk() (*float64, bool)`

GetTotalConsumedThisMonthOk returns a tuple with the TotalConsumedThisMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalConsumedThisMonth

`func (o *UserSessionsQuota) SetTotalConsumedThisMonth(v float64)`

SetTotalConsumedThisMonth sets TotalConsumedThisMonth field to given value.

### HasTotalConsumedThisMonth

`func (o *UserSessionsQuota) HasTotalConsumedThisMonth() bool`

HasTotalConsumedThisMonth returns a boolean if a field has been set.

### GetConsumedUserSessionsWithWebSessionReplayThisMonth

`func (o *UserSessionsQuota) GetConsumedUserSessionsWithWebSessionReplayThisMonth() float64`

GetConsumedUserSessionsWithWebSessionReplayThisMonth returns the ConsumedUserSessionsWithWebSessionReplayThisMonth field if non-nil, zero value otherwise.

### GetConsumedUserSessionsWithWebSessionReplayThisMonthOk

`func (o *UserSessionsQuota) GetConsumedUserSessionsWithWebSessionReplayThisMonthOk() (*float64, bool)`

GetConsumedUserSessionsWithWebSessionReplayThisMonthOk returns a tuple with the ConsumedUserSessionsWithWebSessionReplayThisMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedUserSessionsWithWebSessionReplayThisMonth

`func (o *UserSessionsQuota) SetConsumedUserSessionsWithWebSessionReplayThisMonth(v float64)`

SetConsumedUserSessionsWithWebSessionReplayThisMonth sets ConsumedUserSessionsWithWebSessionReplayThisMonth field to given value.

### HasConsumedUserSessionsWithWebSessionReplayThisMonth

`func (o *UserSessionsQuota) HasConsumedUserSessionsWithWebSessionReplayThisMonth() bool`

HasConsumedUserSessionsWithWebSessionReplayThisMonth returns a boolean if a field has been set.

### GetConsumedMobileSessionsThisYear

`func (o *UserSessionsQuota) GetConsumedMobileSessionsThisYear() float64`

GetConsumedMobileSessionsThisYear returns the ConsumedMobileSessionsThisYear field if non-nil, zero value otherwise.

### GetConsumedMobileSessionsThisYearOk

`func (o *UserSessionsQuota) GetConsumedMobileSessionsThisYearOk() (*float64, bool)`

GetConsumedMobileSessionsThisYearOk returns a tuple with the ConsumedMobileSessionsThisYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedMobileSessionsThisYear

`func (o *UserSessionsQuota) SetConsumedMobileSessionsThisYear(v float64)`

SetConsumedMobileSessionsThisYear sets ConsumedMobileSessionsThisYear field to given value.

### HasConsumedMobileSessionsThisYear

`func (o *UserSessionsQuota) HasConsumedMobileSessionsThisYear() bool`

HasConsumedMobileSessionsThisYear returns a boolean if a field has been set.

### GetConsumedUserSessionsWithMobileSessionReplayThisMonth

`func (o *UserSessionsQuota) GetConsumedUserSessionsWithMobileSessionReplayThisMonth() float64`

GetConsumedUserSessionsWithMobileSessionReplayThisMonth returns the ConsumedUserSessionsWithMobileSessionReplayThisMonth field if non-nil, zero value otherwise.

### GetConsumedUserSessionsWithMobileSessionReplayThisMonthOk

`func (o *UserSessionsQuota) GetConsumedUserSessionsWithMobileSessionReplayThisMonthOk() (*float64, bool)`

GetConsumedUserSessionsWithMobileSessionReplayThisMonthOk returns a tuple with the ConsumedUserSessionsWithMobileSessionReplayThisMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedUserSessionsWithMobileSessionReplayThisMonth

`func (o *UserSessionsQuota) SetConsumedUserSessionsWithMobileSessionReplayThisMonth(v float64)`

SetConsumedUserSessionsWithMobileSessionReplayThisMonth sets ConsumedUserSessionsWithMobileSessionReplayThisMonth field to given value.

### HasConsumedUserSessionsWithMobileSessionReplayThisMonth

`func (o *UserSessionsQuota) HasConsumedUserSessionsWithMobileSessionReplayThisMonth() bool`

HasConsumedUserSessionsWithMobileSessionReplayThisMonth returns a boolean if a field has been set.

### GetConsumedUserSessionsWithMobileSessionReplayThisYear

`func (o *UserSessionsQuota) GetConsumedUserSessionsWithMobileSessionReplayThisYear() float64`

GetConsumedUserSessionsWithMobileSessionReplayThisYear returns the ConsumedUserSessionsWithMobileSessionReplayThisYear field if non-nil, zero value otherwise.

### GetConsumedUserSessionsWithMobileSessionReplayThisYearOk

`func (o *UserSessionsQuota) GetConsumedUserSessionsWithMobileSessionReplayThisYearOk() (*float64, bool)`

GetConsumedUserSessionsWithMobileSessionReplayThisYearOk returns a tuple with the ConsumedUserSessionsWithMobileSessionReplayThisYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedUserSessionsWithMobileSessionReplayThisYear

`func (o *UserSessionsQuota) SetConsumedUserSessionsWithMobileSessionReplayThisYear(v float64)`

SetConsumedUserSessionsWithMobileSessionReplayThisYear sets ConsumedUserSessionsWithMobileSessionReplayThisYear field to given value.

### HasConsumedUserSessionsWithMobileSessionReplayThisYear

`func (o *UserSessionsQuota) HasConsumedUserSessionsWithMobileSessionReplayThisYear() bool`

HasConsumedUserSessionsWithMobileSessionReplayThisYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


