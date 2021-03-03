# ServiceBaselineValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The ID of the service. | 
**DisplayName** | Pointer to **string** | The display name of the service. | [optional] 
**ServiceResponseTimeBaselines** | Pointer to [**[]EntityBaselineData**](EntityBaselineData.md) | The baseline data for the **Response time** duration metric. | [optional] 

## Methods

### NewServiceBaselineValues

`func NewServiceBaselineValues(entityId string, ) *ServiceBaselineValues`

NewServiceBaselineValues instantiates a new ServiceBaselineValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceBaselineValuesWithDefaults

`func NewServiceBaselineValuesWithDefaults() *ServiceBaselineValues`

NewServiceBaselineValuesWithDefaults instantiates a new ServiceBaselineValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *ServiceBaselineValues) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ServiceBaselineValues) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ServiceBaselineValues) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetDisplayName

`func (o *ServiceBaselineValues) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ServiceBaselineValues) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ServiceBaselineValues) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ServiceBaselineValues) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetServiceResponseTimeBaselines

`func (o *ServiceBaselineValues) GetServiceResponseTimeBaselines() []EntityBaselineData`

GetServiceResponseTimeBaselines returns the ServiceResponseTimeBaselines field if non-nil, zero value otherwise.

### GetServiceResponseTimeBaselinesOk

`func (o *ServiceBaselineValues) GetServiceResponseTimeBaselinesOk() (*[]EntityBaselineData, bool)`

GetServiceResponseTimeBaselinesOk returns a tuple with the ServiceResponseTimeBaselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceResponseTimeBaselines

`func (o *ServiceBaselineValues) SetServiceResponseTimeBaselines(v []EntityBaselineData)`

SetServiceResponseTimeBaselines sets ServiceResponseTimeBaselines field to given value.

### HasServiceResponseTimeBaselines

`func (o *ServiceBaselineValues) HasServiceResponseTimeBaselines() bool`

HasServiceResponseTimeBaselines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


