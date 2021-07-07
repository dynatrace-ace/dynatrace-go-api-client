# SupportArchiveDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadState** | Pointer to **string** |  | [optional] 
**Report** | Pointer to [**SupportArchiveReport**](SupportArchiveReport.md) |  | [optional] 

## Methods

### NewSupportArchiveDownload

`func NewSupportArchiveDownload() *SupportArchiveDownload`

NewSupportArchiveDownload instantiates a new SupportArchiveDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportArchiveDownloadWithDefaults

`func NewSupportArchiveDownloadWithDefaults() *SupportArchiveDownload`

NewSupportArchiveDownloadWithDefaults instantiates a new SupportArchiveDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadState

`func (o *SupportArchiveDownload) GetDownloadState() string`

GetDownloadState returns the DownloadState field if non-nil, zero value otherwise.

### GetDownloadStateOk

`func (o *SupportArchiveDownload) GetDownloadStateOk() (*string, bool)`

GetDownloadStateOk returns a tuple with the DownloadState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadState

`func (o *SupportArchiveDownload) SetDownloadState(v string)`

SetDownloadState sets DownloadState field to given value.

### HasDownloadState

`func (o *SupportArchiveDownload) HasDownloadState() bool`

HasDownloadState returns a boolean if a field has been set.

### GetReport

`func (o *SupportArchiveDownload) GetReport() SupportArchiveReport`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *SupportArchiveDownload) GetReportOk() (*SupportArchiveReport, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *SupportArchiveDownload) SetReport(v SupportArchiveReport)`

SetReport sets Report field to given value.

### HasReport

`func (o *SupportArchiveDownload) HasReport() bool`

HasReport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


