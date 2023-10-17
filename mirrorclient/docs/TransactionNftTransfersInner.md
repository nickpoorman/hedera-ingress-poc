# TransactionNftTransfersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsApproval** | **bool** |  | 
**ReceiverAccountId** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**SenderAccountId** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**SerialNumber** | **int64** |  | 
**TokenId** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 

## Methods

### NewTransactionNftTransfersInner

`func NewTransactionNftTransfersInner(isApproval bool, receiverAccountId NullableString, senderAccountId NullableString, serialNumber int64, tokenId NullableString, ) *TransactionNftTransfersInner`

NewTransactionNftTransfersInner instantiates a new TransactionNftTransfersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionNftTransfersInnerWithDefaults

`func NewTransactionNftTransfersInnerWithDefaults() *TransactionNftTransfersInner`

NewTransactionNftTransfersInnerWithDefaults instantiates a new TransactionNftTransfersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsApproval

`func (o *TransactionNftTransfersInner) GetIsApproval() bool`

GetIsApproval returns the IsApproval field if non-nil, zero value otherwise.

### GetIsApprovalOk

`func (o *TransactionNftTransfersInner) GetIsApprovalOk() (*bool, bool)`

GetIsApprovalOk returns a tuple with the IsApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproval

`func (o *TransactionNftTransfersInner) SetIsApproval(v bool)`

SetIsApproval sets IsApproval field to given value.


### GetReceiverAccountId

`func (o *TransactionNftTransfersInner) GetReceiverAccountId() string`

GetReceiverAccountId returns the ReceiverAccountId field if non-nil, zero value otherwise.

### GetReceiverAccountIdOk

`func (o *TransactionNftTransfersInner) GetReceiverAccountIdOk() (*string, bool)`

GetReceiverAccountIdOk returns a tuple with the ReceiverAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverAccountId

`func (o *TransactionNftTransfersInner) SetReceiverAccountId(v string)`

SetReceiverAccountId sets ReceiverAccountId field to given value.


### SetReceiverAccountIdNil

`func (o *TransactionNftTransfersInner) SetReceiverAccountIdNil(b bool)`

 SetReceiverAccountIdNil sets the value for ReceiverAccountId to be an explicit nil

### UnsetReceiverAccountId
`func (o *TransactionNftTransfersInner) UnsetReceiverAccountId()`

UnsetReceiverAccountId ensures that no value is present for ReceiverAccountId, not even an explicit nil
### GetSenderAccountId

`func (o *TransactionNftTransfersInner) GetSenderAccountId() string`

GetSenderAccountId returns the SenderAccountId field if non-nil, zero value otherwise.

### GetSenderAccountIdOk

`func (o *TransactionNftTransfersInner) GetSenderAccountIdOk() (*string, bool)`

GetSenderAccountIdOk returns a tuple with the SenderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAccountId

`func (o *TransactionNftTransfersInner) SetSenderAccountId(v string)`

SetSenderAccountId sets SenderAccountId field to given value.


### SetSenderAccountIdNil

`func (o *TransactionNftTransfersInner) SetSenderAccountIdNil(b bool)`

 SetSenderAccountIdNil sets the value for SenderAccountId to be an explicit nil

### UnsetSenderAccountId
`func (o *TransactionNftTransfersInner) UnsetSenderAccountId()`

UnsetSenderAccountId ensures that no value is present for SenderAccountId, not even an explicit nil
### GetSerialNumber

`func (o *TransactionNftTransfersInner) GetSerialNumber() int64`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *TransactionNftTransfersInner) GetSerialNumberOk() (*int64, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *TransactionNftTransfersInner) SetSerialNumber(v int64)`

SetSerialNumber sets SerialNumber field to given value.


### GetTokenId

`func (o *TransactionNftTransfersInner) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionNftTransfersInner) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionNftTransfersInner) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### SetTokenIdNil

`func (o *TransactionNftTransfersInner) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *TransactionNftTransfersInner) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


