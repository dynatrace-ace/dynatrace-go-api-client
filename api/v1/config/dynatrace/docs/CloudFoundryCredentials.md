# CloudFoundryCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | **string** | The ID of the given credentials configuration. | [optional] [readonly] 
**Active** | **bool** | The monitoring is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;) for given credentials configuration.   If not set on creation, the &#x60;true&#x60; value is used.   If the field is omitted during an update, the old value remains unaffected. | [optional] 
**EndpointStatus** | **string** | The status of the configured endpoint.  ASSIGNED: The credentials are assigned to an ActiveGate which is responsible for processing. UNASSIGNED: The credentials are not yet assigned to an ActiveGate so there is currently no processing. DISABLED: The credentials have been disabled by the user. FASTCHECK_AUTH_ERROR: The credentials are invalid. FASTCHECK_TLS_ERROR: The endpoint TLS certificate is invalid. FASTCHECK_NO_RESPONSE: The endpoint did not return a result until the timeout was reached. FASTCHECK_INVALID_ENDPOINT: The endpoint URL was invalid. FASTCHECK_AUTH_LOCKED: The credentials seem to be locked. UNKNOWN: An unknown error occured.  | [optional] [readonly] 
**EndpointStatusInfo** | **string** | The detailed status info of the configured endpoint. | [optional] [readonly] 
**Name** | **string** | The name of the Cloud Foundry foundation credentials.   Allowed characters are letters, numbers, whitespaces, and the following characters: &#x60;.+-_&#x60;. Leading or trailing whitespace is not allowed. | 
**ApiUrl** | **string** | The URL of the Cloud Foundry foundation credentials.   The URL must be valid according to RFC 2396.   Leading or trailing whitespaces are not allowed. | [optional] 
**LoginUrl** | **string** | The login URL of the Cloud Foundry foundation credentials.   The URL must be valid according to RFC 2396.   Leading or trailing whitespaces are not allowed. | [optional] 
**Username** | **string** | The username of the Cloud Foundry foundation credentials.   Leading and trailing whitespaces are not allowed. | 
**Password** | **string** | The password of the Cloud Foundry foundation credentials. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


