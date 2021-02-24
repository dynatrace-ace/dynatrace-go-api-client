# ApplicationDataPrivacyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Values** | Pointer to [**[]ApplicationDataPrivacy**](ApplicationDataPrivacy.md) |  | [optional] 

## Methods

### NewApplicationDataPrivacyList

`func NewApplicationDataPrivacyList() *ApplicationDataPrivacyList`

NewApplicationDataPrivacyList instantiates a new ApplicationDataPrivacyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDataPrivacyListWithDefaults

`func NewApplicationDataPrivacyListWithDefaults() *ApplicationDataPrivacyList`

NewApplicationDataPrivacyListWithDefaults instantiates a new ApplicationDataPrivacyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ApplicationDataPrivacyList) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationDataPrivacyList) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationDataPrivacyList) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationDataPrivacyList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetValues

`func (o *ApplicationDataPrivacyList) GetValues() []ApplicationDataPrivacy`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ApplicationDataPrivacyList) GetValuesOk() (*[]ApplicationDataPrivacy, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ApplicationDataPrivacyList) SetValues(v []ApplicationDataPrivacy)`

SetValues sets Values field to given value.

### HasValues

`func (o *ApplicationDataPrivacyList) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


