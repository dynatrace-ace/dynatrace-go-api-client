# DavisSecurityAdvice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the advice. | [optional] [readonly] 
**VulnerableComponent** | Pointer to **string** | The vulnerable component to which advice applies. | [optional] [readonly] 
**Technology** | Pointer to **string** | The technology of the vulnerable component. | [optional] [readonly] 
**AdviceType** | Pointer to **string** | The type of the advice. | [optional] [readonly] 
**Critical** | Pointer to **[]string** | IDs of &#x60;critical&#x60; level [security problems](https://dt-url.net/p103u1h) caused by vulnerable component. | [optional] [readonly] 
**High** | Pointer to **[]string** | IDs of &#x60;high&#x60; level [security problems](https://dt-url.net/p103u1h) caused by vulnerable component. | [optional] [readonly] 
**Medium** | Pointer to **[]string** | IDs of &#x60;medium&#x60; level [security problems](https://dt-url.net/p103u1h) caused by vulnerable component. | [optional] [readonly] 
**Low** | Pointer to **[]string** | IDs of &#x60;low&#x60; level [security problems](https://dt-url.net/p103u1h) caused by vulnerable component. | [optional] [readonly] 
**None** | Pointer to **[]string** | IDs of &#x60;none&#x60; level [security problems](https://dt-url.net/p103u1h) caused by vulnerable component. | [optional] [readonly] 

## Methods

### NewDavisSecurityAdvice

`func NewDavisSecurityAdvice() *DavisSecurityAdvice`

NewDavisSecurityAdvice instantiates a new DavisSecurityAdvice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDavisSecurityAdviceWithDefaults

`func NewDavisSecurityAdviceWithDefaults() *DavisSecurityAdvice`

NewDavisSecurityAdviceWithDefaults instantiates a new DavisSecurityAdvice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DavisSecurityAdvice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DavisSecurityAdvice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DavisSecurityAdvice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DavisSecurityAdvice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVulnerableComponent

`func (o *DavisSecurityAdvice) GetVulnerableComponent() string`

GetVulnerableComponent returns the VulnerableComponent field if non-nil, zero value otherwise.

### GetVulnerableComponentOk

`func (o *DavisSecurityAdvice) GetVulnerableComponentOk() (*string, bool)`

GetVulnerableComponentOk returns a tuple with the VulnerableComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableComponent

`func (o *DavisSecurityAdvice) SetVulnerableComponent(v string)`

SetVulnerableComponent sets VulnerableComponent field to given value.

### HasVulnerableComponent

`func (o *DavisSecurityAdvice) HasVulnerableComponent() bool`

HasVulnerableComponent returns a boolean if a field has been set.

### GetTechnology

`func (o *DavisSecurityAdvice) GetTechnology() string`

GetTechnology returns the Technology field if non-nil, zero value otherwise.

### GetTechnologyOk

`func (o *DavisSecurityAdvice) GetTechnologyOk() (*string, bool)`

GetTechnologyOk returns a tuple with the Technology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnology

`func (o *DavisSecurityAdvice) SetTechnology(v string)`

SetTechnology sets Technology field to given value.

### HasTechnology

`func (o *DavisSecurityAdvice) HasTechnology() bool`

HasTechnology returns a boolean if a field has been set.

### GetAdviceType

`func (o *DavisSecurityAdvice) GetAdviceType() string`

GetAdviceType returns the AdviceType field if non-nil, zero value otherwise.

### GetAdviceTypeOk

`func (o *DavisSecurityAdvice) GetAdviceTypeOk() (*string, bool)`

GetAdviceTypeOk returns a tuple with the AdviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdviceType

`func (o *DavisSecurityAdvice) SetAdviceType(v string)`

SetAdviceType sets AdviceType field to given value.

### HasAdviceType

`func (o *DavisSecurityAdvice) HasAdviceType() bool`

HasAdviceType returns a boolean if a field has been set.

### GetCritical

`func (o *DavisSecurityAdvice) GetCritical() []string`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *DavisSecurityAdvice) GetCriticalOk() (*[]string, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *DavisSecurityAdvice) SetCritical(v []string)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *DavisSecurityAdvice) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetHigh

`func (o *DavisSecurityAdvice) GetHigh() []string`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *DavisSecurityAdvice) GetHighOk() (*[]string, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *DavisSecurityAdvice) SetHigh(v []string)`

SetHigh sets High field to given value.

### HasHigh

`func (o *DavisSecurityAdvice) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetMedium

`func (o *DavisSecurityAdvice) GetMedium() []string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *DavisSecurityAdvice) GetMediumOk() (*[]string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *DavisSecurityAdvice) SetMedium(v []string)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *DavisSecurityAdvice) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetLow

`func (o *DavisSecurityAdvice) GetLow() []string`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *DavisSecurityAdvice) GetLowOk() (*[]string, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *DavisSecurityAdvice) SetLow(v []string)`

SetLow sets Low field to given value.

### HasLow

`func (o *DavisSecurityAdvice) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetNone

`func (o *DavisSecurityAdvice) GetNone() []string`

GetNone returns the None field if non-nil, zero value otherwise.

### GetNoneOk

`func (o *DavisSecurityAdvice) GetNoneOk() (*[]string, bool)`

GetNoneOk returns a tuple with the None field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNone

`func (o *DavisSecurityAdvice) SetNone(v []string)`

SetNone sets None field to given value.

### HasNone

`func (o *DavisSecurityAdvice) HasNone() bool`

HasNone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


