# DataPrivacyAndSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**MaskIpAddressesAndGpsCoordinates** | **bool** | Mask IP addresses and GPS coordinates enabled/disabled. | 
**MaskUserActionNames** | **bool** | Mask user action names (Web applications only) enabled/disabled. | 
**MaskPersonalDataInUris** | **bool** | Mask personal data in URIs enabled/disabled. | 
**LogAuditEvents** | Pointer to **bool** | The audit logging is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 

## Methods

### NewDataPrivacyAndSecurity

`func NewDataPrivacyAndSecurity(maskIpAddressesAndGpsCoordinates bool, maskUserActionNames bool, maskPersonalDataInUris bool, ) *DataPrivacyAndSecurity`

NewDataPrivacyAndSecurity instantiates a new DataPrivacyAndSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataPrivacyAndSecurityWithDefaults

`func NewDataPrivacyAndSecurityWithDefaults() *DataPrivacyAndSecurity`

NewDataPrivacyAndSecurityWithDefaults instantiates a new DataPrivacyAndSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *DataPrivacyAndSecurity) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DataPrivacyAndSecurity) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DataPrivacyAndSecurity) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DataPrivacyAndSecurity) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMaskIpAddressesAndGpsCoordinates

`func (o *DataPrivacyAndSecurity) GetMaskIpAddressesAndGpsCoordinates() bool`

GetMaskIpAddressesAndGpsCoordinates returns the MaskIpAddressesAndGpsCoordinates field if non-nil, zero value otherwise.

### GetMaskIpAddressesAndGpsCoordinatesOk

`func (o *DataPrivacyAndSecurity) GetMaskIpAddressesAndGpsCoordinatesOk() (*bool, bool)`

GetMaskIpAddressesAndGpsCoordinatesOk returns a tuple with the MaskIpAddressesAndGpsCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskIpAddressesAndGpsCoordinates

`func (o *DataPrivacyAndSecurity) SetMaskIpAddressesAndGpsCoordinates(v bool)`

SetMaskIpAddressesAndGpsCoordinates sets MaskIpAddressesAndGpsCoordinates field to given value.


### GetMaskUserActionNames

`func (o *DataPrivacyAndSecurity) GetMaskUserActionNames() bool`

GetMaskUserActionNames returns the MaskUserActionNames field if non-nil, zero value otherwise.

### GetMaskUserActionNamesOk

`func (o *DataPrivacyAndSecurity) GetMaskUserActionNamesOk() (*bool, bool)`

GetMaskUserActionNamesOk returns a tuple with the MaskUserActionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskUserActionNames

`func (o *DataPrivacyAndSecurity) SetMaskUserActionNames(v bool)`

SetMaskUserActionNames sets MaskUserActionNames field to given value.


### GetMaskPersonalDataInUris

`func (o *DataPrivacyAndSecurity) GetMaskPersonalDataInUris() bool`

GetMaskPersonalDataInUris returns the MaskPersonalDataInUris field if non-nil, zero value otherwise.

### GetMaskPersonalDataInUrisOk

`func (o *DataPrivacyAndSecurity) GetMaskPersonalDataInUrisOk() (*bool, bool)`

GetMaskPersonalDataInUrisOk returns a tuple with the MaskPersonalDataInUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskPersonalDataInUris

`func (o *DataPrivacyAndSecurity) SetMaskPersonalDataInUris(v bool)`

SetMaskPersonalDataInUris sets MaskPersonalDataInUris field to given value.


### GetLogAuditEvents

`func (o *DataPrivacyAndSecurity) GetLogAuditEvents() bool`

GetLogAuditEvents returns the LogAuditEvents field if non-nil, zero value otherwise.

### GetLogAuditEventsOk

`func (o *DataPrivacyAndSecurity) GetLogAuditEventsOk() (*bool, bool)`

GetLogAuditEventsOk returns a tuple with the LogAuditEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAuditEvents

`func (o *DataPrivacyAndSecurity) SetLogAuditEvents(v bool)`

SetLogAuditEvents sets LogAuditEvents field to given value.

### HasLogAuditEvents

`func (o *DataPrivacyAndSecurity) HasLogAuditEvents() bool`

HasLogAuditEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


