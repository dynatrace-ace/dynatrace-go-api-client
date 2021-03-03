# ProblemCloseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProblemId** | Pointer to **string** | The ID of the problem. | [optional] 
**Comment** | Pointer to [**ProblemComment**](ProblemComment.md) |  | [optional] 
**CloseTimestamp** | Pointer to **int64** | The timestamp when the closure was triggered. | [optional] 
**Closing** | Pointer to **bool** | The problem is in process of closing (&#x60;true&#x60;) or closed (&#x60;false&#x60;). | [optional] 

## Methods

### NewProblemCloseResult

`func NewProblemCloseResult() *ProblemCloseResult`

NewProblemCloseResult instantiates a new ProblemCloseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemCloseResultWithDefaults

`func NewProblemCloseResultWithDefaults() *ProblemCloseResult`

NewProblemCloseResultWithDefaults instantiates a new ProblemCloseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProblemId

`func (o *ProblemCloseResult) GetProblemId() string`

GetProblemId returns the ProblemId field if non-nil, zero value otherwise.

### GetProblemIdOk

`func (o *ProblemCloseResult) GetProblemIdOk() (*string, bool)`

GetProblemIdOk returns a tuple with the ProblemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemId

`func (o *ProblemCloseResult) SetProblemId(v string)`

SetProblemId sets ProblemId field to given value.

### HasProblemId

`func (o *ProblemCloseResult) HasProblemId() bool`

HasProblemId returns a boolean if a field has been set.

### GetComment

`func (o *ProblemCloseResult) GetComment() ProblemComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ProblemCloseResult) GetCommentOk() (*ProblemComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ProblemCloseResult) SetComment(v ProblemComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ProblemCloseResult) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCloseTimestamp

`func (o *ProblemCloseResult) GetCloseTimestamp() int64`

GetCloseTimestamp returns the CloseTimestamp field if non-nil, zero value otherwise.

### GetCloseTimestampOk

`func (o *ProblemCloseResult) GetCloseTimestampOk() (*int64, bool)`

GetCloseTimestampOk returns a tuple with the CloseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTimestamp

`func (o *ProblemCloseResult) SetCloseTimestamp(v int64)`

SetCloseTimestamp sets CloseTimestamp field to given value.

### HasCloseTimestamp

`func (o *ProblemCloseResult) HasCloseTimestamp() bool`

HasCloseTimestamp returns a boolean if a field has been set.

### GetClosing

`func (o *ProblemCloseResult) GetClosing() bool`

GetClosing returns the Closing field if non-nil, zero value otherwise.

### GetClosingOk

`func (o *ProblemCloseResult) GetClosingOk() (*bool, bool)`

GetClosingOk returns a tuple with the Closing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosing

`func (o *ProblemCloseResult) SetClosing(v bool)`

SetClosing sets Closing field to given value.

### HasClosing

`func (o *ProblemCloseResult) HasClosing() bool`

HasClosing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


