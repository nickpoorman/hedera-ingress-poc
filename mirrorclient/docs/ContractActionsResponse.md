# ContractActionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]ContractAction**](ContractAction.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewContractActionsResponse

`func NewContractActionsResponse() *ContractActionsResponse`

NewContractActionsResponse instantiates a new ContractActionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractActionsResponseWithDefaults

`func NewContractActionsResponseWithDefaults() *ContractActionsResponse`

NewContractActionsResponseWithDefaults instantiates a new ContractActionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ContractActionsResponse) GetActions() []ContractAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ContractActionsResponse) GetActionsOk() (*[]ContractAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ContractActionsResponse) SetActions(v []ContractAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ContractActionsResponse) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetLinks

`func (o *ContractActionsResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContractActionsResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContractActionsResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ContractActionsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


