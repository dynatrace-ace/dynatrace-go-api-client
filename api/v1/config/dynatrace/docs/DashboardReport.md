# DashboardReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the report. | [optional] 
**Type** | **string** |  | 
**DashboardId** | **string** | The ID of the associated dashboard. | 
**Enabled** | Pointer to **bool** | The email notifications for the dashboard report are enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 
**Subscriptions** | [**DashboardReportSubscription**](DashboardReportSubscription.md) |  | 

## Methods

### NewDashboardReport

`func NewDashboardReport(type_ string, dashboardId string, subscriptions DashboardReportSubscription, ) *DashboardReport`

NewDashboardReport instantiates a new DashboardReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardReportWithDefaults

`func NewDashboardReportWithDefaults() *DashboardReport`

NewDashboardReportWithDefaults instantiates a new DashboardReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DashboardReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardReport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DashboardReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DashboardReport) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DashboardReport) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DashboardReport) SetType(v string)`

SetType sets Type field to given value.


### GetDashboardId

`func (o *DashboardReport) GetDashboardId() string`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *DashboardReport) GetDashboardIdOk() (*string, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *DashboardReport) SetDashboardId(v string)`

SetDashboardId sets DashboardId field to given value.


### GetEnabled

`func (o *DashboardReport) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DashboardReport) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DashboardReport) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DashboardReport) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSubscriptions

`func (o *DashboardReport) GetSubscriptions() DashboardReportSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *DashboardReport) GetSubscriptionsOk() (*DashboardReportSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *DashboardReport) SetSubscriptions(v DashboardReportSubscription)`

SetSubscriptions sets Subscriptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


