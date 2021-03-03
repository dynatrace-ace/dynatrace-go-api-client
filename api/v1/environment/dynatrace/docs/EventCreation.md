# EventCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** | The type of the event. | 
**Start** | Pointer to **int64** | The start timestamp of the event, in UTC milliseconds.   If not set, the current timestamp is used.   You can report information-only events up to **30 days** into the past.   You can report problem-opening events up to **60 minutes** into the past. | [optional] 
**End** | Pointer to **int64** | The end timestamp of the event, in UTC milliseconds.   If not set, the current time is used for information-only events.   Not applicable to problem-opening events. Such an event stays open until it times out depending on the **timeoutMinutes** parameter. | [optional] 
**TimeoutMinutes** | Pointer to **int32** | The timeout for problem-opening events in minutes. Not applicable to information-only events.   If not set, 15 minutes is used. The maximum allowed value is 120 minutes.   You can refresh the event by sending the same payload again. | [optional] 
**AttachRules** | [**PushEventAttachRules**](PushEventAttachRules.md) |  | 
**CustomProperties** | Pointer to **map[string]map[string]interface{}** | The set of any properties related to the event, in the *\&quot;key\&quot; : \&quot;value\&quot;* format. | [optional] 
**Source** | **string** | The name or ID of the external source of the event. | 
**AnnotationType** | Pointer to **string** | The type of the custom annotation, for example &#x60;DNS route has been changed&#x60;. | [optional] 
**AnnotationDescription** | Pointer to **string** | A detailed description of the custom annotation, for example &#x60;DNS route has been changed to x.mydomain.com&#x60;. | [optional] 
**Description** | Pointer to **string** | The textual description of the configuration change. | [optional] 
**DeploymentName** | Pointer to **string** | The ID of the triggered deployment. | [optional] 
**DeploymentVersion** | Pointer to **string** | The version of the triggered deployment. | [optional] 
**TimeseriesIds** | Pointer to **[]string** | A list of timeseries IDs, related to the event. | [optional] 
**DeploymentProject** | Pointer to **string** | The project name of the deployed package. | [optional] 
**CiBackLink** | Pointer to **string** | The link to the deployed artifact within the 3rd party system. | [optional] 
**RemediationAction** | Pointer to **string** | The link to the deployment related remediation action within the external tool. | [optional] 
**Original** | Pointer to **string** | The previous value of the configuration. | [optional] 
**Changed** | Pointer to **string** | The new value of the configuration that has been set by the event. | [optional] 
**Configuration** | Pointer to **string** | The ID or the name of the configuration that has been changed by the event. | [optional] 
**Title** | Pointer to **string** | The title of the configuration that has been set by the event. | [optional] 
**AllowDavisMerge** | Pointer to **bool** | Allow Davis AI to merge this event into existing problems (true) or force creating a new problem (false).  This only applies to problem-opening event types. | [optional] 

## Methods

### NewEventCreation

`func NewEventCreation(eventType string, attachRules PushEventAttachRules, source string, ) *EventCreation`

NewEventCreation instantiates a new EventCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventCreationWithDefaults

`func NewEventCreationWithDefaults() *EventCreation`

NewEventCreationWithDefaults instantiates a new EventCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *EventCreation) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventCreation) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventCreation) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetStart

`func (o *EventCreation) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EventCreation) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EventCreation) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *EventCreation) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *EventCreation) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *EventCreation) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *EventCreation) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *EventCreation) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetTimeoutMinutes

`func (o *EventCreation) GetTimeoutMinutes() int32`

GetTimeoutMinutes returns the TimeoutMinutes field if non-nil, zero value otherwise.

### GetTimeoutMinutesOk

`func (o *EventCreation) GetTimeoutMinutesOk() (*int32, bool)`

GetTimeoutMinutesOk returns a tuple with the TimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMinutes

`func (o *EventCreation) SetTimeoutMinutes(v int32)`

SetTimeoutMinutes sets TimeoutMinutes field to given value.

### HasTimeoutMinutes

`func (o *EventCreation) HasTimeoutMinutes() bool`

HasTimeoutMinutes returns a boolean if a field has been set.

### GetAttachRules

`func (o *EventCreation) GetAttachRules() PushEventAttachRules`

GetAttachRules returns the AttachRules field if non-nil, zero value otherwise.

### GetAttachRulesOk

`func (o *EventCreation) GetAttachRulesOk() (*PushEventAttachRules, bool)`

GetAttachRulesOk returns a tuple with the AttachRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachRules

`func (o *EventCreation) SetAttachRules(v PushEventAttachRules)`

SetAttachRules sets AttachRules field to given value.


### GetCustomProperties

`func (o *EventCreation) GetCustomProperties() map[string]map[string]interface{}`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *EventCreation) GetCustomPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *EventCreation) SetCustomProperties(v map[string]map[string]interface{})`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *EventCreation) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetSource

`func (o *EventCreation) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EventCreation) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EventCreation) SetSource(v string)`

SetSource sets Source field to given value.


### GetAnnotationType

`func (o *EventCreation) GetAnnotationType() string`

GetAnnotationType returns the AnnotationType field if non-nil, zero value otherwise.

### GetAnnotationTypeOk

`func (o *EventCreation) GetAnnotationTypeOk() (*string, bool)`

GetAnnotationTypeOk returns a tuple with the AnnotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationType

`func (o *EventCreation) SetAnnotationType(v string)`

SetAnnotationType sets AnnotationType field to given value.

### HasAnnotationType

`func (o *EventCreation) HasAnnotationType() bool`

HasAnnotationType returns a boolean if a field has been set.

### GetAnnotationDescription

`func (o *EventCreation) GetAnnotationDescription() string`

GetAnnotationDescription returns the AnnotationDescription field if non-nil, zero value otherwise.

### GetAnnotationDescriptionOk

`func (o *EventCreation) GetAnnotationDescriptionOk() (*string, bool)`

GetAnnotationDescriptionOk returns a tuple with the AnnotationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationDescription

`func (o *EventCreation) SetAnnotationDescription(v string)`

SetAnnotationDescription sets AnnotationDescription field to given value.

### HasAnnotationDescription

`func (o *EventCreation) HasAnnotationDescription() bool`

HasAnnotationDescription returns a boolean if a field has been set.

### GetDescription

`func (o *EventCreation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventCreation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventCreation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventCreation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeploymentName

`func (o *EventCreation) GetDeploymentName() string`

GetDeploymentName returns the DeploymentName field if non-nil, zero value otherwise.

### GetDeploymentNameOk

`func (o *EventCreation) GetDeploymentNameOk() (*string, bool)`

GetDeploymentNameOk returns a tuple with the DeploymentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentName

`func (o *EventCreation) SetDeploymentName(v string)`

SetDeploymentName sets DeploymentName field to given value.

### HasDeploymentName

`func (o *EventCreation) HasDeploymentName() bool`

HasDeploymentName returns a boolean if a field has been set.

### GetDeploymentVersion

`func (o *EventCreation) GetDeploymentVersion() string`

GetDeploymentVersion returns the DeploymentVersion field if non-nil, zero value otherwise.

### GetDeploymentVersionOk

`func (o *EventCreation) GetDeploymentVersionOk() (*string, bool)`

GetDeploymentVersionOk returns a tuple with the DeploymentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentVersion

`func (o *EventCreation) SetDeploymentVersion(v string)`

SetDeploymentVersion sets DeploymentVersion field to given value.

### HasDeploymentVersion

`func (o *EventCreation) HasDeploymentVersion() bool`

HasDeploymentVersion returns a boolean if a field has been set.

### GetTimeseriesIds

`func (o *EventCreation) GetTimeseriesIds() []string`

GetTimeseriesIds returns the TimeseriesIds field if non-nil, zero value otherwise.

### GetTimeseriesIdsOk

`func (o *EventCreation) GetTimeseriesIdsOk() (*[]string, bool)`

GetTimeseriesIdsOk returns a tuple with the TimeseriesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseriesIds

`func (o *EventCreation) SetTimeseriesIds(v []string)`

SetTimeseriesIds sets TimeseriesIds field to given value.

### HasTimeseriesIds

`func (o *EventCreation) HasTimeseriesIds() bool`

HasTimeseriesIds returns a boolean if a field has been set.

### GetDeploymentProject

`func (o *EventCreation) GetDeploymentProject() string`

GetDeploymentProject returns the DeploymentProject field if non-nil, zero value otherwise.

### GetDeploymentProjectOk

`func (o *EventCreation) GetDeploymentProjectOk() (*string, bool)`

GetDeploymentProjectOk returns a tuple with the DeploymentProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProject

`func (o *EventCreation) SetDeploymentProject(v string)`

SetDeploymentProject sets DeploymentProject field to given value.

### HasDeploymentProject

`func (o *EventCreation) HasDeploymentProject() bool`

HasDeploymentProject returns a boolean if a field has been set.

### GetCiBackLink

`func (o *EventCreation) GetCiBackLink() string`

GetCiBackLink returns the CiBackLink field if non-nil, zero value otherwise.

### GetCiBackLinkOk

`func (o *EventCreation) GetCiBackLinkOk() (*string, bool)`

GetCiBackLinkOk returns a tuple with the CiBackLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiBackLink

`func (o *EventCreation) SetCiBackLink(v string)`

SetCiBackLink sets CiBackLink field to given value.

### HasCiBackLink

`func (o *EventCreation) HasCiBackLink() bool`

HasCiBackLink returns a boolean if a field has been set.

### GetRemediationAction

`func (o *EventCreation) GetRemediationAction() string`

GetRemediationAction returns the RemediationAction field if non-nil, zero value otherwise.

### GetRemediationActionOk

`func (o *EventCreation) GetRemediationActionOk() (*string, bool)`

GetRemediationActionOk returns a tuple with the RemediationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationAction

`func (o *EventCreation) SetRemediationAction(v string)`

SetRemediationAction sets RemediationAction field to given value.

### HasRemediationAction

`func (o *EventCreation) HasRemediationAction() bool`

HasRemediationAction returns a boolean if a field has been set.

### GetOriginal

`func (o *EventCreation) GetOriginal() string`

GetOriginal returns the Original field if non-nil, zero value otherwise.

### GetOriginalOk

`func (o *EventCreation) GetOriginalOk() (*string, bool)`

GetOriginalOk returns a tuple with the Original field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginal

`func (o *EventCreation) SetOriginal(v string)`

SetOriginal sets Original field to given value.

### HasOriginal

`func (o *EventCreation) HasOriginal() bool`

HasOriginal returns a boolean if a field has been set.

### GetChanged

`func (o *EventCreation) GetChanged() string`

GetChanged returns the Changed field if non-nil, zero value otherwise.

### GetChangedOk

`func (o *EventCreation) GetChangedOk() (*string, bool)`

GetChangedOk returns a tuple with the Changed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanged

`func (o *EventCreation) SetChanged(v string)`

SetChanged sets Changed field to given value.

### HasChanged

`func (o *EventCreation) HasChanged() bool`

HasChanged returns a boolean if a field has been set.

### GetConfiguration

`func (o *EventCreation) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *EventCreation) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *EventCreation) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *EventCreation) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTitle

`func (o *EventCreation) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventCreation) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EventCreation) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EventCreation) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAllowDavisMerge

`func (o *EventCreation) GetAllowDavisMerge() bool`

GetAllowDavisMerge returns the AllowDavisMerge field if non-nil, zero value otherwise.

### GetAllowDavisMergeOk

`func (o *EventCreation) GetAllowDavisMergeOk() (*bool, bool)`

GetAllowDavisMergeOk returns a tuple with the AllowDavisMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDavisMerge

`func (o *EventCreation) SetAllowDavisMerge(v bool)`

SetAllowDavisMerge sets AllowDavisMerge field to given value.

### HasAllowDavisMerge

`func (o *EventCreation) HasAllowDavisMerge() bool`

HasAllowDavisMerge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


