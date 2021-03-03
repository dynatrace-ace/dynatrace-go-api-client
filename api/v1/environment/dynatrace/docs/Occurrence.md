# Occurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | Value of top parsing field occurrence | [optional] 
**Count** | Pointer to **int64** | Count of top parsing field occurrences | [optional] 

## Methods

### NewOccurrence

`func NewOccurrence() *Occurrence`

NewOccurrence instantiates a new Occurrence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOccurrenceWithDefaults

`func NewOccurrenceWithDefaults() *Occurrence`

NewOccurrenceWithDefaults instantiates a new Occurrence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Occurrence) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Occurrence) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Occurrence) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Occurrence) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCount

`func (o *Occurrence) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Occurrence) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Occurrence) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *Occurrence) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


