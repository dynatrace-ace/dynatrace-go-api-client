# Extension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the extension, for example &#x60;custom.remote.python.demo&#x60;. | [optional] 
**Name** | Pointer to **string** | The name of the extension, displayed in Dynatrace. | [optional] 
**Version** | Pointer to **string** | The version of the extension, displayed in Dynatrace. | [optional] 
**Type** | Pointer to **string** | The type of the extension. It indicates the runtime environment of the extension (for example, ACTIVEGATE). | [optional] 
**MetricGroup** | Pointer to **string** | The metricGroup of the extension used for grouping custom metrics into a hierarchical namespace. | [optional] 
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Properties** | Pointer to [**[]ExtensionProperty**](ExtensionProperty.md) | A list of extension properties. | [optional] 

## Methods

### NewExtension

`func NewExtension() *Extension`

NewExtension instantiates a new Extension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionWithDefaults

`func NewExtensionWithDefaults() *Extension`

NewExtensionWithDefaults instantiates a new Extension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Extension) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Extension) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Extension) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Extension) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Extension) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Extension) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Extension) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Extension) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *Extension) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Extension) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Extension) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Extension) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *Extension) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Extension) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Extension) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Extension) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMetricGroup

`func (o *Extension) GetMetricGroup() string`

GetMetricGroup returns the MetricGroup field if non-nil, zero value otherwise.

### GetMetricGroupOk

`func (o *Extension) GetMetricGroupOk() (*string, bool)`

GetMetricGroupOk returns a tuple with the MetricGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricGroup

`func (o *Extension) SetMetricGroup(v string)`

SetMetricGroup sets MetricGroup field to given value.

### HasMetricGroup

`func (o *Extension) HasMetricGroup() bool`

HasMetricGroup returns a boolean if a field has been set.

### GetMetadata

`func (o *Extension) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Extension) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Extension) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Extension) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *Extension) GetProperties() []ExtensionProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Extension) GetPropertiesOk() (*[]ExtensionProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Extension) SetProperties(v []ExtensionProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Extension) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


