# DatacenterDesc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatacenterMembersList** | Pointer to [**[]DatacenterMember**](DatacenterMember.md) |  | [optional] 
**NewDatacenter** | Pointer to **bool** |  | [optional] 

## Methods

### NewDatacenterDesc

`func NewDatacenterDesc() *DatacenterDesc`

NewDatacenterDesc instantiates a new DatacenterDesc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterDescWithDefaults

`func NewDatacenterDescWithDefaults() *DatacenterDesc`

NewDatacenterDescWithDefaults instantiates a new DatacenterDesc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterMembersList

`func (o *DatacenterDesc) GetDatacenterMembersList() []DatacenterMember`

GetDatacenterMembersList returns the DatacenterMembersList field if non-nil, zero value otherwise.

### GetDatacenterMembersListOk

`func (o *DatacenterDesc) GetDatacenterMembersListOk() (*[]DatacenterMember, bool)`

GetDatacenterMembersListOk returns a tuple with the DatacenterMembersList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterMembersList

`func (o *DatacenterDesc) SetDatacenterMembersList(v []DatacenterMember)`

SetDatacenterMembersList sets DatacenterMembersList field to given value.

### HasDatacenterMembersList

`func (o *DatacenterDesc) HasDatacenterMembersList() bool`

HasDatacenterMembersList returns a boolean if a field has been set.

### GetNewDatacenter

`func (o *DatacenterDesc) GetNewDatacenter() bool`

GetNewDatacenter returns the NewDatacenter field if non-nil, zero value otherwise.

### GetNewDatacenterOk

`func (o *DatacenterDesc) GetNewDatacenterOk() (*bool, bool)`

GetNewDatacenterOk returns a tuple with the NewDatacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDatacenter

`func (o *DatacenterDesc) SetNewDatacenter(v bool)`

SetNewDatacenter sets NewDatacenter field to given value.

### HasNewDatacenter

`func (o *DatacenterDesc) HasNewDatacenter() bool`

HasNewDatacenter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


