# AwsConfigTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The key of the AWS tag. | 
**Value** | **string** | The value of the AWS tag. | 

## Methods

### NewAwsConfigTag

`func NewAwsConfigTag(name string, value string, ) *AwsConfigTag`

NewAwsConfigTag instantiates a new AwsConfigTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsConfigTagWithDefaults

`func NewAwsConfigTagWithDefaults() *AwsConfigTag`

NewAwsConfigTagWithDefaults instantiates a new AwsConfigTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AwsConfigTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsConfigTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsConfigTag) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *AwsConfigTag) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AwsConfigTag) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AwsConfigTag) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


