# ElasticsearchUpgradeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradePossible** | Pointer to **bool** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewElasticsearchUpgradeDTO

`func NewElasticsearchUpgradeDTO() *ElasticsearchUpgradeDTO`

NewElasticsearchUpgradeDTO instantiates a new ElasticsearchUpgradeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElasticsearchUpgradeDTOWithDefaults

`func NewElasticsearchUpgradeDTOWithDefaults() *ElasticsearchUpgradeDTO`

NewElasticsearchUpgradeDTOWithDefaults instantiates a new ElasticsearchUpgradeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradePossible

`func (o *ElasticsearchUpgradeDTO) GetUpgradePossible() bool`

GetUpgradePossible returns the UpgradePossible field if non-nil, zero value otherwise.

### GetUpgradePossibleOk

`func (o *ElasticsearchUpgradeDTO) GetUpgradePossibleOk() (*bool, bool)`

GetUpgradePossibleOk returns a tuple with the UpgradePossible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePossible

`func (o *ElasticsearchUpgradeDTO) SetUpgradePossible(v bool)`

SetUpgradePossible sets UpgradePossible field to given value.

### HasUpgradePossible

`func (o *ElasticsearchUpgradeDTO) HasUpgradePossible() bool`

HasUpgradePossible returns a boolean if a field has been set.

### GetReason

`func (o *ElasticsearchUpgradeDTO) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ElasticsearchUpgradeDTO) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ElasticsearchUpgradeDTO) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ElasticsearchUpgradeDTO) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


