# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **int64** | The timestamp of the event detection, in UTC milliseconds. | [optional] 
**EndTime** | Pointer to **int64** | The timestamp of the event closure, in UTC milliseconds.   Has the &#x60;-1&#x60; value if the event is still open. | [optional] 
**EntityId** | Pointer to **string** | The ID of the affected Dynatrace entity. | [optional] 
**EntityName** | Pointer to **string** | The name of the affected Dynatrace entity. | [optional] 
**SeverityLevel** | Pointer to **string** | The severity of the event. | [optional] 
**ImpactLevel** | Pointer to **string** | The impact level of the event. It shows what is affected by the problem: infrastructure, service, or application. | [optional] 
**EventType** | Pointer to **string** | The type of the event. | [optional] 
**ResourceId** | Pointer to **string** | The id of the resource the event occurred on. | [optional] 
**ResourceName** | Pointer to **string** | The name of the resource the event occurred on. | [optional] 
**Status** | Pointer to **string** | The status of the event. | [optional] 
**Severities** | Pointer to [**[]EventSeverity**](EventSeverity.md) | Additional data on the event severity. | [optional] 
**IsRootCause** | Pointer to **bool** | Indicates if the event is the root cause of a problem. | [optional] 
**DeploymentProject** | Pointer to **string** |  | [optional] 
**CpuLimitInMHz** | Pointer to **int32** |  | [optional] 
**DeploymentParamAdded** | Pointer to **string** |  | [optional] 
**AffectedPrivateSyntheticLocations** | Pointer to **[]string** |  | [optional] 
**IsClusterWide** | Pointer to **bool** | For events with event type &#x60;MONITORING_UNAVAILABLE&#x60;, it may be that the event is occurring on the entire Dynatrace cluster. If this is true, it could be that there are problems on the Dynatrace side. | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**EffectiveEntity** | Pointer to **string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**Artifact** | Pointer to **string** |  | [optional] 
**CpuLoad** | Pointer to **float32** |  | [optional] 
**AffectedRequestsPerMinute** | Pointer to **float32** |  | [optional] 
**AnnotationDescription** | Pointer to **string** |  | [optional] 
**Browser** | Pointer to **string** |  | [optional] 
**AffectedSyntheticLocations** | Pointer to **[]string** |  | [optional] 
**DeploymentName** | Pointer to **string** |  | [optional] 
**DeploymentParamRemoved** | Pointer to **string** |  | [optional] 
**CorrelationId** | Pointer to **string** |  | [optional] 
**UserDefined50thPercentileThreshold** | Pointer to **float32** |  | [optional] 
**SyntheticErrorType** | Pointer to **[]string** | If the event type is one of the synthetic event types then we may have information about the error type. The names of those are returned in this field. | [optional] 
**ServiceMethodGroup** | Pointer to **string** |  | [optional] 
**ReferenceResponseTime90thPercentile** | Pointer to **float32** |  | [optional] 
**UserAction** | Pointer to **string** |  | [optional] 
**MinimumProcessGroupInstanceCountThreshold** | Pointer to **int32** |  | [optional] 
**ReferenceResponseTime50thPercentile** | Pointer to **float32** |  | [optional] 
**Original** | Pointer to **string** |  | [optional] 
**UserDefined90thPercentileThreshold** | Pointer to **float32** |  | [optional] 
**DeploymentVersion** | Pointer to **string** |  | [optional] 
**AnnotationType** | Pointer to **string** |  | [optional] 
**AffectedSyntheticActions** | Pointer to **[]string** | If the event type is one of the synthetic event types then we may have information on which synthetic actions are affected by the event. The names of those are returned in this field. | [optional] 
**AffectedUserActionsPerMinute** | Pointer to **float32** |  | [optional] 
**ActiveMaintenanceWindows** | Pointer to [**[]EntityShortRepresentation**](EntityShortRepresentation.md) |  | [optional] 
**MobileAppVersion** | Pointer to **string** |  | [optional] 
**UserDefinedFailureRateThreshold** | Pointer to **float32** |  | [optional] 
**Percentile** | Pointer to **string** |  | [optional] 
**CustomProperties** | Pointer to **map[string]string** |  | [optional] 
**RemediationAction** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**CiBackLink** | Pointer to **string** |  | [optional] 
**Geolocation** | Pointer to **string** |  | [optional] 
**ServiceMethod** | Pointer to **string** |  | [optional] 
**Changed** | Pointer to **string** |  | [optional] 

## Methods

### NewEvent

`func NewEvent() *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *Event) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Event) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Event) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Event) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *Event) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Event) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Event) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Event) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEntityId

`func (o *Event) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Event) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Event) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *Event) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityName

`func (o *Event) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *Event) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *Event) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *Event) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetSeverityLevel

`func (o *Event) GetSeverityLevel() string`

GetSeverityLevel returns the SeverityLevel field if non-nil, zero value otherwise.

### GetSeverityLevelOk

`func (o *Event) GetSeverityLevelOk() (*string, bool)`

GetSeverityLevelOk returns a tuple with the SeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevel

`func (o *Event) SetSeverityLevel(v string)`

SetSeverityLevel sets SeverityLevel field to given value.

### HasSeverityLevel

`func (o *Event) HasSeverityLevel() bool`

HasSeverityLevel returns a boolean if a field has been set.

### GetImpactLevel

`func (o *Event) GetImpactLevel() string`

GetImpactLevel returns the ImpactLevel field if non-nil, zero value otherwise.

### GetImpactLevelOk

`func (o *Event) GetImpactLevelOk() (*string, bool)`

GetImpactLevelOk returns a tuple with the ImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactLevel

`func (o *Event) SetImpactLevel(v string)`

SetImpactLevel sets ImpactLevel field to given value.

### HasImpactLevel

`func (o *Event) HasImpactLevel() bool`

HasImpactLevel returns a boolean if a field has been set.

### GetEventType

`func (o *Event) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Event) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Event) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *Event) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetResourceId

`func (o *Event) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Event) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Event) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *Event) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceName

`func (o *Event) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *Event) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *Event) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *Event) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetStatus

`func (o *Event) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Event) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Event) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Event) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSeverities

`func (o *Event) GetSeverities() []EventSeverity`

GetSeverities returns the Severities field if non-nil, zero value otherwise.

### GetSeveritiesOk

`func (o *Event) GetSeveritiesOk() (*[]EventSeverity, bool)`

GetSeveritiesOk returns a tuple with the Severities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverities

`func (o *Event) SetSeverities(v []EventSeverity)`

SetSeverities sets Severities field to given value.

### HasSeverities

`func (o *Event) HasSeverities() bool`

HasSeverities returns a boolean if a field has been set.

### GetIsRootCause

`func (o *Event) GetIsRootCause() bool`

GetIsRootCause returns the IsRootCause field if non-nil, zero value otherwise.

### GetIsRootCauseOk

`func (o *Event) GetIsRootCauseOk() (*bool, bool)`

GetIsRootCauseOk returns a tuple with the IsRootCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRootCause

`func (o *Event) SetIsRootCause(v bool)`

SetIsRootCause sets IsRootCause field to given value.

### HasIsRootCause

`func (o *Event) HasIsRootCause() bool`

HasIsRootCause returns a boolean if a field has been set.

### GetDeploymentProject

`func (o *Event) GetDeploymentProject() string`

GetDeploymentProject returns the DeploymentProject field if non-nil, zero value otherwise.

### GetDeploymentProjectOk

`func (o *Event) GetDeploymentProjectOk() (*string, bool)`

GetDeploymentProjectOk returns a tuple with the DeploymentProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProject

`func (o *Event) SetDeploymentProject(v string)`

SetDeploymentProject sets DeploymentProject field to given value.

### HasDeploymentProject

`func (o *Event) HasDeploymentProject() bool`

HasDeploymentProject returns a boolean if a field has been set.

### GetCpuLimitInMHz

`func (o *Event) GetCpuLimitInMHz() int32`

GetCpuLimitInMHz returns the CpuLimitInMHz field if non-nil, zero value otherwise.

### GetCpuLimitInMHzOk

`func (o *Event) GetCpuLimitInMHzOk() (*int32, bool)`

GetCpuLimitInMHzOk returns a tuple with the CpuLimitInMHz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimitInMHz

`func (o *Event) SetCpuLimitInMHz(v int32)`

SetCpuLimitInMHz sets CpuLimitInMHz field to given value.

### HasCpuLimitInMHz

`func (o *Event) HasCpuLimitInMHz() bool`

HasCpuLimitInMHz returns a boolean if a field has been set.

### GetDeploymentParamAdded

`func (o *Event) GetDeploymentParamAdded() string`

GetDeploymentParamAdded returns the DeploymentParamAdded field if non-nil, zero value otherwise.

### GetDeploymentParamAddedOk

`func (o *Event) GetDeploymentParamAddedOk() (*string, bool)`

GetDeploymentParamAddedOk returns a tuple with the DeploymentParamAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentParamAdded

`func (o *Event) SetDeploymentParamAdded(v string)`

SetDeploymentParamAdded sets DeploymentParamAdded field to given value.

### HasDeploymentParamAdded

`func (o *Event) HasDeploymentParamAdded() bool`

HasDeploymentParamAdded returns a boolean if a field has been set.

### GetAffectedPrivateSyntheticLocations

`func (o *Event) GetAffectedPrivateSyntheticLocations() []string`

GetAffectedPrivateSyntheticLocations returns the AffectedPrivateSyntheticLocations field if non-nil, zero value otherwise.

### GetAffectedPrivateSyntheticLocationsOk

`func (o *Event) GetAffectedPrivateSyntheticLocationsOk() (*[]string, bool)`

GetAffectedPrivateSyntheticLocationsOk returns a tuple with the AffectedPrivateSyntheticLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPrivateSyntheticLocations

`func (o *Event) SetAffectedPrivateSyntheticLocations(v []string)`

SetAffectedPrivateSyntheticLocations sets AffectedPrivateSyntheticLocations field to given value.

### HasAffectedPrivateSyntheticLocations

`func (o *Event) HasAffectedPrivateSyntheticLocations() bool`

HasAffectedPrivateSyntheticLocations returns a boolean if a field has been set.

### GetIsClusterWide

`func (o *Event) GetIsClusterWide() bool`

GetIsClusterWide returns the IsClusterWide field if non-nil, zero value otherwise.

### GetIsClusterWideOk

`func (o *Event) GetIsClusterWideOk() (*bool, bool)`

GetIsClusterWideOk returns a tuple with the IsClusterWide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClusterWide

`func (o *Event) SetIsClusterWide(v bool)`

SetIsClusterWide sets IsClusterWide field to given value.

### HasIsClusterWide

`func (o *Event) HasIsClusterWide() bool`

HasIsClusterWide returns a boolean if a field has been set.

### GetSource

`func (o *Event) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Event) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Event) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Event) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetEffectiveEntity

`func (o *Event) GetEffectiveEntity() string`

GetEffectiveEntity returns the EffectiveEntity field if non-nil, zero value otherwise.

### GetEffectiveEntityOk

`func (o *Event) GetEffectiveEntityOk() (*string, bool)`

GetEffectiveEntityOk returns a tuple with the EffectiveEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveEntity

`func (o *Event) SetEffectiveEntity(v string)`

SetEffectiveEntity sets EffectiveEntity field to given value.

### HasEffectiveEntity

`func (o *Event) HasEffectiveEntity() bool`

HasEffectiveEntity returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *Event) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *Event) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *Event) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *Event) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetArtifact

`func (o *Event) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *Event) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *Event) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *Event) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetCpuLoad

`func (o *Event) GetCpuLoad() float32`

GetCpuLoad returns the CpuLoad field if non-nil, zero value otherwise.

### GetCpuLoadOk

`func (o *Event) GetCpuLoadOk() (*float32, bool)`

GetCpuLoadOk returns a tuple with the CpuLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLoad

`func (o *Event) SetCpuLoad(v float32)`

SetCpuLoad sets CpuLoad field to given value.

### HasCpuLoad

`func (o *Event) HasCpuLoad() bool`

HasCpuLoad returns a boolean if a field has been set.

### GetAffectedRequestsPerMinute

`func (o *Event) GetAffectedRequestsPerMinute() float32`

GetAffectedRequestsPerMinute returns the AffectedRequestsPerMinute field if non-nil, zero value otherwise.

### GetAffectedRequestsPerMinuteOk

`func (o *Event) GetAffectedRequestsPerMinuteOk() (*float32, bool)`

GetAffectedRequestsPerMinuteOk returns a tuple with the AffectedRequestsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedRequestsPerMinute

`func (o *Event) SetAffectedRequestsPerMinute(v float32)`

SetAffectedRequestsPerMinute sets AffectedRequestsPerMinute field to given value.

### HasAffectedRequestsPerMinute

`func (o *Event) HasAffectedRequestsPerMinute() bool`

HasAffectedRequestsPerMinute returns a boolean if a field has been set.

### GetAnnotationDescription

`func (o *Event) GetAnnotationDescription() string`

GetAnnotationDescription returns the AnnotationDescription field if non-nil, zero value otherwise.

### GetAnnotationDescriptionOk

`func (o *Event) GetAnnotationDescriptionOk() (*string, bool)`

GetAnnotationDescriptionOk returns a tuple with the AnnotationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationDescription

`func (o *Event) SetAnnotationDescription(v string)`

SetAnnotationDescription sets AnnotationDescription field to given value.

### HasAnnotationDescription

`func (o *Event) HasAnnotationDescription() bool`

HasAnnotationDescription returns a boolean if a field has been set.

### GetBrowser

`func (o *Event) GetBrowser() string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *Event) GetBrowserOk() (*string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *Event) SetBrowser(v string)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *Event) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### GetAffectedSyntheticLocations

`func (o *Event) GetAffectedSyntheticLocations() []string`

GetAffectedSyntheticLocations returns the AffectedSyntheticLocations field if non-nil, zero value otherwise.

### GetAffectedSyntheticLocationsOk

`func (o *Event) GetAffectedSyntheticLocationsOk() (*[]string, bool)`

GetAffectedSyntheticLocationsOk returns a tuple with the AffectedSyntheticLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedSyntheticLocations

`func (o *Event) SetAffectedSyntheticLocations(v []string)`

SetAffectedSyntheticLocations sets AffectedSyntheticLocations field to given value.

### HasAffectedSyntheticLocations

`func (o *Event) HasAffectedSyntheticLocations() bool`

HasAffectedSyntheticLocations returns a boolean if a field has been set.

### GetDeploymentName

`func (o *Event) GetDeploymentName() string`

GetDeploymentName returns the DeploymentName field if non-nil, zero value otherwise.

### GetDeploymentNameOk

`func (o *Event) GetDeploymentNameOk() (*string, bool)`

GetDeploymentNameOk returns a tuple with the DeploymentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentName

`func (o *Event) SetDeploymentName(v string)`

SetDeploymentName sets DeploymentName field to given value.

### HasDeploymentName

`func (o *Event) HasDeploymentName() bool`

HasDeploymentName returns a boolean if a field has been set.

### GetDeploymentParamRemoved

`func (o *Event) GetDeploymentParamRemoved() string`

GetDeploymentParamRemoved returns the DeploymentParamRemoved field if non-nil, zero value otherwise.

### GetDeploymentParamRemovedOk

`func (o *Event) GetDeploymentParamRemovedOk() (*string, bool)`

GetDeploymentParamRemovedOk returns a tuple with the DeploymentParamRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentParamRemoved

`func (o *Event) SetDeploymentParamRemoved(v string)`

SetDeploymentParamRemoved sets DeploymentParamRemoved field to given value.

### HasDeploymentParamRemoved

`func (o *Event) HasDeploymentParamRemoved() bool`

HasDeploymentParamRemoved returns a boolean if a field has been set.

### GetCorrelationId

`func (o *Event) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *Event) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *Event) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *Event) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetUserDefined50thPercentileThreshold

`func (o *Event) GetUserDefined50thPercentileThreshold() float32`

GetUserDefined50thPercentileThreshold returns the UserDefined50thPercentileThreshold field if non-nil, zero value otherwise.

### GetUserDefined50thPercentileThresholdOk

`func (o *Event) GetUserDefined50thPercentileThresholdOk() (*float32, bool)`

GetUserDefined50thPercentileThresholdOk returns a tuple with the UserDefined50thPercentileThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefined50thPercentileThreshold

`func (o *Event) SetUserDefined50thPercentileThreshold(v float32)`

SetUserDefined50thPercentileThreshold sets UserDefined50thPercentileThreshold field to given value.

### HasUserDefined50thPercentileThreshold

`func (o *Event) HasUserDefined50thPercentileThreshold() bool`

HasUserDefined50thPercentileThreshold returns a boolean if a field has been set.

### GetSyntheticErrorType

`func (o *Event) GetSyntheticErrorType() []string`

GetSyntheticErrorType returns the SyntheticErrorType field if non-nil, zero value otherwise.

### GetSyntheticErrorTypeOk

`func (o *Event) GetSyntheticErrorTypeOk() (*[]string, bool)`

GetSyntheticErrorTypeOk returns a tuple with the SyntheticErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticErrorType

`func (o *Event) SetSyntheticErrorType(v []string)`

SetSyntheticErrorType sets SyntheticErrorType field to given value.

### HasSyntheticErrorType

`func (o *Event) HasSyntheticErrorType() bool`

HasSyntheticErrorType returns a boolean if a field has been set.

### GetServiceMethodGroup

`func (o *Event) GetServiceMethodGroup() string`

GetServiceMethodGroup returns the ServiceMethodGroup field if non-nil, zero value otherwise.

### GetServiceMethodGroupOk

`func (o *Event) GetServiceMethodGroupOk() (*string, bool)`

GetServiceMethodGroupOk returns a tuple with the ServiceMethodGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMethodGroup

`func (o *Event) SetServiceMethodGroup(v string)`

SetServiceMethodGroup sets ServiceMethodGroup field to given value.

### HasServiceMethodGroup

`func (o *Event) HasServiceMethodGroup() bool`

HasServiceMethodGroup returns a boolean if a field has been set.

### GetReferenceResponseTime90thPercentile

`func (o *Event) GetReferenceResponseTime90thPercentile() float32`

GetReferenceResponseTime90thPercentile returns the ReferenceResponseTime90thPercentile field if non-nil, zero value otherwise.

### GetReferenceResponseTime90thPercentileOk

`func (o *Event) GetReferenceResponseTime90thPercentileOk() (*float32, bool)`

GetReferenceResponseTime90thPercentileOk returns a tuple with the ReferenceResponseTime90thPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceResponseTime90thPercentile

`func (o *Event) SetReferenceResponseTime90thPercentile(v float32)`

SetReferenceResponseTime90thPercentile sets ReferenceResponseTime90thPercentile field to given value.

### HasReferenceResponseTime90thPercentile

`func (o *Event) HasReferenceResponseTime90thPercentile() bool`

HasReferenceResponseTime90thPercentile returns a boolean if a field has been set.

### GetUserAction

`func (o *Event) GetUserAction() string`

GetUserAction returns the UserAction field if non-nil, zero value otherwise.

### GetUserActionOk

`func (o *Event) GetUserActionOk() (*string, bool)`

GetUserActionOk returns a tuple with the UserAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAction

`func (o *Event) SetUserAction(v string)`

SetUserAction sets UserAction field to given value.

### HasUserAction

`func (o *Event) HasUserAction() bool`

HasUserAction returns a boolean if a field has been set.

### GetMinimumProcessGroupInstanceCountThreshold

`func (o *Event) GetMinimumProcessGroupInstanceCountThreshold() int32`

GetMinimumProcessGroupInstanceCountThreshold returns the MinimumProcessGroupInstanceCountThreshold field if non-nil, zero value otherwise.

### GetMinimumProcessGroupInstanceCountThresholdOk

`func (o *Event) GetMinimumProcessGroupInstanceCountThresholdOk() (*int32, bool)`

GetMinimumProcessGroupInstanceCountThresholdOk returns a tuple with the MinimumProcessGroupInstanceCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumProcessGroupInstanceCountThreshold

`func (o *Event) SetMinimumProcessGroupInstanceCountThreshold(v int32)`

SetMinimumProcessGroupInstanceCountThreshold sets MinimumProcessGroupInstanceCountThreshold field to given value.

### HasMinimumProcessGroupInstanceCountThreshold

`func (o *Event) HasMinimumProcessGroupInstanceCountThreshold() bool`

HasMinimumProcessGroupInstanceCountThreshold returns a boolean if a field has been set.

### GetReferenceResponseTime50thPercentile

`func (o *Event) GetReferenceResponseTime50thPercentile() float32`

GetReferenceResponseTime50thPercentile returns the ReferenceResponseTime50thPercentile field if non-nil, zero value otherwise.

### GetReferenceResponseTime50thPercentileOk

`func (o *Event) GetReferenceResponseTime50thPercentileOk() (*float32, bool)`

GetReferenceResponseTime50thPercentileOk returns a tuple with the ReferenceResponseTime50thPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceResponseTime50thPercentile

`func (o *Event) SetReferenceResponseTime50thPercentile(v float32)`

SetReferenceResponseTime50thPercentile sets ReferenceResponseTime50thPercentile field to given value.

### HasReferenceResponseTime50thPercentile

`func (o *Event) HasReferenceResponseTime50thPercentile() bool`

HasReferenceResponseTime50thPercentile returns a boolean if a field has been set.

### GetOriginal

`func (o *Event) GetOriginal() string`

GetOriginal returns the Original field if non-nil, zero value otherwise.

### GetOriginalOk

`func (o *Event) GetOriginalOk() (*string, bool)`

GetOriginalOk returns a tuple with the Original field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginal

`func (o *Event) SetOriginal(v string)`

SetOriginal sets Original field to given value.

### HasOriginal

`func (o *Event) HasOriginal() bool`

HasOriginal returns a boolean if a field has been set.

### GetUserDefined90thPercentileThreshold

`func (o *Event) GetUserDefined90thPercentileThreshold() float32`

GetUserDefined90thPercentileThreshold returns the UserDefined90thPercentileThreshold field if non-nil, zero value otherwise.

### GetUserDefined90thPercentileThresholdOk

`func (o *Event) GetUserDefined90thPercentileThresholdOk() (*float32, bool)`

GetUserDefined90thPercentileThresholdOk returns a tuple with the UserDefined90thPercentileThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefined90thPercentileThreshold

`func (o *Event) SetUserDefined90thPercentileThreshold(v float32)`

SetUserDefined90thPercentileThreshold sets UserDefined90thPercentileThreshold field to given value.

### HasUserDefined90thPercentileThreshold

`func (o *Event) HasUserDefined90thPercentileThreshold() bool`

HasUserDefined90thPercentileThreshold returns a boolean if a field has been set.

### GetDeploymentVersion

`func (o *Event) GetDeploymentVersion() string`

GetDeploymentVersion returns the DeploymentVersion field if non-nil, zero value otherwise.

### GetDeploymentVersionOk

`func (o *Event) GetDeploymentVersionOk() (*string, bool)`

GetDeploymentVersionOk returns a tuple with the DeploymentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentVersion

`func (o *Event) SetDeploymentVersion(v string)`

SetDeploymentVersion sets DeploymentVersion field to given value.

### HasDeploymentVersion

`func (o *Event) HasDeploymentVersion() bool`

HasDeploymentVersion returns a boolean if a field has been set.

### GetAnnotationType

`func (o *Event) GetAnnotationType() string`

GetAnnotationType returns the AnnotationType field if non-nil, zero value otherwise.

### GetAnnotationTypeOk

`func (o *Event) GetAnnotationTypeOk() (*string, bool)`

GetAnnotationTypeOk returns a tuple with the AnnotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationType

`func (o *Event) SetAnnotationType(v string)`

SetAnnotationType sets AnnotationType field to given value.

### HasAnnotationType

`func (o *Event) HasAnnotationType() bool`

HasAnnotationType returns a boolean if a field has been set.

### GetAffectedSyntheticActions

`func (o *Event) GetAffectedSyntheticActions() []string`

GetAffectedSyntheticActions returns the AffectedSyntheticActions field if non-nil, zero value otherwise.

### GetAffectedSyntheticActionsOk

`func (o *Event) GetAffectedSyntheticActionsOk() (*[]string, bool)`

GetAffectedSyntheticActionsOk returns a tuple with the AffectedSyntheticActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedSyntheticActions

`func (o *Event) SetAffectedSyntheticActions(v []string)`

SetAffectedSyntheticActions sets AffectedSyntheticActions field to given value.

### HasAffectedSyntheticActions

`func (o *Event) HasAffectedSyntheticActions() bool`

HasAffectedSyntheticActions returns a boolean if a field has been set.

### GetAffectedUserActionsPerMinute

`func (o *Event) GetAffectedUserActionsPerMinute() float32`

GetAffectedUserActionsPerMinute returns the AffectedUserActionsPerMinute field if non-nil, zero value otherwise.

### GetAffectedUserActionsPerMinuteOk

`func (o *Event) GetAffectedUserActionsPerMinuteOk() (*float32, bool)`

GetAffectedUserActionsPerMinuteOk returns a tuple with the AffectedUserActionsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedUserActionsPerMinute

`func (o *Event) SetAffectedUserActionsPerMinute(v float32)`

SetAffectedUserActionsPerMinute sets AffectedUserActionsPerMinute field to given value.

### HasAffectedUserActionsPerMinute

`func (o *Event) HasAffectedUserActionsPerMinute() bool`

HasAffectedUserActionsPerMinute returns a boolean if a field has been set.

### GetActiveMaintenanceWindows

`func (o *Event) GetActiveMaintenanceWindows() []EntityShortRepresentation`

GetActiveMaintenanceWindows returns the ActiveMaintenanceWindows field if non-nil, zero value otherwise.

### GetActiveMaintenanceWindowsOk

`func (o *Event) GetActiveMaintenanceWindowsOk() (*[]EntityShortRepresentation, bool)`

GetActiveMaintenanceWindowsOk returns a tuple with the ActiveMaintenanceWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveMaintenanceWindows

`func (o *Event) SetActiveMaintenanceWindows(v []EntityShortRepresentation)`

SetActiveMaintenanceWindows sets ActiveMaintenanceWindows field to given value.

### HasActiveMaintenanceWindows

`func (o *Event) HasActiveMaintenanceWindows() bool`

HasActiveMaintenanceWindows returns a boolean if a field has been set.

### GetMobileAppVersion

`func (o *Event) GetMobileAppVersion() string`

GetMobileAppVersion returns the MobileAppVersion field if non-nil, zero value otherwise.

### GetMobileAppVersionOk

`func (o *Event) GetMobileAppVersionOk() (*string, bool)`

GetMobileAppVersionOk returns a tuple with the MobileAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileAppVersion

`func (o *Event) SetMobileAppVersion(v string)`

SetMobileAppVersion sets MobileAppVersion field to given value.

### HasMobileAppVersion

`func (o *Event) HasMobileAppVersion() bool`

HasMobileAppVersion returns a boolean if a field has been set.

### GetUserDefinedFailureRateThreshold

`func (o *Event) GetUserDefinedFailureRateThreshold() float32`

GetUserDefinedFailureRateThreshold returns the UserDefinedFailureRateThreshold field if non-nil, zero value otherwise.

### GetUserDefinedFailureRateThresholdOk

`func (o *Event) GetUserDefinedFailureRateThresholdOk() (*float32, bool)`

GetUserDefinedFailureRateThresholdOk returns a tuple with the UserDefinedFailureRateThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFailureRateThreshold

`func (o *Event) SetUserDefinedFailureRateThreshold(v float32)`

SetUserDefinedFailureRateThreshold sets UserDefinedFailureRateThreshold field to given value.

### HasUserDefinedFailureRateThreshold

`func (o *Event) HasUserDefinedFailureRateThreshold() bool`

HasUserDefinedFailureRateThreshold returns a boolean if a field has been set.

### GetPercentile

`func (o *Event) GetPercentile() string`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *Event) GetPercentileOk() (*string, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *Event) SetPercentile(v string)`

SetPercentile sets Percentile field to given value.

### HasPercentile

`func (o *Event) HasPercentile() bool`

HasPercentile returns a boolean if a field has been set.

### GetCustomProperties

`func (o *Event) GetCustomProperties() map[string]string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *Event) GetCustomPropertiesOk() (*map[string]string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *Event) SetCustomProperties(v map[string]string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *Event) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetRemediationAction

`func (o *Event) GetRemediationAction() string`

GetRemediationAction returns the RemediationAction field if non-nil, zero value otherwise.

### GetRemediationActionOk

`func (o *Event) GetRemediationActionOk() (*string, bool)`

GetRemediationActionOk returns a tuple with the RemediationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationAction

`func (o *Event) SetRemediationAction(v string)`

SetRemediationAction sets RemediationAction field to given value.

### HasRemediationAction

`func (o *Event) HasRemediationAction() bool`

HasRemediationAction returns a boolean if a field has been set.

### GetService

`func (o *Event) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Event) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Event) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *Event) HasService() bool`

HasService returns a boolean if a field has been set.

### GetCiBackLink

`func (o *Event) GetCiBackLink() string`

GetCiBackLink returns the CiBackLink field if non-nil, zero value otherwise.

### GetCiBackLinkOk

`func (o *Event) GetCiBackLinkOk() (*string, bool)`

GetCiBackLinkOk returns a tuple with the CiBackLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiBackLink

`func (o *Event) SetCiBackLink(v string)`

SetCiBackLink sets CiBackLink field to given value.

### HasCiBackLink

`func (o *Event) HasCiBackLink() bool`

HasCiBackLink returns a boolean if a field has been set.

### GetGeolocation

`func (o *Event) GetGeolocation() string`

GetGeolocation returns the Geolocation field if non-nil, zero value otherwise.

### GetGeolocationOk

`func (o *Event) GetGeolocationOk() (*string, bool)`

GetGeolocationOk returns a tuple with the Geolocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeolocation

`func (o *Event) SetGeolocation(v string)`

SetGeolocation sets Geolocation field to given value.

### HasGeolocation

`func (o *Event) HasGeolocation() bool`

HasGeolocation returns a boolean if a field has been set.

### GetServiceMethod

`func (o *Event) GetServiceMethod() string`

GetServiceMethod returns the ServiceMethod field if non-nil, zero value otherwise.

### GetServiceMethodOk

`func (o *Event) GetServiceMethodOk() (*string, bool)`

GetServiceMethodOk returns a tuple with the ServiceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMethod

`func (o *Event) SetServiceMethod(v string)`

SetServiceMethod sets ServiceMethod field to given value.

### HasServiceMethod

`func (o *Event) HasServiceMethod() bool`

HasServiceMethod returns a boolean if a field has been set.

### GetChanged

`func (o *Event) GetChanged() string`

GetChanged returns the Changed field if non-nil, zero value otherwise.

### GetChangedOk

`func (o *Event) GetChangedOk() (*string, bool)`

GetChangedOk returns a tuple with the Changed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanged

`func (o *Event) SetChanged(v string)`

SetChanged sets Changed field to given value.

### HasChanged

`func (o *Event) HasChanged() bool`

HasChanged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


