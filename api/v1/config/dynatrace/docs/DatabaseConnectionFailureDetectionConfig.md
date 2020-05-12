# DatabaseConnectionFailureDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**ConnectionFailsCount** | **int32** | Number of failed database connections during any **timePeriodMinutes** minutes period to trigger an alert. | [optional] 
**TimePeriodMinutes** | **int32** | The *X* minutes time period during which the **connectionFailsCount** is evaluated. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


