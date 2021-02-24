# MobileCustomAppConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The Dynatrace entity ID of the application. | [optional] [readonly] 
**Name** | **string** | The name of the application. | 
**ApplicationType** | Pointer to **string** | The type of the application. | [optional] [readonly] 
**ApplicationId** | Pointer to **string** | The UUID of the application.    It is used only by OneAgent to send data to Dynatrace. | [optional] [readonly] 
**IconType** | Pointer to **string** | The icon of the custom application.   Mobile applications always use the mobile phone icon&#x60;. | [optional] 
**CostControlUserSessionPercentage** | Pointer to **int32** | The percentage of user sessions to be analyzed. | [optional] 
**ApdexSettings** | Pointer to [**MobileCustomApdex**](MobileCustomApdex.md) |  | [optional] 
**OptInModeEnabled** | Pointer to **bool** | The opt-in mode is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 
**SessionReplayOnCrashEnabled** | Pointer to **bool** | The session replay on crash is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 
**BeaconEndpointType** | Pointer to **string** | The type of the beacon endpoint. | [optional] 
**BeaconEndpointUrl** | Pointer to **string** | The URL of the beacon endpoint.    Only applicable when the **beaconEndpointType** is set to &#x60;ENVIRONMENT_ACTIVE_GATE&#x60; or &#x60;INSTRUMENTED_WEB_SERVER&#x60;. | [optional] 

## Methods

### NewMobileCustomAppConfig

`func NewMobileCustomAppConfig(name string, ) *MobileCustomAppConfig`

NewMobileCustomAppConfig instantiates a new MobileCustomAppConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileCustomAppConfigWithDefaults

`func NewMobileCustomAppConfigWithDefaults() *MobileCustomAppConfig`

NewMobileCustomAppConfigWithDefaults instantiates a new MobileCustomAppConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *MobileCustomAppConfig) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MobileCustomAppConfig) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MobileCustomAppConfig) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *MobileCustomAppConfig) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *MobileCustomAppConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobileCustomAppConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobileCustomAppConfig) SetName(v string)`

SetName sets Name field to given value.


### GetApplicationType

`func (o *MobileCustomAppConfig) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *MobileCustomAppConfig) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *MobileCustomAppConfig) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *MobileCustomAppConfig) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetApplicationId

`func (o *MobileCustomAppConfig) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *MobileCustomAppConfig) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *MobileCustomAppConfig) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *MobileCustomAppConfig) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetIconType

`func (o *MobileCustomAppConfig) GetIconType() string`

GetIconType returns the IconType field if non-nil, zero value otherwise.

### GetIconTypeOk

`func (o *MobileCustomAppConfig) GetIconTypeOk() (*string, bool)`

GetIconTypeOk returns a tuple with the IconType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconType

`func (o *MobileCustomAppConfig) SetIconType(v string)`

SetIconType sets IconType field to given value.

### HasIconType

`func (o *MobileCustomAppConfig) HasIconType() bool`

HasIconType returns a boolean if a field has been set.

### GetCostControlUserSessionPercentage

`func (o *MobileCustomAppConfig) GetCostControlUserSessionPercentage() int32`

GetCostControlUserSessionPercentage returns the CostControlUserSessionPercentage field if non-nil, zero value otherwise.

### GetCostControlUserSessionPercentageOk

`func (o *MobileCustomAppConfig) GetCostControlUserSessionPercentageOk() (*int32, bool)`

GetCostControlUserSessionPercentageOk returns a tuple with the CostControlUserSessionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostControlUserSessionPercentage

`func (o *MobileCustomAppConfig) SetCostControlUserSessionPercentage(v int32)`

SetCostControlUserSessionPercentage sets CostControlUserSessionPercentage field to given value.

### HasCostControlUserSessionPercentage

`func (o *MobileCustomAppConfig) HasCostControlUserSessionPercentage() bool`

HasCostControlUserSessionPercentage returns a boolean if a field has been set.

### GetApdexSettings

`func (o *MobileCustomAppConfig) GetApdexSettings() MobileCustomApdex`

GetApdexSettings returns the ApdexSettings field if non-nil, zero value otherwise.

### GetApdexSettingsOk

`func (o *MobileCustomAppConfig) GetApdexSettingsOk() (*MobileCustomApdex, bool)`

GetApdexSettingsOk returns a tuple with the ApdexSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApdexSettings

`func (o *MobileCustomAppConfig) SetApdexSettings(v MobileCustomApdex)`

SetApdexSettings sets ApdexSettings field to given value.

### HasApdexSettings

`func (o *MobileCustomAppConfig) HasApdexSettings() bool`

HasApdexSettings returns a boolean if a field has been set.

### GetOptInModeEnabled

`func (o *MobileCustomAppConfig) GetOptInModeEnabled() bool`

GetOptInModeEnabled returns the OptInModeEnabled field if non-nil, zero value otherwise.

### GetOptInModeEnabledOk

`func (o *MobileCustomAppConfig) GetOptInModeEnabledOk() (*bool, bool)`

GetOptInModeEnabledOk returns a tuple with the OptInModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptInModeEnabled

`func (o *MobileCustomAppConfig) SetOptInModeEnabled(v bool)`

SetOptInModeEnabled sets OptInModeEnabled field to given value.

### HasOptInModeEnabled

`func (o *MobileCustomAppConfig) HasOptInModeEnabled() bool`

HasOptInModeEnabled returns a boolean if a field has been set.

### GetSessionReplayOnCrashEnabled

`func (o *MobileCustomAppConfig) GetSessionReplayOnCrashEnabled() bool`

GetSessionReplayOnCrashEnabled returns the SessionReplayOnCrashEnabled field if non-nil, zero value otherwise.

### GetSessionReplayOnCrashEnabledOk

`func (o *MobileCustomAppConfig) GetSessionReplayOnCrashEnabledOk() (*bool, bool)`

GetSessionReplayOnCrashEnabledOk returns a tuple with the SessionReplayOnCrashEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionReplayOnCrashEnabled

`func (o *MobileCustomAppConfig) SetSessionReplayOnCrashEnabled(v bool)`

SetSessionReplayOnCrashEnabled sets SessionReplayOnCrashEnabled field to given value.

### HasSessionReplayOnCrashEnabled

`func (o *MobileCustomAppConfig) HasSessionReplayOnCrashEnabled() bool`

HasSessionReplayOnCrashEnabled returns a boolean if a field has been set.

### GetBeaconEndpointType

`func (o *MobileCustomAppConfig) GetBeaconEndpointType() string`

GetBeaconEndpointType returns the BeaconEndpointType field if non-nil, zero value otherwise.

### GetBeaconEndpointTypeOk

`func (o *MobileCustomAppConfig) GetBeaconEndpointTypeOk() (*string, bool)`

GetBeaconEndpointTypeOk returns a tuple with the BeaconEndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconEndpointType

`func (o *MobileCustomAppConfig) SetBeaconEndpointType(v string)`

SetBeaconEndpointType sets BeaconEndpointType field to given value.

### HasBeaconEndpointType

`func (o *MobileCustomAppConfig) HasBeaconEndpointType() bool`

HasBeaconEndpointType returns a boolean if a field has been set.

### GetBeaconEndpointUrl

`func (o *MobileCustomAppConfig) GetBeaconEndpointUrl() string`

GetBeaconEndpointUrl returns the BeaconEndpointUrl field if non-nil, zero value otherwise.

### GetBeaconEndpointUrlOk

`func (o *MobileCustomAppConfig) GetBeaconEndpointUrlOk() (*string, bool)`

GetBeaconEndpointUrlOk returns a tuple with the BeaconEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconEndpointUrl

`func (o *MobileCustomAppConfig) SetBeaconEndpointUrl(v string)`

SetBeaconEndpointUrl sets BeaconEndpointUrl field to given value.

### HasBeaconEndpointUrl

`func (o *MobileCustomAppConfig) HasBeaconEndpointUrl() bool`

HasBeaconEndpointUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


