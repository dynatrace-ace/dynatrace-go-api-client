# KubernetesEventPattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | A label of the events field selector. | 
**FieldSelector** | **string** | The field selector string (url decoding is applied) when storing it. | 
**Active** | **bool** | Whether subscription to this events field selector is enabled (value set to &#x60;true&#x60;). If disabled (value set to &#x60;false&#x60;), Dynatrace will stop fetching events from the Kubernetes API for this events field selector | 

## Methods

### NewKubernetesEventPattern

`func NewKubernetesEventPattern(label string, fieldSelector string, active bool, ) *KubernetesEventPattern`

NewKubernetesEventPattern instantiates a new KubernetesEventPattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEventPatternWithDefaults

`func NewKubernetesEventPatternWithDefaults() *KubernetesEventPattern`

NewKubernetesEventPatternWithDefaults instantiates a new KubernetesEventPattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *KubernetesEventPattern) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *KubernetesEventPattern) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *KubernetesEventPattern) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetFieldSelector

`func (o *KubernetesEventPattern) GetFieldSelector() string`

GetFieldSelector returns the FieldSelector field if non-nil, zero value otherwise.

### GetFieldSelectorOk

`func (o *KubernetesEventPattern) GetFieldSelectorOk() (*string, bool)`

GetFieldSelectorOk returns a tuple with the FieldSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldSelector

`func (o *KubernetesEventPattern) SetFieldSelector(v string)`

SetFieldSelector sets FieldSelector field to given value.


### GetActive

`func (o *KubernetesEventPattern) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *KubernetesEventPattern) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *KubernetesEventPattern) SetActive(v bool)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


