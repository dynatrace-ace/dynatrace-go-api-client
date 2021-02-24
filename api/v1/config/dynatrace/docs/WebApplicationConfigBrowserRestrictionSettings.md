# WebApplicationConfigBrowserRestrictionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | The mode of the list of browser restrictions. | 
**BrowserRestrictions** | Pointer to [**[]WebApplicationConfigBrowserRestriction**](WebApplicationConfigBrowserRestriction.md) | A list of browser restrictions. | [optional] 

## Methods

### NewWebApplicationConfigBrowserRestrictionSettings

`func NewWebApplicationConfigBrowserRestrictionSettings(mode string, ) *WebApplicationConfigBrowserRestrictionSettings`

NewWebApplicationConfigBrowserRestrictionSettings instantiates a new WebApplicationConfigBrowserRestrictionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebApplicationConfigBrowserRestrictionSettingsWithDefaults

`func NewWebApplicationConfigBrowserRestrictionSettingsWithDefaults() *WebApplicationConfigBrowserRestrictionSettings`

NewWebApplicationConfigBrowserRestrictionSettingsWithDefaults instantiates a new WebApplicationConfigBrowserRestrictionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *WebApplicationConfigBrowserRestrictionSettings) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *WebApplicationConfigBrowserRestrictionSettings) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *WebApplicationConfigBrowserRestrictionSettings) SetMode(v string)`

SetMode sets Mode field to given value.


### GetBrowserRestrictions

`func (o *WebApplicationConfigBrowserRestrictionSettings) GetBrowserRestrictions() []WebApplicationConfigBrowserRestriction`

GetBrowserRestrictions returns the BrowserRestrictions field if non-nil, zero value otherwise.

### GetBrowserRestrictionsOk

`func (o *WebApplicationConfigBrowserRestrictionSettings) GetBrowserRestrictionsOk() (*[]WebApplicationConfigBrowserRestriction, bool)`

GetBrowserRestrictionsOk returns a tuple with the BrowserRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserRestrictions

`func (o *WebApplicationConfigBrowserRestrictionSettings) SetBrowserRestrictions(v []WebApplicationConfigBrowserRestriction)`

SetBrowserRestrictions sets BrowserRestrictions field to given value.

### HasBrowserRestrictions

`func (o *WebApplicationConfigBrowserRestrictionSettings) HasBrowserRestrictions() bool`

HasBrowserRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


