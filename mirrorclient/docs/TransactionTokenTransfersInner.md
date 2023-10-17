# TransactionTokenTransfersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**Account** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**Amount** | **int64** |  | 
**IsApproval** | Pointer to **bool** |  | [optional] 

## Methods

### NewTransactionTokenTransfersInner

`func NewTransactionTokenTransfersInner(tokenId NullableString, account NullableString, amount int64, ) *TransactionTokenTransfersInner`

NewTransactionTokenTransfersInner instantiates a new TransactionTokenTransfersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionTokenTransfersInnerWithDefaults

`func NewTransactionTokenTransfersInnerWithDefaults() *TransactionTokenTransfersInner`

NewTransactionTokenTransfersInnerWithDefaults instantiates a new TransactionTokenTransfersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *TransactionTokenTransfersInner) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionTokenTransfersInner) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionTokenTransfersInner) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### SetTokenIdNil

`func (o *TransactionTokenTransfersInner) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *TransactionTokenTransfersInner) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
### GetAccount

`func (o *TransactionTokenTransfersInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TransactionTokenTransfersInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TransactionTokenTransfersInner) SetAccount(v string)`

SetAccount sets Account field to given value.


### SetAccountNil

`func (o *TransactionTokenTransfersInner) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *TransactionTokenTransfersInner) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetAmount

`func (o *TransactionTokenTransfersInner) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionTokenTransfersInner) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionTokenTransfersInner) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetIsApproval

`func (o *TransactionTokenTransfersInner) GetIsApproval() bool`

GetIsApproval returns the IsApproval field if non-nil, zero value otherwise.

### GetIsApprovalOk

`func (o *TransactionTokenTransfersInner) GetIsApprovalOk() (*bool, bool)`

GetIsApprovalOk returns a tuple with the IsApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproval

`func (o *TransactionTokenTransfersInner) SetIsApproval(v bool)`

SetIsApproval sets IsApproval field to given value.

### HasIsApproval

`func (o *TransactionTokenTransfersInner) HasIsApproval() bool`

HasIsApproval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


