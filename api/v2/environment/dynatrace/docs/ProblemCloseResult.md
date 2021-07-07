# ProblemCloseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloseTimestamp** | **int64** | The timestamp when the user triggered the closing. | 
**ProblemId** | **string** | The ID of the problem. | 
**Comment** | Pointer to [**Comment**](Comment.md) |  | [optional] 
**Closing** | **bool** | True, if the problem is being closed. | 

## Methods

### NewProblemCloseResult

`func NewProblemCloseResult(closeTimestamp int64, problemId string, closing bool, ) *ProblemCloseResult`

NewProblemCloseResult instantiates a new ProblemCloseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemCloseResultWithDefaults

`func NewProblemCloseResultWithDefaults() *ProblemCloseResult`

NewProblemCloseResultWithDefaults instantiates a new ProblemCloseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetComment

`func (o *ProblemCloseResult) GetComment() Comment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ProblemCloseResult) GetCommentOk() (*Comment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ProblemCloseResult) SetComment(v Comment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ProblemCloseResult) HasComment() bool`

HasComment returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


