# TransactionByIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | Pointer to [**[]TransactionDetail**](TransactionDetail.md) |  | [optional] 

## Methods

### NewTransactionByIdResponse

`func NewTransactionByIdResponse() *TransactionByIdResponse`

NewTransactionByIdResponse instantiates a new TransactionByIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionByIdResponseWithDefaults

`func NewTransactionByIdResponseWithDefaults() *TransactionByIdResponse`

NewTransactionByIdResponseWithDefaults instantiates a new TransactionByIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *TransactionByIdResponse) GetTransactions() []TransactionDetail`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *TransactionByIdResponse) GetTransactionsOk() (*[]TransactionDetail, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *TransactionByIdResponse) SetTransactions(v []TransactionDetail)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *TransactionByIdResponse) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


