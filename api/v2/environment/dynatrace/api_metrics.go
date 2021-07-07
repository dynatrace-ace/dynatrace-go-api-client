/*
 * Dynatrace Environment API
 *
 *  Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// MetricsApiService MetricsApi service
type MetricsApiService service

type ApiAllMetricsRequest struct {
	ctx _context.Context
	ApiService *MetricsApiService
	nextPageKey *string
	pageSize *int64
	metricSelector *string
	text *string
	fields *string
	writtenSince *string
	metadataSelector *string
}

func (r ApiAllMetricsRequest) NextPageKey(nextPageKey string) ApiAllMetricsRequest {
	r.nextPageKey = &nextPageKey
	return r
}
func (r ApiAllMetricsRequest) PageSize(pageSize int64) ApiAllMetricsRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiAllMetricsRequest) MetricSelector(metricSelector string) ApiAllMetricsRequest {
	r.metricSelector = &metricSelector
	return r
}
func (r ApiAllMetricsRequest) Text(text string) ApiAllMetricsRequest {
	r.text = &text
	return r
}
func (r ApiAllMetricsRequest) Fields(fields string) ApiAllMetricsRequest {
	r.fields = &fields
	return r
}
func (r ApiAllMetricsRequest) WrittenSince(writtenSince string) ApiAllMetricsRequest {
	r.writtenSince = &writtenSince
	return r
}
func (r ApiAllMetricsRequest) MetadataSelector(metadataSelector string) ApiAllMetricsRequest {
	r.metadataSelector = &metadataSelector
	return r
}

func (r ApiAllMetricsRequest) Execute() (MetricDescriptorCollection, *_nethttp.Response, error) {
	return r.ApiService.AllMetricsExecute(r)
}

/*
 * AllMetrics Lists all available metrics
 * You can narrow down the output by selecting metrics in the **metricSelector** field. 

You can additionally limit the output by using pagination: 

1. Specify the number of results per page in the **pageSize** query parameter. 

2. Then use the cursor from the **nextPageKey** field of the response in the **nextPageKey** query parameter to obtain subsequent pages. All other query parameters must be omitted.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiAllMetricsRequest
 */
func (a *MetricsApiService) AllMetrics(ctx _context.Context) ApiAllMetricsRequest {
	return ApiAllMetricsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return MetricDescriptorCollection
 */
func (a *MetricsApiService) AllMetricsExecute(r ApiAllMetricsRequest) (MetricDescriptorCollection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MetricDescriptorCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.AllMetrics")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metrics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.nextPageKey != nil {
		localVarQueryParams.Add("nextPageKey", parameterToString(*r.nextPageKey, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.metricSelector != nil {
		localVarQueryParams.Add("metricSelector", parameterToString(*r.metricSelector, ""))
	}
	if r.text != nil {
		localVarQueryParams.Add("text", parameterToString(*r.text, ""))
	}
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	}
	if r.writtenSince != nil {
		localVarQueryParams.Add("writtenSince", parameterToString(*r.writtenSince, ""))
	}
	if r.metadataSelector != nil {
		localVarQueryParams.Add("metadataSelector", parameterToString(*r.metadataSelector, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=utf-8", "text/csv; header=present; charset=utf-8", "text/csv; header=absent; charset=utf-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteRequest struct {
	ctx _context.Context
	ApiService *MetricsApiService
	metricId string
}


func (r ApiDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
 * Delete Deletes the specified metric
 * Deletion cannot be undone! You can't delete a metric if it has data points ingested within the last two hours.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param metricId
 * @return ApiDeleteRequest
 */
func (a *MetricsApiService) Delete(ctx _context.Context, metricId string) ApiDeleteRequest {
	return ApiDeleteRequest{
		ApiService: a,
		ctx: ctx,
		metricId: metricId,
	}
}

/*
 * Execute executes the request
 */
func (a *MetricsApiService) DeleteExecute(r ApiDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.Delete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metrics/{metricId}"
	localVarPath = strings.Replace(localVarPath, "{"+"metricId"+"}", _neturl.PathEscape(parameterToString(r.metricId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiIngestRequest struct {
	ctx _context.Context
	ApiService *MetricsApiService
	body *string
}

func (r ApiIngestRequest) Body(body string) ApiIngestRequest {
	r.body = &body
	return r
}

func (r ApiIngestRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IngestExecute(r)
}

/*
 * Ingest Pushes metric data points to Dynatrace
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiIngestRequest
 */
func (a *MetricsApiService) Ingest(ctx _context.Context) ApiIngestRequest {
	return ApiIngestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *MetricsApiService) IngestExecute(r ApiIngestRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.Ingest")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metrics/ingest"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain; charset=utf-8"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiMetricRequest struct {
	ctx _context.Context
	ApiService *MetricsApiService
	metricId string
}


func (r ApiMetricRequest) Execute() (MetricDescriptor, *_nethttp.Response, error) {
	return r.ApiService.MetricExecute(r)
}

/*
 * Metric Gets the descriptor of the specified metric
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param metricId The key of the required metric.   You can set additional transformation operators, separated by a colon (`:`). See [Metrics selector transformations](https://dt-url.net/metricSelector) in Dynatrace Documentation for additional information on available result transformations.   The length of the string is limited to 5,000 characters.
 * @return ApiMetricRequest
 */
func (a *MetricsApiService) Metric(ctx _context.Context, metricId string) ApiMetricRequest {
	return ApiMetricRequest{
		ApiService: a,
		ctx: ctx,
		metricId: metricId,
	}
}

/*
 * Execute executes the request
 * @return MetricDescriptor
 */
func (a *MetricsApiService) MetricExecute(r ApiMetricRequest) (MetricDescriptor, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MetricDescriptor
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.Metric")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metrics/{metricId}"
	localVarPath = strings.Replace(localVarPath, "{"+"metricId"+"}", _neturl.PathEscape(parameterToString(r.metricId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=utf-8", "text/csv; header=present; charset=utf-8", "text/csv; header=absent; charset=utf-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiQueryRequest struct {
	ctx _context.Context
	ApiService *MetricsApiService
	metricSelector *string
	resolution *string
	from *string
	to *string
	entitySelector *string
}

func (r ApiQueryRequest) MetricSelector(metricSelector string) ApiQueryRequest {
	r.metricSelector = &metricSelector
	return r
}
func (r ApiQueryRequest) Resolution(resolution string) ApiQueryRequest {
	r.resolution = &resolution
	return r
}
func (r ApiQueryRequest) From(from string) ApiQueryRequest {
	r.from = &from
	return r
}
func (r ApiQueryRequest) To(to string) ApiQueryRequest {
	r.to = &to
	return r
}
func (r ApiQueryRequest) EntitySelector(entitySelector string) ApiQueryRequest {
	r.entitySelector = &entitySelector
	return r
}

func (r ApiQueryRequest) Execute() (MetricData, *_nethttp.Response, error) {
	return r.ApiService.QueryExecute(r)
}

/*
 * Query Gets data points of the specified metrics
 * The following limits apply: 
* The amount of aggregated data points in the response is limited to 1,000 
* The amount of series in the response is limited to 1,000 
   * The amount of data points per series is limited to 10,080 (minutes of one week) 
   * The overall amount of data points is limited to 100,000 

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiQueryRequest
 */
func (a *MetricsApiService) Query(ctx _context.Context) ApiQueryRequest {
	return ApiQueryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return MetricData
 */
func (a *MetricsApiService) QueryExecute(r ApiQueryRequest) (MetricData, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MetricData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.Query")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metrics/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.metricSelector != nil {
		localVarQueryParams.Add("metricSelector", parameterToString(*r.metricSelector, ""))
	}
	if r.resolution != nil {
		localVarQueryParams.Add("resolution", parameterToString(*r.resolution, ""))
	}
	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.entitySelector != nil {
		localVarQueryParams.Add("entitySelector", parameterToString(*r.entitySelector, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=utf-8", "text/csv; header=present; charset=utf-8", "text/csv; header=absent; charset=utf-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}