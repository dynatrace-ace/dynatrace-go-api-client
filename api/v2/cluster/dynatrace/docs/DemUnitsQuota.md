# DemUnitsQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumedThisMonth** | Pointer to **float64** | Monthly environment consumption. Resets each calendar month. | [optional] [readonly] 
**ConsumedThisYear** | Pointer to **float64** | Yearly environment consumption. Resets each year on license creation date anniversary. | [optional] [readonly] 
**MonthlyLimit** | Pointer to **int64** | Monthly environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited. | [optional] 
**AnnualLimit** | Pointer to **int64** | Annual environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited. | [optional] 

## Methods

### NewDemUnitsQuota

`func NewDemUnitsQuota() *DemUnitsQuota`

NewDemUnitsQuota instantiates a new DemUnitsQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDemUnitsQuotaWithDefaults

`func NewDemUnitsQuotaWithDefaults() *DemUnitsQuota`

NewDemUnitsQuotaWithDefaults instantiates a new DemUnitsQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumedThisMonth

`func (o *DemUnitsQuota) GetConsumedThisMonth() float64`

GetConsumedThisMonth returns the ConsumedThisMonth field if non-nil, zero value otherwise.

### GetConsumedThisMonthOk

`func (o *DemUnitsQuota) GetConsumedThisMonthOk() (*float64, bool)`

GetConsumedThisMonthOk returns a tuple with the ConsumedThisMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedThisMonth

`func (o *DemUnitsQuota) SetConsumedThisMonth(v float64)`

SetConsumedThisMonth sets ConsumedThisMonth field to given value.

### HasConsumedThisMonth

`func (o *DemUnitsQuota) HasConsumedThisMonth() bool`

HasConsumedThisMonth returns a boolean if a field has been set.

### GetConsumedThisYear

`func (o *DemUnitsQuota) GetConsumedThisYear() float64`

GetConsumedThisYear returns the ConsumedThisYear field if non-nil, zero value otherwise.

### GetConsumedThisYearOk

`func (o *DemUnitsQuota) GetConsumedThisYearOk() (*float64, bool)`

GetConsumedThisYearOk returns a tuple with the ConsumedThisYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedThisYear

`func (o *DemUnitsQuota) SetConsumedThisYear(v float64)`

SetConsumedThisYear sets ConsumedThisYear field to given value.

### HasConsumedThisYear

`func (o *DemUnitsQuota) HasConsumedThisYear() bool`

HasConsumedThisYear returns a boolean if a field has been set.

### GetMonthlyLimit

`func (o *DemUnitsQuota) GetMonthlyLimit() int64`

GetMonthlyLimit returns the MonthlyLimit field if non-nil, zero value otherwise.

### GetMonthlyLimitOk

`func (o *DemUnitsQuota) GetMonthlyLimitOk() (*int64, bool)`

GetMonthlyLimitOk returns a tuple with the MonthlyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyLimit

`func (o *DemUnitsQuota) SetMonthlyLimit(v int64)`

SetMonthlyLimit sets MonthlyLimit field to given value.

### HasMonthlyLimit

`func (o *DemUnitsQuota) HasMonthlyLimit() bool`

HasMonthlyLimit returns a boolean if a field has been set.

### GetAnnualLimit

`func (o *DemUnitsQuota) GetAnnualLimit() int64`

GetAnnualLimit returns the AnnualLimit field if non-nil, zero value otherwise.

### GetAnnualLimitOk

`func (o *DemUnitsQuota) GetAnnualLimitOk() (*int64, bool)`

GetAnnualLimitOk returns a tuple with the AnnualLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualLimit

`func (o *DemUnitsQuota) SetAnnualLimit(v int64)`

SetAnnualLimit sets AnnualLimit field to given value.

### HasAnnualLimit

`func (o *DemUnitsQuota) HasAnnualLimit() bool`

HasAnnualLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


