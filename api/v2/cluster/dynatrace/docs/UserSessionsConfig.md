# UserSessionsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConcurrentSessionPolicyDto** | [**ConcurrentSessionPolicy**](ConcurrentSessionPolicy.md) |  | 
**AutomaticLogoutDto** | [**AutomaticLogoutConfiguration**](AutomaticLogoutConfiguration.md) |  | 

## Methods

### NewUserSessionsConfig

`func NewUserSessionsConfig(concurrentSessionPolicyDto ConcurrentSessionPolicy, automaticLogoutDto AutomaticLogoutConfiguration, ) *UserSessionsConfig`

NewUserSessionsConfig instantiates a new UserSessionsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionsConfigWithDefaults

`func NewUserSessionsConfigWithDefaults() *UserSessionsConfig`

NewUserSessionsConfigWithDefaults instantiates a new UserSessionsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcurrentSessionPolicyDto

`func (o *UserSessionsConfig) GetConcurrentSessionPolicyDto() ConcurrentSessionPolicy`

GetConcurrentSessionPolicyDto returns the ConcurrentSessionPolicyDto field if non-nil, zero value otherwise.

### GetConcurrentSessionPolicyDtoOk

`func (o *UserSessionsConfig) GetConcurrentSessionPolicyDtoOk() (*ConcurrentSessionPolicy, bool)`

GetConcurrentSessionPolicyDtoOk returns a tuple with the ConcurrentSessionPolicyDto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentSessionPolicyDto

`func (o *UserSessionsConfig) SetConcurrentSessionPolicyDto(v ConcurrentSessionPolicy)`

SetConcurrentSessionPolicyDto sets ConcurrentSessionPolicyDto field to given value.


### GetAutomaticLogoutDto

`func (o *UserSessionsConfig) GetAutomaticLogoutDto() AutomaticLogoutConfiguration`

GetAutomaticLogoutDto returns the AutomaticLogoutDto field if non-nil, zero value otherwise.

### GetAutomaticLogoutDtoOk

`func (o *UserSessionsConfig) GetAutomaticLogoutDtoOk() (*AutomaticLogoutConfiguration, bool)`

GetAutomaticLogoutDtoOk returns a tuple with the AutomaticLogoutDto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticLogoutDto

`func (o *UserSessionsConfig) SetAutomaticLogoutDto(v AutomaticLogoutConfiguration)`

SetAutomaticLogoutDto sets AutomaticLogoutDto field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


