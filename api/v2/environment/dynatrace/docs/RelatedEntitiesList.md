# RelatedEntitiesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]RelatedEntity**](RelatedEntity.md) | A list of related applications. | [optional] [readonly] 
**Services** | Pointer to [**[]RelatedService**](RelatedService.md) | A list of related services. | [optional] [readonly] 
**Hosts** | Pointer to [**[]RelatedEntity**](RelatedEntity.md) | A list of related hosts. | [optional] [readonly] 
**Databases** | Pointer to **[]string** | A list of related databases. | [optional] [readonly] 
**KubernetesWorkloads** | Pointer to [**[]RelatedEntity**](RelatedEntity.md) | A list of related Kubernetes workloads. | [optional] [readonly] 
**KubernetesClusters** | Pointer to [**[]RelatedEntity**](RelatedEntity.md) | A list of related Kubernetes clusters. | [optional] [readonly] 

## Methods

### NewRelatedEntitiesList

`func NewRelatedEntitiesList() *RelatedEntitiesList`

NewRelatedEntitiesList instantiates a new RelatedEntitiesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedEntitiesListWithDefaults

`func NewRelatedEntitiesListWithDefaults() *RelatedEntitiesList`

NewRelatedEntitiesListWithDefaults instantiates a new RelatedEntitiesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *RelatedEntitiesList) GetApplications() []RelatedEntity`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *RelatedEntitiesList) GetApplicationsOk() (*[]RelatedEntity, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *RelatedEntitiesList) SetApplications(v []RelatedEntity)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *RelatedEntitiesList) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetServices

`func (o *RelatedEntitiesList) GetServices() []RelatedService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *RelatedEntitiesList) GetServicesOk() (*[]RelatedService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *RelatedEntitiesList) SetServices(v []RelatedService)`

SetServices sets Services field to given value.

### HasServices

`func (o *RelatedEntitiesList) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetHosts

`func (o *RelatedEntitiesList) GetHosts() []RelatedEntity`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *RelatedEntitiesList) GetHostsOk() (*[]RelatedEntity, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *RelatedEntitiesList) SetHosts(v []RelatedEntity)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *RelatedEntitiesList) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetDatabases

`func (o *RelatedEntitiesList) GetDatabases() []string`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *RelatedEntitiesList) GetDatabasesOk() (*[]string, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *RelatedEntitiesList) SetDatabases(v []string)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *RelatedEntitiesList) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### GetKubernetesWorkloads

`func (o *RelatedEntitiesList) GetKubernetesWorkloads() []RelatedEntity`

GetKubernetesWorkloads returns the KubernetesWorkloads field if non-nil, zero value otherwise.

### GetKubernetesWorkloadsOk

`func (o *RelatedEntitiesList) GetKubernetesWorkloadsOk() (*[]RelatedEntity, bool)`

GetKubernetesWorkloadsOk returns a tuple with the KubernetesWorkloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesWorkloads

`func (o *RelatedEntitiesList) SetKubernetesWorkloads(v []RelatedEntity)`

SetKubernetesWorkloads sets KubernetesWorkloads field to given value.

### HasKubernetesWorkloads

`func (o *RelatedEntitiesList) HasKubernetesWorkloads() bool`

HasKubernetesWorkloads returns a boolean if a field has been set.

### GetKubernetesClusters

`func (o *RelatedEntitiesList) GetKubernetesClusters() []RelatedEntity`

GetKubernetesClusters returns the KubernetesClusters field if non-nil, zero value otherwise.

### GetKubernetesClustersOk

`func (o *RelatedEntitiesList) GetKubernetesClustersOk() (*[]RelatedEntity, bool)`

GetKubernetesClustersOk returns a tuple with the KubernetesClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusters

`func (o *RelatedEntitiesList) SetKubernetesClusters(v []RelatedEntity)`

SetKubernetesClusters sets KubernetesClusters field to given value.

### HasKubernetesClusters

`func (o *RelatedEntitiesList) HasKubernetesClusters() bool`

HasKubernetesClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


