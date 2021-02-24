# RequestAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the request attribute. | [optional] 
**Name** | **string** | The name of the request attribute. | 
**Enabled** | **bool** | The request attribute is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**DataType** | **string** | The data type of the request attribute. | 
**DataSources** | [**[]DataSource**](DataSource.md) | The list of data sources. | 
**Normalization** | **string** | String values transformation.    If the **dataType** is not &#x60;string&#x60;, set the &#x60;Original&#x60; here. | 
**Aggregation** | **string** | Aggregation type for the request values. | 
**Confidential** | **bool** | Confidential data flag. Set &#x60;true&#x60; to treat the captured data as confidential. | 
**SkipPersonalDataMasking** | **bool** | Personal data masking flag. Set &#x60;true&#x60; to skip masking.    Warning: This will potentially access personalized data. | 

## Methods

### NewRequestAttribute

`func NewRequestAttribute(name string, enabled bool, dataType string, dataSources []DataSource, normalization string, aggregation string, confidential bool, skipPersonalDataMasking bool, ) *RequestAttribute`

NewRequestAttribute instantiates a new RequestAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestAttributeWithDefaults

`func NewRequestAttributeWithDefaults() *RequestAttribute`

NewRequestAttributeWithDefaults instantiates a new RequestAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *RequestAttribute) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RequestAttribute) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RequestAttribute) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RequestAttribute) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *RequestAttribute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestAttribute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestAttribute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequestAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RequestAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *RequestAttribute) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RequestAttribute) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RequestAttribute) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDataType

`func (o *RequestAttribute) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *RequestAttribute) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *RequestAttribute) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetDataSources

`func (o *RequestAttribute) GetDataSources() []DataSource`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *RequestAttribute) GetDataSourcesOk() (*[]DataSource, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *RequestAttribute) SetDataSources(v []DataSource)`

SetDataSources sets DataSources field to given value.


### GetNormalization

`func (o *RequestAttribute) GetNormalization() string`

GetNormalization returns the Normalization field if non-nil, zero value otherwise.

### GetNormalizationOk

`func (o *RequestAttribute) GetNormalizationOk() (*string, bool)`

GetNormalizationOk returns a tuple with the Normalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalization

`func (o *RequestAttribute) SetNormalization(v string)`

SetNormalization sets Normalization field to given value.


### GetAggregation

`func (o *RequestAttribute) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *RequestAttribute) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *RequestAttribute) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.


### GetConfidential

`func (o *RequestAttribute) GetConfidential() bool`

GetConfidential returns the Confidential field if non-nil, zero value otherwise.

### GetConfidentialOk

`func (o *RequestAttribute) GetConfidentialOk() (*bool, bool)`

GetConfidentialOk returns a tuple with the Confidential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidential

`func (o *RequestAttribute) SetConfidential(v bool)`

SetConfidential sets Confidential field to given value.


### GetSkipPersonalDataMasking

`func (o *RequestAttribute) GetSkipPersonalDataMasking() bool`

GetSkipPersonalDataMasking returns the SkipPersonalDataMasking field if non-nil, zero value otherwise.

### GetSkipPersonalDataMaskingOk

`func (o *RequestAttribute) GetSkipPersonalDataMaskingOk() (*bool, bool)`

GetSkipPersonalDataMaskingOk returns a tuple with the SkipPersonalDataMasking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPersonalDataMasking

`func (o *RequestAttribute) SetSkipPersonalDataMasking(v bool)`

SetSkipPersonalDataMasking sets SkipPersonalDataMasking field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


