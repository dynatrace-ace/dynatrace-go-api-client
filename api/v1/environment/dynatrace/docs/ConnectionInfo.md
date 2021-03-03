# ConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantUUID** | Pointer to **string** | The ID of your Dynatrace environment. | [optional] 
**TenantToken** | Pointer to **string** | The internal token that is used for authentication when OneAgent connects to the Dynatrace cluster to send data. | [optional] 
**CommunicationEndpoints** | Pointer to **[]string** | The list of endpoints to connect to the Dynatrace environment. The list is sorted by endpoint priority, descending. | [optional] 

## Methods

### NewConnectionInfo

`func NewConnectionInfo() *ConnectionInfo`

NewConnectionInfo instantiates a new ConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionInfoWithDefaults

`func NewConnectionInfoWithDefaults() *ConnectionInfo`

NewConnectionInfoWithDefaults instantiates a new ConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantUUID

`func (o *ConnectionInfo) GetTenantUUID() string`

GetTenantUUID returns the TenantUUID field if non-nil, zero value otherwise.

### GetTenantUUIDOk

`func (o *ConnectionInfo) GetTenantUUIDOk() (*string, bool)`

GetTenantUUIDOk returns a tuple with the TenantUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUUID

`func (o *ConnectionInfo) SetTenantUUID(v string)`

SetTenantUUID sets TenantUUID field to given value.

### HasTenantUUID

`func (o *ConnectionInfo) HasTenantUUID() bool`

HasTenantUUID returns a boolean if a field has been set.

### GetTenantToken

`func (o *ConnectionInfo) GetTenantToken() string`

GetTenantToken returns the TenantToken field if non-nil, zero value otherwise.

### GetTenantTokenOk

`func (o *ConnectionInfo) GetTenantTokenOk() (*string, bool)`

GetTenantTokenOk returns a tuple with the TenantToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantToken

`func (o *ConnectionInfo) SetTenantToken(v string)`

SetTenantToken sets TenantToken field to given value.

### HasTenantToken

`func (o *ConnectionInfo) HasTenantToken() bool`

HasTenantToken returns a boolean if a field has been set.

### GetCommunicationEndpoints

`func (o *ConnectionInfo) GetCommunicationEndpoints() []string`

GetCommunicationEndpoints returns the CommunicationEndpoints field if non-nil, zero value otherwise.

### GetCommunicationEndpointsOk

`func (o *ConnectionInfo) GetCommunicationEndpointsOk() (*[]string, bool)`

GetCommunicationEndpointsOk returns a tuple with the CommunicationEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationEndpoints

`func (o *ConnectionInfo) SetCommunicationEndpoints(v []string)`

SetCommunicationEndpoints sets CommunicationEndpoints field to given value.

### HasCommunicationEndpoints

`func (o *ConnectionInfo) HasCommunicationEndpoints() bool`

HasCommunicationEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


