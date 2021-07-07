# ApiTokenCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersonalAccessToken** | Pointer to **bool** | The token is a personal access token (&#x60;true&#x60;) or an API token (&#x60;false&#x60;).    Personal access tokens are tied to the permissions of their owner. | [optional] 
**ExpirationDate** | Pointer to **string** | The expiration date of the token.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years    | [optional] 
**Scopes** | **[]string** | A list of the scopes to be assigned to the token.  * &#x60;InstallerDownload&#x60;: PaaS integration - Installer download.  * &#x60;DataExport&#x60;: Access problem and event feed, metrics, and topology.  * &#x60;PluginUpload&#x60;: Upload Extension.  * &#x60;SupportAlert&#x60;: PaaS integration - Support alert.  * &#x60;DcrumIntegration&#x60;: Dynatrace module integration - NAM.  * &#x60;AdvancedSyntheticIntegration&#x60;: Dynatrace module integration - Synthetic Classic.  * &#x60;ExternalSyntheticIntegration&#x60;: Create and read synthetic monitors, locations, and nodes.  * &#x60;AppMonIntegration&#x60;: Dynatrace module integration - AppMon.  * &#x60;RumBrowserExtension&#x60;: RUM Browser Extension.  * &#x60;LogExport&#x60;: Read logs.  * &#x60;ReadConfig&#x60;: Read configuration.  * &#x60;WriteConfig&#x60;: Write configuration.  * &#x60;DTAQLAccess&#x60;: User sessions.  * &#x60;UserSessionAnonymization&#x60;: Anonymize user session data for data privacy reasons.  * &#x60;DataPrivacy&#x60;: Change data privacy settings.  * &#x60;CaptureRequestData&#x60;: Capture request data.  * &#x60;Davis&#x60;: Dynatrace module integration - Davis.  * &#x60;DssFileManagement&#x60;: Mobile symbolication file management.  * &#x60;RumJavaScriptTagManagement&#x60;: Real user monitoring JavaScript tag management.  * &#x60;TenantTokenManagement&#x60;: Token management.  * &#x60;ActiveGateCertManagement&#x60;: ActiveGate certificate management.  * &#x60;RestRequestForwarding&#x60;: Fetch data from a remote environment.  * &#x60;ReadSyntheticData&#x60;: Read synthetic monitors, locations, and nodes.  * &#x60;DataImport&#x60;: Data ingest, e.g.: metrics and events.  * &#x60;auditLogs.read&#x60;: Read audit logs.  * &#x60;metrics.read&#x60;: Read metrics.  * &#x60;metrics.write&#x60;: Write metrics.  * &#x60;entities.read&#x60;: Read entities.  * &#x60;entities.write&#x60;: Write entities.  * &#x60;problems.read&#x60;: Read problems.  * &#x60;problems.write&#x60;: Write problems.  * &#x60;networkZones.read&#x60;: Read network zones.  * &#x60;networkZones.write&#x60;: Write network zones.  * &#x60;activeGates.read&#x60;: Read ActiveGates.  * &#x60;activeGates.write&#x60;: Write ActiveGates.  * &#x60;credentialVault.read&#x60;: Read credential vault entries.  * &#x60;credentialVault.write&#x60;: Write credential vault entries.  * &#x60;extensions.read&#x60;: Read extensions.  * &#x60;extensions.write&#x60;: Write extensions.  * &#x60;extensionConfigurations.read&#x60;: Read extension monitoring configurations.  * &#x60;extensionConfigurations.write&#x60;: Write extension monitoring configurations.  * &#x60;extensionEnvironment.read&#x60;: Read extension environment configurations.  * &#x60;extensionEnvironment.write&#x60;: Write extension environment configurations.  * &#x60;metrics.ingest&#x60;: Ingest metrics.  * &#x60;securityProblems.read&#x60;: Read security problems.  * &#x60;securityProblems.write&#x60;: Write security problems.  * &#x60;syntheticLocations.read&#x60;: Read synthetic locations.  * &#x60;syntheticLocations.write&#x60;: Write synthetic locations.  * &#x60;settings.read&#x60;: Read settings.  * &#x60;settings.write&#x60;: Write settings.  * &#x60;tenantTokenRotation.write&#x60;: Tenant token rotation.  * &#x60;slo.read&#x60;: Read SLO.  * &#x60;slo.write&#x60;: Write SLO.  * &#x60;releases.read&#x60;: Read releases.  * &#x60;apiTokens.read&#x60;: Read API tokens.  * &#x60;apiTokens.write&#x60;: Write API tokens.  * &#x60;logs.read&#x60;: Read logs.  * &#x60;logs.ingest&#x60;: Ingest logs.   | 
**Name** | **string** | The name of the token. | 

## Methods

### NewApiTokenCreate

`func NewApiTokenCreate(scopes []string, name string, ) *ApiTokenCreate`

NewApiTokenCreate instantiates a new ApiTokenCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokenCreateWithDefaults

`func NewApiTokenCreateWithDefaults() *ApiTokenCreate`

NewApiTokenCreateWithDefaults instantiates a new ApiTokenCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersonalAccessToken

`func (o *ApiTokenCreate) GetPersonalAccessToken() bool`

GetPersonalAccessToken returns the PersonalAccessToken field if non-nil, zero value otherwise.

### GetPersonalAccessTokenOk

`func (o *ApiTokenCreate) GetPersonalAccessTokenOk() (*bool, bool)`

GetPersonalAccessTokenOk returns a tuple with the PersonalAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalAccessToken

`func (o *ApiTokenCreate) SetPersonalAccessToken(v bool)`

SetPersonalAccessToken sets PersonalAccessToken field to given value.

### HasPersonalAccessToken

`func (o *ApiTokenCreate) HasPersonalAccessToken() bool`

HasPersonalAccessToken returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ApiTokenCreate) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ApiTokenCreate) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ApiTokenCreate) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ApiTokenCreate) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetScopes

`func (o *ApiTokenCreate) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApiTokenCreate) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApiTokenCreate) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetName

`func (o *ApiTokenCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiTokenCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiTokenCreate) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


