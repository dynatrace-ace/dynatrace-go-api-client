# EntityStub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to [**EntityId**](EntityId.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the entity.    Not included in the response in case no entity with the relevant ID was found. | [optional] 

## Methods

### NewEntityStub

`func NewEntityStub() *EntityStub`

NewEntityStub instantiates a new EntityStub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityStubWithDefaults

`func NewEntityStubWithDefaults() *EntityStub`

NewEntityStubWithDefaults instantiates a new EntityStub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *EntityStub) GetEntityId() EntityId`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *EntityStub) GetEntityIdOk() (*EntityId, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *EntityStub) SetEntityId(v EntityId)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *EntityStub) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetName

`func (o *EntityStub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityStub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityStub) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntityStub) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


