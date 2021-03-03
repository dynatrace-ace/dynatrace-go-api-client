# LogFile4pg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The full path to the log. | [optional] 
**Size** | Pointer to **int64** | The size of the log, bytes | [optional] 
**Hosts** | Pointer to [**[]Host4pg**](Host4pg.md) | The distribution of the process group log across hosts. | [optional] 

## Methods

### NewLogFile4pg

`func NewLogFile4pg() *LogFile4pg`

NewLogFile4pg instantiates a new LogFile4pg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogFile4pgWithDefaults

`func NewLogFile4pgWithDefaults() *LogFile4pg`

NewLogFile4pgWithDefaults instantiates a new LogFile4pg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *LogFile4pg) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LogFile4pg) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LogFile4pg) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LogFile4pg) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSize

`func (o *LogFile4pg) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *LogFile4pg) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *LogFile4pg) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *LogFile4pg) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetHosts

`func (o *LogFile4pg) GetHosts() []Host4pg`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *LogFile4pg) GetHostsOk() (*[]Host4pg, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *LogFile4pg) SetHosts(v []Host4pg)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *LogFile4pg) HasHosts() bool`

HasHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


