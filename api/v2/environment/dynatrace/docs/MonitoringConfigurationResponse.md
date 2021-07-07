# MonitoringConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | Pointer to **string** | The identifier of the new configuration | [optional] 
**Code** | Pointer to **int32** | The HTTP Status code | [optional] 

## Methods

### NewMonitoringConfigurationResponse

`func NewMonitoringConfigurationResponse() *MonitoringConfigurationResponse`

NewMonitoringConfigurationResponse instantiates a new MonitoringConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringConfigurationResponseWithDefaults

`func NewMonitoringConfigurationResponseWithDefaults() *MonitoringConfigurationResponse`

NewMonitoringConfigurationResponseWithDefaults instantiates a new MonitoringConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *MonitoringConfigurationResponse) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *MonitoringConfigurationResponse) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *MonitoringConfigurationResponse) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *MonitoringConfigurationResponse) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetCode

`func (o *MonitoringConfigurationResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MonitoringConfigurationResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MonitoringConfigurationResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *MonitoringConfigurationResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


