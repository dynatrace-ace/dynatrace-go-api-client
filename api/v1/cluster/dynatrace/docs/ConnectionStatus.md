# ConnectionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionOk** | Pointer to **bool** | Connection test result | [optional] 
**TestExecuted** | Pointer to **bool** | Indicates whether test was executed at all | [optional] 
**TestExecutionMessage** | Pointer to **string** | Additional comments usually indicates why test was not executed | [optional] 

## Methods

### NewConnectionStatus

`func NewConnectionStatus() *ConnectionStatus`

NewConnectionStatus instantiates a new ConnectionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionStatusWithDefaults

`func NewConnectionStatusWithDefaults() *ConnectionStatus`

NewConnectionStatusWithDefaults instantiates a new ConnectionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionOk

`func (o *ConnectionStatus) GetConnectionOk() bool`

GetConnectionOk returns the ConnectionOk field if non-nil, zero value otherwise.

### GetConnectionOkOk

`func (o *ConnectionStatus) GetConnectionOkOk() (*bool, bool)`

GetConnectionOkOk returns a tuple with the ConnectionOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionOk

`func (o *ConnectionStatus) SetConnectionOk(v bool)`

SetConnectionOk sets ConnectionOk field to given value.

### HasConnectionOk

`func (o *ConnectionStatus) HasConnectionOk() bool`

HasConnectionOk returns a boolean if a field has been set.

### GetTestExecuted

`func (o *ConnectionStatus) GetTestExecuted() bool`

GetTestExecuted returns the TestExecuted field if non-nil, zero value otherwise.

### GetTestExecutedOk

`func (o *ConnectionStatus) GetTestExecutedOk() (*bool, bool)`

GetTestExecutedOk returns a tuple with the TestExecuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestExecuted

`func (o *ConnectionStatus) SetTestExecuted(v bool)`

SetTestExecuted sets TestExecuted field to given value.

### HasTestExecuted

`func (o *ConnectionStatus) HasTestExecuted() bool`

HasTestExecuted returns a boolean if a field has been set.

### GetTestExecutionMessage

`func (o *ConnectionStatus) GetTestExecutionMessage() string`

GetTestExecutionMessage returns the TestExecutionMessage field if non-nil, zero value otherwise.

### GetTestExecutionMessageOk

`func (o *ConnectionStatus) GetTestExecutionMessageOk() (*string, bool)`

GetTestExecutionMessageOk returns a tuple with the TestExecutionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestExecutionMessage

`func (o *ConnectionStatus) SetTestExecutionMessage(v string)`

SetTestExecutionMessage sets TestExecutionMessage field to given value.

### HasTestExecutionMessage

`func (o *ConnectionStatus) HasTestExecutionMessage() bool`

HasTestExecutionMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


