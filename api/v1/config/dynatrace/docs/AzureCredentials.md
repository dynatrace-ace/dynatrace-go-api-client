# AzureCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The Dynatrace entity ID of the Azure credentials configuration. | [optional] [readonly] 
**Label** | **string** | The unique name of the Azure credentials configuration.   Allowed characters are letters, numbers, and spaces. Also the special characters &#x60;.+-_&#x60; are allowed. | 
**AppId** | **string** | The Application ID (also referred to as Client ID)   The combination of Application ID and Directory ID must be unique. | [readonly] 
**DirectoryId** | **string** | The Directory ID (also referred to as Tenant ID)   The combination of Application ID and Directory ID must be unique. | [readonly] 
**Key** | Pointer to **string** | The secret key associated with the Application ID.   For security reasons, GET requests return this field as &#x60;null&#x60;.    Submit your key on creation or update of the configuration. If the field is omitted during an update, the old value remains unaffected. | [optional] 
**Active** | Pointer to **bool** | The monitoring is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;).   If not set on creation, the &#x60;true&#x60; value is used.   If the field is omitted during an update, the old value remains unaffected. | [optional] 
**AutoTagging** | **bool** | The automatic capture of Azure tags is on (&#x60;true&#x60;) or off (&#x60;false&#x60;). | 
**MonitorOnlyTaggedEntities** | **bool** | Monitor only resources that have specified Azure tags (&#x60;true&#x60;) or all resources (&#x60;false&#x60;). | 
**MonitorOnlyTagPairs** | [**[]CloudTag**](CloudTag.md) | A list of Azure tags to be monitored.   You can specify up to 10 tags. A resource tagged with *any* of the specified tags is monitored.   Only applicable when the **monitorOnlyTaggedEntities** parameter is set to &#x60;true&#x60;. | 
**SupportingServices** | Pointer to [**[]AzureSupportingService**](AzureSupportingService.md) | A list of Azure supporting services to be monitored. For each service there&#39;s a sublist of its metrics and the metrics&#39; dimensions that should be monitored. All of these elements (services, metrics, dimensions) must have corresponding static definitions on the server. | [optional] 

## Methods

### NewAzureCredentials

`func NewAzureCredentials(label string, appId string, directoryId string, autoTagging bool, monitorOnlyTaggedEntities bool, monitorOnlyTagPairs []CloudTag, ) *AzureCredentials`

NewAzureCredentials instantiates a new AzureCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCredentialsWithDefaults

`func NewAzureCredentialsWithDefaults() *AzureCredentials`

NewAzureCredentialsWithDefaults instantiates a new AzureCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *AzureCredentials) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AzureCredentials) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AzureCredentials) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AzureCredentials) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *AzureCredentials) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureCredentials) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureCredentials) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureCredentials) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *AzureCredentials) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AzureCredentials) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AzureCredentials) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetAppId

`func (o *AzureCredentials) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AzureCredentials) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AzureCredentials) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetDirectoryId

`func (o *AzureCredentials) GetDirectoryId() string`

GetDirectoryId returns the DirectoryId field if non-nil, zero value otherwise.

### GetDirectoryIdOk

`func (o *AzureCredentials) GetDirectoryIdOk() (*string, bool)`

GetDirectoryIdOk returns a tuple with the DirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryId

`func (o *AzureCredentials) SetDirectoryId(v string)`

SetDirectoryId sets DirectoryId field to given value.


### GetKey

`func (o *AzureCredentials) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AzureCredentials) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AzureCredentials) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AzureCredentials) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetActive

`func (o *AzureCredentials) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AzureCredentials) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AzureCredentials) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AzureCredentials) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAutoTagging

`func (o *AzureCredentials) GetAutoTagging() bool`

GetAutoTagging returns the AutoTagging field if non-nil, zero value otherwise.

### GetAutoTaggingOk

`func (o *AzureCredentials) GetAutoTaggingOk() (*bool, bool)`

GetAutoTaggingOk returns a tuple with the AutoTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTagging

`func (o *AzureCredentials) SetAutoTagging(v bool)`

SetAutoTagging sets AutoTagging field to given value.


### GetMonitorOnlyTaggedEntities

`func (o *AzureCredentials) GetMonitorOnlyTaggedEntities() bool`

GetMonitorOnlyTaggedEntities returns the MonitorOnlyTaggedEntities field if non-nil, zero value otherwise.

### GetMonitorOnlyTaggedEntitiesOk

`func (o *AzureCredentials) GetMonitorOnlyTaggedEntitiesOk() (*bool, bool)`

GetMonitorOnlyTaggedEntitiesOk returns a tuple with the MonitorOnlyTaggedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorOnlyTaggedEntities

`func (o *AzureCredentials) SetMonitorOnlyTaggedEntities(v bool)`

SetMonitorOnlyTaggedEntities sets MonitorOnlyTaggedEntities field to given value.


### GetMonitorOnlyTagPairs

`func (o *AzureCredentials) GetMonitorOnlyTagPairs() []CloudTag`

GetMonitorOnlyTagPairs returns the MonitorOnlyTagPairs field if non-nil, zero value otherwise.

### GetMonitorOnlyTagPairsOk

`func (o *AzureCredentials) GetMonitorOnlyTagPairsOk() (*[]CloudTag, bool)`

GetMonitorOnlyTagPairsOk returns a tuple with the MonitorOnlyTagPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorOnlyTagPairs

`func (o *AzureCredentials) SetMonitorOnlyTagPairs(v []CloudTag)`

SetMonitorOnlyTagPairs sets MonitorOnlyTagPairs field to given value.


### GetSupportingServices

`func (o *AzureCredentials) GetSupportingServices() []AzureSupportingService`

GetSupportingServices returns the SupportingServices field if non-nil, zero value otherwise.

### GetSupportingServicesOk

`func (o *AzureCredentials) GetSupportingServicesOk() (*[]AzureSupportingService, bool)`

GetSupportingServicesOk returns a tuple with the SupportingServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportingServices

`func (o *AzureCredentials) SetSupportingServices(v []AzureSupportingService)`

SetSupportingServices sets SupportingServices field to given value.

### HasSupportingServices

`func (o *AzureCredentials) HasSupportingServices() bool`

HasSupportingServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


