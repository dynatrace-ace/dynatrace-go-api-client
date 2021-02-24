# AllowedBeaconOrigins

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**AllowedBeaconOrigins** | Pointer to [**[]BeaconDomainPattern**](BeaconDomainPattern.md) | A list of allowed beacon origins for CORS requests. | [optional] 
**RejectBeaconsWithoutOriginHeader** | Pointer to **bool** | Discard (&#x60;true&#x60;) or keep (&#x60;false&#x60;) beacons without the **Origin** HTTP header on the BeaconForwarder.   If not set, &#x60;false&#x60; is used. | [optional] 

## Methods

### NewAllowedBeaconOrigins

`func NewAllowedBeaconOrigins() *AllowedBeaconOrigins`

NewAllowedBeaconOrigins instantiates a new AllowedBeaconOrigins object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedBeaconOriginsWithDefaults

`func NewAllowedBeaconOriginsWithDefaults() *AllowedBeaconOrigins`

NewAllowedBeaconOriginsWithDefaults instantiates a new AllowedBeaconOrigins object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *AllowedBeaconOrigins) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AllowedBeaconOrigins) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AllowedBeaconOrigins) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AllowedBeaconOrigins) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAllowedBeaconOrigins

`func (o *AllowedBeaconOrigins) GetAllowedBeaconOrigins() []BeaconDomainPattern`

GetAllowedBeaconOrigins returns the AllowedBeaconOrigins field if non-nil, zero value otherwise.

### GetAllowedBeaconOriginsOk

`func (o *AllowedBeaconOrigins) GetAllowedBeaconOriginsOk() (*[]BeaconDomainPattern, bool)`

GetAllowedBeaconOriginsOk returns a tuple with the AllowedBeaconOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedBeaconOrigins

`func (o *AllowedBeaconOrigins) SetAllowedBeaconOrigins(v []BeaconDomainPattern)`

SetAllowedBeaconOrigins sets AllowedBeaconOrigins field to given value.

### HasAllowedBeaconOrigins

`func (o *AllowedBeaconOrigins) HasAllowedBeaconOrigins() bool`

HasAllowedBeaconOrigins returns a boolean if a field has been set.

### GetRejectBeaconsWithoutOriginHeader

`func (o *AllowedBeaconOrigins) GetRejectBeaconsWithoutOriginHeader() bool`

GetRejectBeaconsWithoutOriginHeader returns the RejectBeaconsWithoutOriginHeader field if non-nil, zero value otherwise.

### GetRejectBeaconsWithoutOriginHeaderOk

`func (o *AllowedBeaconOrigins) GetRejectBeaconsWithoutOriginHeaderOk() (*bool, bool)`

GetRejectBeaconsWithoutOriginHeaderOk returns a tuple with the RejectBeaconsWithoutOriginHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectBeaconsWithoutOriginHeader

`func (o *AllowedBeaconOrigins) SetRejectBeaconsWithoutOriginHeader(v bool)`

SetRejectBeaconsWithoutOriginHeader sets RejectBeaconsWithoutOriginHeader field to given value.

### HasRejectBeaconsWithoutOriginHeader

`func (o *AllowedBeaconOrigins) HasRejectBeaconsWithoutOriginHeader() bool`

HasRejectBeaconsWithoutOriginHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


