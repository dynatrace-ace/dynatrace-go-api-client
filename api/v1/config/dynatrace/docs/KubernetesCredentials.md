# KubernetesCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | **string** | The ID of the given credentials configuration. | [optional] [readonly] 
**Active** | **bool** | The monitoring is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;) for given credentials configuration.   If not set on creation, the &#x60;true&#x60; value is used.   If the field is omitted during an update, the old value remains unaffected. | [optional] 
**EndpointStatus** | **string** | The status of the configured endpoint.  ASSIGNED: The credentials are assigned to an ActiveGate which is responsible for processing. UNASSIGNED: The credentials are not yet assigned to an ActiveGate so there is currently no processing. DISABLED: The credentials have been disabled by the user. FASTCHECK_AUTH_ERROR: The credentials are invalid. FASTCHECK_TLS_ERROR: The endpoint TLS certificate is invalid. FASTCHECK_NO_RESPONSE: The endpoint did not return a result until the timeout was reached. FASTCHECK_INVALID_ENDPOINT: The endpoint URL was invalid. FASTCHECK_AUTH_LOCKED: The credentials seem to be locked. UNKNOWN: An unknown error occured.  | [optional] [readonly] 
**EndpointStatusInfo** | **string** | The detailed status info of the configured endpoint. | [optional] [readonly] 
**Label** | **string** | The name of the Kubernetes credentials configuration.   Allowed characters are letters, numbers, whitespaces, and the following characters: &#x60;.+-_&#x60;. Leading or trailing whitespace is not allowed. | 
**EndpointUrl** | **string** | The URL of the Kubernetes API server.   It must be unique within a Dynatrace environment.   The URL must valid according to RFC 2396. Leading or trailing whitespaces are not allowed. | [optional] 
**AuthToken** | **string** | The service account bearer token for the Kubernetes API server.   Submit your token on creation or update of the configuration. For security reasons, GET requests return this field as &#x60;null&#x60;.   If the field is omitted during an update, the old value remains unaffected. | [optional] 
**EventsIntegrationEnabled** | **bool** | The monitoring of events is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;) for the Kubernetes cluster. Event monitoring depends on the active state of this configuration to be true.   If not set on creation, the &#x60;true&#x60; value is used.   If the field is omitted during an update, the old value remains unaffected. | [optional] 
**WorkloadIntegrationEnabled** | **bool** | Workload and cloud application processing is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;) for the Kubernetes cluster. If the field is omitted during an update, the old value remains unaffected. | [optional] 
**EventsFieldSelectors** | [**[]KubernetesEventPattern**](KubernetesEventPattern.md) | Kubernetes event filters based on field-selectors. If set to &#x60;null&#x60; on creation, no events field selectors are subscribed. If set to &#x60;null&#x60; on update, no change of stored events field selectors is applied. Set an empty list to clear all events field selectors. | [optional] 
**CertificateCheckEnabled** | **bool** | The check of SSL certificates is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;) for the Kubernetes cluster.   If not set on creation, the &#x60;true&#x60; value is used.   If the field is omitted during an update, the old value remains unaffected. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


