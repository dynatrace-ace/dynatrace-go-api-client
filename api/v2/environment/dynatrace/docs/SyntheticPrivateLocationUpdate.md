# SyntheticPrivateLocationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | **[]string** | A list of synthetic nodes belonging to the location.    You can retrieve the list of available nodes with the [GET all nodes](https://dt-url.net/miy3rpl) call. | 
**Name** | **string** | The name of the location. | 
**CountryCode** | Pointer to **string** | The country code of the location.    Use the alpha-2 code of the [ISO 3166-2 standard](https://dt-url.net/iso3166-2), (for example, &#x60;AT&#x60; for Austria or &#x60;PL&#x60; for Poland). | [optional] 
**RegionCode** | Pointer to **string** | The region code of the location.    For the [USA](https://dt-url.net/iso3166us) or [Canada](https://dt-url.net/iso3166ca) use ISO 3166-2 state codes (without &#x60;US-&#x60; or &#x60;CA-&#x60; prefix), for example, &#x60;VA&#x60; for Virginia or &#x60;OR&#x60; for Oregon.    For the rest of the world use [FIPS 10-4 codes](https://dt-url.net/fipscodes). | [optional] 
**City** | Pointer to **string** | The city of the location. | [optional] 
**Latitude** | **float64** | The latitude of the location in &#x60;DDD.dddd&#x60; format. | 
**Longitude** | **float64** | The longitude of the location in &#x60;DDD.dddd&#x60; format. | 
**Status** | Pointer to **string** | The status of the location:   * &#x60;ENABLED&#x60;: The location is displayed as active in the UI. You can assign monitors to the location.  * &#x60;DISABLED&#x60;: The location is displayed as inactive in the UI. You can&#39;t assign monitors to the location. Monitors already assigned to the location will stay there and will be executed from the location.  * &#x60;HIDDEN&#x60;: The location is not displayed in the UI. You can&#39;t assign monitors to the location. You can only set location as &#x60;HIDDEN&#x60; when no monitor is assigned to it. | [optional] 
**AvailabilityLocationOutage** | Pointer to **bool** | The alerting of location outage is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 
**AvailabilityNodeOutage** | Pointer to **bool** | The alerting of node outage is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;).    If enabled, the outage of *any* node in the location triggers an alert. | [optional] 
**LocationNodeOutageDelayInMinutes** | Pointer to **int32** | Alert if the location or node outage lasts longer than *X* minutes.    Only applicable when **availabilityLocationOutage** or **availabilityNodeOutage** is set to &#x60;true&#x60;. | [optional] 
**AvailabilityNotificationsEnabled** | Pointer to **bool** | The notifications of location and node outage is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 

## Methods

### NewSyntheticPrivateLocationUpdate

`func NewSyntheticPrivateLocationUpdate(nodes []string, name string, latitude float64, longitude float64, ) *SyntheticPrivateLocationUpdate`

NewSyntheticPrivateLocationUpdate instantiates a new SyntheticPrivateLocationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticPrivateLocationUpdateWithDefaults

`func NewSyntheticPrivateLocationUpdateWithDefaults() *SyntheticPrivateLocationUpdate`

NewSyntheticPrivateLocationUpdateWithDefaults instantiates a new SyntheticPrivateLocationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *SyntheticPrivateLocationUpdate) GetNodes() []string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *SyntheticPrivateLocationUpdate) GetNodesOk() (*[]string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *SyntheticPrivateLocationUpdate) SetNodes(v []string)`

SetNodes sets Nodes field to given value.


### GetName

`func (o *SyntheticPrivateLocationUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticPrivateLocationUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticPrivateLocationUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetCountryCode

`func (o *SyntheticPrivateLocationUpdate) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SyntheticPrivateLocationUpdate) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SyntheticPrivateLocationUpdate) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SyntheticPrivateLocationUpdate) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetRegionCode

`func (o *SyntheticPrivateLocationUpdate) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *SyntheticPrivateLocationUpdate) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *SyntheticPrivateLocationUpdate) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *SyntheticPrivateLocationUpdate) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetCity

`func (o *SyntheticPrivateLocationUpdate) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SyntheticPrivateLocationUpdate) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SyntheticPrivateLocationUpdate) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SyntheticPrivateLocationUpdate) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetLatitude

`func (o *SyntheticPrivateLocationUpdate) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SyntheticPrivateLocationUpdate) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SyntheticPrivateLocationUpdate) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *SyntheticPrivateLocationUpdate) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SyntheticPrivateLocationUpdate) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SyntheticPrivateLocationUpdate) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetStatus

`func (o *SyntheticPrivateLocationUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticPrivateLocationUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticPrivateLocationUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyntheticPrivateLocationUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAvailabilityLocationOutage

`func (o *SyntheticPrivateLocationUpdate) GetAvailabilityLocationOutage() bool`

GetAvailabilityLocationOutage returns the AvailabilityLocationOutage field if non-nil, zero value otherwise.

### GetAvailabilityLocationOutageOk

`func (o *SyntheticPrivateLocationUpdate) GetAvailabilityLocationOutageOk() (*bool, bool)`

GetAvailabilityLocationOutageOk returns a tuple with the AvailabilityLocationOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityLocationOutage

`func (o *SyntheticPrivateLocationUpdate) SetAvailabilityLocationOutage(v bool)`

SetAvailabilityLocationOutage sets AvailabilityLocationOutage field to given value.

### HasAvailabilityLocationOutage

`func (o *SyntheticPrivateLocationUpdate) HasAvailabilityLocationOutage() bool`

HasAvailabilityLocationOutage returns a boolean if a field has been set.

### GetAvailabilityNodeOutage

`func (o *SyntheticPrivateLocationUpdate) GetAvailabilityNodeOutage() bool`

GetAvailabilityNodeOutage returns the AvailabilityNodeOutage field if non-nil, zero value otherwise.

### GetAvailabilityNodeOutageOk

`func (o *SyntheticPrivateLocationUpdate) GetAvailabilityNodeOutageOk() (*bool, bool)`

GetAvailabilityNodeOutageOk returns a tuple with the AvailabilityNodeOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityNodeOutage

`func (o *SyntheticPrivateLocationUpdate) SetAvailabilityNodeOutage(v bool)`

SetAvailabilityNodeOutage sets AvailabilityNodeOutage field to given value.

### HasAvailabilityNodeOutage

`func (o *SyntheticPrivateLocationUpdate) HasAvailabilityNodeOutage() bool`

HasAvailabilityNodeOutage returns a boolean if a field has been set.

### GetLocationNodeOutageDelayInMinutes

`func (o *SyntheticPrivateLocationUpdate) GetLocationNodeOutageDelayInMinutes() int32`

GetLocationNodeOutageDelayInMinutes returns the LocationNodeOutageDelayInMinutes field if non-nil, zero value otherwise.

### GetLocationNodeOutageDelayInMinutesOk

`func (o *SyntheticPrivateLocationUpdate) GetLocationNodeOutageDelayInMinutesOk() (*int32, bool)`

GetLocationNodeOutageDelayInMinutesOk returns a tuple with the LocationNodeOutageDelayInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationNodeOutageDelayInMinutes

`func (o *SyntheticPrivateLocationUpdate) SetLocationNodeOutageDelayInMinutes(v int32)`

SetLocationNodeOutageDelayInMinutes sets LocationNodeOutageDelayInMinutes field to given value.

### HasLocationNodeOutageDelayInMinutes

`func (o *SyntheticPrivateLocationUpdate) HasLocationNodeOutageDelayInMinutes() bool`

HasLocationNodeOutageDelayInMinutes returns a boolean if a field has been set.

### GetAvailabilityNotificationsEnabled

`func (o *SyntheticPrivateLocationUpdate) GetAvailabilityNotificationsEnabled() bool`

GetAvailabilityNotificationsEnabled returns the AvailabilityNotificationsEnabled field if non-nil, zero value otherwise.

### GetAvailabilityNotificationsEnabledOk

`func (o *SyntheticPrivateLocationUpdate) GetAvailabilityNotificationsEnabledOk() (*bool, bool)`

GetAvailabilityNotificationsEnabledOk returns a tuple with the AvailabilityNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityNotificationsEnabled

`func (o *SyntheticPrivateLocationUpdate) SetAvailabilityNotificationsEnabled(v bool)`

SetAvailabilityNotificationsEnabled sets AvailabilityNotificationsEnabled field to given value.

### HasAvailabilityNotificationsEnabled

`func (o *SyntheticPrivateLocationUpdate) HasAvailabilityNotificationsEnabled() bool`

HasAvailabilityNotificationsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


