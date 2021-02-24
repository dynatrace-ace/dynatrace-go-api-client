# UserActionNamingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Placeholders** | Pointer to [**[]UserActionNamingPlaceholder**](UserActionNamingPlaceholder.md) | User action placeholders. | [optional] 
**LoadActionNamingRules** | Pointer to [**[]UserActionNamingRule**](UserActionNamingRule.md) | User action naming rules for loading actions. | [optional] 
**XhrActionNamingRules** | Pointer to [**[]UserActionNamingRule**](UserActionNamingRule.md) | User action naming rules for xhr actions. | [optional] 
**CustomActionNamingRules** | Pointer to [**[]UserActionNamingRule**](UserActionNamingRule.md) | User action naming rules for custom actions. | [optional] 
**IgnoreCase** | Pointer to **bool** | Case insensitive naming. | [optional] 
**UseFirstDetectedLoadAction** | Pointer to **bool** | First load action found under an XHR action should be used when true. Else the deepest one under the xhr action is used | [optional] 
**SplitUserActionsByDomain** | Pointer to **bool** | Deactivate this setting if different domains should not result in separate user actions. | [optional] 

## Methods

### NewUserActionNamingSettings

`func NewUserActionNamingSettings() *UserActionNamingSettings`

NewUserActionNamingSettings instantiates a new UserActionNamingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActionNamingSettingsWithDefaults

`func NewUserActionNamingSettingsWithDefaults() *UserActionNamingSettings`

NewUserActionNamingSettingsWithDefaults instantiates a new UserActionNamingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaceholders

`func (o *UserActionNamingSettings) GetPlaceholders() []UserActionNamingPlaceholder`

GetPlaceholders returns the Placeholders field if non-nil, zero value otherwise.

### GetPlaceholdersOk

`func (o *UserActionNamingSettings) GetPlaceholdersOk() (*[]UserActionNamingPlaceholder, bool)`

GetPlaceholdersOk returns a tuple with the Placeholders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholders

`func (o *UserActionNamingSettings) SetPlaceholders(v []UserActionNamingPlaceholder)`

SetPlaceholders sets Placeholders field to given value.

### HasPlaceholders

`func (o *UserActionNamingSettings) HasPlaceholders() bool`

HasPlaceholders returns a boolean if a field has been set.

### GetLoadActionNamingRules

`func (o *UserActionNamingSettings) GetLoadActionNamingRules() []UserActionNamingRule`

GetLoadActionNamingRules returns the LoadActionNamingRules field if non-nil, zero value otherwise.

### GetLoadActionNamingRulesOk

`func (o *UserActionNamingSettings) GetLoadActionNamingRulesOk() (*[]UserActionNamingRule, bool)`

GetLoadActionNamingRulesOk returns a tuple with the LoadActionNamingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadActionNamingRules

`func (o *UserActionNamingSettings) SetLoadActionNamingRules(v []UserActionNamingRule)`

SetLoadActionNamingRules sets LoadActionNamingRules field to given value.

### HasLoadActionNamingRules

`func (o *UserActionNamingSettings) HasLoadActionNamingRules() bool`

HasLoadActionNamingRules returns a boolean if a field has been set.

### GetXhrActionNamingRules

`func (o *UserActionNamingSettings) GetXhrActionNamingRules() []UserActionNamingRule`

GetXhrActionNamingRules returns the XhrActionNamingRules field if non-nil, zero value otherwise.

### GetXhrActionNamingRulesOk

`func (o *UserActionNamingSettings) GetXhrActionNamingRulesOk() (*[]UserActionNamingRule, bool)`

GetXhrActionNamingRulesOk returns a tuple with the XhrActionNamingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXhrActionNamingRules

`func (o *UserActionNamingSettings) SetXhrActionNamingRules(v []UserActionNamingRule)`

SetXhrActionNamingRules sets XhrActionNamingRules field to given value.

### HasXhrActionNamingRules

`func (o *UserActionNamingSettings) HasXhrActionNamingRules() bool`

HasXhrActionNamingRules returns a boolean if a field has been set.

### GetCustomActionNamingRules

`func (o *UserActionNamingSettings) GetCustomActionNamingRules() []UserActionNamingRule`

GetCustomActionNamingRules returns the CustomActionNamingRules field if non-nil, zero value otherwise.

### GetCustomActionNamingRulesOk

`func (o *UserActionNamingSettings) GetCustomActionNamingRulesOk() (*[]UserActionNamingRule, bool)`

GetCustomActionNamingRulesOk returns a tuple with the CustomActionNamingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomActionNamingRules

`func (o *UserActionNamingSettings) SetCustomActionNamingRules(v []UserActionNamingRule)`

SetCustomActionNamingRules sets CustomActionNamingRules field to given value.

### HasCustomActionNamingRules

`func (o *UserActionNamingSettings) HasCustomActionNamingRules() bool`

HasCustomActionNamingRules returns a boolean if a field has been set.

### GetIgnoreCase

`func (o *UserActionNamingSettings) GetIgnoreCase() bool`

GetIgnoreCase returns the IgnoreCase field if non-nil, zero value otherwise.

### GetIgnoreCaseOk

`func (o *UserActionNamingSettings) GetIgnoreCaseOk() (*bool, bool)`

GetIgnoreCaseOk returns a tuple with the IgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCase

`func (o *UserActionNamingSettings) SetIgnoreCase(v bool)`

SetIgnoreCase sets IgnoreCase field to given value.

### HasIgnoreCase

`func (o *UserActionNamingSettings) HasIgnoreCase() bool`

HasIgnoreCase returns a boolean if a field has been set.

### GetUseFirstDetectedLoadAction

`func (o *UserActionNamingSettings) GetUseFirstDetectedLoadAction() bool`

GetUseFirstDetectedLoadAction returns the UseFirstDetectedLoadAction field if non-nil, zero value otherwise.

### GetUseFirstDetectedLoadActionOk

`func (o *UserActionNamingSettings) GetUseFirstDetectedLoadActionOk() (*bool, bool)`

GetUseFirstDetectedLoadActionOk returns a tuple with the UseFirstDetectedLoadAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFirstDetectedLoadAction

`func (o *UserActionNamingSettings) SetUseFirstDetectedLoadAction(v bool)`

SetUseFirstDetectedLoadAction sets UseFirstDetectedLoadAction field to given value.

### HasUseFirstDetectedLoadAction

`func (o *UserActionNamingSettings) HasUseFirstDetectedLoadAction() bool`

HasUseFirstDetectedLoadAction returns a boolean if a field has been set.

### GetSplitUserActionsByDomain

`func (o *UserActionNamingSettings) GetSplitUserActionsByDomain() bool`

GetSplitUserActionsByDomain returns the SplitUserActionsByDomain field if non-nil, zero value otherwise.

### GetSplitUserActionsByDomainOk

`func (o *UserActionNamingSettings) GetSplitUserActionsByDomainOk() (*bool, bool)`

GetSplitUserActionsByDomainOk returns a tuple with the SplitUserActionsByDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitUserActionsByDomain

`func (o *UserActionNamingSettings) SetSplitUserActionsByDomain(v bool)`

SetSplitUserActionsByDomain sets SplitUserActionsByDomain field to given value.

### HasSplitUserActionsByDomain

`func (o *UserActionNamingSettings) HasSplitUserActionsByDomain() bool`

HasSplitUserActionsByDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


