# NetworkFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gas** | Pointer to **int64** | gas cost in tinybars | [optional] 
**TransactionType** | Pointer to **string** | type of the transaction | [optional] 

## Methods

### NewNetworkFee

`func NewNetworkFee() *NetworkFee`

NewNetworkFee instantiates a new NetworkFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFeeWithDefaults

`func NewNetworkFeeWithDefaults() *NetworkFee`

NewNetworkFeeWithDefaults instantiates a new NetworkFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGas

`func (o *NetworkFee) GetGas() int64`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *NetworkFee) GetGasOk() (*int64, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *NetworkFee) SetGas(v int64)`

SetGas sets Gas field to given value.

### HasGas

`func (o *NetworkFee) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetTransactionType

`func (o *NetworkFee) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *NetworkFee) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *NetworkFee) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *NetworkFee) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


