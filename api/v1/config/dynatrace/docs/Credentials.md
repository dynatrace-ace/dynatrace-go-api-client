# Credentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the credentials set. | 
**Id** | Pointer to **string** | The ID of the credentials set. | [optional] 
**Description** | Pointer to **string** | A short description of the credentials set. | [optional] 
**OwnerAccessOnly** | Pointer to **bool** | The credentials set is available to every user (&#x60;false&#x60;) or to owner only (&#x60;true&#x60;). | [optional] 
**Type** | Pointer to **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;CERTIFICATE&#x60; -&gt; CertificateCredentials  * &#x60;USERNAME_PASSWORD&#x60; -&gt; UserPasswordCredentials  * &#x60;TOKEN&#x60; -&gt; TokenCredentials   | [optional] 

## Methods

### NewCredentials

`func NewCredentials(name string, ) *Credentials`

NewCredentials instantiates a new Credentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsWithDefaults

`func NewCredentialsWithDefaults() *Credentials`

NewCredentialsWithDefaults instantiates a new Credentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Credentials) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Credentials) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Credentials) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *Credentials) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Credentials) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Credentials) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Credentials) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *Credentials) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Credentials) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Credentials) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Credentials) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwnerAccessOnly

`func (o *Credentials) GetOwnerAccessOnly() bool`

GetOwnerAccessOnly returns the OwnerAccessOnly field if non-nil, zero value otherwise.

### GetOwnerAccessOnlyOk

`func (o *Credentials) GetOwnerAccessOnlyOk() (*bool, bool)`

GetOwnerAccessOnlyOk returns a tuple with the OwnerAccessOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAccessOnly

`func (o *Credentials) SetOwnerAccessOnly(v bool)`

SetOwnerAccessOnly sets OwnerAccessOnly field to given value.

### HasOwnerAccessOnly

`func (o *Credentials) HasOwnerAccessOnly() bool`

HasOwnerAccessOnly returns a boolean if a field has been set.

### GetType

`func (o *Credentials) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Credentials) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Credentials) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Credentials) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


