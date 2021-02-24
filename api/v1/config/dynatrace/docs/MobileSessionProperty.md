# MobileSessionProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The unique key of the session property. | 
**DisplayName** | Pointer to **string** | The display name of the property. | [optional] 
**Type** | **string** | The data type of the property. | 
**Origin** | **string** | The origin of the property | 
**Aggregation** | Pointer to **string** | The aggregation type of the property.     It defines how multiple values of the property are aggregated. | [optional] 
**StoreAsUserActionProperty** | Pointer to **bool** | If &#x60;true&#x60;, the property is stored as a user action property | [optional] 
**StoreAsSessionProperty** | Pointer to **bool** | If &#x60;true&#x60;, the property is stored as a session property | [optional] 
**CleanupRule** | Pointer to **string** | The cleanup rule of the property.   Defines how to extract the data you need from a string value. Specify the [regular expression](https://www.dynatrace.com/support/help/shortlink/regex) for the data you need there. | [optional] 
**ServerSideRequestAttribute** | Pointer to **string** | The ID of the request attribute.   Only applicable when the **origin** is set to &#x60;SERVER_SIDE_REQUEST_ATTRIBUTE&#x60;. | [optional] 
**Name** | Pointer to **string** | The name of the reported value.   Only applicable when the **origin** is set to &#x60;API&#x60;. | [optional] 

## Methods

### NewMobileSessionProperty

`func NewMobileSessionProperty(key string, type_ string, origin string, ) *MobileSessionProperty`

NewMobileSessionProperty instantiates a new MobileSessionProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileSessionPropertyWithDefaults

`func NewMobileSessionPropertyWithDefaults() *MobileSessionProperty`

NewMobileSessionPropertyWithDefaults instantiates a new MobileSessionProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MobileSessionProperty) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MobileSessionProperty) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MobileSessionProperty) SetKey(v string)`

SetKey sets Key field to given value.


### GetDisplayName

`func (o *MobileSessionProperty) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MobileSessionProperty) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MobileSessionProperty) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MobileSessionProperty) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetType

`func (o *MobileSessionProperty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MobileSessionProperty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MobileSessionProperty) SetType(v string)`

SetType sets Type field to given value.


### GetOrigin

`func (o *MobileSessionProperty) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *MobileSessionProperty) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *MobileSessionProperty) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetAggregation

`func (o *MobileSessionProperty) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MobileSessionProperty) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MobileSessionProperty) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *MobileSessionProperty) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetStoreAsUserActionProperty

`func (o *MobileSessionProperty) GetStoreAsUserActionProperty() bool`

GetStoreAsUserActionProperty returns the StoreAsUserActionProperty field if non-nil, zero value otherwise.

### GetStoreAsUserActionPropertyOk

`func (o *MobileSessionProperty) GetStoreAsUserActionPropertyOk() (*bool, bool)`

GetStoreAsUserActionPropertyOk returns a tuple with the StoreAsUserActionProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreAsUserActionProperty

`func (o *MobileSessionProperty) SetStoreAsUserActionProperty(v bool)`

SetStoreAsUserActionProperty sets StoreAsUserActionProperty field to given value.

### HasStoreAsUserActionProperty

`func (o *MobileSessionProperty) HasStoreAsUserActionProperty() bool`

HasStoreAsUserActionProperty returns a boolean if a field has been set.

### GetStoreAsSessionProperty

`func (o *MobileSessionProperty) GetStoreAsSessionProperty() bool`

GetStoreAsSessionProperty returns the StoreAsSessionProperty field if non-nil, zero value otherwise.

### GetStoreAsSessionPropertyOk

`func (o *MobileSessionProperty) GetStoreAsSessionPropertyOk() (*bool, bool)`

GetStoreAsSessionPropertyOk returns a tuple with the StoreAsSessionProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreAsSessionProperty

`func (o *MobileSessionProperty) SetStoreAsSessionProperty(v bool)`

SetStoreAsSessionProperty sets StoreAsSessionProperty field to given value.

### HasStoreAsSessionProperty

`func (o *MobileSessionProperty) HasStoreAsSessionProperty() bool`

HasStoreAsSessionProperty returns a boolean if a field has been set.

### GetCleanupRule

`func (o *MobileSessionProperty) GetCleanupRule() string`

GetCleanupRule returns the CleanupRule field if non-nil, zero value otherwise.

### GetCleanupRuleOk

`func (o *MobileSessionProperty) GetCleanupRuleOk() (*string, bool)`

GetCleanupRuleOk returns a tuple with the CleanupRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupRule

`func (o *MobileSessionProperty) SetCleanupRule(v string)`

SetCleanupRule sets CleanupRule field to given value.

### HasCleanupRule

`func (o *MobileSessionProperty) HasCleanupRule() bool`

HasCleanupRule returns a boolean if a field has been set.

### GetServerSideRequestAttribute

`func (o *MobileSessionProperty) GetServerSideRequestAttribute() string`

GetServerSideRequestAttribute returns the ServerSideRequestAttribute field if non-nil, zero value otherwise.

### GetServerSideRequestAttributeOk

`func (o *MobileSessionProperty) GetServerSideRequestAttributeOk() (*string, bool)`

GetServerSideRequestAttributeOk returns a tuple with the ServerSideRequestAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSideRequestAttribute

`func (o *MobileSessionProperty) SetServerSideRequestAttribute(v string)`

SetServerSideRequestAttribute sets ServerSideRequestAttribute field to given value.

### HasServerSideRequestAttribute

`func (o *MobileSessionProperty) HasServerSideRequestAttribute() bool`

HasServerSideRequestAttribute returns a boolean if a field has been set.

### GetName

`func (o *MobileSessionProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobileSessionProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobileSessionProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MobileSessionProperty) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


