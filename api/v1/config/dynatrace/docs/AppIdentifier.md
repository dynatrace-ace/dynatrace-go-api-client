# AppIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id where the file belongs to | [optional] 
**VersionCode** | **string** | The version code (Android) / bundle version (iOS) the file belongs to | 
**VersionName** | **string** | The version name (Android) / bundle versions string (iOS) the file belongs to | 
**PackageName** | Pointer to **string** | The bundleId (iOS) or package name (Android) the file belongs to | [optional] 
**Os** | Pointer to **string** | The operating system where the file belongs to | [optional] 

## Methods

### NewAppIdentifier

`func NewAppIdentifier(versionCode string, versionName string, ) *AppIdentifier`

NewAppIdentifier instantiates a new AppIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppIdentifierWithDefaults

`func NewAppIdentifierWithDefaults() *AppIdentifier`

NewAppIdentifierWithDefaults instantiates a new AppIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppIdentifier) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppIdentifier) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppIdentifier) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppIdentifier) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionCode

`func (o *AppIdentifier) GetVersionCode() string`

GetVersionCode returns the VersionCode field if non-nil, zero value otherwise.

### GetVersionCodeOk

`func (o *AppIdentifier) GetVersionCodeOk() (*string, bool)`

GetVersionCodeOk returns a tuple with the VersionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCode

`func (o *AppIdentifier) SetVersionCode(v string)`

SetVersionCode sets VersionCode field to given value.


### GetVersionName

`func (o *AppIdentifier) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *AppIdentifier) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *AppIdentifier) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.


### GetPackageName

`func (o *AppIdentifier) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *AppIdentifier) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *AppIdentifier) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *AppIdentifier) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetOs

`func (o *AppIdentifier) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *AppIdentifier) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *AppIdentifier) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *AppIdentifier) HasOs() bool`

HasOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


