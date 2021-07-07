# PrivateSyntheticLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | **[]string** | A list of synthetic nodes belonging to the location.    You can retrieve the list of available nodes with the [GET all nodes](https://dt-url.net/miy3rpl) call. | 
**AvailabilityLocationOutage** | Pointer to **bool** | The alerting of location outage is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 
**AvailabilityNodeOutage** | Pointer to **bool** | The alerting of node outage is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;).    If enabled, the outage of *any* node in the location triggers an alert. | [optional] 
**LocationNodeOutageDelayInMinutes** | Pointer to **int32** | Alert if the location or node outage lasts longer than *X* minutes.    Only applicable when **availabilityLocationOutage** or **availabilityNodeOutage** is set to &#x60;true&#x60;. | [optional] 
**AvailabilityNotificationsEnabled** | Pointer to **bool** | The notifications of location and node outage is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 

## Methods

### NewPrivateSyntheticLocation

`func NewPrivateSyntheticLocation(nodes []string, ) *PrivateSyntheticLocation`

NewPrivateSyntheticLocation instantiates a new PrivateSyntheticLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSyntheticLocationWithDefaults

`func NewPrivateSyntheticLocationWithDefaults() *PrivateSyntheticLocation`

NewPrivateSyntheticLocationWithDefaults instantiates a new PrivateSyntheticLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *PrivateSyntheticLocation) GetNodes() []string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *PrivateSyntheticLocation) GetNodesOk() (*[]string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *PrivateSyntheticLocation) SetNodes(v []string)`

SetNodes sets Nodes field to given value.


### GetAvailabilityLocationOutage

`func (o *PrivateSyntheticLocation) GetAvailabilityLocationOutage() bool`

GetAvailabilityLocationOutage returns the AvailabilityLocationOutage field if non-nil, zero value otherwise.

### GetAvailabilityLocationOutageOk

`func (o *PrivateSyntheticLocation) GetAvailabilityLocationOutageOk() (*bool, bool)`

GetAvailabilityLocationOutageOk returns a tuple with the AvailabilityLocationOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityLocationOutage

`func (o *PrivateSyntheticLocation) SetAvailabilityLocationOutage(v bool)`

SetAvailabilityLocationOutage sets AvailabilityLocationOutage field to given value.

### HasAvailabilityLocationOutage

`func (o *PrivateSyntheticLocation) HasAvailabilityLocationOutage() bool`

HasAvailabilityLocationOutage returns a boolean if a field has been set.

### GetAvailabilityNodeOutage

`func (o *PrivateSyntheticLocation) GetAvailabilityNodeOutage() bool`

GetAvailabilityNodeOutage returns the AvailabilityNodeOutage field if non-nil, zero value otherwise.

### GetAvailabilityNodeOutageOk

`func (o *PrivateSyntheticLocation) GetAvailabilityNodeOutageOk() (*bool, bool)`

GetAvailabilityNodeOutageOk returns a tuple with the AvailabilityNodeOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityNodeOutage

`func (o *PrivateSyntheticLocation) SetAvailabilityNodeOutage(v bool)`

SetAvailabilityNodeOutage sets AvailabilityNodeOutage field to given value.

### HasAvailabilityNodeOutage

`func (o *PrivateSyntheticLocation) HasAvailabilityNodeOutage() bool`

HasAvailabilityNodeOutage returns a boolean if a field has been set.

### GetLocationNodeOutageDelayInMinutes

`func (o *PrivateSyntheticLocation) GetLocationNodeOutageDelayInMinutes() int32`

GetLocationNodeOutageDelayInMinutes returns the LocationNodeOutageDelayInMinutes field if non-nil, zero value otherwise.

### GetLocationNodeOutageDelayInMinutesOk

`func (o *PrivateSyntheticLocation) GetLocationNodeOutageDelayInMinutesOk() (*int32, bool)`

GetLocationNodeOutageDelayInMinutesOk returns a tuple with the LocationNodeOutageDelayInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationNodeOutageDelayInMinutes

`func (o *PrivateSyntheticLocation) SetLocationNodeOutageDelayInMinutes(v int32)`

SetLocationNodeOutageDelayInMinutes sets LocationNodeOutageDelayInMinutes field to given value.

### HasLocationNodeOutageDelayInMinutes

`func (o *PrivateSyntheticLocation) HasLocationNodeOutageDelayInMinutes() bool`

HasLocationNodeOutageDelayInMinutes returns a boolean if a field has been set.

### GetAvailabilityNotificationsEnabled

`func (o *PrivateSyntheticLocation) GetAvailabilityNotificationsEnabled() bool`

GetAvailabilityNotificationsEnabled returns the AvailabilityNotificationsEnabled field if non-nil, zero value otherwise.

### GetAvailabilityNotificationsEnabledOk

`func (o *PrivateSyntheticLocation) GetAvailabilityNotificationsEnabledOk() (*bool, bool)`

GetAvailabilityNotificationsEnabledOk returns a tuple with the AvailabilityNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityNotificationsEnabled

`func (o *PrivateSyntheticLocation) SetAvailabilityNotificationsEnabled(v bool)`

SetAvailabilityNotificationsEnabled sets AvailabilityNotificationsEnabled field to given value.

### HasAvailabilityNotificationsEnabled

`func (o *PrivateSyntheticLocation) HasAvailabilityNotificationsEnabled() bool`

HasAvailabilityNotificationsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


