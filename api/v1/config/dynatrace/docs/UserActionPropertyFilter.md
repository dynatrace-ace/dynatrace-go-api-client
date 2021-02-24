# UserActionPropertyFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key of the action property we&#39;re checking. | [optional] 
**Value** | Pointer to **string** | Only actions that have this value in the specified property are included in the metric calculation.    Only applicable to string values. | [optional] 
**From** | Pointer to **float64** | Only actions that have a value greater than or equal to this are included in the metric calculation.    Only applicable to numerical values. | [optional] 
**To** | Pointer to **float64** | Only actions that have a value less than or equal to this are included in the metric calculation.    Only applicable to numerical values. | [optional] 
**MatchType** | Pointer to **string** | Specifies the match type of a string filter, e.g. using &#x60;Contains&#x60; or &#x60;Equals&#x60;.    Only applicable to string values. | [optional] 

## Methods

### NewUserActionPropertyFilter

`func NewUserActionPropertyFilter() *UserActionPropertyFilter`

NewUserActionPropertyFilter instantiates a new UserActionPropertyFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActionPropertyFilterWithDefaults

`func NewUserActionPropertyFilterWithDefaults() *UserActionPropertyFilter`

NewUserActionPropertyFilterWithDefaults instantiates a new UserActionPropertyFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *UserActionPropertyFilter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UserActionPropertyFilter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UserActionPropertyFilter) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UserActionPropertyFilter) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *UserActionPropertyFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserActionPropertyFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserActionPropertyFilter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UserActionPropertyFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetFrom

`func (o *UserActionPropertyFilter) GetFrom() float64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *UserActionPropertyFilter) GetFromOk() (*float64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *UserActionPropertyFilter) SetFrom(v float64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *UserActionPropertyFilter) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *UserActionPropertyFilter) GetTo() float64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *UserActionPropertyFilter) GetToOk() (*float64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *UserActionPropertyFilter) SetTo(v float64)`

SetTo sets To field to given value.

### HasTo

`func (o *UserActionPropertyFilter) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetMatchType

`func (o *UserActionPropertyFilter) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *UserActionPropertyFilter) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *UserActionPropertyFilter) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *UserActionPropertyFilter) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


