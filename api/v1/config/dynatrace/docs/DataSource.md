# DataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The data source is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**Source** | **string** | The source of the attribute to capture. Works in conjunction with **parameterName** or **methods** and **technology**. | 
**ValueProcessing** | Pointer to [**ValueProcessing**](ValueProcessing.md) |  | [optional] 
**Technology** | Pointer to **string** | The technology of the method to capture if the **source** value is &#x60;METHOD_PARAM&#x60;. \\n\\n Not applicable in other cases. | [optional] 
**SessionAttributeTechnology** | Pointer to **string** | The technology of the session attribute to capture if the **source** value is &#x60;SESSION_ATTRIBUTE&#x60;. \\n\\n Not applicable in other cases. | [optional] 
**Methods** | Pointer to [**[]CapturedMethod**](CapturedMethod.md) | The method specification if the **source** value is &#x60;METHOD_PARAM&#x60;.    Not applicable in other cases. | [optional] 
**ParameterName** | Pointer to **string** | The name of the web request parameter to capture.   Required if the **source** is one of the following: &#x60;POST_PARAMETER&#x60;, &#x60;GET_PARAMETER&#x60;, &#x60;REQUEST_HEADER&#x60;, &#x60;RESPONSE_HEADER&#x60;, &#x60;CUSTOM_ATTRIBUTE&#x60;.   Not applicable in other cases. | [optional] 
**Scope** | Pointer to [**ScopeConditions**](ScopeConditions.md) |  | [optional] 
**CapturingAndStorageLocation** | Pointer to **string** | Specifies the location where the values are captured and stored.   Required if the **source** is one of the following: &#x60;GET_PARAMETER&#x60;, &#x60;URI&#x60;, &#x60;REQUEST_HEADER&#x60;, &#x60;RESPONSE_HEADER&#x60;.    Not applicable in other cases.    If the **source** value is &#x60;REQUEST_HEADER&#x60; or &#x60;RESPONSE_HEADER&#x60;, the &#x60;CAPTURE_AND_STORE_ON_BOTH&#x60; location is not allowed. | [optional] 
**IibNodeType** | Pointer to **string** | The IBM integration bus node type for which the value is captured.   This or &#x60;iibMethodNodeCondition&#x60; is required if the **source** is: &#x60;IIB_NODE&#x60;.   Not applicable in other cases. | [optional] 
**IibMethodNodeCondition** | Pointer to [**ValueCondition**](ValueCondition.md) |  | [optional] 
**CicsSDKMethodNodeCondition** | Pointer to [**ValueCondition**](ValueCondition.md) |  | [optional] 
**IibLabelMethodNodeCondition** | Pointer to [**ValueCondition**](ValueCondition.md) |  | [optional] 
**SpanAttributeKey** | Pointer to **string** | The key of the span attribute to capture.   Required if the **source** is: &#x60;SPAN_ATTRIBUTE&#x60;.   Not applicable in other cases. | [optional] 

## Methods

### NewDataSource

`func NewDataSource(enabled bool, source string, ) *DataSource`

NewDataSource instantiates a new DataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceWithDefaults

`func NewDataSourceWithDefaults() *DataSource`

NewDataSourceWithDefaults instantiates a new DataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DataSource) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DataSource) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DataSource) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSource

`func (o *DataSource) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DataSource) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DataSource) SetSource(v string)`

SetSource sets Source field to given value.


### GetValueProcessing

`func (o *DataSource) GetValueProcessing() ValueProcessing`

GetValueProcessing returns the ValueProcessing field if non-nil, zero value otherwise.

### GetValueProcessingOk

`func (o *DataSource) GetValueProcessingOk() (*ValueProcessing, bool)`

GetValueProcessingOk returns a tuple with the ValueProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueProcessing

`func (o *DataSource) SetValueProcessing(v ValueProcessing)`

SetValueProcessing sets ValueProcessing field to given value.

### HasValueProcessing

`func (o *DataSource) HasValueProcessing() bool`

HasValueProcessing returns a boolean if a field has been set.

### GetTechnology

`func (o *DataSource) GetTechnology() string`

GetTechnology returns the Technology field if non-nil, zero value otherwise.

### GetTechnologyOk

`func (o *DataSource) GetTechnologyOk() (*string, bool)`

GetTechnologyOk returns a tuple with the Technology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnology

`func (o *DataSource) SetTechnology(v string)`

SetTechnology sets Technology field to given value.

### HasTechnology

`func (o *DataSource) HasTechnology() bool`

HasTechnology returns a boolean if a field has been set.

### GetSessionAttributeTechnology

`func (o *DataSource) GetSessionAttributeTechnology() string`

GetSessionAttributeTechnology returns the SessionAttributeTechnology field if non-nil, zero value otherwise.

### GetSessionAttributeTechnologyOk

`func (o *DataSource) GetSessionAttributeTechnologyOk() (*string, bool)`

GetSessionAttributeTechnologyOk returns a tuple with the SessionAttributeTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAttributeTechnology

`func (o *DataSource) SetSessionAttributeTechnology(v string)`

SetSessionAttributeTechnology sets SessionAttributeTechnology field to given value.

### HasSessionAttributeTechnology

`func (o *DataSource) HasSessionAttributeTechnology() bool`

HasSessionAttributeTechnology returns a boolean if a field has been set.

### GetMethods

`func (o *DataSource) GetMethods() []CapturedMethod`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *DataSource) GetMethodsOk() (*[]CapturedMethod, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *DataSource) SetMethods(v []CapturedMethod)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *DataSource) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetParameterName

`func (o *DataSource) GetParameterName() string`

GetParameterName returns the ParameterName field if non-nil, zero value otherwise.

### GetParameterNameOk

`func (o *DataSource) GetParameterNameOk() (*string, bool)`

GetParameterNameOk returns a tuple with the ParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterName

`func (o *DataSource) SetParameterName(v string)`

SetParameterName sets ParameterName field to given value.

### HasParameterName

`func (o *DataSource) HasParameterName() bool`

HasParameterName returns a boolean if a field has been set.

### GetScope

`func (o *DataSource) GetScope() ScopeConditions`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DataSource) GetScopeOk() (*ScopeConditions, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DataSource) SetScope(v ScopeConditions)`

SetScope sets Scope field to given value.

### HasScope

`func (o *DataSource) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetCapturingAndStorageLocation

`func (o *DataSource) GetCapturingAndStorageLocation() string`

GetCapturingAndStorageLocation returns the CapturingAndStorageLocation field if non-nil, zero value otherwise.

### GetCapturingAndStorageLocationOk

`func (o *DataSource) GetCapturingAndStorageLocationOk() (*string, bool)`

GetCapturingAndStorageLocationOk returns a tuple with the CapturingAndStorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturingAndStorageLocation

`func (o *DataSource) SetCapturingAndStorageLocation(v string)`

SetCapturingAndStorageLocation sets CapturingAndStorageLocation field to given value.

### HasCapturingAndStorageLocation

`func (o *DataSource) HasCapturingAndStorageLocation() bool`

HasCapturingAndStorageLocation returns a boolean if a field has been set.

### GetIibNodeType

`func (o *DataSource) GetIibNodeType() string`

GetIibNodeType returns the IibNodeType field if non-nil, zero value otherwise.

### GetIibNodeTypeOk

`func (o *DataSource) GetIibNodeTypeOk() (*string, bool)`

GetIibNodeTypeOk returns a tuple with the IibNodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIibNodeType

`func (o *DataSource) SetIibNodeType(v string)`

SetIibNodeType sets IibNodeType field to given value.

### HasIibNodeType

`func (o *DataSource) HasIibNodeType() bool`

HasIibNodeType returns a boolean if a field has been set.

### GetIibMethodNodeCondition

`func (o *DataSource) GetIibMethodNodeCondition() ValueCondition`

GetIibMethodNodeCondition returns the IibMethodNodeCondition field if non-nil, zero value otherwise.

### GetIibMethodNodeConditionOk

`func (o *DataSource) GetIibMethodNodeConditionOk() (*ValueCondition, bool)`

GetIibMethodNodeConditionOk returns a tuple with the IibMethodNodeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIibMethodNodeCondition

`func (o *DataSource) SetIibMethodNodeCondition(v ValueCondition)`

SetIibMethodNodeCondition sets IibMethodNodeCondition field to given value.

### HasIibMethodNodeCondition

`func (o *DataSource) HasIibMethodNodeCondition() bool`

HasIibMethodNodeCondition returns a boolean if a field has been set.

### GetCicsSDKMethodNodeCondition

`func (o *DataSource) GetCicsSDKMethodNodeCondition() ValueCondition`

GetCicsSDKMethodNodeCondition returns the CicsSDKMethodNodeCondition field if non-nil, zero value otherwise.

### GetCicsSDKMethodNodeConditionOk

`func (o *DataSource) GetCicsSDKMethodNodeConditionOk() (*ValueCondition, bool)`

GetCicsSDKMethodNodeConditionOk returns a tuple with the CicsSDKMethodNodeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCicsSDKMethodNodeCondition

`func (o *DataSource) SetCicsSDKMethodNodeCondition(v ValueCondition)`

SetCicsSDKMethodNodeCondition sets CicsSDKMethodNodeCondition field to given value.

### HasCicsSDKMethodNodeCondition

`func (o *DataSource) HasCicsSDKMethodNodeCondition() bool`

HasCicsSDKMethodNodeCondition returns a boolean if a field has been set.

### GetIibLabelMethodNodeCondition

`func (o *DataSource) GetIibLabelMethodNodeCondition() ValueCondition`

GetIibLabelMethodNodeCondition returns the IibLabelMethodNodeCondition field if non-nil, zero value otherwise.

### GetIibLabelMethodNodeConditionOk

`func (o *DataSource) GetIibLabelMethodNodeConditionOk() (*ValueCondition, bool)`

GetIibLabelMethodNodeConditionOk returns a tuple with the IibLabelMethodNodeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIibLabelMethodNodeCondition

`func (o *DataSource) SetIibLabelMethodNodeCondition(v ValueCondition)`

SetIibLabelMethodNodeCondition sets IibLabelMethodNodeCondition field to given value.

### HasIibLabelMethodNodeCondition

`func (o *DataSource) HasIibLabelMethodNodeCondition() bool`

HasIibLabelMethodNodeCondition returns a boolean if a field has been set.

### GetSpanAttributeKey

`func (o *DataSource) GetSpanAttributeKey() string`

GetSpanAttributeKey returns the SpanAttributeKey field if non-nil, zero value otherwise.

### GetSpanAttributeKeyOk

`func (o *DataSource) GetSpanAttributeKeyOk() (*string, bool)`

GetSpanAttributeKeyOk returns a tuple with the SpanAttributeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanAttributeKey

`func (o *DataSource) SetSpanAttributeKey(v string)`

SetSpanAttributeKey sets SpanAttributeKey field to given value.

### HasSpanAttributeKey

`func (o *DataSource) HasSpanAttributeKey() bool`

HasSpanAttributeKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


