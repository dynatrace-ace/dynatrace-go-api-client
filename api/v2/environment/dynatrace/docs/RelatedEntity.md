# RelatedEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Dynatrace entity ID of the entity. | [optional] [readonly] 
**NumberOfAffectedEntities** | Pointer to **int64** | The number of affected entities related to the entity. | [optional] [readonly] 
**AffectedEntities** | Pointer to **[]string** | A list of affected entities related to the entity. | [optional] [readonly] 

## Methods

### NewRelatedEntity

`func NewRelatedEntity() *RelatedEntity`

NewRelatedEntity instantiates a new RelatedEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedEntityWithDefaults

`func NewRelatedEntityWithDefaults() *RelatedEntity`

NewRelatedEntityWithDefaults instantiates a new RelatedEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RelatedEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelatedEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelatedEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelatedEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumberOfAffectedEntities

`func (o *RelatedEntity) GetNumberOfAffectedEntities() int64`

GetNumberOfAffectedEntities returns the NumberOfAffectedEntities field if non-nil, zero value otherwise.

### GetNumberOfAffectedEntitiesOk

`func (o *RelatedEntity) GetNumberOfAffectedEntitiesOk() (*int64, bool)`

GetNumberOfAffectedEntitiesOk returns a tuple with the NumberOfAffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAffectedEntities

`func (o *RelatedEntity) SetNumberOfAffectedEntities(v int64)`

SetNumberOfAffectedEntities sets NumberOfAffectedEntities field to given value.

### HasNumberOfAffectedEntities

`func (o *RelatedEntity) HasNumberOfAffectedEntities() bool`

HasNumberOfAffectedEntities returns a boolean if a field has been set.

### GetAffectedEntities

`func (o *RelatedEntity) GetAffectedEntities() []string`

GetAffectedEntities returns the AffectedEntities field if non-nil, zero value otherwise.

### GetAffectedEntitiesOk

`func (o *RelatedEntity) GetAffectedEntitiesOk() (*[]string, bool)`

GetAffectedEntitiesOk returns a tuple with the AffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEntities

`func (o *RelatedEntity) SetAffectedEntities(v []string)`

SetAffectedEntities sets AffectedEntities field to given value.

### HasAffectedEntities

`func (o *RelatedEntity) HasAffectedEntities() bool`

HasAffectedEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


