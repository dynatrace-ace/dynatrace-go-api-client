# ComparisonBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | Operator of the comparison. You can reverse it by setting **negate** to &#x60;true&#x60;.   Possible values depend on the **type** of the comparison. Find the list of actuala models in the description of the **type** field and check the description of the model you need. | 
**Value** | Pointer to **map[string]interface{}** | The value to compare to. | [optional] 
**Negate** | **bool** | Reverses the comparison **operator**. For example it turns the **begins with** into **does not begin with**. | 
**CaseSensitive** | Pointer to **bool** | The comparison is case-sensitive (&#x60;true&#x60;) or insensitive (&#x60;false&#x60;). | [optional] 
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;STRING&#x60; -&gt; StringComparison  * &#x60;INDEXED_NAME&#x60; -&gt; IndexedNameComparison  * &#x60;INDEXED_STRING&#x60; -&gt; IndexedStringComparison  * &#x60;INTEGER&#x60; -&gt; IntegerComparison  * &#x60;SERVICE_TYPE&#x60; -&gt; ServiceTypeComparison  * &#x60;PAAS_TYPE&#x60; -&gt; PaasTypeComparison  * &#x60;CLOUD_TYPE&#x60; -&gt; CloudTypeComparison  * &#x60;AZURE_SKU&#x60; -&gt; AzureSkuComparision  * &#x60;AZURE_COMPUTE_MODE&#x60; -&gt; AzureComputeModeComparison  * &#x60;ENTITY_ID&#x60; -&gt; EntityIdComparison  * &#x60;SIMPLE_TECH&#x60; -&gt; SimpleTechComparison  * &#x60;SIMPLE_HOST_TECH&#x60; -&gt; SimpleHostTechComparison  * &#x60;SERVICE_TOPOLOGY&#x60; -&gt; ServiceTopologyComparison  * &#x60;DATABASE_TOPOLOGY&#x60; -&gt; DatabaseTopologyComparison  * &#x60;OS_TYPE&#x60; -&gt; OsTypeComparison  * &#x60;HYPERVISOR_TYPE&#x60; -&gt; HypervisorTypeComparision  * &#x60;IP_ADDRESS&#x60; -&gt; IpAddressComparison  * &#x60;OS_ARCHITECTURE&#x60; -&gt; OsArchitectureComparison  * &#x60;BITNESS&#x60; -&gt; BitnessComparision  * &#x60;APPLICATION_TYPE&#x60; -&gt; ApplicationTypeComparison  * &#x60;MOBILE_PLATFORM&#x60; -&gt; MobilePlatformComparison  * &#x60;CUSTOM_APPLICATION_TYPE&#x60; -&gt; CustomApplicationTypeComparison  * &#x60;DCRUM_DECODER_TYPE&#x60; -&gt; DcrumDecoderComparison  * &#x60;SYNTHETIC_ENGINE_TYPE&#x60; -&gt; SyntheticEngineTypeComparison  * &#x60;TAG&#x60; -&gt; TagComparison  * &#x60;INDEXED_TAG&#x60; -&gt; IndexedTagComparison   | 

## Methods

### NewComparisonBasic

`func NewComparisonBasic(operator string, negate bool, type_ string, ) *ComparisonBasic`

NewComparisonBasic instantiates a new ComparisonBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComparisonBasicWithDefaults

`func NewComparisonBasicWithDefaults() *ComparisonBasic`

NewComparisonBasicWithDefaults instantiates a new ComparisonBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ComparisonBasic) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ComparisonBasic) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ComparisonBasic) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *ComparisonBasic) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ComparisonBasic) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ComparisonBasic) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ComparisonBasic) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetNegate

`func (o *ComparisonBasic) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *ComparisonBasic) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *ComparisonBasic) SetNegate(v bool)`

SetNegate sets Negate field to given value.


### GetCaseSensitive

`func (o *ComparisonBasic) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *ComparisonBasic) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *ComparisonBasic) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *ComparisonBasic) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetType

`func (o *ComparisonBasic) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComparisonBasic) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComparisonBasic) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


