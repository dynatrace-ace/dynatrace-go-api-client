# MonitoringSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FetchRequests** | **bool** | &#x60;fetch()&#x60; request capture enabled/disabled. | 
**XmlHttpRequest** | **bool** | &#x60;XmlHttpRequest&#x60; support enabled/disabled. | 
**JavaScriptFrameworkSupport** | [**JavaScriptFrameworkSupport**](JavaScriptFrameworkSupport.md) |  | 
**ContentCapture** | [**ContentCapture**](ContentCapture.md) |  | 
**ExcludeXhrRegex** | **string** | You can exclude some actions from becoming XHR actions.  Put a regular expression, matching all the required URLs, here.  If noting specified the feature is disabled. | 
**CorrelationHeaderInclusionRegex** | Pointer to **string** | To enable RUM for XHR calls to AWS Lambda, define a regular expression matching these calls, Dynatrace can then automatically add a custom header (x-dtc) to each such request to the respective endpoints in AWS.  Important: These endpoints must accept the x-dtc header, or the requests will fail.  | [optional] 
**InjectionMode** | **string** | JavaScript injection mode. | 
**AddCrossOriginAnonymousAttribute** | Pointer to **bool** | Add the cross origin &#x3D; anonymous attribute to capture JavaScript error messages and W3C resource timings. | [optional] 
**ScriptTagCacheDurationInHours** | Pointer to **int32** | Time duration for the cache settings. | [optional] 
**LibraryFileLocation** | Pointer to **string** | The location of your application’s custom JavaScript library file.  If nothing specified the root directory of your web server is used.  Only supported by auto-injected applications. | [optional] 
**MonitoringDataPath** | Pointer to **string** | The location to send monitoring data from the JavaScript tag.  Specify either a relative or an absolute URL. If you use an absolute URL, data will be sent using CORS.  Only supported by auto-injected applications. | [optional] 
**CustomConfigurationProperties** | **string** | Additional JavaScript tag properties that are specific to your application. To do this, type key&#x3D;value pairs separated using a (|) symbol. | 
**ServerRequestPathId** | **string** | Path to identify the server’s request ID. | 
**SecureCookieAttribute** | **bool** | Secure attribute usage for Dynatrace cookies enabled/disabled. | 
**CookiePlacementDomain** | **string** | Domain for cookie placement. | 
**CacheControlHeaderOptimizations** | **bool** | Optimize the value of cache control headers for use with Dynatrace real user monitoring enabled/disabled. | 
**AdvancedJavaScriptTagSettings** | [**AdvancedJavaScriptTagSettings**](AdvancedJavaScriptTagSettings.md) |  | 
**BrowserRestrictionSettings** | Pointer to [**WebApplicationConfigBrowserRestrictionSettings**](WebApplicationConfigBrowserRestrictionSettings.md) |  | [optional] 
**IpAddressRestrictionSettings** | Pointer to [**WebApplicationConfigIpAddressRestrictionSettings**](WebApplicationConfigIpAddressRestrictionSettings.md) |  | [optional] 
**JavaScriptInjectionRules** | Pointer to [**[]JavaScriptInjectionRules**](JavaScriptInjectionRules.md) | Java script injection rules. | [optional] 

## Methods

### NewMonitoringSettings

`func NewMonitoringSettings(fetchRequests bool, xmlHttpRequest bool, javaScriptFrameworkSupport JavaScriptFrameworkSupport, contentCapture ContentCapture, excludeXhrRegex string, injectionMode string, customConfigurationProperties string, serverRequestPathId string, secureCookieAttribute bool, cookiePlacementDomain string, cacheControlHeaderOptimizations bool, advancedJavaScriptTagSettings AdvancedJavaScriptTagSettings, ) *MonitoringSettings`

NewMonitoringSettings instantiates a new MonitoringSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringSettingsWithDefaults

`func NewMonitoringSettingsWithDefaults() *MonitoringSettings`

NewMonitoringSettingsWithDefaults instantiates a new MonitoringSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFetchRequests

`func (o *MonitoringSettings) GetFetchRequests() bool`

GetFetchRequests returns the FetchRequests field if non-nil, zero value otherwise.

### GetFetchRequestsOk

`func (o *MonitoringSettings) GetFetchRequestsOk() (*bool, bool)`

GetFetchRequestsOk returns a tuple with the FetchRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchRequests

`func (o *MonitoringSettings) SetFetchRequests(v bool)`

SetFetchRequests sets FetchRequests field to given value.


### GetXmlHttpRequest

`func (o *MonitoringSettings) GetXmlHttpRequest() bool`

GetXmlHttpRequest returns the XmlHttpRequest field if non-nil, zero value otherwise.

### GetXmlHttpRequestOk

`func (o *MonitoringSettings) GetXmlHttpRequestOk() (*bool, bool)`

GetXmlHttpRequestOk returns a tuple with the XmlHttpRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlHttpRequest

`func (o *MonitoringSettings) SetXmlHttpRequest(v bool)`

SetXmlHttpRequest sets XmlHttpRequest field to given value.


### GetJavaScriptFrameworkSupport

`func (o *MonitoringSettings) GetJavaScriptFrameworkSupport() JavaScriptFrameworkSupport`

GetJavaScriptFrameworkSupport returns the JavaScriptFrameworkSupport field if non-nil, zero value otherwise.

### GetJavaScriptFrameworkSupportOk

`func (o *MonitoringSettings) GetJavaScriptFrameworkSupportOk() (*JavaScriptFrameworkSupport, bool)`

GetJavaScriptFrameworkSupportOk returns a tuple with the JavaScriptFrameworkSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaScriptFrameworkSupport

`func (o *MonitoringSettings) SetJavaScriptFrameworkSupport(v JavaScriptFrameworkSupport)`

SetJavaScriptFrameworkSupport sets JavaScriptFrameworkSupport field to given value.


### GetContentCapture

`func (o *MonitoringSettings) GetContentCapture() ContentCapture`

GetContentCapture returns the ContentCapture field if non-nil, zero value otherwise.

### GetContentCaptureOk

`func (o *MonitoringSettings) GetContentCaptureOk() (*ContentCapture, bool)`

GetContentCaptureOk returns a tuple with the ContentCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCapture

`func (o *MonitoringSettings) SetContentCapture(v ContentCapture)`

SetContentCapture sets ContentCapture field to given value.


### GetExcludeXhrRegex

`func (o *MonitoringSettings) GetExcludeXhrRegex() string`

GetExcludeXhrRegex returns the ExcludeXhrRegex field if non-nil, zero value otherwise.

### GetExcludeXhrRegexOk

`func (o *MonitoringSettings) GetExcludeXhrRegexOk() (*string, bool)`

GetExcludeXhrRegexOk returns a tuple with the ExcludeXhrRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeXhrRegex

`func (o *MonitoringSettings) SetExcludeXhrRegex(v string)`

SetExcludeXhrRegex sets ExcludeXhrRegex field to given value.


### GetCorrelationHeaderInclusionRegex

`func (o *MonitoringSettings) GetCorrelationHeaderInclusionRegex() string`

GetCorrelationHeaderInclusionRegex returns the CorrelationHeaderInclusionRegex field if non-nil, zero value otherwise.

### GetCorrelationHeaderInclusionRegexOk

`func (o *MonitoringSettings) GetCorrelationHeaderInclusionRegexOk() (*string, bool)`

GetCorrelationHeaderInclusionRegexOk returns a tuple with the CorrelationHeaderInclusionRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationHeaderInclusionRegex

`func (o *MonitoringSettings) SetCorrelationHeaderInclusionRegex(v string)`

SetCorrelationHeaderInclusionRegex sets CorrelationHeaderInclusionRegex field to given value.

### HasCorrelationHeaderInclusionRegex

`func (o *MonitoringSettings) HasCorrelationHeaderInclusionRegex() bool`

HasCorrelationHeaderInclusionRegex returns a boolean if a field has been set.

### GetInjectionMode

`func (o *MonitoringSettings) GetInjectionMode() string`

GetInjectionMode returns the InjectionMode field if non-nil, zero value otherwise.

### GetInjectionModeOk

`func (o *MonitoringSettings) GetInjectionModeOk() (*string, bool)`

GetInjectionModeOk returns a tuple with the InjectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectionMode

`func (o *MonitoringSettings) SetInjectionMode(v string)`

SetInjectionMode sets InjectionMode field to given value.


### GetAddCrossOriginAnonymousAttribute

`func (o *MonitoringSettings) GetAddCrossOriginAnonymousAttribute() bool`

GetAddCrossOriginAnonymousAttribute returns the AddCrossOriginAnonymousAttribute field if non-nil, zero value otherwise.

### GetAddCrossOriginAnonymousAttributeOk

`func (o *MonitoringSettings) GetAddCrossOriginAnonymousAttributeOk() (*bool, bool)`

GetAddCrossOriginAnonymousAttributeOk returns a tuple with the AddCrossOriginAnonymousAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCrossOriginAnonymousAttribute

`func (o *MonitoringSettings) SetAddCrossOriginAnonymousAttribute(v bool)`

SetAddCrossOriginAnonymousAttribute sets AddCrossOriginAnonymousAttribute field to given value.

### HasAddCrossOriginAnonymousAttribute

`func (o *MonitoringSettings) HasAddCrossOriginAnonymousAttribute() bool`

HasAddCrossOriginAnonymousAttribute returns a boolean if a field has been set.

### GetScriptTagCacheDurationInHours

`func (o *MonitoringSettings) GetScriptTagCacheDurationInHours() int32`

GetScriptTagCacheDurationInHours returns the ScriptTagCacheDurationInHours field if non-nil, zero value otherwise.

### GetScriptTagCacheDurationInHoursOk

`func (o *MonitoringSettings) GetScriptTagCacheDurationInHoursOk() (*int32, bool)`

GetScriptTagCacheDurationInHoursOk returns a tuple with the ScriptTagCacheDurationInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptTagCacheDurationInHours

`func (o *MonitoringSettings) SetScriptTagCacheDurationInHours(v int32)`

SetScriptTagCacheDurationInHours sets ScriptTagCacheDurationInHours field to given value.

### HasScriptTagCacheDurationInHours

`func (o *MonitoringSettings) HasScriptTagCacheDurationInHours() bool`

HasScriptTagCacheDurationInHours returns a boolean if a field has been set.

### GetLibraryFileLocation

`func (o *MonitoringSettings) GetLibraryFileLocation() string`

GetLibraryFileLocation returns the LibraryFileLocation field if non-nil, zero value otherwise.

### GetLibraryFileLocationOk

`func (o *MonitoringSettings) GetLibraryFileLocationOk() (*string, bool)`

GetLibraryFileLocationOk returns a tuple with the LibraryFileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryFileLocation

`func (o *MonitoringSettings) SetLibraryFileLocation(v string)`

SetLibraryFileLocation sets LibraryFileLocation field to given value.

### HasLibraryFileLocation

`func (o *MonitoringSettings) HasLibraryFileLocation() bool`

HasLibraryFileLocation returns a boolean if a field has been set.

### GetMonitoringDataPath

`func (o *MonitoringSettings) GetMonitoringDataPath() string`

GetMonitoringDataPath returns the MonitoringDataPath field if non-nil, zero value otherwise.

### GetMonitoringDataPathOk

`func (o *MonitoringSettings) GetMonitoringDataPathOk() (*string, bool)`

GetMonitoringDataPathOk returns a tuple with the MonitoringDataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringDataPath

`func (o *MonitoringSettings) SetMonitoringDataPath(v string)`

SetMonitoringDataPath sets MonitoringDataPath field to given value.

### HasMonitoringDataPath

`func (o *MonitoringSettings) HasMonitoringDataPath() bool`

HasMonitoringDataPath returns a boolean if a field has been set.

### GetCustomConfigurationProperties

`func (o *MonitoringSettings) GetCustomConfigurationProperties() string`

GetCustomConfigurationProperties returns the CustomConfigurationProperties field if non-nil, zero value otherwise.

### GetCustomConfigurationPropertiesOk

`func (o *MonitoringSettings) GetCustomConfigurationPropertiesOk() (*string, bool)`

GetCustomConfigurationPropertiesOk returns a tuple with the CustomConfigurationProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfigurationProperties

`func (o *MonitoringSettings) SetCustomConfigurationProperties(v string)`

SetCustomConfigurationProperties sets CustomConfigurationProperties field to given value.


### GetServerRequestPathId

`func (o *MonitoringSettings) GetServerRequestPathId() string`

GetServerRequestPathId returns the ServerRequestPathId field if non-nil, zero value otherwise.

### GetServerRequestPathIdOk

`func (o *MonitoringSettings) GetServerRequestPathIdOk() (*string, bool)`

GetServerRequestPathIdOk returns a tuple with the ServerRequestPathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRequestPathId

`func (o *MonitoringSettings) SetServerRequestPathId(v string)`

SetServerRequestPathId sets ServerRequestPathId field to given value.


### GetSecureCookieAttribute

`func (o *MonitoringSettings) GetSecureCookieAttribute() bool`

GetSecureCookieAttribute returns the SecureCookieAttribute field if non-nil, zero value otherwise.

### GetSecureCookieAttributeOk

`func (o *MonitoringSettings) GetSecureCookieAttributeOk() (*bool, bool)`

GetSecureCookieAttributeOk returns a tuple with the SecureCookieAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureCookieAttribute

`func (o *MonitoringSettings) SetSecureCookieAttribute(v bool)`

SetSecureCookieAttribute sets SecureCookieAttribute field to given value.


### GetCookiePlacementDomain

`func (o *MonitoringSettings) GetCookiePlacementDomain() string`

GetCookiePlacementDomain returns the CookiePlacementDomain field if non-nil, zero value otherwise.

### GetCookiePlacementDomainOk

`func (o *MonitoringSettings) GetCookiePlacementDomainOk() (*string, bool)`

GetCookiePlacementDomainOk returns a tuple with the CookiePlacementDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookiePlacementDomain

`func (o *MonitoringSettings) SetCookiePlacementDomain(v string)`

SetCookiePlacementDomain sets CookiePlacementDomain field to given value.


### GetCacheControlHeaderOptimizations

`func (o *MonitoringSettings) GetCacheControlHeaderOptimizations() bool`

GetCacheControlHeaderOptimizations returns the CacheControlHeaderOptimizations field if non-nil, zero value otherwise.

### GetCacheControlHeaderOptimizationsOk

`func (o *MonitoringSettings) GetCacheControlHeaderOptimizationsOk() (*bool, bool)`

GetCacheControlHeaderOptimizationsOk returns a tuple with the CacheControlHeaderOptimizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControlHeaderOptimizations

`func (o *MonitoringSettings) SetCacheControlHeaderOptimizations(v bool)`

SetCacheControlHeaderOptimizations sets CacheControlHeaderOptimizations field to given value.


### GetAdvancedJavaScriptTagSettings

`func (o *MonitoringSettings) GetAdvancedJavaScriptTagSettings() AdvancedJavaScriptTagSettings`

GetAdvancedJavaScriptTagSettings returns the AdvancedJavaScriptTagSettings field if non-nil, zero value otherwise.

### GetAdvancedJavaScriptTagSettingsOk

`func (o *MonitoringSettings) GetAdvancedJavaScriptTagSettingsOk() (*AdvancedJavaScriptTagSettings, bool)`

GetAdvancedJavaScriptTagSettingsOk returns a tuple with the AdvancedJavaScriptTagSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedJavaScriptTagSettings

`func (o *MonitoringSettings) SetAdvancedJavaScriptTagSettings(v AdvancedJavaScriptTagSettings)`

SetAdvancedJavaScriptTagSettings sets AdvancedJavaScriptTagSettings field to given value.


### GetBrowserRestrictionSettings

`func (o *MonitoringSettings) GetBrowserRestrictionSettings() WebApplicationConfigBrowserRestrictionSettings`

GetBrowserRestrictionSettings returns the BrowserRestrictionSettings field if non-nil, zero value otherwise.

### GetBrowserRestrictionSettingsOk

`func (o *MonitoringSettings) GetBrowserRestrictionSettingsOk() (*WebApplicationConfigBrowserRestrictionSettings, bool)`

GetBrowserRestrictionSettingsOk returns a tuple with the BrowserRestrictionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserRestrictionSettings

`func (o *MonitoringSettings) SetBrowserRestrictionSettings(v WebApplicationConfigBrowserRestrictionSettings)`

SetBrowserRestrictionSettings sets BrowserRestrictionSettings field to given value.

### HasBrowserRestrictionSettings

`func (o *MonitoringSettings) HasBrowserRestrictionSettings() bool`

HasBrowserRestrictionSettings returns a boolean if a field has been set.

### GetIpAddressRestrictionSettings

`func (o *MonitoringSettings) GetIpAddressRestrictionSettings() WebApplicationConfigIpAddressRestrictionSettings`

GetIpAddressRestrictionSettings returns the IpAddressRestrictionSettings field if non-nil, zero value otherwise.

### GetIpAddressRestrictionSettingsOk

`func (o *MonitoringSettings) GetIpAddressRestrictionSettingsOk() (*WebApplicationConfigIpAddressRestrictionSettings, bool)`

GetIpAddressRestrictionSettingsOk returns a tuple with the IpAddressRestrictionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressRestrictionSettings

`func (o *MonitoringSettings) SetIpAddressRestrictionSettings(v WebApplicationConfigIpAddressRestrictionSettings)`

SetIpAddressRestrictionSettings sets IpAddressRestrictionSettings field to given value.

### HasIpAddressRestrictionSettings

`func (o *MonitoringSettings) HasIpAddressRestrictionSettings() bool`

HasIpAddressRestrictionSettings returns a boolean if a field has been set.

### GetJavaScriptInjectionRules

`func (o *MonitoringSettings) GetJavaScriptInjectionRules() []JavaScriptInjectionRules`

GetJavaScriptInjectionRules returns the JavaScriptInjectionRules field if non-nil, zero value otherwise.

### GetJavaScriptInjectionRulesOk

`func (o *MonitoringSettings) GetJavaScriptInjectionRulesOk() (*[]JavaScriptInjectionRules, bool)`

GetJavaScriptInjectionRulesOk returns a tuple with the JavaScriptInjectionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaScriptInjectionRules

`func (o *MonitoringSettings) SetJavaScriptInjectionRules(v []JavaScriptInjectionRules)`

SetJavaScriptInjectionRules sets JavaScriptInjectionRules field to given value.

### HasJavaScriptInjectionRules

`func (o *MonitoringSettings) HasJavaScriptInjectionRules() bool`

HasJavaScriptInjectionRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


