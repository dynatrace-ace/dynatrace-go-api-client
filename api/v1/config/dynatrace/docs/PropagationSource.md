# PropagationSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagementZone** | Pointer to **string** | Use only request attributes from services that belong to this management zone.. Use either this or &#x60;serviceTag&#x60; | [optional] 
**ServiceTag** | Pointer to [**UniversalTag**](UniversalTag.md) |  | [optional] 

## Methods

### NewPropagationSource

`func NewPropagationSource() *PropagationSource`

NewPropagationSource instantiates a new PropagationSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropagationSourceWithDefaults

`func NewPropagationSourceWithDefaults() *PropagationSource`

NewPropagationSourceWithDefaults instantiates a new PropagationSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagementZone

`func (o *PropagationSource) GetManagementZone() string`

GetManagementZone returns the ManagementZone field if non-nil, zero value otherwise.

### GetManagementZoneOk

`func (o *PropagationSource) GetManagementZoneOk() (*string, bool)`

GetManagementZoneOk returns a tuple with the ManagementZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZone

`func (o *PropagationSource) SetManagementZone(v string)`

SetManagementZone sets ManagementZone field to given value.

### HasManagementZone

`func (o *PropagationSource) HasManagementZone() bool`

HasManagementZone returns a boolean if a field has been set.

### GetServiceTag

`func (o *PropagationSource) GetServiceTag() UniversalTag`

GetServiceTag returns the ServiceTag field if non-nil, zero value otherwise.

### GetServiceTagOk

`func (o *PropagationSource) GetServiceTagOk() (*UniversalTag, bool)`

GetServiceTagOk returns a tuple with the ServiceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTag

`func (o *PropagationSource) SetServiceTag(v UniversalTag)`

SetServiceTag sets ServiceTag field to given value.

### HasServiceTag

`func (o *PropagationSource) HasServiceTag() bool`

HasServiceTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


