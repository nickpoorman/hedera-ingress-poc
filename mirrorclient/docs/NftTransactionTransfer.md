# NftTransactionTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsensusTimestamp** | **string** |  | 
**IsApproval** | **bool** |  | 
**Nonce** | **int32** |  | 
**ReceiverAccountId** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**SenderAccountId** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**TransactionId** | **string** |  | 
**Type** | [**TransactionTypes**](TransactionTypes.md) |  | 

## Methods

### NewNftTransactionTransfer

`func NewNftTransactionTransfer(consensusTimestamp string, isApproval bool, nonce int32, receiverAccountId NullableString, senderAccountId NullableString, transactionId string, type_ TransactionTypes, ) *NftTransactionTransfer`

NewNftTransactionTransfer instantiates a new NftTransactionTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftTransactionTransferWithDefaults

`func NewNftTransactionTransferWithDefaults() *NftTransactionTransfer`

NewNftTransactionTransferWithDefaults instantiates a new NftTransactionTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsensusTimestamp

`func (o *NftTransactionTransfer) GetConsensusTimestamp() string`

GetConsensusTimestamp returns the ConsensusTimestamp field if non-nil, zero value otherwise.

### GetConsensusTimestampOk

`func (o *NftTransactionTransfer) GetConsensusTimestampOk() (*string, bool)`

GetConsensusTimestampOk returns a tuple with the ConsensusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsensusTimestamp

`func (o *NftTransactionTransfer) SetConsensusTimestamp(v string)`

SetConsensusTimestamp sets ConsensusTimestamp field to given value.


### GetIsApproval

`func (o *NftTransactionTransfer) GetIsApproval() bool`

GetIsApproval returns the IsApproval field if non-nil, zero value otherwise.

### GetIsApprovalOk

`func (o *NftTransactionTransfer) GetIsApprovalOk() (*bool, bool)`

GetIsApprovalOk returns a tuple with the IsApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproval

`func (o *NftTransactionTransfer) SetIsApproval(v bool)`

SetIsApproval sets IsApproval field to given value.


### GetNonce

`func (o *NftTransactionTransfer) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *NftTransactionTransfer) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *NftTransactionTransfer) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetReceiverAccountId

`func (o *NftTransactionTransfer) GetReceiverAccountId() string`

GetReceiverAccountId returns the ReceiverAccountId field if non-nil, zero value otherwise.

### GetReceiverAccountIdOk

`func (o *NftTransactionTransfer) GetReceiverAccountIdOk() (*string, bool)`

GetReceiverAccountIdOk returns a tuple with the ReceiverAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverAccountId

`func (o *NftTransactionTransfer) SetReceiverAccountId(v string)`

SetReceiverAccountId sets ReceiverAccountId field to given value.


### SetReceiverAccountIdNil

`func (o *NftTransactionTransfer) SetReceiverAccountIdNil(b bool)`

 SetReceiverAccountIdNil sets the value for ReceiverAccountId to be an explicit nil

### UnsetReceiverAccountId
`func (o *NftTransactionTransfer) UnsetReceiverAccountId()`

UnsetReceiverAccountId ensures that no value is present for ReceiverAccountId, not even an explicit nil
### GetSenderAccountId

`func (o *NftTransactionTransfer) GetSenderAccountId() string`

GetSenderAccountId returns the SenderAccountId field if non-nil, zero value otherwise.

### GetSenderAccountIdOk

`func (o *NftTransactionTransfer) GetSenderAccountIdOk() (*string, bool)`

GetSenderAccountIdOk returns a tuple with the SenderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAccountId

`func (o *NftTransactionTransfer) SetSenderAccountId(v string)`

SetSenderAccountId sets SenderAccountId field to given value.


### SetSenderAccountIdNil

`func (o *NftTransactionTransfer) SetSenderAccountIdNil(b bool)`

 SetSenderAccountIdNil sets the value for SenderAccountId to be an explicit nil

### UnsetSenderAccountId
`func (o *NftTransactionTransfer) UnsetSenderAccountId()`

UnsetSenderAccountId ensures that no value is present for SenderAccountId, not even an explicit nil
### GetTransactionId

`func (o *NftTransactionTransfer) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *NftTransactionTransfer) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *NftTransactionTransfer) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetType

`func (o *NftTransactionTransfer) GetType() TransactionTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NftTransactionTransfer) GetTypeOk() (*TransactionTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NftTransactionTransfer) SetType(v TransactionTypes)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


