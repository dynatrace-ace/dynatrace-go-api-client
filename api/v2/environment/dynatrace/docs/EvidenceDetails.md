# EvidenceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int64** | The total number of evidence of a problem. | 
**Details** | Pointer to [**[]Evidence**](Evidence.md) | A list of all evidence. | [optional] 

## Methods

### NewEvidenceDetails

`func NewEvidenceDetails(totalCount int64, ) *EvidenceDetails`

NewEvidenceDetails instantiates a new EvidenceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidenceDetailsWithDefaults

`func NewEvidenceDetailsWithDefaults() *EvidenceDetails`

NewEvidenceDetailsWithDefaults instantiates a new EvidenceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *EvidenceDetails) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *EvidenceDetails) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *EvidenceDetails) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetDetails

`func (o *EvidenceDetails) GetDetails() []Evidence`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *EvidenceDetails) GetDetailsOk() (*[]Evidence, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *EvidenceDetails) SetDetails(v []Evidence)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *EvidenceDetails) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


