# Preferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateManagementEnabled** | Pointer to **bool** |  | [optional] 
**CertificateManagementPossible** | Pointer to **bool** |  | [optional] 
**SupportSendBilling** | **bool** | If true, usage and billing information will be reported. | 
**SuppressNonBillingRelevantData** | **bool** | If true, usage and billing information will NOT be reported. | 
**SupportSendClusterHealth** | **bool** | If true, Dynatrace cluster health will be reported. | 
**SupportSendEvents** | Pointer to **bool** | If true, Dynatrace cluster health and OneAgent events will be reported. | [optional] 
**SupportAllowRemoteAccess** | Pointer to **bool** | If true, audited remote-access to your Dynatrace configuration is allowed settings. | [optional] 
**RemoteAccessOnDemandOnly** | Pointer to **bool** | If true, audited access to your Dynatrace cluster is allowed by approved Dynatrace employees otherwise by privileged Dynatrace employees. | [optional] 
**CommunityCreateUser** | Pointer to **bool** | If true, each new user will be set up as a Dynatrace Community user upon first login. | [optional] 
**CommunityExternalSearch** | Pointer to **bool** | If true, you can search Dynatrace Community and Dynatrace Help when using the product search. | [optional] 
**BackupEnabled** | Pointer to **bool** | Deprecated. Use the backup configuration endpoint. If true, backups are enabled. | [optional] [readonly] 
**RuxitMonitorsRuxit** | Pointer to **bool** | If true, Dynatrace OneAgent monitors Dynatrace. | [optional] 
**TelemetrySharing** | Pointer to **bool** |  | [optional] 
**HelpChatEnabled** | Pointer to **bool** | If true, Dynatrace ONE live chat is enabled. | [optional] 
**ReadOnlyRemoteAccessAllowed** | Pointer to **bool** | If true, audited, read-only remote access to your Dynatrace configuration settings is allowed. | [optional] 

## Methods

### NewPreferences

`func NewPreferences(supportSendBilling bool, suppressNonBillingRelevantData bool, supportSendClusterHealth bool, ) *Preferences`

NewPreferences instantiates a new Preferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreferencesWithDefaults

`func NewPreferencesWithDefaults() *Preferences`

NewPreferencesWithDefaults instantiates a new Preferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateManagementEnabled

`func (o *Preferences) GetCertificateManagementEnabled() bool`

GetCertificateManagementEnabled returns the CertificateManagementEnabled field if non-nil, zero value otherwise.

### GetCertificateManagementEnabledOk

`func (o *Preferences) GetCertificateManagementEnabledOk() (*bool, bool)`

GetCertificateManagementEnabledOk returns a tuple with the CertificateManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateManagementEnabled

`func (o *Preferences) SetCertificateManagementEnabled(v bool)`

SetCertificateManagementEnabled sets CertificateManagementEnabled field to given value.

### HasCertificateManagementEnabled

`func (o *Preferences) HasCertificateManagementEnabled() bool`

HasCertificateManagementEnabled returns a boolean if a field has been set.

### GetCertificateManagementPossible

`func (o *Preferences) GetCertificateManagementPossible() bool`

GetCertificateManagementPossible returns the CertificateManagementPossible field if non-nil, zero value otherwise.

### GetCertificateManagementPossibleOk

`func (o *Preferences) GetCertificateManagementPossibleOk() (*bool, bool)`

GetCertificateManagementPossibleOk returns a tuple with the CertificateManagementPossible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateManagementPossible

`func (o *Preferences) SetCertificateManagementPossible(v bool)`

SetCertificateManagementPossible sets CertificateManagementPossible field to given value.

### HasCertificateManagementPossible

`func (o *Preferences) HasCertificateManagementPossible() bool`

HasCertificateManagementPossible returns a boolean if a field has been set.

### GetSupportSendBilling

`func (o *Preferences) GetSupportSendBilling() bool`

GetSupportSendBilling returns the SupportSendBilling field if non-nil, zero value otherwise.

### GetSupportSendBillingOk

`func (o *Preferences) GetSupportSendBillingOk() (*bool, bool)`

GetSupportSendBillingOk returns a tuple with the SupportSendBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSendBilling

`func (o *Preferences) SetSupportSendBilling(v bool)`

SetSupportSendBilling sets SupportSendBilling field to given value.


### GetSuppressNonBillingRelevantData

`func (o *Preferences) GetSuppressNonBillingRelevantData() bool`

GetSuppressNonBillingRelevantData returns the SuppressNonBillingRelevantData field if non-nil, zero value otherwise.

### GetSuppressNonBillingRelevantDataOk

`func (o *Preferences) GetSuppressNonBillingRelevantDataOk() (*bool, bool)`

GetSuppressNonBillingRelevantDataOk returns a tuple with the SuppressNonBillingRelevantData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressNonBillingRelevantData

`func (o *Preferences) SetSuppressNonBillingRelevantData(v bool)`

SetSuppressNonBillingRelevantData sets SuppressNonBillingRelevantData field to given value.


### GetSupportSendClusterHealth

`func (o *Preferences) GetSupportSendClusterHealth() bool`

GetSupportSendClusterHealth returns the SupportSendClusterHealth field if non-nil, zero value otherwise.

### GetSupportSendClusterHealthOk

`func (o *Preferences) GetSupportSendClusterHealthOk() (*bool, bool)`

GetSupportSendClusterHealthOk returns a tuple with the SupportSendClusterHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSendClusterHealth

`func (o *Preferences) SetSupportSendClusterHealth(v bool)`

SetSupportSendClusterHealth sets SupportSendClusterHealth field to given value.


### GetSupportSendEvents

`func (o *Preferences) GetSupportSendEvents() bool`

GetSupportSendEvents returns the SupportSendEvents field if non-nil, zero value otherwise.

### GetSupportSendEventsOk

`func (o *Preferences) GetSupportSendEventsOk() (*bool, bool)`

GetSupportSendEventsOk returns a tuple with the SupportSendEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSendEvents

`func (o *Preferences) SetSupportSendEvents(v bool)`

SetSupportSendEvents sets SupportSendEvents field to given value.

### HasSupportSendEvents

`func (o *Preferences) HasSupportSendEvents() bool`

HasSupportSendEvents returns a boolean if a field has been set.

### GetSupportAllowRemoteAccess

`func (o *Preferences) GetSupportAllowRemoteAccess() bool`

GetSupportAllowRemoteAccess returns the SupportAllowRemoteAccess field if non-nil, zero value otherwise.

### GetSupportAllowRemoteAccessOk

`func (o *Preferences) GetSupportAllowRemoteAccessOk() (*bool, bool)`

GetSupportAllowRemoteAccessOk returns a tuple with the SupportAllowRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAllowRemoteAccess

`func (o *Preferences) SetSupportAllowRemoteAccess(v bool)`

SetSupportAllowRemoteAccess sets SupportAllowRemoteAccess field to given value.

### HasSupportAllowRemoteAccess

`func (o *Preferences) HasSupportAllowRemoteAccess() bool`

HasSupportAllowRemoteAccess returns a boolean if a field has been set.

### GetRemoteAccessOnDemandOnly

`func (o *Preferences) GetRemoteAccessOnDemandOnly() bool`

GetRemoteAccessOnDemandOnly returns the RemoteAccessOnDemandOnly field if non-nil, zero value otherwise.

### GetRemoteAccessOnDemandOnlyOk

`func (o *Preferences) GetRemoteAccessOnDemandOnlyOk() (*bool, bool)`

GetRemoteAccessOnDemandOnlyOk returns a tuple with the RemoteAccessOnDemandOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessOnDemandOnly

`func (o *Preferences) SetRemoteAccessOnDemandOnly(v bool)`

SetRemoteAccessOnDemandOnly sets RemoteAccessOnDemandOnly field to given value.

### HasRemoteAccessOnDemandOnly

`func (o *Preferences) HasRemoteAccessOnDemandOnly() bool`

HasRemoteAccessOnDemandOnly returns a boolean if a field has been set.

### GetCommunityCreateUser

`func (o *Preferences) GetCommunityCreateUser() bool`

GetCommunityCreateUser returns the CommunityCreateUser field if non-nil, zero value otherwise.

### GetCommunityCreateUserOk

`func (o *Preferences) GetCommunityCreateUserOk() (*bool, bool)`

GetCommunityCreateUserOk returns a tuple with the CommunityCreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityCreateUser

`func (o *Preferences) SetCommunityCreateUser(v bool)`

SetCommunityCreateUser sets CommunityCreateUser field to given value.

### HasCommunityCreateUser

`func (o *Preferences) HasCommunityCreateUser() bool`

HasCommunityCreateUser returns a boolean if a field has been set.

### GetCommunityExternalSearch

`func (o *Preferences) GetCommunityExternalSearch() bool`

GetCommunityExternalSearch returns the CommunityExternalSearch field if non-nil, zero value otherwise.

### GetCommunityExternalSearchOk

`func (o *Preferences) GetCommunityExternalSearchOk() (*bool, bool)`

GetCommunityExternalSearchOk returns a tuple with the CommunityExternalSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityExternalSearch

`func (o *Preferences) SetCommunityExternalSearch(v bool)`

SetCommunityExternalSearch sets CommunityExternalSearch field to given value.

### HasCommunityExternalSearch

`func (o *Preferences) HasCommunityExternalSearch() bool`

HasCommunityExternalSearch returns a boolean if a field has been set.

### GetBackupEnabled

`func (o *Preferences) GetBackupEnabled() bool`

GetBackupEnabled returns the BackupEnabled field if non-nil, zero value otherwise.

### GetBackupEnabledOk

`func (o *Preferences) GetBackupEnabledOk() (*bool, bool)`

GetBackupEnabledOk returns a tuple with the BackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEnabled

`func (o *Preferences) SetBackupEnabled(v bool)`

SetBackupEnabled sets BackupEnabled field to given value.

### HasBackupEnabled

`func (o *Preferences) HasBackupEnabled() bool`

HasBackupEnabled returns a boolean if a field has been set.

### GetRuxitMonitorsRuxit

`func (o *Preferences) GetRuxitMonitorsRuxit() bool`

GetRuxitMonitorsRuxit returns the RuxitMonitorsRuxit field if non-nil, zero value otherwise.

### GetRuxitMonitorsRuxitOk

`func (o *Preferences) GetRuxitMonitorsRuxitOk() (*bool, bool)`

GetRuxitMonitorsRuxitOk returns a tuple with the RuxitMonitorsRuxit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuxitMonitorsRuxit

`func (o *Preferences) SetRuxitMonitorsRuxit(v bool)`

SetRuxitMonitorsRuxit sets RuxitMonitorsRuxit field to given value.

### HasRuxitMonitorsRuxit

`func (o *Preferences) HasRuxitMonitorsRuxit() bool`

HasRuxitMonitorsRuxit returns a boolean if a field has been set.

### GetTelemetrySharing

`func (o *Preferences) GetTelemetrySharing() bool`

GetTelemetrySharing returns the TelemetrySharing field if non-nil, zero value otherwise.

### GetTelemetrySharingOk

`func (o *Preferences) GetTelemetrySharingOk() (*bool, bool)`

GetTelemetrySharingOk returns a tuple with the TelemetrySharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetrySharing

`func (o *Preferences) SetTelemetrySharing(v bool)`

SetTelemetrySharing sets TelemetrySharing field to given value.

### HasTelemetrySharing

`func (o *Preferences) HasTelemetrySharing() bool`

HasTelemetrySharing returns a boolean if a field has been set.

### GetHelpChatEnabled

`func (o *Preferences) GetHelpChatEnabled() bool`

GetHelpChatEnabled returns the HelpChatEnabled field if non-nil, zero value otherwise.

### GetHelpChatEnabledOk

`func (o *Preferences) GetHelpChatEnabledOk() (*bool, bool)`

GetHelpChatEnabledOk returns a tuple with the HelpChatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpChatEnabled

`func (o *Preferences) SetHelpChatEnabled(v bool)`

SetHelpChatEnabled sets HelpChatEnabled field to given value.

### HasHelpChatEnabled

`func (o *Preferences) HasHelpChatEnabled() bool`

HasHelpChatEnabled returns a boolean if a field has been set.

### GetReadOnlyRemoteAccessAllowed

`func (o *Preferences) GetReadOnlyRemoteAccessAllowed() bool`

GetReadOnlyRemoteAccessAllowed returns the ReadOnlyRemoteAccessAllowed field if non-nil, zero value otherwise.

### GetReadOnlyRemoteAccessAllowedOk

`func (o *Preferences) GetReadOnlyRemoteAccessAllowedOk() (*bool, bool)`

GetReadOnlyRemoteAccessAllowedOk returns a tuple with the ReadOnlyRemoteAccessAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyRemoteAccessAllowed

`func (o *Preferences) SetReadOnlyRemoteAccessAllowed(v bool)`

SetReadOnlyRemoteAccessAllowed sets ReadOnlyRemoteAccessAllowed field to given value.

### HasReadOnlyRemoteAccessAllowed

`func (o *Preferences) HasReadOnlyRemoteAccessAllowed() bool`

HasReadOnlyRemoteAccessAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


