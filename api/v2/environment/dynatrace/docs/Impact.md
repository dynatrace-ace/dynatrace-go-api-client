# Impact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImpactType** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;SERVICE&#x60; -&gt; ServiceImpact  * &#x60;APPLICATION&#x60; -&gt; ApplicationImpact  * &#x60;MOBILE&#x60; -&gt; MobileImpact  * &#x60;CUSTOM_APPLICATION&#x60; -&gt; CustomApplicationImpact   | 
**ImpactedEntity** | [**EntityStub**](EntityStub.md) |  | 
**EstimatedAffectedUsers** | **int64** | The estimated number of affected users. | 

## Methods

### NewImpact

`func NewImpact(impactType string, impactedEntity EntityStub, estimatedAffectedUsers int64, ) *Impact`

NewImpact instantiates a new Impact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImpactWithDefaults

`func NewImpactWithDefaults() *Impact`

NewImpactWithDefaults instantiates a new Impact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpactType

`func (o *Impact) GetImpactType() string`

GetImpactType returns the ImpactType field if non-nil, zero value otherwise.

### GetImpactTypeOk

`func (o *Impact) GetImpactTypeOk() (*string, bool)`

GetImpactTypeOk returns a tuple with the ImpactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactType

`func (o *Impact) SetImpactType(v string)`

SetImpactType sets ImpactType field to given value.


### GetImpactedEntity

`func (o *Impact) GetImpactedEntity() EntityStub`

GetImpactedEntity returns the ImpactedEntity field if non-nil, zero value otherwise.

### GetImpactedEntityOk

`func (o *Impact) GetImpactedEntityOk() (*EntityStub, bool)`

GetImpactedEntityOk returns a tuple with the ImpactedEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactedEntity

`func (o *Impact) SetImpactedEntity(v EntityStub)`

SetImpactedEntity sets ImpactedEntity field to given value.


### GetEstimatedAffectedUsers

`func (o *Impact) GetEstimatedAffectedUsers() int64`

GetEstimatedAffectedUsers returns the EstimatedAffectedUsers field if non-nil, zero value otherwise.

### GetEstimatedAffectedUsersOk

`func (o *Impact) GetEstimatedAffectedUsersOk() (*int64, bool)`

GetEstimatedAffectedUsersOk returns a tuple with the EstimatedAffectedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedAffectedUsers

`func (o *Impact) SetEstimatedAffectedUsers(v int64)`

SetEstimatedAffectedUsers sets EstimatedAffectedUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


