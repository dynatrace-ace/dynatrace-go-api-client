# FilterTopLogRecords

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterQuery** | Pointer to **string** | The query for filtering.   See the [Search patterns in log data and parse results](https://dt-url.net/hf23k34) help page for syntax description. | [optional] 

## Methods

### NewFilterTopLogRecords

`func NewFilterTopLogRecords() *FilterTopLogRecords`

NewFilterTopLogRecords instantiates a new FilterTopLogRecords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterTopLogRecordsWithDefaults

`func NewFilterTopLogRecordsWithDefaults() *FilterTopLogRecords`

NewFilterTopLogRecordsWithDefaults instantiates a new FilterTopLogRecords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterQuery

`func (o *FilterTopLogRecords) GetFilterQuery() string`

GetFilterQuery returns the FilterQuery field if non-nil, zero value otherwise.

### GetFilterQueryOk

`func (o *FilterTopLogRecords) GetFilterQueryOk() (*string, bool)`

GetFilterQueryOk returns a tuple with the FilterQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterQuery

`func (o *FilterTopLogRecords) SetFilterQuery(v string)`

SetFilterQuery sets FilterQuery field to given value.

### HasFilterQuery

`func (o *FilterTopLogRecords) HasFilterQuery() bool`

HasFilterQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


