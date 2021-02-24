# AwsCredentialsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The unique ID of the credentials. | [optional] [readonly] 
**ConnectionStatus** | Pointer to **string** | The status of the connection to the AWS environment.    * &#x60;CONNECTED&#x60;: There was a connection within last 10 minutes.  * &#x60;DISCONNECTED&#x60;: A problem occurred with establishing connection using these credentials. Check whether the data is correct.  * &#x60;UNINITIALIZED&#x60;: The successful connection has never been established for these credentials. | [optional] [readonly] 
**Label** | **string** | The name of the credentials. | 
**PartitionType** | **string** | The type of the AWS partition. | 
**AuthenticationData** | [**AwsAuthenticationData**](AwsAuthenticationData.md) |  | 
**TaggedOnly** | **bool** | Monitor only resources which have specified AWS tags (&#x60;true&#x60;) or all resources (&#x60;false&#x60;). | 
**TagsToMonitor** | [**[]AwsConfigTag**](AwsConfigTag.md) | A list of AWS tags to be monitored.   You can specify up to 10 tags.   Only applicable when the **taggedOnly** parameter is set to &#x60;true&#x60;. | 
**SupportingServicesToMonitor** | Pointer to [**[]AwsSupportingServiceConfig**](AwsSupportingServiceConfig.md) | A list of supporting services to be monitored. | [optional] 

## Methods

### NewAwsCredentialsConfig

`func NewAwsCredentialsConfig(label string, partitionType string, authenticationData AwsAuthenticationData, taggedOnly bool, tagsToMonitor []AwsConfigTag, ) *AwsCredentialsConfig`

NewAwsCredentialsConfig instantiates a new AwsCredentialsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsCredentialsConfigWithDefaults

`func NewAwsCredentialsConfigWithDefaults() *AwsCredentialsConfig`

NewAwsCredentialsConfigWithDefaults instantiates a new AwsCredentialsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *AwsCredentialsConfig) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsCredentialsConfig) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsCredentialsConfig) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AwsCredentialsConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *AwsCredentialsConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsCredentialsConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsCredentialsConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsCredentialsConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *AwsCredentialsConfig) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *AwsCredentialsConfig) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *AwsCredentialsConfig) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *AwsCredentialsConfig) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetLabel

`func (o *AwsCredentialsConfig) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AwsCredentialsConfig) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AwsCredentialsConfig) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPartitionType

`func (o *AwsCredentialsConfig) GetPartitionType() string`

GetPartitionType returns the PartitionType field if non-nil, zero value otherwise.

### GetPartitionTypeOk

`func (o *AwsCredentialsConfig) GetPartitionTypeOk() (*string, bool)`

GetPartitionTypeOk returns a tuple with the PartitionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionType

`func (o *AwsCredentialsConfig) SetPartitionType(v string)`

SetPartitionType sets PartitionType field to given value.


### GetAuthenticationData

`func (o *AwsCredentialsConfig) GetAuthenticationData() AwsAuthenticationData`

GetAuthenticationData returns the AuthenticationData field if non-nil, zero value otherwise.

### GetAuthenticationDataOk

`func (o *AwsCredentialsConfig) GetAuthenticationDataOk() (*AwsAuthenticationData, bool)`

GetAuthenticationDataOk returns a tuple with the AuthenticationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationData

`func (o *AwsCredentialsConfig) SetAuthenticationData(v AwsAuthenticationData)`

SetAuthenticationData sets AuthenticationData field to given value.


### GetTaggedOnly

`func (o *AwsCredentialsConfig) GetTaggedOnly() bool`

GetTaggedOnly returns the TaggedOnly field if non-nil, zero value otherwise.

### GetTaggedOnlyOk

`func (o *AwsCredentialsConfig) GetTaggedOnlyOk() (*bool, bool)`

GetTaggedOnlyOk returns a tuple with the TaggedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedOnly

`func (o *AwsCredentialsConfig) SetTaggedOnly(v bool)`

SetTaggedOnly sets TaggedOnly field to given value.


### GetTagsToMonitor

`func (o *AwsCredentialsConfig) GetTagsToMonitor() []AwsConfigTag`

GetTagsToMonitor returns the TagsToMonitor field if non-nil, zero value otherwise.

### GetTagsToMonitorOk

`func (o *AwsCredentialsConfig) GetTagsToMonitorOk() (*[]AwsConfigTag, bool)`

GetTagsToMonitorOk returns a tuple with the TagsToMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsToMonitor

`func (o *AwsCredentialsConfig) SetTagsToMonitor(v []AwsConfigTag)`

SetTagsToMonitor sets TagsToMonitor field to given value.


### GetSupportingServicesToMonitor

`func (o *AwsCredentialsConfig) GetSupportingServicesToMonitor() []AwsSupportingServiceConfig`

GetSupportingServicesToMonitor returns the SupportingServicesToMonitor field if non-nil, zero value otherwise.

### GetSupportingServicesToMonitorOk

`func (o *AwsCredentialsConfig) GetSupportingServicesToMonitorOk() (*[]AwsSupportingServiceConfig, bool)`

GetSupportingServicesToMonitorOk returns a tuple with the SupportingServicesToMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportingServicesToMonitor

`func (o *AwsCredentialsConfig) SetSupportingServicesToMonitor(v []AwsSupportingServiceConfig)`

SetSupportingServicesToMonitor sets SupportingServicesToMonitor field to given value.

### HasSupportingServicesToMonitor

`func (o *AwsCredentialsConfig) HasSupportingServicesToMonitor() bool`

HasSupportingServicesToMonitor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


