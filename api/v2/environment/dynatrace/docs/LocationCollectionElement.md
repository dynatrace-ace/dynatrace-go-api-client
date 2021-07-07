# LocationCollectionElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the location. | 
**EntityId** | **string** | The Dynatrace entity ID of the location. | 
**Type** | **string** | The type of the location. | 
**CloudPlatform** | Pointer to **string** | The cloud provider where the location is hosted.    Only applicable to &#x60;PUBLIC&#x60; locations. | [optional] 
**Ips** | Pointer to **[]string** | The list of IP addresses assigned to the location.    Only applicable to &#x60;PUBLIC&#x60; locations. | [optional] 
**Stage** | Pointer to **string** | The release stage of the location. | [optional] 
**Status** | Pointer to **string** | The status of the location. | [optional] 
**GeoLocationId** | **string** | The Dynatrace GeoLocation ID of the location. | 

## Methods

### NewLocationCollectionElement

`func NewLocationCollectionElement(name string, entityId string, type_ string, geoLocationId string, ) *LocationCollectionElement`

NewLocationCollectionElement instantiates a new LocationCollectionElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationCollectionElementWithDefaults

`func NewLocationCollectionElementWithDefaults() *LocationCollectionElement`

NewLocationCollectionElementWithDefaults instantiates a new LocationCollectionElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LocationCollectionElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocationCollectionElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocationCollectionElement) SetName(v string)`

SetName sets Name field to given value.


### GetEntityId

`func (o *LocationCollectionElement) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *LocationCollectionElement) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *LocationCollectionElement) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetType

`func (o *LocationCollectionElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LocationCollectionElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LocationCollectionElement) SetType(v string)`

SetType sets Type field to given value.


### GetCloudPlatform

`func (o *LocationCollectionElement) GetCloudPlatform() string`

GetCloudPlatform returns the CloudPlatform field if non-nil, zero value otherwise.

### GetCloudPlatformOk

`func (o *LocationCollectionElement) GetCloudPlatformOk() (*string, bool)`

GetCloudPlatformOk returns a tuple with the CloudPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPlatform

`func (o *LocationCollectionElement) SetCloudPlatform(v string)`

SetCloudPlatform sets CloudPlatform field to given value.

### HasCloudPlatform

`func (o *LocationCollectionElement) HasCloudPlatform() bool`

HasCloudPlatform returns a boolean if a field has been set.

### GetIps

`func (o *LocationCollectionElement) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *LocationCollectionElement) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *LocationCollectionElement) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *LocationCollectionElement) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetStage

`func (o *LocationCollectionElement) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *LocationCollectionElement) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *LocationCollectionElement) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *LocationCollectionElement) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStatus

`func (o *LocationCollectionElement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LocationCollectionElement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LocationCollectionElement) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LocationCollectionElement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGeoLocationId

`func (o *LocationCollectionElement) GetGeoLocationId() string`

GetGeoLocationId returns the GeoLocationId field if non-nil, zero value otherwise.

### GetGeoLocationIdOk

`func (o *LocationCollectionElement) GetGeoLocationIdOk() (*string, bool)`

GetGeoLocationIdOk returns a tuple with the GeoLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoLocationId

`func (o *LocationCollectionElement) SetGeoLocationId(v string)`

SetGeoLocationId sets GeoLocationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


