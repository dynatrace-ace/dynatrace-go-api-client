# EntityIcon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryIconType** | Pointer to **string** | The primary icon of the entity.   Specified by the [barista](https://dt-url.net/u403suy) ID of the icon. | [optional] [readonly] 
**SecondaryIconType** | Pointer to **string** | The secondary icon of the entity.   Specified by the [barista](https://dt-url.net/u403suy) ID of the icon. | [optional] [readonly] 
**CustomIconPath** | Pointer to **string** | The user-defined icon of the entity.   Specify the [barista](https://dt-url.net/u403suy) ID of the icon or a URL of your own icon. | [optional] 

## Methods

### NewEntityIcon

`func NewEntityIcon() *EntityIcon`

NewEntityIcon instantiates a new EntityIcon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityIconWithDefaults

`func NewEntityIconWithDefaults() *EntityIcon`

NewEntityIconWithDefaults instantiates a new EntityIcon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryIconType

`func (o *EntityIcon) GetPrimaryIconType() string`

GetPrimaryIconType returns the PrimaryIconType field if non-nil, zero value otherwise.

### GetPrimaryIconTypeOk

`func (o *EntityIcon) GetPrimaryIconTypeOk() (*string, bool)`

GetPrimaryIconTypeOk returns a tuple with the PrimaryIconType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIconType

`func (o *EntityIcon) SetPrimaryIconType(v string)`

SetPrimaryIconType sets PrimaryIconType field to given value.

### HasPrimaryIconType

`func (o *EntityIcon) HasPrimaryIconType() bool`

HasPrimaryIconType returns a boolean if a field has been set.

### GetSecondaryIconType

`func (o *EntityIcon) GetSecondaryIconType() string`

GetSecondaryIconType returns the SecondaryIconType field if non-nil, zero value otherwise.

### GetSecondaryIconTypeOk

`func (o *EntityIcon) GetSecondaryIconTypeOk() (*string, bool)`

GetSecondaryIconTypeOk returns a tuple with the SecondaryIconType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIconType

`func (o *EntityIcon) SetSecondaryIconType(v string)`

SetSecondaryIconType sets SecondaryIconType field to given value.

### HasSecondaryIconType

`func (o *EntityIcon) HasSecondaryIconType() bool`

HasSecondaryIconType returns a boolean if a field has been set.

### GetCustomIconPath

`func (o *EntityIcon) GetCustomIconPath() string`

GetCustomIconPath returns the CustomIconPath field if non-nil, zero value otherwise.

### GetCustomIconPathOk

`func (o *EntityIcon) GetCustomIconPathOk() (*string, bool)`

GetCustomIconPathOk returns a tuple with the CustomIconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomIconPath

`func (o *EntityIcon) SetCustomIconPath(v string)`

SetCustomIconPath sets CustomIconPath field to given value.

### HasCustomIconPath

`func (o *EntityIcon) HasCustomIconPath() bool`

HasCustomIconPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


