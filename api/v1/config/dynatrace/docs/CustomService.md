# CustomService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the custom service. | [optional] 
**Name** | **string** | The name of the custom service, displayed in the UI. | 
**Order** | Pointer to **string** | The order string. Sorting custom services alphabetically by their order string determines their relative ordering.  Typically this is managed by Dynatrace internally and will not be present in GET responses. | [optional] 
**Enabled** | **bool** | Custom service enabled/disabled. | 
**Rules** | [**[]DetectionRule**](DetectionRule.md) | The list of rules defining the custom service. | 
**QueueEntryPoint** | **bool** | The queue entry point flag.   Set to &#x60;true&#x60; for custom messaging services. | 
**QueueEntryPointType** | Pointer to **string** | The queue entry point type.. | [optional] 
**ProcessGroups** | Pointer to **[]string** | The list of process groups the custom service should belong to. | [optional] 

## Methods

### NewCustomService

`func NewCustomService(name string, enabled bool, rules []DetectionRule, queueEntryPoint bool, ) *CustomService`

NewCustomService instantiates a new CustomService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomServiceWithDefaults

`func NewCustomServiceWithDefaults() *CustomService`

NewCustomServiceWithDefaults instantiates a new CustomService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CustomService) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomService) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomService) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomService) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *CustomService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CustomService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomService) SetName(v string)`

SetName sets Name field to given value.


### GetOrder

`func (o *CustomService) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CustomService) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CustomService) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CustomService) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetEnabled

`func (o *CustomService) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomService) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomService) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRules

`func (o *CustomService) GetRules() []DetectionRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CustomService) GetRulesOk() (*[]DetectionRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CustomService) SetRules(v []DetectionRule)`

SetRules sets Rules field to given value.


### GetQueueEntryPoint

`func (o *CustomService) GetQueueEntryPoint() bool`

GetQueueEntryPoint returns the QueueEntryPoint field if non-nil, zero value otherwise.

### GetQueueEntryPointOk

`func (o *CustomService) GetQueueEntryPointOk() (*bool, bool)`

GetQueueEntryPointOk returns a tuple with the QueueEntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueEntryPoint

`func (o *CustomService) SetQueueEntryPoint(v bool)`

SetQueueEntryPoint sets QueueEntryPoint field to given value.


### GetQueueEntryPointType

`func (o *CustomService) GetQueueEntryPointType() string`

GetQueueEntryPointType returns the QueueEntryPointType field if non-nil, zero value otherwise.

### GetQueueEntryPointTypeOk

`func (o *CustomService) GetQueueEntryPointTypeOk() (*string, bool)`

GetQueueEntryPointTypeOk returns a tuple with the QueueEntryPointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueEntryPointType

`func (o *CustomService) SetQueueEntryPointType(v string)`

SetQueueEntryPointType sets QueueEntryPointType field to given value.

### HasQueueEntryPointType

`func (o *CustomService) HasQueueEntryPointType() bool`

HasQueueEntryPointType returns a boolean if a field has been set.

### GetProcessGroups

`func (o *CustomService) GetProcessGroups() []string`

GetProcessGroups returns the ProcessGroups field if non-nil, zero value otherwise.

### GetProcessGroupsOk

`func (o *CustomService) GetProcessGroupsOk() (*[]string, bool)`

GetProcessGroupsOk returns a tuple with the ProcessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroups

`func (o *CustomService) SetProcessGroups(v []string)`

SetProcessGroups sets ProcessGroups field to given value.

### HasProcessGroups

`func (o *CustomService) HasProcessGroups() bool`

HasProcessGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


