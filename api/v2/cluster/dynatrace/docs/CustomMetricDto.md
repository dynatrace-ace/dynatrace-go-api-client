# CustomMetricDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewCustomMetricDto

`func NewCustomMetricDto() *CustomMetricDto`

NewCustomMetricDto instantiates a new CustomMetricDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomMetricDtoWithDefaults

`func NewCustomMetricDtoWithDefaults() *CustomMetricDto`

NewCustomMetricDtoWithDefaults instantiates a new CustomMetricDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *CustomMetricDto) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CustomMetricDto) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CustomMetricDto) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CustomMetricDto) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTotal

`func (o *CustomMetricDto) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CustomMetricDto) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CustomMetricDto) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CustomMetricDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


