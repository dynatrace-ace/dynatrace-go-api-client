# UserSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** | The ID of the Dynatrace environment that captured the user session.   This field can not be queried via the User Session Query Language. | [optional] 
**UserSessionId** | Pointer to **string** | The unique ID of the user session. | [optional] 
**StartTime** | Pointer to **int64** | The timestamp of the first user action in the user session, in UTC milliseconds. | [optional] 
**EndTime** | Pointer to **int64** | The timestamp of the last user action in the user session, in UTC milliseconds. | [optional] 
**Duration** | Pointer to **int64** | The duration of the user session, in milliseconds.    This is calculated as the amount of time between the start of the first user action and the end of the last user action. | [optional] 
**InternalUserId** | Pointer to **string** | The unique ID of the user that triggered the user session. | [optional] 
**UserType** | Pointer to **string** | The type of the user. Indicates either a real human user (&#x60;REAL_USER&#x60;) or a robot (&#x60;ROBOT&#x60; or &#x60;SYNTHETIC&#x60;). | [optional] 
**ReasonForNoSessionReplay** | Pointer to **string** | The reason for no session replay. | [optional] 
**ApplicationType** | Pointer to **string** | The type of the application used in the user session. | [optional] 
**Bounce** | Pointer to **bool** | The user session has (&#x60;true&#x60;) or doesn&#39;t have (&#x60;false&#x60;) a bounce.    A bounce means there is only one user action in the user session. | [optional] 
**NewUser** | Pointer to **bool** | The user is a first-time (&#x60;true&#x60;) or a returning user (&#x60;false&#x60;). | [optional] 
**UserActionCount** | Pointer to **int32** | The number of user actions in the user session. | [optional] 
**TotalErrorCount** | Pointer to **int32** | The number of errors detected in the user session. | [optional] 
**TotalLicenseCreditCount** | Pointer to **int32** | Count of the license credits | [optional] 
**MatchingConversionGoalsCount** | Pointer to **int32** | The number of conversion goals achieved by the user session. | [optional] 
**MatchingConversionGoals** | Pointer to **[]string** | A list of conversion goals achieved by the user session.    Additionally, you can define conversion goals for a single user action. | [optional] 
**Ip** | Pointer to **string** | The IP address (IPv4 or IPv6) from which the user session originates. | [optional] 
**Continent** | Pointer to **string** | The continent from which the user session originates (based on the IP address). | [optional] 
**Country** | Pointer to **string** | The country from which the user session originates (based on the IP address). | [optional] 
**Region** | Pointer to **string** | The region from which the user session originates (based on the IP address). | [optional] 
**City** | Pointer to **string** | The city from which the user session originates (based on the IP address). | [optional] 
**BrowserType** | Pointer to **string** | The type of browser used for the user session. | [optional] 
**BrowserFamily** | Pointer to **string** | The family of the browser used for the user session. | [optional] 
**BrowserMajorVersion** | Pointer to **string** | The version of the browser used for the user session. | [optional] 
**OsFamily** | Pointer to **string** | The type of operating system used for the user session. | [optional] 
**OsVersion** | Pointer to **string** | The version of the operating system used for the user session. | [optional] 
**Manufacturer** | Pointer to **string** | The detected manufacturer of the device used for the user session. | [optional] 
**Device** | Pointer to **string** | The detected device used for the user session. | [optional] 
**UserId** | Pointer to **string** | The user ID provided for the user session by session tagging. | [optional] 
**ScreenHeight** | Pointer to **int32** | The detected screen height of the device used for the user session. | [optional] 
**ScreenWidth** | Pointer to **int32** | The detected screen width of the device used for the user session. | [optional] 
**ScreenOrientation** | Pointer to **string** | The detected screen orientation of the device used on the device for the user session. | [optional] 
**DisplayResolution** | Pointer to **string** | The detected screen resolution of the device used for the user session. | [optional] 
**HasCrash** | Pointer to **bool** | The user session includes (&#x60;true&#x60;) or doesn&#39;t include (&#x60;false&#x60;) a crash. | [optional] 
**HasSessionReplay** | Pointer to **bool** | Indicates whether the user session has Session Replay. | [optional] 
**Isp** | Pointer to **string** | The internet service provider from which the user session originates (based on the IP address). | [optional] 
**ClientType** | Pointer to **string** | Additional information about the client.   This field can not be queried via the user session query language. Use the **browserType** field instead. | [optional] 
**BrowserMonitorId** | Pointer to **string** | Reports the id of the browser monitor | [optional] 
**BrowserMonitorName** | Pointer to **string** | Reports the name of the browser monitor | [optional] 
**StringProperties** | Pointer to [**[]StringProperty**](StringProperty.md) | A list of custom properties of the user session with string values. | [optional] 
**LongProperties** | Pointer to [**[]LongProperty**](LongProperty.md) | A list of custom properties of the user session with integer (short or long) values. | [optional] 
**DoubleProperties** | Pointer to [**[]DoubleProperty**](DoubleProperty.md) | A list of custom properties of the user session with floating-point numerical values. | [optional] 
**DateProperties** | Pointer to [**[]DateProperty**](DateProperty.md) | A list of custom properties of the user session with date values. | [optional] 
**UserActions** | Pointer to [**[]UserSessionUserAction**](UserSessionUserAction.md) | A list of user actions recorded in the user session. | [optional] 
**Events** | Pointer to [**[]UserSessionEvents**](UserSessionEvents.md) | A list of additional events recorded in the user session. | [optional] 
**Errors** | Pointer to [**[]UserSessionErrors**](UserSessionErrors.md) | A list of errors recorded in the user session. | [optional] 
**SyntheticEvents** | Pointer to [**[]UserSessionSyntheticEvent**](UserSessionSyntheticEvent.md) | A list of synthetic events recorded in the user session. | [optional] 
**AppVersion** | Pointer to **string** | The version of the application where the user session has been recorded.    This information is provided by another integration, such as OpenKit. | [optional] 
**EndReason** | Pointer to **string** | The reason for the end of the user session. | [optional] 
**NumberOfRageClicks** | Pointer to **int32** | The number of rage clicks detected in the user session. | [optional] 
**UserExperienceScore** | Pointer to **string** | The user experience score of the user session. | [optional] 
**Carrier** | Pointer to **string** | The carrier information of the mobile user session. | [optional] 
**NetworkTechnology** | Pointer to **string** | The network technology information of the mobile user session. | [optional] 
**ConnectionType** | Pointer to **string** | The serialized connection type of the mobile user session. | [optional] 
**ReplayStart** | Pointer to **int64** | Epoch time when replay starts | [optional] 
**ReplayEnd** | Pointer to **int64** | Epoch time when replay ends | [optional] 
**ClientTimeOffset** | Pointer to **int32** | Client time offset in milliseconds | [optional] 

## Methods

### NewUserSession

`func NewUserSession() *UserSession`

NewUserSession instantiates a new UserSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionWithDefaults

`func NewUserSessionWithDefaults() *UserSession`

NewUserSessionWithDefaults instantiates a new UserSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *UserSession) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UserSession) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UserSession) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UserSession) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUserSessionId

`func (o *UserSession) GetUserSessionId() string`

GetUserSessionId returns the UserSessionId field if non-nil, zero value otherwise.

### GetUserSessionIdOk

`func (o *UserSession) GetUserSessionIdOk() (*string, bool)`

GetUserSessionIdOk returns a tuple with the UserSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSessionId

`func (o *UserSession) SetUserSessionId(v string)`

SetUserSessionId sets UserSessionId field to given value.

### HasUserSessionId

`func (o *UserSession) HasUserSessionId() bool`

HasUserSessionId returns a boolean if a field has been set.

### GetStartTime

`func (o *UserSession) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UserSession) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UserSession) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UserSession) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *UserSession) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UserSession) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UserSession) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UserSession) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetDuration

`func (o *UserSession) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UserSession) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UserSession) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *UserSession) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetInternalUserId

`func (o *UserSession) GetInternalUserId() string`

GetInternalUserId returns the InternalUserId field if non-nil, zero value otherwise.

### GetInternalUserIdOk

`func (o *UserSession) GetInternalUserIdOk() (*string, bool)`

GetInternalUserIdOk returns a tuple with the InternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalUserId

`func (o *UserSession) SetInternalUserId(v string)`

SetInternalUserId sets InternalUserId field to given value.

### HasInternalUserId

`func (o *UserSession) HasInternalUserId() bool`

HasInternalUserId returns a boolean if a field has been set.

### GetUserType

`func (o *UserSession) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserSession) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserSession) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserSession) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetReasonForNoSessionReplay

`func (o *UserSession) GetReasonForNoSessionReplay() string`

GetReasonForNoSessionReplay returns the ReasonForNoSessionReplay field if non-nil, zero value otherwise.

### GetReasonForNoSessionReplayOk

`func (o *UserSession) GetReasonForNoSessionReplayOk() (*string, bool)`

GetReasonForNoSessionReplayOk returns a tuple with the ReasonForNoSessionReplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForNoSessionReplay

`func (o *UserSession) SetReasonForNoSessionReplay(v string)`

SetReasonForNoSessionReplay sets ReasonForNoSessionReplay field to given value.

### HasReasonForNoSessionReplay

`func (o *UserSession) HasReasonForNoSessionReplay() bool`

HasReasonForNoSessionReplay returns a boolean if a field has been set.

### GetApplicationType

`func (o *UserSession) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *UserSession) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *UserSession) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *UserSession) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetBounce

`func (o *UserSession) GetBounce() bool`

GetBounce returns the Bounce field if non-nil, zero value otherwise.

### GetBounceOk

`func (o *UserSession) GetBounceOk() (*bool, bool)`

GetBounceOk returns a tuple with the Bounce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounce

`func (o *UserSession) SetBounce(v bool)`

SetBounce sets Bounce field to given value.

### HasBounce

`func (o *UserSession) HasBounce() bool`

HasBounce returns a boolean if a field has been set.

### GetNewUser

`func (o *UserSession) GetNewUser() bool`

GetNewUser returns the NewUser field if non-nil, zero value otherwise.

### GetNewUserOk

`func (o *UserSession) GetNewUserOk() (*bool, bool)`

GetNewUserOk returns a tuple with the NewUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewUser

`func (o *UserSession) SetNewUser(v bool)`

SetNewUser sets NewUser field to given value.

### HasNewUser

`func (o *UserSession) HasNewUser() bool`

HasNewUser returns a boolean if a field has been set.

### GetUserActionCount

`func (o *UserSession) GetUserActionCount() int32`

GetUserActionCount returns the UserActionCount field if non-nil, zero value otherwise.

### GetUserActionCountOk

`func (o *UserSession) GetUserActionCountOk() (*int32, bool)`

GetUserActionCountOk returns a tuple with the UserActionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionCount

`func (o *UserSession) SetUserActionCount(v int32)`

SetUserActionCount sets UserActionCount field to given value.

### HasUserActionCount

`func (o *UserSession) HasUserActionCount() bool`

HasUserActionCount returns a boolean if a field has been set.

### GetTotalErrorCount

`func (o *UserSession) GetTotalErrorCount() int32`

GetTotalErrorCount returns the TotalErrorCount field if non-nil, zero value otherwise.

### GetTotalErrorCountOk

`func (o *UserSession) GetTotalErrorCountOk() (*int32, bool)`

GetTotalErrorCountOk returns a tuple with the TotalErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalErrorCount

`func (o *UserSession) SetTotalErrorCount(v int32)`

SetTotalErrorCount sets TotalErrorCount field to given value.

### HasTotalErrorCount

`func (o *UserSession) HasTotalErrorCount() bool`

HasTotalErrorCount returns a boolean if a field has been set.

### GetTotalLicenseCreditCount

`func (o *UserSession) GetTotalLicenseCreditCount() int32`

GetTotalLicenseCreditCount returns the TotalLicenseCreditCount field if non-nil, zero value otherwise.

### GetTotalLicenseCreditCountOk

`func (o *UserSession) GetTotalLicenseCreditCountOk() (*int32, bool)`

GetTotalLicenseCreditCountOk returns a tuple with the TotalLicenseCreditCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLicenseCreditCount

`func (o *UserSession) SetTotalLicenseCreditCount(v int32)`

SetTotalLicenseCreditCount sets TotalLicenseCreditCount field to given value.

### HasTotalLicenseCreditCount

`func (o *UserSession) HasTotalLicenseCreditCount() bool`

HasTotalLicenseCreditCount returns a boolean if a field has been set.

### GetMatchingConversionGoalsCount

`func (o *UserSession) GetMatchingConversionGoalsCount() int32`

GetMatchingConversionGoalsCount returns the MatchingConversionGoalsCount field if non-nil, zero value otherwise.

### GetMatchingConversionGoalsCountOk

`func (o *UserSession) GetMatchingConversionGoalsCountOk() (*int32, bool)`

GetMatchingConversionGoalsCountOk returns a tuple with the MatchingConversionGoalsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingConversionGoalsCount

`func (o *UserSession) SetMatchingConversionGoalsCount(v int32)`

SetMatchingConversionGoalsCount sets MatchingConversionGoalsCount field to given value.

### HasMatchingConversionGoalsCount

`func (o *UserSession) HasMatchingConversionGoalsCount() bool`

HasMatchingConversionGoalsCount returns a boolean if a field has been set.

### GetMatchingConversionGoals

`func (o *UserSession) GetMatchingConversionGoals() []string`

GetMatchingConversionGoals returns the MatchingConversionGoals field if non-nil, zero value otherwise.

### GetMatchingConversionGoalsOk

`func (o *UserSession) GetMatchingConversionGoalsOk() (*[]string, bool)`

GetMatchingConversionGoalsOk returns a tuple with the MatchingConversionGoals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingConversionGoals

`func (o *UserSession) SetMatchingConversionGoals(v []string)`

SetMatchingConversionGoals sets MatchingConversionGoals field to given value.

### HasMatchingConversionGoals

`func (o *UserSession) HasMatchingConversionGoals() bool`

HasMatchingConversionGoals returns a boolean if a field has been set.

### GetIp

`func (o *UserSession) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *UserSession) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *UserSession) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *UserSession) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetContinent

`func (o *UserSession) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *UserSession) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *UserSession) SetContinent(v string)`

SetContinent sets Continent field to given value.

### HasContinent

`func (o *UserSession) HasContinent() bool`

HasContinent returns a boolean if a field has been set.

### GetCountry

`func (o *UserSession) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserSession) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserSession) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UserSession) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetRegion

`func (o *UserSession) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UserSession) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UserSession) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UserSession) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCity

`func (o *UserSession) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UserSession) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UserSession) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UserSession) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetBrowserType

`func (o *UserSession) GetBrowserType() string`

GetBrowserType returns the BrowserType field if non-nil, zero value otherwise.

### GetBrowserTypeOk

`func (o *UserSession) GetBrowserTypeOk() (*string, bool)`

GetBrowserTypeOk returns a tuple with the BrowserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserType

`func (o *UserSession) SetBrowserType(v string)`

SetBrowserType sets BrowserType field to given value.

### HasBrowserType

`func (o *UserSession) HasBrowserType() bool`

HasBrowserType returns a boolean if a field has been set.

### GetBrowserFamily

`func (o *UserSession) GetBrowserFamily() string`

GetBrowserFamily returns the BrowserFamily field if non-nil, zero value otherwise.

### GetBrowserFamilyOk

`func (o *UserSession) GetBrowserFamilyOk() (*string, bool)`

GetBrowserFamilyOk returns a tuple with the BrowserFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserFamily

`func (o *UserSession) SetBrowserFamily(v string)`

SetBrowserFamily sets BrowserFamily field to given value.

### HasBrowserFamily

`func (o *UserSession) HasBrowserFamily() bool`

HasBrowserFamily returns a boolean if a field has been set.

### GetBrowserMajorVersion

`func (o *UserSession) GetBrowserMajorVersion() string`

GetBrowserMajorVersion returns the BrowserMajorVersion field if non-nil, zero value otherwise.

### GetBrowserMajorVersionOk

`func (o *UserSession) GetBrowserMajorVersionOk() (*string, bool)`

GetBrowserMajorVersionOk returns a tuple with the BrowserMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserMajorVersion

`func (o *UserSession) SetBrowserMajorVersion(v string)`

SetBrowserMajorVersion sets BrowserMajorVersion field to given value.

### HasBrowserMajorVersion

`func (o *UserSession) HasBrowserMajorVersion() bool`

HasBrowserMajorVersion returns a boolean if a field has been set.

### GetOsFamily

`func (o *UserSession) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *UserSession) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *UserSession) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *UserSession) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetOsVersion

`func (o *UserSession) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *UserSession) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *UserSession) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *UserSession) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetManufacturer

`func (o *UserSession) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *UserSession) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *UserSession) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *UserSession) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetDevice

`func (o *UserSession) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *UserSession) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *UserSession) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *UserSession) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetUserId

`func (o *UserSession) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserSession) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserSession) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserSession) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetScreenHeight

`func (o *UserSession) GetScreenHeight() int32`

GetScreenHeight returns the ScreenHeight field if non-nil, zero value otherwise.

### GetScreenHeightOk

`func (o *UserSession) GetScreenHeightOk() (*int32, bool)`

GetScreenHeightOk returns a tuple with the ScreenHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenHeight

`func (o *UserSession) SetScreenHeight(v int32)`

SetScreenHeight sets ScreenHeight field to given value.

### HasScreenHeight

`func (o *UserSession) HasScreenHeight() bool`

HasScreenHeight returns a boolean if a field has been set.

### GetScreenWidth

`func (o *UserSession) GetScreenWidth() int32`

GetScreenWidth returns the ScreenWidth field if non-nil, zero value otherwise.

### GetScreenWidthOk

`func (o *UserSession) GetScreenWidthOk() (*int32, bool)`

GetScreenWidthOk returns a tuple with the ScreenWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenWidth

`func (o *UserSession) SetScreenWidth(v int32)`

SetScreenWidth sets ScreenWidth field to given value.

### HasScreenWidth

`func (o *UserSession) HasScreenWidth() bool`

HasScreenWidth returns a boolean if a field has been set.

### GetScreenOrientation

`func (o *UserSession) GetScreenOrientation() string`

GetScreenOrientation returns the ScreenOrientation field if non-nil, zero value otherwise.

### GetScreenOrientationOk

`func (o *UserSession) GetScreenOrientationOk() (*string, bool)`

GetScreenOrientationOk returns a tuple with the ScreenOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenOrientation

`func (o *UserSession) SetScreenOrientation(v string)`

SetScreenOrientation sets ScreenOrientation field to given value.

### HasScreenOrientation

`func (o *UserSession) HasScreenOrientation() bool`

HasScreenOrientation returns a boolean if a field has been set.

### GetDisplayResolution

`func (o *UserSession) GetDisplayResolution() string`

GetDisplayResolution returns the DisplayResolution field if non-nil, zero value otherwise.

### GetDisplayResolutionOk

`func (o *UserSession) GetDisplayResolutionOk() (*string, bool)`

GetDisplayResolutionOk returns a tuple with the DisplayResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayResolution

`func (o *UserSession) SetDisplayResolution(v string)`

SetDisplayResolution sets DisplayResolution field to given value.

### HasDisplayResolution

`func (o *UserSession) HasDisplayResolution() bool`

HasDisplayResolution returns a boolean if a field has been set.

### GetHasCrash

`func (o *UserSession) GetHasCrash() bool`

GetHasCrash returns the HasCrash field if non-nil, zero value otherwise.

### GetHasCrashOk

`func (o *UserSession) GetHasCrashOk() (*bool, bool)`

GetHasCrashOk returns a tuple with the HasCrash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCrash

`func (o *UserSession) SetHasCrash(v bool)`

SetHasCrash sets HasCrash field to given value.

### HasHasCrash

`func (o *UserSession) HasHasCrash() bool`

HasHasCrash returns a boolean if a field has been set.

### GetHasSessionReplay

`func (o *UserSession) GetHasSessionReplay() bool`

GetHasSessionReplay returns the HasSessionReplay field if non-nil, zero value otherwise.

### GetHasSessionReplayOk

`func (o *UserSession) GetHasSessionReplayOk() (*bool, bool)`

GetHasSessionReplayOk returns a tuple with the HasSessionReplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSessionReplay

`func (o *UserSession) SetHasSessionReplay(v bool)`

SetHasSessionReplay sets HasSessionReplay field to given value.

### HasHasSessionReplay

`func (o *UserSession) HasHasSessionReplay() bool`

HasHasSessionReplay returns a boolean if a field has been set.

### GetIsp

`func (o *UserSession) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *UserSession) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *UserSession) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *UserSession) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetClientType

`func (o *UserSession) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *UserSession) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *UserSession) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *UserSession) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetBrowserMonitorId

`func (o *UserSession) GetBrowserMonitorId() string`

GetBrowserMonitorId returns the BrowserMonitorId field if non-nil, zero value otherwise.

### GetBrowserMonitorIdOk

`func (o *UserSession) GetBrowserMonitorIdOk() (*string, bool)`

GetBrowserMonitorIdOk returns a tuple with the BrowserMonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserMonitorId

`func (o *UserSession) SetBrowserMonitorId(v string)`

SetBrowserMonitorId sets BrowserMonitorId field to given value.

### HasBrowserMonitorId

`func (o *UserSession) HasBrowserMonitorId() bool`

HasBrowserMonitorId returns a boolean if a field has been set.

### GetBrowserMonitorName

`func (o *UserSession) GetBrowserMonitorName() string`

GetBrowserMonitorName returns the BrowserMonitorName field if non-nil, zero value otherwise.

### GetBrowserMonitorNameOk

`func (o *UserSession) GetBrowserMonitorNameOk() (*string, bool)`

GetBrowserMonitorNameOk returns a tuple with the BrowserMonitorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserMonitorName

`func (o *UserSession) SetBrowserMonitorName(v string)`

SetBrowserMonitorName sets BrowserMonitorName field to given value.

### HasBrowserMonitorName

`func (o *UserSession) HasBrowserMonitorName() bool`

HasBrowserMonitorName returns a boolean if a field has been set.

### GetStringProperties

`func (o *UserSession) GetStringProperties() []StringProperty`

GetStringProperties returns the StringProperties field if non-nil, zero value otherwise.

### GetStringPropertiesOk

`func (o *UserSession) GetStringPropertiesOk() (*[]StringProperty, bool)`

GetStringPropertiesOk returns a tuple with the StringProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringProperties

`func (o *UserSession) SetStringProperties(v []StringProperty)`

SetStringProperties sets StringProperties field to given value.

### HasStringProperties

`func (o *UserSession) HasStringProperties() bool`

HasStringProperties returns a boolean if a field has been set.

### GetLongProperties

`func (o *UserSession) GetLongProperties() []LongProperty`

GetLongProperties returns the LongProperties field if non-nil, zero value otherwise.

### GetLongPropertiesOk

`func (o *UserSession) GetLongPropertiesOk() (*[]LongProperty, bool)`

GetLongPropertiesOk returns a tuple with the LongProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongProperties

`func (o *UserSession) SetLongProperties(v []LongProperty)`

SetLongProperties sets LongProperties field to given value.

### HasLongProperties

`func (o *UserSession) HasLongProperties() bool`

HasLongProperties returns a boolean if a field has been set.

### GetDoubleProperties

`func (o *UserSession) GetDoubleProperties() []DoubleProperty`

GetDoubleProperties returns the DoubleProperties field if non-nil, zero value otherwise.

### GetDoublePropertiesOk

`func (o *UserSession) GetDoublePropertiesOk() (*[]DoubleProperty, bool)`

GetDoublePropertiesOk returns a tuple with the DoubleProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleProperties

`func (o *UserSession) SetDoubleProperties(v []DoubleProperty)`

SetDoubleProperties sets DoubleProperties field to given value.

### HasDoubleProperties

`func (o *UserSession) HasDoubleProperties() bool`

HasDoubleProperties returns a boolean if a field has been set.

### GetDateProperties

`func (o *UserSession) GetDateProperties() []DateProperty`

GetDateProperties returns the DateProperties field if non-nil, zero value otherwise.

### GetDatePropertiesOk

`func (o *UserSession) GetDatePropertiesOk() (*[]DateProperty, bool)`

GetDatePropertiesOk returns a tuple with the DateProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateProperties

`func (o *UserSession) SetDateProperties(v []DateProperty)`

SetDateProperties sets DateProperties field to given value.

### HasDateProperties

`func (o *UserSession) HasDateProperties() bool`

HasDateProperties returns a boolean if a field has been set.

### GetUserActions

`func (o *UserSession) GetUserActions() []UserSessionUserAction`

GetUserActions returns the UserActions field if non-nil, zero value otherwise.

### GetUserActionsOk

`func (o *UserSession) GetUserActionsOk() (*[]UserSessionUserAction, bool)`

GetUserActionsOk returns a tuple with the UserActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActions

`func (o *UserSession) SetUserActions(v []UserSessionUserAction)`

SetUserActions sets UserActions field to given value.

### HasUserActions

`func (o *UserSession) HasUserActions() bool`

HasUserActions returns a boolean if a field has been set.

### GetEvents

`func (o *UserSession) GetEvents() []UserSessionEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *UserSession) GetEventsOk() (*[]UserSessionEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *UserSession) SetEvents(v []UserSessionEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *UserSession) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetErrors

`func (o *UserSession) GetErrors() []UserSessionErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UserSession) GetErrorsOk() (*[]UserSessionErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UserSession) SetErrors(v []UserSessionErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *UserSession) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSyntheticEvents

`func (o *UserSession) GetSyntheticEvents() []UserSessionSyntheticEvent`

GetSyntheticEvents returns the SyntheticEvents field if non-nil, zero value otherwise.

### GetSyntheticEventsOk

`func (o *UserSession) GetSyntheticEventsOk() (*[]UserSessionSyntheticEvent, bool)`

GetSyntheticEventsOk returns a tuple with the SyntheticEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticEvents

`func (o *UserSession) SetSyntheticEvents(v []UserSessionSyntheticEvent)`

SetSyntheticEvents sets SyntheticEvents field to given value.

### HasSyntheticEvents

`func (o *UserSession) HasSyntheticEvents() bool`

HasSyntheticEvents returns a boolean if a field has been set.

### GetAppVersion

`func (o *UserSession) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *UserSession) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *UserSession) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *UserSession) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetEndReason

`func (o *UserSession) GetEndReason() string`

GetEndReason returns the EndReason field if non-nil, zero value otherwise.

### GetEndReasonOk

`func (o *UserSession) GetEndReasonOk() (*string, bool)`

GetEndReasonOk returns a tuple with the EndReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndReason

`func (o *UserSession) SetEndReason(v string)`

SetEndReason sets EndReason field to given value.

### HasEndReason

`func (o *UserSession) HasEndReason() bool`

HasEndReason returns a boolean if a field has been set.

### GetNumberOfRageClicks

`func (o *UserSession) GetNumberOfRageClicks() int32`

GetNumberOfRageClicks returns the NumberOfRageClicks field if non-nil, zero value otherwise.

### GetNumberOfRageClicksOk

`func (o *UserSession) GetNumberOfRageClicksOk() (*int32, bool)`

GetNumberOfRageClicksOk returns a tuple with the NumberOfRageClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRageClicks

`func (o *UserSession) SetNumberOfRageClicks(v int32)`

SetNumberOfRageClicks sets NumberOfRageClicks field to given value.

### HasNumberOfRageClicks

`func (o *UserSession) HasNumberOfRageClicks() bool`

HasNumberOfRageClicks returns a boolean if a field has been set.

### GetUserExperienceScore

`func (o *UserSession) GetUserExperienceScore() string`

GetUserExperienceScore returns the UserExperienceScore field if non-nil, zero value otherwise.

### GetUserExperienceScoreOk

`func (o *UserSession) GetUserExperienceScoreOk() (*string, bool)`

GetUserExperienceScoreOk returns a tuple with the UserExperienceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserExperienceScore

`func (o *UserSession) SetUserExperienceScore(v string)`

SetUserExperienceScore sets UserExperienceScore field to given value.

### HasUserExperienceScore

`func (o *UserSession) HasUserExperienceScore() bool`

HasUserExperienceScore returns a boolean if a field has been set.

### GetCarrier

`func (o *UserSession) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *UserSession) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *UserSession) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *UserSession) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetNetworkTechnology

`func (o *UserSession) GetNetworkTechnology() string`

GetNetworkTechnology returns the NetworkTechnology field if non-nil, zero value otherwise.

### GetNetworkTechnologyOk

`func (o *UserSession) GetNetworkTechnologyOk() (*string, bool)`

GetNetworkTechnologyOk returns a tuple with the NetworkTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTechnology

`func (o *UserSession) SetNetworkTechnology(v string)`

SetNetworkTechnology sets NetworkTechnology field to given value.

### HasNetworkTechnology

`func (o *UserSession) HasNetworkTechnology() bool`

HasNetworkTechnology returns a boolean if a field has been set.

### GetConnectionType

`func (o *UserSession) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *UserSession) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *UserSession) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *UserSession) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetReplayStart

`func (o *UserSession) GetReplayStart() int64`

GetReplayStart returns the ReplayStart field if non-nil, zero value otherwise.

### GetReplayStartOk

`func (o *UserSession) GetReplayStartOk() (*int64, bool)`

GetReplayStartOk returns a tuple with the ReplayStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayStart

`func (o *UserSession) SetReplayStart(v int64)`

SetReplayStart sets ReplayStart field to given value.

### HasReplayStart

`func (o *UserSession) HasReplayStart() bool`

HasReplayStart returns a boolean if a field has been set.

### GetReplayEnd

`func (o *UserSession) GetReplayEnd() int64`

GetReplayEnd returns the ReplayEnd field if non-nil, zero value otherwise.

### GetReplayEndOk

`func (o *UserSession) GetReplayEndOk() (*int64, bool)`

GetReplayEndOk returns a tuple with the ReplayEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayEnd

`func (o *UserSession) SetReplayEnd(v int64)`

SetReplayEnd sets ReplayEnd field to given value.

### HasReplayEnd

`func (o *UserSession) HasReplayEnd() bool`

HasReplayEnd returns a boolean if a field has been set.

### GetClientTimeOffset

`func (o *UserSession) GetClientTimeOffset() int32`

GetClientTimeOffset returns the ClientTimeOffset field if non-nil, zero value otherwise.

### GetClientTimeOffsetOk

`func (o *UserSession) GetClientTimeOffsetOk() (*int32, bool)`

GetClientTimeOffsetOk returns a tuple with the ClientTimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTimeOffset

`func (o *UserSession) SetClientTimeOffset(v int32)`

SetClientTimeOffset sets ClientTimeOffset field to given value.

### HasClientTimeOffset

`func (o *UserSession) HasClientTimeOffset() bool`

HasClientTimeOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


