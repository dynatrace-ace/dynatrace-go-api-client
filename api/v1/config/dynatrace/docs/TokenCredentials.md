# TokenCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the credentials set. | 
**Id** | **string** | The ID of the credentials set. | [optional] 
**Description** | **string** | A short description of the credentials set.. | 
**OwnerAccessOnly** | **bool** | The credentials set is available to every user (&#x60;false&#x60;) or to owner only (&#x60;true&#x60;). | [optional] 
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;CERTIFICATE&#x60; -&gt; CertificateCredentials  * &#x60;USERNAME_PASSWORD&#x60; -&gt; UserPasswordCredentials  * &#x60;TOKEN&#x60; -&gt; TokenCredentials   | [optional] 
**Token** | **string** | Token in the string format. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


