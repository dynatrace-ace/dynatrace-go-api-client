# CloudTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the tag. | [optional] 
**Value** | Pointer to **string** | The value of the tag.    If set to &#x60;null&#x60;, then resources with any value of the tag are monitored. | [optional] 

## Methods

### NewCloudTag

`func NewCloudTag() *CloudTag`

NewCloudTag instantiates a new CloudTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTagWithDefaults

`func NewCloudTagWithDefaults() *CloudTag`

NewCloudTagWithDefaults instantiates a new CloudTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudTag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudTag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *CloudTag) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CloudTag) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CloudTag) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CloudTag) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


