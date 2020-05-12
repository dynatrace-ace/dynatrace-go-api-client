# TimeoutSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimedActionSupport** | **bool** | Timed action support enabled/disabled.   Enable to detect actions that trigger sending of XHRs via *setTimout* methods. | 
**TemporaryActionLimit** | **int32** | Defines how deep temporary actions may cascade. 0 disables temporary actions completely. Recommended value if enabled is 3. | 
**TemporaryActionTotalTimeout** | **int32** | The total timeout of all cascaded timeouts that should still be able to create a temporary action | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


