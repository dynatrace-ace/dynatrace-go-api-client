# PathDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definition** | **string** | The path to the required log path.    If the **type** is set to &#x60;WILDCARD&#x60;, it may contain wildcard characters (&#x60;*&#x60;). | 
**Type** | **string** | The type of the log path **definition**: fixed or an expression with wildcards. | 

## Methods

### NewPathDefinition

`func NewPathDefinition(definition string, type_ string, ) *PathDefinition`

NewPathDefinition instantiates a new PathDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathDefinitionWithDefaults

`func NewPathDefinitionWithDefaults() *PathDefinition`

NewPathDefinitionWithDefaults instantiates a new PathDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *PathDefinition) GetDefinition() string`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *PathDefinition) GetDefinitionOk() (*string, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *PathDefinition) SetDefinition(v string)`

SetDefinition sets Definition field to given value.


### GetType

`func (o *PathDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PathDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PathDefinition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


