# ApplicationDataPrivacy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Identifier** | **string** | Dynatrace entity ID of the web application. | [optional] [readonly] 
**DataCaptureOptInEnabled** | **bool** | Set to &#x60;true&#x60; to disable data capture and cookies until JavaScriptAPI dtrum.enable() is called. | 
**PersistentCookieForUserTracking** | **bool** | Set to &#x60;true&#x60; to set persistent cookie in order to recognize returning devices. | 
**DoNotTrackBehaviour** | **string** | How to handle browsers&#39; \&quot;Do Not Track\&quot;: &#x60;IGNORE_DO_NOT_TRACK&#x60;, or &#x60;CAPTURE_ANONYMIZED&#x60;, or &#x60;DO_NOT_CAPTURE&#x60;. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


