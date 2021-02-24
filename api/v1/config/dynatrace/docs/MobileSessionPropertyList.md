# MobileSessionPropertyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionProperties** | Pointer to [**[]MobileSessionPropertyShort**](MobileSessionPropertyShort.md) | A list of short representations of mobile session properties. | [optional] 

## Methods

### NewMobileSessionPropertyList

`func NewMobileSessionPropertyList() *MobileSessionPropertyList`

NewMobileSessionPropertyList instantiates a new MobileSessionPropertyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileSessionPropertyListWithDefaults

`func NewMobileSessionPropertyListWithDefaults() *MobileSessionPropertyList`

NewMobileSessionPropertyListWithDefaults instantiates a new MobileSessionPropertyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionProperties

`func (o *MobileSessionPropertyList) GetSessionProperties() []MobileSessionPropertyShort`

GetSessionProperties returns the SessionProperties field if non-nil, zero value otherwise.

### GetSessionPropertiesOk

`func (o *MobileSessionPropertyList) GetSessionPropertiesOk() (*[]MobileSessionPropertyShort, bool)`

GetSessionPropertiesOk returns a tuple with the SessionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionProperties

`func (o *MobileSessionPropertyList) SetSessionProperties(v []MobileSessionPropertyShort)`

SetSessionProperties sets SessionProperties field to given value.

### HasSessionProperties

`func (o *MobileSessionPropertyList) HasSessionProperties() bool`

HasSessionProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


