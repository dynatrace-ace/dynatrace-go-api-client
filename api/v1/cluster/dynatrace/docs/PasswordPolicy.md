# PasswordPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RealmId** | Pointer to **string** |  | [optional] 
**MinPasswordLength** | **int64** | Minimum password length | 
**MinNumberOfUppercaseChars** | **int64** | Minimum number of uppercase characters | 
**MinNumberOfLowercaseChars** | **int64** | Minimum number of lowercase characters | 
**MinNumberOfDigits** | **int64** | Minimum number of digits | 
**MinNumberOfNonAlphanumericChars** | **int64** | Minimum number of non-alphanumeric characters | 

## Methods

### NewPasswordPolicy

`func NewPasswordPolicy(minPasswordLength int64, minNumberOfUppercaseChars int64, minNumberOfLowercaseChars int64, minNumberOfDigits int64, minNumberOfNonAlphanumericChars int64, ) *PasswordPolicy`

NewPasswordPolicy instantiates a new PasswordPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyWithDefaults

`func NewPasswordPolicyWithDefaults() *PasswordPolicy`

NewPasswordPolicyWithDefaults instantiates a new PasswordPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRealmId

`func (o *PasswordPolicy) GetRealmId() string`

GetRealmId returns the RealmId field if non-nil, zero value otherwise.

### GetRealmIdOk

`func (o *PasswordPolicy) GetRealmIdOk() (*string, bool)`

GetRealmIdOk returns a tuple with the RealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmId

`func (o *PasswordPolicy) SetRealmId(v string)`

SetRealmId sets RealmId field to given value.

### HasRealmId

`func (o *PasswordPolicy) HasRealmId() bool`

HasRealmId returns a boolean if a field has been set.

### GetMinPasswordLength

`func (o *PasswordPolicy) GetMinPasswordLength() int64`

GetMinPasswordLength returns the MinPasswordLength field if non-nil, zero value otherwise.

### GetMinPasswordLengthOk

`func (o *PasswordPolicy) GetMinPasswordLengthOk() (*int64, bool)`

GetMinPasswordLengthOk returns a tuple with the MinPasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPasswordLength

`func (o *PasswordPolicy) SetMinPasswordLength(v int64)`

SetMinPasswordLength sets MinPasswordLength field to given value.


### GetMinNumberOfUppercaseChars

`func (o *PasswordPolicy) GetMinNumberOfUppercaseChars() int64`

GetMinNumberOfUppercaseChars returns the MinNumberOfUppercaseChars field if non-nil, zero value otherwise.

### GetMinNumberOfUppercaseCharsOk

`func (o *PasswordPolicy) GetMinNumberOfUppercaseCharsOk() (*int64, bool)`

GetMinNumberOfUppercaseCharsOk returns a tuple with the MinNumberOfUppercaseChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumberOfUppercaseChars

`func (o *PasswordPolicy) SetMinNumberOfUppercaseChars(v int64)`

SetMinNumberOfUppercaseChars sets MinNumberOfUppercaseChars field to given value.


### GetMinNumberOfLowercaseChars

`func (o *PasswordPolicy) GetMinNumberOfLowercaseChars() int64`

GetMinNumberOfLowercaseChars returns the MinNumberOfLowercaseChars field if non-nil, zero value otherwise.

### GetMinNumberOfLowercaseCharsOk

`func (o *PasswordPolicy) GetMinNumberOfLowercaseCharsOk() (*int64, bool)`

GetMinNumberOfLowercaseCharsOk returns a tuple with the MinNumberOfLowercaseChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumberOfLowercaseChars

`func (o *PasswordPolicy) SetMinNumberOfLowercaseChars(v int64)`

SetMinNumberOfLowercaseChars sets MinNumberOfLowercaseChars field to given value.


### GetMinNumberOfDigits

`func (o *PasswordPolicy) GetMinNumberOfDigits() int64`

GetMinNumberOfDigits returns the MinNumberOfDigits field if non-nil, zero value otherwise.

### GetMinNumberOfDigitsOk

`func (o *PasswordPolicy) GetMinNumberOfDigitsOk() (*int64, bool)`

GetMinNumberOfDigitsOk returns a tuple with the MinNumberOfDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumberOfDigits

`func (o *PasswordPolicy) SetMinNumberOfDigits(v int64)`

SetMinNumberOfDigits sets MinNumberOfDigits field to given value.


### GetMinNumberOfNonAlphanumericChars

`func (o *PasswordPolicy) GetMinNumberOfNonAlphanumericChars() int64`

GetMinNumberOfNonAlphanumericChars returns the MinNumberOfNonAlphanumericChars field if non-nil, zero value otherwise.

### GetMinNumberOfNonAlphanumericCharsOk

`func (o *PasswordPolicy) GetMinNumberOfNonAlphanumericCharsOk() (*int64, bool)`

GetMinNumberOfNonAlphanumericCharsOk returns a tuple with the MinNumberOfNonAlphanumericChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumberOfNonAlphanumericChars

`func (o *PasswordPolicy) SetMinNumberOfNonAlphanumericChars(v int64)`

SetMinNumberOfNonAlphanumericChars sets MinNumberOfNonAlphanumericChars field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


