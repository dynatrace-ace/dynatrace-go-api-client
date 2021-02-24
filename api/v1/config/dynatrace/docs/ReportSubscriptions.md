# ReportSubscriptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | **string** | The schedule of the subscription.    * Weekly subscribers receive the report every Monday at midnight.   * Monthly subscribers receive the report on the first Monday of the month at midnight. | 
**Recipients** | **[]string** | A list of the recipients.   You can specify email addresses or Dynatrace user IDs here. | 

## Methods

### NewReportSubscriptions

`func NewReportSubscriptions(schedule string, recipients []string, ) *ReportSubscriptions`

NewReportSubscriptions instantiates a new ReportSubscriptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportSubscriptionsWithDefaults

`func NewReportSubscriptionsWithDefaults() *ReportSubscriptions`

NewReportSubscriptionsWithDefaults instantiates a new ReportSubscriptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *ReportSubscriptions) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ReportSubscriptions) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ReportSubscriptions) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetRecipients

`func (o *ReportSubscriptions) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ReportSubscriptions) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ReportSubscriptions) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


