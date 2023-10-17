# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **NullableString** |  | [optional] 
**ChargedTxFee** | Pointer to **int64** |  | [optional] 
**ConsensusTimestamp** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**MaxFee** | Pointer to **string** |  | [optional] 
**MemoBase64** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to [**TransactionTypes**](TransactionTypes.md) |  | [optional] 
**NftTransfers** | Pointer to [**[]TransactionNftTransfersInner**](TransactionNftTransfersInner.md) |  | [optional] 
**Node** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Nonce** | Pointer to **int32** |  | [optional] 
**ParentConsensusTimestamp** | Pointer to **NullableString** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**Scheduled** | Pointer to **bool** |  | [optional] 
**StakingRewardTransfers** | Pointer to [**[]StakingRewardTransfer**](StakingRewardTransfer.md) |  | [optional] 
**TokenTransfers** | Pointer to [**[]TransactionTokenTransfersInner**](TransactionTokenTransfersInner.md) |  | [optional] 
**TransactionHash** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**Transfers** | Pointer to [**[]TransactionTransfersInner**](TransactionTransfersInner.md) |  | [optional] 
**ValidDurationSeconds** | Pointer to **string** |  | [optional] 
**ValidStartTimestamp** | Pointer to **string** |  | [optional] 

## Methods

### NewTransaction

`func NewTransaction() *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *Transaction) GetBytes() string`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *Transaction) GetBytesOk() (*string, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *Transaction) SetBytes(v string)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *Transaction) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### SetBytesNil

`func (o *Transaction) SetBytesNil(b bool)`

 SetBytesNil sets the value for Bytes to be an explicit nil

### UnsetBytes
`func (o *Transaction) UnsetBytes()`

UnsetBytes ensures that no value is present for Bytes, not even an explicit nil
### GetChargedTxFee

`func (o *Transaction) GetChargedTxFee() int64`

GetChargedTxFee returns the ChargedTxFee field if non-nil, zero value otherwise.

### GetChargedTxFeeOk

`func (o *Transaction) GetChargedTxFeeOk() (*int64, bool)`

GetChargedTxFeeOk returns a tuple with the ChargedTxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargedTxFee

`func (o *Transaction) SetChargedTxFee(v int64)`

SetChargedTxFee sets ChargedTxFee field to given value.

### HasChargedTxFee

`func (o *Transaction) HasChargedTxFee() bool`

HasChargedTxFee returns a boolean if a field has been set.

### GetConsensusTimestamp

`func (o *Transaction) GetConsensusTimestamp() string`

GetConsensusTimestamp returns the ConsensusTimestamp field if non-nil, zero value otherwise.

### GetConsensusTimestampOk

`func (o *Transaction) GetConsensusTimestampOk() (*string, bool)`

GetConsensusTimestampOk returns a tuple with the ConsensusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsensusTimestamp

`func (o *Transaction) SetConsensusTimestamp(v string)`

SetConsensusTimestamp sets ConsensusTimestamp field to given value.

### HasConsensusTimestamp

`func (o *Transaction) HasConsensusTimestamp() bool`

HasConsensusTimestamp returns a boolean if a field has been set.

### GetEntityId

`func (o *Transaction) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Transaction) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Transaction) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *Transaction) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *Transaction) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *Transaction) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetMaxFee

`func (o *Transaction) GetMaxFee() string`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *Transaction) GetMaxFeeOk() (*string, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *Transaction) SetMaxFee(v string)`

SetMaxFee sets MaxFee field to given value.

### HasMaxFee

`func (o *Transaction) HasMaxFee() bool`

HasMaxFee returns a boolean if a field has been set.

### GetMemoBase64

`func (o *Transaction) GetMemoBase64() string`

GetMemoBase64 returns the MemoBase64 field if non-nil, zero value otherwise.

### GetMemoBase64Ok

`func (o *Transaction) GetMemoBase64Ok() (*string, bool)`

GetMemoBase64Ok returns a tuple with the MemoBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoBase64

`func (o *Transaction) SetMemoBase64(v string)`

SetMemoBase64 sets MemoBase64 field to given value.

### HasMemoBase64

`func (o *Transaction) HasMemoBase64() bool`

HasMemoBase64 returns a boolean if a field has been set.

### SetMemoBase64Nil

`func (o *Transaction) SetMemoBase64Nil(b bool)`

 SetMemoBase64Nil sets the value for MemoBase64 to be an explicit nil

### UnsetMemoBase64
`func (o *Transaction) UnsetMemoBase64()`

UnsetMemoBase64 ensures that no value is present for MemoBase64, not even an explicit nil
### GetName

`func (o *Transaction) GetName() TransactionTypes`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Transaction) GetNameOk() (*TransactionTypes, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Transaction) SetName(v TransactionTypes)`

SetName sets Name field to given value.

### HasName

`func (o *Transaction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNftTransfers

`func (o *Transaction) GetNftTransfers() []TransactionNftTransfersInner`

GetNftTransfers returns the NftTransfers field if non-nil, zero value otherwise.

### GetNftTransfersOk

`func (o *Transaction) GetNftTransfersOk() (*[]TransactionNftTransfersInner, bool)`

GetNftTransfersOk returns a tuple with the NftTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftTransfers

`func (o *Transaction) SetNftTransfers(v []TransactionNftTransfersInner)`

SetNftTransfers sets NftTransfers field to given value.

### HasNftTransfers

`func (o *Transaction) HasNftTransfers() bool`

HasNftTransfers returns a boolean if a field has been set.

### GetNode

`func (o *Transaction) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *Transaction) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *Transaction) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *Transaction) HasNode() bool`

HasNode returns a boolean if a field has been set.

### SetNodeNil

`func (o *Transaction) SetNodeNil(b bool)`

 SetNodeNil sets the value for Node to be an explicit nil

### UnsetNode
`func (o *Transaction) UnsetNode()`

UnsetNode ensures that no value is present for Node, not even an explicit nil
### GetNonce

`func (o *Transaction) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *Transaction) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *Transaction) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *Transaction) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetParentConsensusTimestamp

`func (o *Transaction) GetParentConsensusTimestamp() string`

GetParentConsensusTimestamp returns the ParentConsensusTimestamp field if non-nil, zero value otherwise.

### GetParentConsensusTimestampOk

`func (o *Transaction) GetParentConsensusTimestampOk() (*string, bool)`

GetParentConsensusTimestampOk returns a tuple with the ParentConsensusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConsensusTimestamp

`func (o *Transaction) SetParentConsensusTimestamp(v string)`

SetParentConsensusTimestamp sets ParentConsensusTimestamp field to given value.

### HasParentConsensusTimestamp

`func (o *Transaction) HasParentConsensusTimestamp() bool`

HasParentConsensusTimestamp returns a boolean if a field has been set.

### SetParentConsensusTimestampNil

`func (o *Transaction) SetParentConsensusTimestampNil(b bool)`

 SetParentConsensusTimestampNil sets the value for ParentConsensusTimestamp to be an explicit nil

### UnsetParentConsensusTimestamp
`func (o *Transaction) UnsetParentConsensusTimestamp()`

UnsetParentConsensusTimestamp ensures that no value is present for ParentConsensusTimestamp, not even an explicit nil
### GetResult

`func (o *Transaction) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Transaction) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Transaction) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *Transaction) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetScheduled

`func (o *Transaction) GetScheduled() bool`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *Transaction) GetScheduledOk() (*bool, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *Transaction) SetScheduled(v bool)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *Transaction) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.

### GetStakingRewardTransfers

`func (o *Transaction) GetStakingRewardTransfers() []StakingRewardTransfer`

GetStakingRewardTransfers returns the StakingRewardTransfers field if non-nil, zero value otherwise.

### GetStakingRewardTransfersOk

`func (o *Transaction) GetStakingRewardTransfersOk() (*[]StakingRewardTransfer, bool)`

GetStakingRewardTransfersOk returns a tuple with the StakingRewardTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingRewardTransfers

`func (o *Transaction) SetStakingRewardTransfers(v []StakingRewardTransfer)`

SetStakingRewardTransfers sets StakingRewardTransfers field to given value.

### HasStakingRewardTransfers

`func (o *Transaction) HasStakingRewardTransfers() bool`

HasStakingRewardTransfers returns a boolean if a field has been set.

### GetTokenTransfers

`func (o *Transaction) GetTokenTransfers() []TransactionTokenTransfersInner`

GetTokenTransfers returns the TokenTransfers field if non-nil, zero value otherwise.

### GetTokenTransfersOk

`func (o *Transaction) GetTokenTransfersOk() (*[]TransactionTokenTransfersInner, bool)`

GetTokenTransfersOk returns a tuple with the TokenTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTransfers

`func (o *Transaction) SetTokenTransfers(v []TransactionTokenTransfersInner)`

SetTokenTransfers sets TokenTransfers field to given value.

### HasTokenTransfers

`func (o *Transaction) HasTokenTransfers() bool`

HasTokenTransfers returns a boolean if a field has been set.

### GetTransactionHash

`func (o *Transaction) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *Transaction) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *Transaction) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *Transaction) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### GetTransactionId

`func (o *Transaction) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Transaction) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Transaction) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *Transaction) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransfers

`func (o *Transaction) GetTransfers() []TransactionTransfersInner`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *Transaction) GetTransfersOk() (*[]TransactionTransfersInner, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *Transaction) SetTransfers(v []TransactionTransfersInner)`

SetTransfers sets Transfers field to given value.

### HasTransfers

`func (o *Transaction) HasTransfers() bool`

HasTransfers returns a boolean if a field has been set.

### GetValidDurationSeconds

`func (o *Transaction) GetValidDurationSeconds() string`

GetValidDurationSeconds returns the ValidDurationSeconds field if non-nil, zero value otherwise.

### GetValidDurationSecondsOk

`func (o *Transaction) GetValidDurationSecondsOk() (*string, bool)`

GetValidDurationSecondsOk returns a tuple with the ValidDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDurationSeconds

`func (o *Transaction) SetValidDurationSeconds(v string)`

SetValidDurationSeconds sets ValidDurationSeconds field to given value.

### HasValidDurationSeconds

`func (o *Transaction) HasValidDurationSeconds() bool`

HasValidDurationSeconds returns a boolean if a field has been set.

### GetValidStartTimestamp

`func (o *Transaction) GetValidStartTimestamp() string`

GetValidStartTimestamp returns the ValidStartTimestamp field if non-nil, zero value otherwise.

### GetValidStartTimestampOk

`func (o *Transaction) GetValidStartTimestampOk() (*string, bool)`

GetValidStartTimestampOk returns a tuple with the ValidStartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidStartTimestamp

`func (o *Transaction) SetValidStartTimestamp(v string)`

SetValidStartTimestamp sets ValidStartTimestamp field to given value.

### HasValidStartTimestamp

`func (o *Transaction) HasValidStartTimestamp() bool`

HasValidStartTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


