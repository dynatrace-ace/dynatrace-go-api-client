# ApplicationToRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitors** | Pointer to **[]string** | Represents relationship between a synthetic monitor and an application | [optional] 

## Methods

### NewApplicationToRelationships

`func NewApplicationToRelationships() *ApplicationToRelationships`

NewApplicationToRelationships instantiates a new ApplicationToRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationToRelationshipsWithDefaults

`func NewApplicationToRelationshipsWithDefaults() *ApplicationToRelationships`

NewApplicationToRelationshipsWithDefaults instantiates a new ApplicationToRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitors

`func (o *ApplicationToRelationships) GetMonitors() []string`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *ApplicationToRelationships) GetMonitorsOk() (*[]string, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *ApplicationToRelationships) SetMonitors(v []string)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *ApplicationToRelationships) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


