# MetaDataCapturing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the meta data to capture. | 
**CapturingName** | **string** | The name of the meta data to capture. | 
**Name** | **string** | Name for displaying the captured values in Dynatrace. | 
**UniqueId** | Pointer to **int32** | The unique id of the meta data to capture. | [optional] 
**PublicMetadata** | Pointer to **bool** | True if this metadata should be captured regardless of the privacy settings | [optional] 

## Methods

### NewMetaDataCapturing

`func NewMetaDataCapturing(type_ string, capturingName string, name string, ) *MetaDataCapturing`

NewMetaDataCapturing instantiates a new MetaDataCapturing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaDataCapturingWithDefaults

`func NewMetaDataCapturingWithDefaults() *MetaDataCapturing`

NewMetaDataCapturingWithDefaults instantiates a new MetaDataCapturing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetaDataCapturing) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetaDataCapturing) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetaDataCapturing) SetType(v string)`

SetType sets Type field to given value.


### GetCapturingName

`func (o *MetaDataCapturing) GetCapturingName() string`

GetCapturingName returns the CapturingName field if non-nil, zero value otherwise.

### GetCapturingNameOk

`func (o *MetaDataCapturing) GetCapturingNameOk() (*string, bool)`

GetCapturingNameOk returns a tuple with the CapturingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturingName

`func (o *MetaDataCapturing) SetCapturingName(v string)`

SetCapturingName sets CapturingName field to given value.


### GetName

`func (o *MetaDataCapturing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaDataCapturing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaDataCapturing) SetName(v string)`

SetName sets Name field to given value.


### GetUniqueId

`func (o *MetaDataCapturing) GetUniqueId() int32`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *MetaDataCapturing) GetUniqueIdOk() (*int32, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *MetaDataCapturing) SetUniqueId(v int32)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *MetaDataCapturing) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetPublicMetadata

`func (o *MetaDataCapturing) GetPublicMetadata() bool`

GetPublicMetadata returns the PublicMetadata field if non-nil, zero value otherwise.

### GetPublicMetadataOk

`func (o *MetaDataCapturing) GetPublicMetadataOk() (*bool, bool)`

GetPublicMetadataOk returns a tuple with the PublicMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicMetadata

`func (o *MetaDataCapturing) SetPublicMetadata(v bool)`

SetPublicMetadata sets PublicMetadata field to given value.

### HasPublicMetadata

`func (o *MetaDataCapturing) HasPublicMetadata() bool`

HasPublicMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


