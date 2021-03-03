# \OneAgentOnAHostApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHostsWithSpecificAgents**](OneAgentOnAHostApi.md#GetHostsWithSpecificAgents) | **Get** /oneagents | Gets the list of hosts with OneAgent deployment information for each host



## GetHostsWithSpecificAgents

> HostsListPage GetHostsWithSpecificAgents(ctx).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).RelativeTime(relativeTime).Tag(tag).Entity(entity).ManagementZoneId(managementZoneId).ManagementZone(managementZone).NetworkZoneId(networkZoneId).HostGroupId(hostGroupId).HostGroupName(hostGroupName).OsType(osType).AvailabilityState(availabilityState).MonitoringType(monitoringType).AgentVersionIs(agentVersionIs).AgentVersionNumber(agentVersionNumber).AutoUpdateSetting(autoUpdateSetting).UpdateStatus(updateStatus).FaultyVersion(faultyVersion).TechnologyModuleType(technologyModuleType).TechnologyModuleVersionIs(technologyModuleVersionIs).TechnologyModuleVersionNumber(technologyModuleVersionNumber).TechnologyModuleFaultyVersion(technologyModuleFaultyVersion).PluginName(pluginName).PluginVersionIs(pluginVersionIs).PluginVersionNumber(pluginVersionNumber).PluginState(pluginState).NextPageKey(nextPageKey).Execute()

Gets the list of hosts with OneAgent deployment information for each host



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    startTimestamp := int64(789) // int64 | The start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then 72 hours behind from now is used. (optional)
    endTimestamp := int64(789) // int64 | The end timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then the current timestamp is used.   The timeframe must not exceed 3 months (92 days). (optional)
    relativeTime := "relativeTime_example" // string | The relative timeframe, back from now.   If you need to specify relative timeframe that is not presented in the list of possible values, specify the **startTimestamp** (up to 92 days back from now) and leave **endTimestamp** and **relativeTime** empty. (optional)
    tag := []string{"Inner_example"} // []string | Filters the resulting set of hosts by the specified tag. You can specify several tags in the following format: `tag=tag1&tag=tag2`. The host has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: `tag=[context]key:value`. For custom key-value tags, omit the context: `tag=key:value`. (optional)
    entity := []string{"Inner_example"} // []string | Filters result to the specified hosts only.    To specify several hosts use the following format: `entity=ID1&entity=ID2`. (optional)
    managementZoneId := int64(789) // int64 | Only return hosts that are part of the specified management zone.   Specify the management zone ID here. (optional)
    managementZone := "managementZone_example" // string | Only return hosts that are part of the specified management zone.   Specify the management zone name here.   If the **managementZoneId** parameter is set, this parameter is ignored. (optional)
    networkZoneId := "networkZoneId_example" // string | Filters the resulting set of hosts by the specified network zone.    Specify the Dynatrace entity ID of the required network zone. You can fetch the list of available network zones with the [GET all network zones](https://dt-url.net/u4o3r7z) call. (optional)
    hostGroupId := "hostGroupId_example" // string | Filters the resulting set of hosts by the specified host group.    Specify the Dynatrace entity ID of the required host group. (optional)
    hostGroupName := "hostGroupName_example" // string | Filters the resulting set of hosts by the specified host group.    Specify the name of the required host group. (optional)
    osType := "osType_example" // string | Filters the resulting set of hosts by the OS type. (optional)
    availabilityState := "availabilityState_example" // string | Filters the resulting set of hosts by the availability state of OneAgent.   * `MONITORED`: Hosts where OneAgent is enabled and active. * `UNMONITORED`: Hosts where OneAgent is disabled and inactive. * `CRASHED`: Hosts where OneAgent has returned a crash status code. * `LOST`: Hosts where it is impossible to establish connection with OneAgent. * `PRE_MONITORED`: Hosts where OneAgent is being initialized for monitoring. * `SHUTDOWN`: Hosts where OneAgent is shutting down in a controlled process. * `UNEXPECTED_SHUTDOWN`: Hosts where OneAgent is shutting down in an uncontrolled process. * `UNKNOWN`: Hosts where the state of OneAgent is unknown. (optional)
    monitoringType := "monitoringType_example" // string | Filters the resulting set of hosts by monitoring mode of OneAgent deployed on the host. (optional)
    agentVersionIs := "agentVersionIs_example" // string | Filters the resulting set of hosts to those that have a certain OneAgent version deployed on the host.   Specify the comparison operator here. (optional)
    agentVersionNumber := "agentVersionNumber_example" // string | Filters the resulting set of hosts to those that have a certain OneAgent version deployed on the host.   Specify the version in the `<major>.<minor>.<revision>` format, for example `1.182.0`. You can fetch the list of available versions with the [GET available versions](https://dt-url.net/fo23rb5) call. (optional)
    autoUpdateSetting := "autoUpdateSetting_example" // string | Filters the resulting set of hosts by the actual state of the auto-update setting of deployed OneAgents. (optional)
    updateStatus := "updateStatus_example" // string | Filters the resulting set of hosts by the update status of OneAgent deployed on the host. (optional)
    faultyVersion := true // bool | Filters the resulting set of hosts to those that run OneAgent version that is marked as faulty. (optional)
    technologyModuleType := "technologyModuleType_example" // string | Filters the resulting set of hosts to those that run the specified OneAgent code module.   If several code module filters are set, the code module has to match **all** the filters. (optional)
    technologyModuleVersionIs := "technologyModuleVersionIs_example" // string | Filters the resulting set of hosts to those that have a certain code module version deployed on the host.   Specify the comparison operator here.   If several code module filters are set, the code module has to match **all** the filters. (optional)
    technologyModuleVersionNumber := "technologyModuleVersionNumber_example" // string | Filters the resulting set of hosts to those that have a certain code module version deployed on the host.   Specify the version in the `<major>.<minor>.<revision>` format, for example `1.182.0`. You can fetch the list of available versions with the [GET available versions](https://dt-url.net/fo23rb5) call.   If several code module filters are set, the code module has to match **all** the filters. (optional)
    technologyModuleFaultyVersion := true // bool | Filters the resulting set of hosts to those that run the code module version that is marked as faulty.   If several code module filters are set, the code module has to match **all** the filters. (optional)
    pluginName := "pluginName_example" // string | Filters the resulting set of hosts to those that run the plugin with the specified name.   The **CONTAINS** operator is applied to the specified value.   If several plugin filters are set, the plugin has to match **all** the filters. (optional)
    pluginVersionIs := "pluginVersionIs_example" // string | Filters the resulting set of hosts to those that have a certain plugin version deployed on the host.   Specify the comparison operator here.   If several plugin filters are set, the plugin has to match **all** the filters. (optional)
    pluginVersionNumber := "pluginVersionNumber_example" // string | Filters the resulting set of hosts to those that have a certain plugin version deployed on the host.   Specify the version in the `<major>.<minor>.<revision>` format, for example `1.182.0`. You can fetch the list of available versions with the [GET available versions](https://dt-url.net/fo23rb5) call.   `<minor>` and `<revision>` parts of the version number are optional.   If several plugin filters are set, the plugin has to match **all** the filters. (optional)
    pluginState := "pluginState_example" // string | Filters the resulting set of hosts to those that run the plugin with the specified state. (optional)
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify this cursor.   You must keep all other query parameters as they were in the first request to obtain subsequent pages. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OneAgentOnAHostApi.GetHostsWithSpecificAgents(context.Background()).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).RelativeTime(relativeTime).Tag(tag).Entity(entity).ManagementZoneId(managementZoneId).ManagementZone(managementZone).NetworkZoneId(networkZoneId).HostGroupId(hostGroupId).HostGroupName(hostGroupName).OsType(osType).AvailabilityState(availabilityState).MonitoringType(monitoringType).AgentVersionIs(agentVersionIs).AgentVersionNumber(agentVersionNumber).AutoUpdateSetting(autoUpdateSetting).UpdateStatus(updateStatus).FaultyVersion(faultyVersion).TechnologyModuleType(technologyModuleType).TechnologyModuleVersionIs(technologyModuleVersionIs).TechnologyModuleVersionNumber(technologyModuleVersionNumber).TechnologyModuleFaultyVersion(technologyModuleFaultyVersion).PluginName(pluginName).PluginVersionIs(pluginVersionIs).PluginVersionNumber(pluginVersionNumber).PluginState(pluginState).NextPageKey(nextPageKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostApi.GetHostsWithSpecificAgents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostsWithSpecificAgents`: HostsListPage
    fmt.Fprintf(os.Stdout, "Response from `OneAgentOnAHostApi.GetHostsWithSpecificAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostsWithSpecificAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimestamp** | **int64** | The start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then 72 hours behind from now is used. | 
 **endTimestamp** | **int64** | The end timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then the current timestamp is used.   The timeframe must not exceed 3 months (92 days). | 
 **relativeTime** | **string** | The relative timeframe, back from now.   If you need to specify relative timeframe that is not presented in the list of possible values, specify the **startTimestamp** (up to 92 days back from now) and leave **endTimestamp** and **relativeTime** empty. | 
 **tag** | **[]string** | Filters the resulting set of hosts by the specified tag. You can specify several tags in the following format: &#x60;tag&#x3D;tag1&amp;tag&#x3D;tag2&#x60;. The host has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: &#x60;tag&#x3D;[context]key:value&#x60;. For custom key-value tags, omit the context: &#x60;tag&#x3D;key:value&#x60;. | 
 **entity** | **[]string** | Filters result to the specified hosts only.    To specify several hosts use the following format: &#x60;entity&#x3D;ID1&amp;entity&#x3D;ID2&#x60;. | 
 **managementZoneId** | **int64** | Only return hosts that are part of the specified management zone.   Specify the management zone ID here. | 
 **managementZone** | **string** | Only return hosts that are part of the specified management zone.   Specify the management zone name here.   If the **managementZoneId** parameter is set, this parameter is ignored. | 
 **networkZoneId** | **string** | Filters the resulting set of hosts by the specified network zone.    Specify the Dynatrace entity ID of the required network zone. You can fetch the list of available network zones with the [GET all network zones](https://dt-url.net/u4o3r7z) call. | 
 **hostGroupId** | **string** | Filters the resulting set of hosts by the specified host group.    Specify the Dynatrace entity ID of the required host group. | 
 **hostGroupName** | **string** | Filters the resulting set of hosts by the specified host group.    Specify the name of the required host group. | 
 **osType** | **string** | Filters the resulting set of hosts by the OS type. | 
 **availabilityState** | **string** | Filters the resulting set of hosts by the availability state of OneAgent.   * &#x60;MONITORED&#x60;: Hosts where OneAgent is enabled and active. * &#x60;UNMONITORED&#x60;: Hosts where OneAgent is disabled and inactive. * &#x60;CRASHED&#x60;: Hosts where OneAgent has returned a crash status code. * &#x60;LOST&#x60;: Hosts where it is impossible to establish connection with OneAgent. * &#x60;PRE_MONITORED&#x60;: Hosts where OneAgent is being initialized for monitoring. * &#x60;SHUTDOWN&#x60;: Hosts where OneAgent is shutting down in a controlled process. * &#x60;UNEXPECTED_SHUTDOWN&#x60;: Hosts where OneAgent is shutting down in an uncontrolled process. * &#x60;UNKNOWN&#x60;: Hosts where the state of OneAgent is unknown. | 
 **monitoringType** | **string** | Filters the resulting set of hosts by monitoring mode of OneAgent deployed on the host. | 
 **agentVersionIs** | **string** | Filters the resulting set of hosts to those that have a certain OneAgent version deployed on the host.   Specify the comparison operator here. | 
 **agentVersionNumber** | **string** | Filters the resulting set of hosts to those that have a certain OneAgent version deployed on the host.   Specify the version in the &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;&#x60; format, for example &#x60;1.182.0&#x60;. You can fetch the list of available versions with the [GET available versions](https://dt-url.net/fo23rb5) call. | 
 **autoUpdateSetting** | **string** | Filters the resulting set of hosts by the actual state of the auto-update setting of deployed OneAgents. | 
 **updateStatus** | **string** | Filters the resulting set of hosts by the update status of OneAgent deployed on the host. | 
 **faultyVersion** | **bool** | Filters the resulting set of hosts to those that run OneAgent version that is marked as faulty. | 
 **technologyModuleType** | **string** | Filters the resulting set of hosts to those that run the specified OneAgent code module.   If several code module filters are set, the code module has to match **all** the filters. | 
 **technologyModuleVersionIs** | **string** | Filters the resulting set of hosts to those that have a certain code module version deployed on the host.   Specify the comparison operator here.   If several code module filters are set, the code module has to match **all** the filters. | 
 **technologyModuleVersionNumber** | **string** | Filters the resulting set of hosts to those that have a certain code module version deployed on the host.   Specify the version in the &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;&#x60; format, for example &#x60;1.182.0&#x60;. You can fetch the list of available versions with the [GET available versions](https://dt-url.net/fo23rb5) call.   If several code module filters are set, the code module has to match **all** the filters. | 
 **technologyModuleFaultyVersion** | **bool** | Filters the resulting set of hosts to those that run the code module version that is marked as faulty.   If several code module filters are set, the code module has to match **all** the filters. | 
 **pluginName** | **string** | Filters the resulting set of hosts to those that run the plugin with the specified name.   The **CONTAINS** operator is applied to the specified value.   If several plugin filters are set, the plugin has to match **all** the filters. | 
 **pluginVersionIs** | **string** | Filters the resulting set of hosts to those that have a certain plugin version deployed on the host.   Specify the comparison operator here.   If several plugin filters are set, the plugin has to match **all** the filters. | 
 **pluginVersionNumber** | **string** | Filters the resulting set of hosts to those that have a certain plugin version deployed on the host.   Specify the version in the &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;&#x60; format, for example &#x60;1.182.0&#x60;. You can fetch the list of available versions with the [GET available versions](https://dt-url.net/fo23rb5) call.   &#x60;&lt;minor&gt;&#x60; and &#x60;&lt;revision&gt;&#x60; parts of the version number are optional.   If several plugin filters are set, the plugin has to match **all** the filters. | 
 **pluginState** | **string** | Filters the resulting set of hosts to those that run the plugin with the specified state. | 
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify this cursor.   You must keep all other query parameters as they were in the first request to obtain subsequent pages. | 

### Return type

[**HostsListPage**](HostsListPage.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

