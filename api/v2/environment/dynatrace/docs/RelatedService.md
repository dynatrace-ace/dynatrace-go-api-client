# RelatedService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Dynatrace entity ID of the entity. | [optional] [readonly] 
**NumberOfAffectedEntities** | Pointer to **int64** | The number of affected entities related to the entity. | [optional] [readonly] 
**AffectedEntities** | Pointer to **[]string** | A list of affected entities related to the entity. | [optional] [readonly] 
**Exposure** | Pointer to **string** | The level of exposure of the service. | [optional] [readonly] 

## Methods

### NewRelatedService

`func NewRelatedService() *RelatedService`

NewRelatedService instantiates a new RelatedService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedServiceWithDefaults

`func NewRelatedServiceWithDefaults() *RelatedService`

NewRelatedServiceWithDefaults instantiates a new RelatedService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RelatedService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelatedService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelatedService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelatedService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumberOfAffectedEntities

`func (o *RelatedService) GetNumberOfAffectedEntities() int64`

GetNumberOfAffectedEntities returns the NumberOfAffectedEntities field if non-nil, zero value otherwise.

### GetNumberOfAffectedEntitiesOk

`func (o *RelatedService) GetNumberOfAffectedEntitiesOk() (*int64, bool)`

GetNumberOfAffectedEntitiesOk returns a tuple with the NumberOfAffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAffectedEntities

`func (o *RelatedService) SetNumberOfAffectedEntities(v int64)`

SetNumberOfAffectedEntities sets NumberOfAffectedEntities field to given value.

### HasNumberOfAffectedEntities

`func (o *RelatedService) HasNumberOfAffectedEntities() bool`

HasNumberOfAffectedEntities returns a boolean if a field has been set.

### GetAffectedEntities

`func (o *RelatedService) GetAffectedEntities() []string`

GetAffectedEntities returns the AffectedEntities field if non-nil, zero value otherwise.

### GetAffectedEntitiesOk

`func (o *RelatedService) GetAffectedEntitiesOk() (*[]string, bool)`

GetAffectedEntitiesOk returns a tuple with the AffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEntities

`func (o *RelatedService) SetAffectedEntities(v []string)`

SetAffectedEntities sets AffectedEntities field to given value.

### HasAffectedEntities

`func (o *RelatedService) HasAffectedEntities() bool`

HasAffectedEntities returns a boolean if a field has been set.

### GetExposure

`func (o *RelatedService) GetExposure() string`

GetExposure returns the Exposure field if non-nil, zero value otherwise.

### GetExposureOk

`func (o *RelatedService) GetExposureOk() (*string, bool)`

GetExposureOk returns a tuple with the Exposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposure

`func (o *RelatedService) SetExposure(v string)`

SetExposure sets Exposure field to given value.

### HasExposure

`func (o *RelatedService) HasExposure() bool`

HasExposure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


