# ApplicationBaselineValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The Dynatrace entity ID of the application. | 
**DisplayName** | Pointer to **string** | The name of the application as displayed in the UI. | [optional] 
**ApplicationDomInteractiveBaselines** | Pointer to [**[]EntityBaselineData**](EntityBaselineData.md) | The baseline data for the **DOM interactive** duration metric. | [optional] 
**ApplicationHtmlDownloadedBaselines** | Pointer to [**[]EntityBaselineData**](EntityBaselineData.md) | The baseline data for the **HTML downloaded** duration metric. | [optional] 
**ApplicationLoadEventEndBaselines** | Pointer to [**[]EntityBaselineData**](EntityBaselineData.md) | The baseline data for the **Load event end** duration metric. | [optional] 
**ApplicationLoadEventStartBaselines** | Pointer to [**[]EntityBaselineData**](EntityBaselineData.md) | The baseline data for the **Load event start** duration metric. | [optional] 
**ApplicationResponseTimeBaselines** | Pointer to [**[]EntityBaselineData**](EntityBaselineData.md) | The baseline data for the **Response time** duration metric. | [optional] 
**ApplicationSpeedIndexBaselines** | Pointer to [**[]EntityBaselineData**](EntityBaselineData.md) | The baseline data for the **Speed index** duration metric. | [optional] 
**ApplicationTimeToFirstByteBaselines** | Pointer to [**[]EntityBaselineData**](EntityBaselineData.md) | The baseline data for the **Time to first byte** duration metric. | [optional] 
**ApplicationVisualCompleteBaselines** | Pointer to [**[]EntityBaselineData**](EntityBaselineData.md) | The baseline data for the **Visually complete** duration metric. | [optional] 

## Methods

### NewApplicationBaselineValues

`func NewApplicationBaselineValues(entityId string, ) *ApplicationBaselineValues`

NewApplicationBaselineValues instantiates a new ApplicationBaselineValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationBaselineValuesWithDefaults

`func NewApplicationBaselineValuesWithDefaults() *ApplicationBaselineValues`

NewApplicationBaselineValuesWithDefaults instantiates a new ApplicationBaselineValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *ApplicationBaselineValues) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ApplicationBaselineValues) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ApplicationBaselineValues) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetDisplayName

`func (o *ApplicationBaselineValues) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApplicationBaselineValues) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApplicationBaselineValues) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApplicationBaselineValues) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetApplicationDomInteractiveBaselines

`func (o *ApplicationBaselineValues) GetApplicationDomInteractiveBaselines() []EntityBaselineData`

GetApplicationDomInteractiveBaselines returns the ApplicationDomInteractiveBaselines field if non-nil, zero value otherwise.

### GetApplicationDomInteractiveBaselinesOk

`func (o *ApplicationBaselineValues) GetApplicationDomInteractiveBaselinesOk() (*[]EntityBaselineData, bool)`

GetApplicationDomInteractiveBaselinesOk returns a tuple with the ApplicationDomInteractiveBaselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDomInteractiveBaselines

`func (o *ApplicationBaselineValues) SetApplicationDomInteractiveBaselines(v []EntityBaselineData)`

SetApplicationDomInteractiveBaselines sets ApplicationDomInteractiveBaselines field to given value.

### HasApplicationDomInteractiveBaselines

`func (o *ApplicationBaselineValues) HasApplicationDomInteractiveBaselines() bool`

HasApplicationDomInteractiveBaselines returns a boolean if a field has been set.

### GetApplicationHtmlDownloadedBaselines

`func (o *ApplicationBaselineValues) GetApplicationHtmlDownloadedBaselines() []EntityBaselineData`

GetApplicationHtmlDownloadedBaselines returns the ApplicationHtmlDownloadedBaselines field if non-nil, zero value otherwise.

### GetApplicationHtmlDownloadedBaselinesOk

`func (o *ApplicationBaselineValues) GetApplicationHtmlDownloadedBaselinesOk() (*[]EntityBaselineData, bool)`

GetApplicationHtmlDownloadedBaselinesOk returns a tuple with the ApplicationHtmlDownloadedBaselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationHtmlDownloadedBaselines

`func (o *ApplicationBaselineValues) SetApplicationHtmlDownloadedBaselines(v []EntityBaselineData)`

SetApplicationHtmlDownloadedBaselines sets ApplicationHtmlDownloadedBaselines field to given value.

### HasApplicationHtmlDownloadedBaselines

`func (o *ApplicationBaselineValues) HasApplicationHtmlDownloadedBaselines() bool`

HasApplicationHtmlDownloadedBaselines returns a boolean if a field has been set.

### GetApplicationLoadEventEndBaselines

`func (o *ApplicationBaselineValues) GetApplicationLoadEventEndBaselines() []EntityBaselineData`

GetApplicationLoadEventEndBaselines returns the ApplicationLoadEventEndBaselines field if non-nil, zero value otherwise.

### GetApplicationLoadEventEndBaselinesOk

`func (o *ApplicationBaselineValues) GetApplicationLoadEventEndBaselinesOk() (*[]EntityBaselineData, bool)`

GetApplicationLoadEventEndBaselinesOk returns a tuple with the ApplicationLoadEventEndBaselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationLoadEventEndBaselines

`func (o *ApplicationBaselineValues) SetApplicationLoadEventEndBaselines(v []EntityBaselineData)`

SetApplicationLoadEventEndBaselines sets ApplicationLoadEventEndBaselines field to given value.

### HasApplicationLoadEventEndBaselines

`func (o *ApplicationBaselineValues) HasApplicationLoadEventEndBaselines() bool`

HasApplicationLoadEventEndBaselines returns a boolean if a field has been set.

### GetApplicationLoadEventStartBaselines

`func (o *ApplicationBaselineValues) GetApplicationLoadEventStartBaselines() []EntityBaselineData`

GetApplicationLoadEventStartBaselines returns the ApplicationLoadEventStartBaselines field if non-nil, zero value otherwise.

### GetApplicationLoadEventStartBaselinesOk

`func (o *ApplicationBaselineValues) GetApplicationLoadEventStartBaselinesOk() (*[]EntityBaselineData, bool)`

GetApplicationLoadEventStartBaselinesOk returns a tuple with the ApplicationLoadEventStartBaselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationLoadEventStartBaselines

`func (o *ApplicationBaselineValues) SetApplicationLoadEventStartBaselines(v []EntityBaselineData)`

SetApplicationLoadEventStartBaselines sets ApplicationLoadEventStartBaselines field to given value.

### HasApplicationLoadEventStartBaselines

`func (o *ApplicationBaselineValues) HasApplicationLoadEventStartBaselines() bool`

HasApplicationLoadEventStartBaselines returns a boolean if a field has been set.

### GetApplicationResponseTimeBaselines

`func (o *ApplicationBaselineValues) GetApplicationResponseTimeBaselines() []EntityBaselineData`

GetApplicationResponseTimeBaselines returns the ApplicationResponseTimeBaselines field if non-nil, zero value otherwise.

### GetApplicationResponseTimeBaselinesOk

`func (o *ApplicationBaselineValues) GetApplicationResponseTimeBaselinesOk() (*[]EntityBaselineData, bool)`

GetApplicationResponseTimeBaselinesOk returns a tuple with the ApplicationResponseTimeBaselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationResponseTimeBaselines

`func (o *ApplicationBaselineValues) SetApplicationResponseTimeBaselines(v []EntityBaselineData)`

SetApplicationResponseTimeBaselines sets ApplicationResponseTimeBaselines field to given value.

### HasApplicationResponseTimeBaselines

`func (o *ApplicationBaselineValues) HasApplicationResponseTimeBaselines() bool`

HasApplicationResponseTimeBaselines returns a boolean if a field has been set.

### GetApplicationSpeedIndexBaselines

`func (o *ApplicationBaselineValues) GetApplicationSpeedIndexBaselines() []EntityBaselineData`

GetApplicationSpeedIndexBaselines returns the ApplicationSpeedIndexBaselines field if non-nil, zero value otherwise.

### GetApplicationSpeedIndexBaselinesOk

`func (o *ApplicationBaselineValues) GetApplicationSpeedIndexBaselinesOk() (*[]EntityBaselineData, bool)`

GetApplicationSpeedIndexBaselinesOk returns a tuple with the ApplicationSpeedIndexBaselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSpeedIndexBaselines

`func (o *ApplicationBaselineValues) SetApplicationSpeedIndexBaselines(v []EntityBaselineData)`

SetApplicationSpeedIndexBaselines sets ApplicationSpeedIndexBaselines field to given value.

### HasApplicationSpeedIndexBaselines

`func (o *ApplicationBaselineValues) HasApplicationSpeedIndexBaselines() bool`

HasApplicationSpeedIndexBaselines returns a boolean if a field has been set.

### GetApplicationTimeToFirstByteBaselines

`func (o *ApplicationBaselineValues) GetApplicationTimeToFirstByteBaselines() []EntityBaselineData`

GetApplicationTimeToFirstByteBaselines returns the ApplicationTimeToFirstByteBaselines field if non-nil, zero value otherwise.

### GetApplicationTimeToFirstByteBaselinesOk

`func (o *ApplicationBaselineValues) GetApplicationTimeToFirstByteBaselinesOk() (*[]EntityBaselineData, bool)`

GetApplicationTimeToFirstByteBaselinesOk returns a tuple with the ApplicationTimeToFirstByteBaselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTimeToFirstByteBaselines

`func (o *ApplicationBaselineValues) SetApplicationTimeToFirstByteBaselines(v []EntityBaselineData)`

SetApplicationTimeToFirstByteBaselines sets ApplicationTimeToFirstByteBaselines field to given value.

### HasApplicationTimeToFirstByteBaselines

`func (o *ApplicationBaselineValues) HasApplicationTimeToFirstByteBaselines() bool`

HasApplicationTimeToFirstByteBaselines returns a boolean if a field has been set.

### GetApplicationVisualCompleteBaselines

`func (o *ApplicationBaselineValues) GetApplicationVisualCompleteBaselines() []EntityBaselineData`

GetApplicationVisualCompleteBaselines returns the ApplicationVisualCompleteBaselines field if non-nil, zero value otherwise.

### GetApplicationVisualCompleteBaselinesOk

`func (o *ApplicationBaselineValues) GetApplicationVisualCompleteBaselinesOk() (*[]EntityBaselineData, bool)`

GetApplicationVisualCompleteBaselinesOk returns a tuple with the ApplicationVisualCompleteBaselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVisualCompleteBaselines

`func (o *ApplicationBaselineValues) SetApplicationVisualCompleteBaselines(v []EntityBaselineData)`

SetApplicationVisualCompleteBaselines sets ApplicationVisualCompleteBaselines field to given value.

### HasApplicationVisualCompleteBaselines

`func (o *ApplicationBaselineValues) HasApplicationVisualCompleteBaselines() bool`

HasApplicationVisualCompleteBaselines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


