# CredentialsResponseElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the credentials set. | 
**Id** | Pointer to **string** | The ID of the credentials set. | [optional] 
**Type** | **string** | The type of the credentials set. | 
**Description** | **string** | A short description of the credentials set. | 
**Owner** | **string** | The owner of the credential (user for which used API token was created). | 
**OwnerAccessOnly** | **bool** | Flag indicating that this credential is visible only to the owner. | 

## Methods

### NewCredentialsResponseElement

`func NewCredentialsResponseElement(name string, type_ string, description string, owner string, ownerAccessOnly bool, ) *CredentialsResponseElement`

NewCredentialsResponseElement instantiates a new CredentialsResponseElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsResponseElementWithDefaults

`func NewCredentialsResponseElementWithDefaults() *CredentialsResponseElement`

NewCredentialsResponseElementWithDefaults instantiates a new CredentialsResponseElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CredentialsResponseElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialsResponseElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialsResponseElement) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *CredentialsResponseElement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialsResponseElement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialsResponseElement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CredentialsResponseElement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CredentialsResponseElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialsResponseElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialsResponseElement) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *CredentialsResponseElement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialsResponseElement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialsResponseElement) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOwner

`func (o *CredentialsResponseElement) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CredentialsResponseElement) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CredentialsResponseElement) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetOwnerAccessOnly

`func (o *CredentialsResponseElement) GetOwnerAccessOnly() bool`

GetOwnerAccessOnly returns the OwnerAccessOnly field if non-nil, zero value otherwise.

### GetOwnerAccessOnlyOk

`func (o *CredentialsResponseElement) GetOwnerAccessOnlyOk() (*bool, bool)`

GetOwnerAccessOnlyOk returns a tuple with the OwnerAccessOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAccessOnly

`func (o *CredentialsResponseElement) SetOwnerAccessOnly(v bool)`

SetOwnerAccessOnly sets OwnerAccessOnly field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


