# ContractResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]ContractResult**](ContractResult.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewContractResultsResponse

`func NewContractResultsResponse() *ContractResultsResponse`

NewContractResultsResponse instantiates a new ContractResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractResultsResponseWithDefaults

`func NewContractResultsResponseWithDefaults() *ContractResultsResponse`

NewContractResultsResponseWithDefaults instantiates a new ContractResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ContractResultsResponse) GetResults() []ContractResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ContractResultsResponse) GetResultsOk() (*[]ContractResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ContractResultsResponse) SetResults(v []ContractResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *ContractResultsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetLinks

`func (o *ContractResultsResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContractResultsResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContractResultsResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ContractResultsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


