# MultiDatacenterTopology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatacenterTopology** | Pointer to [**map[string]DatacenterDesc**](DatacenterDesc.md) | Map of multidatacenter topology | [optional] 
**OldDatacenterName** | Pointer to **string** |  | [optional] 
**NewDatacenterName** | Pointer to **string** |  | [optional] 

## Methods

### NewMultiDatacenterTopology

`func NewMultiDatacenterTopology() *MultiDatacenterTopology`

NewMultiDatacenterTopology instantiates a new MultiDatacenterTopology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiDatacenterTopologyWithDefaults

`func NewMultiDatacenterTopologyWithDefaults() *MultiDatacenterTopology`

NewMultiDatacenterTopologyWithDefaults instantiates a new MultiDatacenterTopology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterTopology

`func (o *MultiDatacenterTopology) GetDatacenterTopology() map[string]DatacenterDesc`

GetDatacenterTopology returns the DatacenterTopology field if non-nil, zero value otherwise.

### GetDatacenterTopologyOk

`func (o *MultiDatacenterTopology) GetDatacenterTopologyOk() (*map[string]DatacenterDesc, bool)`

GetDatacenterTopologyOk returns a tuple with the DatacenterTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterTopology

`func (o *MultiDatacenterTopology) SetDatacenterTopology(v map[string]DatacenterDesc)`

SetDatacenterTopology sets DatacenterTopology field to given value.

### HasDatacenterTopology

`func (o *MultiDatacenterTopology) HasDatacenterTopology() bool`

HasDatacenterTopology returns a boolean if a field has been set.

### GetOldDatacenterName

`func (o *MultiDatacenterTopology) GetOldDatacenterName() string`

GetOldDatacenterName returns the OldDatacenterName field if non-nil, zero value otherwise.

### GetOldDatacenterNameOk

`func (o *MultiDatacenterTopology) GetOldDatacenterNameOk() (*string, bool)`

GetOldDatacenterNameOk returns a tuple with the OldDatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldDatacenterName

`func (o *MultiDatacenterTopology) SetOldDatacenterName(v string)`

SetOldDatacenterName sets OldDatacenterName field to given value.

### HasOldDatacenterName

`func (o *MultiDatacenterTopology) HasOldDatacenterName() bool`

HasOldDatacenterName returns a boolean if a field has been set.

### GetNewDatacenterName

`func (o *MultiDatacenterTopology) GetNewDatacenterName() string`

GetNewDatacenterName returns the NewDatacenterName field if non-nil, zero value otherwise.

### GetNewDatacenterNameOk

`func (o *MultiDatacenterTopology) GetNewDatacenterNameOk() (*string, bool)`

GetNewDatacenterNameOk returns a tuple with the NewDatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDatacenterName

`func (o *MultiDatacenterTopology) SetNewDatacenterName(v string)`

SetNewDatacenterName sets NewDatacenterName field to given value.

### HasNewDatacenterName

`func (o *MultiDatacenterTopology) HasNewDatacenterName() bool`

HasNewDatacenterName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


