# UserSessionErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of error. | [optional] 
**Name** | Pointer to **string** | The name of the error. | [optional] 
**Domain** | Pointer to **string** | The DNS domain where the error has been recorded. | [optional] 
**StartTime** | Pointer to **int64** | The timestamp of the error, in UTC milliseconds. | [optional] 
**Application** | Pointer to **string** | The name of the application, based on the configured detection rules. | [optional] 
**InternalApplicationId** | Pointer to **string** | The Dynatrace entity ID of the application.    This information is useful when calling various REST APIs, for example, as a key for time series queries. | [optional] 

## Methods

### NewUserSessionErrors

`func NewUserSessionErrors() *UserSessionErrors`

NewUserSessionErrors instantiates a new UserSessionErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionErrorsWithDefaults

`func NewUserSessionErrorsWithDefaults() *UserSessionErrors`

NewUserSessionErrorsWithDefaults instantiates a new UserSessionErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserSessionErrors) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSessionErrors) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSessionErrors) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserSessionErrors) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UserSessionErrors) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSessionErrors) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSessionErrors) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserSessionErrors) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomain

`func (o *UserSessionErrors) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UserSessionErrors) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UserSessionErrors) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UserSessionErrors) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetStartTime

`func (o *UserSessionErrors) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UserSessionErrors) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UserSessionErrors) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UserSessionErrors) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetApplication

`func (o *UserSessionErrors) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *UserSessionErrors) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *UserSessionErrors) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *UserSessionErrors) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetInternalApplicationId

`func (o *UserSessionErrors) GetInternalApplicationId() string`

GetInternalApplicationId returns the InternalApplicationId field if non-nil, zero value otherwise.

### GetInternalApplicationIdOk

`func (o *UserSessionErrors) GetInternalApplicationIdOk() (*string, bool)`

GetInternalApplicationIdOk returns a tuple with the InternalApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalApplicationId

`func (o *UserSessionErrors) SetInternalApplicationId(v string)`

SetInternalApplicationId sets InternalApplicationId field to given value.

### HasInternalApplicationId

`func (o *UserSessionErrors) HasInternalApplicationId() bool`

HasInternalApplicationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


