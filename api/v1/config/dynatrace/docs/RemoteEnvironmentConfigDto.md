# RemoteEnvironmentConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | The API token granting access to the remote environment.   The token must have the **Fetch data from a remote environment** (&#x60;RestRequestForwarding&#x60;) scope. You can create such a token only via [Tokens API](https://www.dynatrace.com/support/help/shortlink/api-tokens-post-new).   For security reasons, GET requests return this field as &#x60;null&#x60;. | [optional] 
**NetworkScope** | **string** | The network scope of the remote environment: * &#x60;EXTERNAL&#x60;: The remote environment is located in an another network.  * &#x60;INTERNAL&#x60;: The remote environment is located in the same network.  * &#x60;CLUSTER&#x60;: The remote environment is located in the same cluster.   Dynatrace SaaS can only use &#x60;EXTERNAL&#x60;.  If not set, &#x60;EXTERNAL&#x60; is used. | [optional] 
**Id** | **string** | The ID of the configuration. | [optional] 
**DisplayName** | **string** | The display name of the remote environment. | 
**Uri** | **string** | The URI of the remote environment. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


