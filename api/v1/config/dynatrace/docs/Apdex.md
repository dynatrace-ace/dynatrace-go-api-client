# Apdex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToleratedThreshold** | **int32** | Maximal length of an action, in milliseconds, which is considered as satisfied user experience.    You can use values between 100 and 60000. | [optional] 
**FrustratingThreshold** | **int32** | Maximal length of an action, in milliseconds, which is considered as tolerable user experience.    You can use values between 100 and 240000. | [optional] 
**ToleratedFallbackThreshold** | **int32** | Fallback threshold of an XHR action, in milliseconds, defining a satisfied user experience, when the configured KPM is not available.    Values between 100 and 60000 are allowed. | [optional] 
**FrustratingFallbackThreshold** | **int32** | Fallback threshold of an XHR action, in milliseconds, defining a tolerable user experience, when the configured KPM is not available.    Values between 100 and 240000 are allowed. | [optional] 
**ConsiderJavaScriptErrors** | **bool** | Consider JavaScript errors in Apdex calculations enabled/disabled | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


