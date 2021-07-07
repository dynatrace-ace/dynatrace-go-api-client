# \MonitoredEntitiesApi

All URIs are relative to *http://https:/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEntities**](MonitoredEntitiesApi.md#GetEntities) | **Get** /entities | Gets the information about monitored entities
[**GetEntity**](MonitoredEntitiesApi.md#GetEntity) | **Get** /entities/{entityId} | Gets the properties of the specified monitored entity
[**GetEntityType**](MonitoredEntitiesApi.md#GetEntityType) | **Get** /entityTypes/{type} | Gets a list of properties for the specified entity type
[**GetEntityTypes**](MonitoredEntitiesApi.md#GetEntityTypes) | **Get** /entityTypes | Gets a list of properties for all entity types
[**PushCustomDevice**](MonitoredEntitiesApi.md#PushCustomDevice) | **Post** /entities/custom | Creates or updates a custom device



## GetEntities

> EntitiesList GetEntities(ctx).NextPageKey(nextPageKey).PageSize(pageSize).EntitySelector(entitySelector).From(from).To(to).Fields(fields).Sort(sort).Execute()

Gets the information about monitored entities



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
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of entities.   If not set, 50 is used. (optional)
    entitySelector := "entitySelector_example" // string | Defines the scope of the query. Only entities matching the specified criteria are included into response.   You must set one of these criteria:   * Entity type: `type(\"TYPE\")`  * Dynatrace entity ID: `entityId(\"id\")`. You can specify several IDs, separated by a comma (`entityId(\"id-1\",\"id-2\")`).   You can add one or several of the following criteria. Values are case-sensitive and the `EQUALS` operator is used unless otherwise specified.   * Tag: `tag(\"value\")`. Tags in `[context]key:value`, `key:value`, and `value` formats are detected and parsed automatically. Any colons (`:`) that are part of the key or value must be escaped with a backslash(`\\`), otherwise, it will be interpreted as the separator between the key and the value. All tag values are case-sensitive.  * Management zone ID: `mzId(123)`  * Management zone name: `mzName(\"value\")` * Entity name: `entityName(\"value\")`. By default this filters for entities, whose name **contains** the given value and is **not** case sensitive.   The following modifications are available:  * `entityName.equals`: changes the operator to `EQUALS`.   * `entityName.startsWith`: changes the operator to `BEGINS WITH`.   * `entityName.in`: enables you to provide multiple values. The `EQUALS` operator applies.   * `caseSensitive(entityName(\"value\"))`: takes any entity name criterion as an arguments and makes the value case-sensitive. * Health state (HEALTHY,UNHEALTHY): `healthState(\"HEALTHY\")` * First seen timestamp: `firstSeenTms.<operator>(now-3h)`. Use any timestamp format from the **from**_/_**to** parameters.   The following operators are available:  * `lte`: earlier than or at the specified time  * `lt`: earlier than the specified time  * `gte`: later than or at the specified time  * `gt`: later than the specified time * Entity attribute: `<attribute>(\"value1\",\"value2\")` and `<attribute>.exists()`. To fetch the list of available attributes, execute the [GET all entity types](https://dt-url.net/dw03s7h) and check the **properties** field.  * Relationships: `fromRelationships.<relationshipName>()` and `toRelationships.<relationshipName>()`. The criterion takes an entity selector as an attribute. To fetch the list of available relationships, execute the [GET all entity types](https://dt-url.net/dw03s7h) and check the **fromRelationships** and **toRelationships** fields. * Negation: `not(<criterion>)`. Inverts any criterion except for **type**.   For more information, see [Entity selector](https://dt-url.net/apientityselector) in Dynatrace Documentation.   To set several criteria, separate them with a comma (`,`). For example, `type(\"HOST\"),healthState(\"HEALTHY\")`. Only results matching **all** criteria are included in response.   The length of the string is limited to 10,000 characters.   The field is **required** when you're querying the first page of results. (optional)
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of three days is used (`now-3d`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    fields := "fields_example" // string | Defines the list of entity properties included in the response. The ID and the name of an entity are **always** included to the response.   To add properties, list them with leading plus `+`. You can specify several properties, separated by a comma (for example `fields=+lastSeenTms,+properties.BITNESS`).   Use the [GET entity type](https://dt-url.net/2ka3ivt) request to fetch the list of properties available for your entity type. Fields from the **properties** object must be specified in the `properties.FIELD` format (for example, `properties.BITNESS`). (optional)
    sort := "sort_example" // string | Defines the ordering of the entities returned.   This field is **optional**, each field has a sign prefix (+/-), which corresponds to sorting order ( + for ascending and - for descending). If no sign prefix is set, then default ascending sorting order will be applied.   Currently ordering is only available for the display name (for example `sort=name or sort =+name for ascending, sort=-name for descending`) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitoredEntitiesApi.GetEntities(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).EntitySelector(entitySelector).From(from).To(to).Fields(fields).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoredEntitiesApi.GetEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntities`: EntitiesList
    fmt.Fprintf(os.Stdout, "Response from `MonitoredEntitiesApi.GetEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of entities.   If not set, 50 is used. | 
 **entitySelector** | **string** | Defines the scope of the query. Only entities matching the specified criteria are included into response.   You must set one of these criteria:   * Entity type: &#x60;type(\&quot;TYPE\&quot;)&#x60;  * Dynatrace entity ID: &#x60;entityId(\&quot;id\&quot;)&#x60;. You can specify several IDs, separated by a comma (&#x60;entityId(\&quot;id-1\&quot;,\&quot;id-2\&quot;)&#x60;).   You can add one or several of the following criteria. Values are case-sensitive and the &#x60;EQUALS&#x60; operator is used unless otherwise specified.   * Tag: &#x60;tag(\&quot;value\&quot;)&#x60;. Tags in &#x60;[context]key:value&#x60;, &#x60;key:value&#x60;, and &#x60;value&#x60; formats are detected and parsed automatically. Any colons (&#x60;:&#x60;) that are part of the key or value must be escaped with a backslash(&#x60;\\&#x60;), otherwise, it will be interpreted as the separator between the key and the value. All tag values are case-sensitive.  * Management zone ID: &#x60;mzId(123)&#x60;  * Management zone name: &#x60;mzName(\&quot;value\&quot;)&#x60; * Entity name: &#x60;entityName(\&quot;value\&quot;)&#x60;. By default this filters for entities, whose name **contains** the given value and is **not** case sensitive.   The following modifications are available:  * &#x60;entityName.equals&#x60;: changes the operator to &#x60;EQUALS&#x60;.   * &#x60;entityName.startsWith&#x60;: changes the operator to &#x60;BEGINS WITH&#x60;.   * &#x60;entityName.in&#x60;: enables you to provide multiple values. The &#x60;EQUALS&#x60; operator applies.   * &#x60;caseSensitive(entityName(\&quot;value\&quot;))&#x60;: takes any entity name criterion as an arguments and makes the value case-sensitive. * Health state (HEALTHY,UNHEALTHY): &#x60;healthState(\&quot;HEALTHY\&quot;)&#x60; * First seen timestamp: &#x60;firstSeenTms.&lt;operator&gt;(now-3h)&#x60;. Use any timestamp format from the **from**_/_**to** parameters.   The following operators are available:  * &#x60;lte&#x60;: earlier than or at the specified time  * &#x60;lt&#x60;: earlier than the specified time  * &#x60;gte&#x60;: later than or at the specified time  * &#x60;gt&#x60;: later than the specified time * Entity attribute: &#x60;&lt;attribute&gt;(\&quot;value1\&quot;,\&quot;value2\&quot;)&#x60; and &#x60;&lt;attribute&gt;.exists()&#x60;. To fetch the list of available attributes, execute the [GET all entity types](https://dt-url.net/dw03s7h) and check the **properties** field.  * Relationships: &#x60;fromRelationships.&lt;relationshipName&gt;()&#x60; and &#x60;toRelationships.&lt;relationshipName&gt;()&#x60;. The criterion takes an entity selector as an attribute. To fetch the list of available relationships, execute the [GET all entity types](https://dt-url.net/dw03s7h) and check the **fromRelationships** and **toRelationships** fields. * Negation: &#x60;not(&lt;criterion&gt;)&#x60;. Inverts any criterion except for **type**.   For more information, see [Entity selector](https://dt-url.net/apientityselector) in Dynatrace Documentation.   To set several criteria, separate them with a comma (&#x60;,&#x60;). For example, &#x60;type(\&quot;HOST\&quot;),healthState(\&quot;HEALTHY\&quot;)&#x60;. Only results matching **all** criteria are included in response.   The length of the string is limited to 10,000 characters.   The field is **required** when you&#39;re querying the first page of results. | 
 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of three days is used (&#x60;now-3d&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **fields** | **string** | Defines the list of entity properties included in the response. The ID and the name of an entity are **always** included to the response.   To add properties, list them with leading plus &#x60;+&#x60;. You can specify several properties, separated by a comma (for example &#x60;fields&#x3D;+lastSeenTms,+properties.BITNESS&#x60;).   Use the [GET entity type](https://dt-url.net/2ka3ivt) request to fetch the list of properties available for your entity type. Fields from the **properties** object must be specified in the &#x60;properties.FIELD&#x60; format (for example, &#x60;properties.BITNESS&#x60;). | 
 **sort** | **string** | Defines the ordering of the entities returned.   This field is **optional**, each field has a sign prefix (+/-), which corresponds to sorting order ( + for ascending and - for descending). If no sign prefix is set, then default ascending sorting order will be applied.   Currently ordering is only available for the display name (for example &#x60;sort&#x3D;name or sort &#x3D;+name for ascending, sort&#x3D;-name for descending&#x60;) | 

### Return type

[**EntitiesList**](EntitiesList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntity

> Entity GetEntity(ctx, entityId).From(from).To(to).Fields(fields).Execute()

Gets the properties of the specified monitored entity

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
    entityId := "entityId_example" // string | The ID of the required entity.
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of three days is used (`now-3d`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    fields := "fields_example" // string | Defines the list of entity properties included in the response. The ID and the name of an entity are **always** included to the response.   To add properties, list them with leading plus `+`. You can specify several properties, separated by a comma (for example `fields=+lastSeenTms,+properties.BITNESS`).   Use the [GET entity type](https://dt-url.net/2ka3ivt) request to fetch the list of properties available for your entity type. Fields from the **properties** object must be specified in the `properties.FIELD` format (for example, `properties.BITNESS`). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitoredEntitiesApi.GetEntity(context.Background(), entityId).From(from).To(to).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoredEntitiesApi.GetEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntity`: Entity
    fmt.Fprintf(os.Stdout, "Response from `MonitoredEntitiesApi.GetEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | The ID of the required entity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of three days is used (&#x60;now-3d&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **fields** | **string** | Defines the list of entity properties included in the response. The ID and the name of an entity are **always** included to the response.   To add properties, list them with leading plus &#x60;+&#x60;. You can specify several properties, separated by a comma (for example &#x60;fields&#x3D;+lastSeenTms,+properties.BITNESS&#x60;).   Use the [GET entity type](https://dt-url.net/2ka3ivt) request to fetch the list of properties available for your entity type. Fields from the **properties** object must be specified in the &#x60;properties.FIELD&#x60; format (for example, &#x60;properties.BITNESS&#x60;). | 

### Return type

[**Entity**](Entity.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntityType

> EntityType GetEntityType(ctx, type_).Execute()

Gets a list of properties for the specified entity type

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
    type_ := "type__example" // string | The required entity type.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitoredEntitiesApi.GetEntityType(context.Background(), type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoredEntitiesApi.GetEntityType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntityType`: EntityType
    fmt.Fprintf(os.Stdout, "Response from `MonitoredEntitiesApi.GetEntityType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The required entity type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntityType**](EntityType.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntityTypes

> EntityTypeList GetEntityTypes(ctx).NextPageKey(nextPageKey).PageSize(pageSize).Execute()

Gets a list of properties for all entity types



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
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of entity types in a single response payload.   The maximal allowed page size is 500.   If not set, 50 is used. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitoredEntitiesApi.GetEntityTypes(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoredEntitiesApi.GetEntityTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntityTypes`: EntityTypeList
    fmt.Fprintf(os.Stdout, "Response from `MonitoredEntitiesApi.GetEntityTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of entity types in a single response payload.   The maximal allowed page size is 500.   If not set, 50 is used. | 

### Return type

[**EntityTypeList**](EntityTypeList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PushCustomDevice

> CustomDeviceCreationResult PushCustomDevice(ctx).CustomDeviceCreation(customDeviceCreation).Execute()

Creates or updates a custom device

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
    customDeviceCreation := *openapiclient.NewCustomDeviceCreation("CustomDeviceId_example", "DisplayName_example") // CustomDeviceCreation | The JSON body of the request. Contains parameters of a custom device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitoredEntitiesApi.PushCustomDevice(context.Background()).CustomDeviceCreation(customDeviceCreation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoredEntitiesApi.PushCustomDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PushCustomDevice`: CustomDeviceCreationResult
    fmt.Fprintf(os.Stdout, "Response from `MonitoredEntitiesApi.PushCustomDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPushCustomDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customDeviceCreation** | [**CustomDeviceCreation**](CustomDeviceCreation.md) | The JSON body of the request. Contains parameters of a custom device. | 

### Return type

[**CustomDeviceCreationResult**](CustomDeviceCreationResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

