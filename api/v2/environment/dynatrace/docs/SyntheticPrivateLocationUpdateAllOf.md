# SyntheticPrivateLocationUpdateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to **[]string** | A list of synthetic nodes belonging to the location.    You can retrieve the list of available nodes with the [GET all nodes](https://dt-url.net/miy3rpl) call. | [optional] 
**Name** | Pointer to **string** | The name of the location. | [optional] 
**CountryCode** | Pointer to **string** | The country code of the location.    Use the alpha-2 code of the [ISO 3166-2 standard](https://dt-url.net/iso3166-2), (for example, &#x60;AT&#x60; for Austria or &#x60;PL&#x60; for Poland). | [optional] 
**RegionCode** | Pointer to **string** | The region code of the location.    For the [USA](https://dt-url.net/iso3166us) or [Canada](https://dt-url.net/iso3166ca) use ISO 3166-2 state codes (without &#x60;US-&#x60; or &#x60;CA-&#x60; prefix), for example, &#x60;VA&#x60; for Virginia or &#x60;OR&#x60; for Oregon.    For the rest of the world use [FIPS 10-4 codes](https://dt-url.net/fipscodes). | [optional] 
**City** | Pointer to **string** | The city of the location. | [optional] 
**Latitude** | Pointer to **float64** | The latitude of the location in &#x60;DDD.dddd&#x60; format. | [optional] 
**Longitude** | Pointer to **float64** | The longitude of the location in &#x60;DDD.dddd&#x60; format. | [optional] 
**Status** | Pointer to **string** | The status of the location:   * &#x60;ENABLED&#x60;: The location is displayed as active in the UI. You can assign monitors to the location.  * &#x60;DISABLED&#x60;: The location is displayed as inactive in the UI. You can&#39;t assign monitors to the location. Monitors already assigned to the location will stay there and will be executed from the location.  * &#x60;HIDDEN&#x60;: The location is not displayed in the UI. You can&#39;t assign monitors to the location. You can only set location as &#x60;HIDDEN&#x60; when no monitor is assigned to it. | [optional] 
**AvailabilityLocationOutage** | Pointer to **bool** | The alerting of location outage is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 
**AvailabilityNodeOutage** | Pointer to **bool** | The alerting of node outage is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;).    If enabled, the outage of *any* node in the location triggers an alert. | [optional] 
**LocationNodeOutageDelayInMinutes** | Pointer to **int32** | Alert if the location or node outage lasts longer than *X* minutes.    Only applicable when **availabilityLocationOutage** or **availabilityNodeOutage** is set to &#x60;true&#x60;. | [optional] 
**AvailabilityNotificationsEnabled** | Pointer to **bool** | The notifications of location and node outage is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 

## Methods

### NewSyntheticPrivateLocationUpdateAllOf

`func NewSyntheticPrivateLocationUpdateAllOf() *SyntheticPrivateLocationUpdateAllOf`

NewSyntheticPrivateLocationUpdateAllOf instantiates a new SyntheticPrivateLocationUpdateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticPrivateLocationUpdateAllOfWithDefaults

`func NewSyntheticPrivateLocationUpdateAllOfWithDefaults() *SyntheticPrivateLocationUpdateAllOf`

NewSyntheticPrivateLocationUpdateAllOfWithDefaults instantiates a new SyntheticPrivateLocationUpdateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *SyntheticPrivateLocationUpdateAllOf) GetNodes() []string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *SyntheticPrivateLocationUpdateAllOf) GetNodesOk() (*[]string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *SyntheticPrivateLocationUpdateAllOf) SetNodes(v []string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *SyntheticPrivateLocationUpdateAllOf) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetName

`func (o *SyntheticPrivateLocationUpdateAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticPrivateLocationUpdateAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticPrivateLocationUpdateAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyntheticPrivateLocationUpdateAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCountryCode

`func (o *SyntheticPrivateLocationUpdateAllOf) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SyntheticPrivateLocationUpdateAllOf) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SyntheticPrivateLocationUpdateAllOf) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SyntheticPrivateLocationUpdateAllOf) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetRegionCode

`func (o *SyntheticPrivateLocationUpdateAllOf) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *SyntheticPrivateLocationUpdateAllOf) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *SyntheticPrivateLocationUpdateAllOf) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *SyntheticPrivateLocationUpdateAllOf) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetCity

`func (o *SyntheticPrivateLocationUpdateAllOf) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SyntheticPrivateLocationUpdateAllOf) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SyntheticPrivateLocationUpdateAllOf) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SyntheticPrivateLocationUpdateAllOf) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetLatitude

`func (o *SyntheticPrivateLocationUpdateAllOf) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SyntheticPrivateLocationUpdateAllOf) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SyntheticPrivateLocationUpdateAllOf) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *SyntheticPrivateLocationUpdateAllOf) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *SyntheticPrivateLocationUpdateAllOf) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SyntheticPrivateLocationUpdateAllOf) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SyntheticPrivateLocationUpdateAllOf) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *SyntheticPrivateLocationUpdateAllOf) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetStatus

`func (o *SyntheticPrivateLocationUpdateAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticPrivateLocationUpdateAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticPrivateLocationUpdateAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyntheticPrivateLocationUpdateAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAvailabilityLocationOutage

`func (o *SyntheticPrivateLocationUpdateAllOf) GetAvailabilityLocationOutage() bool`

GetAvailabilityLocationOutage returns the AvailabilityLocationOutage field if non-nil, zero value otherwise.

### GetAvailabilityLocationOutageOk

`func (o *SyntheticPrivateLocationUpdateAllOf) GetAvailabilityLocationOutageOk() (*bool, bool)`

GetAvailabilityLocationOutageOk returns a tuple with the AvailabilityLocationOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityLocationOutage

`func (o *SyntheticPrivateLocationUpdateAllOf) SetAvailabilityLocationOutage(v bool)`

SetAvailabilityLocationOutage sets AvailabilityLocationOutage field to given value.

### HasAvailabilityLocationOutage

`func (o *SyntheticPrivateLocationUpdateAllOf) HasAvailabilityLocationOutage() bool`

HasAvailabilityLocationOutage returns a boolean if a field has been set.

### GetAvailabilityNodeOutage

`func (o *SyntheticPrivateLocationUpdateAllOf) GetAvailabilityNodeOutage() bool`

GetAvailabilityNodeOutage returns the AvailabilityNodeOutage field if non-nil, zero value otherwise.

### GetAvailabilityNodeOutageOk

`func (o *SyntheticPrivateLocationUpdateAllOf) GetAvailabilityNodeOutageOk() (*bool, bool)`

GetAvailabilityNodeOutageOk returns a tuple with the AvailabilityNodeOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityNodeOutage

`func (o *SyntheticPrivateLocationUpdateAllOf) SetAvailabilityNodeOutage(v bool)`

SetAvailabilityNodeOutage sets AvailabilityNodeOutage field to given value.

### HasAvailabilityNodeOutage

`func (o *SyntheticPrivateLocationUpdateAllOf) HasAvailabilityNodeOutage() bool`

HasAvailabilityNodeOutage returns a boolean if a field has been set.

### GetLocationNodeOutageDelayInMinutes

`func (o *SyntheticPrivateLocationUpdateAllOf) GetLocationNodeOutageDelayInMinutes() int32`

GetLocationNodeOutageDelayInMinutes returns the LocationNodeOutageDelayInMinutes field if non-nil, zero value otherwise.

### GetLocationNodeOutageDelayInMinutesOk

`func (o *SyntheticPrivateLocationUpdateAllOf) GetLocationNodeOutageDelayInMinutesOk() (*int32, bool)`

GetLocationNodeOutageDelayInMinutesOk returns a tuple with the LocationNodeOutageDelayInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationNodeOutageDelayInMinutes

`func (o *SyntheticPrivateLocationUpdateAllOf) SetLocationNodeOutageDelayInMinutes(v int32)`

SetLocationNodeOutageDelayInMinutes sets LocationNodeOutageDelayInMinutes field to given value.

### HasLocationNodeOutageDelayInMinutes

`func (o *SyntheticPrivateLocationUpdateAllOf) HasLocationNodeOutageDelayInMinutes() bool`

HasLocationNodeOutageDelayInMinutes returns a boolean if a field has been set.

### GetAvailabilityNotificationsEnabled

`func (o *SyntheticPrivateLocationUpdateAllOf) GetAvailabilityNotificationsEnabled() bool`

GetAvailabilityNotificationsEnabled returns the AvailabilityNotificationsEnabled field if non-nil, zero value otherwise.

### GetAvailabilityNotificationsEnabledOk

`func (o *SyntheticPrivateLocationUpdateAllOf) GetAvailabilityNotificationsEnabledOk() (*bool, bool)`

GetAvailabilityNotificationsEnabledOk returns a tuple with the AvailabilityNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityNotificationsEnabled

`func (o *SyntheticPrivateLocationUpdateAllOf) SetAvailabilityNotificationsEnabled(v bool)`

SetAvailabilityNotificationsEnabled sets AvailabilityNotificationsEnabled field to given value.

### HasAvailabilityNotificationsEnabled

`func (o *SyntheticPrivateLocationUpdateAllOf) HasAvailabilityNotificationsEnabled() bool`

HasAvailabilityNotificationsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


