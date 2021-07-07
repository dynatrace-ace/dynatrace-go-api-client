# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Node ID | [optional] 
**ClusterMemberAddress** | Pointer to **string** | Cluster member address | [optional] 
**OperationState** | Pointer to **string** | Operation state | [optional] 
**BuildVersion** | Pointer to **string** | Server version | [optional] 
**OsInfo** | Pointer to **string** | OS info | [optional] 
**JvmInfo** | Pointer to **string** | JVM info | [optional] 
**DnsEntryPointUris** | Pointer to **[]string** | DNS entry point URIs | [optional] 
**RestServiceRootUris** | Pointer to **[]string** | REST service root URIs | [optional] 
**CommunicationUris** | Pointer to **[]string** | Communication URIs | [optional] 

## Methods

### NewCluster

`func NewCluster() *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cluster) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Cluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClusterMemberAddress

`func (o *Cluster) GetClusterMemberAddress() string`

GetClusterMemberAddress returns the ClusterMemberAddress field if non-nil, zero value otherwise.

### GetClusterMemberAddressOk

`func (o *Cluster) GetClusterMemberAddressOk() (*string, bool)`

GetClusterMemberAddressOk returns a tuple with the ClusterMemberAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMemberAddress

`func (o *Cluster) SetClusterMemberAddress(v string)`

SetClusterMemberAddress sets ClusterMemberAddress field to given value.

### HasClusterMemberAddress

`func (o *Cluster) HasClusterMemberAddress() bool`

HasClusterMemberAddress returns a boolean if a field has been set.

### GetOperationState

`func (o *Cluster) GetOperationState() string`

GetOperationState returns the OperationState field if non-nil, zero value otherwise.

### GetOperationStateOk

`func (o *Cluster) GetOperationStateOk() (*string, bool)`

GetOperationStateOk returns a tuple with the OperationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationState

`func (o *Cluster) SetOperationState(v string)`

SetOperationState sets OperationState field to given value.

### HasOperationState

`func (o *Cluster) HasOperationState() bool`

HasOperationState returns a boolean if a field has been set.

### GetBuildVersion

`func (o *Cluster) GetBuildVersion() string`

GetBuildVersion returns the BuildVersion field if non-nil, zero value otherwise.

### GetBuildVersionOk

`func (o *Cluster) GetBuildVersionOk() (*string, bool)`

GetBuildVersionOk returns a tuple with the BuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersion

`func (o *Cluster) SetBuildVersion(v string)`

SetBuildVersion sets BuildVersion field to given value.

### HasBuildVersion

`func (o *Cluster) HasBuildVersion() bool`

HasBuildVersion returns a boolean if a field has been set.

### GetOsInfo

`func (o *Cluster) GetOsInfo() string`

GetOsInfo returns the OsInfo field if non-nil, zero value otherwise.

### GetOsInfoOk

`func (o *Cluster) GetOsInfoOk() (*string, bool)`

GetOsInfoOk returns a tuple with the OsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInfo

`func (o *Cluster) SetOsInfo(v string)`

SetOsInfo sets OsInfo field to given value.

### HasOsInfo

`func (o *Cluster) HasOsInfo() bool`

HasOsInfo returns a boolean if a field has been set.

### GetJvmInfo

`func (o *Cluster) GetJvmInfo() string`

GetJvmInfo returns the JvmInfo field if non-nil, zero value otherwise.

### GetJvmInfoOk

`func (o *Cluster) GetJvmInfoOk() (*string, bool)`

GetJvmInfoOk returns a tuple with the JvmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJvmInfo

`func (o *Cluster) SetJvmInfo(v string)`

SetJvmInfo sets JvmInfo field to given value.

### HasJvmInfo

`func (o *Cluster) HasJvmInfo() bool`

HasJvmInfo returns a boolean if a field has been set.

### GetDnsEntryPointUris

`func (o *Cluster) GetDnsEntryPointUris() []string`

GetDnsEntryPointUris returns the DnsEntryPointUris field if non-nil, zero value otherwise.

### GetDnsEntryPointUrisOk

`func (o *Cluster) GetDnsEntryPointUrisOk() (*[]string, bool)`

GetDnsEntryPointUrisOk returns a tuple with the DnsEntryPointUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsEntryPointUris

`func (o *Cluster) SetDnsEntryPointUris(v []string)`

SetDnsEntryPointUris sets DnsEntryPointUris field to given value.

### HasDnsEntryPointUris

`func (o *Cluster) HasDnsEntryPointUris() bool`

HasDnsEntryPointUris returns a boolean if a field has been set.

### GetRestServiceRootUris

`func (o *Cluster) GetRestServiceRootUris() []string`

GetRestServiceRootUris returns the RestServiceRootUris field if non-nil, zero value otherwise.

### GetRestServiceRootUrisOk

`func (o *Cluster) GetRestServiceRootUrisOk() (*[]string, bool)`

GetRestServiceRootUrisOk returns a tuple with the RestServiceRootUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestServiceRootUris

`func (o *Cluster) SetRestServiceRootUris(v []string)`

SetRestServiceRootUris sets RestServiceRootUris field to given value.

### HasRestServiceRootUris

`func (o *Cluster) HasRestServiceRootUris() bool`

HasRestServiceRootUris returns a boolean if a field has been set.

### GetCommunicationUris

`func (o *Cluster) GetCommunicationUris() []string`

GetCommunicationUris returns the CommunicationUris field if non-nil, zero value otherwise.

### GetCommunicationUrisOk

`func (o *Cluster) GetCommunicationUrisOk() (*[]string, bool)`

GetCommunicationUrisOk returns a tuple with the CommunicationUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationUris

`func (o *Cluster) SetCommunicationUris(v []string)`

SetCommunicationUris sets CommunicationUris field to given value.

### HasCommunicationUris

`func (o *Cluster) HasCommunicationUris() bool`

HasCommunicationUris returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


