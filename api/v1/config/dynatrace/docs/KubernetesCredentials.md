# KubernetesCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the given credentials configuration. | [optional] [readonly] 
**Active** | Pointer to **bool** | The monitoring is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;) for given credentials configuration.   If not set on creation, the &#x60;true&#x60; value is used.   If the field is omitted during an update, the old value remains unaffected. | [optional] 
**EndpointStatus** | Pointer to **string** | The status of the configured endpoint.  ASSIGNED: The credentials are assigned to an ActiveGate which is responsible for processing. UNASSIGNED: The credentials are not yet assigned to an ActiveGate so there is currently no processing. DISABLED: The credentials have been disabled by the user. FASTCHECK_AUTH_ERROR: The credentials are invalid. FASTCHECK_TLS_ERROR: The endpoint TLS certificate is invalid. FASTCHECK_NO_RESPONSE: The endpoint did not return a result until the timeout was reached. FASTCHECK_INVALID_ENDPOINT: The endpoint URL was invalid. FASTCHECK_AUTH_LOCKED: The credentials seem to be locked. UNKNOWN: An unknown error occured.  | [optional] [readonly] 
**EndpointStatusInfo** | Pointer to **string** | The detailed status info of the configured endpoint. | [optional] [readonly] 
**Label** | **string** | The name of the Kubernetes credentials configuration.   Allowed characters are letters, numbers, whitespaces, and the following characters: &#x60;.+-_&#x60;. Leading or trailing whitespace is not allowed. | 
**EndpointUrl** | **string** | The URL of the Kubernetes API server.   It must be unique within a Dynatrace environment.   The URL must valid according to RFC 2396. Leading or trailing whitespaces are not allowed. | 
**AuthToken** | Pointer to **string** | The service account bearer token for the Kubernetes API server.   Submit your token on creation or update of the configuration. For security reasons, GET requests return this field as &#x60;null&#x60;.   If the field is omitted during an update, the old value remains unaffected. | [optional] 
**EventsIntegrationEnabled** | Pointer to **bool** | The monitoring of events is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;) for the Kubernetes cluster. Event monitoring depends on the active state of this configuration to be true.   If not set on creation, the &#x60;false&#x60; value is used.   If the field is omitted during an update, the old value remains unaffected. | [optional] 
**WorkloadIntegrationEnabled** | Pointer to **bool** | Workload and cloud application processing is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;) for the Kubernetes cluster. If the field is omitted during an update, the old value remains unaffected. | [optional] 
**EventsFieldSelectors** | Pointer to [**[]KubernetesEventPattern**](KubernetesEventPattern.md) | Kubernetes event filters based on field-selectors. If set to &#x60;null&#x60; on creation, no events field selectors are subscribed. If set to &#x60;null&#x60; on update, no change of stored events field selectors is applied. Set an empty list to clear all events field selectors. | [optional] 
**CertificateCheckEnabled** | Pointer to **bool** | The check of SSL certificates is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;) for the Kubernetes cluster.   If not set on creation, the &#x60;true&#x60; value is used.   If the field is omitted during an update, the old value remains unaffected. | [optional] 

## Methods

### NewKubernetesCredentials

`func NewKubernetesCredentials(label string, endpointUrl string, ) *KubernetesCredentials`

NewKubernetesCredentials instantiates a new KubernetesCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCredentialsWithDefaults

`func NewKubernetesCredentialsWithDefaults() *KubernetesCredentials`

NewKubernetesCredentialsWithDefaults instantiates a new KubernetesCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *KubernetesCredentials) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubernetesCredentials) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubernetesCredentials) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubernetesCredentials) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *KubernetesCredentials) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesCredentials) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesCredentials) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesCredentials) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActive

`func (o *KubernetesCredentials) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *KubernetesCredentials) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *KubernetesCredentials) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *KubernetesCredentials) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEndpointStatus

`func (o *KubernetesCredentials) GetEndpointStatus() string`

GetEndpointStatus returns the EndpointStatus field if non-nil, zero value otherwise.

### GetEndpointStatusOk

`func (o *KubernetesCredentials) GetEndpointStatusOk() (*string, bool)`

GetEndpointStatusOk returns a tuple with the EndpointStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointStatus

`func (o *KubernetesCredentials) SetEndpointStatus(v string)`

SetEndpointStatus sets EndpointStatus field to given value.

### HasEndpointStatus

`func (o *KubernetesCredentials) HasEndpointStatus() bool`

HasEndpointStatus returns a boolean if a field has been set.

### GetEndpointStatusInfo

`func (o *KubernetesCredentials) GetEndpointStatusInfo() string`

GetEndpointStatusInfo returns the EndpointStatusInfo field if non-nil, zero value otherwise.

### GetEndpointStatusInfoOk

`func (o *KubernetesCredentials) GetEndpointStatusInfoOk() (*string, bool)`

GetEndpointStatusInfoOk returns a tuple with the EndpointStatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointStatusInfo

`func (o *KubernetesCredentials) SetEndpointStatusInfo(v string)`

SetEndpointStatusInfo sets EndpointStatusInfo field to given value.

### HasEndpointStatusInfo

`func (o *KubernetesCredentials) HasEndpointStatusInfo() bool`

HasEndpointStatusInfo returns a boolean if a field has been set.

### GetLabel

`func (o *KubernetesCredentials) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *KubernetesCredentials) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *KubernetesCredentials) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetEndpointUrl

`func (o *KubernetesCredentials) GetEndpointUrl() string`

GetEndpointUrl returns the EndpointUrl field if non-nil, zero value otherwise.

### GetEndpointUrlOk

`func (o *KubernetesCredentials) GetEndpointUrlOk() (*string, bool)`

GetEndpointUrlOk returns a tuple with the EndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUrl

`func (o *KubernetesCredentials) SetEndpointUrl(v string)`

SetEndpointUrl sets EndpointUrl field to given value.


### GetAuthToken

`func (o *KubernetesCredentials) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *KubernetesCredentials) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *KubernetesCredentials) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *KubernetesCredentials) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetEventsIntegrationEnabled

`func (o *KubernetesCredentials) GetEventsIntegrationEnabled() bool`

GetEventsIntegrationEnabled returns the EventsIntegrationEnabled field if non-nil, zero value otherwise.

### GetEventsIntegrationEnabledOk

`func (o *KubernetesCredentials) GetEventsIntegrationEnabledOk() (*bool, bool)`

GetEventsIntegrationEnabledOk returns a tuple with the EventsIntegrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsIntegrationEnabled

`func (o *KubernetesCredentials) SetEventsIntegrationEnabled(v bool)`

SetEventsIntegrationEnabled sets EventsIntegrationEnabled field to given value.

### HasEventsIntegrationEnabled

`func (o *KubernetesCredentials) HasEventsIntegrationEnabled() bool`

HasEventsIntegrationEnabled returns a boolean if a field has been set.

### GetWorkloadIntegrationEnabled

`func (o *KubernetesCredentials) GetWorkloadIntegrationEnabled() bool`

GetWorkloadIntegrationEnabled returns the WorkloadIntegrationEnabled field if non-nil, zero value otherwise.

### GetWorkloadIntegrationEnabledOk

`func (o *KubernetesCredentials) GetWorkloadIntegrationEnabledOk() (*bool, bool)`

GetWorkloadIntegrationEnabledOk returns a tuple with the WorkloadIntegrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadIntegrationEnabled

`func (o *KubernetesCredentials) SetWorkloadIntegrationEnabled(v bool)`

SetWorkloadIntegrationEnabled sets WorkloadIntegrationEnabled field to given value.

### HasWorkloadIntegrationEnabled

`func (o *KubernetesCredentials) HasWorkloadIntegrationEnabled() bool`

HasWorkloadIntegrationEnabled returns a boolean if a field has been set.

### GetEventsFieldSelectors

`func (o *KubernetesCredentials) GetEventsFieldSelectors() []KubernetesEventPattern`

GetEventsFieldSelectors returns the EventsFieldSelectors field if non-nil, zero value otherwise.

### GetEventsFieldSelectorsOk

`func (o *KubernetesCredentials) GetEventsFieldSelectorsOk() (*[]KubernetesEventPattern, bool)`

GetEventsFieldSelectorsOk returns a tuple with the EventsFieldSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsFieldSelectors

`func (o *KubernetesCredentials) SetEventsFieldSelectors(v []KubernetesEventPattern)`

SetEventsFieldSelectors sets EventsFieldSelectors field to given value.

### HasEventsFieldSelectors

`func (o *KubernetesCredentials) HasEventsFieldSelectors() bool`

HasEventsFieldSelectors returns a boolean if a field has been set.

### GetCertificateCheckEnabled

`func (o *KubernetesCredentials) GetCertificateCheckEnabled() bool`

GetCertificateCheckEnabled returns the CertificateCheckEnabled field if non-nil, zero value otherwise.

### GetCertificateCheckEnabledOk

`func (o *KubernetesCredentials) GetCertificateCheckEnabledOk() (*bool, bool)`

GetCertificateCheckEnabledOk returns a tuple with the CertificateCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCheckEnabled

`func (o *KubernetesCredentials) SetCertificateCheckEnabled(v bool)`

SetCertificateCheckEnabled sets CertificateCheckEnabled field to given value.

### HasCertificateCheckEnabled

`func (o *KubernetesCredentials) HasCertificateCheckEnabled() bool`

HasCertificateCheckEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


