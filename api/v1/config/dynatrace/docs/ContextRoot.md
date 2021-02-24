# ContextRoot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transformations** | Pointer to [**[]ContextRootTransformation**](ContextRootTransformation.md) | Transformations to be applied to the detected value. | [optional] 
**SegmentsToCopyFromUrlPath** | Pointer to **int32** | The number of segments of the URL to be kept.   The URL is divided by slashes (&#x60;/&#x60;), the indexing starts with &#x60;1&#x60; at context root.   For example, if you specify &#x60;2&#x60; for the &#x60;www.dynatrace.com/support/help/dynatrace-api/&#x60; URL, the value of &#x60;support/help&#x60; is used. | [optional] 

## Methods

### NewContextRoot

`func NewContextRoot() *ContextRoot`

NewContextRoot instantiates a new ContextRoot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextRootWithDefaults

`func NewContextRootWithDefaults() *ContextRoot`

NewContextRootWithDefaults instantiates a new ContextRoot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransformations

`func (o *ContextRoot) GetTransformations() []ContextRootTransformation`

GetTransformations returns the Transformations field if non-nil, zero value otherwise.

### GetTransformationsOk

`func (o *ContextRoot) GetTransformationsOk() (*[]ContextRootTransformation, bool)`

GetTransformationsOk returns a tuple with the Transformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformations

`func (o *ContextRoot) SetTransformations(v []ContextRootTransformation)`

SetTransformations sets Transformations field to given value.

### HasTransformations

`func (o *ContextRoot) HasTransformations() bool`

HasTransformations returns a boolean if a field has been set.

### GetSegmentsToCopyFromUrlPath

`func (o *ContextRoot) GetSegmentsToCopyFromUrlPath() int32`

GetSegmentsToCopyFromUrlPath returns the SegmentsToCopyFromUrlPath field if non-nil, zero value otherwise.

### GetSegmentsToCopyFromUrlPathOk

`func (o *ContextRoot) GetSegmentsToCopyFromUrlPathOk() (*int32, bool)`

GetSegmentsToCopyFromUrlPathOk returns a tuple with the SegmentsToCopyFromUrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentsToCopyFromUrlPath

`func (o *ContextRoot) SetSegmentsToCopyFromUrlPath(v int32)`

SetSegmentsToCopyFromUrlPath sets SegmentsToCopyFromUrlPath field to given value.

### HasSegmentsToCopyFromUrlPath

`func (o *ContextRoot) HasSegmentsToCopyFromUrlPath() bool`

HasSegmentsToCopyFromUrlPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


