# ApplicationFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | **string** | The value to look for. | 
**ApplicationMatchType** | **string** | The operator of the matching. | 
**ApplicationMatchTarget** | **string** | Where to look for the the **pattern** value. | 

## Methods

### NewApplicationFilter

`func NewApplicationFilter(pattern string, applicationMatchType string, applicationMatchTarget string, ) *ApplicationFilter`

NewApplicationFilter instantiates a new ApplicationFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationFilterWithDefaults

`func NewApplicationFilterWithDefaults() *ApplicationFilter`

NewApplicationFilterWithDefaults instantiates a new ApplicationFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *ApplicationFilter) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ApplicationFilter) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ApplicationFilter) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetApplicationMatchType

`func (o *ApplicationFilter) GetApplicationMatchType() string`

GetApplicationMatchType returns the ApplicationMatchType field if non-nil, zero value otherwise.

### GetApplicationMatchTypeOk

`func (o *ApplicationFilter) GetApplicationMatchTypeOk() (*string, bool)`

GetApplicationMatchTypeOk returns a tuple with the ApplicationMatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationMatchType

`func (o *ApplicationFilter) SetApplicationMatchType(v string)`

SetApplicationMatchType sets ApplicationMatchType field to given value.


### GetApplicationMatchTarget

`func (o *ApplicationFilter) GetApplicationMatchTarget() string`

GetApplicationMatchTarget returns the ApplicationMatchTarget field if non-nil, zero value otherwise.

### GetApplicationMatchTargetOk

`func (o *ApplicationFilter) GetApplicationMatchTargetOk() (*string, bool)`

GetApplicationMatchTargetOk returns a tuple with the ApplicationMatchTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationMatchTarget

`func (o *ApplicationFilter) SetApplicationMatchTarget(v string)`

SetApplicationMatchTarget sets ApplicationMatchTarget field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


