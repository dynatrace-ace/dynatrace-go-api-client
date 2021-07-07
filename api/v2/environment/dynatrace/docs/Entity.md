# Entity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagementZones** | Pointer to [**[]ManagementZone**](ManagementZone.md) | A set of management zones to which the entity belongs. | [optional] 
**FromRelationships** | Pointer to [**map[string][]EntityId**](array.md) | A list of relationships where the entity occupies the FROM position. | [optional] 
**ToRelationships** | Pointer to [**map[string][]EntityId**](array.md) | A list of relationships where the entity occupies the TO position. | [optional] 
**EntityId** | Pointer to **string** | The ID of the entity. | [optional] 
**LastSeenTms** | Pointer to **int64** | The timestamp at which the entity was last seen, in UTC milliseconds. | [optional] 
**FirstSeenTms** | Pointer to **int64** | The timestamp at which the entity was first seen, in UTC milliseconds. | [optional] 
**Icon** | Pointer to [**EntityIcon**](EntityIcon.md) |  | [optional] 
**Properties** | Pointer to **map[string]map[string]interface{}** | A list of additional properties of the entity. | [optional] 
**DisplayName** | Pointer to **string** | The name of the entity, displayed in the UI. | [optional] 
**Tags** | Pointer to [**[]METag**](METag.md) | A set of tags assigned to the entity. | [optional] 

## Methods

### NewEntity

`func NewEntity() *Entity`

NewEntity instantiates a new Entity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityWithDefaults

`func NewEntityWithDefaults() *Entity`

NewEntityWithDefaults instantiates a new Entity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagementZones

`func (o *Entity) GetManagementZones() []ManagementZone`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *Entity) GetManagementZonesOk() (*[]ManagementZone, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *Entity) SetManagementZones(v []ManagementZone)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *Entity) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetFromRelationships

`func (o *Entity) GetFromRelationships() map[string][]EntityId`

GetFromRelationships returns the FromRelationships field if non-nil, zero value otherwise.

### GetFromRelationshipsOk

`func (o *Entity) GetFromRelationshipsOk() (*map[string][]EntityId, bool)`

GetFromRelationshipsOk returns a tuple with the FromRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromRelationships

`func (o *Entity) SetFromRelationships(v map[string][]EntityId)`

SetFromRelationships sets FromRelationships field to given value.

### HasFromRelationships

`func (o *Entity) HasFromRelationships() bool`

HasFromRelationships returns a boolean if a field has been set.

### GetToRelationships

`func (o *Entity) GetToRelationships() map[string][]EntityId`

GetToRelationships returns the ToRelationships field if non-nil, zero value otherwise.

### GetToRelationshipsOk

`func (o *Entity) GetToRelationshipsOk() (*map[string][]EntityId, bool)`

GetToRelationshipsOk returns a tuple with the ToRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRelationships

`func (o *Entity) SetToRelationships(v map[string][]EntityId)`

SetToRelationships sets ToRelationships field to given value.

### HasToRelationships

`func (o *Entity) HasToRelationships() bool`

HasToRelationships returns a boolean if a field has been set.

### GetEntityId

`func (o *Entity) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Entity) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Entity) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *Entity) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetLastSeenTms

`func (o *Entity) GetLastSeenTms() int64`

GetLastSeenTms returns the LastSeenTms field if non-nil, zero value otherwise.

### GetLastSeenTmsOk

`func (o *Entity) GetLastSeenTmsOk() (*int64, bool)`

GetLastSeenTmsOk returns a tuple with the LastSeenTms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenTms

`func (o *Entity) SetLastSeenTms(v int64)`

SetLastSeenTms sets LastSeenTms field to given value.

### HasLastSeenTms

`func (o *Entity) HasLastSeenTms() bool`

HasLastSeenTms returns a boolean if a field has been set.

### GetFirstSeenTms

`func (o *Entity) GetFirstSeenTms() int64`

GetFirstSeenTms returns the FirstSeenTms field if non-nil, zero value otherwise.

### GetFirstSeenTmsOk

`func (o *Entity) GetFirstSeenTmsOk() (*int64, bool)`

GetFirstSeenTmsOk returns a tuple with the FirstSeenTms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenTms

`func (o *Entity) SetFirstSeenTms(v int64)`

SetFirstSeenTms sets FirstSeenTms field to given value.

### HasFirstSeenTms

`func (o *Entity) HasFirstSeenTms() bool`

HasFirstSeenTms returns a boolean if a field has been set.

### GetIcon

`func (o *Entity) GetIcon() EntityIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Entity) GetIconOk() (*EntityIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Entity) SetIcon(v EntityIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Entity) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetProperties

`func (o *Entity) GetProperties() map[string]map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Entity) GetPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Entity) SetProperties(v map[string]map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Entity) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetDisplayName

`func (o *Entity) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Entity) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Entity) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Entity) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTags

`func (o *Entity) GetTags() []METag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Entity) GetTagsOk() (*[]METag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Entity) SetTags(v []METag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Entity) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


