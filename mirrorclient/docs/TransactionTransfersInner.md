# TransactionTransfersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**Amount** | **int64** |  | 
**IsApproval** | Pointer to **bool** |  | [optional] 

## Methods

### NewTransactionTransfersInner

`func NewTransactionTransfersInner(account NullableString, amount int64, ) *TransactionTransfersInner`

NewTransactionTransfersInner instantiates a new TransactionTransfersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionTransfersInnerWithDefaults

`func NewTransactionTransfersInnerWithDefaults() *TransactionTransfersInner`

NewTransactionTransfersInnerWithDefaults instantiates a new TransactionTransfersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *TransactionTransfersInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TransactionTransfersInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TransactionTransfersInner) SetAccount(v string)`

SetAccount sets Account field to given value.


### SetAccountNil

`func (o *TransactionTransfersInner) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *TransactionTransfersInner) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetAmount

`func (o *TransactionTransfersInner) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionTransfersInner) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionTransfersInner) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetIsApproval

`func (o *TransactionTransfersInner) GetIsApproval() bool`

GetIsApproval returns the IsApproval field if non-nil, zero value otherwise.

### GetIsApprovalOk

`func (o *TransactionTransfersInner) GetIsApprovalOk() (*bool, bool)`

GetIsApprovalOk returns a tuple with the IsApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproval

`func (o *TransactionTransfersInner) SetIsApproval(v bool)`

SetIsApproval sets IsApproval field to given value.

### HasIsApproval

`func (o *TransactionTransfersInner) HasIsApproval() bool`

HasIsApproval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


