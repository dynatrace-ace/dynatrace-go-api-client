# TransactionTrafficQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxLimit** | Pointer to **int32** | Maximum traffic [units per minute] | [optional] 

## Methods

### NewTransactionTrafficQuota

`func NewTransactionTrafficQuota() *TransactionTrafficQuota`

NewTransactionTrafficQuota instantiates a new TransactionTrafficQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionTrafficQuotaWithDefaults

`func NewTransactionTrafficQuotaWithDefaults() *TransactionTrafficQuota`

NewTransactionTrafficQuotaWithDefaults instantiates a new TransactionTrafficQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxLimit

`func (o *TransactionTrafficQuota) GetMaxLimit() int32`

GetMaxLimit returns the MaxLimit field if non-nil, zero value otherwise.

### GetMaxLimitOk

`func (o *TransactionTrafficQuota) GetMaxLimitOk() (*int32, bool)`

GetMaxLimitOk returns a tuple with the MaxLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLimit

`func (o *TransactionTrafficQuota) SetMaxLimit(v int32)`

SetMaxLimit sets MaxLimit field to given value.

### HasMaxLimit

`func (o *TransactionTrafficQuota) HasMaxLimit() bool`

HasMaxLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


