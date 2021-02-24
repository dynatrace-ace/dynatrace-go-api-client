# PagerDutyNotificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | The name of the PagerDuty account. | 
**ServiceApiKey** | Pointer to **string** | The API key to access PagerDuty. | [optional] 
**ServiceName** | **string** | The name of the service. | 

## Methods

### NewPagerDutyNotificationConfig

`func NewPagerDutyNotificationConfig(account string, serviceName string, ) *PagerDutyNotificationConfig`

NewPagerDutyNotificationConfig instantiates a new PagerDutyNotificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagerDutyNotificationConfigWithDefaults

`func NewPagerDutyNotificationConfigWithDefaults() *PagerDutyNotificationConfig`

NewPagerDutyNotificationConfigWithDefaults instantiates a new PagerDutyNotificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *PagerDutyNotificationConfig) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PagerDutyNotificationConfig) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PagerDutyNotificationConfig) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetServiceApiKey

`func (o *PagerDutyNotificationConfig) GetServiceApiKey() string`

GetServiceApiKey returns the ServiceApiKey field if non-nil, zero value otherwise.

### GetServiceApiKeyOk

`func (o *PagerDutyNotificationConfig) GetServiceApiKeyOk() (*string, bool)`

GetServiceApiKeyOk returns a tuple with the ServiceApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceApiKey

`func (o *PagerDutyNotificationConfig) SetServiceApiKey(v string)`

SetServiceApiKey sets ServiceApiKey field to given value.

### HasServiceApiKey

`func (o *PagerDutyNotificationConfig) HasServiceApiKey() bool`

HasServiceApiKey returns a boolean if a field has been set.

### GetServiceName

`func (o *PagerDutyNotificationConfig) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *PagerDutyNotificationConfig) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *PagerDutyNotificationConfig) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


