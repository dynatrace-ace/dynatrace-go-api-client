# CalculatedMobileMetricUserActionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasReportedError** | Pointer to **bool** | The error status of the actions to be included in the metric calculation:   * &#x60;true&#x60;: Only actions with reported errors are included.    * &#x60;false&#x60;: All actions are included. | [optional] 
**HasHttpError** | Pointer to **bool** | The HTTP error status of the actions to be included in the metric calculation:   * &#x60;true&#x60;: Only actions with HTTP errors are included.    * &#x60;false&#x60;: All actions are included. | [optional] 
**UserActionName** | Pointer to **string** | Only actions with this name are included in the metric calculation.    The EQUALS operator applies. | [optional] 
**AppVersion** | Pointer to **string** | Only actions coming from this app version are included in the metric calculation.    The EQUALS operator applies. | [optional] 
**Device** | Pointer to **string** | Only actions coming from this app version are included in the metric calculation.    The EQUALS operator applies. | [optional] 
**Manufacturer** | Pointer to **string** | Only actions coming from devices of this manufacturer are included in the metric calculation.    The EQUALS operator applies. | [optional] 
**Apdex** | Pointer to **string** | Only actions with the specified Apdex score are included in the metric calculation. | [optional] 
**OsFamily** | Pointer to **string** | Only actions coming from this OS family are included in the metric calculation.    Specify the OS ID here. | [optional] 
**OsVersion** | Pointer to **string** | Only actions coming from this OS version are included in the metric calculation.    Specify the OS ID here. | [optional] 
**City** | Pointer to **string** | Only actions of users from this city are included in the metric calculation.    Specify geolocation ID here. | [optional] 
**Continent** | Pointer to **string** | Only actions of users from this continent are included in the metric calculation.    Specify geolocation ID here. | [optional] 
**Country** | Pointer to **string** | Only actions of users from this country are included in the metric calculation.    Specify geolocation ID here. | [optional] 
**Region** | Pointer to **string** | Only actions of users from this region are included in the metric calculation.    Specify geolocation ID here. | [optional] 
**ActionDurationFromMilliseconds** | Pointer to **int32** | Only actions with a duration more than or equal to this value (in milliseconds) are included in the metric calculation. | [optional] 
**ActionDurationToMilliseconds** | Pointer to **int32** | Only actions with a duration less than or equal to this value (in milliseconds) are included in the metric calculation. | [optional] 
**Carrier** | Pointer to **string** | Only actions coming from this carrier type are included in the metric calculation. | [optional] 
**ConnectionType** | Pointer to **string** | Only actions coming from this connection type are included in the metric calculation. | [optional] 
**NetworkTechnology** | Pointer to **string** | Filter by network technology | [optional] 
**Isp** | Pointer to **string** | Only actions coming from this internet service provider are included in the metric calculation.     The EQUALS operator applies. | [optional] 
**Orientation** | Pointer to **string** | Only actions coming from devices with this display orientation are included in the metric calculation. | [optional] 
**Resolution** | Pointer to **string** | Only actions coming from devices with this display resolution are included in the metric calculation. | [optional] 

## Methods

### NewCalculatedMobileMetricUserActionFilter

`func NewCalculatedMobileMetricUserActionFilter() *CalculatedMobileMetricUserActionFilter`

NewCalculatedMobileMetricUserActionFilter instantiates a new CalculatedMobileMetricUserActionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalculatedMobileMetricUserActionFilterWithDefaults

`func NewCalculatedMobileMetricUserActionFilterWithDefaults() *CalculatedMobileMetricUserActionFilter`

NewCalculatedMobileMetricUserActionFilterWithDefaults instantiates a new CalculatedMobileMetricUserActionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasReportedError

`func (o *CalculatedMobileMetricUserActionFilter) GetHasReportedError() bool`

GetHasReportedError returns the HasReportedError field if non-nil, zero value otherwise.

### GetHasReportedErrorOk

`func (o *CalculatedMobileMetricUserActionFilter) GetHasReportedErrorOk() (*bool, bool)`

GetHasReportedErrorOk returns a tuple with the HasReportedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReportedError

`func (o *CalculatedMobileMetricUserActionFilter) SetHasReportedError(v bool)`

SetHasReportedError sets HasReportedError field to given value.

### HasHasReportedError

`func (o *CalculatedMobileMetricUserActionFilter) HasHasReportedError() bool`

HasHasReportedError returns a boolean if a field has been set.

### GetHasHttpError

`func (o *CalculatedMobileMetricUserActionFilter) GetHasHttpError() bool`

GetHasHttpError returns the HasHttpError field if non-nil, zero value otherwise.

### GetHasHttpErrorOk

`func (o *CalculatedMobileMetricUserActionFilter) GetHasHttpErrorOk() (*bool, bool)`

GetHasHttpErrorOk returns a tuple with the HasHttpError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHttpError

`func (o *CalculatedMobileMetricUserActionFilter) SetHasHttpError(v bool)`

SetHasHttpError sets HasHttpError field to given value.

### HasHasHttpError

`func (o *CalculatedMobileMetricUserActionFilter) HasHasHttpError() bool`

HasHasHttpError returns a boolean if a field has been set.

### GetUserActionName

`func (o *CalculatedMobileMetricUserActionFilter) GetUserActionName() string`

GetUserActionName returns the UserActionName field if non-nil, zero value otherwise.

### GetUserActionNameOk

`func (o *CalculatedMobileMetricUserActionFilter) GetUserActionNameOk() (*string, bool)`

GetUserActionNameOk returns a tuple with the UserActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionName

`func (o *CalculatedMobileMetricUserActionFilter) SetUserActionName(v string)`

SetUserActionName sets UserActionName field to given value.

### HasUserActionName

`func (o *CalculatedMobileMetricUserActionFilter) HasUserActionName() bool`

HasUserActionName returns a boolean if a field has been set.

### GetAppVersion

`func (o *CalculatedMobileMetricUserActionFilter) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *CalculatedMobileMetricUserActionFilter) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *CalculatedMobileMetricUserActionFilter) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *CalculatedMobileMetricUserActionFilter) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetDevice

`func (o *CalculatedMobileMetricUserActionFilter) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CalculatedMobileMetricUserActionFilter) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CalculatedMobileMetricUserActionFilter) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *CalculatedMobileMetricUserActionFilter) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetManufacturer

`func (o *CalculatedMobileMetricUserActionFilter) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *CalculatedMobileMetricUserActionFilter) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *CalculatedMobileMetricUserActionFilter) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *CalculatedMobileMetricUserActionFilter) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetApdex

`func (o *CalculatedMobileMetricUserActionFilter) GetApdex() string`

GetApdex returns the Apdex field if non-nil, zero value otherwise.

### GetApdexOk

`func (o *CalculatedMobileMetricUserActionFilter) GetApdexOk() (*string, bool)`

GetApdexOk returns a tuple with the Apdex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApdex

`func (o *CalculatedMobileMetricUserActionFilter) SetApdex(v string)`

SetApdex sets Apdex field to given value.

### HasApdex

`func (o *CalculatedMobileMetricUserActionFilter) HasApdex() bool`

HasApdex returns a boolean if a field has been set.

### GetOsFamily

`func (o *CalculatedMobileMetricUserActionFilter) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *CalculatedMobileMetricUserActionFilter) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *CalculatedMobileMetricUserActionFilter) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *CalculatedMobileMetricUserActionFilter) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetOsVersion

`func (o *CalculatedMobileMetricUserActionFilter) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *CalculatedMobileMetricUserActionFilter) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *CalculatedMobileMetricUserActionFilter) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *CalculatedMobileMetricUserActionFilter) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetCity

`func (o *CalculatedMobileMetricUserActionFilter) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CalculatedMobileMetricUserActionFilter) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CalculatedMobileMetricUserActionFilter) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CalculatedMobileMetricUserActionFilter) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetContinent

`func (o *CalculatedMobileMetricUserActionFilter) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *CalculatedMobileMetricUserActionFilter) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *CalculatedMobileMetricUserActionFilter) SetContinent(v string)`

SetContinent sets Continent field to given value.

### HasContinent

`func (o *CalculatedMobileMetricUserActionFilter) HasContinent() bool`

HasContinent returns a boolean if a field has been set.

### GetCountry

`func (o *CalculatedMobileMetricUserActionFilter) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CalculatedMobileMetricUserActionFilter) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CalculatedMobileMetricUserActionFilter) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CalculatedMobileMetricUserActionFilter) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetRegion

`func (o *CalculatedMobileMetricUserActionFilter) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CalculatedMobileMetricUserActionFilter) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CalculatedMobileMetricUserActionFilter) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CalculatedMobileMetricUserActionFilter) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetActionDurationFromMilliseconds

`func (o *CalculatedMobileMetricUserActionFilter) GetActionDurationFromMilliseconds() int32`

GetActionDurationFromMilliseconds returns the ActionDurationFromMilliseconds field if non-nil, zero value otherwise.

### GetActionDurationFromMillisecondsOk

`func (o *CalculatedMobileMetricUserActionFilter) GetActionDurationFromMillisecondsOk() (*int32, bool)`

GetActionDurationFromMillisecondsOk returns a tuple with the ActionDurationFromMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDurationFromMilliseconds

`func (o *CalculatedMobileMetricUserActionFilter) SetActionDurationFromMilliseconds(v int32)`

SetActionDurationFromMilliseconds sets ActionDurationFromMilliseconds field to given value.

### HasActionDurationFromMilliseconds

`func (o *CalculatedMobileMetricUserActionFilter) HasActionDurationFromMilliseconds() bool`

HasActionDurationFromMilliseconds returns a boolean if a field has been set.

### GetActionDurationToMilliseconds

`func (o *CalculatedMobileMetricUserActionFilter) GetActionDurationToMilliseconds() int32`

GetActionDurationToMilliseconds returns the ActionDurationToMilliseconds field if non-nil, zero value otherwise.

### GetActionDurationToMillisecondsOk

`func (o *CalculatedMobileMetricUserActionFilter) GetActionDurationToMillisecondsOk() (*int32, bool)`

GetActionDurationToMillisecondsOk returns a tuple with the ActionDurationToMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDurationToMilliseconds

`func (o *CalculatedMobileMetricUserActionFilter) SetActionDurationToMilliseconds(v int32)`

SetActionDurationToMilliseconds sets ActionDurationToMilliseconds field to given value.

### HasActionDurationToMilliseconds

`func (o *CalculatedMobileMetricUserActionFilter) HasActionDurationToMilliseconds() bool`

HasActionDurationToMilliseconds returns a boolean if a field has been set.

### GetCarrier

`func (o *CalculatedMobileMetricUserActionFilter) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *CalculatedMobileMetricUserActionFilter) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *CalculatedMobileMetricUserActionFilter) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *CalculatedMobileMetricUserActionFilter) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetConnectionType

`func (o *CalculatedMobileMetricUserActionFilter) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CalculatedMobileMetricUserActionFilter) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CalculatedMobileMetricUserActionFilter) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *CalculatedMobileMetricUserActionFilter) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetNetworkTechnology

`func (o *CalculatedMobileMetricUserActionFilter) GetNetworkTechnology() string`

GetNetworkTechnology returns the NetworkTechnology field if non-nil, zero value otherwise.

### GetNetworkTechnologyOk

`func (o *CalculatedMobileMetricUserActionFilter) GetNetworkTechnologyOk() (*string, bool)`

GetNetworkTechnologyOk returns a tuple with the NetworkTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTechnology

`func (o *CalculatedMobileMetricUserActionFilter) SetNetworkTechnology(v string)`

SetNetworkTechnology sets NetworkTechnology field to given value.

### HasNetworkTechnology

`func (o *CalculatedMobileMetricUserActionFilter) HasNetworkTechnology() bool`

HasNetworkTechnology returns a boolean if a field has been set.

### GetIsp

`func (o *CalculatedMobileMetricUserActionFilter) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *CalculatedMobileMetricUserActionFilter) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *CalculatedMobileMetricUserActionFilter) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *CalculatedMobileMetricUserActionFilter) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetOrientation

`func (o *CalculatedMobileMetricUserActionFilter) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *CalculatedMobileMetricUserActionFilter) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *CalculatedMobileMetricUserActionFilter) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *CalculatedMobileMetricUserActionFilter) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetResolution

`func (o *CalculatedMobileMetricUserActionFilter) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *CalculatedMobileMetricUserActionFilter) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *CalculatedMobileMetricUserActionFilter) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *CalculatedMobileMetricUserActionFilter) HasResolution() bool`

HasResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


