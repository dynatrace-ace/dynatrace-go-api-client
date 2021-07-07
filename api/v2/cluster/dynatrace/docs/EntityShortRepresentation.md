# EntityShortRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the Dynatrace entity. | 
**Name** | Pointer to **string** | The name of the Dynatrace entity. | [optional] [readonly] 
**Description** | Pointer to **string** | A short description of the Dynatrace entity. | [optional] [readonly] 

## Methods

### NewEntityShortRepresentation

`func NewEntityShortRepresentation(id string, ) *EntityShortRepresentation`

NewEntityShortRepresentation instantiates a new EntityShortRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityShortRepresentationWithDefaults

`func NewEntityShortRepresentationWithDefaults() *EntityShortRepresentation`

NewEntityShortRepresentationWithDefaults instantiates a new EntityShortRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityShortRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityShortRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityShortRepresentation) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EntityShortRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityShortRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityShortRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntityShortRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EntityShortRepresentation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityShortRepresentation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityShortRepresentation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityShortRepresentation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


