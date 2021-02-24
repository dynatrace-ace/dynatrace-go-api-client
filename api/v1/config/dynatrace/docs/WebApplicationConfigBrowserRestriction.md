# WebApplicationConfigBrowserRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrowserVersion** | Pointer to **string** | The version of the browser that is used. | [optional] 
**BrowserType** | **string** | The type of the browser that is used. | 
**Platform** | Pointer to **string** | The platform on which the browser is being used. | [optional] 
**Comparator** | Pointer to **string** | Compares different browsers together. | [optional] 

## Methods

### NewWebApplicationConfigBrowserRestriction

`func NewWebApplicationConfigBrowserRestriction(browserType string, ) *WebApplicationConfigBrowserRestriction`

NewWebApplicationConfigBrowserRestriction instantiates a new WebApplicationConfigBrowserRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebApplicationConfigBrowserRestrictionWithDefaults

`func NewWebApplicationConfigBrowserRestrictionWithDefaults() *WebApplicationConfigBrowserRestriction`

NewWebApplicationConfigBrowserRestrictionWithDefaults instantiates a new WebApplicationConfigBrowserRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowserVersion

`func (o *WebApplicationConfigBrowserRestriction) GetBrowserVersion() string`

GetBrowserVersion returns the BrowserVersion field if non-nil, zero value otherwise.

### GetBrowserVersionOk

`func (o *WebApplicationConfigBrowserRestriction) GetBrowserVersionOk() (*string, bool)`

GetBrowserVersionOk returns a tuple with the BrowserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserVersion

`func (o *WebApplicationConfigBrowserRestriction) SetBrowserVersion(v string)`

SetBrowserVersion sets BrowserVersion field to given value.

### HasBrowserVersion

`func (o *WebApplicationConfigBrowserRestriction) HasBrowserVersion() bool`

HasBrowserVersion returns a boolean if a field has been set.

### GetBrowserType

`func (o *WebApplicationConfigBrowserRestriction) GetBrowserType() string`

GetBrowserType returns the BrowserType field if non-nil, zero value otherwise.

### GetBrowserTypeOk

`func (o *WebApplicationConfigBrowserRestriction) GetBrowserTypeOk() (*string, bool)`

GetBrowserTypeOk returns a tuple with the BrowserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserType

`func (o *WebApplicationConfigBrowserRestriction) SetBrowserType(v string)`

SetBrowserType sets BrowserType field to given value.


### GetPlatform

`func (o *WebApplicationConfigBrowserRestriction) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *WebApplicationConfigBrowserRestriction) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *WebApplicationConfigBrowserRestriction) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *WebApplicationConfigBrowserRestriction) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetComparator

`func (o *WebApplicationConfigBrowserRestriction) GetComparator() string`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *WebApplicationConfigBrowserRestriction) GetComparatorOk() (*string, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *WebApplicationConfigBrowserRestriction) SetComparator(v string)`

SetComparator sets Comparator field to given value.

### HasComparator

`func (o *WebApplicationConfigBrowserRestriction) HasComparator() bool`

HasComparator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


