# JiraNotificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the notification configuration. | [optional] 
**Name** | **string** | The name of the notification configuration. | 
**AlertingProfile** | **string** | The ID of the associated alerting profile. | 
**Active** | **bool** | The configuration is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;EMAIL&#x60; -&gt; EmailNotificationConfig  * &#x60;PAGER_DUTY&#x60; -&gt; PagerDutyNotificationConfig  * &#x60;WEBHOOK&#x60; -&gt; WebHookNotificationConfig  * &#x60;SLACK&#x60; -&gt; SlackNotificationConfig  * &#x60;HIPCHAT&#x60; -&gt; HipChatNotificationConfig  * &#x60;VICTOROPS&#x60; -&gt; VictorOpsNotificationConfig  * &#x60;SERVICE_NOW&#x60; -&gt; ServiceNowNotificationConfig  * &#x60;XMATTERS&#x60; -&gt; XMattersNotificationConfig  * &#x60;ANSIBLETOWER&#x60; -&gt; AnsibleTowerNotificationConfig  * &#x60;OPS_GENIE&#x60; -&gt; OpsGenieNotificationConfig  * &#x60;JIRA&#x60; -&gt; JiraNotificationConfig  * &#x60;TRELLO&#x60; -&gt; TrelloNotificationConfig   | 
**Username** | **string** | The username of the Jira profile. | 
**Password** | **string** | The password for the Jira profile. | [optional] 
**Url** | **string** | The URL of the Jira API endpoint. | 
**ProjectKey** | **string** | The project key of the Jira issue to be created by this notification. | 
**IssueType** | **string** | The type of the Jira issue to be created by this notification. | 
**Summary** | **string** | The summary of the Jira issue to be created by this notification.   You can use the following placeholders:  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsText}&#x60;: All problem event details, including root cause, as a text-formatted string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;. If the problem has been merged into another problem, it has the &#x60;MERGED&#x60; value.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | 
**Description** | **string** | The description of the Jira issue to be created by this notification.    You can use same placeholders as in issue summary. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


