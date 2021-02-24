# ColumnDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;CUSTOM&#x60; -&gt; CustomColumnDefinition  * &#x60;JSON&#x60; -&gt; JsonColumnDefinition   | 

## Methods

### NewColumnDefinition

`func NewColumnDefinition(name string, type_ string, ) *ColumnDefinition`

NewColumnDefinition instantiates a new ColumnDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnDefinitionWithDefaults

`func NewColumnDefinitionWithDefaults() *ColumnDefinition`

NewColumnDefinitionWithDefaults instantiates a new ColumnDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ColumnDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ColumnDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ColumnDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ColumnDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ColumnDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ColumnDefinition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


