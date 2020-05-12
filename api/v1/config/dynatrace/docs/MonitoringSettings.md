# MonitoringSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FetchRequests** | **bool** | &#x60;fetch()&#x60; request capture enabled/disabled. | 
**XmlHttpRequest** | **bool** | &#x60;XmlHttpRequest&#x60; support enabled/disabled. | 
**JavaScriptFrameworkSupport** | [**JavaScriptFrameworkSupport**](JavaScriptFrameworkSupport.md) |  | 
**ContentCapture** | [**ContentCapture**](ContentCapture.md) |  | 
**ExcludeXhrRegex** | **string** | You can exclude some actions from becoming XHR actions.  Put a regular expression, matching all the required URLs, here.  If noting specified the feature is disabled. | 
**InjectionMode** | **string** | JavaScript injection mode. | 
**AddCrossOriginAnonymousAttribute** | **bool** | Add the cross origin &#x3D; anonymous attribute to capture JavaScript error messages and W3C resource timings. | [optional] 
**ScriptTagCacheDurationInHours** | **int32** | Time duration for the cache settings. | [optional] 
**LibraryFileLocation** | **string** | The location of your application’s custom JavaScript library file.  If nothing specified the root directory of your web server is used. | [optional] 
**MonitoringDataPath** | **string** | The location to send monitoring data from the JavaScript tag.  Specify either a relative or an absolute URL. If you use an absolute URL, data will be sent using CORS. | [optional] 
**CustomConfigurationProperties** | **string** | Additional JavaScript tag properties that are specific to your application. To do this, type key&#x3D;value pairs separated using a (|) symbol. | 
**ServerRequestPathId** | **string** | Path to identify the server’s request ID. | 
**SecureCookieAttribute** | **bool** | Secure attribute usage for Dynatrace cookies enabled/disabled. | 
**CookiePlacementDomain** | **string** | Domain for cookie placement. | 
**CacheControlHeaderOptimizations** | **bool** | Optimize the value of cache control headers for use with Dynatrace real user monitoring enabled/disabled. | 
**AdvancedJavaScriptTagSettings** | [**AdvancedJavaScriptTagSettings**](AdvancedJavaScriptTagSettings.md) |  | 
**BrowserRestrictionSettings** | [**WebApplicationConfigBrowserRestrictionSettings**](WebApplicationConfigBrowserRestrictionSettings.md) |  | [optional] 
**IpAddressRestrictionSettings** | [**WebApplicationConfigIpAddressRestrictionSettings**](WebApplicationConfigIpAddressRestrictionSettings.md) |  | [optional] 
**JavaScriptInjectionRules** | [**[]JavaScriptInjectionRules**](JavaScriptInjectionRules.md) | Java script injection rules. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


