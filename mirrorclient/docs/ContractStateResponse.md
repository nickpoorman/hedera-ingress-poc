# ContractStateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**[]ContractState**](ContractState.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewContractStateResponse

`func NewContractStateResponse() *ContractStateResponse`

NewContractStateResponse instantiates a new ContractStateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractStateResponseWithDefaults

`func NewContractStateResponseWithDefaults() *ContractStateResponse`

NewContractStateResponseWithDefaults instantiates a new ContractStateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ContractStateResponse) GetState() []ContractState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContractStateResponse) GetStateOk() (*[]ContractState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContractStateResponse) SetState(v []ContractState)`

SetState sets State field to given value.

### HasState

`func (o *ContractStateResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLinks

`func (o *ContractStateResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContractStateResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContractStateResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ContractStateResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


