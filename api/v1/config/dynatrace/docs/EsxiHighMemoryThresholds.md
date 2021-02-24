# EsxiHighMemoryThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompressionDecompressionRate** | **float32** | Alert if ESXi host swap IN/OUT or compression/decompression rate is higher *X* kilobytes per second in 3 out of 5 samples. | 

## Methods

### NewEsxiHighMemoryThresholds

`func NewEsxiHighMemoryThresholds(compressionDecompressionRate float32, ) *EsxiHighMemoryThresholds`

NewEsxiHighMemoryThresholds instantiates a new EsxiHighMemoryThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsxiHighMemoryThresholdsWithDefaults

`func NewEsxiHighMemoryThresholdsWithDefaults() *EsxiHighMemoryThresholds`

NewEsxiHighMemoryThresholdsWithDefaults instantiates a new EsxiHighMemoryThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompressionDecompressionRate

`func (o *EsxiHighMemoryThresholds) GetCompressionDecompressionRate() float32`

GetCompressionDecompressionRate returns the CompressionDecompressionRate field if non-nil, zero value otherwise.

### GetCompressionDecompressionRateOk

`func (o *EsxiHighMemoryThresholds) GetCompressionDecompressionRateOk() (*float32, bool)`

GetCompressionDecompressionRateOk returns a tuple with the CompressionDecompressionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionDecompressionRate

`func (o *EsxiHighMemoryThresholds) SetCompressionDecompressionRate(v float32)`

SetCompressionDecompressionRate sets CompressionDecompressionRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


