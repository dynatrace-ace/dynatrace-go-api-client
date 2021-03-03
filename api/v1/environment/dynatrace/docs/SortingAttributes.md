# SortingAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SortAscending** | Pointer to **bool** | Sort ascending (&#x60;true&#x60;) or descending (&#x60;false&#x60;). | [optional] 
**SortFieldName** | Pointer to **string** | The field to sort by. | [optional] 

## Methods

### NewSortingAttributes

`func NewSortingAttributes() *SortingAttributes`

NewSortingAttributes instantiates a new SortingAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortingAttributesWithDefaults

`func NewSortingAttributesWithDefaults() *SortingAttributes`

NewSortingAttributesWithDefaults instantiates a new SortingAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSortAscending

`func (o *SortingAttributes) GetSortAscending() bool`

GetSortAscending returns the SortAscending field if non-nil, zero value otherwise.

### GetSortAscendingOk

`func (o *SortingAttributes) GetSortAscendingOk() (*bool, bool)`

GetSortAscendingOk returns a tuple with the SortAscending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAscending

`func (o *SortingAttributes) SetSortAscending(v bool)`

SetSortAscending sets SortAscending field to given value.

### HasSortAscending

`func (o *SortingAttributes) HasSortAscending() bool`

HasSortAscending returns a boolean if a field has been set.

### GetSortFieldName

`func (o *SortingAttributes) GetSortFieldName() string`

GetSortFieldName returns the SortFieldName field if non-nil, zero value otherwise.

### GetSortFieldNameOk

`func (o *SortingAttributes) GetSortFieldNameOk() (*string, bool)`

GetSortFieldNameOk returns a tuple with the SortFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortFieldName

`func (o *SortingAttributes) SetSortFieldName(v string)`

SetSortFieldName sets SortFieldName field to given value.

### HasSortFieldName

`func (o *SortingAttributes) HasSortFieldName() bool`

HasSortFieldName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


