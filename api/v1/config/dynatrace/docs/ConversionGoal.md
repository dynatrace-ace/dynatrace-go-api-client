# ConversionGoal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the conversion goal. | 
**Id** | Pointer to **string** | The ID of conversion goal.    Omit it while creating a new conversion goal. | [optional] 
**Type** | Pointer to **string** | The type of the conversion goal. | [optional] 
**DestinationDetails** | Pointer to [**DestinationDetails**](DestinationDetails.md) |  | [optional] 
**UserActionDetails** | Pointer to [**UserActionDetails**](UserActionDetails.md) |  | [optional] 
**VisitDurationDetails** | Pointer to [**VisitDurationDetails**](VisitDurationDetails.md) |  | [optional] 
**VisitNumActionDetails** | Pointer to [**VisitNumActionDetails**](VisitNumActionDetails.md) |  | [optional] 

## Methods

### NewConversionGoal

`func NewConversionGoal(name string, ) *ConversionGoal`

NewConversionGoal instantiates a new ConversionGoal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionGoalWithDefaults

`func NewConversionGoalWithDefaults() *ConversionGoal`

NewConversionGoalWithDefaults instantiates a new ConversionGoal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConversionGoal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConversionGoal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConversionGoal) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *ConversionGoal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConversionGoal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConversionGoal) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConversionGoal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ConversionGoal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConversionGoal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConversionGoal) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConversionGoal) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDestinationDetails

`func (o *ConversionGoal) GetDestinationDetails() DestinationDetails`

GetDestinationDetails returns the DestinationDetails field if non-nil, zero value otherwise.

### GetDestinationDetailsOk

`func (o *ConversionGoal) GetDestinationDetailsOk() (*DestinationDetails, bool)`

GetDestinationDetailsOk returns a tuple with the DestinationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDetails

`func (o *ConversionGoal) SetDestinationDetails(v DestinationDetails)`

SetDestinationDetails sets DestinationDetails field to given value.

### HasDestinationDetails

`func (o *ConversionGoal) HasDestinationDetails() bool`

HasDestinationDetails returns a boolean if a field has been set.

### GetUserActionDetails

`func (o *ConversionGoal) GetUserActionDetails() UserActionDetails`

GetUserActionDetails returns the UserActionDetails field if non-nil, zero value otherwise.

### GetUserActionDetailsOk

`func (o *ConversionGoal) GetUserActionDetailsOk() (*UserActionDetails, bool)`

GetUserActionDetailsOk returns a tuple with the UserActionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionDetails

`func (o *ConversionGoal) SetUserActionDetails(v UserActionDetails)`

SetUserActionDetails sets UserActionDetails field to given value.

### HasUserActionDetails

`func (o *ConversionGoal) HasUserActionDetails() bool`

HasUserActionDetails returns a boolean if a field has been set.

### GetVisitDurationDetails

`func (o *ConversionGoal) GetVisitDurationDetails() VisitDurationDetails`

GetVisitDurationDetails returns the VisitDurationDetails field if non-nil, zero value otherwise.

### GetVisitDurationDetailsOk

`func (o *ConversionGoal) GetVisitDurationDetailsOk() (*VisitDurationDetails, bool)`

GetVisitDurationDetailsOk returns a tuple with the VisitDurationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitDurationDetails

`func (o *ConversionGoal) SetVisitDurationDetails(v VisitDurationDetails)`

SetVisitDurationDetails sets VisitDurationDetails field to given value.

### HasVisitDurationDetails

`func (o *ConversionGoal) HasVisitDurationDetails() bool`

HasVisitDurationDetails returns a boolean if a field has been set.

### GetVisitNumActionDetails

`func (o *ConversionGoal) GetVisitNumActionDetails() VisitNumActionDetails`

GetVisitNumActionDetails returns the VisitNumActionDetails field if non-nil, zero value otherwise.

### GetVisitNumActionDetailsOk

`func (o *ConversionGoal) GetVisitNumActionDetailsOk() (*VisitNumActionDetails, bool)`

GetVisitNumActionDetailsOk returns a tuple with the VisitNumActionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitNumActionDetails

`func (o *ConversionGoal) SetVisitNumActionDetails(v VisitNumActionDetails)`

SetVisitNumActionDetails sets VisitNumActionDetails field to given value.

### HasVisitNumActionDetails

`func (o *ConversionGoal) HasVisitNumActionDetails() bool`

HasVisitNumActionDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


