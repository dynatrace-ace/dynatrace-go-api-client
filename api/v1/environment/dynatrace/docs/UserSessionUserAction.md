# UserSessionUserAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the user action.    Typically, this is the name of the page that is loaded as part of a user action or a textual description of the action, such as a mouse click. | [optional] 
**Domain** | Pointer to **string** | The DNS domain where the user action has been recorded. | [optional] 
**TargetUrl** | Pointer to **string** | The target URL of the user action. | [optional] 
**Type** | Pointer to **string** | The type of the user action. | [optional] 
**StartTime** | Pointer to **int64** | The start timestamp of the user action, in UTC milliseconds. | [optional] 
**EndTime** | Pointer to **int64** | The end timestamp of the user action, in UTC milliseconds. | [optional] 
**Duration** | Pointer to **int64** | The duration of the user action, in milliseconds.    This is calculated as the of time between the start and the end timestamps of the user action. | [optional] 
**Application** | Pointer to **string** | The name of the application where the user action has been recorded. | [optional] 
**InternalApplicationId** | Pointer to **string** | The Dynatrace entity ID of the application where the user action has been recorded.    This information is useful when calling various REST APIs, for example as a key for time series queries. | [optional] 
**InternalKeyUserActionId** | Pointer to **string** | The Dynatrace entity ID of the key user action | [optional] 
**SpeedIndex** | Pointer to **int32** | The [speed index](https://dt-url.net/qk1a3r19) of the user action, in milliseconds.    This is calculated as average time it takes for all visible parts of a page to display. | [optional] 
**ErrorCount** | Pointer to **int32** | The number of errors detected in the user action.   This value is deprecated and only available for tenants created before version 206. It will be removed eventually. | [optional] 
**ApdexCategory** | Pointer to **string** | The [user experience index](https://dt-url.net/apdexdoc) of the user action. | [optional] 
**MatchingConversionGoals** | Pointer to **[]string** | A list of conversion goals achieved by the user action.    Additionally, you can define conversion goals for a user session as a whole. | [optional] 
**NetworkTime** | Pointer to **int32** | The amount of time spent on the data transfer for the user action, in milliseconds. | [optional] 
**ServerTime** | Pointer to **int32** | The amount of time spent on the server-side processing for the user action, in milliseconds. | [optional] 
**FrontendTime** | Pointer to **int32** | The amount of time spent on the frontend rendering for the user action, in milliseconds. | [optional] 
**DocumentInteractiveTime** | Pointer to **int32** | The amount of time spent until the document for the user action became interactive, in milliseconds. | [optional] 
**FailedImages** | Pointer to **int32** | The number of failed image loads in the user action.   This value is deprecated and only available for tenants created before version 206. It will be removed eventually. | [optional] 
**FailedXhrRequests** | Pointer to **int32** | The number of failed AJAX requests in the user action.   This value is deprecated and only available for tenants created before version 206. It will be removed eventually. | [optional] 
**HttpRequestsWithErrors** | Pointer to **int32** | The number of HTTP requests with erroneous response codes in the user action where the response code indicates a failed state.   This value is deprecated and only available for tenants created before version 206. It will be removed eventually. | [optional] 
**ThirdPartyResources** | Pointer to **int32** | The number of third party resources loaded for the user action. | [optional] 
**ThirdPartyBusyTime** | Pointer to **int32** | The time spent waiting for third party resources for the user action, in milliseconds. | [optional] 
**CdnResources** | Pointer to **int32** | The number of resources fetched from a CDN for the user action. | [optional] 
**CdnBusyTime** | Pointer to **int32** | The time spent waiting for CDN resources for the user action, in milliseconds. | [optional] 
**FirstPartyResources** | Pointer to **int32** | The number of resources fetched from the originating server for the user action. | [optional] 
**FirstPartyBusyTime** | Pointer to **int32** | The time spent waiting for resources from the originating server for the user action, in milliseconds. | [optional] 
**HasCrash** | Pointer to **bool** | The user action has (&#x60;true&#x60;) or doesn&#39;t have (&#x60;false&#x60;) a crash. | [optional] 
**DomCompleteTime** | Pointer to **int32** | The amount of time until the DOM tree is completed, in milliseconds. | [optional] 
**DomContentLoadedTime** | Pointer to **int32** | The amount of time until the DOM tree is loaded, in milliseconds. | [optional] 
**LoadEventStart** | Pointer to **int32** | The amount of time until the load event started, in milliseconds. | [optional] 
**LoadEventEnd** | Pointer to **int32** | The amount of time until the load event ended, in milliseconds. | [optional] 
**NavigationStart** | Pointer to **int64** | The amount of time until the navigation started, in milliseconds. | [optional] 
**RequestStart** | Pointer to **int32** | The amount of time until the request started, in milliseconds. | [optional] 
**ResponseStart** | Pointer to **int32** | The amount of time until the response started, in milliseconds. | [optional] 
**ResponseEnd** | Pointer to **int32** | The amount of time until the response ended, in milliseconds. | [optional] 
**VisuallyCompleteTime** | Pointer to **int32** | The amount of time until the page is [visually complete](https://dt-url.net/qk1a3r19), in milliseconds. | [optional] 
**SyntheticEvent** | Pointer to **string** | The name of the [Synthetic event](https://dt-url.net/dq1e3rmm) that triggered the user action. | [optional] 
**SyntheticEventId** | Pointer to **string** | The ID of the [Synthetic event](https://dt-url.net/dq1e3rmm) that triggered the user action. | [optional] 
**KeyUserAction** | Pointer to **bool** | Whether the action is a key action | [optional] 
**StringProperties** | Pointer to [**[]StringProperty**](StringProperty.md) | A list of custom properties of the user session with string values. | [optional] 
**LongProperties** | Pointer to [**[]LongProperty**](LongProperty.md) | A list of custom properties of the user session with integer (short or long) values. | [optional] 
**DoubleProperties** | Pointer to [**[]DoubleProperty**](DoubleProperty.md) | A list of custom properties of the user session with floating-point numerical values. | [optional] 
**DateProperties** | Pointer to [**[]DateProperty**](DateProperty.md) | A list of custom properties of the user session with date values. | [optional] 
**UserActionPropertyCount** | Pointer to **int32** | The total number of user action properties in the user action. | [optional] 
**CustomErrorCount** | Pointer to **int32** | The total number of custom errors during the user action. | [optional] 
**JavascriptErrorCount** | Pointer to **int32** | The total number of Javascript errors during the user action. | [optional] 
**RequestErrorCount** | Pointer to **int32** | The total number of request errors during the user action. | [optional] 
**LargestContentfulPaint** | Pointer to **int32** | The largest contentful paint (LCP) is the time that the largest element on the page took to render, in milliseconds. It is an important, user-centric metric for measuring perceived load speed because it marks the point in the page load timeline when the page&#39;s main content has likely loaded. A fast LCP helps reassure the user that the page is useful. | [optional] 
**FirstInputDelay** | Pointer to **int32** | The first input delay (FID) is the time the browser took to respond to the first user input event, in milliseconds. It is an important, user-centric metric for measuring load responsiveness because it quantifies the experience users feel when trying to interact with unresponsive pages. A low FID helps ensure that the page is usable. | [optional] 
**TotalBlockingTime** | Pointer to **int32** | The measurement of the total amount of time between first contentful paint (FCP) and time to interactive (TTI) where the browser was blocked for long enough to prevent input responsiveness. | [optional] 
**CumulativeLayoutShift** | Pointer to **float64** | The cumulative layout shift (CLS) is the sum total of all individual layout shift scores for every unexpected layout shift that occurs during the entire lifespan of the page. It is an important, user-centric metric for measuring visual stability because it helps quantify how often users experience unexpected layout shifts. A low CLS helps ensure that the page is delightful. | [optional] 

## Methods

### NewUserSessionUserAction

`func NewUserSessionUserAction() *UserSessionUserAction`

NewUserSessionUserAction instantiates a new UserSessionUserAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionUserActionWithDefaults

`func NewUserSessionUserActionWithDefaults() *UserSessionUserAction`

NewUserSessionUserActionWithDefaults instantiates a new UserSessionUserAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserSessionUserAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSessionUserAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSessionUserAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserSessionUserAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomain

`func (o *UserSessionUserAction) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UserSessionUserAction) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UserSessionUserAction) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UserSessionUserAction) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTargetUrl

`func (o *UserSessionUserAction) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *UserSessionUserAction) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *UserSessionUserAction) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *UserSessionUserAction) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetType

`func (o *UserSessionUserAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSessionUserAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSessionUserAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserSessionUserAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStartTime

`func (o *UserSessionUserAction) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UserSessionUserAction) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UserSessionUserAction) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UserSessionUserAction) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *UserSessionUserAction) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UserSessionUserAction) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UserSessionUserAction) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UserSessionUserAction) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetDuration

`func (o *UserSessionUserAction) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UserSessionUserAction) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UserSessionUserAction) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *UserSessionUserAction) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetApplication

`func (o *UserSessionUserAction) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *UserSessionUserAction) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *UserSessionUserAction) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *UserSessionUserAction) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetInternalApplicationId

`func (o *UserSessionUserAction) GetInternalApplicationId() string`

GetInternalApplicationId returns the InternalApplicationId field if non-nil, zero value otherwise.

### GetInternalApplicationIdOk

`func (o *UserSessionUserAction) GetInternalApplicationIdOk() (*string, bool)`

GetInternalApplicationIdOk returns a tuple with the InternalApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalApplicationId

`func (o *UserSessionUserAction) SetInternalApplicationId(v string)`

SetInternalApplicationId sets InternalApplicationId field to given value.

### HasInternalApplicationId

`func (o *UserSessionUserAction) HasInternalApplicationId() bool`

HasInternalApplicationId returns a boolean if a field has been set.

### GetInternalKeyUserActionId

`func (o *UserSessionUserAction) GetInternalKeyUserActionId() string`

GetInternalKeyUserActionId returns the InternalKeyUserActionId field if non-nil, zero value otherwise.

### GetInternalKeyUserActionIdOk

`func (o *UserSessionUserAction) GetInternalKeyUserActionIdOk() (*string, bool)`

GetInternalKeyUserActionIdOk returns a tuple with the InternalKeyUserActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalKeyUserActionId

`func (o *UserSessionUserAction) SetInternalKeyUserActionId(v string)`

SetInternalKeyUserActionId sets InternalKeyUserActionId field to given value.

### HasInternalKeyUserActionId

`func (o *UserSessionUserAction) HasInternalKeyUserActionId() bool`

HasInternalKeyUserActionId returns a boolean if a field has been set.

### GetSpeedIndex

`func (o *UserSessionUserAction) GetSpeedIndex() int32`

GetSpeedIndex returns the SpeedIndex field if non-nil, zero value otherwise.

### GetSpeedIndexOk

`func (o *UserSessionUserAction) GetSpeedIndexOk() (*int32, bool)`

GetSpeedIndexOk returns a tuple with the SpeedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedIndex

`func (o *UserSessionUserAction) SetSpeedIndex(v int32)`

SetSpeedIndex sets SpeedIndex field to given value.

### HasSpeedIndex

`func (o *UserSessionUserAction) HasSpeedIndex() bool`

HasSpeedIndex returns a boolean if a field has been set.

### GetErrorCount

`func (o *UserSessionUserAction) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *UserSessionUserAction) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *UserSessionUserAction) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.

### HasErrorCount

`func (o *UserSessionUserAction) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### GetApdexCategory

`func (o *UserSessionUserAction) GetApdexCategory() string`

GetApdexCategory returns the ApdexCategory field if non-nil, zero value otherwise.

### GetApdexCategoryOk

`func (o *UserSessionUserAction) GetApdexCategoryOk() (*string, bool)`

GetApdexCategoryOk returns a tuple with the ApdexCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApdexCategory

`func (o *UserSessionUserAction) SetApdexCategory(v string)`

SetApdexCategory sets ApdexCategory field to given value.

### HasApdexCategory

`func (o *UserSessionUserAction) HasApdexCategory() bool`

HasApdexCategory returns a boolean if a field has been set.

### GetMatchingConversionGoals

`func (o *UserSessionUserAction) GetMatchingConversionGoals() []string`

GetMatchingConversionGoals returns the MatchingConversionGoals field if non-nil, zero value otherwise.

### GetMatchingConversionGoalsOk

`func (o *UserSessionUserAction) GetMatchingConversionGoalsOk() (*[]string, bool)`

GetMatchingConversionGoalsOk returns a tuple with the MatchingConversionGoals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingConversionGoals

`func (o *UserSessionUserAction) SetMatchingConversionGoals(v []string)`

SetMatchingConversionGoals sets MatchingConversionGoals field to given value.

### HasMatchingConversionGoals

`func (o *UserSessionUserAction) HasMatchingConversionGoals() bool`

HasMatchingConversionGoals returns a boolean if a field has been set.

### GetNetworkTime

`func (o *UserSessionUserAction) GetNetworkTime() int32`

GetNetworkTime returns the NetworkTime field if non-nil, zero value otherwise.

### GetNetworkTimeOk

`func (o *UserSessionUserAction) GetNetworkTimeOk() (*int32, bool)`

GetNetworkTimeOk returns a tuple with the NetworkTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTime

`func (o *UserSessionUserAction) SetNetworkTime(v int32)`

SetNetworkTime sets NetworkTime field to given value.

### HasNetworkTime

`func (o *UserSessionUserAction) HasNetworkTime() bool`

HasNetworkTime returns a boolean if a field has been set.

### GetServerTime

`func (o *UserSessionUserAction) GetServerTime() int32`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *UserSessionUserAction) GetServerTimeOk() (*int32, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *UserSessionUserAction) SetServerTime(v int32)`

SetServerTime sets ServerTime field to given value.

### HasServerTime

`func (o *UserSessionUserAction) HasServerTime() bool`

HasServerTime returns a boolean if a field has been set.

### GetFrontendTime

`func (o *UserSessionUserAction) GetFrontendTime() int32`

GetFrontendTime returns the FrontendTime field if non-nil, zero value otherwise.

### GetFrontendTimeOk

`func (o *UserSessionUserAction) GetFrontendTimeOk() (*int32, bool)`

GetFrontendTimeOk returns a tuple with the FrontendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontendTime

`func (o *UserSessionUserAction) SetFrontendTime(v int32)`

SetFrontendTime sets FrontendTime field to given value.

### HasFrontendTime

`func (o *UserSessionUserAction) HasFrontendTime() bool`

HasFrontendTime returns a boolean if a field has been set.

### GetDocumentInteractiveTime

`func (o *UserSessionUserAction) GetDocumentInteractiveTime() int32`

GetDocumentInteractiveTime returns the DocumentInteractiveTime field if non-nil, zero value otherwise.

### GetDocumentInteractiveTimeOk

`func (o *UserSessionUserAction) GetDocumentInteractiveTimeOk() (*int32, bool)`

GetDocumentInteractiveTimeOk returns a tuple with the DocumentInteractiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentInteractiveTime

`func (o *UserSessionUserAction) SetDocumentInteractiveTime(v int32)`

SetDocumentInteractiveTime sets DocumentInteractiveTime field to given value.

### HasDocumentInteractiveTime

`func (o *UserSessionUserAction) HasDocumentInteractiveTime() bool`

HasDocumentInteractiveTime returns a boolean if a field has been set.

### GetFailedImages

`func (o *UserSessionUserAction) GetFailedImages() int32`

GetFailedImages returns the FailedImages field if non-nil, zero value otherwise.

### GetFailedImagesOk

`func (o *UserSessionUserAction) GetFailedImagesOk() (*int32, bool)`

GetFailedImagesOk returns a tuple with the FailedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedImages

`func (o *UserSessionUserAction) SetFailedImages(v int32)`

SetFailedImages sets FailedImages field to given value.

### HasFailedImages

`func (o *UserSessionUserAction) HasFailedImages() bool`

HasFailedImages returns a boolean if a field has been set.

### GetFailedXhrRequests

`func (o *UserSessionUserAction) GetFailedXhrRequests() int32`

GetFailedXhrRequests returns the FailedXhrRequests field if non-nil, zero value otherwise.

### GetFailedXhrRequestsOk

`func (o *UserSessionUserAction) GetFailedXhrRequestsOk() (*int32, bool)`

GetFailedXhrRequestsOk returns a tuple with the FailedXhrRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedXhrRequests

`func (o *UserSessionUserAction) SetFailedXhrRequests(v int32)`

SetFailedXhrRequests sets FailedXhrRequests field to given value.

### HasFailedXhrRequests

`func (o *UserSessionUserAction) HasFailedXhrRequests() bool`

HasFailedXhrRequests returns a boolean if a field has been set.

### GetHttpRequestsWithErrors

`func (o *UserSessionUserAction) GetHttpRequestsWithErrors() int32`

GetHttpRequestsWithErrors returns the HttpRequestsWithErrors field if non-nil, zero value otherwise.

### GetHttpRequestsWithErrorsOk

`func (o *UserSessionUserAction) GetHttpRequestsWithErrorsOk() (*int32, bool)`

GetHttpRequestsWithErrorsOk returns a tuple with the HttpRequestsWithErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestsWithErrors

`func (o *UserSessionUserAction) SetHttpRequestsWithErrors(v int32)`

SetHttpRequestsWithErrors sets HttpRequestsWithErrors field to given value.

### HasHttpRequestsWithErrors

`func (o *UserSessionUserAction) HasHttpRequestsWithErrors() bool`

HasHttpRequestsWithErrors returns a boolean if a field has been set.

### GetThirdPartyResources

`func (o *UserSessionUserAction) GetThirdPartyResources() int32`

GetThirdPartyResources returns the ThirdPartyResources field if non-nil, zero value otherwise.

### GetThirdPartyResourcesOk

`func (o *UserSessionUserAction) GetThirdPartyResourcesOk() (*int32, bool)`

GetThirdPartyResourcesOk returns a tuple with the ThirdPartyResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyResources

`func (o *UserSessionUserAction) SetThirdPartyResources(v int32)`

SetThirdPartyResources sets ThirdPartyResources field to given value.

### HasThirdPartyResources

`func (o *UserSessionUserAction) HasThirdPartyResources() bool`

HasThirdPartyResources returns a boolean if a field has been set.

### GetThirdPartyBusyTime

`func (o *UserSessionUserAction) GetThirdPartyBusyTime() int32`

GetThirdPartyBusyTime returns the ThirdPartyBusyTime field if non-nil, zero value otherwise.

### GetThirdPartyBusyTimeOk

`func (o *UserSessionUserAction) GetThirdPartyBusyTimeOk() (*int32, bool)`

GetThirdPartyBusyTimeOk returns a tuple with the ThirdPartyBusyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyBusyTime

`func (o *UserSessionUserAction) SetThirdPartyBusyTime(v int32)`

SetThirdPartyBusyTime sets ThirdPartyBusyTime field to given value.

### HasThirdPartyBusyTime

`func (o *UserSessionUserAction) HasThirdPartyBusyTime() bool`

HasThirdPartyBusyTime returns a boolean if a field has been set.

### GetCdnResources

`func (o *UserSessionUserAction) GetCdnResources() int32`

GetCdnResources returns the CdnResources field if non-nil, zero value otherwise.

### GetCdnResourcesOk

`func (o *UserSessionUserAction) GetCdnResourcesOk() (*int32, bool)`

GetCdnResourcesOk returns a tuple with the CdnResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnResources

`func (o *UserSessionUserAction) SetCdnResources(v int32)`

SetCdnResources sets CdnResources field to given value.

### HasCdnResources

`func (o *UserSessionUserAction) HasCdnResources() bool`

HasCdnResources returns a boolean if a field has been set.

### GetCdnBusyTime

`func (o *UserSessionUserAction) GetCdnBusyTime() int32`

GetCdnBusyTime returns the CdnBusyTime field if non-nil, zero value otherwise.

### GetCdnBusyTimeOk

`func (o *UserSessionUserAction) GetCdnBusyTimeOk() (*int32, bool)`

GetCdnBusyTimeOk returns a tuple with the CdnBusyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnBusyTime

`func (o *UserSessionUserAction) SetCdnBusyTime(v int32)`

SetCdnBusyTime sets CdnBusyTime field to given value.

### HasCdnBusyTime

`func (o *UserSessionUserAction) HasCdnBusyTime() bool`

HasCdnBusyTime returns a boolean if a field has been set.

### GetFirstPartyResources

`func (o *UserSessionUserAction) GetFirstPartyResources() int32`

GetFirstPartyResources returns the FirstPartyResources field if non-nil, zero value otherwise.

### GetFirstPartyResourcesOk

`func (o *UserSessionUserAction) GetFirstPartyResourcesOk() (*int32, bool)`

GetFirstPartyResourcesOk returns a tuple with the FirstPartyResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPartyResources

`func (o *UserSessionUserAction) SetFirstPartyResources(v int32)`

SetFirstPartyResources sets FirstPartyResources field to given value.

### HasFirstPartyResources

`func (o *UserSessionUserAction) HasFirstPartyResources() bool`

HasFirstPartyResources returns a boolean if a field has been set.

### GetFirstPartyBusyTime

`func (o *UserSessionUserAction) GetFirstPartyBusyTime() int32`

GetFirstPartyBusyTime returns the FirstPartyBusyTime field if non-nil, zero value otherwise.

### GetFirstPartyBusyTimeOk

`func (o *UserSessionUserAction) GetFirstPartyBusyTimeOk() (*int32, bool)`

GetFirstPartyBusyTimeOk returns a tuple with the FirstPartyBusyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPartyBusyTime

`func (o *UserSessionUserAction) SetFirstPartyBusyTime(v int32)`

SetFirstPartyBusyTime sets FirstPartyBusyTime field to given value.

### HasFirstPartyBusyTime

`func (o *UserSessionUserAction) HasFirstPartyBusyTime() bool`

HasFirstPartyBusyTime returns a boolean if a field has been set.

### GetHasCrash

`func (o *UserSessionUserAction) GetHasCrash() bool`

GetHasCrash returns the HasCrash field if non-nil, zero value otherwise.

### GetHasCrashOk

`func (o *UserSessionUserAction) GetHasCrashOk() (*bool, bool)`

GetHasCrashOk returns a tuple with the HasCrash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCrash

`func (o *UserSessionUserAction) SetHasCrash(v bool)`

SetHasCrash sets HasCrash field to given value.

### HasHasCrash

`func (o *UserSessionUserAction) HasHasCrash() bool`

HasHasCrash returns a boolean if a field has been set.

### GetDomCompleteTime

`func (o *UserSessionUserAction) GetDomCompleteTime() int32`

GetDomCompleteTime returns the DomCompleteTime field if non-nil, zero value otherwise.

### GetDomCompleteTimeOk

`func (o *UserSessionUserAction) GetDomCompleteTimeOk() (*int32, bool)`

GetDomCompleteTimeOk returns a tuple with the DomCompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomCompleteTime

`func (o *UserSessionUserAction) SetDomCompleteTime(v int32)`

SetDomCompleteTime sets DomCompleteTime field to given value.

### HasDomCompleteTime

`func (o *UserSessionUserAction) HasDomCompleteTime() bool`

HasDomCompleteTime returns a boolean if a field has been set.

### GetDomContentLoadedTime

`func (o *UserSessionUserAction) GetDomContentLoadedTime() int32`

GetDomContentLoadedTime returns the DomContentLoadedTime field if non-nil, zero value otherwise.

### GetDomContentLoadedTimeOk

`func (o *UserSessionUserAction) GetDomContentLoadedTimeOk() (*int32, bool)`

GetDomContentLoadedTimeOk returns a tuple with the DomContentLoadedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomContentLoadedTime

`func (o *UserSessionUserAction) SetDomContentLoadedTime(v int32)`

SetDomContentLoadedTime sets DomContentLoadedTime field to given value.

### HasDomContentLoadedTime

`func (o *UserSessionUserAction) HasDomContentLoadedTime() bool`

HasDomContentLoadedTime returns a boolean if a field has been set.

### GetLoadEventStart

`func (o *UserSessionUserAction) GetLoadEventStart() int32`

GetLoadEventStart returns the LoadEventStart field if non-nil, zero value otherwise.

### GetLoadEventStartOk

`func (o *UserSessionUserAction) GetLoadEventStartOk() (*int32, bool)`

GetLoadEventStartOk returns a tuple with the LoadEventStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadEventStart

`func (o *UserSessionUserAction) SetLoadEventStart(v int32)`

SetLoadEventStart sets LoadEventStart field to given value.

### HasLoadEventStart

`func (o *UserSessionUserAction) HasLoadEventStart() bool`

HasLoadEventStart returns a boolean if a field has been set.

### GetLoadEventEnd

`func (o *UserSessionUserAction) GetLoadEventEnd() int32`

GetLoadEventEnd returns the LoadEventEnd field if non-nil, zero value otherwise.

### GetLoadEventEndOk

`func (o *UserSessionUserAction) GetLoadEventEndOk() (*int32, bool)`

GetLoadEventEndOk returns a tuple with the LoadEventEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadEventEnd

`func (o *UserSessionUserAction) SetLoadEventEnd(v int32)`

SetLoadEventEnd sets LoadEventEnd field to given value.

### HasLoadEventEnd

`func (o *UserSessionUserAction) HasLoadEventEnd() bool`

HasLoadEventEnd returns a boolean if a field has been set.

### GetNavigationStart

`func (o *UserSessionUserAction) GetNavigationStart() int64`

GetNavigationStart returns the NavigationStart field if non-nil, zero value otherwise.

### GetNavigationStartOk

`func (o *UserSessionUserAction) GetNavigationStartOk() (*int64, bool)`

GetNavigationStartOk returns a tuple with the NavigationStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigationStart

`func (o *UserSessionUserAction) SetNavigationStart(v int64)`

SetNavigationStart sets NavigationStart field to given value.

### HasNavigationStart

`func (o *UserSessionUserAction) HasNavigationStart() bool`

HasNavigationStart returns a boolean if a field has been set.

### GetRequestStart

`func (o *UserSessionUserAction) GetRequestStart() int32`

GetRequestStart returns the RequestStart field if non-nil, zero value otherwise.

### GetRequestStartOk

`func (o *UserSessionUserAction) GetRequestStartOk() (*int32, bool)`

GetRequestStartOk returns a tuple with the RequestStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStart

`func (o *UserSessionUserAction) SetRequestStart(v int32)`

SetRequestStart sets RequestStart field to given value.

### HasRequestStart

`func (o *UserSessionUserAction) HasRequestStart() bool`

HasRequestStart returns a boolean if a field has been set.

### GetResponseStart

`func (o *UserSessionUserAction) GetResponseStart() int32`

GetResponseStart returns the ResponseStart field if non-nil, zero value otherwise.

### GetResponseStartOk

`func (o *UserSessionUserAction) GetResponseStartOk() (*int32, bool)`

GetResponseStartOk returns a tuple with the ResponseStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStart

`func (o *UserSessionUserAction) SetResponseStart(v int32)`

SetResponseStart sets ResponseStart field to given value.

### HasResponseStart

`func (o *UserSessionUserAction) HasResponseStart() bool`

HasResponseStart returns a boolean if a field has been set.

### GetResponseEnd

`func (o *UserSessionUserAction) GetResponseEnd() int32`

GetResponseEnd returns the ResponseEnd field if non-nil, zero value otherwise.

### GetResponseEndOk

`func (o *UserSessionUserAction) GetResponseEndOk() (*int32, bool)`

GetResponseEndOk returns a tuple with the ResponseEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseEnd

`func (o *UserSessionUserAction) SetResponseEnd(v int32)`

SetResponseEnd sets ResponseEnd field to given value.

### HasResponseEnd

`func (o *UserSessionUserAction) HasResponseEnd() bool`

HasResponseEnd returns a boolean if a field has been set.

### GetVisuallyCompleteTime

`func (o *UserSessionUserAction) GetVisuallyCompleteTime() int32`

GetVisuallyCompleteTime returns the VisuallyCompleteTime field if non-nil, zero value otherwise.

### GetVisuallyCompleteTimeOk

`func (o *UserSessionUserAction) GetVisuallyCompleteTimeOk() (*int32, bool)`

GetVisuallyCompleteTimeOk returns a tuple with the VisuallyCompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisuallyCompleteTime

`func (o *UserSessionUserAction) SetVisuallyCompleteTime(v int32)`

SetVisuallyCompleteTime sets VisuallyCompleteTime field to given value.

### HasVisuallyCompleteTime

`func (o *UserSessionUserAction) HasVisuallyCompleteTime() bool`

HasVisuallyCompleteTime returns a boolean if a field has been set.

### GetSyntheticEvent

`func (o *UserSessionUserAction) GetSyntheticEvent() string`

GetSyntheticEvent returns the SyntheticEvent field if non-nil, zero value otherwise.

### GetSyntheticEventOk

`func (o *UserSessionUserAction) GetSyntheticEventOk() (*string, bool)`

GetSyntheticEventOk returns a tuple with the SyntheticEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticEvent

`func (o *UserSessionUserAction) SetSyntheticEvent(v string)`

SetSyntheticEvent sets SyntheticEvent field to given value.

### HasSyntheticEvent

`func (o *UserSessionUserAction) HasSyntheticEvent() bool`

HasSyntheticEvent returns a boolean if a field has been set.

### GetSyntheticEventId

`func (o *UserSessionUserAction) GetSyntheticEventId() string`

GetSyntheticEventId returns the SyntheticEventId field if non-nil, zero value otherwise.

### GetSyntheticEventIdOk

`func (o *UserSessionUserAction) GetSyntheticEventIdOk() (*string, bool)`

GetSyntheticEventIdOk returns a tuple with the SyntheticEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticEventId

`func (o *UserSessionUserAction) SetSyntheticEventId(v string)`

SetSyntheticEventId sets SyntheticEventId field to given value.

### HasSyntheticEventId

`func (o *UserSessionUserAction) HasSyntheticEventId() bool`

HasSyntheticEventId returns a boolean if a field has been set.

### GetKeyUserAction

`func (o *UserSessionUserAction) GetKeyUserAction() bool`

GetKeyUserAction returns the KeyUserAction field if non-nil, zero value otherwise.

### GetKeyUserActionOk

`func (o *UserSessionUserAction) GetKeyUserActionOk() (*bool, bool)`

GetKeyUserActionOk returns a tuple with the KeyUserAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUserAction

`func (o *UserSessionUserAction) SetKeyUserAction(v bool)`

SetKeyUserAction sets KeyUserAction field to given value.

### HasKeyUserAction

`func (o *UserSessionUserAction) HasKeyUserAction() bool`

HasKeyUserAction returns a boolean if a field has been set.

### GetStringProperties

`func (o *UserSessionUserAction) GetStringProperties() []StringProperty`

GetStringProperties returns the StringProperties field if non-nil, zero value otherwise.

### GetStringPropertiesOk

`func (o *UserSessionUserAction) GetStringPropertiesOk() (*[]StringProperty, bool)`

GetStringPropertiesOk returns a tuple with the StringProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringProperties

`func (o *UserSessionUserAction) SetStringProperties(v []StringProperty)`

SetStringProperties sets StringProperties field to given value.

### HasStringProperties

`func (o *UserSessionUserAction) HasStringProperties() bool`

HasStringProperties returns a boolean if a field has been set.

### GetLongProperties

`func (o *UserSessionUserAction) GetLongProperties() []LongProperty`

GetLongProperties returns the LongProperties field if non-nil, zero value otherwise.

### GetLongPropertiesOk

`func (o *UserSessionUserAction) GetLongPropertiesOk() (*[]LongProperty, bool)`

GetLongPropertiesOk returns a tuple with the LongProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongProperties

`func (o *UserSessionUserAction) SetLongProperties(v []LongProperty)`

SetLongProperties sets LongProperties field to given value.

### HasLongProperties

`func (o *UserSessionUserAction) HasLongProperties() bool`

HasLongProperties returns a boolean if a field has been set.

### GetDoubleProperties

`func (o *UserSessionUserAction) GetDoubleProperties() []DoubleProperty`

GetDoubleProperties returns the DoubleProperties field if non-nil, zero value otherwise.

### GetDoublePropertiesOk

`func (o *UserSessionUserAction) GetDoublePropertiesOk() (*[]DoubleProperty, bool)`

GetDoublePropertiesOk returns a tuple with the DoubleProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleProperties

`func (o *UserSessionUserAction) SetDoubleProperties(v []DoubleProperty)`

SetDoubleProperties sets DoubleProperties field to given value.

### HasDoubleProperties

`func (o *UserSessionUserAction) HasDoubleProperties() bool`

HasDoubleProperties returns a boolean if a field has been set.

### GetDateProperties

`func (o *UserSessionUserAction) GetDateProperties() []DateProperty`

GetDateProperties returns the DateProperties field if non-nil, zero value otherwise.

### GetDatePropertiesOk

`func (o *UserSessionUserAction) GetDatePropertiesOk() (*[]DateProperty, bool)`

GetDatePropertiesOk returns a tuple with the DateProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateProperties

`func (o *UserSessionUserAction) SetDateProperties(v []DateProperty)`

SetDateProperties sets DateProperties field to given value.

### HasDateProperties

`func (o *UserSessionUserAction) HasDateProperties() bool`

HasDateProperties returns a boolean if a field has been set.

### GetUserActionPropertyCount

`func (o *UserSessionUserAction) GetUserActionPropertyCount() int32`

GetUserActionPropertyCount returns the UserActionPropertyCount field if non-nil, zero value otherwise.

### GetUserActionPropertyCountOk

`func (o *UserSessionUserAction) GetUserActionPropertyCountOk() (*int32, bool)`

GetUserActionPropertyCountOk returns a tuple with the UserActionPropertyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionPropertyCount

`func (o *UserSessionUserAction) SetUserActionPropertyCount(v int32)`

SetUserActionPropertyCount sets UserActionPropertyCount field to given value.

### HasUserActionPropertyCount

`func (o *UserSessionUserAction) HasUserActionPropertyCount() bool`

HasUserActionPropertyCount returns a boolean if a field has been set.

### GetCustomErrorCount

`func (o *UserSessionUserAction) GetCustomErrorCount() int32`

GetCustomErrorCount returns the CustomErrorCount field if non-nil, zero value otherwise.

### GetCustomErrorCountOk

`func (o *UserSessionUserAction) GetCustomErrorCountOk() (*int32, bool)`

GetCustomErrorCountOk returns a tuple with the CustomErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomErrorCount

`func (o *UserSessionUserAction) SetCustomErrorCount(v int32)`

SetCustomErrorCount sets CustomErrorCount field to given value.

### HasCustomErrorCount

`func (o *UserSessionUserAction) HasCustomErrorCount() bool`

HasCustomErrorCount returns a boolean if a field has been set.

### GetJavascriptErrorCount

`func (o *UserSessionUserAction) GetJavascriptErrorCount() int32`

GetJavascriptErrorCount returns the JavascriptErrorCount field if non-nil, zero value otherwise.

### GetJavascriptErrorCountOk

`func (o *UserSessionUserAction) GetJavascriptErrorCountOk() (*int32, bool)`

GetJavascriptErrorCountOk returns a tuple with the JavascriptErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavascriptErrorCount

`func (o *UserSessionUserAction) SetJavascriptErrorCount(v int32)`

SetJavascriptErrorCount sets JavascriptErrorCount field to given value.

### HasJavascriptErrorCount

`func (o *UserSessionUserAction) HasJavascriptErrorCount() bool`

HasJavascriptErrorCount returns a boolean if a field has been set.

### GetRequestErrorCount

`func (o *UserSessionUserAction) GetRequestErrorCount() int32`

GetRequestErrorCount returns the RequestErrorCount field if non-nil, zero value otherwise.

### GetRequestErrorCountOk

`func (o *UserSessionUserAction) GetRequestErrorCountOk() (*int32, bool)`

GetRequestErrorCountOk returns a tuple with the RequestErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestErrorCount

`func (o *UserSessionUserAction) SetRequestErrorCount(v int32)`

SetRequestErrorCount sets RequestErrorCount field to given value.

### HasRequestErrorCount

`func (o *UserSessionUserAction) HasRequestErrorCount() bool`

HasRequestErrorCount returns a boolean if a field has been set.

### GetLargestContentfulPaint

`func (o *UserSessionUserAction) GetLargestContentfulPaint() int32`

GetLargestContentfulPaint returns the LargestContentfulPaint field if non-nil, zero value otherwise.

### GetLargestContentfulPaintOk

`func (o *UserSessionUserAction) GetLargestContentfulPaintOk() (*int32, bool)`

GetLargestContentfulPaintOk returns a tuple with the LargestContentfulPaint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargestContentfulPaint

`func (o *UserSessionUserAction) SetLargestContentfulPaint(v int32)`

SetLargestContentfulPaint sets LargestContentfulPaint field to given value.

### HasLargestContentfulPaint

`func (o *UserSessionUserAction) HasLargestContentfulPaint() bool`

HasLargestContentfulPaint returns a boolean if a field has been set.

### GetFirstInputDelay

`func (o *UserSessionUserAction) GetFirstInputDelay() int32`

GetFirstInputDelay returns the FirstInputDelay field if non-nil, zero value otherwise.

### GetFirstInputDelayOk

`func (o *UserSessionUserAction) GetFirstInputDelayOk() (*int32, bool)`

GetFirstInputDelayOk returns a tuple with the FirstInputDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstInputDelay

`func (o *UserSessionUserAction) SetFirstInputDelay(v int32)`

SetFirstInputDelay sets FirstInputDelay field to given value.

### HasFirstInputDelay

`func (o *UserSessionUserAction) HasFirstInputDelay() bool`

HasFirstInputDelay returns a boolean if a field has been set.

### GetTotalBlockingTime

`func (o *UserSessionUserAction) GetTotalBlockingTime() int32`

GetTotalBlockingTime returns the TotalBlockingTime field if non-nil, zero value otherwise.

### GetTotalBlockingTimeOk

`func (o *UserSessionUserAction) GetTotalBlockingTimeOk() (*int32, bool)`

GetTotalBlockingTimeOk returns a tuple with the TotalBlockingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBlockingTime

`func (o *UserSessionUserAction) SetTotalBlockingTime(v int32)`

SetTotalBlockingTime sets TotalBlockingTime field to given value.

### HasTotalBlockingTime

`func (o *UserSessionUserAction) HasTotalBlockingTime() bool`

HasTotalBlockingTime returns a boolean if a field has been set.

### GetCumulativeLayoutShift

`func (o *UserSessionUserAction) GetCumulativeLayoutShift() float64`

GetCumulativeLayoutShift returns the CumulativeLayoutShift field if non-nil, zero value otherwise.

### GetCumulativeLayoutShiftOk

`func (o *UserSessionUserAction) GetCumulativeLayoutShiftOk() (*float64, bool)`

GetCumulativeLayoutShiftOk returns a tuple with the CumulativeLayoutShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeLayoutShift

`func (o *UserSessionUserAction) SetCumulativeLayoutShift(v float64)`

SetCumulativeLayoutShift sets CumulativeLayoutShift field to given value.

### HasCumulativeLayoutShift

`func (o *UserSessionUserAction) HasCumulativeLayoutShift() bool`

HasCumulativeLayoutShift returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


