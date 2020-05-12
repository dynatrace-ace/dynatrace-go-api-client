# AlertingPredefinedEventFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** | The type of the predefined event. | 
**Negate** | **bool** | The alert triggers when the problem of specified severity arises while the specified event **is** happening (&#x60;false&#x60;) or while the specified event is **not** happening (&#x60;true&#x60;).    For example, if you chose the Slowdown (&#x60;PERFORMANCE&#x60;) severity and Unexpected high traffic (&#x60;APPLICATION_UNEXPECTED_HIGH_LOAD&#x60;) event with **negate** set to &#x60;true&#x60;, the alerting profile will trigger only when the slowdown problem is raised while there is no unexpected high traffic event.   Consider the following use case as an example. The Slowdown (&#x60;PERFORMANCE&#x60;) severity rule is set. Depending on the configuration of the event filter (Unexpected high traffic (&#x60;APPLICATION_UNEXPECTED_HIGH_LOAD&#x60;) event is used as an example), the behavior of the alerting profile is one of the following:* **negate** is set to &#x60;false&#x60;: The alert triggers when the slowdown problem is raised while unexpected high traffic event is happening.  * **negate** is set to &#x60;true&#x60;: The alert triggers when the slowdown problem is raised while there is no unexpected high traffic event.   * no event rule is set: The alert triggers when the slowdown problem is raised, regardless of any events. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


