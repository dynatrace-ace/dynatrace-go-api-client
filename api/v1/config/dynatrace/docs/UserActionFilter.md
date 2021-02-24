# UserActionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionDurationFromMilliseconds** | Pointer to **int32** | Only actions with a duration more than or equal to this value (in milliseconds) are included in the metric calculation. | [optional] 
**ActionDurationToMilliseconds** | Pointer to **int32** | Only actions with a duration less than or equal to this value (in milliseconds) are included in the metric calculation. | [optional] 
**LoadAction** | Pointer to **bool** | The status of load actions in the metric calculation:   * &#x60;true&#x60;: Only load actions are included.  * &#x60;false&#x60;: All actions are included. | [optional] 
**XhrAction** | Pointer to **bool** | The status of XHR actions in the metric calculation:   * &#x60;true&#x60;: Only XHR actions are included.  * &#x60;false&#x60;: All actions are included. | [optional] 
**CustomAction** | Pointer to **bool** | The status of custom actions in the metric calculation:   * &#x60;true&#x60;: Only custom actions are included.  * &#x60;false&#x60;: All actions are included. | [optional] 
**Apdex** | Pointer to **string** | Only actions with the specified Apdex score are included in the metric calculation. | [optional] 
**Domain** | Pointer to **string** | Only user actions coming from the specified domain are included in the metric calculation.    The EQUALS operator applies. | [optional] 
**UserActionName** | Pointer to **string** | Only actions with this name are included in the metric calculation.    The EQUALS operator applies. | [optional] 
**RealUser** | Pointer to **bool** | The status of actions coming from real users in the metric calculation:   * &#x60;true&#x60;: Only actions from real users are included.  * &#x60;false&#x60;: All actions are included. | [optional] 
**Robot** | Pointer to **bool** | The status of actions coming from robots in the metric calculation:   * &#x60;true&#x60;: Only actions from robots are included.  * &#x60;false&#x60;: All actions are included. | [optional] 
**Synthetic** | Pointer to **bool** | The status of actions coming from synthetic monitors in the metric calculation:   * &#x60;true&#x60;: Only actions from synthetic monitors are included.  * &#x60;false&#x60;: All actions are included. | [optional] 
**BrowserFamily** | Pointer to **string** | Only user actions coming from the specified browser family are included in the metric calculation.     The EQUALS operator applies. | [optional] 
**BrowserType** | Pointer to **string** | Only user actions coming from the specified browser type are included in the metric calculation.     The EQUALS operator applies. | [optional] 
**BrowserVersion** | Pointer to **string** | Only user actions coming from the specified browser version are included in the metric calculation.     The EQUALS operator applies. | [optional] 
**HasCustomErrors** | Pointer to **bool** | The custom error status of the actions to be included in the metric calculation:   * &#x60;true&#x60;: Only actions with custom errors are included.  * &#x60;false&#x60;: All actions are included. | [optional] 
**HasAnyError** | Pointer to **bool** | The error status of the actions to be included in the metric calculation:   * &#x60;true&#x60;: Only actions that have any errors are included.  * &#x60;false&#x60;: All actions are included. | [optional] 
**HasHttpErrors** | Pointer to **bool** | The HTTP error status of the actions to be included in the metric calculation:   * &#x60;true&#x60;: Only actions with HTTP errors are included.  * &#x60;false&#x60;: All actions are included. | [optional] 
**HasJavascriptErrors** | Pointer to **bool** | The JavaScript error status of the actions to be included in the metric calculation:   * &#x60;true&#x60;: Only actions with JavaScript errors are included.  * &#x60;false&#x60;: All actions are included. | [optional] 
**City** | Pointer to **string** | Only actions of users from this city are included in the metric calculation.    Specify geolocation ID here. | [optional] 
**Continent** | Pointer to **string** | Only actions of users from this continent are included in the metric calculation.    Specify geolocation ID here. | [optional] 
**Country** | Pointer to **string** | Only actions of users from this country are included in the metric calculation.    Specify geolocation ID here. | [optional] 
**Region** | Pointer to **string** | Only actions of users from this region are included in the metric calculation.    Specify geolocation ID here. | [optional] 
**Ip** | Pointer to **string** | Only actions coming from this IP address are included in the metric calculation.     The EQUALS operator applies. | [optional] 
**IpV6Traffic** | Pointer to **bool** | The IPv6 status of the actions to be included in the metric calculation:   * &#x60;true&#x60;: Only actions coming from IPv6 are included.  * &#x60;false&#x60;: All actions are included. | [optional] 
**OsFamily** | Pointer to **string** | Only actions coming from this OS family are included in the metric calculation.    Specify the OS ID here. | [optional] 
**OsVersion** | Pointer to **string** | Only actions coming from this OS version are included in the metric calculation.    Specify the OS ID here. | [optional] 
**HttpErrorCode** | Pointer to **int32** | The HTTP error status code of the actions to be included in the metric calculation. | [optional] 
**HttpPath** | Pointer to **string** | The request path that has been determined to be the origin of an HTTP error of the actions to be included in the metric calculation. | [optional] 
**CustomErrorType** | Pointer to **string** | The custom error type of the actions to be included in the metric calculation. | [optional] 
**CustomErrorName** | Pointer to **string** | The custom error name of the actions to be included in the metric calculation. | [optional] 
**UserActionProperties** | Pointer to [**[]UserActionPropertyFilter**](UserActionPropertyFilter.md) | Only actions with the specified properties are included in the metric calculation. | [optional] 
**TargetViewName** | Pointer to **string** | Only actions on the specified view are included in the metric calculation.    The EQUALS operator applies. | [optional] 

## Methods

### NewUserActionFilter

`func NewUserActionFilter() *UserActionFilter`

NewUserActionFilter instantiates a new UserActionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActionFilterWithDefaults

`func NewUserActionFilterWithDefaults() *UserActionFilter`

NewUserActionFilterWithDefaults instantiates a new UserActionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionDurationFromMilliseconds

`func (o *UserActionFilter) GetActionDurationFromMilliseconds() int32`

GetActionDurationFromMilliseconds returns the ActionDurationFromMilliseconds field if non-nil, zero value otherwise.

### GetActionDurationFromMillisecondsOk

`func (o *UserActionFilter) GetActionDurationFromMillisecondsOk() (*int32, bool)`

GetActionDurationFromMillisecondsOk returns a tuple with the ActionDurationFromMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDurationFromMilliseconds

`func (o *UserActionFilter) SetActionDurationFromMilliseconds(v int32)`

SetActionDurationFromMilliseconds sets ActionDurationFromMilliseconds field to given value.

### HasActionDurationFromMilliseconds

`func (o *UserActionFilter) HasActionDurationFromMilliseconds() bool`

HasActionDurationFromMilliseconds returns a boolean if a field has been set.

### GetActionDurationToMilliseconds

`func (o *UserActionFilter) GetActionDurationToMilliseconds() int32`

GetActionDurationToMilliseconds returns the ActionDurationToMilliseconds field if non-nil, zero value otherwise.

### GetActionDurationToMillisecondsOk

`func (o *UserActionFilter) GetActionDurationToMillisecondsOk() (*int32, bool)`

GetActionDurationToMillisecondsOk returns a tuple with the ActionDurationToMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDurationToMilliseconds

`func (o *UserActionFilter) SetActionDurationToMilliseconds(v int32)`

SetActionDurationToMilliseconds sets ActionDurationToMilliseconds field to given value.

### HasActionDurationToMilliseconds

`func (o *UserActionFilter) HasActionDurationToMilliseconds() bool`

HasActionDurationToMilliseconds returns a boolean if a field has been set.

### GetLoadAction

`func (o *UserActionFilter) GetLoadAction() bool`

GetLoadAction returns the LoadAction field if non-nil, zero value otherwise.

### GetLoadActionOk

`func (o *UserActionFilter) GetLoadActionOk() (*bool, bool)`

GetLoadActionOk returns a tuple with the LoadAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadAction

`func (o *UserActionFilter) SetLoadAction(v bool)`

SetLoadAction sets LoadAction field to given value.

### HasLoadAction

`func (o *UserActionFilter) HasLoadAction() bool`

HasLoadAction returns a boolean if a field has been set.

### GetXhrAction

`func (o *UserActionFilter) GetXhrAction() bool`

GetXhrAction returns the XhrAction field if non-nil, zero value otherwise.

### GetXhrActionOk

`func (o *UserActionFilter) GetXhrActionOk() (*bool, bool)`

GetXhrActionOk returns a tuple with the XhrAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXhrAction

`func (o *UserActionFilter) SetXhrAction(v bool)`

SetXhrAction sets XhrAction field to given value.

### HasXhrAction

`func (o *UserActionFilter) HasXhrAction() bool`

HasXhrAction returns a boolean if a field has been set.

### GetCustomAction

`func (o *UserActionFilter) GetCustomAction() bool`

GetCustomAction returns the CustomAction field if non-nil, zero value otherwise.

### GetCustomActionOk

`func (o *UserActionFilter) GetCustomActionOk() (*bool, bool)`

GetCustomActionOk returns a tuple with the CustomAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAction

`func (o *UserActionFilter) SetCustomAction(v bool)`

SetCustomAction sets CustomAction field to given value.

### HasCustomAction

`func (o *UserActionFilter) HasCustomAction() bool`

HasCustomAction returns a boolean if a field has been set.

### GetApdex

`func (o *UserActionFilter) GetApdex() string`

GetApdex returns the Apdex field if non-nil, zero value otherwise.

### GetApdexOk

`func (o *UserActionFilter) GetApdexOk() (*string, bool)`

GetApdexOk returns a tuple with the Apdex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApdex

`func (o *UserActionFilter) SetApdex(v string)`

SetApdex sets Apdex field to given value.

### HasApdex

`func (o *UserActionFilter) HasApdex() bool`

HasApdex returns a boolean if a field has been set.

### GetDomain

`func (o *UserActionFilter) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UserActionFilter) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UserActionFilter) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UserActionFilter) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetUserActionName

`func (o *UserActionFilter) GetUserActionName() string`

GetUserActionName returns the UserActionName field if non-nil, zero value otherwise.

### GetUserActionNameOk

`func (o *UserActionFilter) GetUserActionNameOk() (*string, bool)`

GetUserActionNameOk returns a tuple with the UserActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionName

`func (o *UserActionFilter) SetUserActionName(v string)`

SetUserActionName sets UserActionName field to given value.

### HasUserActionName

`func (o *UserActionFilter) HasUserActionName() bool`

HasUserActionName returns a boolean if a field has been set.

### GetRealUser

`func (o *UserActionFilter) GetRealUser() bool`

GetRealUser returns the RealUser field if non-nil, zero value otherwise.

### GetRealUserOk

`func (o *UserActionFilter) GetRealUserOk() (*bool, bool)`

GetRealUserOk returns a tuple with the RealUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealUser

`func (o *UserActionFilter) SetRealUser(v bool)`

SetRealUser sets RealUser field to given value.

### HasRealUser

`func (o *UserActionFilter) HasRealUser() bool`

HasRealUser returns a boolean if a field has been set.

### GetRobot

`func (o *UserActionFilter) GetRobot() bool`

GetRobot returns the Robot field if non-nil, zero value otherwise.

### GetRobotOk

`func (o *UserActionFilter) GetRobotOk() (*bool, bool)`

GetRobotOk returns a tuple with the Robot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRobot

`func (o *UserActionFilter) SetRobot(v bool)`

SetRobot sets Robot field to given value.

### HasRobot

`func (o *UserActionFilter) HasRobot() bool`

HasRobot returns a boolean if a field has been set.

### GetSynthetic

`func (o *UserActionFilter) GetSynthetic() bool`

GetSynthetic returns the Synthetic field if non-nil, zero value otherwise.

### GetSyntheticOk

`func (o *UserActionFilter) GetSyntheticOk() (*bool, bool)`

GetSyntheticOk returns a tuple with the Synthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetic

`func (o *UserActionFilter) SetSynthetic(v bool)`

SetSynthetic sets Synthetic field to given value.

### HasSynthetic

`func (o *UserActionFilter) HasSynthetic() bool`

HasSynthetic returns a boolean if a field has been set.

### GetBrowserFamily

`func (o *UserActionFilter) GetBrowserFamily() string`

GetBrowserFamily returns the BrowserFamily field if non-nil, zero value otherwise.

### GetBrowserFamilyOk

`func (o *UserActionFilter) GetBrowserFamilyOk() (*string, bool)`

GetBrowserFamilyOk returns a tuple with the BrowserFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserFamily

`func (o *UserActionFilter) SetBrowserFamily(v string)`

SetBrowserFamily sets BrowserFamily field to given value.

### HasBrowserFamily

`func (o *UserActionFilter) HasBrowserFamily() bool`

HasBrowserFamily returns a boolean if a field has been set.

### GetBrowserType

`func (o *UserActionFilter) GetBrowserType() string`

GetBrowserType returns the BrowserType field if non-nil, zero value otherwise.

### GetBrowserTypeOk

`func (o *UserActionFilter) GetBrowserTypeOk() (*string, bool)`

GetBrowserTypeOk returns a tuple with the BrowserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserType

`func (o *UserActionFilter) SetBrowserType(v string)`

SetBrowserType sets BrowserType field to given value.

### HasBrowserType

`func (o *UserActionFilter) HasBrowserType() bool`

HasBrowserType returns a boolean if a field has been set.

### GetBrowserVersion

`func (o *UserActionFilter) GetBrowserVersion() string`

GetBrowserVersion returns the BrowserVersion field if non-nil, zero value otherwise.

### GetBrowserVersionOk

`func (o *UserActionFilter) GetBrowserVersionOk() (*string, bool)`

GetBrowserVersionOk returns a tuple with the BrowserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserVersion

`func (o *UserActionFilter) SetBrowserVersion(v string)`

SetBrowserVersion sets BrowserVersion field to given value.

### HasBrowserVersion

`func (o *UserActionFilter) HasBrowserVersion() bool`

HasBrowserVersion returns a boolean if a field has been set.

### GetHasCustomErrors

`func (o *UserActionFilter) GetHasCustomErrors() bool`

GetHasCustomErrors returns the HasCustomErrors field if non-nil, zero value otherwise.

### GetHasCustomErrorsOk

`func (o *UserActionFilter) GetHasCustomErrorsOk() (*bool, bool)`

GetHasCustomErrorsOk returns a tuple with the HasCustomErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCustomErrors

`func (o *UserActionFilter) SetHasCustomErrors(v bool)`

SetHasCustomErrors sets HasCustomErrors field to given value.

### HasHasCustomErrors

`func (o *UserActionFilter) HasHasCustomErrors() bool`

HasHasCustomErrors returns a boolean if a field has been set.

### GetHasAnyError

`func (o *UserActionFilter) GetHasAnyError() bool`

GetHasAnyError returns the HasAnyError field if non-nil, zero value otherwise.

### GetHasAnyErrorOk

`func (o *UserActionFilter) GetHasAnyErrorOk() (*bool, bool)`

GetHasAnyErrorOk returns a tuple with the HasAnyError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAnyError

`func (o *UserActionFilter) SetHasAnyError(v bool)`

SetHasAnyError sets HasAnyError field to given value.

### HasHasAnyError

`func (o *UserActionFilter) HasHasAnyError() bool`

HasHasAnyError returns a boolean if a field has been set.

### GetHasHttpErrors

`func (o *UserActionFilter) GetHasHttpErrors() bool`

GetHasHttpErrors returns the HasHttpErrors field if non-nil, zero value otherwise.

### GetHasHttpErrorsOk

`func (o *UserActionFilter) GetHasHttpErrorsOk() (*bool, bool)`

GetHasHttpErrorsOk returns a tuple with the HasHttpErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHttpErrors

`func (o *UserActionFilter) SetHasHttpErrors(v bool)`

SetHasHttpErrors sets HasHttpErrors field to given value.

### HasHasHttpErrors

`func (o *UserActionFilter) HasHasHttpErrors() bool`

HasHasHttpErrors returns a boolean if a field has been set.

### GetHasJavascriptErrors

`func (o *UserActionFilter) GetHasJavascriptErrors() bool`

GetHasJavascriptErrors returns the HasJavascriptErrors field if non-nil, zero value otherwise.

### GetHasJavascriptErrorsOk

`func (o *UserActionFilter) GetHasJavascriptErrorsOk() (*bool, bool)`

GetHasJavascriptErrorsOk returns a tuple with the HasJavascriptErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasJavascriptErrors

`func (o *UserActionFilter) SetHasJavascriptErrors(v bool)`

SetHasJavascriptErrors sets HasJavascriptErrors field to given value.

### HasHasJavascriptErrors

`func (o *UserActionFilter) HasHasJavascriptErrors() bool`

HasHasJavascriptErrors returns a boolean if a field has been set.

### GetCity

`func (o *UserActionFilter) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UserActionFilter) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UserActionFilter) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UserActionFilter) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetContinent

`func (o *UserActionFilter) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *UserActionFilter) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *UserActionFilter) SetContinent(v string)`

SetContinent sets Continent field to given value.

### HasContinent

`func (o *UserActionFilter) HasContinent() bool`

HasContinent returns a boolean if a field has been set.

### GetCountry

`func (o *UserActionFilter) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserActionFilter) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserActionFilter) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UserActionFilter) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetRegion

`func (o *UserActionFilter) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UserActionFilter) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UserActionFilter) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UserActionFilter) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetIp

`func (o *UserActionFilter) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *UserActionFilter) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *UserActionFilter) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *UserActionFilter) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpV6Traffic

`func (o *UserActionFilter) GetIpV6Traffic() bool`

GetIpV6Traffic returns the IpV6Traffic field if non-nil, zero value otherwise.

### GetIpV6TrafficOk

`func (o *UserActionFilter) GetIpV6TrafficOk() (*bool, bool)`

GetIpV6TrafficOk returns a tuple with the IpV6Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Traffic

`func (o *UserActionFilter) SetIpV6Traffic(v bool)`

SetIpV6Traffic sets IpV6Traffic field to given value.

### HasIpV6Traffic

`func (o *UserActionFilter) HasIpV6Traffic() bool`

HasIpV6Traffic returns a boolean if a field has been set.

### GetOsFamily

`func (o *UserActionFilter) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *UserActionFilter) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *UserActionFilter) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *UserActionFilter) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetOsVersion

`func (o *UserActionFilter) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *UserActionFilter) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *UserActionFilter) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *UserActionFilter) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetHttpErrorCode

`func (o *UserActionFilter) GetHttpErrorCode() int32`

GetHttpErrorCode returns the HttpErrorCode field if non-nil, zero value otherwise.

### GetHttpErrorCodeOk

`func (o *UserActionFilter) GetHttpErrorCodeOk() (*int32, bool)`

GetHttpErrorCodeOk returns a tuple with the HttpErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpErrorCode

`func (o *UserActionFilter) SetHttpErrorCode(v int32)`

SetHttpErrorCode sets HttpErrorCode field to given value.

### HasHttpErrorCode

`func (o *UserActionFilter) HasHttpErrorCode() bool`

HasHttpErrorCode returns a boolean if a field has been set.

### GetHttpPath

`func (o *UserActionFilter) GetHttpPath() string`

GetHttpPath returns the HttpPath field if non-nil, zero value otherwise.

### GetHttpPathOk

`func (o *UserActionFilter) GetHttpPathOk() (*string, bool)`

GetHttpPathOk returns a tuple with the HttpPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPath

`func (o *UserActionFilter) SetHttpPath(v string)`

SetHttpPath sets HttpPath field to given value.

### HasHttpPath

`func (o *UserActionFilter) HasHttpPath() bool`

HasHttpPath returns a boolean if a field has been set.

### GetCustomErrorType

`func (o *UserActionFilter) GetCustomErrorType() string`

GetCustomErrorType returns the CustomErrorType field if non-nil, zero value otherwise.

### GetCustomErrorTypeOk

`func (o *UserActionFilter) GetCustomErrorTypeOk() (*string, bool)`

GetCustomErrorTypeOk returns a tuple with the CustomErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomErrorType

`func (o *UserActionFilter) SetCustomErrorType(v string)`

SetCustomErrorType sets CustomErrorType field to given value.

### HasCustomErrorType

`func (o *UserActionFilter) HasCustomErrorType() bool`

HasCustomErrorType returns a boolean if a field has been set.

### GetCustomErrorName

`func (o *UserActionFilter) GetCustomErrorName() string`

GetCustomErrorName returns the CustomErrorName field if non-nil, zero value otherwise.

### GetCustomErrorNameOk

`func (o *UserActionFilter) GetCustomErrorNameOk() (*string, bool)`

GetCustomErrorNameOk returns a tuple with the CustomErrorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomErrorName

`func (o *UserActionFilter) SetCustomErrorName(v string)`

SetCustomErrorName sets CustomErrorName field to given value.

### HasCustomErrorName

`func (o *UserActionFilter) HasCustomErrorName() bool`

HasCustomErrorName returns a boolean if a field has been set.

### GetUserActionProperties

`func (o *UserActionFilter) GetUserActionProperties() []UserActionPropertyFilter`

GetUserActionProperties returns the UserActionProperties field if non-nil, zero value otherwise.

### GetUserActionPropertiesOk

`func (o *UserActionFilter) GetUserActionPropertiesOk() (*[]UserActionPropertyFilter, bool)`

GetUserActionPropertiesOk returns a tuple with the UserActionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionProperties

`func (o *UserActionFilter) SetUserActionProperties(v []UserActionPropertyFilter)`

SetUserActionProperties sets UserActionProperties field to given value.

### HasUserActionProperties

`func (o *UserActionFilter) HasUserActionProperties() bool`

HasUserActionProperties returns a boolean if a field has been set.

### GetTargetViewName

`func (o *UserActionFilter) GetTargetViewName() string`

GetTargetViewName returns the TargetViewName field if non-nil, zero value otherwise.

### GetTargetViewNameOk

`func (o *UserActionFilter) GetTargetViewNameOk() (*string, bool)`

GetTargetViewNameOk returns a tuple with the TargetViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetViewName

`func (o *UserActionFilter) SetTargetViewName(v string)`

SetTargetViewName sets TargetViewName field to given value.

### HasTargetViewName

`func (o *UserActionFilter) HasTargetViewName() bool`

HasTargetViewName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


