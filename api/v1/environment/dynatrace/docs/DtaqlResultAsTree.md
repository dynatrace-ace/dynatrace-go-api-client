# DtaqlResultAsTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtrapolationLevel** | Pointer to **int32** | The extrapolation level of the result.   To improve performance, some results may be calculated from a subset of actual data. The extrapolation level indicates the share of actual data in the result.   The number is the denominator of a fraction and indicates the amount of actual data. The value &#x60;1&#x60; means that the result contains only the actual data. The value &#x60;4&#x60; means that result is calculated using 1/4 of the actual data.   If you need the analysis to be based on the actual data, reduce the timeframe of your query. For example, in case of extrapolation level of &#x60;4&#x60;, try to use 1/4 of the original timeframe. | [optional] 
**AdditionalColumnNames** | Pointer to **[]string** | A list of columns in the additionalValues table.    Only present if the endpoint was called with &#x60;deepLinkFields&#x3D;true&#x60; parameter. | [optional] 
**AdditionalValues** | Pointer to **[][]map[string]interface{}** | A list of data rows.    Each array element represents a row in the table of additionally linked fields.   The size of each data row and the order of the elements correspond to the **additionalColumnNames** content.   Only present if the endpoint was called with &#x60;deepLinkFields&#x3D;true&#x60; parameter. | [optional] 
**Explanations** | Pointer to **[]string** | Additional information about the query and the result, that helps to understand the query and how the result was calculated.   Only appears when the **explain** parameter is set to &#x60;true&#x60;.   Example: The number of results was limited to the default of 50. Use the &#x60;LIMIT&#x60; clause to increase or decrease this limit. | [optional] 
**BranchNames** | Pointer to **[]string** | A list of branches of the tree.    Typically, these are fields from the &#x60;SELECT&#x60; clause, that have been used in the &#x60;GROUP BY&#x60; clause. | [optional] 
**LeafNames** | Pointer to **[]string** | A list of leaves on each tree branch.    Typically, these are fields from the &#x60;SELECT&#x60; clause, that have not been used in the &#x60;GROUP BY&#x60; clause. | [optional] 
**Values** | Pointer to **map[string]interface{}** | The user session query result as a tree. | [optional] 

## Methods

### NewDtaqlResultAsTree

`func NewDtaqlResultAsTree() *DtaqlResultAsTree`

NewDtaqlResultAsTree instantiates a new DtaqlResultAsTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtaqlResultAsTreeWithDefaults

`func NewDtaqlResultAsTreeWithDefaults() *DtaqlResultAsTree`

NewDtaqlResultAsTreeWithDefaults instantiates a new DtaqlResultAsTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtrapolationLevel

`func (o *DtaqlResultAsTree) GetExtrapolationLevel() int32`

GetExtrapolationLevel returns the ExtrapolationLevel field if non-nil, zero value otherwise.

### GetExtrapolationLevelOk

`func (o *DtaqlResultAsTree) GetExtrapolationLevelOk() (*int32, bool)`

GetExtrapolationLevelOk returns a tuple with the ExtrapolationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtrapolationLevel

`func (o *DtaqlResultAsTree) SetExtrapolationLevel(v int32)`

SetExtrapolationLevel sets ExtrapolationLevel field to given value.

### HasExtrapolationLevel

`func (o *DtaqlResultAsTree) HasExtrapolationLevel() bool`

HasExtrapolationLevel returns a boolean if a field has been set.

### GetAdditionalColumnNames

`func (o *DtaqlResultAsTree) GetAdditionalColumnNames() []string`

GetAdditionalColumnNames returns the AdditionalColumnNames field if non-nil, zero value otherwise.

### GetAdditionalColumnNamesOk

`func (o *DtaqlResultAsTree) GetAdditionalColumnNamesOk() (*[]string, bool)`

GetAdditionalColumnNamesOk returns a tuple with the AdditionalColumnNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalColumnNames

`func (o *DtaqlResultAsTree) SetAdditionalColumnNames(v []string)`

SetAdditionalColumnNames sets AdditionalColumnNames field to given value.

### HasAdditionalColumnNames

`func (o *DtaqlResultAsTree) HasAdditionalColumnNames() bool`

HasAdditionalColumnNames returns a boolean if a field has been set.

### GetAdditionalValues

`func (o *DtaqlResultAsTree) GetAdditionalValues() [][]map[string]interface{}`

GetAdditionalValues returns the AdditionalValues field if non-nil, zero value otherwise.

### GetAdditionalValuesOk

`func (o *DtaqlResultAsTree) GetAdditionalValuesOk() (*[][]map[string]interface{}, bool)`

GetAdditionalValuesOk returns a tuple with the AdditionalValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalValues

`func (o *DtaqlResultAsTree) SetAdditionalValues(v [][]map[string]interface{})`

SetAdditionalValues sets AdditionalValues field to given value.

### HasAdditionalValues

`func (o *DtaqlResultAsTree) HasAdditionalValues() bool`

HasAdditionalValues returns a boolean if a field has been set.

### GetExplanations

`func (o *DtaqlResultAsTree) GetExplanations() []string`

GetExplanations returns the Explanations field if non-nil, zero value otherwise.

### GetExplanationsOk

`func (o *DtaqlResultAsTree) GetExplanationsOk() (*[]string, bool)`

GetExplanationsOk returns a tuple with the Explanations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanations

`func (o *DtaqlResultAsTree) SetExplanations(v []string)`

SetExplanations sets Explanations field to given value.

### HasExplanations

`func (o *DtaqlResultAsTree) HasExplanations() bool`

HasExplanations returns a boolean if a field has been set.

### GetBranchNames

`func (o *DtaqlResultAsTree) GetBranchNames() []string`

GetBranchNames returns the BranchNames field if non-nil, zero value otherwise.

### GetBranchNamesOk

`func (o *DtaqlResultAsTree) GetBranchNamesOk() (*[]string, bool)`

GetBranchNamesOk returns a tuple with the BranchNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchNames

`func (o *DtaqlResultAsTree) SetBranchNames(v []string)`

SetBranchNames sets BranchNames field to given value.

### HasBranchNames

`func (o *DtaqlResultAsTree) HasBranchNames() bool`

HasBranchNames returns a boolean if a field has been set.

### GetLeafNames

`func (o *DtaqlResultAsTree) GetLeafNames() []string`

GetLeafNames returns the LeafNames field if non-nil, zero value otherwise.

### GetLeafNamesOk

`func (o *DtaqlResultAsTree) GetLeafNamesOk() (*[]string, bool)`

GetLeafNamesOk returns a tuple with the LeafNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafNames

`func (o *DtaqlResultAsTree) SetLeafNames(v []string)`

SetLeafNames sets LeafNames field to given value.

### HasLeafNames

`func (o *DtaqlResultAsTree) HasLeafNames() bool`

HasLeafNames returns a boolean if a field has been set.

### GetValues

`func (o *DtaqlResultAsTree) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *DtaqlResultAsTree) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *DtaqlResultAsTree) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *DtaqlResultAsTree) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


