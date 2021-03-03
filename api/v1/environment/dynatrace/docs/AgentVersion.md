# AgentVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Major** | Pointer to **int32** | The major version number. | [optional] 
**Minor** | Pointer to **int32** | The minor version number. | [optional] 
**Revision** | Pointer to **int32** | The revision number. | [optional] 
**Timestamp** | Pointer to **string** | A timestamp string: format \&quot;yyyymmdd-hhmmss | [optional] 
**SourceRevision** | Pointer to **string** | A string representation of the SVN revision number. | [optional] 

## Methods

### NewAgentVersion

`func NewAgentVersion() *AgentVersion`

NewAgentVersion instantiates a new AgentVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentVersionWithDefaults

`func NewAgentVersionWithDefaults() *AgentVersion`

NewAgentVersionWithDefaults instantiates a new AgentVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMajor

`func (o *AgentVersion) GetMajor() int32`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *AgentVersion) GetMajorOk() (*int32, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *AgentVersion) SetMajor(v int32)`

SetMajor sets Major field to given value.

### HasMajor

`func (o *AgentVersion) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetMinor

`func (o *AgentVersion) GetMinor() int32`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *AgentVersion) GetMinorOk() (*int32, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *AgentVersion) SetMinor(v int32)`

SetMinor sets Minor field to given value.

### HasMinor

`func (o *AgentVersion) HasMinor() bool`

HasMinor returns a boolean if a field has been set.

### GetRevision

`func (o *AgentVersion) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *AgentVersion) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *AgentVersion) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *AgentVersion) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetTimestamp

`func (o *AgentVersion) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AgentVersion) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AgentVersion) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AgentVersion) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetSourceRevision

`func (o *AgentVersion) GetSourceRevision() string`

GetSourceRevision returns the SourceRevision field if non-nil, zero value otherwise.

### GetSourceRevisionOk

`func (o *AgentVersion) GetSourceRevisionOk() (*string, bool)`

GetSourceRevisionOk returns a tuple with the SourceRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRevision

`func (o *AgentVersion) SetSourceRevision(v string)`

SetSourceRevision sets SourceRevision field to given value.

### HasSourceRevision

`func (o *AgentVersion) HasSourceRevision() bool`

HasSourceRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


