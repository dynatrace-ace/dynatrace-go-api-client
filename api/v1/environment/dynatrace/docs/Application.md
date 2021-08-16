# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** | The Dynatrace entity ID of the required entity. | [optional] 
**DisplayName** | Pointer to **string** | The name of the Dynatrace entity as displayed in the UI. | [optional] 
**CustomizedName** | Pointer to **string** | The customized name of the entity | [optional] 
**DiscoveredName** | Pointer to **string** | The discovered name of the entity | [optional] 
**FirstSeenTimestamp** | Pointer to **int64** | The timestamp of when the entity was first detected, in UTC milliseconds | [optional] 
**LastSeenTimestamp** | Pointer to **int64** | The timestamp of when the entity was last detected, in UTC milliseconds | [optional] 
**Tags** | Pointer to [**[]TagInfo**](TagInfo.md) | The list of entity tags. | [optional] 
**FromRelationships** | Pointer to [**ApplicationFromRelationships**](ApplicationFromRelationships.md) |  | [optional] 
**ToRelationships** | Pointer to [**ApplicationToRelationships**](ApplicationToRelationships.md) |  | [optional] 
**ApplicationType** | Pointer to **string** |  | [optional] 
**RuleAppliedMatchType** | Pointer to **string** |  | [optional] 
**RuleAppliedPattern** | Pointer to **string** |  | [optional] 
**ApplicationMatchTarget** | Pointer to **string** |  | [optional] 
**ManagementZones** | Pointer to [**[]EntityShortRepresentation**](EntityShortRepresentation.md) | The management zones that the entity is part of. | [optional] 

## Methods

### NewApplication

`func NewApplication() *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *Application) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Application) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Application) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *Application) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetDisplayName

`func (o *Application) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Application) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Application) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Application) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCustomizedName

`func (o *Application) GetCustomizedName() string`

GetCustomizedName returns the CustomizedName field if non-nil, zero value otherwise.

### GetCustomizedNameOk

`func (o *Application) GetCustomizedNameOk() (*string, bool)`

GetCustomizedNameOk returns a tuple with the CustomizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizedName

`func (o *Application) SetCustomizedName(v string)`

SetCustomizedName sets CustomizedName field to given value.

### HasCustomizedName

`func (o *Application) HasCustomizedName() bool`

HasCustomizedName returns a boolean if a field has been set.

### GetDiscoveredName

`func (o *Application) GetDiscoveredName() string`

GetDiscoveredName returns the DiscoveredName field if non-nil, zero value otherwise.

### GetDiscoveredNameOk

`func (o *Application) GetDiscoveredNameOk() (*string, bool)`

GetDiscoveredNameOk returns a tuple with the DiscoveredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredName

`func (o *Application) SetDiscoveredName(v string)`

SetDiscoveredName sets DiscoveredName field to given value.

### HasDiscoveredName

`func (o *Application) HasDiscoveredName() bool`

HasDiscoveredName returns a boolean if a field has been set.

### GetFirstSeenTimestamp

`func (o *Application) GetFirstSeenTimestamp() int64`

GetFirstSeenTimestamp returns the FirstSeenTimestamp field if non-nil, zero value otherwise.

### GetFirstSeenTimestampOk

`func (o *Application) GetFirstSeenTimestampOk() (*int64, bool)`

GetFirstSeenTimestampOk returns a tuple with the FirstSeenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenTimestamp

`func (o *Application) SetFirstSeenTimestamp(v int64)`

SetFirstSeenTimestamp sets FirstSeenTimestamp field to given value.

### HasFirstSeenTimestamp

`func (o *Application) HasFirstSeenTimestamp() bool`

HasFirstSeenTimestamp returns a boolean if a field has been set.

### GetLastSeenTimestamp

`func (o *Application) GetLastSeenTimestamp() int64`

GetLastSeenTimestamp returns the LastSeenTimestamp field if non-nil, zero value otherwise.

### GetLastSeenTimestampOk

`func (o *Application) GetLastSeenTimestampOk() (*int64, bool)`

GetLastSeenTimestampOk returns a tuple with the LastSeenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenTimestamp

`func (o *Application) SetLastSeenTimestamp(v int64)`

SetLastSeenTimestamp sets LastSeenTimestamp field to given value.

### HasLastSeenTimestamp

`func (o *Application) HasLastSeenTimestamp() bool`

HasLastSeenTimestamp returns a boolean if a field has been set.

### GetTags

`func (o *Application) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Application) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Application) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Application) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFromRelationships

`func (o *Application) GetFromRelationships() ApplicationFromRelationships`

GetFromRelationships returns the FromRelationships field if non-nil, zero value otherwise.

### GetFromRelationshipsOk

`func (o *Application) GetFromRelationshipsOk() (*ApplicationFromRelationships, bool)`

GetFromRelationshipsOk returns a tuple with the FromRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromRelationships

`func (o *Application) SetFromRelationships(v ApplicationFromRelationships)`

SetFromRelationships sets FromRelationships field to given value.

### HasFromRelationships

`func (o *Application) HasFromRelationships() bool`

HasFromRelationships returns a boolean if a field has been set.

### GetToRelationships

`func (o *Application) GetToRelationships() ApplicationToRelationships`

GetToRelationships returns the ToRelationships field if non-nil, zero value otherwise.

### GetToRelationshipsOk

`func (o *Application) GetToRelationshipsOk() (*ApplicationToRelationships, bool)`

GetToRelationshipsOk returns a tuple with the ToRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRelationships

`func (o *Application) SetToRelationships(v ApplicationToRelationships)`

SetToRelationships sets ToRelationships field to given value.

### HasToRelationships

`func (o *Application) HasToRelationships() bool`

HasToRelationships returns a boolean if a field has been set.

### GetApplicationType

`func (o *Application) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *Application) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *Application) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *Application) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetRuleAppliedMatchType

`func (o *Application) GetRuleAppliedMatchType() string`

GetRuleAppliedMatchType returns the RuleAppliedMatchType field if non-nil, zero value otherwise.

### GetRuleAppliedMatchTypeOk

`func (o *Application) GetRuleAppliedMatchTypeOk() (*string, bool)`

GetRuleAppliedMatchTypeOk returns a tuple with the RuleAppliedMatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleAppliedMatchType

`func (o *Application) SetRuleAppliedMatchType(v string)`

SetRuleAppliedMatchType sets RuleAppliedMatchType field to given value.

### HasRuleAppliedMatchType

`func (o *Application) HasRuleAppliedMatchType() bool`

HasRuleAppliedMatchType returns a boolean if a field has been set.

### GetRuleAppliedPattern

`func (o *Application) GetRuleAppliedPattern() string`

GetRuleAppliedPattern returns the RuleAppliedPattern field if non-nil, zero value otherwise.

### GetRuleAppliedPatternOk

`func (o *Application) GetRuleAppliedPatternOk() (*string, bool)`

GetRuleAppliedPatternOk returns a tuple with the RuleAppliedPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleAppliedPattern

`func (o *Application) SetRuleAppliedPattern(v string)`

SetRuleAppliedPattern sets RuleAppliedPattern field to given value.

### HasRuleAppliedPattern

`func (o *Application) HasRuleAppliedPattern() bool`

HasRuleAppliedPattern returns a boolean if a field has been set.

### GetApplicationMatchTarget

`func (o *Application) GetApplicationMatchTarget() string`

GetApplicationMatchTarget returns the ApplicationMatchTarget field if non-nil, zero value otherwise.

### GetApplicationMatchTargetOk

`func (o *Application) GetApplicationMatchTargetOk() (*string, bool)`

GetApplicationMatchTargetOk returns a tuple with the ApplicationMatchTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationMatchTarget

`func (o *Application) SetApplicationMatchTarget(v string)`

SetApplicationMatchTarget sets ApplicationMatchTarget field to given value.

### HasApplicationMatchTarget

`func (o *Application) HasApplicationMatchTarget() bool`

HasApplicationMatchTarget returns a boolean if a field has been set.

### GetManagementZones

`func (o *Application) GetManagementZones() []EntityShortRepresentation`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *Application) GetManagementZonesOk() (*[]EntityShortRepresentation, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *Application) SetManagementZones(v []EntityShortRepresentation)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *Application) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


