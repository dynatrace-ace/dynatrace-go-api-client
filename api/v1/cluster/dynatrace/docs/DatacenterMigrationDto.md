# DatacenterMigrationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewDatacenterName** | Pointer to **string** |  | [optional] 
**NodesIp** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDatacenterMigrationDto

`func NewDatacenterMigrationDto() *DatacenterMigrationDto`

NewDatacenterMigrationDto instantiates a new DatacenterMigrationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterMigrationDtoWithDefaults

`func NewDatacenterMigrationDtoWithDefaults() *DatacenterMigrationDto`

NewDatacenterMigrationDtoWithDefaults instantiates a new DatacenterMigrationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewDatacenterName

`func (o *DatacenterMigrationDto) GetNewDatacenterName() string`

GetNewDatacenterName returns the NewDatacenterName field if non-nil, zero value otherwise.

### GetNewDatacenterNameOk

`func (o *DatacenterMigrationDto) GetNewDatacenterNameOk() (*string, bool)`

GetNewDatacenterNameOk returns a tuple with the NewDatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDatacenterName

`func (o *DatacenterMigrationDto) SetNewDatacenterName(v string)`

SetNewDatacenterName sets NewDatacenterName field to given value.

### HasNewDatacenterName

`func (o *DatacenterMigrationDto) HasNewDatacenterName() bool`

HasNewDatacenterName returns a boolean if a field has been set.

### GetNodesIp

`func (o *DatacenterMigrationDto) GetNodesIp() []string`

GetNodesIp returns the NodesIp field if non-nil, zero value otherwise.

### GetNodesIpOk

`func (o *DatacenterMigrationDto) GetNodesIpOk() (*[]string, bool)`

GetNodesIpOk returns a tuple with the NodesIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesIp

`func (o *DatacenterMigrationDto) SetNodesIp(v []string)`

SetNodesIp sets NodesIp field to given value.

### HasNodesIp

`func (o *DatacenterMigrationDto) HasNodesIp() bool`

HasNodesIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


