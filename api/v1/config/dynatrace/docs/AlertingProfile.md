# AlertingProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the alerting profile. | [optional] 
**DisplayName** | **string** | The name of the alerting profile, displayed in the UI. | 
**Rules** | Pointer to [**[]AlertingProfileSeverityRule**](AlertingProfileSeverityRule.md) | A list of severity rules.    The rules are evaluated from top to bottom. The first matching rule applies and further evaluation stops.   If you specify both severity rule and event filter, the AND logic applies. | [optional] 
**MzId** | Pointer to **string** | The ID of the management zone to which the alerting profile applies. | [optional] 
**EventTypeFilters** | Pointer to [**[]AlertingEventTypeFilter**](AlertingEventTypeFilter.md) | The list of event filters.   For all filters that are *negated* inside of these event filters, that is all \&quot;Predefined\&quot; as well as \&quot;Custom\&quot; (Title and/or Description) ones the AND logic applies. For all *non-negated* ones the OR logic applies. Between these two groups, negated and non-negated, the AND logic applies.   If you specify both severity rule and event filter, the AND logic applies. | [optional] 

## Methods

### NewAlertingProfile

`func NewAlertingProfile(displayName string, ) *AlertingProfile`

NewAlertingProfile instantiates a new AlertingProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingProfileWithDefaults

`func NewAlertingProfileWithDefaults() *AlertingProfile`

NewAlertingProfileWithDefaults instantiates a new AlertingProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *AlertingProfile) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AlertingProfile) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AlertingProfile) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AlertingProfile) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *AlertingProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertingProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertingProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AlertingProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *AlertingProfile) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AlertingProfile) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AlertingProfile) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetRules

`func (o *AlertingProfile) GetRules() []AlertingProfileSeverityRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AlertingProfile) GetRulesOk() (*[]AlertingProfileSeverityRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AlertingProfile) SetRules(v []AlertingProfileSeverityRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *AlertingProfile) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetMzId

`func (o *AlertingProfile) GetMzId() string`

GetMzId returns the MzId field if non-nil, zero value otherwise.

### GetMzIdOk

`func (o *AlertingProfile) GetMzIdOk() (*string, bool)`

GetMzIdOk returns a tuple with the MzId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMzId

`func (o *AlertingProfile) SetMzId(v string)`

SetMzId sets MzId field to given value.

### HasMzId

`func (o *AlertingProfile) HasMzId() bool`

HasMzId returns a boolean if a field has been set.

### GetEventTypeFilters

`func (o *AlertingProfile) GetEventTypeFilters() []AlertingEventTypeFilter`

GetEventTypeFilters returns the EventTypeFilters field if non-nil, zero value otherwise.

### GetEventTypeFiltersOk

`func (o *AlertingProfile) GetEventTypeFiltersOk() (*[]AlertingEventTypeFilter, bool)`

GetEventTypeFiltersOk returns a tuple with the EventTypeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeFilters

`func (o *AlertingProfile) SetEventTypeFilters(v []AlertingEventTypeFilter)`

SetEventTypeFilters sets EventTypeFilters field to given value.

### HasEventTypeFilters

`func (o *AlertingProfile) HasEventTypeFilters() bool`

HasEventTypeFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


