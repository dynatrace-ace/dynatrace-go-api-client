# MethodReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visibility** | **string** | The visibility of the method to capture. | 
**Modifiers** | **[]string** | The modifiers of the method to capture. | 
**ClassName** | **string** | The class name where the method to capture resides.    Either this or the **fileName** must be set. | [optional] 
**FileName** | **string** | The file name where the method to capture resides.    Either this or **className** must be set. | [optional] 
**FileNameMatcher** | **string** | The operator of the comparison.    If not set, &#x60;EQUALS&#x60; is used. | [optional] 
**MethodName** | **string** | The name of the method to capture. | 
**ArgumentTypes** | **[]string** | The list of argument types. | 
**ReturnType** | **string** | The return type. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


