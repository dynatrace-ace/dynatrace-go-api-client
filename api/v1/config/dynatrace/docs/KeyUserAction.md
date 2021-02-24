# KeyUserAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the action. | 
**ActionType** | **string** | The type of the action. | 
**Domain** | Pointer to **string** | The domain where the action is performed. | [optional] 
**MeIdentifier** | Pointer to **string** | The Dynatrace entity ID of the action. | [optional] [readonly] 

## Methods

### NewKeyUserAction

`func NewKeyUserAction(name string, actionType string, ) *KeyUserAction`

NewKeyUserAction instantiates a new KeyUserAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyUserActionWithDefaults

`func NewKeyUserActionWithDefaults() *KeyUserAction`

NewKeyUserActionWithDefaults instantiates a new KeyUserAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyUserAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyUserAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyUserAction) SetName(v string)`

SetName sets Name field to given value.


### GetActionType

`func (o *KeyUserAction) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *KeyUserAction) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *KeyUserAction) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetDomain

`func (o *KeyUserAction) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *KeyUserAction) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *KeyUserAction) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *KeyUserAction) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetMeIdentifier

`func (o *KeyUserAction) GetMeIdentifier() string`

GetMeIdentifier returns the MeIdentifier field if non-nil, zero value otherwise.

### GetMeIdentifierOk

`func (o *KeyUserAction) GetMeIdentifierOk() (*string, bool)`

GetMeIdentifierOk returns a tuple with the MeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeIdentifier

`func (o *KeyUserAction) SetMeIdentifier(v string)`

SetMeIdentifier sets MeIdentifier field to given value.

### HasMeIdentifier

`func (o *KeyUserAction) HasMeIdentifier() bool`

HasMeIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


