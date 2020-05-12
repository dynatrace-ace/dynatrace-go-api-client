# WebApplicationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Identifier** | **string** | Dynatrace entity ID of the web application. | [optional] [readonly] 
**Name** | **string** | The name of the web application, displayed in the UI. | 
**Type** | **string** | The type of the web application. | [optional] 
**RealUserMonitoringEnabled** | **bool** | Real user monitoring enabled/disabled. | 
**CostControlUserSessionPercentage** | **float32** | Analize *X*% of user sessions. | 
**LoadActionKeyPerformanceMetric** | **string** | The key performance metric of load actions. | 
**XhrActionKeyPerformanceMetric** | **string** | The key performance metric of XHR actions. | 
**LoadActionApdexSettings** | [**Apdex**](Apdex.md) |  | 
**XhrActionApdexSettings** | [**Apdex**](Apdex.md) |  | 
**CustomActionApdexSettings** | [**Apdex**](Apdex.md) |  | 
**WaterfallSettings** | [**WaterfallSettings**](WaterfallSettings.md) |  | 
**MonitoringSettings** | [**MonitoringSettings**](MonitoringSettings.md) |  | 
**UserActionNamingSettings** | [**UserActionNamingSettings**](UserActionNamingSettings.md) |  | [optional] 
**MetaDataCaptureSettings** | [**[]MetaDataCapturing**](MetaDataCapturing.md) | Java script agent meta data capture settings. | [optional] 
**ConversionGoals** | [**[]ConversionGoal**](ConversionGoal.md) | A list of conversion goals of the application. | [optional] 
**UrlInjectionPattern** | **string** | Url injection pattern for manual web application. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


