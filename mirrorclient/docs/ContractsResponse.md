# ContractsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contracts** | Pointer to [**[]Contract**](Contract.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewContractsResponse

`func NewContractsResponse() *ContractsResponse`

NewContractsResponse instantiates a new ContractsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractsResponseWithDefaults

`func NewContractsResponseWithDefaults() *ContractsResponse`

NewContractsResponseWithDefaults instantiates a new ContractsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContracts

`func (o *ContractsResponse) GetContracts() []Contract`

GetContracts returns the Contracts field if non-nil, zero value otherwise.

### GetContractsOk

`func (o *ContractsResponse) GetContractsOk() (*[]Contract, bool)`

GetContractsOk returns a tuple with the Contracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContracts

`func (o *ContractsResponse) SetContracts(v []Contract)`

SetContracts sets Contracts field to given value.

### HasContracts

`func (o *ContractsResponse) HasContracts() bool`

HasContracts returns a boolean if a field has been set.

### GetLinks

`func (o *ContractsResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContractsResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContractsResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ContractsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


