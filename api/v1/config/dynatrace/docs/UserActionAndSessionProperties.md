# UserActionAndSessionProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display name of the property. | [optional] 
**Type** | **string** | The data type of the property. | 
**Origin** | **string** | The origin of the property | 
**Aggregation** | Pointer to **string** | The aggregation type of the property.     It defines how multiple values of the property are aggregated. | [optional] 
**StoreAsUserActionProperty** | Pointer to **bool** | If &#x60;true&#x60;, the property is stored as a user action property | [optional] 
**StoreAsSessionProperty** | Pointer to **bool** | If &#x60;true&#x60;, the property is stored as a session property | [optional] 
**CleanupRule** | Pointer to **string** | The cleanup rule of the property.   Defines how to extract the data you need from a string value. Specify the [regular expression](https://www.dynatrace.com/support/help/shortlink/regex) for the data you need there. | [optional] 
**ServerSideRequestAttribute** | Pointer to **string** | The ID of the request attribute.   Only applicable when the **origin** is set to &#x60;SERVER_SIDE_REQUEST_ATTRIBUTE&#x60;. | [optional] 
**UniqueId** | **int32** | Unique id among all userTags and properties of this application | 
**Key** | **string** | Key of the property | 
**MetadataId** | Pointer to **int32** | If the origin is META_DATA, metaData id of the property | [optional] 
**IgnoreCase** | Pointer to **bool** | If true, the value of this property will always be stored in lower case. Defaults to false. | [optional] 

## Methods

### NewUserActionAndSessionProperties

`func NewUserActionAndSessionProperties(type_ string, origin string, uniqueId int32, key string, ) *UserActionAndSessionProperties`

NewUserActionAndSessionProperties instantiates a new UserActionAndSessionProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActionAndSessionPropertiesWithDefaults

`func NewUserActionAndSessionPropertiesWithDefaults() *UserActionAndSessionProperties`

NewUserActionAndSessionPropertiesWithDefaults instantiates a new UserActionAndSessionProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *UserActionAndSessionProperties) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserActionAndSessionProperties) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserActionAndSessionProperties) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserActionAndSessionProperties) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetType

`func (o *UserActionAndSessionProperties) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserActionAndSessionProperties) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserActionAndSessionProperties) SetType(v string)`

SetType sets Type field to given value.


### GetOrigin

`func (o *UserActionAndSessionProperties) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *UserActionAndSessionProperties) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *UserActionAndSessionProperties) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetAggregation

`func (o *UserActionAndSessionProperties) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *UserActionAndSessionProperties) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *UserActionAndSessionProperties) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *UserActionAndSessionProperties) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetStoreAsUserActionProperty

`func (o *UserActionAndSessionProperties) GetStoreAsUserActionProperty() bool`

GetStoreAsUserActionProperty returns the StoreAsUserActionProperty field if non-nil, zero value otherwise.

### GetStoreAsUserActionPropertyOk

`func (o *UserActionAndSessionProperties) GetStoreAsUserActionPropertyOk() (*bool, bool)`

GetStoreAsUserActionPropertyOk returns a tuple with the StoreAsUserActionProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreAsUserActionProperty

`func (o *UserActionAndSessionProperties) SetStoreAsUserActionProperty(v bool)`

SetStoreAsUserActionProperty sets StoreAsUserActionProperty field to given value.

### HasStoreAsUserActionProperty

`func (o *UserActionAndSessionProperties) HasStoreAsUserActionProperty() bool`

HasStoreAsUserActionProperty returns a boolean if a field has been set.

### GetStoreAsSessionProperty

`func (o *UserActionAndSessionProperties) GetStoreAsSessionProperty() bool`

GetStoreAsSessionProperty returns the StoreAsSessionProperty field if non-nil, zero value otherwise.

### GetStoreAsSessionPropertyOk

`func (o *UserActionAndSessionProperties) GetStoreAsSessionPropertyOk() (*bool, bool)`

GetStoreAsSessionPropertyOk returns a tuple with the StoreAsSessionProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreAsSessionProperty

`func (o *UserActionAndSessionProperties) SetStoreAsSessionProperty(v bool)`

SetStoreAsSessionProperty sets StoreAsSessionProperty field to given value.

### HasStoreAsSessionProperty

`func (o *UserActionAndSessionProperties) HasStoreAsSessionProperty() bool`

HasStoreAsSessionProperty returns a boolean if a field has been set.

### GetCleanupRule

`func (o *UserActionAndSessionProperties) GetCleanupRule() string`

GetCleanupRule returns the CleanupRule field if non-nil, zero value otherwise.

### GetCleanupRuleOk

`func (o *UserActionAndSessionProperties) GetCleanupRuleOk() (*string, bool)`

GetCleanupRuleOk returns a tuple with the CleanupRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupRule

`func (o *UserActionAndSessionProperties) SetCleanupRule(v string)`

SetCleanupRule sets CleanupRule field to given value.

### HasCleanupRule

`func (o *UserActionAndSessionProperties) HasCleanupRule() bool`

HasCleanupRule returns a boolean if a field has been set.

### GetServerSideRequestAttribute

`func (o *UserActionAndSessionProperties) GetServerSideRequestAttribute() string`

GetServerSideRequestAttribute returns the ServerSideRequestAttribute field if non-nil, zero value otherwise.

### GetServerSideRequestAttributeOk

`func (o *UserActionAndSessionProperties) GetServerSideRequestAttributeOk() (*string, bool)`

GetServerSideRequestAttributeOk returns a tuple with the ServerSideRequestAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSideRequestAttribute

`func (o *UserActionAndSessionProperties) SetServerSideRequestAttribute(v string)`

SetServerSideRequestAttribute sets ServerSideRequestAttribute field to given value.

### HasServerSideRequestAttribute

`func (o *UserActionAndSessionProperties) HasServerSideRequestAttribute() bool`

HasServerSideRequestAttribute returns a boolean if a field has been set.

### GetUniqueId

`func (o *UserActionAndSessionProperties) GetUniqueId() int32`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UserActionAndSessionProperties) GetUniqueIdOk() (*int32, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UserActionAndSessionProperties) SetUniqueId(v int32)`

SetUniqueId sets UniqueId field to given value.


### GetKey

`func (o *UserActionAndSessionProperties) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UserActionAndSessionProperties) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UserActionAndSessionProperties) SetKey(v string)`

SetKey sets Key field to given value.


### GetMetadataId

`func (o *UserActionAndSessionProperties) GetMetadataId() int32`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *UserActionAndSessionProperties) GetMetadataIdOk() (*int32, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *UserActionAndSessionProperties) SetMetadataId(v int32)`

SetMetadataId sets MetadataId field to given value.

### HasMetadataId

`func (o *UserActionAndSessionProperties) HasMetadataId() bool`

HasMetadataId returns a boolean if a field has been set.

### GetIgnoreCase

`func (o *UserActionAndSessionProperties) GetIgnoreCase() bool`

GetIgnoreCase returns the IgnoreCase field if non-nil, zero value otherwise.

### GetIgnoreCaseOk

`func (o *UserActionAndSessionProperties) GetIgnoreCaseOk() (*bool, bool)`

GetIgnoreCaseOk returns a tuple with the IgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCase

`func (o *UserActionAndSessionProperties) SetIgnoreCase(v bool)`

SetIgnoreCase sets IgnoreCase field to given value.

### HasIgnoreCase

`func (o *UserActionAndSessionProperties) HasIgnoreCase() bool`

HasIgnoreCase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


