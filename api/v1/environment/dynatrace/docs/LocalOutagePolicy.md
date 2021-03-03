# LocalOutagePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedLocations** | **int32** | The number of affected locations to trigger an alert. | 
**ConsecutiveRuns** | **int32** | The number of consecutive fails to trigger an alert. | 

## Methods

### NewLocalOutagePolicy

`func NewLocalOutagePolicy(affectedLocations int32, consecutiveRuns int32, ) *LocalOutagePolicy`

NewLocalOutagePolicy instantiates a new LocalOutagePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalOutagePolicyWithDefaults

`func NewLocalOutagePolicyWithDefaults() *LocalOutagePolicy`

NewLocalOutagePolicyWithDefaults instantiates a new LocalOutagePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedLocations

`func (o *LocalOutagePolicy) GetAffectedLocations() int32`

GetAffectedLocations returns the AffectedLocations field if non-nil, zero value otherwise.

### GetAffectedLocationsOk

`func (o *LocalOutagePolicy) GetAffectedLocationsOk() (*int32, bool)`

GetAffectedLocationsOk returns a tuple with the AffectedLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedLocations

`func (o *LocalOutagePolicy) SetAffectedLocations(v int32)`

SetAffectedLocations sets AffectedLocations field to given value.


### GetConsecutiveRuns

`func (o *LocalOutagePolicy) GetConsecutiveRuns() int32`

GetConsecutiveRuns returns the ConsecutiveRuns field if non-nil, zero value otherwise.

### GetConsecutiveRunsOk

`func (o *LocalOutagePolicy) GetConsecutiveRunsOk() (*int32, bool)`

GetConsecutiveRunsOk returns a tuple with the ConsecutiveRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveRuns

`func (o *LocalOutagePolicy) SetConsecutiveRuns(v int32)`

SetConsecutiveRuns sets ConsecutiveRuns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


