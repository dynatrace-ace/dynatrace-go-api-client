# AlertingProfileSeverityRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeverityLevel** | **string** | The severity level to trigger the alert. | 
**TagFilter** | [**AlertingProfileTagFilter**](AlertingProfileTagFilter.md) |  | 
**DelayInMinutes** | **int32** | Send a notification if a problem remains open longer than *X* minutes. | 

## Methods

### NewAlertingProfileSeverityRule

`func NewAlertingProfileSeverityRule(severityLevel string, tagFilter AlertingProfileTagFilter, delayInMinutes int32, ) *AlertingProfileSeverityRule`

NewAlertingProfileSeverityRule instantiates a new AlertingProfileSeverityRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingProfileSeverityRuleWithDefaults

`func NewAlertingProfileSeverityRuleWithDefaults() *AlertingProfileSeverityRule`

NewAlertingProfileSeverityRuleWithDefaults instantiates a new AlertingProfileSeverityRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeverityLevel

`func (o *AlertingProfileSeverityRule) GetSeverityLevel() string`

GetSeverityLevel returns the SeverityLevel field if non-nil, zero value otherwise.

### GetSeverityLevelOk

`func (o *AlertingProfileSeverityRule) GetSeverityLevelOk() (*string, bool)`

GetSeverityLevelOk returns a tuple with the SeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevel

`func (o *AlertingProfileSeverityRule) SetSeverityLevel(v string)`

SetSeverityLevel sets SeverityLevel field to given value.


### GetTagFilter

`func (o *AlertingProfileSeverityRule) GetTagFilter() AlertingProfileTagFilter`

GetTagFilter returns the TagFilter field if non-nil, zero value otherwise.

### GetTagFilterOk

`func (o *AlertingProfileSeverityRule) GetTagFilterOk() (*AlertingProfileTagFilter, bool)`

GetTagFilterOk returns a tuple with the TagFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilter

`func (o *AlertingProfileSeverityRule) SetTagFilter(v AlertingProfileTagFilter)`

SetTagFilter sets TagFilter field to given value.


### GetDelayInMinutes

`func (o *AlertingProfileSeverityRule) GetDelayInMinutes() int32`

GetDelayInMinutes returns the DelayInMinutes field if non-nil, zero value otherwise.

### GetDelayInMinutesOk

`func (o *AlertingProfileSeverityRule) GetDelayInMinutesOk() (*int32, bool)`

GetDelayInMinutesOk returns a tuple with the DelayInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayInMinutes

`func (o *AlertingProfileSeverityRule) SetDelayInMinutes(v int32)`

SetDelayInMinutes sets DelayInMinutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


