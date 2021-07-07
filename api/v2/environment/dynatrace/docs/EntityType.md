# EntityType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DimensionKey** | Pointer to **string** | The dimension key used within metrics for this monitored entity. | [optional] 
**EntityLimitExceeded** | Pointer to **bool** | Whether the entity creation limit for the given type has been exceeded | [optional] 
**ManagementZones** | Pointer to **string** | The placeholder for the list of management zones of an actual entity. | [optional] 
**FromRelationships** | Pointer to [**[]ToPosition**](ToPosition.md) | A list of possible relationships where the monitored entity type occupies the FROM position | [optional] 
**ToRelationships** | Pointer to [**[]FromPosition**](FromPosition.md) | A list of possible relationships where the monitored entity type occupies the TO position. | [optional] 
**Properties** | Pointer to [**[]EntityTypePropertyDto**](EntityTypePropertyDto.md) | A list of additional properties of the monitored entity type. | [optional] 
**Type** | Pointer to **string** | The type of the monitored entity. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the monitored entity. | [optional] 
**Tags** | Pointer to **string** | The placeholder for the list of tags of an actual entity. | [optional] 

## Methods

### NewEntityType

`func NewEntityType() *EntityType`

NewEntityType instantiates a new EntityType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityTypeWithDefaults

`func NewEntityTypeWithDefaults() *EntityType`

NewEntityTypeWithDefaults instantiates a new EntityType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensionKey

`func (o *EntityType) GetDimensionKey() string`

GetDimensionKey returns the DimensionKey field if non-nil, zero value otherwise.

### GetDimensionKeyOk

`func (o *EntityType) GetDimensionKeyOk() (*string, bool)`

GetDimensionKeyOk returns a tuple with the DimensionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionKey

`func (o *EntityType) SetDimensionKey(v string)`

SetDimensionKey sets DimensionKey field to given value.

### HasDimensionKey

`func (o *EntityType) HasDimensionKey() bool`

HasDimensionKey returns a boolean if a field has been set.

### GetEntityLimitExceeded

`func (o *EntityType) GetEntityLimitExceeded() bool`

GetEntityLimitExceeded returns the EntityLimitExceeded field if non-nil, zero value otherwise.

### GetEntityLimitExceededOk

`func (o *EntityType) GetEntityLimitExceededOk() (*bool, bool)`

GetEntityLimitExceededOk returns a tuple with the EntityLimitExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityLimitExceeded

`func (o *EntityType) SetEntityLimitExceeded(v bool)`

SetEntityLimitExceeded sets EntityLimitExceeded field to given value.

### HasEntityLimitExceeded

`func (o *EntityType) HasEntityLimitExceeded() bool`

HasEntityLimitExceeded returns a boolean if a field has been set.

### GetManagementZones

`func (o *EntityType) GetManagementZones() string`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *EntityType) GetManagementZonesOk() (*string, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *EntityType) SetManagementZones(v string)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *EntityType) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetFromRelationships

`func (o *EntityType) GetFromRelationships() []ToPosition`

GetFromRelationships returns the FromRelationships field if non-nil, zero value otherwise.

### GetFromRelationshipsOk

`func (o *EntityType) GetFromRelationshipsOk() (*[]ToPosition, bool)`

GetFromRelationshipsOk returns a tuple with the FromRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromRelationships

`func (o *EntityType) SetFromRelationships(v []ToPosition)`

SetFromRelationships sets FromRelationships field to given value.

### HasFromRelationships

`func (o *EntityType) HasFromRelationships() bool`

HasFromRelationships returns a boolean if a field has been set.

### GetToRelationships

`func (o *EntityType) GetToRelationships() []FromPosition`

GetToRelationships returns the ToRelationships field if non-nil, zero value otherwise.

### GetToRelationshipsOk

`func (o *EntityType) GetToRelationshipsOk() (*[]FromPosition, bool)`

GetToRelationshipsOk returns a tuple with the ToRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRelationships

`func (o *EntityType) SetToRelationships(v []FromPosition)`

SetToRelationships sets ToRelationships field to given value.

### HasToRelationships

`func (o *EntityType) HasToRelationships() bool`

HasToRelationships returns a boolean if a field has been set.

### GetProperties

`func (o *EntityType) GetProperties() []EntityTypePropertyDto`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EntityType) GetPropertiesOk() (*[]EntityTypePropertyDto, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EntityType) SetProperties(v []EntityTypePropertyDto)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *EntityType) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetType

`func (o *EntityType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntityType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayName

`func (o *EntityType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EntityType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EntityType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EntityType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTags

`func (o *EntityType) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EntityType) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EntityType) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EntityType) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


