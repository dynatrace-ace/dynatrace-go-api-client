# MobileSessionPropertyUpdate

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
**Name** | Pointer to **string** | The name of the reported value.   Only applicable when the **origin** is set to &#x60;API&#x60;. | [optional] 

## Methods

### NewMobileSessionPropertyUpdate

`func NewMobileSessionPropertyUpdate(type_ string, origin string, ) *MobileSessionPropertyUpdate`

NewMobileSessionPropertyUpdate instantiates a new MobileSessionPropertyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileSessionPropertyUpdateWithDefaults

`func NewMobileSessionPropertyUpdateWithDefaults() *MobileSessionPropertyUpdate`

NewMobileSessionPropertyUpdateWithDefaults instantiates a new MobileSessionPropertyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MobileSessionPropertyUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MobileSessionPropertyUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MobileSessionPropertyUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MobileSessionPropertyUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetType

`func (o *MobileSessionPropertyUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MobileSessionPropertyUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MobileSessionPropertyUpdate) SetType(v string)`

SetType sets Type field to given value.


### GetOrigin

`func (o *MobileSessionPropertyUpdate) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *MobileSessionPropertyUpdate) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *MobileSessionPropertyUpdate) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetAggregation

`func (o *MobileSessionPropertyUpdate) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MobileSessionPropertyUpdate) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MobileSessionPropertyUpdate) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *MobileSessionPropertyUpdate) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetStoreAsUserActionProperty

`func (o *MobileSessionPropertyUpdate) GetStoreAsUserActionProperty() bool`

GetStoreAsUserActionProperty returns the StoreAsUserActionProperty field if non-nil, zero value otherwise.

### GetStoreAsUserActionPropertyOk

`func (o *MobileSessionPropertyUpdate) GetStoreAsUserActionPropertyOk() (*bool, bool)`

GetStoreAsUserActionPropertyOk returns a tuple with the StoreAsUserActionProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreAsUserActionProperty

`func (o *MobileSessionPropertyUpdate) SetStoreAsUserActionProperty(v bool)`

SetStoreAsUserActionProperty sets StoreAsUserActionProperty field to given value.

### HasStoreAsUserActionProperty

`func (o *MobileSessionPropertyUpdate) HasStoreAsUserActionProperty() bool`

HasStoreAsUserActionProperty returns a boolean if a field has been set.

### GetStoreAsSessionProperty

`func (o *MobileSessionPropertyUpdate) GetStoreAsSessionProperty() bool`

GetStoreAsSessionProperty returns the StoreAsSessionProperty field if non-nil, zero value otherwise.

### GetStoreAsSessionPropertyOk

`func (o *MobileSessionPropertyUpdate) GetStoreAsSessionPropertyOk() (*bool, bool)`

GetStoreAsSessionPropertyOk returns a tuple with the StoreAsSessionProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreAsSessionProperty

`func (o *MobileSessionPropertyUpdate) SetStoreAsSessionProperty(v bool)`

SetStoreAsSessionProperty sets StoreAsSessionProperty field to given value.

### HasStoreAsSessionProperty

`func (o *MobileSessionPropertyUpdate) HasStoreAsSessionProperty() bool`

HasStoreAsSessionProperty returns a boolean if a field has been set.

### GetCleanupRule

`func (o *MobileSessionPropertyUpdate) GetCleanupRule() string`

GetCleanupRule returns the CleanupRule field if non-nil, zero value otherwise.

### GetCleanupRuleOk

`func (o *MobileSessionPropertyUpdate) GetCleanupRuleOk() (*string, bool)`

GetCleanupRuleOk returns a tuple with the CleanupRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupRule

`func (o *MobileSessionPropertyUpdate) SetCleanupRule(v string)`

SetCleanupRule sets CleanupRule field to given value.

### HasCleanupRule

`func (o *MobileSessionPropertyUpdate) HasCleanupRule() bool`

HasCleanupRule returns a boolean if a field has been set.

### GetServerSideRequestAttribute

`func (o *MobileSessionPropertyUpdate) GetServerSideRequestAttribute() string`

GetServerSideRequestAttribute returns the ServerSideRequestAttribute field if non-nil, zero value otherwise.

### GetServerSideRequestAttributeOk

`func (o *MobileSessionPropertyUpdate) GetServerSideRequestAttributeOk() (*string, bool)`

GetServerSideRequestAttributeOk returns a tuple with the ServerSideRequestAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSideRequestAttribute

`func (o *MobileSessionPropertyUpdate) SetServerSideRequestAttribute(v string)`

SetServerSideRequestAttribute sets ServerSideRequestAttribute field to given value.

### HasServerSideRequestAttribute

`func (o *MobileSessionPropertyUpdate) HasServerSideRequestAttribute() bool`

HasServerSideRequestAttribute returns a boolean if a field has been set.

### GetName

`func (o *MobileSessionPropertyUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobileSessionPropertyUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobileSessionPropertyUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MobileSessionPropertyUpdate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


