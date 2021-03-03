# PublicSyntheticLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudPlatform** | **string** | The cloud provider where the location is hosted. | 
**Ips** | **[]string** | The list of IP addresses assigned to the location. | 
**Stage** | **string** | The stage of the location. | 
**BrowserType** | **string** | The type of the browser the location is using to execute browser monitors. | 
**BrowserVersion** | **string** | The version of the browser the location is using to execute browser monitors. | 
**Capabilities** | Pointer to **[]string** | A list of location capabilities. | [optional] 

## Methods

### NewPublicSyntheticLocation

`func NewPublicSyntheticLocation(cloudPlatform string, ips []string, stage string, browserType string, browserVersion string, ) *PublicSyntheticLocation`

NewPublicSyntheticLocation instantiates a new PublicSyntheticLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSyntheticLocationWithDefaults

`func NewPublicSyntheticLocationWithDefaults() *PublicSyntheticLocation`

NewPublicSyntheticLocationWithDefaults instantiates a new PublicSyntheticLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudPlatform

`func (o *PublicSyntheticLocation) GetCloudPlatform() string`

GetCloudPlatform returns the CloudPlatform field if non-nil, zero value otherwise.

### GetCloudPlatformOk

`func (o *PublicSyntheticLocation) GetCloudPlatformOk() (*string, bool)`

GetCloudPlatformOk returns a tuple with the CloudPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPlatform

`func (o *PublicSyntheticLocation) SetCloudPlatform(v string)`

SetCloudPlatform sets CloudPlatform field to given value.


### GetIps

`func (o *PublicSyntheticLocation) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *PublicSyntheticLocation) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *PublicSyntheticLocation) SetIps(v []string)`

SetIps sets Ips field to given value.


### GetStage

`func (o *PublicSyntheticLocation) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *PublicSyntheticLocation) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *PublicSyntheticLocation) SetStage(v string)`

SetStage sets Stage field to given value.


### GetBrowserType

`func (o *PublicSyntheticLocation) GetBrowserType() string`

GetBrowserType returns the BrowserType field if non-nil, zero value otherwise.

### GetBrowserTypeOk

`func (o *PublicSyntheticLocation) GetBrowserTypeOk() (*string, bool)`

GetBrowserTypeOk returns a tuple with the BrowserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserType

`func (o *PublicSyntheticLocation) SetBrowserType(v string)`

SetBrowserType sets BrowserType field to given value.


### GetBrowserVersion

`func (o *PublicSyntheticLocation) GetBrowserVersion() string`

GetBrowserVersion returns the BrowserVersion field if non-nil, zero value otherwise.

### GetBrowserVersionOk

`func (o *PublicSyntheticLocation) GetBrowserVersionOk() (*string, bool)`

GetBrowserVersionOk returns a tuple with the BrowserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserVersion

`func (o *PublicSyntheticLocation) SetBrowserVersion(v string)`

SetBrowserVersion sets BrowserVersion field to given value.


### GetCapabilities

`func (o *PublicSyntheticLocation) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *PublicSyntheticLocation) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *PublicSyntheticLocation) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *PublicSyntheticLocation) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


