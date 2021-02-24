# DashboardReportSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WEEK** | Pointer to **[]string** | A list of weekly subscribers.   Weekly subscribers receive the report every Monday at midnight.    You can specify email addresses or Dynatrace user IDs here. | [optional] 
**MONTH** | Pointer to **[]string** | A list of monthly subscribers.   Monthly subscribers receive the report on the first Monday of the month at midnight.   You can specify email addresses or Dynatrace user IDs here. | [optional] 

## Methods

### NewDashboardReportSubscription

`func NewDashboardReportSubscription() *DashboardReportSubscription`

NewDashboardReportSubscription instantiates a new DashboardReportSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardReportSubscriptionWithDefaults

`func NewDashboardReportSubscriptionWithDefaults() *DashboardReportSubscription`

NewDashboardReportSubscriptionWithDefaults instantiates a new DashboardReportSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWEEK

`func (o *DashboardReportSubscription) GetWEEK() []string`

GetWEEK returns the WEEK field if non-nil, zero value otherwise.

### GetWEEKOk

`func (o *DashboardReportSubscription) GetWEEKOk() (*[]string, bool)`

GetWEEKOk returns a tuple with the WEEK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWEEK

`func (o *DashboardReportSubscription) SetWEEK(v []string)`

SetWEEK sets WEEK field to given value.

### HasWEEK

`func (o *DashboardReportSubscription) HasWEEK() bool`

HasWEEK returns a boolean if a field has been set.

### GetMONTH

`func (o *DashboardReportSubscription) GetMONTH() []string`

GetMONTH returns the MONTH field if non-nil, zero value otherwise.

### GetMONTHOk

`func (o *DashboardReportSubscription) GetMONTHOk() (*[]string, bool)`

GetMONTHOk returns a tuple with the MONTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMONTH

`func (o *DashboardReportSubscription) SetMONTH(v []string)`

SetMONTH sets MONTH field to given value.

### HasMONTH

`func (o *DashboardReportSubscription) HasMONTH() bool`

HasMONTH returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


