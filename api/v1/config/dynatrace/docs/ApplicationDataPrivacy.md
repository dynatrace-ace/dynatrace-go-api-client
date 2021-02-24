# ApplicationDataPrivacy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Identifier** | Pointer to **string** | Dynatrace entity ID of the web application. | [optional] [readonly] 
**DataCaptureOptInEnabled** | **bool** | Set to &#x60;true&#x60; to disable data capture and cookies until JavaScriptAPI dtrum.enable() is called. | 
**PersistentCookieForUserTracking** | **bool** | Set to &#x60;true&#x60; to set persistent cookie in order to recognize returning devices. | 
**DoNotTrackBehaviour** | **string** | How to handle browsers&#39; \&quot;Do Not Track\&quot;: &#x60;IGNORE_DO_NOT_TRACK&#x60;, or &#x60;CAPTURE_ANONYMIZED&#x60;, or &#x60;DO_NOT_CAPTURE&#x60;. | 

## Methods

### NewApplicationDataPrivacy

`func NewApplicationDataPrivacy(dataCaptureOptInEnabled bool, persistentCookieForUserTracking bool, doNotTrackBehaviour string, ) *ApplicationDataPrivacy`

NewApplicationDataPrivacy instantiates a new ApplicationDataPrivacy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDataPrivacyWithDefaults

`func NewApplicationDataPrivacyWithDefaults() *ApplicationDataPrivacy`

NewApplicationDataPrivacyWithDefaults instantiates a new ApplicationDataPrivacy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ApplicationDataPrivacy) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationDataPrivacy) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationDataPrivacy) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationDataPrivacy) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetIdentifier

`func (o *ApplicationDataPrivacy) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ApplicationDataPrivacy) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ApplicationDataPrivacy) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ApplicationDataPrivacy) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetDataCaptureOptInEnabled

`func (o *ApplicationDataPrivacy) GetDataCaptureOptInEnabled() bool`

GetDataCaptureOptInEnabled returns the DataCaptureOptInEnabled field if non-nil, zero value otherwise.

### GetDataCaptureOptInEnabledOk

`func (o *ApplicationDataPrivacy) GetDataCaptureOptInEnabledOk() (*bool, bool)`

GetDataCaptureOptInEnabledOk returns a tuple with the DataCaptureOptInEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCaptureOptInEnabled

`func (o *ApplicationDataPrivacy) SetDataCaptureOptInEnabled(v bool)`

SetDataCaptureOptInEnabled sets DataCaptureOptInEnabled field to given value.


### GetPersistentCookieForUserTracking

`func (o *ApplicationDataPrivacy) GetPersistentCookieForUserTracking() bool`

GetPersistentCookieForUserTracking returns the PersistentCookieForUserTracking field if non-nil, zero value otherwise.

### GetPersistentCookieForUserTrackingOk

`func (o *ApplicationDataPrivacy) GetPersistentCookieForUserTrackingOk() (*bool, bool)`

GetPersistentCookieForUserTrackingOk returns a tuple with the PersistentCookieForUserTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentCookieForUserTracking

`func (o *ApplicationDataPrivacy) SetPersistentCookieForUserTracking(v bool)`

SetPersistentCookieForUserTracking sets PersistentCookieForUserTracking field to given value.


### GetDoNotTrackBehaviour

`func (o *ApplicationDataPrivacy) GetDoNotTrackBehaviour() string`

GetDoNotTrackBehaviour returns the DoNotTrackBehaviour field if non-nil, zero value otherwise.

### GetDoNotTrackBehaviourOk

`func (o *ApplicationDataPrivacy) GetDoNotTrackBehaviourOk() (*string, bool)`

GetDoNotTrackBehaviourOk returns a tuple with the DoNotTrackBehaviour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrackBehaviour

`func (o *ApplicationDataPrivacy) SetDoNotTrackBehaviour(v string)`

SetDoNotTrackBehaviour sets DoNotTrackBehaviour field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


