# Scope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | **[]string** | A list of Dynatrace entities (for example, hosts or services) to be included in the scope.   Allowed values are Dynatrace entity IDs. | 
**Matches** | [**[]MonitoredEntityFilter**](MonitoredEntityFilter.md) | A list of matching rules for dynamic scope formation.   If several rules are set, the OR logic applies. | 

## Methods

### NewScope

`func NewScope(entities []string, matches []MonitoredEntityFilter, ) *Scope`

NewScope instantiates a new Scope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeWithDefaults

`func NewScopeWithDefaults() *Scope`

NewScopeWithDefaults instantiates a new Scope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *Scope) GetEntities() []string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *Scope) GetEntitiesOk() (*[]string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *Scope) SetEntities(v []string)`

SetEntities sets Entities field to given value.


### GetMatches

`func (o *Scope) GetMatches() []MonitoredEntityFilter`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *Scope) GetMatchesOk() (*[]MonitoredEntityFilter, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *Scope) SetMatches(v []MonitoredEntityFilter)`

SetMatches sets Matches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


