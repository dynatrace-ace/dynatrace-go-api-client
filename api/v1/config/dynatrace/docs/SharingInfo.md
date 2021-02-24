# SharingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkShared** | Pointer to **bool** | If &#x60;true&#x60;, the dashboard is shared via link and authenticated users with the link can view. | [optional] 
**Published** | Pointer to **bool** | If &#x60;true&#x60;, the dashboard is published to anyone on this environment. | [optional] 

## Methods

### NewSharingInfo

`func NewSharingInfo() *SharingInfo`

NewSharingInfo instantiates a new SharingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharingInfoWithDefaults

`func NewSharingInfoWithDefaults() *SharingInfo`

NewSharingInfoWithDefaults instantiates a new SharingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkShared

`func (o *SharingInfo) GetLinkShared() bool`

GetLinkShared returns the LinkShared field if non-nil, zero value otherwise.

### GetLinkSharedOk

`func (o *SharingInfo) GetLinkSharedOk() (*bool, bool)`

GetLinkSharedOk returns a tuple with the LinkShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkShared

`func (o *SharingInfo) SetLinkShared(v bool)`

SetLinkShared sets LinkShared field to given value.

### HasLinkShared

`func (o *SharingInfo) HasLinkShared() bool`

HasLinkShared returns a boolean if a field has been set.

### GetPublished

`func (o *SharingInfo) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *SharingInfo) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *SharingInfo) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *SharingInfo) HasPublished() bool`

HasPublished returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


