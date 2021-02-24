# DestinationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UrlOrPath** | **string** | The path to be reached to hit the conversion goal. | 
**MatchType** | Pointer to **string** | The operator of the match. | [optional] 
**CaseSensitive** | Pointer to **bool** | The match is case-sensitive (&#x60;true&#x60;) or (&#x60;false&#x60;). | [optional] 

## Methods

### NewDestinationDetails

`func NewDestinationDetails(urlOrPath string, ) *DestinationDetails`

NewDestinationDetails instantiates a new DestinationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationDetailsWithDefaults

`func NewDestinationDetailsWithDefaults() *DestinationDetails`

NewDestinationDetailsWithDefaults instantiates a new DestinationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrlOrPath

`func (o *DestinationDetails) GetUrlOrPath() string`

GetUrlOrPath returns the UrlOrPath field if non-nil, zero value otherwise.

### GetUrlOrPathOk

`func (o *DestinationDetails) GetUrlOrPathOk() (*string, bool)`

GetUrlOrPathOk returns a tuple with the UrlOrPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlOrPath

`func (o *DestinationDetails) SetUrlOrPath(v string)`

SetUrlOrPath sets UrlOrPath field to given value.


### GetMatchType

`func (o *DestinationDetails) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *DestinationDetails) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *DestinationDetails) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *DestinationDetails) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *DestinationDetails) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *DestinationDetails) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *DestinationDetails) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *DestinationDetails) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


