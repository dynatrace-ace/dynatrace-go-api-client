# RequestNaming

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the request naming rule. | [optional] 
**Order** | Pointer to **string** | The order string. Sorting request namings alphabetically by their order string determines their relative ordering.  Typically this is managed by Dynatrace internally and will not be present in GET responses nor used if present in PUT/POST requests, except where noted otherwise. | [optional] 
**Enabled** | **bool** | The rule is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**NamingPattern** | **string** | The name to be assigned to matching requests. | 
**ManagementZones** | Pointer to **[]string** | Specifies the management zones for which this rule should be applied. | [optional] 
**Conditions** | [**[]Condition**](Condition.md) | The set of conditions for the request naming rule usage.    You can specify several conditions. The request has to match **all** the specified conditions for the rule to trigger. | 
**Placeholders** | Pointer to [**[]Placeholder**](Placeholder.md) | The list of custom placeholders to be used in the naming pattern.    It enables you to extract a request attribute value or other request attribute and use it in the request naming pattern. | [optional] 

## Methods

### NewRequestNaming

`func NewRequestNaming(enabled bool, namingPattern string, conditions []Condition, ) *RequestNaming`

NewRequestNaming instantiates a new RequestNaming object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestNamingWithDefaults

`func NewRequestNamingWithDefaults() *RequestNaming`

NewRequestNamingWithDefaults instantiates a new RequestNaming object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *RequestNaming) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RequestNaming) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RequestNaming) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RequestNaming) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *RequestNaming) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestNaming) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestNaming) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequestNaming) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrder

`func (o *RequestNaming) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RequestNaming) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RequestNaming) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RequestNaming) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetEnabled

`func (o *RequestNaming) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RequestNaming) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RequestNaming) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetNamingPattern

`func (o *RequestNaming) GetNamingPattern() string`

GetNamingPattern returns the NamingPattern field if non-nil, zero value otherwise.

### GetNamingPatternOk

`func (o *RequestNaming) GetNamingPatternOk() (*string, bool)`

GetNamingPatternOk returns a tuple with the NamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingPattern

`func (o *RequestNaming) SetNamingPattern(v string)`

SetNamingPattern sets NamingPattern field to given value.


### GetManagementZones

`func (o *RequestNaming) GetManagementZones() []string`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *RequestNaming) GetManagementZonesOk() (*[]string, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *RequestNaming) SetManagementZones(v []string)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *RequestNaming) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetConditions

`func (o *RequestNaming) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *RequestNaming) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *RequestNaming) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.


### GetPlaceholders

`func (o *RequestNaming) GetPlaceholders() []Placeholder`

GetPlaceholders returns the Placeholders field if non-nil, zero value otherwise.

### GetPlaceholdersOk

`func (o *RequestNaming) GetPlaceholdersOk() (*[]Placeholder, bool)`

GetPlaceholdersOk returns a tuple with the Placeholders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholders

`func (o *RequestNaming) SetPlaceholders(v []Placeholder)`

SetPlaceholders sets Placeholders field to given value.

### HasPlaceholders

`func (o *RequestNaming) HasPlaceholders() bool`

HasPlaceholders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


