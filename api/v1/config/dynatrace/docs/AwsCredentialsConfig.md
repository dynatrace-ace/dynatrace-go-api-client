# AwsCredentialsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | **string** | The unique ID of the credentials. | [optional] [readonly] 
**ConnectionStatus** | **string** | The status of the connection to the AWS environment.    * &#x60;CONNECTED&#x60;: There was a connection within last 10 minutes.  * &#x60;DISCONNECTED&#x60;: A problem occurred with establishing connection using these credentials. Check whether the data is correct.  * &#x60;UNINITIALIZED&#x60;: The successful connection has never been established for these credentials. | [optional] [readonly] 
**Label** | **string** | The name of the credentials. | 
**PartitionType** | **string** | The type of the AWS partition. | 
**AuthenticationData** | [**AwsAuthenticationData**](AwsAuthenticationData.md) |  | 
**TaggedOnly** | **bool** | Monitor only resources which have specified AWS tags (&#x60;true&#x60;) or all resources (&#x60;false&#x60;). | 
**TagsToMonitor** | [**[]AwsConfigTag**](AwsConfigTag.md) | A list of AWS tags to be monitored.   You can specify up to 10 tags.   Only applicable when the **taggedOnly** parameter is set to &#x60;true&#x60;. | 
**SupportingServicesToMonitor** | [**[]AwsSupportingServiceConfig**](AwsSupportingServiceConfig.md) | A list of supporting services to be monitored. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


