# SessionPropertiesQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumedThisMonth** | Pointer to **float64** | Monthly environment consumption. Resets each calendar month. | [optional] [readonly] 
**ConsumedThisYear** | Pointer to **float64** | Yearly environment consumption. Resets each year on license creation date anniversary. | [optional] [readonly] 

## Methods

### NewSessionPropertiesQuota

`func NewSessionPropertiesQuota() *SessionPropertiesQuota`

NewSessionPropertiesQuota instantiates a new SessionPropertiesQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionPropertiesQuotaWithDefaults

`func NewSessionPropertiesQuotaWithDefaults() *SessionPropertiesQuota`

NewSessionPropertiesQuotaWithDefaults instantiates a new SessionPropertiesQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumedThisMonth

`func (o *SessionPropertiesQuota) GetConsumedThisMonth() float64`

GetConsumedThisMonth returns the ConsumedThisMonth field if non-nil, zero value otherwise.

### GetConsumedThisMonthOk

`func (o *SessionPropertiesQuota) GetConsumedThisMonthOk() (*float64, bool)`

GetConsumedThisMonthOk returns a tuple with the ConsumedThisMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedThisMonth

`func (o *SessionPropertiesQuota) SetConsumedThisMonth(v float64)`

SetConsumedThisMonth sets ConsumedThisMonth field to given value.

### HasConsumedThisMonth

`func (o *SessionPropertiesQuota) HasConsumedThisMonth() bool`

HasConsumedThisMonth returns a boolean if a field has been set.

### GetConsumedThisYear

`func (o *SessionPropertiesQuota) GetConsumedThisYear() float64`

GetConsumedThisYear returns the ConsumedThisYear field if non-nil, zero value otherwise.

### GetConsumedThisYearOk

`func (o *SessionPropertiesQuota) GetConsumedThisYearOk() (*float64, bool)`

GetConsumedThisYearOk returns a tuple with the ConsumedThisYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedThisYear

`func (o *SessionPropertiesQuota) SetConsumedThisYear(v float64)`

SetConsumedThisYear sets ConsumedThisYear field to given value.

### HasConsumedThisYear

`func (o *SessionPropertiesQuota) HasConsumedThisYear() bool`

HasConsumedThisYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


