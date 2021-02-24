# UserActionNamingPlaceholder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Placeholder name. | 
**Input** | **string** | Input. | 
**ProcessingPart** | **string** | Part. | 
**ProcessingSteps** | Pointer to [**[]UserActionNamingPlaceholderProcessingStep**](UserActionNamingPlaceholderProcessingStep.md) | Processing actions. | [optional] 
**MetadataId** | Pointer to **int32** | Id of the metadata. | [optional] 
**UseGuessedElementIdentifier** | **bool** | Use the element identifier that was selected by Dynatrace. | 

## Methods

### NewUserActionNamingPlaceholder

`func NewUserActionNamingPlaceholder(name string, input string, processingPart string, useGuessedElementIdentifier bool, ) *UserActionNamingPlaceholder`

NewUserActionNamingPlaceholder instantiates a new UserActionNamingPlaceholder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActionNamingPlaceholderWithDefaults

`func NewUserActionNamingPlaceholderWithDefaults() *UserActionNamingPlaceholder`

NewUserActionNamingPlaceholderWithDefaults instantiates a new UserActionNamingPlaceholder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserActionNamingPlaceholder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserActionNamingPlaceholder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserActionNamingPlaceholder) SetName(v string)`

SetName sets Name field to given value.


### GetInput

`func (o *UserActionNamingPlaceholder) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *UserActionNamingPlaceholder) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *UserActionNamingPlaceholder) SetInput(v string)`

SetInput sets Input field to given value.


### GetProcessingPart

`func (o *UserActionNamingPlaceholder) GetProcessingPart() string`

GetProcessingPart returns the ProcessingPart field if non-nil, zero value otherwise.

### GetProcessingPartOk

`func (o *UserActionNamingPlaceholder) GetProcessingPartOk() (*string, bool)`

GetProcessingPartOk returns a tuple with the ProcessingPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingPart

`func (o *UserActionNamingPlaceholder) SetProcessingPart(v string)`

SetProcessingPart sets ProcessingPart field to given value.


### GetProcessingSteps

`func (o *UserActionNamingPlaceholder) GetProcessingSteps() []UserActionNamingPlaceholderProcessingStep`

GetProcessingSteps returns the ProcessingSteps field if non-nil, zero value otherwise.

### GetProcessingStepsOk

`func (o *UserActionNamingPlaceholder) GetProcessingStepsOk() (*[]UserActionNamingPlaceholderProcessingStep, bool)`

GetProcessingStepsOk returns a tuple with the ProcessingSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingSteps

`func (o *UserActionNamingPlaceholder) SetProcessingSteps(v []UserActionNamingPlaceholderProcessingStep)`

SetProcessingSteps sets ProcessingSteps field to given value.

### HasProcessingSteps

`func (o *UserActionNamingPlaceholder) HasProcessingSteps() bool`

HasProcessingSteps returns a boolean if a field has been set.

### GetMetadataId

`func (o *UserActionNamingPlaceholder) GetMetadataId() int32`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *UserActionNamingPlaceholder) GetMetadataIdOk() (*int32, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *UserActionNamingPlaceholder) SetMetadataId(v int32)`

SetMetadataId sets MetadataId field to given value.

### HasMetadataId

`func (o *UserActionNamingPlaceholder) HasMetadataId() bool`

HasMetadataId returns a boolean if a field has been set.

### GetUseGuessedElementIdentifier

`func (o *UserActionNamingPlaceholder) GetUseGuessedElementIdentifier() bool`

GetUseGuessedElementIdentifier returns the UseGuessedElementIdentifier field if non-nil, zero value otherwise.

### GetUseGuessedElementIdentifierOk

`func (o *UserActionNamingPlaceholder) GetUseGuessedElementIdentifierOk() (*bool, bool)`

GetUseGuessedElementIdentifierOk returns a tuple with the UseGuessedElementIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGuessedElementIdentifier

`func (o *UserActionNamingPlaceholder) SetUseGuessedElementIdentifier(v bool)`

SetUseGuessedElementIdentifier sets UseGuessedElementIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


