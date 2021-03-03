# Host4pg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | The entity ID of the host. | [optional] 
**LogSize** | Pointer to **int64** | The size of the PG log for the host, bytes. | [optional] 
**ContentAccess** | Pointer to **bool** | The access to the log content is granted (true) or denied (false). | [optional] 
**AvailableForAnalysis** | Pointer to **bool** | The log is available (true) or not available (false) for analysis. | [optional] 

## Methods

### NewHost4pg

`func NewHost4pg() *Host4pg`

NewHost4pg instantiates a new Host4pg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHost4pgWithDefaults

`func NewHost4pgWithDefaults() *Host4pg`

NewHost4pgWithDefaults instantiates a new Host4pg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *Host4pg) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *Host4pg) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *Host4pg) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *Host4pg) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetLogSize

`func (o *Host4pg) GetLogSize() int64`

GetLogSize returns the LogSize field if non-nil, zero value otherwise.

### GetLogSizeOk

`func (o *Host4pg) GetLogSizeOk() (*int64, bool)`

GetLogSizeOk returns a tuple with the LogSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSize

`func (o *Host4pg) SetLogSize(v int64)`

SetLogSize sets LogSize field to given value.

### HasLogSize

`func (o *Host4pg) HasLogSize() bool`

HasLogSize returns a boolean if a field has been set.

### GetContentAccess

`func (o *Host4pg) GetContentAccess() bool`

GetContentAccess returns the ContentAccess field if non-nil, zero value otherwise.

### GetContentAccessOk

`func (o *Host4pg) GetContentAccessOk() (*bool, bool)`

GetContentAccessOk returns a tuple with the ContentAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAccess

`func (o *Host4pg) SetContentAccess(v bool)`

SetContentAccess sets ContentAccess field to given value.

### HasContentAccess

`func (o *Host4pg) HasContentAccess() bool`

HasContentAccess returns a boolean if a field has been set.

### GetAvailableForAnalysis

`func (o *Host4pg) GetAvailableForAnalysis() bool`

GetAvailableForAnalysis returns the AvailableForAnalysis field if non-nil, zero value otherwise.

### GetAvailableForAnalysisOk

`func (o *Host4pg) GetAvailableForAnalysisOk() (*bool, bool)`

GetAvailableForAnalysisOk returns a tuple with the AvailableForAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForAnalysis

`func (o *Host4pg) SetAvailableForAnalysis(v bool)`

SetAvailableForAnalysis sets AvailableForAnalysis field to given value.

### HasAvailableForAnalysis

`func (o *Host4pg) HasAvailableForAnalysis() bool`

HasAvailableForAnalysis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


