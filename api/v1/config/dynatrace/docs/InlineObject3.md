# InlineObject3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | [***os.File**](*os.File.md) | Plugin ZIP file to be uploaded.    File name must match the **name** field in the &#x60;plugin.json&#x60; file.   For example, for the plugin whose **name** is &#x60;custom.remote.python.demo&#x60;, the name of the plugin file must be &#x60;custom.remote.python.demo.zip&#x60;. | 

## Methods

### NewInlineObject3

`func NewInlineObject3(file *os.File, ) *InlineObject3`

NewInlineObject3 instantiates a new InlineObject3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject3WithDefaults

`func NewInlineObject3WithDefaults() *InlineObject3`

NewInlineObject3WithDefaults instantiates a new InlineObject3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *InlineObject3) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *InlineObject3) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *InlineObject3) SetFile(v *os.File)`

SetFile sets File field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


