# AlertingProfileTagFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeMode** | **string** | The filtering mode:  * &#x60;INCLUDE_ANY&#x60;: The rule applies to monitored entities that have at least one of the specified tags. You can specify up to 100 tags.  * &#x60;INCLUDE_ALL&#x60;: The rule applies to monitored entities that have **all** of the specified tags. You can specify up to 10 tags.  * &#x60;NONE&#x60;: The rule applies to all monitored entities. | 
**TagFilters** | Pointer to [**[]TagFilter**](TagFilter.md) | A list of required tags. | [optional] 

## Methods

### NewAlertingProfileTagFilter

`func NewAlertingProfileTagFilter(includeMode string, ) *AlertingProfileTagFilter`

NewAlertingProfileTagFilter instantiates a new AlertingProfileTagFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingProfileTagFilterWithDefaults

`func NewAlertingProfileTagFilterWithDefaults() *AlertingProfileTagFilter`

NewAlertingProfileTagFilterWithDefaults instantiates a new AlertingProfileTagFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeMode

`func (o *AlertingProfileTagFilter) GetIncludeMode() string`

GetIncludeMode returns the IncludeMode field if non-nil, zero value otherwise.

### GetIncludeModeOk

`func (o *AlertingProfileTagFilter) GetIncludeModeOk() (*string, bool)`

GetIncludeModeOk returns a tuple with the IncludeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMode

`func (o *AlertingProfileTagFilter) SetIncludeMode(v string)`

SetIncludeMode sets IncludeMode field to given value.


### GetTagFilters

`func (o *AlertingProfileTagFilter) GetTagFilters() []TagFilter`

GetTagFilters returns the TagFilters field if non-nil, zero value otherwise.

### GetTagFiltersOk

`func (o *AlertingProfileTagFilter) GetTagFiltersOk() (*[]TagFilter, bool)`

GetTagFiltersOk returns a tuple with the TagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilters

`func (o *AlertingProfileTagFilter) SetTagFilters(v []TagFilter)`

SetTagFilters sets TagFilters field to given value.

### HasTagFilters

`func (o *AlertingProfileTagFilter) HasTagFilters() bool`

HasTagFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


