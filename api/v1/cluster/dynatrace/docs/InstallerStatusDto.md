# InstallerStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 

## Methods

### NewInstallerStatusDto

`func NewInstallerStatusDto() *InstallerStatusDto`

NewInstallerStatusDto instantiates a new InstallerStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallerStatusDtoWithDefaults

`func NewInstallerStatusDtoWithDefaults() *InstallerStatusDto`

NewInstallerStatusDtoWithDefaults instantiates a new InstallerStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InstallerStatusDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstallerStatusDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstallerStatusDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstallerStatusDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetails

`func (o *InstallerStatusDto) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InstallerStatusDto) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InstallerStatusDto) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InstallerStatusDto) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


