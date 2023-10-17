# AccountBalanceTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**Alias** | **NullableString** | RFC4648 no-padding base32 encoded account alias | 
**AutoRenewPeriod** | **NullableInt64** |  | 
**Balance** | [**NullableBalance**](Balance.md) |  | 
**CreatedTimestamp** | **NullableString** |  | 
**DeclineReward** | **bool** | Whether the account declines receiving a staking reward | 
**Deleted** | **NullableBool** |  | 
**EthereumNonce** | **NullableInt64** |  | 
**EvmAddress** | **Nullable*os.File** | A network entity encoded as an EVM address in hex. | 
**ExpiryTimestamp** | **NullableString** |  | 
**Key** | [**NullableKey**](Key.md) |  | 
**MaxAutomaticTokenAssociations** | **NullableInt32** |  | 
**Memo** | **NullableString** |  | 
**PendingReward** | Pointer to **int64** | The pending reward in tinybars the account will receive in the next reward payout. Note the value is updated at the end of each staking period and there may be delay to reflect the changes in the past staking period.  | [optional] 
**ReceiverSigRequired** | **NullableBool** |  | 
**StakedAccountId** | **string** |  | 
**StakedNodeId** | **NullableInt64** | The id of the node to which this account is staking | 
**StakePeriodStart** | **string** |  | 
**Transactions** | [**[]Transaction**](Transaction.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewAccountBalanceTransactions

`func NewAccountBalanceTransactions(account NullableString, alias NullableString, autoRenewPeriod NullableInt64, balance NullableBalance, createdTimestamp NullableString, declineReward bool, deleted NullableBool, ethereumNonce NullableInt64, evmAddress Nullable*os.File, expiryTimestamp NullableString, key NullableKey, maxAutomaticTokenAssociations NullableInt32, memo NullableString, receiverSigRequired NullableBool, stakedAccountId string, stakedNodeId NullableInt64, stakePeriodStart string, transactions []Transaction, links Links, ) *AccountBalanceTransactions`

NewAccountBalanceTransactions instantiates a new AccountBalanceTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountBalanceTransactionsWithDefaults

`func NewAccountBalanceTransactionsWithDefaults() *AccountBalanceTransactions`

NewAccountBalanceTransactionsWithDefaults instantiates a new AccountBalanceTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AccountBalanceTransactions) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccountBalanceTransactions) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccountBalanceTransactions) SetAccount(v string)`

SetAccount sets Account field to given value.


### SetAccountNil

`func (o *AccountBalanceTransactions) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *AccountBalanceTransactions) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetAlias

`func (o *AccountBalanceTransactions) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *AccountBalanceTransactions) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *AccountBalanceTransactions) SetAlias(v string)`

SetAlias sets Alias field to given value.


### SetAliasNil

`func (o *AccountBalanceTransactions) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *AccountBalanceTransactions) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetAutoRenewPeriod

`func (o *AccountBalanceTransactions) GetAutoRenewPeriod() int64`

GetAutoRenewPeriod returns the AutoRenewPeriod field if non-nil, zero value otherwise.

### GetAutoRenewPeriodOk

`func (o *AccountBalanceTransactions) GetAutoRenewPeriodOk() (*int64, bool)`

GetAutoRenewPeriodOk returns a tuple with the AutoRenewPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewPeriod

`func (o *AccountBalanceTransactions) SetAutoRenewPeriod(v int64)`

SetAutoRenewPeriod sets AutoRenewPeriod field to given value.


### SetAutoRenewPeriodNil

`func (o *AccountBalanceTransactions) SetAutoRenewPeriodNil(b bool)`

 SetAutoRenewPeriodNil sets the value for AutoRenewPeriod to be an explicit nil

### UnsetAutoRenewPeriod
`func (o *AccountBalanceTransactions) UnsetAutoRenewPeriod()`

UnsetAutoRenewPeriod ensures that no value is present for AutoRenewPeriod, not even an explicit nil
### GetBalance

`func (o *AccountBalanceTransactions) GetBalance() Balance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AccountBalanceTransactions) GetBalanceOk() (*Balance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AccountBalanceTransactions) SetBalance(v Balance)`

SetBalance sets Balance field to given value.


### SetBalanceNil

`func (o *AccountBalanceTransactions) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *AccountBalanceTransactions) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetCreatedTimestamp

`func (o *AccountBalanceTransactions) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *AccountBalanceTransactions) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *AccountBalanceTransactions) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### SetCreatedTimestampNil

`func (o *AccountBalanceTransactions) SetCreatedTimestampNil(b bool)`

 SetCreatedTimestampNil sets the value for CreatedTimestamp to be an explicit nil

### UnsetCreatedTimestamp
`func (o *AccountBalanceTransactions) UnsetCreatedTimestamp()`

UnsetCreatedTimestamp ensures that no value is present for CreatedTimestamp, not even an explicit nil
### GetDeclineReward

`func (o *AccountBalanceTransactions) GetDeclineReward() bool`

GetDeclineReward returns the DeclineReward field if non-nil, zero value otherwise.

### GetDeclineRewardOk

`func (o *AccountBalanceTransactions) GetDeclineRewardOk() (*bool, bool)`

GetDeclineRewardOk returns a tuple with the DeclineReward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReward

`func (o *AccountBalanceTransactions) SetDeclineReward(v bool)`

SetDeclineReward sets DeclineReward field to given value.


### GetDeleted

`func (o *AccountBalanceTransactions) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *AccountBalanceTransactions) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *AccountBalanceTransactions) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### SetDeletedNil

`func (o *AccountBalanceTransactions) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *AccountBalanceTransactions) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetEthereumNonce

`func (o *AccountBalanceTransactions) GetEthereumNonce() int64`

GetEthereumNonce returns the EthereumNonce field if non-nil, zero value otherwise.

### GetEthereumNonceOk

`func (o *AccountBalanceTransactions) GetEthereumNonceOk() (*int64, bool)`

GetEthereumNonceOk returns a tuple with the EthereumNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthereumNonce

`func (o *AccountBalanceTransactions) SetEthereumNonce(v int64)`

SetEthereumNonce sets EthereumNonce field to given value.


### SetEthereumNonceNil

`func (o *AccountBalanceTransactions) SetEthereumNonceNil(b bool)`

 SetEthereumNonceNil sets the value for EthereumNonce to be an explicit nil

### UnsetEthereumNonce
`func (o *AccountBalanceTransactions) UnsetEthereumNonce()`

UnsetEthereumNonce ensures that no value is present for EthereumNonce, not even an explicit nil
### GetEvmAddress

`func (o *AccountBalanceTransactions) GetEvmAddress() *os.File`

GetEvmAddress returns the EvmAddress field if non-nil, zero value otherwise.

### GetEvmAddressOk

`func (o *AccountBalanceTransactions) GetEvmAddressOk() (**os.File, bool)`

GetEvmAddressOk returns a tuple with the EvmAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvmAddress

`func (o *AccountBalanceTransactions) SetEvmAddress(v *os.File)`

SetEvmAddress sets EvmAddress field to given value.


### SetEvmAddressNil

`func (o *AccountBalanceTransactions) SetEvmAddressNil(b bool)`

 SetEvmAddressNil sets the value for EvmAddress to be an explicit nil

### UnsetEvmAddress
`func (o *AccountBalanceTransactions) UnsetEvmAddress()`

UnsetEvmAddress ensures that no value is present for EvmAddress, not even an explicit nil
### GetExpiryTimestamp

`func (o *AccountBalanceTransactions) GetExpiryTimestamp() string`

GetExpiryTimestamp returns the ExpiryTimestamp field if non-nil, zero value otherwise.

### GetExpiryTimestampOk

`func (o *AccountBalanceTransactions) GetExpiryTimestampOk() (*string, bool)`

GetExpiryTimestampOk returns a tuple with the ExpiryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimestamp

`func (o *AccountBalanceTransactions) SetExpiryTimestamp(v string)`

SetExpiryTimestamp sets ExpiryTimestamp field to given value.


### SetExpiryTimestampNil

`func (o *AccountBalanceTransactions) SetExpiryTimestampNil(b bool)`

 SetExpiryTimestampNil sets the value for ExpiryTimestamp to be an explicit nil

### UnsetExpiryTimestamp
`func (o *AccountBalanceTransactions) UnsetExpiryTimestamp()`

UnsetExpiryTimestamp ensures that no value is present for ExpiryTimestamp, not even an explicit nil
### GetKey

`func (o *AccountBalanceTransactions) GetKey() Key`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AccountBalanceTransactions) GetKeyOk() (*Key, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AccountBalanceTransactions) SetKey(v Key)`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *AccountBalanceTransactions) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *AccountBalanceTransactions) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetMaxAutomaticTokenAssociations

`func (o *AccountBalanceTransactions) GetMaxAutomaticTokenAssociations() int32`

GetMaxAutomaticTokenAssociations returns the MaxAutomaticTokenAssociations field if non-nil, zero value otherwise.

### GetMaxAutomaticTokenAssociationsOk

`func (o *AccountBalanceTransactions) GetMaxAutomaticTokenAssociationsOk() (*int32, bool)`

GetMaxAutomaticTokenAssociationsOk returns a tuple with the MaxAutomaticTokenAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAutomaticTokenAssociations

`func (o *AccountBalanceTransactions) SetMaxAutomaticTokenAssociations(v int32)`

SetMaxAutomaticTokenAssociations sets MaxAutomaticTokenAssociations field to given value.


### SetMaxAutomaticTokenAssociationsNil

`func (o *AccountBalanceTransactions) SetMaxAutomaticTokenAssociationsNil(b bool)`

 SetMaxAutomaticTokenAssociationsNil sets the value for MaxAutomaticTokenAssociations to be an explicit nil

### UnsetMaxAutomaticTokenAssociations
`func (o *AccountBalanceTransactions) UnsetMaxAutomaticTokenAssociations()`

UnsetMaxAutomaticTokenAssociations ensures that no value is present for MaxAutomaticTokenAssociations, not even an explicit nil
### GetMemo

`func (o *AccountBalanceTransactions) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *AccountBalanceTransactions) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *AccountBalanceTransactions) SetMemo(v string)`

SetMemo sets Memo field to given value.


### SetMemoNil

`func (o *AccountBalanceTransactions) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *AccountBalanceTransactions) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetPendingReward

`func (o *AccountBalanceTransactions) GetPendingReward() int64`

GetPendingReward returns the PendingReward field if non-nil, zero value otherwise.

### GetPendingRewardOk

`func (o *AccountBalanceTransactions) GetPendingRewardOk() (*int64, bool)`

GetPendingRewardOk returns a tuple with the PendingReward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingReward

`func (o *AccountBalanceTransactions) SetPendingReward(v int64)`

SetPendingReward sets PendingReward field to given value.

### HasPendingReward

`func (o *AccountBalanceTransactions) HasPendingReward() bool`

HasPendingReward returns a boolean if a field has been set.

### GetReceiverSigRequired

`func (o *AccountBalanceTransactions) GetReceiverSigRequired() bool`

GetReceiverSigRequired returns the ReceiverSigRequired field if non-nil, zero value otherwise.

### GetReceiverSigRequiredOk

`func (o *AccountBalanceTransactions) GetReceiverSigRequiredOk() (*bool, bool)`

GetReceiverSigRequiredOk returns a tuple with the ReceiverSigRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverSigRequired

`func (o *AccountBalanceTransactions) SetReceiverSigRequired(v bool)`

SetReceiverSigRequired sets ReceiverSigRequired field to given value.


### SetReceiverSigRequiredNil

`func (o *AccountBalanceTransactions) SetReceiverSigRequiredNil(b bool)`

 SetReceiverSigRequiredNil sets the value for ReceiverSigRequired to be an explicit nil

### UnsetReceiverSigRequired
`func (o *AccountBalanceTransactions) UnsetReceiverSigRequired()`

UnsetReceiverSigRequired ensures that no value is present for ReceiverSigRequired, not even an explicit nil
### GetStakedAccountId

`func (o *AccountBalanceTransactions) GetStakedAccountId() string`

GetStakedAccountId returns the StakedAccountId field if non-nil, zero value otherwise.

### GetStakedAccountIdOk

`func (o *AccountBalanceTransactions) GetStakedAccountIdOk() (*string, bool)`

GetStakedAccountIdOk returns a tuple with the StakedAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakedAccountId

`func (o *AccountBalanceTransactions) SetStakedAccountId(v string)`

SetStakedAccountId sets StakedAccountId field to given value.


### GetStakedNodeId

`func (o *AccountBalanceTransactions) GetStakedNodeId() int64`

GetStakedNodeId returns the StakedNodeId field if non-nil, zero value otherwise.

### GetStakedNodeIdOk

`func (o *AccountBalanceTransactions) GetStakedNodeIdOk() (*int64, bool)`

GetStakedNodeIdOk returns a tuple with the StakedNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakedNodeId

`func (o *AccountBalanceTransactions) SetStakedNodeId(v int64)`

SetStakedNodeId sets StakedNodeId field to given value.


### SetStakedNodeIdNil

`func (o *AccountBalanceTransactions) SetStakedNodeIdNil(b bool)`

 SetStakedNodeIdNil sets the value for StakedNodeId to be an explicit nil

### UnsetStakedNodeId
`func (o *AccountBalanceTransactions) UnsetStakedNodeId()`

UnsetStakedNodeId ensures that no value is present for StakedNodeId, not even an explicit nil
### GetStakePeriodStart

`func (o *AccountBalanceTransactions) GetStakePeriodStart() string`

GetStakePeriodStart returns the StakePeriodStart field if non-nil, zero value otherwise.

### GetStakePeriodStartOk

`func (o *AccountBalanceTransactions) GetStakePeriodStartOk() (*string, bool)`

GetStakePeriodStartOk returns a tuple with the StakePeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakePeriodStart

`func (o *AccountBalanceTransactions) SetStakePeriodStart(v string)`

SetStakePeriodStart sets StakePeriodStart field to given value.


### GetTransactions

`func (o *AccountBalanceTransactions) GetTransactions() []Transaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *AccountBalanceTransactions) GetTransactionsOk() (*[]Transaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *AccountBalanceTransactions) SetTransactions(v []Transaction)`

SetTransactions sets Transactions field to given value.


### GetLinks

`func (o *AccountBalanceTransactions) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AccountBalanceTransactions) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AccountBalanceTransactions) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


