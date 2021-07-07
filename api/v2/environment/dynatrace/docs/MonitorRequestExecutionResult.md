# MonitorRequestExecutionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | Request id. | [optional] 
**RequestName** | Pointer to **string** | Request name. | [optional] 
**SequenceNumber** | Pointer to **int32** | Request&#39;s sequence number. | [optional] 
**StartTimestamp** | Pointer to **int64** | Request start timestamp. | [optional] 
**EngineId** | Pointer to **int64** | VUC&#39;s id on which monitor&#39;s request was executed. | [optional] 
**PublicLocation** | Pointer to **bool** | Flag informs whether request was executed on public location. | [optional] 
**Url** | Pointer to **string** | Request URL address. | [optional] 
**Method** | Pointer to **string** | Request method type. | [optional] 
**RequestBody** | Pointer to **string** | Request&#39;s request body. | [optional] 
**RequestHeaders** | Pointer to [**[]MonitorRequestHeader**](MonitorRequestHeader.md) | A list of request&#39;s headers | [optional] 
**ResponseStatusCode** | Pointer to **int32** | Request&#39;s response status code. | [optional] 
**ResponseBody** | Pointer to **string** | Request&#39;s response body. | [optional] 
**ResponseSize** | Pointer to **int64** | Request&#39;s response size in bytes. | [optional] 
**ResponseBodySizeLimitExceeded** | Pointer to **bool** | A flag indicating that the response payload size limit of 10MB has been exceeded. | [optional] 
**ResponseHeaders** | Pointer to [**[]MonitorRequestHeader**](MonitorRequestHeader.md) | A list of request&#39;s response headers | [optional] 
**ResolvedIps** | Pointer to **[]string** | Request&#39;s resolved ips.&#39; | [optional] 
**HealthStatusCode** | Pointer to **int32** | Request&#39;s health status code. | [optional] 
**HealthStatus** | Pointer to **string** | Request&#39;s health status. | [optional] 
**ResponseMessage** | Pointer to **string** | Request&#39;s response message.&#39; | [optional] 
**PeerCertificateExpiryDate** | Pointer to **int64** | An expiry date of the first SSL certificate from the certificate chain. | [optional] 
**PeerCertificateDetails** | Pointer to **string** | Request&#39;s certificate details. | [optional] 
**TotalTime** | Pointer to **int64** | A total request time measured in ms. | [optional] 
**HostNameResolutionTime** | Pointer to **int64** | A hostname resolution time measured in ms. | [optional] 
**TcpConnectTime** | Pointer to **int64** | A TCP connect time measured in ms. | [optional] 
**TlsHandshakeTime** | Pointer to **int64** | A TLS handshake time measured in ms. | [optional] 
**TimeToFirstByte** | Pointer to **int64** | A time to first byte measured in ms. | [optional] 
**RedirectionTime** | Pointer to **int64** | Total number of milliseconds spent on handling all redirect requests, measured in ms. | [optional] 
**WaitingTime** | Pointer to **int64** | Waiting time (time to first byte - (DNS lookup time + TCP connect time + TLS handshake time), measured in ms. | [optional] 
**RedirectsCount** | Pointer to **int32** | Number of request&#39;s redirects. | [optional] 
**FailureMessage** | Pointer to **string** | Request&#39;s failure message. | [optional] 

## Methods

### NewMonitorRequestExecutionResult

`func NewMonitorRequestExecutionResult() *MonitorRequestExecutionResult`

NewMonitorRequestExecutionResult instantiates a new MonitorRequestExecutionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorRequestExecutionResultWithDefaults

`func NewMonitorRequestExecutionResultWithDefaults() *MonitorRequestExecutionResult`

NewMonitorRequestExecutionResultWithDefaults instantiates a new MonitorRequestExecutionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *MonitorRequestExecutionResult) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *MonitorRequestExecutionResult) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *MonitorRequestExecutionResult) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *MonitorRequestExecutionResult) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestName

`func (o *MonitorRequestExecutionResult) GetRequestName() string`

GetRequestName returns the RequestName field if non-nil, zero value otherwise.

### GetRequestNameOk

`func (o *MonitorRequestExecutionResult) GetRequestNameOk() (*string, bool)`

GetRequestNameOk returns a tuple with the RequestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestName

`func (o *MonitorRequestExecutionResult) SetRequestName(v string)`

SetRequestName sets RequestName field to given value.

### HasRequestName

`func (o *MonitorRequestExecutionResult) HasRequestName() bool`

HasRequestName returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *MonitorRequestExecutionResult) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *MonitorRequestExecutionResult) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *MonitorRequestExecutionResult) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *MonitorRequestExecutionResult) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *MonitorRequestExecutionResult) GetStartTimestamp() int64`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *MonitorRequestExecutionResult) GetStartTimestampOk() (*int64, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *MonitorRequestExecutionResult) SetStartTimestamp(v int64)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *MonitorRequestExecutionResult) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetEngineId

`func (o *MonitorRequestExecutionResult) GetEngineId() int64`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *MonitorRequestExecutionResult) GetEngineIdOk() (*int64, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *MonitorRequestExecutionResult) SetEngineId(v int64)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *MonitorRequestExecutionResult) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetPublicLocation

`func (o *MonitorRequestExecutionResult) GetPublicLocation() bool`

GetPublicLocation returns the PublicLocation field if non-nil, zero value otherwise.

### GetPublicLocationOk

`func (o *MonitorRequestExecutionResult) GetPublicLocationOk() (*bool, bool)`

GetPublicLocationOk returns a tuple with the PublicLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicLocation

`func (o *MonitorRequestExecutionResult) SetPublicLocation(v bool)`

SetPublicLocation sets PublicLocation field to given value.

### HasPublicLocation

`func (o *MonitorRequestExecutionResult) HasPublicLocation() bool`

HasPublicLocation returns a boolean if a field has been set.

### GetUrl

`func (o *MonitorRequestExecutionResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MonitorRequestExecutionResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MonitorRequestExecutionResult) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MonitorRequestExecutionResult) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMethod

`func (o *MonitorRequestExecutionResult) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *MonitorRequestExecutionResult) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *MonitorRequestExecutionResult) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *MonitorRequestExecutionResult) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRequestBody

`func (o *MonitorRequestExecutionResult) GetRequestBody() string`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *MonitorRequestExecutionResult) GetRequestBodyOk() (*string, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *MonitorRequestExecutionResult) SetRequestBody(v string)`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *MonitorRequestExecutionResult) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *MonitorRequestExecutionResult) GetRequestHeaders() []MonitorRequestHeader`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *MonitorRequestExecutionResult) GetRequestHeadersOk() (*[]MonitorRequestHeader, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *MonitorRequestExecutionResult) SetRequestHeaders(v []MonitorRequestHeader)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *MonitorRequestExecutionResult) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### GetResponseStatusCode

`func (o *MonitorRequestExecutionResult) GetResponseStatusCode() int32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *MonitorRequestExecutionResult) GetResponseStatusCodeOk() (*int32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *MonitorRequestExecutionResult) SetResponseStatusCode(v int32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *MonitorRequestExecutionResult) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetResponseBody

`func (o *MonitorRequestExecutionResult) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *MonitorRequestExecutionResult) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *MonitorRequestExecutionResult) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *MonitorRequestExecutionResult) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetResponseSize

`func (o *MonitorRequestExecutionResult) GetResponseSize() int64`

GetResponseSize returns the ResponseSize field if non-nil, zero value otherwise.

### GetResponseSizeOk

`func (o *MonitorRequestExecutionResult) GetResponseSizeOk() (*int64, bool)`

GetResponseSizeOk returns a tuple with the ResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSize

`func (o *MonitorRequestExecutionResult) SetResponseSize(v int64)`

SetResponseSize sets ResponseSize field to given value.

### HasResponseSize

`func (o *MonitorRequestExecutionResult) HasResponseSize() bool`

HasResponseSize returns a boolean if a field has been set.

### GetResponseBodySizeLimitExceeded

`func (o *MonitorRequestExecutionResult) GetResponseBodySizeLimitExceeded() bool`

GetResponseBodySizeLimitExceeded returns the ResponseBodySizeLimitExceeded field if non-nil, zero value otherwise.

### GetResponseBodySizeLimitExceededOk

`func (o *MonitorRequestExecutionResult) GetResponseBodySizeLimitExceededOk() (*bool, bool)`

GetResponseBodySizeLimitExceededOk returns a tuple with the ResponseBodySizeLimitExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBodySizeLimitExceeded

`func (o *MonitorRequestExecutionResult) SetResponseBodySizeLimitExceeded(v bool)`

SetResponseBodySizeLimitExceeded sets ResponseBodySizeLimitExceeded field to given value.

### HasResponseBodySizeLimitExceeded

`func (o *MonitorRequestExecutionResult) HasResponseBodySizeLimitExceeded() bool`

HasResponseBodySizeLimitExceeded returns a boolean if a field has been set.

### GetResponseHeaders

`func (o *MonitorRequestExecutionResult) GetResponseHeaders() []MonitorRequestHeader`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *MonitorRequestExecutionResult) GetResponseHeadersOk() (*[]MonitorRequestHeader, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *MonitorRequestExecutionResult) SetResponseHeaders(v []MonitorRequestHeader)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *MonitorRequestExecutionResult) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### GetResolvedIps

`func (o *MonitorRequestExecutionResult) GetResolvedIps() []string`

GetResolvedIps returns the ResolvedIps field if non-nil, zero value otherwise.

### GetResolvedIpsOk

`func (o *MonitorRequestExecutionResult) GetResolvedIpsOk() (*[]string, bool)`

GetResolvedIpsOk returns a tuple with the ResolvedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedIps

`func (o *MonitorRequestExecutionResult) SetResolvedIps(v []string)`

SetResolvedIps sets ResolvedIps field to given value.

### HasResolvedIps

`func (o *MonitorRequestExecutionResult) HasResolvedIps() bool`

HasResolvedIps returns a boolean if a field has been set.

### GetHealthStatusCode

`func (o *MonitorRequestExecutionResult) GetHealthStatusCode() int32`

GetHealthStatusCode returns the HealthStatusCode field if non-nil, zero value otherwise.

### GetHealthStatusCodeOk

`func (o *MonitorRequestExecutionResult) GetHealthStatusCodeOk() (*int32, bool)`

GetHealthStatusCodeOk returns a tuple with the HealthStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatusCode

`func (o *MonitorRequestExecutionResult) SetHealthStatusCode(v int32)`

SetHealthStatusCode sets HealthStatusCode field to given value.

### HasHealthStatusCode

`func (o *MonitorRequestExecutionResult) HasHealthStatusCode() bool`

HasHealthStatusCode returns a boolean if a field has been set.

### GetHealthStatus

`func (o *MonitorRequestExecutionResult) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *MonitorRequestExecutionResult) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *MonitorRequestExecutionResult) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *MonitorRequestExecutionResult) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### GetResponseMessage

`func (o *MonitorRequestExecutionResult) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *MonitorRequestExecutionResult) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *MonitorRequestExecutionResult) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.

### HasResponseMessage

`func (o *MonitorRequestExecutionResult) HasResponseMessage() bool`

HasResponseMessage returns a boolean if a field has been set.

### GetPeerCertificateExpiryDate

`func (o *MonitorRequestExecutionResult) GetPeerCertificateExpiryDate() int64`

GetPeerCertificateExpiryDate returns the PeerCertificateExpiryDate field if non-nil, zero value otherwise.

### GetPeerCertificateExpiryDateOk

`func (o *MonitorRequestExecutionResult) GetPeerCertificateExpiryDateOk() (*int64, bool)`

GetPeerCertificateExpiryDateOk returns a tuple with the PeerCertificateExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerCertificateExpiryDate

`func (o *MonitorRequestExecutionResult) SetPeerCertificateExpiryDate(v int64)`

SetPeerCertificateExpiryDate sets PeerCertificateExpiryDate field to given value.

### HasPeerCertificateExpiryDate

`func (o *MonitorRequestExecutionResult) HasPeerCertificateExpiryDate() bool`

HasPeerCertificateExpiryDate returns a boolean if a field has been set.

### GetPeerCertificateDetails

`func (o *MonitorRequestExecutionResult) GetPeerCertificateDetails() string`

GetPeerCertificateDetails returns the PeerCertificateDetails field if non-nil, zero value otherwise.

### GetPeerCertificateDetailsOk

`func (o *MonitorRequestExecutionResult) GetPeerCertificateDetailsOk() (*string, bool)`

GetPeerCertificateDetailsOk returns a tuple with the PeerCertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerCertificateDetails

`func (o *MonitorRequestExecutionResult) SetPeerCertificateDetails(v string)`

SetPeerCertificateDetails sets PeerCertificateDetails field to given value.

### HasPeerCertificateDetails

`func (o *MonitorRequestExecutionResult) HasPeerCertificateDetails() bool`

HasPeerCertificateDetails returns a boolean if a field has been set.

### GetTotalTime

`func (o *MonitorRequestExecutionResult) GetTotalTime() int64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *MonitorRequestExecutionResult) GetTotalTimeOk() (*int64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *MonitorRequestExecutionResult) SetTotalTime(v int64)`

SetTotalTime sets TotalTime field to given value.

### HasTotalTime

`func (o *MonitorRequestExecutionResult) HasTotalTime() bool`

HasTotalTime returns a boolean if a field has been set.

### GetHostNameResolutionTime

`func (o *MonitorRequestExecutionResult) GetHostNameResolutionTime() int64`

GetHostNameResolutionTime returns the HostNameResolutionTime field if non-nil, zero value otherwise.

### GetHostNameResolutionTimeOk

`func (o *MonitorRequestExecutionResult) GetHostNameResolutionTimeOk() (*int64, bool)`

GetHostNameResolutionTimeOk returns a tuple with the HostNameResolutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNameResolutionTime

`func (o *MonitorRequestExecutionResult) SetHostNameResolutionTime(v int64)`

SetHostNameResolutionTime sets HostNameResolutionTime field to given value.

### HasHostNameResolutionTime

`func (o *MonitorRequestExecutionResult) HasHostNameResolutionTime() bool`

HasHostNameResolutionTime returns a boolean if a field has been set.

### GetTcpConnectTime

`func (o *MonitorRequestExecutionResult) GetTcpConnectTime() int64`

GetTcpConnectTime returns the TcpConnectTime field if non-nil, zero value otherwise.

### GetTcpConnectTimeOk

`func (o *MonitorRequestExecutionResult) GetTcpConnectTimeOk() (*int64, bool)`

GetTcpConnectTimeOk returns a tuple with the TcpConnectTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpConnectTime

`func (o *MonitorRequestExecutionResult) SetTcpConnectTime(v int64)`

SetTcpConnectTime sets TcpConnectTime field to given value.

### HasTcpConnectTime

`func (o *MonitorRequestExecutionResult) HasTcpConnectTime() bool`

HasTcpConnectTime returns a boolean if a field has been set.

### GetTlsHandshakeTime

`func (o *MonitorRequestExecutionResult) GetTlsHandshakeTime() int64`

GetTlsHandshakeTime returns the TlsHandshakeTime field if non-nil, zero value otherwise.

### GetTlsHandshakeTimeOk

`func (o *MonitorRequestExecutionResult) GetTlsHandshakeTimeOk() (*int64, bool)`

GetTlsHandshakeTimeOk returns a tuple with the TlsHandshakeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsHandshakeTime

`func (o *MonitorRequestExecutionResult) SetTlsHandshakeTime(v int64)`

SetTlsHandshakeTime sets TlsHandshakeTime field to given value.

### HasTlsHandshakeTime

`func (o *MonitorRequestExecutionResult) HasTlsHandshakeTime() bool`

HasTlsHandshakeTime returns a boolean if a field has been set.

### GetTimeToFirstByte

`func (o *MonitorRequestExecutionResult) GetTimeToFirstByte() int64`

GetTimeToFirstByte returns the TimeToFirstByte field if non-nil, zero value otherwise.

### GetTimeToFirstByteOk

`func (o *MonitorRequestExecutionResult) GetTimeToFirstByteOk() (*int64, bool)`

GetTimeToFirstByteOk returns a tuple with the TimeToFirstByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToFirstByte

`func (o *MonitorRequestExecutionResult) SetTimeToFirstByte(v int64)`

SetTimeToFirstByte sets TimeToFirstByte field to given value.

### HasTimeToFirstByte

`func (o *MonitorRequestExecutionResult) HasTimeToFirstByte() bool`

HasTimeToFirstByte returns a boolean if a field has been set.

### GetRedirectionTime

`func (o *MonitorRequestExecutionResult) GetRedirectionTime() int64`

GetRedirectionTime returns the RedirectionTime field if non-nil, zero value otherwise.

### GetRedirectionTimeOk

`func (o *MonitorRequestExecutionResult) GetRedirectionTimeOk() (*int64, bool)`

GetRedirectionTimeOk returns a tuple with the RedirectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectionTime

`func (o *MonitorRequestExecutionResult) SetRedirectionTime(v int64)`

SetRedirectionTime sets RedirectionTime field to given value.

### HasRedirectionTime

`func (o *MonitorRequestExecutionResult) HasRedirectionTime() bool`

HasRedirectionTime returns a boolean if a field has been set.

### GetWaitingTime

`func (o *MonitorRequestExecutionResult) GetWaitingTime() int64`

GetWaitingTime returns the WaitingTime field if non-nil, zero value otherwise.

### GetWaitingTimeOk

`func (o *MonitorRequestExecutionResult) GetWaitingTimeOk() (*int64, bool)`

GetWaitingTimeOk returns a tuple with the WaitingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingTime

`func (o *MonitorRequestExecutionResult) SetWaitingTime(v int64)`

SetWaitingTime sets WaitingTime field to given value.

### HasWaitingTime

`func (o *MonitorRequestExecutionResult) HasWaitingTime() bool`

HasWaitingTime returns a boolean if a field has been set.

### GetRedirectsCount

`func (o *MonitorRequestExecutionResult) GetRedirectsCount() int32`

GetRedirectsCount returns the RedirectsCount field if non-nil, zero value otherwise.

### GetRedirectsCountOk

`func (o *MonitorRequestExecutionResult) GetRedirectsCountOk() (*int32, bool)`

GetRedirectsCountOk returns a tuple with the RedirectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectsCount

`func (o *MonitorRequestExecutionResult) SetRedirectsCount(v int32)`

SetRedirectsCount sets RedirectsCount field to given value.

### HasRedirectsCount

`func (o *MonitorRequestExecutionResult) HasRedirectsCount() bool`

HasRedirectsCount returns a boolean if a field has been set.

### GetFailureMessage

`func (o *MonitorRequestExecutionResult) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *MonitorRequestExecutionResult) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *MonitorRequestExecutionResult) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *MonitorRequestExecutionResult) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


