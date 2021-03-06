/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// MonitoringSettings Real user monitoring settings.
type MonitoringSettings struct {
	// `fetch()` request capture enabled/disabled.
	FetchRequests bool `json:"fetchRequests"`
	// `XmlHttpRequest` support enabled/disabled.
	XmlHttpRequest bool `json:"xmlHttpRequest"`
	JavaScriptFrameworkSupport JavaScriptFrameworkSupport `json:"javaScriptFrameworkSupport"`
	ContentCapture ContentCapture `json:"contentCapture"`
	// You can exclude some actions from becoming XHR actions.  Put a regular expression, matching all the required URLs, here.  If noting specified the feature is disabled.
	ExcludeXhrRegex string `json:"excludeXhrRegex"`
	// To enable RUM for XHR calls to AWS Lambda, define a regular expression matching these calls, Dynatrace can then automatically add a custom header (x-dtc) to each such request to the respective endpoints in AWS.  Important: These endpoints must accept the x-dtc header, or the requests will fail. 
	CorrelationHeaderInclusionRegex *string `json:"correlationHeaderInclusionRegex,omitempty"`
	// JavaScript injection mode.
	InjectionMode string `json:"injectionMode"`
	// Add the cross origin = anonymous attribute to capture JavaScript error messages and W3C resource timings.
	AddCrossOriginAnonymousAttribute *bool `json:"addCrossOriginAnonymousAttribute,omitempty"`
	// Time duration for the cache settings.
	ScriptTagCacheDurationInHours *int32 `json:"scriptTagCacheDurationInHours,omitempty"`
	// The location of your application’s custom JavaScript library file.  If nothing specified the root directory of your web server is used.  Only supported by auto-injected applications.
	LibraryFileLocation *string `json:"libraryFileLocation,omitempty"`
	// The location to send monitoring data from the JavaScript tag.  Specify either a relative or an absolute URL. If you use an absolute URL, data will be sent using CORS.  Only supported by auto-injected applications.
	MonitoringDataPath *string `json:"monitoringDataPath,omitempty"`
	// Additional JavaScript tag properties that are specific to your application. To do this, type key=value pairs separated using a (|) symbol.
	CustomConfigurationProperties string `json:"customConfigurationProperties"`
	// Path to identify the server’s request ID.
	ServerRequestPathId string `json:"serverRequestPathId"`
	// Secure attribute usage for Dynatrace cookies enabled/disabled.
	SecureCookieAttribute bool `json:"secureCookieAttribute"`
	// Domain for cookie placement.
	CookiePlacementDomain string `json:"cookiePlacementDomain"`
	// Optimize the value of cache control headers for use with Dynatrace real user monitoring enabled/disabled.
	CacheControlHeaderOptimizations bool `json:"cacheControlHeaderOptimizations"`
	AdvancedJavaScriptTagSettings AdvancedJavaScriptTagSettings `json:"advancedJavaScriptTagSettings"`
	BrowserRestrictionSettings *WebApplicationConfigBrowserRestrictionSettings `json:"browserRestrictionSettings,omitempty"`
	IpAddressRestrictionSettings *WebApplicationConfigIpAddressRestrictionSettings `json:"ipAddressRestrictionSettings,omitempty"`
	// Java script injection rules.
	JavaScriptInjectionRules *[]JavaScriptInjectionRules `json:"javaScriptInjectionRules,omitempty"`
}

// NewMonitoringSettings instantiates a new MonitoringSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringSettings(fetchRequests bool, xmlHttpRequest bool, javaScriptFrameworkSupport JavaScriptFrameworkSupport, contentCapture ContentCapture, excludeXhrRegex string, injectionMode string, customConfigurationProperties string, serverRequestPathId string, secureCookieAttribute bool, cookiePlacementDomain string, cacheControlHeaderOptimizations bool, advancedJavaScriptTagSettings AdvancedJavaScriptTagSettings, ) *MonitoringSettings {
	this := MonitoringSettings{}
	this.FetchRequests = fetchRequests
	this.XmlHttpRequest = xmlHttpRequest
	this.JavaScriptFrameworkSupport = javaScriptFrameworkSupport
	this.ContentCapture = contentCapture
	this.ExcludeXhrRegex = excludeXhrRegex
	this.InjectionMode = injectionMode
	this.CustomConfigurationProperties = customConfigurationProperties
	this.ServerRequestPathId = serverRequestPathId
	this.SecureCookieAttribute = secureCookieAttribute
	this.CookiePlacementDomain = cookiePlacementDomain
	this.CacheControlHeaderOptimizations = cacheControlHeaderOptimizations
	this.AdvancedJavaScriptTagSettings = advancedJavaScriptTagSettings
	return &this
}

// NewMonitoringSettingsWithDefaults instantiates a new MonitoringSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringSettingsWithDefaults() *MonitoringSettings {
	this := MonitoringSettings{}
	return &this
}

// GetFetchRequests returns the FetchRequests field value
func (o *MonitoringSettings) GetFetchRequests() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.FetchRequests
}

// GetFetchRequestsOk returns a tuple with the FetchRequests field value
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetFetchRequestsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FetchRequests, true
}

// SetFetchRequests sets field value
func (o *MonitoringSettings) SetFetchRequests(v bool) {
	o.FetchRequests = v
}

// GetXmlHttpRequest returns the XmlHttpRequest field value
func (o *MonitoringSettings) GetXmlHttpRequest() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.XmlHttpRequest
}

// GetXmlHttpRequestOk returns a tuple with the XmlHttpRequest field value
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetXmlHttpRequestOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.XmlHttpRequest, true
}

// SetXmlHttpRequest sets field value
func (o *MonitoringSettings) SetXmlHttpRequest(v bool) {
	o.XmlHttpRequest = v
}

// GetJavaScriptFrameworkSupport returns the JavaScriptFrameworkSupport field value
func (o *MonitoringSettings) GetJavaScriptFrameworkSupport() JavaScriptFrameworkSupport {
	if o == nil  {
		var ret JavaScriptFrameworkSupport
		return ret
	}

	return o.JavaScriptFrameworkSupport
}

// GetJavaScriptFrameworkSupportOk returns a tuple with the JavaScriptFrameworkSupport field value
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetJavaScriptFrameworkSupportOk() (*JavaScriptFrameworkSupport, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.JavaScriptFrameworkSupport, true
}

// SetJavaScriptFrameworkSupport sets field value
func (o *MonitoringSettings) SetJavaScriptFrameworkSupport(v JavaScriptFrameworkSupport) {
	o.JavaScriptFrameworkSupport = v
}

// GetContentCapture returns the ContentCapture field value
func (o *MonitoringSettings) GetContentCapture() ContentCapture {
	if o == nil  {
		var ret ContentCapture
		return ret
	}

	return o.ContentCapture
}

// GetContentCaptureOk returns a tuple with the ContentCapture field value
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetContentCaptureOk() (*ContentCapture, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContentCapture, true
}

// SetContentCapture sets field value
func (o *MonitoringSettings) SetContentCapture(v ContentCapture) {
	o.ContentCapture = v
}

// GetExcludeXhrRegex returns the ExcludeXhrRegex field value
func (o *MonitoringSettings) GetExcludeXhrRegex() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ExcludeXhrRegex
}

// GetExcludeXhrRegexOk returns a tuple with the ExcludeXhrRegex field value
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetExcludeXhrRegexOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExcludeXhrRegex, true
}

// SetExcludeXhrRegex sets field value
func (o *MonitoringSettings) SetExcludeXhrRegex(v string) {
	o.ExcludeXhrRegex = v
}

// GetCorrelationHeaderInclusionRegex returns the CorrelationHeaderInclusionRegex field value if set, zero value otherwise.
func (o *MonitoringSettings) GetCorrelationHeaderInclusionRegex() string {
	if o == nil || o.CorrelationHeaderInclusionRegex == nil {
		var ret string
		return ret
	}
	return *o.CorrelationHeaderInclusionRegex
}

// GetCorrelationHeaderInclusionRegexOk returns a tuple with the CorrelationHeaderInclusionRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetCorrelationHeaderInclusionRegexOk() (*string, bool) {
	if o == nil || o.CorrelationHeaderInclusionRegex == nil {
		return nil, false
	}
	return o.CorrelationHeaderInclusionRegex, true
}

// HasCorrelationHeaderInclusionRegex returns a boolean if a field has been set.
func (o *MonitoringSettings) HasCorrelationHeaderInclusionRegex() bool {
	if o != nil && o.CorrelationHeaderInclusionRegex != nil {
		return true
	}

	return false
}

// SetCorrelationHeaderInclusionRegex gets a reference to the given string and assigns it to the CorrelationHeaderInclusionRegex field.
func (o *MonitoringSettings) SetCorrelationHeaderInclusionRegex(v string) {
	o.CorrelationHeaderInclusionRegex = &v
}

// GetInjectionMode returns the InjectionMode field value
func (o *MonitoringSettings) GetInjectionMode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.InjectionMode
}

// GetInjectionModeOk returns a tuple with the InjectionMode field value
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetInjectionModeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InjectionMode, true
}

// SetInjectionMode sets field value
func (o *MonitoringSettings) SetInjectionMode(v string) {
	o.InjectionMode = v
}

// GetAddCrossOriginAnonymousAttribute returns the AddCrossOriginAnonymousAttribute field value if set, zero value otherwise.
func (o *MonitoringSettings) GetAddCrossOriginAnonymousAttribute() bool {
	if o == nil || o.AddCrossOriginAnonymousAttribute == nil {
		var ret bool
		return ret
	}
	return *o.AddCrossOriginAnonymousAttribute
}

// GetAddCrossOriginAnonymousAttributeOk returns a tuple with the AddCrossOriginAnonymousAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetAddCrossOriginAnonymousAttributeOk() (*bool, bool) {
	if o == nil || o.AddCrossOriginAnonymousAttribute == nil {
		return nil, false
	}
	return o.AddCrossOriginAnonymousAttribute, true
}

// HasAddCrossOriginAnonymousAttribute returns a boolean if a field has been set.
func (o *MonitoringSettings) HasAddCrossOriginAnonymousAttribute() bool {
	if o != nil && o.AddCrossOriginAnonymousAttribute != nil {
		return true
	}

	return false
}

// SetAddCrossOriginAnonymousAttribute gets a reference to the given bool and assigns it to the AddCrossOriginAnonymousAttribute field.
func (o *MonitoringSettings) SetAddCrossOriginAnonymousAttribute(v bool) {
	o.AddCrossOriginAnonymousAttribute = &v
}

// GetScriptTagCacheDurationInHours returns the ScriptTagCacheDurationInHours field value if set, zero value otherwise.
func (o *MonitoringSettings) GetScriptTagCacheDurationInHours() int32 {
	if o == nil || o.ScriptTagCacheDurationInHours == nil {
		var ret int32
		return ret
	}
	return *o.ScriptTagCacheDurationInHours
}

// GetScriptTagCacheDurationInHoursOk returns a tuple with the ScriptTagCacheDurationInHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetScriptTagCacheDurationInHoursOk() (*int32, bool) {
	if o == nil || o.ScriptTagCacheDurationInHours == nil {
		return nil, false
	}
	return o.ScriptTagCacheDurationInHours, true
}

// HasScriptTagCacheDurationInHours returns a boolean if a field has been set.
func (o *MonitoringSettings) HasScriptTagCacheDurationInHours() bool {
	if o != nil && o.ScriptTagCacheDurationInHours != nil {
		return true
	}

	return false
}

// SetScriptTagCacheDurationInHours gets a reference to the given int32 and assigns it to the ScriptTagCacheDurationInHours field.
func (o *MonitoringSettings) SetScriptTagCacheDurationInHours(v int32) {
	o.ScriptTagCacheDurationInHours = &v
}

// GetLibraryFileLocation returns the LibraryFileLocation field value if set, zero value otherwise.
func (o *MonitoringSettings) GetLibraryFileLocation() string {
	if o == nil || o.LibraryFileLocation == nil {
		var ret string
		return ret
	}
	return *o.LibraryFileLocation
}

// GetLibraryFileLocationOk returns a tuple with the LibraryFileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetLibraryFileLocationOk() (*string, bool) {
	if o == nil || o.LibraryFileLocation == nil {
		return nil, false
	}
	return o.LibraryFileLocation, true
}

// HasLibraryFileLocation returns a boolean if a field has been set.
func (o *MonitoringSettings) HasLibraryFileLocation() bool {
	if o != nil && o.LibraryFileLocation != nil {
		return true
	}

	return false
}

// SetLibraryFileLocation gets a reference to the given string and assigns it to the LibraryFileLocation field.
func (o *MonitoringSettings) SetLibraryFileLocation(v string) {
	o.LibraryFileLocation = &v
}

// GetMonitoringDataPath returns the MonitoringDataPath field value if set, zero value otherwise.
func (o *MonitoringSettings) GetMonitoringDataPath() string {
	if o == nil || o.MonitoringDataPath == nil {
		var ret string
		return ret
	}
	return *o.MonitoringDataPath
}

// GetMonitoringDataPathOk returns a tuple with the MonitoringDataPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetMonitoringDataPathOk() (*string, bool) {
	if o == nil || o.MonitoringDataPath == nil {
		return nil, false
	}
	return o.MonitoringDataPath, true
}

// HasMonitoringDataPath returns a boolean if a field has been set.
func (o *MonitoringSettings) HasMonitoringDataPath() bool {
	if o != nil && o.MonitoringDataPath != nil {
		return true
	}

	return false
}

// SetMonitoringDataPath gets a reference to the given string and assigns it to the MonitoringDataPath field.
func (o *MonitoringSettings) SetMonitoringDataPath(v string) {
	o.MonitoringDataPath = &v
}

// GetCustomConfigurationProperties returns the CustomConfigurationProperties field value
func (o *MonitoringSettings) GetCustomConfigurationProperties() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CustomConfigurationProperties
}

// GetCustomConfigurationPropertiesOk returns a tuple with the CustomConfigurationProperties field value
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetCustomConfigurationPropertiesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomConfigurationProperties, true
}

// SetCustomConfigurationProperties sets field value
func (o *MonitoringSettings) SetCustomConfigurationProperties(v string) {
	o.CustomConfigurationProperties = v
}

// GetServerRequestPathId returns the ServerRequestPathId field value
func (o *MonitoringSettings) GetServerRequestPathId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ServerRequestPathId
}

// GetServerRequestPathIdOk returns a tuple with the ServerRequestPathId field value
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetServerRequestPathIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServerRequestPathId, true
}

// SetServerRequestPathId sets field value
func (o *MonitoringSettings) SetServerRequestPathId(v string) {
	o.ServerRequestPathId = v
}

// GetSecureCookieAttribute returns the SecureCookieAttribute field value
func (o *MonitoringSettings) GetSecureCookieAttribute() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.SecureCookieAttribute
}

// GetSecureCookieAttributeOk returns a tuple with the SecureCookieAttribute field value
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetSecureCookieAttributeOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecureCookieAttribute, true
}

// SetSecureCookieAttribute sets field value
func (o *MonitoringSettings) SetSecureCookieAttribute(v bool) {
	o.SecureCookieAttribute = v
}

// GetCookiePlacementDomain returns the CookiePlacementDomain field value
func (o *MonitoringSettings) GetCookiePlacementDomain() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CookiePlacementDomain
}

// GetCookiePlacementDomainOk returns a tuple with the CookiePlacementDomain field value
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetCookiePlacementDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CookiePlacementDomain, true
}

// SetCookiePlacementDomain sets field value
func (o *MonitoringSettings) SetCookiePlacementDomain(v string) {
	o.CookiePlacementDomain = v
}

// GetCacheControlHeaderOptimizations returns the CacheControlHeaderOptimizations field value
func (o *MonitoringSettings) GetCacheControlHeaderOptimizations() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.CacheControlHeaderOptimizations
}

// GetCacheControlHeaderOptimizationsOk returns a tuple with the CacheControlHeaderOptimizations field value
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetCacheControlHeaderOptimizationsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CacheControlHeaderOptimizations, true
}

// SetCacheControlHeaderOptimizations sets field value
func (o *MonitoringSettings) SetCacheControlHeaderOptimizations(v bool) {
	o.CacheControlHeaderOptimizations = v
}

// GetAdvancedJavaScriptTagSettings returns the AdvancedJavaScriptTagSettings field value
func (o *MonitoringSettings) GetAdvancedJavaScriptTagSettings() AdvancedJavaScriptTagSettings {
	if o == nil  {
		var ret AdvancedJavaScriptTagSettings
		return ret
	}

	return o.AdvancedJavaScriptTagSettings
}

// GetAdvancedJavaScriptTagSettingsOk returns a tuple with the AdvancedJavaScriptTagSettings field value
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetAdvancedJavaScriptTagSettingsOk() (*AdvancedJavaScriptTagSettings, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AdvancedJavaScriptTagSettings, true
}

// SetAdvancedJavaScriptTagSettings sets field value
func (o *MonitoringSettings) SetAdvancedJavaScriptTagSettings(v AdvancedJavaScriptTagSettings) {
	o.AdvancedJavaScriptTagSettings = v
}

// GetBrowserRestrictionSettings returns the BrowserRestrictionSettings field value if set, zero value otherwise.
func (o *MonitoringSettings) GetBrowserRestrictionSettings() WebApplicationConfigBrowserRestrictionSettings {
	if o == nil || o.BrowserRestrictionSettings == nil {
		var ret WebApplicationConfigBrowserRestrictionSettings
		return ret
	}
	return *o.BrowserRestrictionSettings
}

// GetBrowserRestrictionSettingsOk returns a tuple with the BrowserRestrictionSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetBrowserRestrictionSettingsOk() (*WebApplicationConfigBrowserRestrictionSettings, bool) {
	if o == nil || o.BrowserRestrictionSettings == nil {
		return nil, false
	}
	return o.BrowserRestrictionSettings, true
}

// HasBrowserRestrictionSettings returns a boolean if a field has been set.
func (o *MonitoringSettings) HasBrowserRestrictionSettings() bool {
	if o != nil && o.BrowserRestrictionSettings != nil {
		return true
	}

	return false
}

// SetBrowserRestrictionSettings gets a reference to the given WebApplicationConfigBrowserRestrictionSettings and assigns it to the BrowserRestrictionSettings field.
func (o *MonitoringSettings) SetBrowserRestrictionSettings(v WebApplicationConfigBrowserRestrictionSettings) {
	o.BrowserRestrictionSettings = &v
}

// GetIpAddressRestrictionSettings returns the IpAddressRestrictionSettings field value if set, zero value otherwise.
func (o *MonitoringSettings) GetIpAddressRestrictionSettings() WebApplicationConfigIpAddressRestrictionSettings {
	if o == nil || o.IpAddressRestrictionSettings == nil {
		var ret WebApplicationConfigIpAddressRestrictionSettings
		return ret
	}
	return *o.IpAddressRestrictionSettings
}

// GetIpAddressRestrictionSettingsOk returns a tuple with the IpAddressRestrictionSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetIpAddressRestrictionSettingsOk() (*WebApplicationConfigIpAddressRestrictionSettings, bool) {
	if o == nil || o.IpAddressRestrictionSettings == nil {
		return nil, false
	}
	return o.IpAddressRestrictionSettings, true
}

// HasIpAddressRestrictionSettings returns a boolean if a field has been set.
func (o *MonitoringSettings) HasIpAddressRestrictionSettings() bool {
	if o != nil && o.IpAddressRestrictionSettings != nil {
		return true
	}

	return false
}

// SetIpAddressRestrictionSettings gets a reference to the given WebApplicationConfigIpAddressRestrictionSettings and assigns it to the IpAddressRestrictionSettings field.
func (o *MonitoringSettings) SetIpAddressRestrictionSettings(v WebApplicationConfigIpAddressRestrictionSettings) {
	o.IpAddressRestrictionSettings = &v
}

// GetJavaScriptInjectionRules returns the JavaScriptInjectionRules field value if set, zero value otherwise.
func (o *MonitoringSettings) GetJavaScriptInjectionRules() []JavaScriptInjectionRules {
	if o == nil || o.JavaScriptInjectionRules == nil {
		var ret []JavaScriptInjectionRules
		return ret
	}
	return *o.JavaScriptInjectionRules
}

// GetJavaScriptInjectionRulesOk returns a tuple with the JavaScriptInjectionRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetJavaScriptInjectionRulesOk() (*[]JavaScriptInjectionRules, bool) {
	if o == nil || o.JavaScriptInjectionRules == nil {
		return nil, false
	}
	return o.JavaScriptInjectionRules, true
}

// HasJavaScriptInjectionRules returns a boolean if a field has been set.
func (o *MonitoringSettings) HasJavaScriptInjectionRules() bool {
	if o != nil && o.JavaScriptInjectionRules != nil {
		return true
	}

	return false
}

// SetJavaScriptInjectionRules gets a reference to the given []JavaScriptInjectionRules and assigns it to the JavaScriptInjectionRules field.
func (o *MonitoringSettings) SetJavaScriptInjectionRules(v []JavaScriptInjectionRules) {
	o.JavaScriptInjectionRules = &v
}

func (o MonitoringSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fetchRequests"] = o.FetchRequests
	}
	if true {
		toSerialize["xmlHttpRequest"] = o.XmlHttpRequest
	}
	if true {
		toSerialize["javaScriptFrameworkSupport"] = o.JavaScriptFrameworkSupport
	}
	if true {
		toSerialize["contentCapture"] = o.ContentCapture
	}
	if true {
		toSerialize["excludeXhrRegex"] = o.ExcludeXhrRegex
	}
	if o.CorrelationHeaderInclusionRegex != nil {
		toSerialize["correlationHeaderInclusionRegex"] = o.CorrelationHeaderInclusionRegex
	}
	if true {
		toSerialize["injectionMode"] = o.InjectionMode
	}
	if o.AddCrossOriginAnonymousAttribute != nil {
		toSerialize["addCrossOriginAnonymousAttribute"] = o.AddCrossOriginAnonymousAttribute
	}
	if o.ScriptTagCacheDurationInHours != nil {
		toSerialize["scriptTagCacheDurationInHours"] = o.ScriptTagCacheDurationInHours
	}
	if o.LibraryFileLocation != nil {
		toSerialize["libraryFileLocation"] = o.LibraryFileLocation
	}
	if o.MonitoringDataPath != nil {
		toSerialize["monitoringDataPath"] = o.MonitoringDataPath
	}
	if true {
		toSerialize["customConfigurationProperties"] = o.CustomConfigurationProperties
	}
	if true {
		toSerialize["serverRequestPathId"] = o.ServerRequestPathId
	}
	if true {
		toSerialize["secureCookieAttribute"] = o.SecureCookieAttribute
	}
	if true {
		toSerialize["cookiePlacementDomain"] = o.CookiePlacementDomain
	}
	if true {
		toSerialize["cacheControlHeaderOptimizations"] = o.CacheControlHeaderOptimizations
	}
	if true {
		toSerialize["advancedJavaScriptTagSettings"] = o.AdvancedJavaScriptTagSettings
	}
	if o.BrowserRestrictionSettings != nil {
		toSerialize["browserRestrictionSettings"] = o.BrowserRestrictionSettings
	}
	if o.IpAddressRestrictionSettings != nil {
		toSerialize["ipAddressRestrictionSettings"] = o.IpAddressRestrictionSettings
	}
	if o.JavaScriptInjectionRules != nil {
		toSerialize["javaScriptInjectionRules"] = o.JavaScriptInjectionRules
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringSettings struct {
	value *MonitoringSettings
	isSet bool
}

func (v NullableMonitoringSettings) Get() *MonitoringSettings {
	return v.value
}

func (v *NullableMonitoringSettings) Set(val *MonitoringSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringSettings(val *MonitoringSettings) *NullableMonitoringSettings {
	return &NullableMonitoringSettings{value: val, isSet: true}
}

func (v NullableMonitoringSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


