# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferencedType** | Pointer to **string** | The type referenced by the item&#39;s value. | [optional] 
**Documentation** | Pointer to **string** | An extended description and/or links to documentation. | [optional] 
**SubType** | Pointer to **string** | The subtype of the item&#39;s value. | [optional] 
**Type** | Pointer to **map[string]interface{}** | The type of the item&#39;s value. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the item. | [optional] 
**Description** | Pointer to **string** | A short description of the item. | [optional] 
**Constraints** | Pointer to [**[]Constraint**](Constraint.md) | A list of constraints limiting the values to be accepted. | [optional] 

## Methods

### NewItem

`func NewItem() *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferencedType

`func (o *Item) GetReferencedType() string`

GetReferencedType returns the ReferencedType field if non-nil, zero value otherwise.

### GetReferencedTypeOk

`func (o *Item) GetReferencedTypeOk() (*string, bool)`

GetReferencedTypeOk returns a tuple with the ReferencedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedType

`func (o *Item) SetReferencedType(v string)`

SetReferencedType sets ReferencedType field to given value.

### HasReferencedType

`func (o *Item) HasReferencedType() bool`

HasReferencedType returns a boolean if a field has been set.

### GetDocumentation

`func (o *Item) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *Item) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *Item) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *Item) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetSubType

`func (o *Item) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *Item) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *Item) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *Item) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetType

`func (o *Item) GetType() map[string]interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Item) GetTypeOk() (*map[string]interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Item) SetType(v map[string]interface{})`

SetType sets Type field to given value.

### HasType

`func (o *Item) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayName

`func (o *Item) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Item) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Item) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Item) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *Item) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Item) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Item) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Item) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConstraints

`func (o *Item) GetConstraints() []Constraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *Item) GetConstraintsOk() (*[]Constraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *Item) SetConstraints(v []Constraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *Item) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


