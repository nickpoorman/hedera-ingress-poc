# NftTransactionHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | [**[]NftTransactionTransfer**](NftTransactionTransfer.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewNftTransactionHistory

`func NewNftTransactionHistory(transactions []NftTransactionTransfer, links Links, ) *NftTransactionHistory`

NewNftTransactionHistory instantiates a new NftTransactionHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftTransactionHistoryWithDefaults

`func NewNftTransactionHistoryWithDefaults() *NftTransactionHistory`

NewNftTransactionHistoryWithDefaults instantiates a new NftTransactionHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *NftTransactionHistory) GetTransactions() []NftTransactionTransfer`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *NftTransactionHistory) GetTransactionsOk() (*[]NftTransactionTransfer, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *NftTransactionHistory) SetTransactions(v []NftTransactionTransfer)`

SetTransactions sets Transactions field to given value.


### GetLinks

`func (o *NftTransactionHistory) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NftTransactionHistory) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NftTransactionHistory) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


